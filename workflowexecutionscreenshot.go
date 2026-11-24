// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
)

// WorkflowExecutionScreenshotService contains methods and other services that help
// with interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowExecutionScreenshotService] method instead.
type WorkflowExecutionScreenshotService struct {
	options []option.RequestOption
}

// NewWorkflowExecutionScreenshotService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkflowExecutionScreenshotService(opts ...option.RequestOption) (r WorkflowExecutionScreenshotService) {
	r = WorkflowExecutionScreenshotService{}
	r.options = opts
	return
}

// Redirects (302) to a freshly signed, time-limited image URL for a screenshot
// captured during this execution. Owner-gated: 404 if the execution or the
// screenshot does not belong to the caller.
func (r *WorkflowExecutionScreenshotService) Get(ctx context.Context, screenshotID string, query WorkflowExecutionScreenshotGetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.ExecutionID == "" {
		err = errors.New("missing required executionId parameter")
		return err
	}
	if screenshotID == "" {
		err = errors.New("missing required screenshotId parameter")
		return err
	}
	path := fmt.Sprintf("executions/%s/screenshots/%s", query.ExecutionID, screenshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

type WorkflowExecutionScreenshotGetParams struct {
	ExecutionID string `path:"executionId" api:"required" format:"uuid" json:"-"`
	paramObj
}
