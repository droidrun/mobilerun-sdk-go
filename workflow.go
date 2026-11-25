// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"encoding/json"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
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

type RecordingPolicy struct {
	// Any of "off", "flow", "selected_steps".
	Mode RecordingPolicyMode `json:"mode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingPolicy) RawJSON() string { return r.JSON.raw }
func (r *RecordingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingPolicyMode string

const (
	RecordingPolicyModeOff           RecordingPolicyMode = "off"
	RecordingPolicyModeFlow          RecordingPolicyMode = "flow"
	RecordingPolicyModeSelectedSteps RecordingPolicyMode = "selected_steps"
)

type ScheduleRule struct {
	// Any of "once", "cron", "recurring".
	Type ScheduleRuleType `json:"type" api:"required"`
	// ISO 8601 datetime (for type=once)
	DateTime string `json:"dateTime"`
	// Cron expression (for type=cron)
	Expression string `json:"expression"`
	// Optional per-occurrence random window around the nominal schedule time
	Jitter ScheduleRuleJitter `json:"jitter"`
	// RRULE string (for type=recurring)
	Rrule string `json:"rrule"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		DateTime    respjson.Field
		Expression  respjson.Field
		Jitter      respjson.Field
		Rrule       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleRule) RawJSON() string { return r.JSON.raw }
func (r *ScheduleRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ScheduleRule to a ScheduleRuleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ScheduleRuleParam.Overrides()
func (r ScheduleRule) ToParam() ScheduleRuleParam {
	return param.Override[ScheduleRuleParam](json.RawMessage(r.RawJSON()))
}

type ScheduleRuleType string

const (
	ScheduleRuleTypeOnce      ScheduleRuleType = "once"
	ScheduleRuleTypeCron      ScheduleRuleType = "cron"
	ScheduleRuleTypeRecurring ScheduleRuleType = "recurring"
)

// Optional per-occurrence random window around the nominal schedule time
type ScheduleRuleJitter struct {
	AfterMinutes  int64 `json:"afterMinutes"`
	BeforeMinutes int64 `json:"beforeMinutes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterMinutes  respjson.Field
		BeforeMinutes respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleRuleJitter) RawJSON() string { return r.JSON.raw }
func (r *ScheduleRuleJitter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ScheduleRuleParam struct {
	// Any of "once", "cron", "recurring".
	Type ScheduleRuleType `json:"type,omitzero" api:"required"`
	// ISO 8601 datetime (for type=once)
	DateTime param.Opt[string] `json:"dateTime,omitzero"`
	// Cron expression (for type=cron)
	Expression param.Opt[string] `json:"expression,omitzero"`
	// RRULE string (for type=recurring)
	Rrule param.Opt[string] `json:"rrule,omitzero"`
	// Optional per-occurrence random window around the nominal schedule time
	Jitter ScheduleRuleJitterParam `json:"jitter,omitzero"`
	paramObj
}

func (r ScheduleRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow ScheduleRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScheduleRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional per-occurrence random window around the nominal schedule time
type ScheduleRuleJitterParam struct {
	AfterMinutes  param.Opt[int64] `json:"afterMinutes,omitzero"`
	BeforeMinutes param.Opt[int64] `json:"beforeMinutes,omitzero"`
	paramObj
}

func (r ScheduleRuleJitterParam) MarshalJSON() (data []byte, err error) {
	type shadow ScheduleRuleJitterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScheduleRuleJitterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Trigger struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "event", "schedule", "custom".
	Activation          TriggerActivation `json:"activation" api:"required"`
	CreatedAt           time.Time         `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy           string            `json:"createdBy" api:"required" format:"uuid"`
	CustomPayloadSchema map[string]any    `json:"customPayloadSchema" api:"required"`
	Description         string            `json:"description" api:"required"`
	EventType           string            `json:"eventType" api:"required"`
	Name                string            `json:"name" api:"required"`
	OwnerID             string            `json:"ownerId" api:"required" format:"uuid"`
	Payload             map[string]any    `json:"payload" api:"required"`
	ScheduleRule        ScheduleRule      `json:"scheduleRule" api:"required"`
	Timezone            string            `json:"timezone" api:"required"`
	UpdatedAt           time.Time         `json:"updatedAt" api:"required" format:"date-time"`
	// Deprecated: use ownerId (tenancy) / createdBy (actor).
	//
	// Deprecated: deprecated
	UserID       string `json:"userId" api:"required" format:"uuid"`
	Conditions   any    `json:"conditions"`
	NextFireTime string `json:"nextFireTime" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Activation          respjson.Field
		CreatedAt           respjson.Field
		CreatedBy           respjson.Field
		CustomPayloadSchema respjson.Field
		Description         respjson.Field
		EventType           respjson.Field
		Name                respjson.Field
		OwnerID             respjson.Field
		Payload             respjson.Field
		ScheduleRule        respjson.Field
		Timezone            respjson.Field
		UpdatedAt           respjson.Field
		UserID              respjson.Field
		Conditions          respjson.Field
		NextFireTime        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Trigger) RawJSON() string { return r.JSON.raw }
func (r *Trigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TriggerActivation string

const (
	TriggerActivationEvent    TriggerActivation = "event"
	TriggerActivationSchedule TriggerActivation = "schedule"
	TriggerActivationCustom   TriggerActivation = "custom"
)
