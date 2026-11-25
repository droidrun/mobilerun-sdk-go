// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// WorkflowEventService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowEventService] method instead.
type WorkflowEventService struct {
	options []option.RequestOption
}

// NewWorkflowEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowEventService(opts ...option.RequestOption) (r WorkflowEventService) {
	r = WorkflowEventService{}
	r.options = opts
	return
}

// Simulate an event against all configured flows. Returns which flows would match
// and what actions would run, without storing the event or enqueuing jobs.
func (r *WorkflowEventService) DryRun(ctx context.Context, body WorkflowEventDryRunParams, opts ...option.RequestOption) (res *WorkflowEventDryRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "events/dry-run"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Ingest an event for trigger evaluation. Returns immediately with 202 Accepted.
func (r *WorkflowEventService) Ingest(ctx context.Context, body WorkflowEventIngestParams, opts ...option.RequestOption) (res *WorkflowEventIngestResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "events/ingest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WorkflowEventDryRunResponse struct {
	Data WorkflowEventDryRunResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventDryRunResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEventDryRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventDryRunResponseData struct {
	MatchedFlows []WorkflowEventDryRunResponseDataMatchedFlow `json:"matchedFlows" api:"required"`
	Validation   WorkflowEventDryRunResponseDataValidation    `json:"validation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MatchedFlows respjson.Field
		Validation   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventDryRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEventDryRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventDryRunResponseDataMatchedFlow struct {
	Actions   []WorkflowEventDryRunResponseDataMatchedFlowAction `json:"actions" api:"required"`
	Flow      WorkflowEventDryRunResponseDataMatchedFlowFlow     `json:"flow" api:"required"`
	Gates     WorkflowEventDryRunResponseDataMatchedFlowGates    `json:"gates" api:"required"`
	Trigger   Trigger                                            `json:"trigger" api:"required"`
	WouldFire bool                                               `json:"wouldFire" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions     respjson.Field
		Flow        respjson.Field
		Gates       respjson.Field
		Trigger     respjson.Field
		WouldFire   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventDryRunResponseDataMatchedFlow) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEventDryRunResponseDataMatchedFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventDryRunResponseDataMatchedFlowAction struct {
	ContinueOnError bool   `json:"continueOnError" api:"required"`
	FlowActionID    string `json:"flowActionId" api:"required" format:"uuid"`
	Key             string `json:"key" api:"required"`
	// Deprecated: use uses.
	//
	// Deprecated: deprecated
	Method           string `json:"method" api:"required"`
	Name             string `json:"name" api:"required"`
	RecordingEnabled bool   `json:"recordingEnabled" api:"required"`
	// Deprecated: use uses.
	//
	// Any of "tasks_api", "devices_api", "agents_api", "webhooks", "integrations_api",
	// "triggers_api", "sandbox_api".
	//
	// Deprecated: deprecated
	Service string `json:"service" api:"required"`
	// Authored step name; null when this step has no name. Includes hidden executable
	// steps.
	Uses string `json:"uses" api:"required"`
	// Nested child actions (loop/branch bodies), each the same shape as a
	// ResolvedAction, including nullable uses and deprecated service/method.
	Children []any          `json:"children"`
	Params   map[string]any `json:"params"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContinueOnError  respjson.Field
		FlowActionID     respjson.Field
		Key              respjson.Field
		Method           respjson.Field
		Name             respjson.Field
		RecordingEnabled respjson.Field
		Service          respjson.Field
		Uses             respjson.Field
		Children         respjson.Field
		Params           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventDryRunResponseDataMatchedFlowAction) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEventDryRunResponseDataMatchedFlowAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventDryRunResponseDataMatchedFlowFlow struct {
	ID                  string    `json:"id" api:"required" format:"uuid"`
	AgentID             string    `json:"agentId" api:"required" format:"uuid"`
	ArchivedAt          time.Time `json:"archivedAt" api:"required" format:"date-time"`
	BlockedAt           time.Time `json:"blockedAt" api:"required" format:"date-time"`
	ConsecutiveFailures int64     `json:"consecutiveFailures" api:"required"`
	// Any of "flow", "device".
	CooldownScope   string                                                 `json:"cooldownScope" api:"required"`
	CooldownSeconds int64                                                  `json:"cooldownSeconds" api:"required"`
	CreatedAt       time.Time                                              `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy       string                                                 `json:"createdBy" api:"required" format:"uuid"`
	Delivery        WorkflowEventDryRunResponseDataMatchedFlowFlowDelivery `json:"delivery" api:"required"`
	Description     string                                                 `json:"description" api:"required"`
	DeviceIDs       []string                                               `json:"deviceIds" api:"required" format:"uuid"`
	// Compatibility projection of lifecycleStatus; true only when lifecycleStatus is
	// enabled.
	Enabled                 bool      `json:"enabled" api:"required"`
	HealthMonitoringEnabled bool      `json:"healthMonitoringEnabled" api:"required"`
	LastFailureAt           time.Time `json:"lastFailureAt" api:"required" format:"date-time"`
	// Curated flow failure class. device_unavailable blocks immediately for a
	// terminated device or a device confirmed gone; event deviceState="gone" is
	// triggers-defined for a deleted/not-found device.
	//
	// Any of "device_not_found", "device_unavailable", "permission_denied",
	// "client_error", "transient", "logic", "invalid_config".
	LastFailureCode string    `json:"lastFailureCode" api:"required"`
	LastTriggeredAt time.Time `json:"lastTriggeredAt" api:"required" format:"date-time"`
	// Any of "enabled", "disabled", "archived".
	LifecycleStatus string `json:"lifecycleStatus" api:"required"`
	Name            string `json:"name" api:"required"`
	NotifyOnFailure bool   `json:"notifyOnFailure" api:"required"`
	NotifyOnSuccess bool   `json:"notifyOnSuccess" api:"required"`
	NotifyWebhookID string `json:"notifyWebhookId" api:"required" format:"uuid"`
	OwnerID         string `json:"ownerId" api:"required" format:"uuid"`
	// Defaults to final.
	//
	// Any of "final", "off".
	OwnerWake string `json:"ownerWake" api:"required"`
	// Deprecated: use recordingPolicy.mode ("flow" = recordingEnabled=true, "off" =
	// recordingEnabled=false).
	//
	// Deprecated: deprecated
	RecordingEnabled bool            `json:"recordingEnabled" api:"required"`
	RecordingPolicy  RecordingPolicy `json:"recordingPolicy" api:"required"`
	// Effective routine result delivery: the stored setting, else message for a chat
	// sink, else none.
	//
	// Any of "none", "message", "wake".
	ResultDelivery         string `json:"resultDelivery" api:"required"`
	SaveSandboxLogs        bool   `json:"saveSandboxLogs" api:"required"`
	SelfHealingEnabled     bool   `json:"selfHealingEnabled" api:"required"`
	SelfHealingMaxAttempts int64  `json:"selfHealingMaxAttempts" api:"required"`
	// Any of "healthy", "failing", "blocked".
	Status string `json:"status" api:"required"`
	// Template-resolver semantics this flow runs under. 1 = legacy
	// (missing/forbidden/null all resolve to ”). 2 = typed (missing/forbidden throw,
	// a whole-token null stays JSON null). 3 = typed, key-based (steps are addressed
	// as {{steps.<key>...}} instead of by name; trigger.payload is not available).
	// Existing flows stay 1; new flows default to the current version.
	TemplateResolutionVersion int64     `json:"templateResolutionVersion" api:"required"`
	TriggerID                 string    `json:"triggerId" api:"required" format:"uuid"`
	UpdatedAt                 time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Deprecated: use ownerId (tenancy) / createdBy (actor).
	//
	// Deprecated: deprecated
	UserID string `json:"userId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		AgentID                   respjson.Field
		ArchivedAt                respjson.Field
		BlockedAt                 respjson.Field
		ConsecutiveFailures       respjson.Field
		CooldownScope             respjson.Field
		CooldownSeconds           respjson.Field
		CreatedAt                 respjson.Field
		CreatedBy                 respjson.Field
		Delivery                  respjson.Field
		Description               respjson.Field
		DeviceIDs                 respjson.Field
		Enabled                   respjson.Field
		HealthMonitoringEnabled   respjson.Field
		LastFailureAt             respjson.Field
		LastFailureCode           respjson.Field
		LastTriggeredAt           respjson.Field
		LifecycleStatus           respjson.Field
		Name                      respjson.Field
		NotifyOnFailure           respjson.Field
		NotifyOnSuccess           respjson.Field
		NotifyWebhookID           respjson.Field
		OwnerID                   respjson.Field
		OwnerWake                 respjson.Field
		RecordingEnabled          respjson.Field
		RecordingPolicy           respjson.Field
		ResultDelivery            respjson.Field
		SaveSandboxLogs           respjson.Field
		SelfHealingEnabled        respjson.Field
		SelfHealingMaxAttempts    respjson.Field
		Status                    respjson.Field
		TemplateResolutionVersion respjson.Field
		TriggerID                 respjson.Field
		UpdatedAt                 respjson.Field
		UserID                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventDryRunResponseDataMatchedFlowFlow) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEventDryRunResponseDataMatchedFlowFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventDryRunResponseDataMatchedFlowFlowDelivery struct {
	// Any of "mobilerun", "one_drive", "google_drive".
	Destination string `json:"destination" api:"required"`
	// Absolute destination folder. For destination mobilerun, a folder turns on saving
	// a copy of the run's recording, screenshots and files.upload outputs in Mobilerun
	// storage under that folder (counts against storage). May contain {{meta.<name>}}
	// (for example {{meta.shortDate}} = YYMMDD, {{meta.fileTime}} = HH-mm, in the
	// trigger timezone) and {{event.payload.<field>}} placeholders, resolved once per
	// run so every artifact of that run lands in the same folder. Each substituted
	// value must be a single safe folder name (no "/", "\", "..", control characters);
	// otherwise delivery fails with invalid_delivery_folder. For OneDrive and Google
	// Drive, screenshots are delivered directly into the resolved folder with a
	// placeholder folder, and into a per-run "<flow>-<timestamp>-<id>-screenshots"
	// subfolder with a static folder. Mobilerun copies always go directly into the
	// folder.
	Folder      string                                                          `json:"folder"`
	Recording   WorkflowEventDryRunResponseDataMatchedFlowFlowDeliveryRecording `json:"recording"`
	Screenshots any                                                             `json:"screenshots"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Destination respjson.Field
		Folder      respjson.Field
		Recording   respjson.Field
		Screenshots respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventDryRunResponseDataMatchedFlowFlowDelivery) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEventDryRunResponseDataMatchedFlowFlowDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventDryRunResponseDataMatchedFlowFlowDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventDryRunResponseDataMatchedFlowFlowDeliveryRecording) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowEventDryRunResponseDataMatchedFlowFlowDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventDryRunResponseDataMatchedFlowGates struct {
	Blocked        bool     `json:"blocked" api:"required"`
	CooldownActive bool     `json:"cooldownActive" api:"required"`
	DeviceAttached bool     `json:"deviceAttached" api:"required"`
	DeviceIDs      []string `json:"deviceIds" api:"required" format:"uuid"`
	Enabled        bool     `json:"enabled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blocked        respjson.Field
		CooldownActive respjson.Field
		DeviceAttached respjson.Field
		DeviceIDs      respjson.Field
		Enabled        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventDryRunResponseDataMatchedFlowGates) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEventDryRunResponseDataMatchedFlowGates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventDryRunResponseDataValidation struct {
	Valid  bool                                             `json:"valid" api:"required"`
	Errors []WorkflowEventDryRunResponseDataValidationError `json:"errors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Valid       respjson.Field
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventDryRunResponseDataValidation) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEventDryRunResponseDataValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventDryRunResponseDataValidationError struct {
	Field   string `json:"field" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventDryRunResponseDataValidationError) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEventDryRunResponseDataValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventIngestResponse struct {
	EventID string `json:"eventId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventIngestResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEventIngestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventDryRunParams struct {
	EventType string            `json:"eventType" api:"required"`
	DeviceID  param.Opt[string] `json:"deviceId,omitzero" format:"uuid"`
	Payload   map[string]any    `json:"payload,omitzero"`
	paramObj
}

func (r WorkflowEventDryRunParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowEventDryRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowEventDryRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowEventIngestParams struct {
	EventType string            `json:"eventType" api:"required"`
	DeviceID  param.Opt[string] `json:"deviceId,omitzero" format:"uuid"`
	Payload   map[string]any    `json:"payload,omitzero"`
	paramObj
}

func (r WorkflowEventIngestParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowEventIngestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowEventIngestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
