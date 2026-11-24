// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
)

// WorkflowService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	options       []option.RequestOption
	Triggers      WorkflowTriggerService
	ActionCatalog WorkflowActionCatalogService
	Actions       WorkflowActionService
	Flows         WorkflowFlowService
	Events        WorkflowEventService
	Executions    WorkflowExecutionService
	Timezones     WorkflowTimezoneService
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r WorkflowService) {
	r = WorkflowService{}
	r.options = opts
	r.Triggers = NewWorkflowTriggerService(opts...)
	r.ActionCatalog = NewWorkflowActionCatalogService(opts...)
	r.Actions = NewWorkflowActionService(opts...)
	r.Flows = NewWorkflowFlowService(opts...)
	r.Events = NewWorkflowEventService(opts...)
	r.Executions = NewWorkflowExecutionService(opts...)
	r.Timezones = NewWorkflowTimezoneService(opts...)
	return
}
