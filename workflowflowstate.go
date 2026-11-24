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
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// WorkflowFlowStateService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowFlowStateService] method instead.
type WorkflowFlowStateService struct {
	options []option.RequestOption
}

// NewWorkflowFlowStateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowFlowStateService(opts ...option.RequestOption) (r WorkflowFlowStateService) {
	r = WorkflowFlowStateService{}
	r.options = opts
	return
}

// List the keys a routine has saved with `state.save`, with the metadata of each
// key's newest version.
func (r *WorkflowFlowStateService) List(ctx context.Context, flowID string, opts ...option.RequestOption) (res *WorkflowFlowStateListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/state", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete every version of one state key. The next `state.load` of the key reports
// `found: false`. Deleting an unknown key succeeds.
func (r *WorkflowFlowStateService) Delete(ctx context.Context, key string, body WorkflowFlowStateDeleteParams, opts ...option.RequestOption) (res *WorkflowFlowStateDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.FlowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/state/%s", body.FlowID, url.PathEscape(key))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Return the retained versions of one state key, newest first, including each
// saved value. An unknown key returns an empty list.
func (r *WorkflowFlowStateService) ListVersions(ctx context.Context, key string, query WorkflowFlowStateListVersionsParams, opts ...option.RequestOption) (res *WorkflowFlowStateListVersionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.FlowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/state/%s/versions", query.FlowID, url.PathEscape(key))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WorkflowFlowStateListResponse struct {
	Data []WorkflowFlowStateListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowStateListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowStateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowStateListResponseData struct {
	Key       string    `json:"key" api:"required"`
	SavedAt   time.Time `json:"savedAt" api:"required" format:"date-time"`
	SizeBytes int64     `json:"sizeBytes" api:"required"`
	// Newest saved version of this key
	Version int64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		SavedAt     respjson.Field
		SizeBytes   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowStateListResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowStateListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowStateDeleteResponse struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowStateDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowStateDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowStateListVersionsResponse struct {
	Data []WorkflowFlowStateListVersionsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowStateListVersionsResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowStateListVersionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowStateListVersionsResponseData struct {
	// Execution that saved this version
	ExecutionID string    `json:"executionId" api:"required" format:"uuid"`
	SavedAt     time.Time `json:"savedAt" api:"required" format:"date-time"`
	SizeBytes   int64     `json:"sizeBytes" api:"required"`
	Version     int64     `json:"version" api:"required"`
	// The saved JSON value
	Value any `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExecutionID respjson.Field
		SavedAt     respjson.Field
		SizeBytes   respjson.Field
		Version     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowStateListVersionsResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowStateListVersionsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowStateDeleteParams struct {
	FlowID string `path:"flowId" api:"required" format:"uuid" json:"-"`
	paramObj
}

type WorkflowFlowStateListVersionsParams struct {
	FlowID string `path:"flowId" api:"required" format:"uuid" json:"-"`
	paramObj
}
