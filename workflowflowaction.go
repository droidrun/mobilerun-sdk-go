// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// WorkflowFlowActionService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowFlowActionService] method instead.
type WorkflowFlowActionService struct {
	options []option.RequestOption
}

// NewWorkflowFlowActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowFlowActionService(opts ...option.RequestOption) (r WorkflowFlowActionService) {
	r = WorkflowFlowActionService{}
	r.options = opts
	return
}

// Return the ordered list of actions attached to a flow, including any nested
// child actions. Returns 404 if the flow does not exist.
func (r *WorkflowFlowActionService) List(ctx context.Context, flowID string, opts ...option.RequestOption) (res *WorkflowFlowActionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/actions", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Append a single action to a flow at the given `position`, optionally nesting it
// under a `parentFlowActionId` or supplying its own `children`. Supports a
// `nameOverride`, param `overrides`, and `continueOnError`. Returns 404 if the
// flow does not exist.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). Replays with the same key and an identical body return the original
// 201; a changed body under the same key returns 422 `idempotency_key_reused`.
func (r *WorkflowFlowActionService) Add(ctx context.Context, flowID string, body WorkflowFlowActionAddParams, opts ...option.RequestOption) (res *WorkflowFlowActionAddResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/actions", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove a single action from a flow by its `flowActionId`. Returns 404 if the
// flow or flow action does not exist.
func (r *WorkflowFlowActionService) Remove(ctx context.Context, flowActionID string, body WorkflowFlowActionRemoveParams, opts ...option.RequestOption) (res *WorkflowFlowActionRemoveResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.FlowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	if flowActionID == "" {
		err = errors.New("missing required flowActionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/actions/%s", body.FlowID, flowActionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Replace a flow's entire action list with the supplied set (at least one
// required). Each action references an `actionId` and a unique `position`, and may
// include nested `children`, a `nameOverride`, param `overrides`, and a
// `continueOnError` flag. Returns 404 if the flow does not exist.
func (r *WorkflowFlowActionService) Replace(ctx context.Context, flowID string, body WorkflowFlowActionReplaceParams, opts ...option.RequestOption) (res *WorkflowFlowActionReplaceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/actions", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type FlowAction struct {
	ID                 string              `json:"id" api:"required" format:"uuid"`
	ActionID           string              `json:"actionId" api:"required" format:"uuid"`
	ContinueOnError    bool                `json:"continueOnError" api:"required"`
	CreatedAt          time.Time           `json:"createdAt" api:"required" format:"date-time"`
	FlowID             string              `json:"flowId" api:"required" format:"uuid"`
	Key                string              `json:"key" api:"required"`
	NameOverride       string              `json:"nameOverride" api:"required"`
	Overrides          FlowActionOverrides `json:"overrides" api:"required"`
	ParentFlowActionID string              `json:"parentFlowActionId" api:"required" format:"uuid"`
	Position           int64               `json:"position" api:"required"`
	// Selected for per-step recording under the flow's "selected_steps" recording
	// policy.
	RecordingEnabled bool `json:"recordingEnabled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		ActionID           respjson.Field
		ContinueOnError    respjson.Field
		CreatedAt          respjson.Field
		FlowID             respjson.Field
		Key                respjson.Field
		NameOverride       respjson.Field
		Overrides          respjson.Field
		ParentFlowActionID respjson.Field
		Position           respjson.Field
		RecordingEnabled   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlowAction) RawJSON() string { return r.JSON.raw }
func (r *FlowAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FlowActionOverrides struct {
	Params map[string]any `json:"params"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Params      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlowActionOverrides) RawJSON() string { return r.JSON.raw }
func (r *FlowActionOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActionListResponse struct {
	Data []FlowAction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowActionListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowActionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActionAddResponse struct {
	Data FlowAction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowActionAddResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowActionAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActionRemoveResponse struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowActionRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowActionRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActionReplaceResponse struct {
	Data []FlowAction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowActionReplaceResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowActionReplaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActionAddParams struct {
	ActionID           string            `json:"actionId" api:"required" format:"uuid"`
	Position           int64             `json:"position" api:"required"`
	ParentFlowActionID param.Opt[string] `json:"parentFlowActionId,omitzero" format:"uuid"`
	ContinueOnError    param.Opt[bool]   `json:"continueOnError,omitzero"`
	// Stable identifier used by template resolution v2+ to address this step as
	// {{key.field}}. Omitted keys are derived from the step name whenever a flow is
	// created. For template resolution v3, replacing an existing flow's full action
	// tree requires every step to carry an explicit key — a missing key is rejected
	// there, not derived, so a name change can never silently move a step's key.
	Key               param.Opt[string]                    `json:"key,omitzero"`
	NameOverride      param.Opt[string]                    `json:"nameOverride,omitzero"`
	RecordingEnabled  param.Opt[bool]                      `json:"recordingEnabled,omitzero"`
	OverridesProperty WorkflowFlowActionAddParamsOverrides `json:"overrides,omitzero"`
	Children          []WorkflowFlowActionAddParamsChild   `json:"children,omitzero"`
	paramObj
}

func (r WorkflowFlowActionAddParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowActionAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowActionAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Position are required.
type WorkflowFlowActionAddParamsChild struct {
	ActionID        string          `json:"actionId" api:"required" format:"uuid"`
	Position        int64           `json:"position" api:"required"`
	ContinueOnError param.Opt[bool] `json:"continueOnError,omitzero"`
	// Stable identifier used by template resolution v2+ to address this step as
	// {{key.field}}. Omitted keys are derived from the step name whenever a flow is
	// created. For template resolution v3, replacing an existing flow's full action
	// tree requires every step to carry an explicit key — a missing key is rejected
	// there, not derived, so a name change can never silently move a step's key.
	Key               param.Opt[string]                         `json:"key,omitzero"`
	NameOverride      param.Opt[string]                         `json:"nameOverride,omitzero"`
	RecordingEnabled  param.Opt[bool]                           `json:"recordingEnabled,omitzero"`
	OverridesProperty WorkflowFlowActionAddParamsChildOverrides `json:"overrides,omitzero"`
	paramObj
}

func (r WorkflowFlowActionAddParamsChild) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowActionAddParamsChild
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowActionAddParamsChild) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActionAddParamsChildOverrides struct {
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r WorkflowFlowActionAddParamsChildOverrides) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowActionAddParamsChildOverrides
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowActionAddParamsChildOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActionAddParamsOverrides struct {
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r WorkflowFlowActionAddParamsOverrides) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowActionAddParamsOverrides
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowActionAddParamsOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActionRemoveParams struct {
	FlowID string `path:"flowId" api:"required" format:"uuid" json:"-"`
	paramObj
}

type WorkflowFlowActionReplaceParams struct {
	Actions []WorkflowFlowActionReplaceParamsAction `json:"actions,omitzero" api:"required"`
	paramObj
}

func (r WorkflowFlowActionReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowActionReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowActionReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Position are required.
type WorkflowFlowActionReplaceParamsAction struct {
	ActionID        string          `json:"actionId" api:"required" format:"uuid"`
	Position        int64           `json:"position" api:"required"`
	ContinueOnError param.Opt[bool] `json:"continueOnError,omitzero"`
	// Stable identifier used by template resolution v2+ to address this step as
	// {{key.field}}. Omitted keys are derived from the step name whenever a flow is
	// created. For template resolution v3, replacing an existing flow's full action
	// tree requires every step to carry an explicit key — a missing key is rejected
	// there, not derived, so a name change can never silently move a step's key.
	Key               param.Opt[string]                              `json:"key,omitzero"`
	NameOverride      param.Opt[string]                              `json:"nameOverride,omitzero"`
	RecordingEnabled  param.Opt[bool]                                `json:"recordingEnabled,omitzero"`
	OverridesProperty WorkflowFlowActionReplaceParamsActionOverrides `json:"overrides,omitzero"`
	Children          []WorkflowFlowActionReplaceParamsActionChild   `json:"children,omitzero"`
	paramObj
}

func (r WorkflowFlowActionReplaceParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowActionReplaceParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowActionReplaceParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Position are required.
type WorkflowFlowActionReplaceParamsActionChild struct {
	ActionID        string          `json:"actionId" api:"required" format:"uuid"`
	Position        int64           `json:"position" api:"required"`
	ContinueOnError param.Opt[bool] `json:"continueOnError,omitzero"`
	// Stable identifier used by template resolution v2+ to address this step as
	// {{key.field}}. Omitted keys are derived from the step name whenever a flow is
	// created. For template resolution v3, replacing an existing flow's full action
	// tree requires every step to carry an explicit key — a missing key is rejected
	// there, not derived, so a name change can never silently move a step's key.
	Key               param.Opt[string]                                   `json:"key,omitzero"`
	NameOverride      param.Opt[string]                                   `json:"nameOverride,omitzero"`
	RecordingEnabled  param.Opt[bool]                                     `json:"recordingEnabled,omitzero"`
	OverridesProperty WorkflowFlowActionReplaceParamsActionChildOverrides `json:"overrides,omitzero"`
	paramObj
}

func (r WorkflowFlowActionReplaceParamsActionChild) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowActionReplaceParamsActionChild
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowActionReplaceParamsActionChild) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActionReplaceParamsActionChildOverrides struct {
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r WorkflowFlowActionReplaceParamsActionChildOverrides) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowActionReplaceParamsActionChildOverrides
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowActionReplaceParamsActionChildOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActionReplaceParamsActionOverrides struct {
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r WorkflowFlowActionReplaceParamsActionOverrides) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowActionReplaceParamsActionOverrides
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowActionReplaceParamsActionOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
