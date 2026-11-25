// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sdk

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/droidrun/mobilerun-sdk-go/v6"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/tidwall/gjson"
)

// clientRef carries the raw client and the options it was built with. The options
// are needed for presigned transfers, which must not send API credentials.
type clientRef struct {
	api     *mobilerun.Client
	options []option.RequestOption
}

// objectControl is a request option understood only by the object layer. Raw
// client methods ignore it.
type objectControl func(*objectControls)

func (objectControl) Apply(*requestconfig.RequestConfig) error { return nil }

type objectControls struct {
	noWaitReady        bool
	readyTimeout       time.Duration
	keepOnReadyFailure bool
	noWait             bool
	waitTimeout        time.Duration
	pollInterval       time.Duration
	err                error
}

// WithoutWaitReady returns the handle immediately after create (or after a method
// that waits for readiness) instead of waiting until the object is ready.
func WithoutWaitReady() option.RequestOption {
	return objectControl(func(c *objectControls) { c.noWaitReady = true })
}

// WithReadyTimeout overrides the declared readiness timeout.
func WithReadyTimeout(timeout time.Duration) option.RequestOption {
	return objectControl(func(c *objectControls) {
		if timeout <= 0 {
			c.err = errors.New("sdk: ready timeout must be positive")
		}
		c.readyTimeout = timeout
	})
}

// WithoutDisposeOnReadyFailure keeps a created object when its readiness wait
// fails. By default the object is disposed once, best-effort.
func WithoutDisposeOnReadyFailure() option.RequestOption {
	return objectControl(func(c *objectControls) { c.keepOnReadyFailure = true })
}

// WithoutWait returns the handle immediately after a method that runs a declared
// wait afterwards.
func WithoutWait() option.RequestOption {
	return objectControl(func(c *objectControls) { c.noWait = true })
}

// WithWaitTimeout overrides the declared timeout of a wait.
func WithWaitTimeout(timeout time.Duration) option.RequestOption {
	return objectControl(func(c *objectControls) {
		if timeout <= 0 {
			c.err = errors.New("sdk: wait timeout must be positive")
		}
		c.waitTimeout = timeout
	})
}

// WithPollInterval overrides the declared poll interval of a wait.
func WithPollInterval(interval time.Duration) option.RequestOption {
	return objectControl(func(c *objectControls) {
		if interval <= 0 {
			c.err = errors.New("sdk: poll interval must be positive")
		}
		c.pollInterval = interval
	})
}

func splitOptions(opts []option.RequestOption) (objectControls, []option.RequestOption, error) {
	var controls objectControls
	request := make([]option.RequestOption, 0, len(opts))
	for _, opt := range opts {
		if control, ok := opt.(objectControl); ok {
			control(&controls)
			continue
		}
		request = append(request, opt)
	}
	return controls, request, controls.err
}

func durationOr(value, fallback time.Duration) time.Duration {
	if value > 0 {
		return value
	}
	return fallback
}

// WaitTimeoutError reports that a wait passed its deadline. Cause is the last
// retried error, if any.
type WaitTimeoutError struct {
	Cause error
}

func (e *WaitTimeoutError) Error() string {
	if e.Cause != nil {
		return "sdk: wait timed out: " + e.Cause.Error()
	}
	return "sdk: wait timed out"
}

func (e *WaitTimeoutError) Unwrap() error { return e.Cause }

// WaitFailureError reports that a wait observed a failure (or unlisted) state.
// Response is the last poll response.
type WaitFailureError struct {
	State    string
	Response any
}

func (e *WaitFailureError) Error() string {
	return fmt.Sprintf("sdk: object entered failure state %q", e.State)
}

// ObjectReadyError reports a failed readiness wait. Object and ID identify the
// handle. Disposed is true when a created object was cleaned up; CleanupErr is
// the cleanup failure, which never replaces Cause.
type ObjectReadyError struct {
	Object     any
	ID         string
	Cause      error
	Disposed   bool
	CleanupErr error
}

func (e *ObjectReadyError) Error() string {
	return fmt.Sprintf("sdk: readiness of %s failed: %v", e.ID, e.Cause)
}

func (e *ObjectReadyError) Unwrap() error { return e.Cause }

// OutputValidationError reports a missing or undecodable wait output.
type OutputValidationError struct {
	ObjectID string
	Response any
	Cause    error
}

func (e *OutputValidationError) Error() string {
	return fmt.Sprintf("sdk: output validation failed for %s: %v", e.ObjectID, e.Cause)
}

func (e *OutputValidationError) Unwrap() error { return e.Cause }

// rawJSON returns the JSON the API sent for a response value.
func rawJSON(value any) string {
	if value == nil {
		return ""
	}
	if rv := reflect.ValueOf(value); rv.Kind() == reflect.Pointer && rv.IsNil() {
		return ""
	}
	switch v := value.(type) {
	case interface{ RawJSON() string }:
		return v.RawJSON()
	case json.RawMessage:
		return string(v)
	case string:
		return v
	case []byte:
		return string(v)
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(encoded)
}

func jsonPath(parts ...string) string {
	kept := make([]string, 0, len(parts))
	for _, part := range parts {
		if part != "" {
			kept = append(kept, part)
		}
	}
	return strings.Join(kept, ".")
}

func lookup(value any, path string) gjson.Result {
	raw := rawJSON(value)
	if raw == "" {
		return gjson.Result{}
	}
	if path == "" {
		return gjson.Parse(raw)
	}
	return gjson.Get(raw, path)
}

func objectID(value any, path string) string {
	result := lookup(value, path)
	if result.Type != gjson.String {
		return ""
	}
	return result.Str
}

func resultString(result gjson.Result) string {
	if result.Type != gjson.String {
		return ""
	}
	return result.Str
}

func resultInt(result gjson.Result) int64 {
	if result.Type != gjson.Number {
		return 0
	}
	return result.Int()
}

func resultFloat(result gjson.Result) float64 {
	if result.Type != gjson.Number {
		return 0
	}
	return result.Float()
}

func resultBool(result gjson.Result) bool {
	return result.Type == gjson.True
}

func resultValue(result gjson.Result) any {
	if !result.Exists() {
		return nil
	}
	return result.Value()
}

// objectState guards the last response of a handle. Handles may be shared
// between goroutines; methods that refresh info lock it.
type objectState struct {
	mu       sync.Mutex
	info     any
	infoRoot string
	owned    bool
}

func (s *objectState) setInfo(info any, root string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.info = info
	s.infoRoot = root
}

func (s *objectState) snapshot() (any, string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.info, s.infoRoot
}

func (s *objectState) field(path string) gjson.Result {
	info, root := s.snapshot()
	return lookup(info, jsonPath(root, path))
}

func (s *objectState) setOwned(owned bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.owned = owned
}

func (s *objectState) takeOwnership() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	owned := s.owned
	s.owned = false
	return owned
}

func objectString(class string, id *string, fields ...any) string {
	var b strings.Builder
	b.WriteString(class)
	b.WriteString("{")
	sep := ""
	if id != nil {
		b.WriteString("id=")
		b.WriteString(strconv.Quote(*id))
		sep = ", "
	}
	for i := 0; i+1 < len(fields); i += 2 {
		fmt.Fprintf(&b, "%s%v=%#v", sep, fields[i], fields[i+1])
		sep = ", "
	}
	b.WriteString("}")
	return b.String()
}

// objectRows returns the rows at path ("" for a root array) as raw JSON.
func objectRows(response any, path string) []json.RawMessage {
	rows := lookup(response, path)
	if !rows.IsArray() {
		return nil
	}
	result := make([]json.RawMessage, 0, len(rows.Array()))
	for _, row := range rows.Array() {
		result = append(result, json.RawMessage(row.Raw))
	}
	return result
}

// firstRowID returns the ID of the first row, or fallback for an empty response.
func firstRowID(response any, rowsPath, idPath, fallback string) string {
	rows := lookup(response, rowsPath)
	if !rows.IsArray() || len(rows.Array()) == 0 {
		return fallback
	}
	if id := rows.Array()[0].Get(idPath); id.Type == gjson.String {
		return id.Str
	}
	return fallback
}

// readyOptions forwards the caller's request options and readiness controls to
// the readiness wait that follows create or a ready-after method.
func readyOptions(c objectControls, opts []option.RequestOption) []option.RequestOption {
	result := slices.Clone(opts)
	if c.readyTimeout > 0 {
		result = append(result, WithReadyTimeout(c.readyTimeout))
	}
	return result
}

// waitOptions forwards the caller's request options and wait controls to the
// declared wait that follows a wait-after method.
func waitOptions(c objectControls, opts []option.RequestOption) []option.RequestOption {
	result := slices.Clone(opts)
	if c.waitTimeout > 0 {
		result = append(result, WithWaitTimeout(c.waitTimeout))
	}
	if c.pollInterval > 0 {
		result = append(result, WithPollInterval(c.pollInterval))
	}
	return result
}

// cleanupTimeout bounds best-effort cleanup after a failure; cleanup ignores the
// caller's cancellation but never blocks longer than this.
var cleanupTimeout = 10000 * time.Millisecond

func cleanupContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.WithoutCancel(ctx), cleanupTimeout)
}

// callerContext keeps ctx on every request of an SDK pager, including later
// pages, which the SDK fetches with a background context.
func callerContext(ctx context.Context) option.RequestOption {
	return option.WithMiddleware(func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		merged, cancel := context.WithCancel(req.Context())
		stop := context.AfterFunc(ctx, cancel)
		res, err := next(req.WithContext(merged))
		if err != nil || res == nil || res.Body == nil {
			stop()
			cancel()
			return res, err
		}
		res.Body = &releaseOnClose{ReadCloser: res.Body, release: func() { stop(); cancel() }}
		return res, nil
	})
}

type releaseOnClose struct {
	io.ReadCloser
	release func()
	once    sync.Once
}

func (r *releaseOnClose) Close() error {
	err := r.ReadCloser.Close()
	r.once.Do(r.release)
	return err
}

type binding struct {
	name     string
	location string
	value    any
}

// bindOptions applies bound and derived parameters after the caller's options,
// so they cannot be overridden.
func bindOptions(opts []option.RequestOption, bindings ...binding) []option.RequestOption {
	result := make([]option.RequestOption, 0, len(opts)+len(bindings))
	result = append(result, opts...)
	for _, b := range bindings {
		switch b.location {
		case "query":
			result = append(result, option.WithQuery(b.name, queryValue(b.value)))
		case "header":
			result = append(result, option.WithHeader(b.name, queryValue(b.value)))
		case "body":
			result = append(result, option.WithJSONSet(jsonKey(b.name), b.value))
		}
	}
	return result
}

func queryValue(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case time.Time:
		return v.Format(time.RFC3339Nano)
	case nil:
		return ""
	}
	return fmt.Sprint(value)
}

func jsonKey(name string) string {
	replacer := strings.NewReplacer(".", "\\.", "*", "\\*", "?", "\\?", "|", "\\|", "#", "\\#", "@", "\\@")
	return replacer.Replace(name)
}

// ObjectPage is one page of object handles. Response is the raw API response of
// the page.
type ObjectPage[T any, R any] struct {
	Items    []T
	Response R
	next     func() (*ObjectPage[T, R], error)
}

// HasNextPage reports whether another page may follow. Pages of SDK-paginated
// lists only know this after the next request.
func (p *ObjectPage[T, R]) HasNextPage() bool { return p != nil && p.next != nil }

// GetNextPage fetches the next page. After the last page it returns nil and no
// error.
func (p *ObjectPage[T, R]) GetNextPage() (*ObjectPage[T, R], error) {
	if !p.HasNextPage() {
		return nil, nil
	}
	return p.next()
}

func (p *ObjectPage[T, R]) objectItems() []T { return p.Items }

func (p *ObjectPage[T, R]) nextObjectPage() (objectPageSource[T], error) {
	next, err := p.GetNextPage()
	if next == nil {
		return nil, err
	}
	return next, err
}

type objectPageSource[T any] interface {
	objectItems() []T
	nextObjectPage() (objectPageSource[T], error)
}

// ObjectAutoPager iterates the handles of every page:
//
//	for pager.Next() {
//		handle := pager.Current()
//	}
//	if err := pager.Err(); err != nil { ... }
type ObjectAutoPager[T any] struct {
	page objectPageSource[T]
	cur  T
	idx  int
	err  error
}

func newObjectAutoPager[T any, R any](page *ObjectPage[T, R], err error) *ObjectAutoPager[T] {
	pager := &ObjectAutoPager[T]{err: err}
	if page != nil {
		pager.page = page
	}
	return pager
}

// Next advances to the next handle, fetching pages as needed.
func (r *ObjectAutoPager[T]) Next() bool {
	for r.page != nil && r.err == nil {
		items := r.page.objectItems()
		if r.idx < len(items) {
			r.cur = items[r.idx]
			r.idx++
			return true
		}
		r.page, r.err = r.page.nextObjectPage()
		r.idx = 0
	}
	return false
}

// Current returns the handle Next advanced to.
func (r *ObjectAutoPager[T]) Current() T { return r.cur }

// Err returns the error that stopped the iteration, if any.
func (r *ObjectAutoPager[T]) Err() error { return r.err }

const singleCallInterval = time.Millisecond

var (
	retryStatuses     = []int{408, 409, 429}
	immediateStatuses = []int{408}
)

const (
	serverErrorMin   = 500
	initialBackoff   = 500 * time.Millisecond
	maxBackoff       = 8000 * time.Millisecond
	backoffJitterMin = 0.75
)

type retryFunc func(err error, attempt int, interval time.Duration) (time.Duration, bool)

// waitFor polls until poll reports done, a non-retryable error occurs or the
// deadline passes. Each poll runs without native retries and with a request
// timeout capped at the remaining deadline.
func waitFor[T any](
	ctx context.Context,
	timeout time.Duration,
	interval time.Duration,
	pollTimeout time.Duration,
	poll func(ctx context.Context, opts []option.RequestOption) (T, bool, error),
	retry retryFunc,
) (T, error) {
	var zero T
	if timeout <= 0 || interval <= 0 {
		return zero, errors.New("sdk: wait timeout and poll interval must be positive")
	}
	deadline := time.Now().Add(timeout)
	waitCtx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	var lastErr error
	attempt := 0
	expired := func(err error) (T, error) {
		if ctx.Err() != nil {
			return zero, ctx.Err()
		}
		if lastErr != nil {
			err = lastErr
		}
		return zero, &WaitTimeoutError{Cause: err}
	}
	for {
		if ctx.Err() != nil {
			return zero, ctx.Err()
		}
		remaining := time.Until(deadline)
		if remaining <= 0 {
			return expired(nil)
		}
		requestTimeout := remaining
		if pollTimeout > 0 && pollTimeout < requestTimeout {
			requestTimeout = pollTimeout
		}
		result, done, err := poll(waitCtx, []option.RequestOption{option.WithMaxRetries(0), option.WithRequestTimeout(requestTimeout)})
		if err != nil {
			if ctx.Err() != nil {
				return zero, ctx.Err()
			}
			if waitCtx.Err() != nil {
				return expired(err)
			}
			delay, retryable := retry(err, attempt, interval)
			attempt++
			if !retryable {
				return zero, err
			}
			lastErr = err
			if !sleepUntil(waitCtx, delay, deadline) {
				return expired(nil)
			}
			continue
		}
		attempt = 0
		if done {
			return result, nil
		}
		if !sleepUntil(waitCtx, interval, deadline) {
			return expired(nil)
		}
	}
}

func sleepUntil(ctx context.Context, delay time.Duration, deadline time.Time) bool {
	if remaining := time.Until(deadline); delay > remaining {
		delay = remaining
	}
	if delay < 0 {
		delay = 0
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}

func isTimeout(err error) bool {
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}
	var netErr net.Error
	return errors.As(err, &netErr) && netErr.Timeout()
}

func isTransport(err error) bool {
	var urlErr *url.Error
	var netErr net.Error
	return isTimeout(err) || errors.As(err, &urlErr) || errors.As(err, &netErr)
}

func retryableStatus(status int) bool {
	return slices.Contains(retryStatuses, status) || status >= serverErrorMin
}

func errorHeader(apiErr *mobilerun.Error, name string) string {
	if apiErr.Response == nil {
		return ""
	}
	return apiErr.Response.Header.Get(name)
}

// retryReadiness retries request timeouts and 408 immediately, and 409, 429 and
// server errors with backoff; x-should-retry overrides the status decision.
func retryReadiness(err error, attempt int, _ time.Duration) (time.Duration, bool) {
	var apiErr *mobilerun.Error
	if !errors.As(err, &apiErr) {
		return 0, isTimeout(err)
	}
	hint := errorHeader(apiErr, "x-should-retry")
	if hint == "false" {
		return 0, false
	}
	if hint != "true" && !retryableStatus(apiErr.StatusCode) {
		return 0, false
	}
	if slices.Contains(immediateStatuses, apiErr.StatusCode) {
		return 0, true
	}
	return retryBackoff(apiErr, attempt), true
}

// retryPoll retries transport errors after the poll interval, and 408, 409, 429
// and server errors with backoff.
func retryPoll(err error, attempt int, interval time.Duration) (time.Duration, bool) {
	var apiErr *mobilerun.Error
	if !errors.As(err, &apiErr) {
		return interval, isTransport(err)
	}
	if !retryableStatus(apiErr.StatusCode) {
		return 0, false
	}
	return retryBackoff(apiErr, attempt), true
}

func retryBackoff(apiErr *mobilerun.Error, attempt int) time.Duration {
	if value := strings.TrimSpace(errorHeader(apiErr, "retry-after-ms")); value != "" {
		if ms, err := strconv.ParseFloat(value, 64); err == nil && ms >= 0 && !math.IsInf(ms, 0) {
			return time.Duration(ms * float64(time.Millisecond))
		}
	}
	if value := strings.TrimSpace(errorHeader(apiErr, "retry-after")); value != "" {
		if seconds, err := strconv.ParseFloat(value, 64); err == nil {
			if seconds >= 0 && !math.IsInf(seconds, 0) {
				return time.Duration(seconds * float64(time.Second))
			}
		} else if at, err := http.ParseTime(value); err == nil {
			if delay := time.Until(at); delay >= 0 {
				return delay
			}
		}
	}
	if attempt > 10 {
		attempt = 10
	}
	base := initialBackoff << attempt
	if base > maxBackoff {
		base = maxBackoff
	}
	return time.Duration(float64(base) * (backoffJitterMin + rand.Float64()*(1-backoffJitterMin)))
}

func enumState(response any, path string, success, failure, pending []string, unlistedPending bool) (bool, error) {
	state := resultString(lookup(response, path))
	if slices.Contains(success, state) {
		return true, nil
	}
	if slices.Contains(failure, state) || (!unlistedPending && !slices.Contains(pending, state)) {
		return false, &WaitFailureError{State: state, Response: response}
	}
	return false, nil
}

func conditionMet(result gjson.Result, until string) bool {
	if until == "present" {
		return result.Exists() && result.Type != gjson.Null
	}
	if result.Type == gjson.String {
		return result.Str != ""
	}
	return result.IsArray() && len(result.Array()) > 0
}

func decodeOutput(objectID string, response any, path string, out any) error {
	value := lookup(response, path)
	if !value.Exists() || value.Type == gjson.Null {
		return &OutputValidationError{ObjectID: objectID, Response: response, Cause: errors.New("completed object has no output")}
	}
	if err := json.Unmarshal([]byte(value.Raw), out); err != nil {
		return &OutputValidationError{ObjectID: objectID, Response: response, Cause: err}
	}
	return nil
}

// nextPageNumber reads page-number pagination metadata. requested is 0 for the
// caller's first request, whose page is not known to the object layer.
func nextPageNumber(response any, currentPath, hasNextPath, totalPath string, requested int64) (int64, bool, error) {
	current := lookup(response, currentPath)
	hasNext := lookup(response, hasNextPath)
	total := lookup(response, totalPath)
	invalid := errors.New("sdk: invalid object pagination metadata")
	if current.Type != gjson.Number || total.Type != gjson.Number || (hasNext.Type != gjson.True && hasNext.Type != gjson.False) {
		return 0, false, invalid
	}
	if current.Float() != float64(current.Int()) || total.Float() != float64(total.Int()) || current.Int() < 0 || total.Int() < 0 {
		return 0, false, invalid
	}
	if !hasNext.Bool() || current.Int() >= total.Int() {
		return 0, false, nil
	}
	if requested > 0 && current.Int() < requested {
		return 0, false, errors.New("sdk: object pagination made no forward progress")
	}
	return current.Int() + 1, true, nil
}

// nextObjectCursor reads cursor pagination metadata. An absent, null or empty
// cursor ends the iteration; a repeated cursor is rejected.
func nextObjectCursor(response any, path string, seen map[string]bool) (string, bool, error) {
	next := lookup(response, path)
	if !next.Exists() || next.Type == gjson.Null || (next.Type == gjson.String && next.Str == "") {
		return "", false, nil
	}
	if next.Type != gjson.String {
		return "", false, errors.New("sdk: invalid object cursor pagination metadata")
	}
	if seen[next.Str] {
		return "", false, errors.New("sdk: object pagination made no forward progress")
	}
	return next.Str, true, nil
}

// BinaryResult holds the bytes of a binary response.
type BinaryResult struct {
	Bytes       []byte
	ContentType string
}

// Base64 returns the bytes as standard base64.
func (b *BinaryResult) Base64() string { return base64.StdEncoding.EncodeToString(b.Bytes) }

// DataURL returns a data URL with the response content type.
func (b *BinaryResult) DataURL() string {
	contentType := b.ContentType
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	return "data:" + contentType + ";base64," + b.Base64()
}

// Save writes the bytes to path.
func (b *BinaryResult) Save(path string) error { return os.WriteFile(path, b.Bytes, 0o644) }

func headerBool(value string) *bool {
	switch value {
	case "true":
		result := true
		return &result
	case "false":
		result := false
		return &result
	}
	return nil
}

// WebhookVerificationError reports why a signed delivery was rejected. Reason is
// missing_header, invalid_signature, timestamp_out_of_tolerance, malformed or
// invalid_json (the signature is valid but the body is not JSON; respond 400).
type WebhookVerificationError struct {
	Reason string
}

func (e *WebhookVerificationError) Error() string {
	return "sdk: webhook verification failed: " + e.Reason
}

// SignatureVerifyOptions overrides the declared tolerance or the current time.
type SignatureVerifyOptions struct {
	// Tolerance accepted in either time direction; zero uses the declared default.
	Tolerance time.Duration
	// Now is the verification time; zero uses time.Now().
	Now time.Time
}

type signatureScheme struct {
	signatureHeader string
	timestampHeader string
	prefix          string
	unit            string
	tolerance       time.Duration
}

var signatureDigest = regexp.MustCompile("^[0-9a-fA-F]{64}$")
var signatureTimestamp = regexp.MustCompile("^[0-9]+$")

func signatureHeaderValue(headers http.Header, name string) (string, error) {
	var values []string
	for key, value := range headers {
		if strings.EqualFold(key, name) {
			if values != nil {
				return "", &WebhookVerificationError{Reason: "malformed"}
			}
			values = value
		}
	}
	if values == nil {
		return "", &WebhookVerificationError{Reason: "missing_header"}
	}
	if len(values) != 1 {
		return "", &WebhookVerificationError{Reason: "malformed"}
	}
	return values[0], nil
}

// verifySignature checks an HMAC-SHA256 signature over "{timestamp}.{body}" and
// returns the parsed JSON body. Pass the exact raw bytes before JSON parsing.
func verifySignature(body []byte, headers http.Header, secret string, opts *SignatureVerifyOptions, scheme signatureScheme) (any, error) {
	timestamp, err := signatureHeaderValue(headers, scheme.timestampHeader)
	if err != nil {
		return nil, err
	}
	signature, err := signatureHeaderValue(headers, scheme.signatureHeader)
	if err != nil {
		return nil, err
	}
	tolerance := scheme.tolerance
	now := time.Now()
	if opts != nil {
		if opts.Tolerance < 0 {
			return nil, &WebhookVerificationError{Reason: "malformed"}
		}
		if opts.Tolerance > 0 {
			tolerance = opts.Tolerance
		}
		if !opts.Now.IsZero() {
			now = opts.Now
		}
	}
	if !signatureTimestamp.MatchString(timestamp) || secret == "" || !strings.HasPrefix(signature, scheme.prefix) {
		return nil, &WebhookVerificationError{Reason: "malformed"}
	}
	digest := strings.TrimPrefix(signature, scheme.prefix)
	if !signatureDigest.MatchString(digest) {
		return nil, &WebhookVerificationError{Reason: "malformed"}
	}
	numeric, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return nil, &WebhookVerificationError{Reason: "malformed"}
	}
	var signedAt time.Time
	if scheme.unit == "s" {
		if numeric > math.MaxInt64/int64(time.Second) {
			return nil, &WebhookVerificationError{Reason: "malformed"}
		}
		signedAt = time.Unix(numeric, 0)
	} else {
		signedAt = time.UnixMilli(numeric)
	}
	skew := now.Sub(signedAt)
	if skew < 0 {
		skew = -skew
	}
	if skew > tolerance {
		return nil, &WebhookVerificationError{Reason: "timestamp_out_of_tolerance"}
	}
	supplied, err := hex.DecodeString(digest)
	if err != nil {
		return nil, &WebhookVerificationError{Reason: "malformed"}
	}
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + "."))
	mac.Write(body)
	if !hmac.Equal(mac.Sum(nil), supplied) {
		return nil, &WebhookVerificationError{Reason: "invalid_signature"}
	}
	decoded := bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	if !utf8.Valid(decoded) {
		return nil, &WebhookVerificationError{Reason: "invalid_json"}
	}
	var value any
	if err := json.Unmarshal(decoded, &value); err != nil {
		return nil, &WebhookVerificationError{Reason: "invalid_json"}
	}
	return value, nil
}

// UploadError wraps a failed presigned upload. Cleanup failures never replace
// Cause.
type UploadError struct {
	Cause error
}

func (e *UploadError) Error() string { return "sdk: presigned upload failed: " + e.Cause.Error() }

func (e *UploadError) Unwrap() error { return e.Cause }

// UploadFile is one file to upload. Set Path to stream a local file (reopened
// for each retry), or Data for bytes in memory. Filename is required for Data;
// ContentType defaults from the file extension; Size, when set, must match the
// measured byte count.
type UploadFile struct {
	Path        string
	Data        []byte
	Filename    string
	ContentType string
	Size        int64
}

// UploadPath returns an UploadFile streaming the file at path.
func UploadPath(path string) UploadFile { return UploadFile{Path: path} }

// UploadBytes returns an UploadFile for bytes in memory.
func UploadBytes(filename string, data []byte) UploadFile {
	return UploadFile{Filename: filename, Data: data}
}

type preparedUpload struct {
	filename    string
	contentType string
	size        int64
	open        func() (io.ReadCloser, error)
}

var uploadContentTypes = map[string]string{
	"apk": "application/vnd.android.package-archive", "xapk": "application/zip", "apks": "application/zip",
	"ipa": "application/octet-stream", "json": "application/json", "txt": "text/plain", "csv": "text/csv",
	"pdf": "application/pdf", "png": "image/png", "jpg": "image/jpeg", "jpeg": "image/jpeg", "zip": "application/zip",
}

func uploadContentType(filename string) string {
	extension := strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
	if contentType, ok := uploadContentTypes[extension]; ok {
		return contentType
	}
	return "application/octet-stream"
}

func prepareUploads(files []UploadFile, allowed []string) ([]preparedUpload, error) {
	prepared := make([]preparedUpload, 0, len(files))
	names := map[string]bool{}
	for _, file := range files {
		upload := preparedUpload{filename: file.Filename, contentType: file.ContentType}
		var measured int64
		if file.Path != "" {
			if file.Data != nil {
				return nil, errors.New("sdk: an upload takes either Path or Data")
			}
			stat, err := os.Stat(file.Path)
			if err != nil {
				return nil, err
			}
			if !stat.Mode().IsRegular() {
				return nil, errors.New("sdk: upload path must name a file")
			}
			if upload.filename == "" {
				upload.filename = filepath.Base(file.Path)
			}
			measured = stat.Size()
			path := file.Path
			upload.open = func() (io.ReadCloser, error) { return os.Open(path) }
		} else {
			data := file.Data
			measured = int64(len(data))
			upload.open = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(data)), nil }
		}
		if upload.filename == "" {
			return nil, errors.New("sdk: uploads from bytes require a filename")
		}
		if file.Size != 0 && file.Size != measured {
			return nil, errors.New("sdk: upload size must match the measured byte count")
		}
		upload.size = measured
		if upload.contentType == "" {
			upload.contentType = uploadContentType(upload.filename)
		}
		if allowed != nil && !slices.Contains(allowed, upload.contentType) {
			return nil, fmt.Errorf("sdk: unsupported upload content type %q", upload.contentType)
		}
		if names[upload.filename] {
			return nil, errors.New("sdk: uploads need unique, nonempty file names")
		}
		names[upload.filename] = true
		prepared = append(prepared, upload)
	}
	if len(prepared) == 0 {
		return nil, errors.New("sdk: uploads need unique, nonempty file names")
	}
	return prepared, nil
}

// uploadURLs matches every file to exactly one presigned URL.
func uploadURLs(response any, files []preparedUpload, urlPath, urlsPath, nameField, urlField string) ([]string, error) {
	if urlPath != "" {
		url := lookup(response, urlPath)
		if len(files) != 1 || url.Type != gjson.String {
			return nil, errors.New("sdk: invalid presigned upload URL")
		}
		return []string{url.Str}, nil
	}
	rows := lookup(response, urlsPath)
	if !rows.IsArray() || len(rows.Array()) != len(files) {
		return nil, errors.New("sdk: upload URL count or file names do not match")
	}
	byName := map[string]string{}
	for _, row := range rows.Array() {
		name := row.Get(nameField)
		if name.Type != gjson.String {
			return nil, errors.New("sdk: upload URL row without a file name")
		}
		if _, ok := byName[name.Str]; ok {
			return nil, errors.New("sdk: duplicate upload URL name")
		}
		url := row.Get(urlField)
		if url.Type != gjson.String {
			return nil, errors.New("sdk: missing upload URL for a file")
		}
		byName[name.Str] = url.Str
	}
	urls := make([]string, len(files))
	for i, file := range files {
		url, ok := byName[file.filename]
		if !ok {
			return nil, errors.New("sdk: missing upload URL for a file")
		}
		urls[i] = url
	}
	return urls, nil
}

// definiteConfirmFailure reports a confirm failure after which nothing was stored.
func definiteConfirmFailure(err error) bool {
	var apiErr *mobilerun.Error
	return errors.As(err, &apiErr) && apiErr.StatusCode >= 400 && apiErr.StatusCode < 500 && apiErr.StatusCode != 408
}

// putUpload sends one file to a presigned URL through the SDK's own request
// execution (HTTP client or custom doer, timeout, retry policy and backoff), with
// every API header removed, redirects rejected and the signed query string kept
// out of errors.
func putUpload(ctx context.Context, client *clientRef, file preparedUpload, target string, opts []option.RequestOption) error {
	parsed, err := url.Parse(target)
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") || parsed.User != nil {
		return errors.New("sdk: invalid presigned URL")
	}
	var raw *http.Response
	options := append(append(slices.Clone(client.options), opts...), option.WithMiddleware(storageRequest(file.contentType)), option.WithResponseInto(&raw))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, target, nil, nil, options...)
	if err != nil {
		return err
	}
	if cfg.CustomHTTPDoer == nil && cfg.HTTPClient != nil {
		noRedirect := *cfg.HTTPClient
		noRedirect.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
		cfg.HTTPClient = &noRedirect
	}
	body, err := file.open()
	if err != nil {
		return err
	}
	cfg.Request.Body = body
	cfg.Request.ContentLength = file.size
	cfg.Request.GetBody = file.open
	err = cfg.Execute()
	if raw != nil && raw.Body != nil {
		_, _ = io.Copy(io.Discard, raw.Body)
		_ = raw.Body.Close()
	}
	if err != nil {
		return sanitizeStorageError(err, parsed)
	}
	if raw == nil || raw.StatusCode >= 300 || raw.Request == nil || raw.Request.Method != http.MethodPut || raw.Request.URL.String() != parsed.String() {
		status := 0
		if raw != nil {
			status = raw.StatusCode
		}
		return fmt.Errorf("sdk: presigned upload to %s was redirected or not accepted (status %d)", sanitizedURL(parsed), status)
	}
	return nil
}

// storageRequest runs closest to the transport: storage receives only the
// content headers, and storage error bodies are never parsed or retained.
func storageRequest(contentType string) option.Middleware {
	return func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		req.Header = http.Header{"Content-Type": []string{contentType}}
		res, err := next(req)
		if res != nil && res.StatusCode >= 300 && res.Body != nil {
			_, _ = io.Copy(io.Discard, res.Body)
			_ = res.Body.Close()
			res.Body = io.NopCloser(strings.NewReader("{}"))
		}
		return res, err
	}
}

// storageError carries a storage failure with the signed URL removed from its
// message and from every wrapped error.
type storageError struct {
	msg   string
	cause error
}

func (e *storageError) Error() string { return e.msg }
func (e *storageError) Unwrap() error { return e.cause }

func sanitizeStorageError(err error, target *url.URL) error {
	clean := sanitizedURL(target)
	var apiErr *mobilerun.Error
	if errors.As(err, &apiErr) {
		if apiErr.Request != nil {
			sanitized := apiErr.Request.Clone(context.Background())
			sanitized.URL, _ = url.Parse(clean)
			apiErr.Request = sanitized
		}
		if apiErr.Response != nil {
			apiErr.Response.Request = apiErr.Request
		}
		return err
	}
	root := err
	for {
		var urlErr *url.Error
		if !errors.As(root, &urlErr) {
			break
		}
		root = urlErr.Err
	}
	msg := strings.ReplaceAll(err.Error(), target.String(), clean)
	if target.RawQuery != "" {
		msg = strings.ReplaceAll(msg, target.RawQuery, "")
	}
	cause := root
	if target.RawQuery != "" && cause != nil && strings.Contains(cause.Error(), target.RawQuery) {
		cause = nil
	}
	return &storageError{msg: "sdk: presigned upload failed: " + msg, cause: cause}
}

func sanitizedURL(u *url.URL) string {
	return (&url.URL{Scheme: u.Scheme, Host: u.Host, Path: u.Path}).String()
}
