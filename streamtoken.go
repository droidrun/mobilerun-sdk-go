// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// StreamTokenService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamTokenService] method instead.
type StreamTokenService struct {
	options []option.RequestOption
}

// NewStreamTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamTokenService(opts ...option.RequestOption) (r StreamTokenService) {
	r = StreamTokenService{}
	r.options = opts
	return
}

// Mint a short-lived, single-use stream token for a device. The token permits
// streaming and control.
func (r *StreamTokenService) New(ctx context.Context, body StreamTokenNewParams, opts ...option.RequestOption) (res *StreamTokenCreateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "identity/stream-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Revoke a stream token by its id. The token can no longer be used to start a
// stream.
func (r *StreamTokenService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("identity/stream-keys/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type StreamTokenCreateResponse struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	Key       string    `json:"key" api:"required"`
	KeyID     string    `json:"keyId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Key         respjson.Field
		KeyID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamTokenCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *StreamTokenCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamTokenNewParams struct {
	DeviceID   string           `json:"deviceId" api:"required" format:"uuid"`
	TtlSeconds param.Opt[int64] `json:"ttlSeconds,omitzero"`
	// Exact `https://host[:port]` origins (and `http://localhost[:port]`) allowed to
	// connect.
	AllowedOrigins []string `json:"allowedOrigins,omitzero"`
	paramObj
}

func (r StreamTokenNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StreamTokenNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StreamTokenNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
