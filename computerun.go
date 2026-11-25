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
	Data FunctionRun `json:"data" api:"required"`
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
func (r ComputeRunListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeRunListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeRunCancelResponse struct {
	Data FunctionRun `json:"data" api:"required"`
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
