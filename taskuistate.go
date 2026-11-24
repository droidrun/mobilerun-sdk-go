// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// Tasks API
//
// TaskUiStateService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskUiStateService] method instead.
type TaskUiStateService struct {
	options []option.RequestOption
}

// NewTaskUiStateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTaskUiStateService(opts ...option.RequestOption) (r TaskUiStateService) {
	r = TaskUiStateService{}
	r.options = opts
	return
}

// Get a specific UI state by index.
func (r *TaskUiStateService) Get(ctx context.Context, index int64, query TaskUiStateGetParams, opts ...option.RequestOption) (res *TaskUiStateGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.TaskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tasks/%s/ui_states/%v", query.TaskID, index)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all UI state URLs for a task.
func (r *TaskUiStateService) List(ctx context.Context, taskID string, opts ...option.RequestOption) (res *TaskUiStateListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tasks/%s/ui_states", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type TaskUiStateGetResponse struct {
	// The URL of the media
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskUiStateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskUiStateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskUiStateListResponse struct {
	// The list of media URLs
	URLs []string `json:"urls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URLs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskUiStateListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskUiStateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskUiStateGetParams struct {
	TaskID string `path:"task_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
