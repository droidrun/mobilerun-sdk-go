// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

// ComputeRunService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeRunService] method instead.
type ComputeRunService struct {
	options []option.RequestOption
}

// NewComputeRunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeRunService(opts ...option.RequestOption) (r ComputeRunService) {
	r = ComputeRunService{}
	r.options = opts
	return
}

// Get a run
func (r *ComputeRunService) Get(ctx context.Context, runID string, query ComputeRunGetParams, opts ...option.RequestOption) (res *ComputeRunGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List the runs started on the public routes, newest first
func (r *ComputeRunService) List(ctx context.Context, query ComputeRunListParams, opts ...option.RequestOption) (res *ComputeRunListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "compute/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Cancel a run
func (r *ComputeRunService) Cancel(ctx context.Context, runID string, opts ...option.RequestOption) (res *ComputeRunCancelResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/runs/%s/cancel", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get a short-lived download link for the full log of a run
func (r *ComputeRunService) Logs(ctx context.Context, runID string, opts ...option.RequestOption) (res *ComputeRunLogsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/runs/%s/logs", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ComputeRunGetResponse struct {
	Data ComputeRunGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunGetResponseData struct {
	ID         string    `json:"id" api:"required" format:"uuid"`
	DurationMs int64     `json:"durationMs" api:"required"`
	EndedAt    time.Time `json:"endedAt" api:"required" format:"date-time"`
	// Any of "function_error", "input_invalid", "timeout", "cancelled",
	// "sandbox_lost", "sandbox_start_failed", "function_not_found", "internal_error",
	// "output_invalid", "result_too_large", "state_too_large", "dataset_too_large",
	// "executor_unavailable".
	ErrorCode  string         `json:"errorCode" api:"required"`
	FunctionID string         `json:"functionId" api:"required" format:"uuid"`
	Input      map[string]any `json:"input" api:"required"`
	// Any of "none", "pending", "available", "gone".
	Logs string `json:"logs" api:"required"`
	// Any of "succeeded", "partial", "failed".
	Outcome   string                               `json:"outcome" api:"required"`
	Result    ComputeRunGetResponseDataResultUnion `json:"result" api:"required"`
	StartedAt time.Time                            `json:"startedAt" api:"required" format:"date-time"`
	// Any of "running", "done", "failed".
	Status      string                               `json:"status" api:"required"`
	Stderr      string                               `json:"stderr" api:"required"`
	Stdout      string                               `json:"stdout" api:"required"`
	Truncated   bool                                 `json:"truncated" api:"required"`
	UsedSeconds int64                                `json:"usedSeconds" api:"required"`
	Version     int64                                `json:"version" api:"required"`
	APIKey      ComputeRunGetResponseDataAPIKeyUnion `json:"apiKey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DurationMs  respjson.Field
		EndedAt     respjson.Field
		ErrorCode   respjson.Field
		FunctionID  respjson.Field
		Input       respjson.Field
		Logs        respjson.Field
		Outcome     respjson.Field
		Result      respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Stderr      respjson.Field
		Stdout      respjson.Field
		Truncated   respjson.Field
		UsedSeconds respjson.Field
		Version     respjson.Field
		APIKey      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeRunGetResponseDataResultUnion contains all possible properties and values
// from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfComputeRunGetResponseDataResultMapItem OfAnyArray OfString
// OfFloat OfBool]
type ComputeRunGetResponseDataResultUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfComputeRunGetResponseDataResultMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfComputeRunGetResponseDataResultMapItem respjson.Field
		OfAnyArray                               respjson.Field
		OfString                                 respjson.Field
		OfFloat                                  respjson.Field
		OfBool                                   respjson.Field
		raw                                      string
	} `json:"-"`
}

func (u ComputeRunGetResponseDataResultUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeRunGetResponseDataResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeRunGetResponseDataResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeRunGetResponseDataResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeRunGetResponseDataResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeRunGetResponseDataResultUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeRunGetResponseDataResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeRunGetResponseDataAPIKeyUnion contains all possible properties and values
// from [ComputeRunGetResponseDataAPIKeyObject],
// [ComputeRunGetResponseDataAPIKeyObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeRunGetResponseDataAPIKeyUnion struct {
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant [ComputeRunGetResponseDataAPIKeyObject].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [ComputeRunGetResponseDataAPIKeyObject2].
	Full bool `json:"full"`
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeRunGetResponseDataAPIKeyUnion) AsComputeRunGetResponseDataAPIKeyObject() (v ComputeRunGetResponseDataAPIKeyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeRunGetResponseDataAPIKeyUnion) AsComputeRunGetResponseDataAPIKeyObject2() (v ComputeRunGetResponseDataAPIKeyObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeRunGetResponseDataAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeRunGetResponseDataAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunGetResponseDataAPIKeyObject struct {
	ExpiresAt   time.Time           `json:"expiresAt" api:"required" format:"date-time"`
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunGetResponseDataAPIKeyObject) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunGetResponseDataAPIKeyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunGetResponseDataAPIKeyObject2 struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Any of true.
	Full bool `json:"full" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Full        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunGetResponseDataAPIKeyObject2) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunGetResponseDataAPIKeyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunListResponse struct {
	Data ComputeRunListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunListResponseData struct {
	Items      []ComputeRunListResponseDataItem `json:"items" api:"required"`
	Pagination shared.Pagination                `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunListResponseDataItem struct {
	ID         string    `json:"id" api:"required" format:"uuid"`
	DurationMs int64     `json:"durationMs" api:"required"`
	EndedAt    time.Time `json:"endedAt" api:"required" format:"date-time"`
	// Any of "function_error", "input_invalid", "timeout", "cancelled",
	// "sandbox_lost", "sandbox_start_failed", "function_not_found", "internal_error",
	// "output_invalid", "result_too_large", "state_too_large", "dataset_too_large",
	// "executor_unavailable".
	ErrorCode  string `json:"errorCode" api:"required"`
	FunctionID string `json:"functionId" api:"required" format:"uuid"`
	// Any of "none", "pending", "available", "gone".
	Logs string `json:"logs" api:"required"`
	// Any of "succeeded", "partial", "failed".
	Outcome   string    `json:"outcome" api:"required"`
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Any of "running", "done", "failed".
	Status      string                                    `json:"status" api:"required"`
	UsedSeconds int64                                     `json:"usedSeconds" api:"required"`
	Version     int64                                     `json:"version" api:"required"`
	APIKey      ComputeRunListResponseDataItemAPIKeyUnion `json:"apiKey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DurationMs  respjson.Field
		EndedAt     respjson.Field
		ErrorCode   respjson.Field
		FunctionID  respjson.Field
		Logs        respjson.Field
		Outcome     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		UsedSeconds respjson.Field
		Version     respjson.Field
		APIKey      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunListResponseDataItem) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunListResponseDataItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeRunListResponseDataItemAPIKeyUnion contains all possible properties and
// values from [ComputeRunListResponseDataItemAPIKeyObject],
// [ComputeRunListResponseDataItemAPIKeyObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeRunListResponseDataItemAPIKeyUnion struct {
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant [ComputeRunListResponseDataItemAPIKeyObject].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [ComputeRunListResponseDataItemAPIKeyObject2].
	Full bool `json:"full"`
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeRunListResponseDataItemAPIKeyUnion) AsComputeRunListResponseDataItemAPIKeyObject() (v ComputeRunListResponseDataItemAPIKeyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeRunListResponseDataItemAPIKeyUnion) AsComputeRunListResponseDataItemAPIKeyObject2() (v ComputeRunListResponseDataItemAPIKeyObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeRunListResponseDataItemAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeRunListResponseDataItemAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunListResponseDataItemAPIKeyObject struct {
	ExpiresAt   time.Time           `json:"expiresAt" api:"required" format:"date-time"`
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunListResponseDataItemAPIKeyObject) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunListResponseDataItemAPIKeyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunListResponseDataItemAPIKeyObject2 struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Any of true.
	Full bool `json:"full" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Full        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunListResponseDataItemAPIKeyObject2) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunListResponseDataItemAPIKeyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunCancelResponse struct {
	Data ComputeRunCancelResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunCancelResponseData struct {
	ID         string    `json:"id" api:"required" format:"uuid"`
	DurationMs int64     `json:"durationMs" api:"required"`
	EndedAt    time.Time `json:"endedAt" api:"required" format:"date-time"`
	// Any of "function_error", "input_invalid", "timeout", "cancelled",
	// "sandbox_lost", "sandbox_start_failed", "function_not_found", "internal_error",
	// "output_invalid", "result_too_large", "state_too_large", "dataset_too_large",
	// "executor_unavailable".
	ErrorCode  string         `json:"errorCode" api:"required"`
	FunctionID string         `json:"functionId" api:"required" format:"uuid"`
	Input      map[string]any `json:"input" api:"required"`
	// Any of "none", "pending", "available", "gone".
	Logs string `json:"logs" api:"required"`
	// Any of "succeeded", "partial", "failed".
	Outcome   string                                  `json:"outcome" api:"required"`
	Result    ComputeRunCancelResponseDataResultUnion `json:"result" api:"required"`
	StartedAt time.Time                               `json:"startedAt" api:"required" format:"date-time"`
	// Any of "running", "done", "failed".
	Status      string                                  `json:"status" api:"required"`
	Stderr      string                                  `json:"stderr" api:"required"`
	Stdout      string                                  `json:"stdout" api:"required"`
	Truncated   bool                                    `json:"truncated" api:"required"`
	UsedSeconds int64                                   `json:"usedSeconds" api:"required"`
	Version     int64                                   `json:"version" api:"required"`
	APIKey      ComputeRunCancelResponseDataAPIKeyUnion `json:"apiKey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DurationMs  respjson.Field
		EndedAt     respjson.Field
		ErrorCode   respjson.Field
		FunctionID  respjson.Field
		Input       respjson.Field
		Logs        respjson.Field
		Outcome     respjson.Field
		Result      respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Stderr      respjson.Field
		Stdout      respjson.Field
		Truncated   respjson.Field
		UsedSeconds respjson.Field
		Version     respjson.Field
		APIKey      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunCancelResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunCancelResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeRunCancelResponseDataResultUnion contains all possible properties and
// values from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfComputeRunCancelResponseDataResultMapItem OfAnyArray OfString
// OfFloat OfBool]
type ComputeRunCancelResponseDataResultUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfComputeRunCancelResponseDataResultMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfComputeRunCancelResponseDataResultMapItem respjson.Field
		OfAnyArray                                  respjson.Field
		OfString                                    respjson.Field
		OfFloat                                     respjson.Field
		OfBool                                      respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u ComputeRunCancelResponseDataResultUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeRunCancelResponseDataResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeRunCancelResponseDataResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeRunCancelResponseDataResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeRunCancelResponseDataResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeRunCancelResponseDataResultUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeRunCancelResponseDataResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeRunCancelResponseDataAPIKeyUnion contains all possible properties and
// values from [ComputeRunCancelResponseDataAPIKeyObject],
// [ComputeRunCancelResponseDataAPIKeyObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeRunCancelResponseDataAPIKeyUnion struct {
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant [ComputeRunCancelResponseDataAPIKeyObject].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [ComputeRunCancelResponseDataAPIKeyObject2].
	Full bool `json:"full"`
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeRunCancelResponseDataAPIKeyUnion) AsComputeRunCancelResponseDataAPIKeyObject() (v ComputeRunCancelResponseDataAPIKeyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeRunCancelResponseDataAPIKeyUnion) AsComputeRunCancelResponseDataAPIKeyObject2() (v ComputeRunCancelResponseDataAPIKeyObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeRunCancelResponseDataAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeRunCancelResponseDataAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunCancelResponseDataAPIKeyObject struct {
	ExpiresAt   time.Time           `json:"expiresAt" api:"required" format:"date-time"`
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunCancelResponseDataAPIKeyObject) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunCancelResponseDataAPIKeyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunCancelResponseDataAPIKeyObject2 struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Any of true.
	Full bool `json:"full" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Full        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunCancelResponseDataAPIKeyObject2) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunCancelResponseDataAPIKeyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunLogsResponse struct {
	Data ComputeRunLogsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunLogsResponseData struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	SizeBytes int64     `json:"sizeBytes" api:"required"`
	URL       string    `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		SizeBytes   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeRunLogsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunLogsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunGetParams struct {
	// Seconds the reply waits for the run to end; 202 when it is still running.
	WaitS param.Opt[int64] `query:"waitS,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeRunGetParams]'s query parameters as `url.Values`.
func (r ComputeRunGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeRunListParams struct {
	FunctionID param.Opt[string] `query:"functionId,omitzero" format:"uuid" json:"-"`
	Page       param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	// Any of "running", "done", "failed".
	Status ComputeRunListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeRunListParams]'s query parameters as `url.Values`.
func (r ComputeRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeRunListParamsStatus string

const (
	ComputeRunListParamsStatusRunning ComputeRunListParamsStatus = "running"
	ComputeRunListParamsStatusDone    ComputeRunListParamsStatus = "done"
	ComputeRunListParamsStatusFailed  ComputeRunListParamsStatus = "failed"
)
