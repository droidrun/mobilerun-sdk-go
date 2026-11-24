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

// ComputeFunctionRunService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeFunctionRunService] method instead.
type ComputeFunctionRunService struct {
	options []option.RequestOption
}

// NewComputeFunctionRunService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewComputeFunctionRunService(opts ...option.RequestOption) (r ComputeFunctionRunService) {
	r = ComputeFunctionRunService{}
	r.options = opts
	return
}

// Returns 200 with the finished run when it ends within `waitS`, else 202 with the
// run.
func (r *ComputeFunctionRunService) New(ctx context.Context, functionID string, body ComputeFunctionRunNewParams, opts ...option.RequestOption) (res *ComputeFunctionRunNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/functions/%s/runs", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List the runs of a function, newest first
func (r *ComputeFunctionRunService) List(ctx context.Context, functionID string, query ComputeFunctionRunListParams, opts ...option.RequestOption) (res *ComputeFunctionRunListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/functions/%s/runs", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ComputeFunctionRunNewResponse struct {
	Data ComputeFunctionRunNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionRunNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionRunNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionRunNewResponseData struct {
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
	Outcome   string                                       `json:"outcome" api:"required"`
	Result    ComputeFunctionRunNewResponseDataResultUnion `json:"result" api:"required"`
	StartedAt time.Time                                    `json:"startedAt" api:"required" format:"date-time"`
	// Any of "running", "done", "failed".
	Status      string                                       `json:"status" api:"required"`
	Stderr      string                                       `json:"stderr" api:"required"`
	Stdout      string                                       `json:"stdout" api:"required"`
	Truncated   bool                                         `json:"truncated" api:"required"`
	UsedSeconds int64                                        `json:"usedSeconds" api:"required"`
	Version     int64                                        `json:"version" api:"required"`
	APIKey      ComputeFunctionRunNewResponseDataAPIKeyUnion `json:"apiKey"`
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
func (r ComputeFunctionRunNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionRunNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeFunctionRunNewResponseDataResultUnion contains all possible properties
// and values from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfComputeFunctionRunNewResponseDataResultMapItem OfAnyArray
// OfString OfFloat OfBool]
type ComputeFunctionRunNewResponseDataResultUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfComputeFunctionRunNewResponseDataResultMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfComputeFunctionRunNewResponseDataResultMapItem respjson.Field
		OfAnyArray                                       respjson.Field
		OfString                                         respjson.Field
		OfFloat                                          respjson.Field
		OfBool                                           respjson.Field
		raw                                              string
	} `json:"-"`
}

func (u ComputeFunctionRunNewResponseDataResultUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeFunctionRunNewResponseDataResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeFunctionRunNewResponseDataResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeFunctionRunNewResponseDataResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeFunctionRunNewResponseDataResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeFunctionRunNewResponseDataResultUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeFunctionRunNewResponseDataResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeFunctionRunNewResponseDataAPIKeyUnion contains all possible properties
// and values from [ComputeFunctionRunNewResponseDataAPIKeyObject],
// [ComputeFunctionRunNewResponseDataAPIKeyObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeFunctionRunNewResponseDataAPIKeyUnion struct {
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant [ComputeFunctionRunNewResponseDataAPIKeyObject].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [ComputeFunctionRunNewResponseDataAPIKeyObject2].
	Full bool `json:"full"`
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeFunctionRunNewResponseDataAPIKeyUnion) AsComputeFunctionRunNewResponseDataAPIKeyObject() (v ComputeFunctionRunNewResponseDataAPIKeyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeFunctionRunNewResponseDataAPIKeyUnion) AsComputeFunctionRunNewResponseDataAPIKeyObject2() (v ComputeFunctionRunNewResponseDataAPIKeyObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeFunctionRunNewResponseDataAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeFunctionRunNewResponseDataAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionRunNewResponseDataAPIKeyObject struct {
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
func (r ComputeFunctionRunNewResponseDataAPIKeyObject) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionRunNewResponseDataAPIKeyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionRunNewResponseDataAPIKeyObject2 struct {
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
func (r ComputeFunctionRunNewResponseDataAPIKeyObject2) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionRunNewResponseDataAPIKeyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionRunListResponse struct {
	Data ComputeFunctionRunListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionRunListResponseData struct {
	Items      []ComputeFunctionRunListResponseDataItem `json:"items" api:"required"`
	Pagination shared.Pagination                        `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionRunListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionRunListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionRunListResponseDataItem struct {
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
	Status      string                                            `json:"status" api:"required"`
	UsedSeconds int64                                             `json:"usedSeconds" api:"required"`
	Version     int64                                             `json:"version" api:"required"`
	APIKey      ComputeFunctionRunListResponseDataItemAPIKeyUnion `json:"apiKey"`
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
func (r ComputeFunctionRunListResponseDataItem) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionRunListResponseDataItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeFunctionRunListResponseDataItemAPIKeyUnion contains all possible
// properties and values from [ComputeFunctionRunListResponseDataItemAPIKeyObject],
// [ComputeFunctionRunListResponseDataItemAPIKeyObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeFunctionRunListResponseDataItemAPIKeyUnion struct {
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant [ComputeFunctionRunListResponseDataItemAPIKeyObject].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant
	// [ComputeFunctionRunListResponseDataItemAPIKeyObject2].
	Full bool `json:"full"`
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u ComputeFunctionRunListResponseDataItemAPIKeyUnion) AsComputeFunctionRunListResponseDataItemAPIKeyObject() (v ComputeFunctionRunListResponseDataItemAPIKeyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeFunctionRunListResponseDataItemAPIKeyUnion) AsComputeFunctionRunListResponseDataItemAPIKeyObject2() (v ComputeFunctionRunListResponseDataItemAPIKeyObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeFunctionRunListResponseDataItemAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeFunctionRunListResponseDataItemAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionRunListResponseDataItemAPIKeyObject struct {
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
func (r ComputeFunctionRunListResponseDataItemAPIKeyObject) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionRunListResponseDataItemAPIKeyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionRunListResponseDataItemAPIKeyObject2 struct {
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
func (r ComputeFunctionRunListResponseDataItemAPIKeyObject2) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionRunListResponseDataItemAPIKeyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionRunNewParams struct {
	Input map[string]any `json:"input,omitzero" api:"required"`
	// Omitted: 300 s.
	TimeoutS param.Opt[int64] `json:"timeoutS,omitzero"`
	// Seconds the reply waits for the run to end; 202 when it is still running.
	WaitS param.Opt[int64] `json:"waitS,omitzero"`
	// Latest selects the newest ready version of the function; a number pins a
	// version.
	Version ComputeFunctionRunNewParamsVersionUnion `json:"version,omitzero"`
	paramObj
}

func (r ComputeFunctionRunNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionRunNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionRunNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ComputeFunctionRunNewParamsVersionUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfComputeFunctionRunNewsVersionString)
	OfComputeFunctionRunNewsVersionString param.Opt[string] `json:",omitzero,inline"`
	OfInt                                 param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u ComputeFunctionRunNewParamsVersionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfComputeFunctionRunNewsVersionString, u.OfInt)
}
func (u *ComputeFunctionRunNewParamsVersionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ComputeFunctionRunNewParamsVersionString string

const (
	ComputeFunctionRunNewParamsVersionStringLatest ComputeFunctionRunNewParamsVersionString = "latest"
)

type ComputeFunctionRunListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Any of "running", "done", "failed".
	Status ComputeFunctionRunListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeFunctionRunListParams]'s query parameters as
// `url.Values`.
func (r ComputeFunctionRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeFunctionRunListParamsStatus string

const (
	ComputeFunctionRunListParamsStatusRunning ComputeFunctionRunListParamsStatus = "running"
	ComputeFunctionRunListParamsStatusDone    ComputeFunctionRunListParamsStatus = "done"
	ComputeFunctionRunListParamsStatusFailed  ComputeFunctionRunListParamsStatus = "failed"
)
