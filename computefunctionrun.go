// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

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
	Data FunctionRun `json:"data" api:"required"`
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
	Items      []FunctionRunSummary `json:"items" api:"required"`
	Pagination shared.Pagination    `json:"pagination" api:"required"`
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
