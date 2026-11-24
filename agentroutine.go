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

// AgentRoutineService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentRoutineService] method instead.
type AgentRoutineService struct {
	options []option.RequestOption
}

// NewAgentRoutineService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentRoutineService(opts ...option.RequestOption) (r AgentRoutineService) {
	r = AgentRoutineService{}
	r.options = opts
	return
}

// Forwards the scheduler CreateFlowBody unchanged and assigns the Agent selected
// in the path. Requires Idempotency-Key; reuse it after a transient failure.
func (r *AgentRoutineService) New(ctx context.Context, agentID string, body AgentRoutineNewParams, opts ...option.RequestOption) (res *AgentRoutineNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/routines", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The source must belong to the authenticated owner and be assigned to the same
// Agent or unassigned. The target is the Agent in the path. Requires
// Idempotency-Key.
func (r *AgentRoutineService) Clone(ctx context.Context, routineID string, params AgentRoutineCloneParams, opts ...option.RequestOption) (res *AgentRoutineCloneResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.AgentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if routineID == "" {
		err = errors.New("missing required routineId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/routines/%s/clone", params.AgentID, routineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type AgentRoutineNewResponse struct {
	Data AgentRoutineNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRoutineNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRoutineNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRoutineNewResponseData struct {
	ID          string         `json:"id" api:"required"`
	AgentID     string         `json:"agentId" api:"required" format:"uuid"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AgentID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRoutineNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *AgentRoutineNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRoutineCloneResponse struct {
	Data AgentRoutineCloneResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRoutineCloneResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRoutineCloneResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRoutineCloneResponseData struct {
	ID          string         `json:"id" api:"required"`
	AgentID     string         `json:"agentId" api:"required" format:"uuid"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AgentID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRoutineCloneResponseData) RawJSON() string { return r.JSON.raw }
func (r *AgentRoutineCloneResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRoutineNewParams struct {
	Actions                 []AgentRoutineNewParamsAction `json:"actions,omitzero" api:"required"`
	Name                    string                        `json:"name" api:"required"`
	TriggerID               string                        `json:"triggerId" api:"required" format:"uuid"`
	CooldownSeconds         param.Opt[int64]              `json:"cooldownSeconds,omitzero"`
	NotifyWebhookID         param.Opt[string]             `json:"notifyWebhookId,omitzero" format:"uuid"`
	Description             param.Opt[string]             `json:"description,omitzero"`
	Enabled                 param.Opt[bool]               `json:"enabled,omitzero"`
	HealthMonitoringEnabled param.Opt[bool]               `json:"healthMonitoringEnabled,omitzero"`
	NotifyOnFailure         param.Opt[bool]               `json:"notifyOnFailure,omitzero"`
	NotifyOnSuccess         param.Opt[bool]               `json:"notifyOnSuccess,omitzero"`
	RecordingEnabled        param.Opt[bool]               `json:"recordingEnabled,omitzero"`
	SelfHealingEnabled      param.Opt[bool]               `json:"selfHealingEnabled,omitzero"`
	SelfHealingMaxAttempts  param.Opt[int64]              `json:"selfHealingMaxAttempts,omitzero"`
	// Any of "flow", "device".
	CooldownScope   AgentRoutineNewParamsCooldownScope   `json:"cooldownScope,omitzero"`
	Delivery        AgentRoutineNewParamsDelivery        `json:"delivery,omitzero"`
	DeviceIDs       []string                             `json:"deviceIds,omitzero" format:"uuid"`
	RecordingPolicy AgentRoutineNewParamsRecordingPolicy `json:"recordingPolicy,omitzero"`
	paramObj
}

func (r AgentRoutineNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentRoutineNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRoutineNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Position are required.
type AgentRoutineNewParamsAction struct {
	ActionID    string         `json:"actionId" api:"required" format:"uuid"`
	Position    int64          `json:"position" api:"required"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r AgentRoutineNewParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow AgentRoutineNewParamsAction
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *AgentRoutineNewParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRoutineNewParamsCooldownScope string

const (
	AgentRoutineNewParamsCooldownScopeFlow   AgentRoutineNewParamsCooldownScope = "flow"
	AgentRoutineNewParamsCooldownScopeDevice AgentRoutineNewParamsCooldownScope = "device"
)

// The property Destination is required.
type AgentRoutineNewParamsDelivery struct {
	// Any of "one_drive", "google_drive", "mobilerun".
	Destination string                                 `json:"destination,omitzero" api:"required"`
	Folder      param.Opt[string]                      `json:"folder,omitzero"`
	Recording   AgentRoutineNewParamsDeliveryRecording `json:"recording,omitzero"`
	Screenshots any                                    `json:"screenshots,omitzero"`
	paramObj
}

func (r AgentRoutineNewParamsDelivery) MarshalJSON() (data []byte, err error) {
	type shadow AgentRoutineNewParamsDelivery
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRoutineNewParamsDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentRoutineNewParamsDelivery](
		"destination", "one_drive", "google_drive", "mobilerun",
	)
}

// The property Filename is required.
type AgentRoutineNewParamsDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	paramObj
}

func (r AgentRoutineNewParamsDeliveryRecording) MarshalJSON() (data []byte, err error) {
	type shadow AgentRoutineNewParamsDeliveryRecording
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRoutineNewParamsDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Mode is required.
type AgentRoutineNewParamsRecordingPolicy struct {
	// Any of "off", "flow", "selected_steps".
	Mode string `json:"mode,omitzero" api:"required"`
	paramObj
}

func (r AgentRoutineNewParamsRecordingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow AgentRoutineNewParamsRecordingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRoutineNewParamsRecordingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentRoutineNewParamsRecordingPolicy](
		"mode", "off", "flow", "selected_steps",
	)
}

type AgentRoutineCloneParams struct {
	AgentID   string            `path:"agentId" api:"required" format:"uuid" json:"-"`
	Name      param.Opt[string] `json:"name,omitzero"`
	DeviceIDs []string          `json:"deviceIds,omitzero" format:"uuid"`
	paramObj
}

func (r AgentRoutineCloneParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentRoutineCloneParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRoutineCloneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
