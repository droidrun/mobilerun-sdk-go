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
// TaskScreenshotService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskScreenshotService] method instead.
type TaskScreenshotService struct {
	options []option.RequestOption
}

// NewTaskScreenshotService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTaskScreenshotService(opts ...option.RequestOption) (r TaskScreenshotService) {
	r = TaskScreenshotService{}
	r.options = opts
	return
}

// Get a specific screenshot by index.
func (r *TaskScreenshotService) Get(ctx context.Context, index int64, query TaskScreenshotGetParams, opts ...option.RequestOption) (res *TaskScreenshotGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.TaskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tasks/%s/screenshots/%v", query.TaskID, index)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all screenshot URLs for a task.
func (r *TaskScreenshotService) List(ctx context.Context, taskID string, opts ...option.RequestOption) (res *TaskScreenshotListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tasks/%s/screenshots", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type TaskScreenshotGetResponse struct {
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
func (r TaskScreenshotGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskScreenshotGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskScreenshotListResponse struct {
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
func (r TaskScreenshotListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskScreenshotListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskScreenshotGetParams struct {
	TaskID string `path:"task_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
