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
)

// ComputeSandboxCommandService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeSandboxCommandService] method instead.
type ComputeSandboxCommandService struct {
	options []option.RequestOption
}

// NewComputeSandboxCommandService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewComputeSandboxCommandService(opts ...option.RequestOption) (r ComputeSandboxCommandService) {
	r = ComputeSandboxCommandService{}
	r.options = opts
	return
}

// Get a command result
func (r *ComputeSandboxCommandService) Get(ctx context.Context, commandID string, params ComputeSandboxCommandGetParams, opts ...option.RequestOption) (res *ComputeSandboxCommandGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.SandboxID == "" {
		err = errors.New("missing required sandboxId parameter")
		return nil, err
	}
	if commandID == "" {
		err = errors.New("missing required commandId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/sandboxes/%s/commands/%s", params.SandboxID, commandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List the newest commands
func (r *ComputeSandboxCommandService) List(ctx context.Context, sandboxID string, opts ...option.RequestOption) (res *ComputeSandboxCommandListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if sandboxID == "" {
		err = errors.New("missing required sandboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/sandboxes/%s/commands", sandboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ComputeSandboxCommandGetResponse struct {
	Data ComputeSandboxCommandGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxCommandGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxCommandGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxCommandGetResponseData struct {
	CommandID  string    `json:"commandId" api:"required" format:"uuid"`
	DurationMs float64   `json:"durationMs" api:"required"`
	ExitCode   int64     `json:"exitCode" api:"required"`
	StartedAt  time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Any of "running", "exited".
	Status    string `json:"status" api:"required"`
	Stderr    string `json:"stderr" api:"required"`
	Stdout    string `json:"stdout" api:"required"`
	TimedOut  bool   `json:"timedOut" api:"required"`
	Truncated bool   `json:"truncated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommandID   respjson.Field
		DurationMs  respjson.Field
		ExitCode    respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Stderr      respjson.Field
		Stdout      respjson.Field
		TimedOut    respjson.Field
		Truncated   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxCommandGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxCommandGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxCommandListResponse struct {
	Data ComputeSandboxCommandListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxCommandListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxCommandListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxCommandListResponseData struct {
	Items []ComputeSandboxCommandListResponseDataItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxCommandListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxCommandListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxCommandListResponseDataItem struct {
	Command    string    `json:"command" api:"required"`
	CommandID  string    `json:"commandId" api:"required" format:"uuid"`
	DurationMs float64   `json:"durationMs" api:"required"`
	ExitCode   int64     `json:"exitCode" api:"required"`
	StartedAt  time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Any of "running", "exited".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command     respjson.Field
		CommandID   respjson.Field
		DurationMs  respjson.Field
		ExitCode    respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSandboxCommandListResponseDataItem) RawJSON() string { return r.JSON.raw }
func (r *ComputeSandboxCommandListResponseDataItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSandboxCommandGetParams struct {
	SandboxID string           `path:"sandboxId" api:"required" format:"uuid" json:"-"`
	WaitS     param.Opt[int64] `query:"waitS,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeSandboxCommandGetParams]'s query parameters as
// `url.Values`.
func (r ComputeSandboxCommandGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
