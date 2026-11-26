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
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// Tasks API
//
// TaskApprovalService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskApprovalService] method instead.
type TaskApprovalService struct {
	options []option.RequestOption
}

// NewTaskApprovalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTaskApprovalService(opts ...option.RequestOption) (r TaskApprovalService) {
	r = TaskApprovalService{}
	r.options = opts
	return
}

// Approve or deny the approval a running task is waiting for (see
// `pendingApproval` on the task). Deny ends the task as failed.
func (r *TaskApprovalService) Decide(ctx context.Context, approvalID string, params TaskApprovalDecideParams, opts ...option.RequestOption) (res *TaskApprovalDecideResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.TaskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	if approvalID == "" {
		err = errors.New("missing required approval_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tasks/%s/approvals/%s", params.TaskID, approvalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type TaskApprovalDecideResponse struct {
	// The recorded decision
	//
	// Any of "approve", "deny".
	Decision TaskApprovalDecideResponseDecision `json:"decision" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Decision    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskApprovalDecideResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskApprovalDecideResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The recorded decision
type TaskApprovalDecideResponseDecision string

const (
	TaskApprovalDecideResponseDecisionApprove TaskApprovalDecideResponseDecision = "approve"
	TaskApprovalDecideResponseDecisionDeny    TaskApprovalDecideResponseDecision = "deny"
)

type TaskApprovalDecideParams struct {
	TaskID string `path:"task_id" api:"required" format:"uuid" json:"-"`
	// Approve lets the task continue; deny ends it as failed
	//
	// Any of "approve", "deny".
	Decision TaskApprovalDecideParamsDecision `json:"decision,omitzero" api:"required"`
	paramObj
}

func (r TaskApprovalDecideParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskApprovalDecideParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskApprovalDecideParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Approve lets the task continue; deny ends it as failed
type TaskApprovalDecideParamsDecision string

const (
	TaskApprovalDecideParamsDecisionApprove TaskApprovalDecideParamsDecision = "approve"
	TaskApprovalDecideParamsDecisionDeny    TaskApprovalDecideParamsDecision = "deny"
)
