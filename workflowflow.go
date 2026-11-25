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
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/pagination"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// WorkflowFlowService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowFlowService] method instead.
type WorkflowFlowService struct {
	options       []option.RequestOption
	ScheduledRuns WorkflowFlowScheduledRunService
	Actions       WorkflowFlowActionService
	State         WorkflowFlowStateService
}

// NewWorkflowFlowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowFlowService(opts ...option.RequestOption) (r WorkflowFlowService) {
	r = WorkflowFlowService{}
	r.options = opts
	r.ScheduledRuns = NewWorkflowFlowScheduledRunService(opts...)
	r.Actions = NewWorkflowFlowActionService(opts...)
	r.State = NewWorkflowFlowStateService(opts...)
	return
}

// Create a flow that binds a trigger (`triggerId`) to an ordered list of actions.
// The list may be empty or omitted: a routine without steps is stored disabled and
// cannot be enabled, run or verified until it has a step. Enabling it returns 400
// with "Add a step before enabling this routine". Optional settings include target
// `deviceIds`, a cooldown (`cooldownSeconds`/`cooldownScope`), and webhook
// notifications on success or failure.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). Replays with the same key and an identical body return the original
// 201; a changed body under the same key returns 422 `idempotency_key_reused`.
func (r *WorkflowFlowService) New(ctx context.Context, body WorkflowFlowNewParams, opts ...option.RequestOption) (res *WorkflowFlowNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "flows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single flow by its ID, including its trigger binding, configuration, and
// current status. Returns 404 if no flow matches.
func (r *WorkflowFlowService) Get(ctx context.Context, flowID string, opts ...option.RequestOption) (res *WorkflowFlowGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a flow's settings — name, trigger binding, enabled state,
// target devices, cooldown, or notifications; all fields are optional. The owner
// can enable a routine directly with `enabled: true` or
// `lifecycleStatus: "enabled"`, without verification; at least one step is
// required. For the verified path used by agent deploys, use POST
// /flows/{flowId}/verify then POST /flows/{flowId}/activate instead. Actions are
// managed through the flow-actions endpoints, not here. Returns 404 if the flow
// does not exist.
func (r *WorkflowFlowService) Update(ctx context.Context, flowID string, body WorkflowFlowUpdateParams, opts ...option.RequestOption) (res *WorkflowFlowUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Return a paginated list of flows. Supports filtering by `agentId` (UUID or
// `unassigned`; omit for all buckets), `triggerId`, `enabled`, one or more health
// `status` values (healthy, failing, blocked), `mine` (flows created by the
// calling actor), `createdBy` (flows created by a given actor id — mutually
// exclusive with `mine`), plus free-text `search` and ordering.
func (r *WorkflowFlowService) List(ctx context.Context, query WorkflowFlowListParams, opts ...option.RequestOption) (res *pagination.Page[WorkflowFlowListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "flows"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Return a paginated list of flows. Supports filtering by `agentId` (UUID or
// `unassigned`; omit for all buckets), `triggerId`, `enabled`, one or more health
// `status` values (healthy, failing, blocked), `mine` (flows created by the
// calling actor), `createdBy` (flows created by a given actor id — mutually
// exclusive with `mine`), plus free-text `search` and ordering.
func (r *WorkflowFlowService) ListAutoPaging(ctx context.Context, query WorkflowFlowListParams, opts ...option.RequestOption) *pagination.PageAutoPager[WorkflowFlowListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Terminally archive a flow by its ID. Archived flows cannot be restored and are
// hidden from customer reads. Repeating the request is idempotent for the owner.
// Files saved by its runs are deleted with it.
func (r *WorkflowFlowService) Delete(ctx context.Context, flowID string, opts ...option.RequestOption) (res *WorkflowFlowDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// The verified activation path used by agent deploys: first run POST
// /flows/{flowId}/verify, then activate the disabled flow only when the supplied
// verification execution succeeded and the flow graph has not changed since
// verification. Repeating activation is idempotent for an already-enabled flow.
// Alternatively, the owner can enable a routine with at least one step directly
// through PATCH /flows/{flowId} using `enabled: true` or
// `lifecycleStatus: "enabled"`, without verification.
func (r *WorkflowFlowService) Activate(ctx context.Context, flowID string, body WorkflowFlowActivateParams, opts ...option.RequestOption) (res *WorkflowFlowActivateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/activate", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the environment routine limit and remaining capacity for this owner and
// agent bucket, counting enabled and disabled non-deleted flows. Omit agentId (or
// use unassigned) for the unassigned bucket. This is advisory; create and clone
// enforce admission under a bucket lock.
func (r *WorkflowFlowService) Capacity(ctx context.Context, query WorkflowFlowCapacityParams, opts ...option.RequestOption) (res *WorkflowFlowCapacityResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "flows/capacity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Create a copy of an existing flow, including its actions and settings. The
// optional body can override the new flow's `name` and target `deviceIds`. Returns
// 404 if the source flow does not exist.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). Replays with the same key and an identical body return the original
// 201; a changed body under the same key returns 422 `idempotency_key_reused`.
func (r *WorkflowFlowService) Clone(ctx context.Context, flowID string, body WorkflowFlowCloneParams, opts ...option.RequestOption) (res *WorkflowFlowCloneResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/clone", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Atomically creates the trigger, actions and disabled flow. Body: {document}.
// Optional Idempotency-Key fingerprints the canonical document; identical retries
// reload the resource, changed documents return idempotency_key_reused. The owner
// can enable a routine with at least one step directly through PATCH
// /flows/{flowId} using enabled: true or lifecycleStatus: "enabled", without
// verification. For the verified path used by agent deploys, call POST
// /flows/{flowId}/verify then POST /flows/{flowId}/activate; activation requires a
// successful verification of the current graph.
func (r *WorkflowFlowService) NewFromDocument(ctx context.Context, body WorkflowFlowNewFromDocumentParams, opts ...option.RequestOption) (res *WorkflowFlowNewFromDocumentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "flows/document"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Return the recording/delivery readiness for a flow (recording mode, whether
// delivery can include the recording, whether a files.upload action exists, the
// first direct OneDrive/Google Drive upload step if any) plus, per destination,
// whether the flow owner has an active connection. Returns 404 if the flow does
// not exist.
func (r *WorkflowFlowService) DeliveryOptions(ctx context.Context, flowID string, opts ...option.RequestOption) (res *WorkflowFlowDeliveryOptionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/delivery-options", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The optional agent_routine profile lists only agent-eligible steps. This is
// authoring guidance; backend validation remains authoritative.
func (r *WorkflowFlowService) DocumentSchema(ctx context.Context, query WorkflowFlowDocumentSchemaParams, opts ...option.RequestOption) (res *WorkflowFlowDocumentSchemaResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "flows/document-schema"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Simulate this flow firing without storing events, enqueuing jobs, or consuming
// cooldown/rate-limit slots.
//
// Works for every trigger activation type:
//
//   - `event`: validates the payload against the event catalog schema and evaluates
//     the trigger conditions.
//   - `custom`: validates the payload against the custom payload schema (conditions
//     do not apply).
//   - `schedule`: validates the stored trigger payload and its step references, and
//     reports the next fire time; request payload overrides are ignored.
//
// The response reports `wouldFire` — whether the flow would actually run right now
// — alongside the gates that decide it (enabled, device attached, blocked,
// cooldown). `rateLimited` is informational and is not folded into `wouldFire`.
func (r *WorkflowFlowService) DryRun(ctx context.Context, flowID string, body WorkflowFlowDryRunParams, opts ...option.RequestOption) (res *WorkflowFlowDryRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/dry-run", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List self-healing repair episodes
func (r *WorkflowFlowService) ListRepairs(ctx context.Context, flowID string, opts ...option.RequestOption) (res *WorkflowFlowListRepairsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/repairs", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Body: {document}. Preserves enabled state. Omitted settings preserves all stored
// settings; present settings replaces them, with absent fields defaulted. Agent
// routines require a new version deployed through the agent. Omitted description,
// devices and trigger options clear their values. Shared triggers return
// trigger_shared and agent-managed routines return routine_agent_managed (409);
// shared action definitions are never modified. To enable a disabled routine with
// at least one step, the owner can PATCH /flows/{flowId} with enabled: true or
// lifecycleStatus: "enabled", without verification. For the verified path used by
// agent deploys, call POST /flows/{flowId}/verify then POST
// /flows/{flowId}/activate; activation requires a successful verification of the
// current graph.
func (r *WorkflowFlowService) ReplaceDocument(ctx context.Context, flowID string, body WorkflowFlowReplaceDocumentParams, opts ...option.RequestOption) (res *WorkflowFlowReplaceDocumentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/document", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns the canonical routine document, including delivery, recording, health,
// cooldown, owner wake and scheduled inputs. Unsupported stored behavior returns
// document_unsupported with JSON Pointer paths.
func (r *WorkflowFlowService) GetDocument(ctx context.Context, flowID string, opts ...option.RequestOption) (res *WorkflowFlowGetDocumentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/document", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get a self-healing repair episode
func (r *WorkflowFlowService) GetRepair(ctx context.Context, episodeID string, query WorkflowFlowGetRepairParams, opts ...option.RequestOption) (res *WorkflowFlowGetRepairResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.FlowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	if episodeID == "" {
		err = errors.New("missing required episodeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/repairs/%s", query.FlowID, episodeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Immediately enqueue one live execution on every device bound to the flow,
// regardless of whether its trigger is event-based, custom, or scheduled. A
// routine without steps is refused with 409 `flow_has_no_steps`.
//
// Supports an optional `Idempotency-Key` header (1-255 printable ASCII
// characters). The run identity is derived from the key, so a retry never starts a
// second run. Replays with the same key and an identical body return the original
// status and the original `executionIds`; a changed body under the same key
// returns 422 `idempotency_key_reused`.
func (r *WorkflowFlowService) Run(ctx context.Context, flowID string, body WorkflowFlowRunParams, opts ...option.RequestOption) (res *WorkflowFlowRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/run", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get flow template context
func (r *WorkflowFlowService) TemplateContext(ctx context.Context, query WorkflowFlowTemplateContextParams, opts ...option.RequestOption) (res *WorkflowFlowTemplateContextResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "flows/template-context"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Clear a flow's blocked status after fixing the underlying issue. Idempotent —
// safe to call on already-healthy flows.
func (r *WorkflowFlowService) Unblock(ctx context.Context, flowID string, opts ...option.RequestOption) (res *WorkflowFlowUnblockResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/unblock", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Validate a flow
func (r *WorkflowFlowService) Validate(ctx context.Context, body WorkflowFlowValidateParams, opts ...option.RequestOption) (res *WorkflowFlowValidateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "flows/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Validates graph, settings and scheduled inputs without writing. Disconnected
// delivery destinations produce warnings; create and replace enforce the existing
// delivery connection checks.
func (r *WorkflowFlowService) ValidateDocument(ctx context.Context, body WorkflowFlowValidateDocumentParams, opts ...option.RequestOption) (res *WorkflowFlowValidateDocumentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "flows/document/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Run a bound, deduplicated **verification** of this flow on exactly one device as
// a real run.
//
// Unlike a normal trigger firing, verification:
//
//   - runs a flow that is `enabled` OR `disabled` (an archived flow returns 404);
//   - refuses routines without steps with 409 `flow_has_no_steps`;
//   - targets one device (`deviceId`, or the single bound device);
//   - is idempotent per `invocationId`: a repeat returns the existing run
//     (`deduplicated: true`, HTTP 200) instead of enqueuing another (HTTP 202);
//   - never mutates flow health, `lastTriggeredAt`, or the repair loop.
//
// Per-activation payload gate:
//
//   - `custom`: validates the payload against the custom payload schema (422 on
//     failure).
//   - `event`: validates the payload against the event catalog schema and evaluates
//     the trigger conditions; a failing gate returns 409 `conditions_not_met` with
//     the dry-run report.
//   - `schedule`: uses the stored trigger payload as event.payload; request payload
//     overrides are ignored.
func (r *WorkflowFlowService) Verify(ctx context.Context, flowID string, body WorkflowFlowVerifyParams, opts ...option.RequestOption) (res *WorkflowFlowVerifyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("flows/%s/verify", flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WorkflowFlowNewResponse struct {
	Data WorkflowFlowNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewResponseData struct {
	ID                  string    `json:"id" api:"required" format:"uuid"`
	AgentID             string    `json:"agentId" api:"required" format:"uuid"`
	ArchivedAt          time.Time `json:"archivedAt" api:"required" format:"date-time"`
	BlockedAt           time.Time `json:"blockedAt" api:"required" format:"date-time"`
	ConsecutiveFailures int64     `json:"consecutiveFailures" api:"required"`
	// Any of "flow", "device".
	CooldownScope   string                              `json:"cooldownScope" api:"required"`
	CooldownSeconds int64                               `json:"cooldownSeconds" api:"required"`
	CreatedAt       time.Time                           `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy       string                              `json:"createdBy" api:"required" format:"uuid"`
	Delivery        WorkflowFlowNewResponseDataDelivery `json:"delivery" api:"required"`
	Description     string                              `json:"description" api:"required"`
	DeviceIDs       []string                            `json:"deviceIds" api:"required" format:"uuid"`
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
func (r WorkflowFlowNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewResponseDataDelivery struct {
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
	Folder      string                                       `json:"folder"`
	Recording   WorkflowFlowNewResponseDataDeliveryRecording `json:"recording"`
	Screenshots any                                          `json:"screenshots"`
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
func (r WorkflowFlowNewResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowNewResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewResponseDataDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowNewResponseDataDeliveryRecording) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowNewResponseDataDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowGetResponse struct {
	Data WorkflowFlowGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowGetResponseData struct {
	ID                  string    `json:"id" api:"required" format:"uuid"`
	AgentID             string    `json:"agentId" api:"required" format:"uuid"`
	ArchivedAt          time.Time `json:"archivedAt" api:"required" format:"date-time"`
	BlockedAt           time.Time `json:"blockedAt" api:"required" format:"date-time"`
	ConsecutiveFailures int64     `json:"consecutiveFailures" api:"required"`
	// Any of "flow", "device".
	CooldownScope   string                              `json:"cooldownScope" api:"required"`
	CooldownSeconds int64                               `json:"cooldownSeconds" api:"required"`
	CreatedAt       time.Time                           `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy       string                              `json:"createdBy" api:"required" format:"uuid"`
	Delivery        WorkflowFlowGetResponseDataDelivery `json:"delivery" api:"required"`
	Description     string                              `json:"description" api:"required"`
	DeviceIDs       []string                            `json:"deviceIds" api:"required" format:"uuid"`
	// Compatibility projection of lifecycleStatus; true only when lifecycleStatus is
	// enabled.
	Enabled                 bool                              `json:"enabled" api:"required"`
	Health                  WorkflowFlowGetResponseDataHealth `json:"health" api:"required"`
	HealthMonitoringEnabled bool                              `json:"healthMonitoringEnabled" api:"required"`
	LastFailureAt           time.Time                         `json:"lastFailureAt" api:"required" format:"date-time"`
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
		Health                    respjson.Field
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
func (r WorkflowFlowGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowGetResponseDataDelivery struct {
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
	Folder      string                                       `json:"folder"`
	Recording   WorkflowFlowGetResponseDataDeliveryRecording `json:"recording"`
	Screenshots any                                          `json:"screenshots"`
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
func (r WorkflowFlowGetResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowGetResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowGetResponseDataDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowGetResponseDataDeliveryRecording) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowGetResponseDataDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowGetResponseDataHealth struct {
	// Whether automatic health-status transitions are enabled; healthMonitoringEnabled
	// must also be true.
	AutoBlockEnabled bool  `json:"autoBlockEnabled" api:"required"`
	BlockThreshold   int64 `json:"blockThreshold" api:"required"`
	FailingThreshold int64 `json:"failingThreshold" api:"required"`
	// Default for new routines; this flow uses selfHealingMaxAttempts.
	MaxAttemptsDefault int64 `json:"maxAttemptsDefault" api:"required"`
	// Permanent failure classes count toward the thresholds; invalid_config, preflight
	// device_not_found, and device_unavailable with a terminated or confirmed gone
	// device block immediately. Event deviceState="gone" is triggers-defined for a
	// deleted/not-found device. Other device_unavailable states use normal thresholds;
	// transient never counts.
	//
	// Any of "device_not_found", "device_unavailable", "permission_denied",
	// "client_error", "transient", "logic", "invalid_config".
	PermanentFailureCodes []*string `json:"permanentFailureCodes" api:"required"`
	// Configuration eligibility: health monitoring and self-healing are both enabled,
	// and the flow is enabled and not deleted. Actual repair additionally requires an
	// enabled repair worker and an eligible script execution/timeout failure.
	SelfHealingEligible bool `json:"selfHealingEligible" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoBlockEnabled      respjson.Field
		BlockThreshold        respjson.Field
		FailingThreshold      respjson.Field
		MaxAttemptsDefault    respjson.Field
		PermanentFailureCodes respjson.Field
		SelfHealingEligible   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowGetResponseDataHealth) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowGetResponseDataHealth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowUpdateResponse struct {
	Data WorkflowFlowUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowUpdateResponseData struct {
	ID                  string    `json:"id" api:"required" format:"uuid"`
	AgentID             string    `json:"agentId" api:"required" format:"uuid"`
	ArchivedAt          time.Time `json:"archivedAt" api:"required" format:"date-time"`
	BlockedAt           time.Time `json:"blockedAt" api:"required" format:"date-time"`
	ConsecutiveFailures int64     `json:"consecutiveFailures" api:"required"`
	// Any of "flow", "device".
	CooldownScope   string                                 `json:"cooldownScope" api:"required"`
	CooldownSeconds int64                                  `json:"cooldownSeconds" api:"required"`
	CreatedAt       time.Time                              `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy       string                                 `json:"createdBy" api:"required" format:"uuid"`
	Delivery        WorkflowFlowUpdateResponseDataDelivery `json:"delivery" api:"required"`
	Description     string                                 `json:"description" api:"required"`
	DeviceIDs       []string                               `json:"deviceIds" api:"required" format:"uuid"`
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
func (r WorkflowFlowUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowUpdateResponseDataDelivery struct {
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
	Folder      string                                          `json:"folder"`
	Recording   WorkflowFlowUpdateResponseDataDeliveryRecording `json:"recording"`
	Screenshots any                                             `json:"screenshots"`
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
func (r WorkflowFlowUpdateResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowUpdateResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowUpdateResponseDataDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowUpdateResponseDataDeliveryRecording) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowUpdateResponseDataDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowListResponse struct {
	ID                  string    `json:"id" api:"required" format:"uuid"`
	AgentID             string    `json:"agentId" api:"required" format:"uuid"`
	ArchivedAt          time.Time `json:"archivedAt" api:"required" format:"date-time"`
	BlockedAt           time.Time `json:"blockedAt" api:"required" format:"date-time"`
	ConsecutiveFailures int64     `json:"consecutiveFailures" api:"required"`
	// Any of "flow", "device".
	CooldownScope   WorkflowFlowListResponseCooldownScope `json:"cooldownScope" api:"required"`
	CooldownSeconds int64                                 `json:"cooldownSeconds" api:"required"`
	CreatedAt       time.Time                             `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy       string                                `json:"createdBy" api:"required" format:"uuid"`
	Delivery        WorkflowFlowListResponseDelivery      `json:"delivery" api:"required"`
	Description     string                                `json:"description" api:"required"`
	DeviceIDs       []string                              `json:"deviceIds" api:"required" format:"uuid"`
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
	LastFailureCode WorkflowFlowListResponseLastFailureCode `json:"lastFailureCode" api:"required"`
	LastTriggeredAt time.Time                               `json:"lastTriggeredAt" api:"required" format:"date-time"`
	// Any of "enabled", "disabled", "archived".
	LifecycleStatus WorkflowFlowListResponseLifecycleStatus `json:"lifecycleStatus" api:"required"`
	Name            string                                  `json:"name" api:"required"`
	NotifyOnFailure bool                                    `json:"notifyOnFailure" api:"required"`
	NotifyOnSuccess bool                                    `json:"notifyOnSuccess" api:"required"`
	NotifyWebhookID string                                  `json:"notifyWebhookId" api:"required" format:"uuid"`
	OwnerID         string                                  `json:"ownerId" api:"required" format:"uuid"`
	// Defaults to final.
	//
	// Any of "final", "off".
	OwnerWake WorkflowFlowListResponseOwnerWake `json:"ownerWake" api:"required"`
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
	ResultDelivery         WorkflowFlowListResponseResultDelivery `json:"resultDelivery" api:"required"`
	SaveSandboxLogs        bool                                   `json:"saveSandboxLogs" api:"required"`
	SelfHealingEnabled     bool                                   `json:"selfHealingEnabled" api:"required"`
	SelfHealingMaxAttempts int64                                  `json:"selfHealingMaxAttempts" api:"required"`
	// Any of "healthy", "failing", "blocked".
	Status WorkflowFlowListResponseStatus `json:"status" api:"required"`
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
func (r WorkflowFlowListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowListResponseCooldownScope string

const (
	WorkflowFlowListResponseCooldownScopeFlow   WorkflowFlowListResponseCooldownScope = "flow"
	WorkflowFlowListResponseCooldownScopeDevice WorkflowFlowListResponseCooldownScope = "device"
)

type WorkflowFlowListResponseDelivery struct {
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
	Folder      string                                    `json:"folder"`
	Recording   WorkflowFlowListResponseDeliveryRecording `json:"recording"`
	Screenshots any                                       `json:"screenshots"`
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
func (r WorkflowFlowListResponseDelivery) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowListResponseDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowListResponseDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowListResponseDeliveryRecording) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowListResponseDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Curated flow failure class. device_unavailable blocks immediately for a
// terminated device or a device confirmed gone; event deviceState="gone" is
// triggers-defined for a deleted/not-found device.
type WorkflowFlowListResponseLastFailureCode string

const (
	WorkflowFlowListResponseLastFailureCodeDeviceNotFound    WorkflowFlowListResponseLastFailureCode = "device_not_found"
	WorkflowFlowListResponseLastFailureCodeDeviceUnavailable WorkflowFlowListResponseLastFailureCode = "device_unavailable"
	WorkflowFlowListResponseLastFailureCodePermissionDenied  WorkflowFlowListResponseLastFailureCode = "permission_denied"
	WorkflowFlowListResponseLastFailureCodeClientError       WorkflowFlowListResponseLastFailureCode = "client_error"
	WorkflowFlowListResponseLastFailureCodeTransient         WorkflowFlowListResponseLastFailureCode = "transient"
	WorkflowFlowListResponseLastFailureCodeLogic             WorkflowFlowListResponseLastFailureCode = "logic"
	WorkflowFlowListResponseLastFailureCodeInvalidConfig     WorkflowFlowListResponseLastFailureCode = "invalid_config"
)

type WorkflowFlowListResponseLifecycleStatus string

const (
	WorkflowFlowListResponseLifecycleStatusEnabled  WorkflowFlowListResponseLifecycleStatus = "enabled"
	WorkflowFlowListResponseLifecycleStatusDisabled WorkflowFlowListResponseLifecycleStatus = "disabled"
	WorkflowFlowListResponseLifecycleStatusArchived WorkflowFlowListResponseLifecycleStatus = "archived"
)

// Defaults to final.
type WorkflowFlowListResponseOwnerWake string

const (
	WorkflowFlowListResponseOwnerWakeFinal WorkflowFlowListResponseOwnerWake = "final"
	WorkflowFlowListResponseOwnerWakeOff   WorkflowFlowListResponseOwnerWake = "off"
)

// Effective routine result delivery: the stored setting, else message for a chat
// sink, else none.
type WorkflowFlowListResponseResultDelivery string

const (
	WorkflowFlowListResponseResultDeliveryNone    WorkflowFlowListResponseResultDelivery = "none"
	WorkflowFlowListResponseResultDeliveryMessage WorkflowFlowListResponseResultDelivery = "message"
	WorkflowFlowListResponseResultDeliveryWake    WorkflowFlowListResponseResultDelivery = "wake"
)

type WorkflowFlowListResponseStatus string

const (
	WorkflowFlowListResponseStatusHealthy WorkflowFlowListResponseStatus = "healthy"
	WorkflowFlowListResponseStatusFailing WorkflowFlowListResponseStatus = "failing"
	WorkflowFlowListResponseStatusBlocked WorkflowFlowListResponseStatus = "blocked"
)

type WorkflowFlowDeleteResponse struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActivateResponse struct {
	Data WorkflowFlowActivateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowActivateResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowActivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActivateResponseData struct {
	ID                  string    `json:"id" api:"required" format:"uuid"`
	AgentID             string    `json:"agentId" api:"required" format:"uuid"`
	ArchivedAt          time.Time `json:"archivedAt" api:"required" format:"date-time"`
	BlockedAt           time.Time `json:"blockedAt" api:"required" format:"date-time"`
	ConsecutiveFailures int64     `json:"consecutiveFailures" api:"required"`
	// Any of "flow", "device".
	CooldownScope   string                                   `json:"cooldownScope" api:"required"`
	CooldownSeconds int64                                    `json:"cooldownSeconds" api:"required"`
	CreatedAt       time.Time                                `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy       string                                   `json:"createdBy" api:"required" format:"uuid"`
	Delivery        WorkflowFlowActivateResponseDataDelivery `json:"delivery" api:"required"`
	Description     string                                   `json:"description" api:"required"`
	DeviceIDs       []string                                 `json:"deviceIds" api:"required" format:"uuid"`
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
func (r WorkflowFlowActivateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowActivateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActivateResponseDataDelivery struct {
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
	Folder      string                                            `json:"folder"`
	Recording   WorkflowFlowActivateResponseDataDeliveryRecording `json:"recording"`
	Screenshots any                                               `json:"screenshots"`
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
func (r WorkflowFlowActivateResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowActivateResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowActivateResponseDataDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowActivateResponseDataDeliveryRecording) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowActivateResponseDataDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowCapacityResponse struct {
	Data WorkflowFlowCapacityResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowCapacityResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowCapacityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowCapacityResponseData struct {
	Included  int64 `json:"included" api:"required"`
	Remaining int64 `json:"remaining" api:"required"`
	// Any of "available", "exhausted".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Included    respjson.Field
		Remaining   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowCapacityResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowCapacityResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowCloneResponse struct {
	Data WorkflowFlowCloneResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowCloneResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowCloneResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowCloneResponseData struct {
	ID                  string    `json:"id" api:"required" format:"uuid"`
	AgentID             string    `json:"agentId" api:"required" format:"uuid"`
	ArchivedAt          time.Time `json:"archivedAt" api:"required" format:"date-time"`
	BlockedAt           time.Time `json:"blockedAt" api:"required" format:"date-time"`
	ConsecutiveFailures int64     `json:"consecutiveFailures" api:"required"`
	// Any of "flow", "device".
	CooldownScope   string                                `json:"cooldownScope" api:"required"`
	CooldownSeconds int64                                 `json:"cooldownSeconds" api:"required"`
	CreatedAt       time.Time                             `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy       string                                `json:"createdBy" api:"required" format:"uuid"`
	Delivery        WorkflowFlowCloneResponseDataDelivery `json:"delivery" api:"required"`
	Description     string                                `json:"description" api:"required"`
	DeviceIDs       []string                              `json:"deviceIds" api:"required" format:"uuid"`
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
func (r WorkflowFlowCloneResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowCloneResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowCloneResponseDataDelivery struct {
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
	Folder      string                                         `json:"folder"`
	Recording   WorkflowFlowCloneResponseDataDeliveryRecording `json:"recording"`
	Screenshots any                                            `json:"screenshots"`
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
func (r WorkflowFlowCloneResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowCloneResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowCloneResponseDataDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowCloneResponseDataDeliveryRecording) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowCloneResponseDataDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewFromDocumentResponse struct {
	Data WorkflowFlowNewFromDocumentResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowNewFromDocumentResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowNewFromDocumentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewFromDocumentResponseData struct {
	Document map[string]any                              `json:"document" api:"required"`
	Flow     WorkflowFlowNewFromDocumentResponseDataFlow `json:"flow" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document    respjson.Field
		Flow        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowNewFromDocumentResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowNewFromDocumentResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewFromDocumentResponseDataFlow struct {
	ID                  string    `json:"id" api:"required" format:"uuid"`
	AgentID             string    `json:"agentId" api:"required" format:"uuid"`
	ArchivedAt          time.Time `json:"archivedAt" api:"required" format:"date-time"`
	BlockedAt           time.Time `json:"blockedAt" api:"required" format:"date-time"`
	ConsecutiveFailures int64     `json:"consecutiveFailures" api:"required"`
	// Any of "flow", "device".
	CooldownScope   string                                              `json:"cooldownScope" api:"required"`
	CooldownSeconds int64                                               `json:"cooldownSeconds" api:"required"`
	CreatedAt       time.Time                                           `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy       string                                              `json:"createdBy" api:"required" format:"uuid"`
	Delivery        WorkflowFlowNewFromDocumentResponseDataFlowDelivery `json:"delivery" api:"required"`
	Description     string                                              `json:"description" api:"required"`
	DeviceIDs       []string                                            `json:"deviceIds" api:"required" format:"uuid"`
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
func (r WorkflowFlowNewFromDocumentResponseDataFlow) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowNewFromDocumentResponseDataFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewFromDocumentResponseDataFlowDelivery struct {
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
	Folder      string                                                       `json:"folder"`
	Recording   WorkflowFlowNewFromDocumentResponseDataFlowDeliveryRecording `json:"recording"`
	Screenshots any                                                          `json:"screenshots"`
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
func (r WorkflowFlowNewFromDocumentResponseDataFlowDelivery) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowNewFromDocumentResponseDataFlowDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewFromDocumentResponseDataFlowDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowNewFromDocumentResponseDataFlowDeliveryRecording) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowFlowNewFromDocumentResponseDataFlowDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDeliveryOptionsResponse struct {
	Data WorkflowFlowDeliveryOptionsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowDeliveryOptionsResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDeliveryOptionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDeliveryOptionsResponseData struct {
	CanDeliverRecording bool `json:"canDeliverRecording" api:"required"`
	// Delivery destinations in their canonical order.
	Destinations     []WorkflowFlowDeliveryOptionsResponseDataDestination    `json:"destinations" api:"required"`
	DirectUploadStep WorkflowFlowDeliveryOptionsResponseDataDirectUploadStep `json:"directUploadStep" api:"required"`
	HasFilesUpload   bool                                                    `json:"hasFilesUpload" api:"required"`
	// Any of "off", "flow", "selected_steps".
	RecordingMode string `json:"recordingMode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanDeliverRecording respjson.Field
		Destinations        respjson.Field
		DirectUploadStep    respjson.Field
		HasFilesUpload      respjson.Field
		RecordingMode       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowDeliveryOptionsResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDeliveryOptionsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDeliveryOptionsResponseDataDestination struct {
	// Any of "connected", "not_connected", "expired", "unknown".
	Connection string `json:"connection" api:"required"`
	ConnectKey string `json:"connectKey" api:"required"`
	IconURL    string `json:"iconUrl" api:"required" format:"uri"`
	// Any of "mobilerun", "one_drive", "google_drive".
	Key   string `json:"key" api:"required"`
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connection  respjson.Field
		ConnectKey  respjson.Field
		IconURL     respjson.Field
		Key         respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowDeliveryOptionsResponseDataDestination) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDeliveryOptionsResponseDataDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDeliveryOptionsResponseDataDirectUploadStep struct {
	FlowActionID string `json:"flowActionId" api:"required" format:"uuid"`
	Name         string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlowActionID respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowDeliveryOptionsResponseDataDirectUploadStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDeliveryOptionsResponseDataDirectUploadStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDocumentSchemaResponse map[string]any

type WorkflowFlowDryRunResponse struct {
	Data WorkflowFlowDryRunResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowDryRunResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDryRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDryRunResponseData struct {
	Actions []WorkflowFlowDryRunResponseDataAction `json:"actions" api:"required"`
	// Any of "event", "schedule", "custom".
	Activation       string                                   `json:"activation" api:"required"`
	ConditionsPassed bool                                     `json:"conditionsPassed" api:"required"`
	Gates            WorkflowFlowDryRunResponseDataGates      `json:"gates" api:"required"`
	NextFireTime     string                                   `json:"nextFireTime" api:"required"`
	RateLimited      bool                                     `json:"rateLimited" api:"required"`
	TestRun          WorkflowFlowDryRunResponseDataTestRun    `json:"testRun" api:"required"`
	Validation       WorkflowFlowDryRunResponseDataValidation `json:"validation" api:"required"`
	WouldFire        bool                                     `json:"wouldFire" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions          respjson.Field
		Activation       respjson.Field
		ConditionsPassed respjson.Field
		Gates            respjson.Field
		NextFireTime     respjson.Field
		RateLimited      respjson.Field
		TestRun          respjson.Field
		Validation       respjson.Field
		WouldFire        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowDryRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDryRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDryRunResponseDataAction struct {
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
func (r WorkflowFlowDryRunResponseDataAction) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDryRunResponseDataAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDryRunResponseDataGates struct {
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
func (r WorkflowFlowDryRunResponseDataGates) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDryRunResponseDataGates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDryRunResponseDataTestRun struct {
	// Reasons a manual test cannot start; empty when ready is true.
	Blockers []WorkflowFlowDryRunResponseDataTestRunBlocker `json:"blockers" api:"required"`
	// Absent required or referenced input paths to supply at test time; empty for
	// stored schedule inputs.
	MissingInputs []string `json:"missingInputs" api:"required"`
	// Whether verification can start, allowing paused or blocked flows and inputs
	// supplied at test time.
	Ready bool `json:"ready" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockers      respjson.Field
		MissingInputs respjson.Field
		Ready         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowDryRunResponseDataTestRun) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDryRunResponseDataTestRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDryRunResponseDataTestRunBlocker struct {
	// Blocking reason: invalid payload, failed conditions, required device absent,
	// account rate limit reached, or a deployment of the flow in progress,
	// respectively.
	//
	// Any of "validation", "conditions_not_met", "no_device", "rate_limited",
	// "deploy_in_progress".
	Code string `json:"code" api:"required"`
	// Validation error field path, in the same format as validation.errors[].field.
	Field string `json:"field"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Field       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowDryRunResponseDataTestRunBlocker) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDryRunResponseDataTestRunBlocker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDryRunResponseDataValidation struct {
	Valid  bool                                            `json:"valid" api:"required"`
	Errors []WorkflowFlowDryRunResponseDataValidationError `json:"errors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Valid       respjson.Field
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowDryRunResponseDataValidation) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDryRunResponseDataValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDryRunResponseDataValidationError struct {
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
func (r WorkflowFlowDryRunResponseDataValidationError) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowDryRunResponseDataValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowListRepairsResponse struct {
	Data []WorkflowFlowListRepairsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowListRepairsResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowListRepairsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowListRepairsResponseData struct {
	ID                string    `json:"id" api:"required" format:"uuid"`
	AgentRunID        string    `json:"agentRunId" api:"required"`
	Attempt           int64     `json:"attempt" api:"required"`
	CandidateSlug     string    `json:"candidateSlug" api:"required"`
	ChatSessionID     string    `json:"chatSessionId" api:"required" format:"uuid"`
	CreatedAt         time.Time `json:"createdAt" api:"required" format:"date-time"`
	DeviceID          string    `json:"deviceId" api:"required" format:"uuid"`
	Episode           int64     `json:"episode" api:"required"`
	Error             string    `json:"error" api:"required"`
	FailedStepIndex   int64     `json:"failedStepIndex" api:"required"`
	FinishedAt        time.Time `json:"finishedAt" api:"required" format:"date-time"`
	FlowID            string    `json:"flowId" api:"required" format:"uuid"`
	MaxAttempts       int64     `json:"maxAttempts" api:"required"`
	OriginalSlug      string    `json:"originalSlug" api:"required"`
	SourceExecutionID string    `json:"sourceExecutionId" api:"required" format:"uuid"`
	StartedAt         time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Any of "pending", "running", "canary", "promoting", "repaired", "failed",
	// "escalated".
	Status    string                                     `json:"status" api:"required"`
	UpdatedAt time.Time                                  `json:"updatedAt" api:"required" format:"date-time"`
	Verdict   WorkflowFlowListRepairsResponseDataVerdict `json:"verdict" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AgentRunID        respjson.Field
		Attempt           respjson.Field
		CandidateSlug     respjson.Field
		ChatSessionID     respjson.Field
		CreatedAt         respjson.Field
		DeviceID          respjson.Field
		Episode           respjson.Field
		Error             respjson.Field
		FailedStepIndex   respjson.Field
		FinishedAt        respjson.Field
		FlowID            respjson.Field
		MaxAttempts       respjson.Field
		OriginalSlug      respjson.Field
		SourceExecutionID respjson.Field
		StartedAt         respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		Verdict           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowListRepairsResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowListRepairsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowListRepairsResponseDataVerdict struct {
	// Any of "ok", "flaky", "broken".
	Outcome string `json:"outcome" api:"required"`
	Summary string `json:"summary" api:"required"`
	Reason  string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Outcome     respjson.Field
		Summary     respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowListRepairsResponseDataVerdict) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowListRepairsResponseDataVerdict) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowReplaceDocumentResponse struct {
	Data WorkflowFlowReplaceDocumentResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowReplaceDocumentResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowReplaceDocumentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowReplaceDocumentResponseData struct {
	Document map[string]any                              `json:"document" api:"required"`
	Flow     WorkflowFlowReplaceDocumentResponseDataFlow `json:"flow" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document    respjson.Field
		Flow        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowReplaceDocumentResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowReplaceDocumentResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowReplaceDocumentResponseDataFlow struct {
	ID                  string    `json:"id" api:"required" format:"uuid"`
	AgentID             string    `json:"agentId" api:"required" format:"uuid"`
	ArchivedAt          time.Time `json:"archivedAt" api:"required" format:"date-time"`
	BlockedAt           time.Time `json:"blockedAt" api:"required" format:"date-time"`
	ConsecutiveFailures int64     `json:"consecutiveFailures" api:"required"`
	// Any of "flow", "device".
	CooldownScope   string                                              `json:"cooldownScope" api:"required"`
	CooldownSeconds int64                                               `json:"cooldownSeconds" api:"required"`
	CreatedAt       time.Time                                           `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy       string                                              `json:"createdBy" api:"required" format:"uuid"`
	Delivery        WorkflowFlowReplaceDocumentResponseDataFlowDelivery `json:"delivery" api:"required"`
	Description     string                                              `json:"description" api:"required"`
	DeviceIDs       []string                                            `json:"deviceIds" api:"required" format:"uuid"`
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
func (r WorkflowFlowReplaceDocumentResponseDataFlow) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowReplaceDocumentResponseDataFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowReplaceDocumentResponseDataFlowDelivery struct {
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
	Folder      string                                                       `json:"folder"`
	Recording   WorkflowFlowReplaceDocumentResponseDataFlowDeliveryRecording `json:"recording"`
	Screenshots any                                                          `json:"screenshots"`
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
func (r WorkflowFlowReplaceDocumentResponseDataFlowDelivery) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowReplaceDocumentResponseDataFlowDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowReplaceDocumentResponseDataFlowDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowReplaceDocumentResponseDataFlowDeliveryRecording) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowFlowReplaceDocumentResponseDataFlowDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowGetDocumentResponse struct {
	Data map[string]any `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowGetDocumentResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowGetDocumentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowGetRepairResponse struct {
	Data WorkflowFlowGetRepairResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowGetRepairResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowGetRepairResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowGetRepairResponseData struct {
	ID              string    `json:"id" api:"required" format:"uuid"`
	AgentRunID      string    `json:"agentRunId" api:"required"`
	Attempt         int64     `json:"attempt" api:"required"`
	CandidateSlug   string    `json:"candidateSlug" api:"required"`
	ChatSessionID   string    `json:"chatSessionId" api:"required" format:"uuid"`
	CreatedAt       time.Time `json:"createdAt" api:"required" format:"date-time"`
	DeviceID        string    `json:"deviceId" api:"required" format:"uuid"`
	Episode         int64     `json:"episode" api:"required"`
	Error           string    `json:"error" api:"required"`
	FailedStepIndex int64     `json:"failedStepIndex" api:"required"`
	FinishedAt      time.Time `json:"finishedAt" api:"required" format:"date-time"`
	FlowID          string    `json:"flowId" api:"required" format:"uuid"`
	// Any of "healthy", "failing", "blocked".
	FlowStatus        string    `json:"flowStatus" api:"required"`
	MaxAttempts       int64     `json:"maxAttempts" api:"required"`
	OriginalSlug      string    `json:"originalSlug" api:"required"`
	SourceExecutionID string    `json:"sourceExecutionId" api:"required" format:"uuid"`
	StartedAt         time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Any of "pending", "running", "canary", "promoting", "repaired", "failed",
	// "escalated".
	Status    string                                   `json:"status" api:"required"`
	UpdatedAt time.Time                                `json:"updatedAt" api:"required" format:"date-time"`
	Verdict   WorkflowFlowGetRepairResponseDataVerdict `json:"verdict" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AgentRunID        respjson.Field
		Attempt           respjson.Field
		CandidateSlug     respjson.Field
		ChatSessionID     respjson.Field
		CreatedAt         respjson.Field
		DeviceID          respjson.Field
		Episode           respjson.Field
		Error             respjson.Field
		FailedStepIndex   respjson.Field
		FinishedAt        respjson.Field
		FlowID            respjson.Field
		FlowStatus        respjson.Field
		MaxAttempts       respjson.Field
		OriginalSlug      respjson.Field
		SourceExecutionID respjson.Field
		StartedAt         respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		Verdict           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowGetRepairResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowGetRepairResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowGetRepairResponseDataVerdict struct {
	// Any of "ok", "flaky", "broken".
	Outcome string `json:"outcome" api:"required"`
	Summary string `json:"summary" api:"required"`
	Reason  string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Outcome     respjson.Field
		Summary     respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowGetRepairResponseDataVerdict) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowGetRepairResponseDataVerdict) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowRunResponse struct {
	Data WorkflowFlowRunResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowRunResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowRunResponseData struct {
	EnqueuedCount int64 `json:"enqueuedCount" api:"required"`
	// Ids of the executions this run created. An idempotent replay returns the ids of
	// the original run.
	ExecutionIDs []string `json:"executionIds" api:"required" format:"uuid"`
	// At least one admitted execution failed queue submission. This is an enqueue
	// error, not a skip. Keyed runs report prepared executions; post-commit submission
	// failures are recorded on the execution instead.
	//
	// Any of true.
	EnqueueFailed bool `json:"enqueueFailed"`
	// Why no execution was admitted: device_required means no device and actions are
	// not device-free eligible; flow_disabled means disabled at admission;
	// flow_unavailable means missing or archived at admission. Blocked flows keep
	// their existing error response. Manual runs bypass trigger conditions.
	Skipped WorkflowFlowRunResponseDataSkipped `json:"skipped"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnqueuedCount respjson.Field
		ExecutionIDs  respjson.Field
		EnqueueFailed respjson.Field
		Skipped       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why no execution was admitted: device_required means no device and actions are
// not device-free eligible; flow_disabled means disabled at admission;
// flow_unavailable means missing or archived at admission. Blocked flows keep
// their existing error response. Manual runs bypass trigger conditions.
type WorkflowFlowRunResponseDataSkipped struct {
	// Any of "device_required", "flow_disabled", "flow_unavailable".
	Reason string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowRunResponseDataSkipped) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowRunResponseDataSkipped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowTemplateContextResponse struct {
	Data WorkflowFlowTemplateContextResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowTemplateContextResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowTemplateContextResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowTemplateContextResponseData struct {
	Grammar                   string                                               `json:"grammar" api:"required"`
	Loop                      WorkflowFlowTemplateContextResponseDataLoop          `json:"loop" api:"required"`
	PayloadAliases            []string                                             `json:"payloadAliases" api:"required"`
	Roots                     []WorkflowFlowTemplateContextResponseDataRoot        `json:"roots" api:"required"`
	StepReference             WorkflowFlowTemplateContextResponseDataStepReference `json:"stepReference" api:"required"`
	TemplateResolutionVersion int64                                                `json:"templateResolutionVersion" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Grammar                   respjson.Field
		Loop                      respjson.Field
		PayloadAliases            respjson.Field
		Roots                     respjson.Field
		StepReference             respjson.Field
		TemplateResolutionVersion respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowTemplateContextResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowTemplateContextResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowTemplateContextResponseDataLoop struct {
	Context             []string `json:"context" api:"required"`
	ItemVariableDefault string   `json:"itemVariableDefault" api:"required"`
	Name                string   `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context             respjson.Field
		ItemVariableDefault respjson.Field
		Name                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowTemplateContextResponseDataLoop) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowTemplateContextResponseDataLoop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowTemplateContextResponseDataRoot struct {
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	// Known root members; meta lists all execution calendar variables
	Shape []WorkflowFlowTemplateContextResponseDataRootShape `json:"shape" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		Shape       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowTemplateContextResponseDataRoot) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowTemplateContextResponseDataRoot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowTemplateContextResponseDataRootShape struct {
	Description string `json:"description" api:"required"`
	Example     string `json:"example" api:"required"`
	Name        string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Example     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowTemplateContextResponseDataRootShape) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowTemplateContextResponseDataRootShape) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowTemplateContextResponseDataStepReference struct {
	Note     string `json:"note" api:"required"`
	Pattern  string `json:"pattern" api:"required"`
	KeyRule  string `json:"keyRule"`
	SlugRule string `json:"slugRule"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Note        respjson.Field
		Pattern     respjson.Field
		KeyRule     respjson.Field
		SlugRule    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowTemplateContextResponseDataStepReference) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowTemplateContextResponseDataStepReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowUnblockResponse struct {
	Data WorkflowFlowUnblockResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowUnblockResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowUnblockResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowUnblockResponseData struct {
	ID                  string    `json:"id" api:"required" format:"uuid"`
	AgentID             string    `json:"agentId" api:"required" format:"uuid"`
	ArchivedAt          time.Time `json:"archivedAt" api:"required" format:"date-time"`
	BlockedAt           time.Time `json:"blockedAt" api:"required" format:"date-time"`
	ConsecutiveFailures int64     `json:"consecutiveFailures" api:"required"`
	// Any of "flow", "device".
	CooldownScope   string                                  `json:"cooldownScope" api:"required"`
	CooldownSeconds int64                                   `json:"cooldownSeconds" api:"required"`
	CreatedAt       time.Time                               `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy       string                                  `json:"createdBy" api:"required" format:"uuid"`
	Delivery        WorkflowFlowUnblockResponseDataDelivery `json:"delivery" api:"required"`
	Description     string                                  `json:"description" api:"required"`
	DeviceIDs       []string                                `json:"deviceIds" api:"required" format:"uuid"`
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
func (r WorkflowFlowUnblockResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowUnblockResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowUnblockResponseDataDelivery struct {
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
	Folder      string                                           `json:"folder"`
	Recording   WorkflowFlowUnblockResponseDataDeliveryRecording `json:"recording"`
	Screenshots any                                              `json:"screenshots"`
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
func (r WorkflowFlowUnblockResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowUnblockResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowUnblockResponseDataDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowUnblockResponseDataDeliveryRecording) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowUnblockResponseDataDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateResponse struct {
	Data WorkflowFlowValidateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateResponseData struct {
	Errors   []WorkflowFlowValidateResponseDataError   `json:"errors" api:"required"`
	Valid    bool                                      `json:"valid" api:"required"`
	Warnings []WorkflowFlowValidateResponseDataWarning `json:"warnings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		Valid       respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowValidateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowValidateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateResponseDataError struct {
	Code       string                                      `json:"code" api:"required"`
	Field      string                                      `json:"field" api:"required"`
	Message    string                                      `json:"message" api:"required"`
	ActionID   string                                      `json:"actionId"`
	KnownSteps []string                                    `json:"knownSteps"`
	Reference  string                                      `json:"reference"`
	Repair     WorkflowFlowValidateResponseDataErrorRepair `json:"repair"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Field       respjson.Field
		Message     respjson.Field
		ActionID    respjson.Field
		KnownSteps  respjson.Field
		Reference   respjson.Field
		Repair      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowValidateResponseDataError) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowValidateResponseDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateResponseDataErrorRepair struct {
	// Any of "fix_spec_and_restart", "ask_user", "retry_later".
	Action string `json:"action" api:"required"`
	Hint   string `json:"hint" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Hint        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowValidateResponseDataErrorRepair) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowValidateResponseDataErrorRepair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateResponseDataWarning struct {
	Code        string `json:"code" api:"required"`
	Message     string `json:"message" api:"required"`
	ConnectKey  string `json:"connectKey"`
	Destination string `json:"destination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ConnectKey  respjson.Field
		Destination respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowValidateResponseDataWarning) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowValidateResponseDataWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateDocumentResponse struct {
	Data WorkflowFlowValidateDocumentResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowValidateDocumentResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowValidateDocumentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateDocumentResponseData struct {
	Errors   []WorkflowFlowValidateDocumentResponseDataError   `json:"errors" api:"required"`
	Valid    bool                                              `json:"valid" api:"required"`
	Warnings []WorkflowFlowValidateDocumentResponseDataWarning `json:"warnings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		Valid       respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowValidateDocumentResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowValidateDocumentResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateDocumentResponseDataError struct {
	Code string `json:"code" api:"required"`
	// RFC 6901 JSON Pointer; empty string is the document root
	Field   string                                              `json:"field" api:"required"`
	Message string                                              `json:"message" api:"required"`
	Repair  WorkflowFlowValidateDocumentResponseDataErrorRepair `json:"repair" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Field       respjson.Field
		Message     respjson.Field
		Repair      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowValidateDocumentResponseDataError) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowValidateDocumentResponseDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateDocumentResponseDataErrorRepair struct {
	// Any of "fix_spec_and_restart", "ask_user", "retry_later".
	Action string `json:"action" api:"required"`
	Hint   string `json:"hint" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Hint        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowValidateDocumentResponseDataErrorRepair) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowValidateDocumentResponseDataErrorRepair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateDocumentResponseDataWarning struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowValidateDocumentResponseDataWarning) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowValidateDocumentResponseDataWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowVerifyResponse struct {
	Data WorkflowFlowVerifyResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowVerifyResponseData struct {
	// True when this run already existed for the (flow, invocationId) key (HTTP 200);
	// false for a newly enqueued run (HTTP 202).
	Deduplicated bool `json:"deduplicated" api:"required"`
	// Device the verification runs on; null for a device-free verification (a 0-device
	// flow whose actions are all device-free allowlisted).
	DeviceID     string `json:"deviceId" api:"required" format:"uuid"`
	ExecutionID  string `json:"executionId" api:"required" format:"uuid"`
	FlowID       string `json:"flowId" api:"required" format:"uuid"`
	InvocationID string `json:"invocationId" api:"required" format:"uuid"`
	// Any of "verification".
	Kind string `json:"kind" api:"required"`
	// Any of "pending", "running", "success", "failed", "cancelled", "skipped",
	// "invalid".
	Status    string `json:"status" api:"required"`
	TriggerID string `json:"triggerId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deduplicated respjson.Field
		DeviceID     respjson.Field
		ExecutionID  respjson.Field
		FlowID       respjson.Field
		InvocationID respjson.Field
		Kind         respjson.Field
		Status       respjson.Field
		TriggerID    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFlowVerifyResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFlowVerifyResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewParams struct {
	Name            string            `json:"name" api:"required"`
	TriggerID       string            `json:"triggerId" api:"required" format:"uuid"`
	CooldownSeconds param.Opt[int64]  `json:"cooldownSeconds,omitzero"`
	NotifyWebhookID param.Opt[string] `json:"notifyWebhookId,omitzero" format:"uuid"`
	Description     param.Opt[string] `json:"description,omitzero"`
	// Defaults to true when steps exist. A routine without steps is created disabled;
	// explicit true is rejected.
	Enabled                 param.Opt[bool]               `json:"enabled,omitzero"`
	HealthMonitoringEnabled param.Opt[bool]               `json:"healthMonitoringEnabled,omitzero"`
	NotifyOnFailure         param.Opt[bool]               `json:"notifyOnFailure,omitzero"`
	NotifyOnSuccess         param.Opt[bool]               `json:"notifyOnSuccess,omitzero"`
	RecordingEnabled        param.Opt[bool]               `json:"recordingEnabled,omitzero"`
	SaveSandboxLogs         param.Opt[bool]               `json:"saveSandboxLogs,omitzero"`
	SelfHealingEnabled      param.Opt[bool]               `json:"selfHealingEnabled,omitzero"`
	SelfHealingMaxAttempts  param.Opt[int64]              `json:"selfHealingMaxAttempts,omitzero"`
	Actions                 []WorkflowFlowNewParamsAction `json:"actions,omitzero"`
	// Any of "flow", "device".
	CooldownScope WorkflowFlowNewParamsCooldownScope `json:"cooldownScope,omitzero"`
	Delivery      WorkflowFlowNewParamsDelivery      `json:"delivery,omitzero"`
	DeviceIDs     []string                           `json:"deviceIds,omitzero" format:"uuid"`
	// Defaults to final.
	//
	// Any of "final", "off".
	OwnerWake       WorkflowFlowNewParamsOwnerWake       `json:"ownerWake,omitzero"`
	RecordingPolicy WorkflowFlowNewParamsRecordingPolicy `json:"recordingPolicy,omitzero"`
	paramObj
}

func (r WorkflowFlowNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Position are required.
type WorkflowFlowNewParamsAction struct {
	ActionID        string          `json:"actionId" api:"required" format:"uuid"`
	Position        int64           `json:"position" api:"required"`
	ContinueOnError param.Opt[bool] `json:"continueOnError,omitzero"`
	// Stable identifier used by template resolution v2+ to address this step as
	// {{key.field}}. Omitted keys are derived from the step name whenever a flow is
	// created. For template resolution v3, replacing an existing flow's full action
	// tree requires every step to carry an explicit key — a missing key is rejected
	// there, not derived, so a name change can never silently move a step's key.
	Key               param.Opt[string]                    `json:"key,omitzero"`
	NameOverride      param.Opt[string]                    `json:"nameOverride,omitzero"`
	RecordingEnabled  param.Opt[bool]                      `json:"recordingEnabled,omitzero"`
	OverridesProperty WorkflowFlowNewParamsActionOverrides `json:"overrides,omitzero"`
	Children          []WorkflowFlowNewParamsActionChild   `json:"children,omitzero"`
	paramObj
}

func (r WorkflowFlowNewParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowNewParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowNewParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Position are required.
type WorkflowFlowNewParamsActionChild struct {
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
	OverridesProperty WorkflowFlowNewParamsActionChildOverrides `json:"overrides,omitzero"`
	paramObj
}

func (r WorkflowFlowNewParamsActionChild) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowNewParamsActionChild
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowNewParamsActionChild) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewParamsActionChildOverrides struct {
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r WorkflowFlowNewParamsActionChildOverrides) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowNewParamsActionChildOverrides
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowNewParamsActionChildOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewParamsActionOverrides struct {
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r WorkflowFlowNewParamsActionOverrides) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowNewParamsActionOverrides
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowNewParamsActionOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewParamsCooldownScope string

const (
	WorkflowFlowNewParamsCooldownScopeFlow   WorkflowFlowNewParamsCooldownScope = "flow"
	WorkflowFlowNewParamsCooldownScopeDevice WorkflowFlowNewParamsCooldownScope = "device"
)

// The property Destination is required.
type WorkflowFlowNewParamsDelivery struct {
	// Any of "mobilerun", "one_drive", "google_drive".
	Destination string `json:"destination,omitzero" api:"required"`
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
	Folder      param.Opt[string]                      `json:"folder,omitzero"`
	Recording   WorkflowFlowNewParamsDeliveryRecording `json:"recording,omitzero"`
	Screenshots any                                    `json:"screenshots,omitzero"`
	paramObj
}

func (r WorkflowFlowNewParamsDelivery) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowNewParamsDelivery
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowNewParamsDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowNewParamsDelivery](
		"destination", "mobilerun", "one_drive", "google_drive",
	)
}

// The property Filename is required.
type WorkflowFlowNewParamsDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	paramObj
}

func (r WorkflowFlowNewParamsDeliveryRecording) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowNewParamsDeliveryRecording
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowNewParamsDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defaults to final.
type WorkflowFlowNewParamsOwnerWake string

const (
	WorkflowFlowNewParamsOwnerWakeFinal WorkflowFlowNewParamsOwnerWake = "final"
	WorkflowFlowNewParamsOwnerWakeOff   WorkflowFlowNewParamsOwnerWake = "off"
)

// The property Mode is required.
type WorkflowFlowNewParamsRecordingPolicy struct {
	// Any of "off", "flow", "selected_steps".
	Mode string `json:"mode,omitzero" api:"required"`
	paramObj
}

func (r WorkflowFlowNewParamsRecordingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowNewParamsRecordingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowNewParamsRecordingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowNewParamsRecordingPolicy](
		"mode", "off", "flow", "selected_steps",
	)
}

type WorkflowFlowUpdateParams struct {
	CooldownSeconds         param.Opt[int64]  `json:"cooldownSeconds,omitzero"`
	NotifyWebhookID         param.Opt[string] `json:"notifyWebhookId,omitzero" format:"uuid"`
	Description             param.Opt[string] `json:"description,omitzero"`
	Enabled                 param.Opt[bool]   `json:"enabled,omitzero"`
	HealthMonitoringEnabled param.Opt[bool]   `json:"healthMonitoringEnabled,omitzero"`
	Name                    param.Opt[string] `json:"name,omitzero"`
	NotifyOnFailure         param.Opt[bool]   `json:"notifyOnFailure,omitzero"`
	NotifyOnSuccess         param.Opt[bool]   `json:"notifyOnSuccess,omitzero"`
	// Deprecated compatibility field. true maps to recordingPolicy.mode="flow"; false
	// maps to "off".
	RecordingEnabled       param.Opt[bool]                  `json:"recordingEnabled,omitzero"`
	SaveSandboxLogs        param.Opt[bool]                  `json:"saveSandboxLogs,omitzero"`
	SelfHealingEnabled     param.Opt[bool]                  `json:"selfHealingEnabled,omitzero"`
	SelfHealingMaxAttempts param.Opt[int64]                 `json:"selfHealingMaxAttempts,omitzero"`
	TriggerID              param.Opt[string]                `json:"triggerId,omitzero" format:"uuid"`
	Delivery               WorkflowFlowUpdateParamsDelivery `json:"delivery,omitzero"`
	// Any of "flow", "device".
	CooldownScope WorkflowFlowUpdateParamsCooldownScope `json:"cooldownScope,omitzero"`
	DeviceIDs     []string                              `json:"deviceIds,omitzero" format:"uuid"`
	// Set the visible agent lifecycle. Archive remains available only through DELETE.
	//
	// Any of "enabled", "disabled".
	LifecycleStatus WorkflowFlowUpdateParamsLifecycleStatus `json:"lifecycleStatus,omitzero"`
	// Defaults to final.
	//
	// Any of "final", "off".
	OwnerWake       WorkflowFlowUpdateParamsOwnerWake       `json:"ownerWake,omitzero"`
	RecordingPolicy WorkflowFlowUpdateParamsRecordingPolicy `json:"recordingPolicy,omitzero"`
	paramObj
}

func (r WorkflowFlowUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowUpdateParamsCooldownScope string

const (
	WorkflowFlowUpdateParamsCooldownScopeFlow   WorkflowFlowUpdateParamsCooldownScope = "flow"
	WorkflowFlowUpdateParamsCooldownScopeDevice WorkflowFlowUpdateParamsCooldownScope = "device"
)

// The property Destination is required.
type WorkflowFlowUpdateParamsDelivery struct {
	// Any of "mobilerun", "one_drive", "google_drive".
	Destination string `json:"destination,omitzero" api:"required"`
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
	Folder      param.Opt[string]                         `json:"folder,omitzero"`
	Recording   WorkflowFlowUpdateParamsDeliveryRecording `json:"recording,omitzero"`
	Screenshots any                                       `json:"screenshots,omitzero"`
	paramObj
}

func (r WorkflowFlowUpdateParamsDelivery) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowUpdateParamsDelivery
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowUpdateParamsDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowUpdateParamsDelivery](
		"destination", "mobilerun", "one_drive", "google_drive",
	)
}

// The property Filename is required.
type WorkflowFlowUpdateParamsDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	paramObj
}

func (r WorkflowFlowUpdateParamsDeliveryRecording) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowUpdateParamsDeliveryRecording
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowUpdateParamsDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set the visible agent lifecycle. Archive remains available only through DELETE.
type WorkflowFlowUpdateParamsLifecycleStatus string

const (
	WorkflowFlowUpdateParamsLifecycleStatusEnabled  WorkflowFlowUpdateParamsLifecycleStatus = "enabled"
	WorkflowFlowUpdateParamsLifecycleStatusDisabled WorkflowFlowUpdateParamsLifecycleStatus = "disabled"
)

// Defaults to final.
type WorkflowFlowUpdateParamsOwnerWake string

const (
	WorkflowFlowUpdateParamsOwnerWakeFinal WorkflowFlowUpdateParamsOwnerWake = "final"
	WorkflowFlowUpdateParamsOwnerWakeOff   WorkflowFlowUpdateParamsOwnerWake = "off"
)

// The property Mode is required.
type WorkflowFlowUpdateParamsRecordingPolicy struct {
	// Any of "off", "flow", "selected_steps".
	Mode string `json:"mode,omitzero" api:"required"`
	paramObj
}

func (r WorkflowFlowUpdateParamsRecordingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowUpdateParamsRecordingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowUpdateParamsRecordingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowUpdateParamsRecordingPolicy](
		"mode", "off", "flow", "selected_steps",
	)
}

type WorkflowFlowListParams struct {
	CreatedBy param.Opt[string] `query:"createdBy,omitzero" format:"uuid" json:"-"`
	Page      param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize  param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	Search    param.Opt[string] `query:"search,omitzero" json:"-"`
	TriggerID param.Opt[string] `query:"triggerId,omitzero" format:"uuid" json:"-"`
	// Agent bucket UUID or unassigned. Omit on list to include every bucket; omit on
	// capacity for unassigned.
	AgentID WorkflowFlowListParamsAgentID `query:"agentId,omitzero" json:"-"`
	// Only include flows with this enabled state.
	//
	// Any of "true", "false".
	Enabled WorkflowFlowListParamsEnabled `query:"enabled,omitzero" json:"-"`
	// Only include flows created by you.
	//
	// Any of "true", "false".
	Mine WorkflowFlowListParamsMine `query:"mine,omitzero" json:"-"`
	// Any of "name", "createdAt", "updatedAt".
	OrderBy WorkflowFlowListParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// Any of "asc", "desc".
	OrderByDirection WorkflowFlowListParamsOrderByDirection `query:"orderByDirection,omitzero" json:"-"`
	// Any of "healthy", "failing", "blocked".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowFlowListParams]'s query parameters as `url.Values`.
func (r WorkflowFlowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Agent bucket UUID or unassigned. Omit on list to include every bucket; omit on
// capacity for unassigned.
type WorkflowFlowListParamsAgentID string

const (
	WorkflowFlowListParamsAgentIDUnassigned WorkflowFlowListParamsAgentID = "unassigned"
)

// Only include flows with this enabled state.
type WorkflowFlowListParamsEnabled string

const (
	WorkflowFlowListParamsEnabledTrue  WorkflowFlowListParamsEnabled = "true"
	WorkflowFlowListParamsEnabledFalse WorkflowFlowListParamsEnabled = "false"
)

// Only include flows created by you.
type WorkflowFlowListParamsMine string

const (
	WorkflowFlowListParamsMineTrue  WorkflowFlowListParamsMine = "true"
	WorkflowFlowListParamsMineFalse WorkflowFlowListParamsMine = "false"
)

type WorkflowFlowListParamsOrderBy string

const (
	WorkflowFlowListParamsOrderByName      WorkflowFlowListParamsOrderBy = "name"
	WorkflowFlowListParamsOrderByCreatedAt WorkflowFlowListParamsOrderBy = "createdAt"
	WorkflowFlowListParamsOrderByUpdatedAt WorkflowFlowListParamsOrderBy = "updatedAt"
)

type WorkflowFlowListParamsOrderByDirection string

const (
	WorkflowFlowListParamsOrderByDirectionAsc  WorkflowFlowListParamsOrderByDirection = "asc"
	WorkflowFlowListParamsOrderByDirectionDesc WorkflowFlowListParamsOrderByDirection = "desc"
)

type WorkflowFlowActivateParams struct {
	VerificationExecutionID string `json:"verificationExecutionId" api:"required" format:"uuid"`
	paramObj
}

func (r WorkflowFlowActivateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowActivateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowActivateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowCapacityParams struct {
	// Agent bucket UUID or unassigned. Omit on list to include every bucket; omit on
	// capacity for unassigned.
	AgentID WorkflowFlowCapacityParamsAgentID `query:"agentId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowFlowCapacityParams]'s query parameters as
// `url.Values`.
func (r WorkflowFlowCapacityParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Agent bucket UUID or unassigned. Omit on list to include every bucket; omit on
// capacity for unassigned.
type WorkflowFlowCapacityParamsAgentID string

const (
	WorkflowFlowCapacityParamsAgentIDUnassigned WorkflowFlowCapacityParamsAgentID = "unassigned"
)

type WorkflowFlowCloneParams struct {
	Name      param.Opt[string] `json:"name,omitzero"`
	DeviceIDs []string          `json:"deviceIds,omitzero" format:"uuid"`
	paramObj
}

func (r WorkflowFlowCloneParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowCloneParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowCloneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowNewFromDocumentParams struct {
	// Routine document; format at GET /flows/document-schema
	Document any `json:"document,omitzero"`
	paramObj
}

func (r WorkflowFlowNewFromDocumentParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowNewFromDocumentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowNewFromDocumentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowDocumentSchemaParams struct {
	// Any of "agent_routine".
	Profile WorkflowFlowDocumentSchemaParamsProfile `query:"profile,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowFlowDocumentSchemaParams]'s query parameters as
// `url.Values`.
func (r WorkflowFlowDocumentSchemaParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowFlowDocumentSchemaParamsProfile string

const (
	WorkflowFlowDocumentSchemaParamsProfileAgentRoutine WorkflowFlowDocumentSchemaParamsProfile = "agent_routine"
)

type WorkflowFlowDryRunParams struct {
	Payload map[string]any `json:"payload,omitzero"`
	paramObj
}

func (r WorkflowFlowDryRunParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowDryRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowDryRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowReplaceDocumentParams struct {
	// Routine document; format at GET /flows/document-schema
	Document any `json:"document,omitzero"`
	paramObj
}

func (r WorkflowFlowReplaceDocumentParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowReplaceDocumentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowReplaceDocumentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowGetRepairParams struct {
	FlowID string `path:"flowId" api:"required" format:"uuid" json:"-"`
	paramObj
}

type WorkflowFlowRunParams struct {
	// Record and keep this run's scripts.run screenshots
	// (device.screenshot(save=True)) even without flow delivery. Used for a one-off
	// proof run.
	CaptureScreenshots param.Opt[bool] `json:"captureScreenshots,omitzero"`
	Payload            map[string]any  `json:"payload,omitzero"`
	paramObj
}

func (r WorkflowFlowRunParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowTemplateContextParams struct {
	TemplateResolutionVersion param.Opt[int64] `query:"templateResolutionVersion,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowFlowTemplateContextParams]'s query parameters as
// `url.Values`.
func (r WorkflowFlowTemplateContextParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowFlowValidateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfObject *WorkflowFlowValidateParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWorkflowFlowValidatesBodyObject2 *WorkflowFlowValidateParamsBodyObject2 `json:",inline"`

	paramObj
}

func (u WorkflowFlowValidateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfObject, u.OfWorkflowFlowValidatesBodyObject2)
}
func (r *WorkflowFlowValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Actions, Name are required.
type WorkflowFlowValidateParamsBodyObject struct {
	Actions                 []WorkflowFlowValidateParamsBodyObjectActionUnion `json:"actions,omitzero" api:"required"`
	Name                    string                                            `json:"name" api:"required"`
	CooldownSeconds         param.Opt[int64]                                  `json:"cooldownSeconds,omitzero"`
	NotifyWebhookID         param.Opt[string]                                 `json:"notifyWebhookId,omitzero" format:"uuid"`
	Description             param.Opt[string]                                 `json:"description,omitzero"`
	Enabled                 param.Opt[bool]                                   `json:"enabled,omitzero"`
	HealthMonitoringEnabled param.Opt[bool]                                   `json:"healthMonitoringEnabled,omitzero"`
	NotifyOnFailure         param.Opt[bool]                                   `json:"notifyOnFailure,omitzero"`
	NotifyOnSuccess         param.Opt[bool]                                   `json:"notifyOnSuccess,omitzero"`
	RecordingEnabled        param.Opt[bool]                                   `json:"recordingEnabled,omitzero"`
	SaveSandboxLogs         param.Opt[bool]                                   `json:"saveSandboxLogs,omitzero"`
	SelfHealingEnabled      param.Opt[bool]                                   `json:"selfHealingEnabled,omitzero"`
	SelfHealingMaxAttempts  param.Opt[int64]                                  `json:"selfHealingMaxAttempts,omitzero"`
	TriggerID               param.Opt[string]                                 `json:"triggerId,omitzero" format:"uuid"`
	// Any of "flow", "device".
	CooldownScope string                                       `json:"cooldownScope,omitzero"`
	Delivery      WorkflowFlowValidateParamsBodyObjectDelivery `json:"delivery,omitzero"`
	DeviceIDs     []string                                     `json:"deviceIds,omitzero" format:"uuid"`
	// Defaults to final.
	//
	// Any of "final", "off".
	OwnerWake       string                                              `json:"ownerWake,omitzero"`
	RecordingPolicy WorkflowFlowValidateParamsBodyObjectRecordingPolicy `json:"recordingPolicy,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject](
		"cooldownScope", "flow", "device",
	)
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject](
		"ownerWake", "final", "off",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowFlowValidateParamsBodyObjectActionUnion struct {
	OfWorkflowFlowValidatesBodyObjectActionFlowActionInput      *WorkflowFlowValidateParamsBodyObjectActionFlowActionInput      `json:",omitzero,inline"`
	OfWorkflowFlowValidatesBodyObjectActionFlowDraftActionInput *WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInput `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowFlowValidateParamsBodyObjectActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWorkflowFlowValidatesBodyObjectActionFlowActionInput, u.OfWorkflowFlowValidatesBodyObjectActionFlowDraftActionInput)
}
func (u *WorkflowFlowValidateParamsBodyObjectActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ActionID, Position are required.
type WorkflowFlowValidateParamsBodyObjectActionFlowActionInput struct {
	ActionID        string          `json:"actionId" api:"required" format:"uuid"`
	Position        int64           `json:"position" api:"required"`
	ContinueOnError param.Opt[bool] `json:"continueOnError,omitzero"`
	// Stable identifier used by template resolution v2+ to address this step as
	// {{key.field}}. Omitted keys are derived from the step name whenever a flow is
	// created. For template resolution v3, replacing an existing flow's full action
	// tree requires every step to carry an explicit key — a missing key is rejected
	// there, not derived, so a name change can never silently move a step's key.
	Key               param.Opt[string]                                                  `json:"key,omitzero"`
	NameOverride      param.Opt[string]                                                  `json:"nameOverride,omitzero"`
	RecordingEnabled  param.Opt[bool]                                                    `json:"recordingEnabled,omitzero"`
	OverridesProperty WorkflowFlowValidateParamsBodyObjectActionFlowActionInputOverrides `json:"overrides,omitzero"`
	Children          []WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChild   `json:"children,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObjectActionFlowActionInput) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObjectActionFlowActionInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObjectActionFlowActionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Position are required.
type WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChild struct {
	ActionID        string          `json:"actionId" api:"required" format:"uuid"`
	Position        int64           `json:"position" api:"required"`
	ContinueOnError param.Opt[bool] `json:"continueOnError,omitzero"`
	// Stable identifier used by template resolution v2+ to address this step as
	// {{key.field}}. Omitted keys are derived from the step name whenever a flow is
	// created. For template resolution v3, replacing an existing flow's full action
	// tree requires every step to carry an explicit key — a missing key is rejected
	// there, not derived, so a name change can never silently move a step's key.
	Key               param.Opt[string]                                                       `json:"key,omitzero"`
	NameOverride      param.Opt[string]                                                       `json:"nameOverride,omitzero"`
	RecordingEnabled  param.Opt[bool]                                                         `json:"recordingEnabled,omitzero"`
	OverridesProperty WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChildOverrides `json:"overrides,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChild) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChild
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChild) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChildOverrides struct {
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChildOverrides) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChildOverrides
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObjectActionFlowActionInputChildOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateParamsBodyObjectActionFlowActionInputOverrides struct {
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObjectActionFlowActionInputOverrides) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObjectActionFlowActionInputOverrides
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObjectActionFlowActionInputOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Draft, Position are required.
type WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInput struct {
	Draft           WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInputDraft `json:"draft,omitzero" api:"required"`
	Position        int64                                                               `json:"position" api:"required"`
	ContinueOnError param.Opt[bool]                                                     `json:"continueOnError,omitzero"`
	// Stable identifier used by template resolution v2+ to address this step as
	// {{key.field}}. Omitted keys are derived from the step name whenever a flow is
	// created. For template resolution v3, replacing an existing flow's full action
	// tree requires every step to carry an explicit key — a missing key is rejected
	// there, not derived, so a name change can never silently move a step's key.
	Key               param.Opt[string]                                                       `json:"key,omitzero"`
	NameOverride      param.Opt[string]                                                       `json:"nameOverride,omitzero"`
	RecordingEnabled  param.Opt[bool]                                                         `json:"recordingEnabled,omitzero"`
	OverridesProperty WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInputOverrides `json:"overrides,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInput) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Method, Name, Service are required.
type WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInputDraft struct {
	Method  string         `json:"method" api:"required"`
	Name    string         `json:"name" api:"required"`
	Service string         `json:"service" api:"required"`
	Params  map[string]any `json:"params,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInputDraft) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInputDraft
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInputDraft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInputOverrides struct {
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInputOverrides) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInputOverrides
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObjectActionFlowDraftActionInputOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Destination is required.
type WorkflowFlowValidateParamsBodyObjectDelivery struct {
	// Any of "mobilerun", "one_drive", "google_drive".
	Destination string `json:"destination,omitzero" api:"required"`
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
	Folder      param.Opt[string]                                     `json:"folder,omitzero"`
	Recording   WorkflowFlowValidateParamsBodyObjectDeliveryRecording `json:"recording,omitzero"`
	Screenshots any                                                   `json:"screenshots,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObjectDelivery) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObjectDelivery
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObjectDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObjectDelivery](
		"destination", "mobilerun", "one_drive", "google_drive",
	)
}

// The property Filename is required.
type WorkflowFlowValidateParamsBodyObjectDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObjectDeliveryRecording) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObjectDeliveryRecording
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObjectDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Mode is required.
type WorkflowFlowValidateParamsBodyObjectRecordingPolicy struct {
	// Any of "off", "flow", "selected_steps".
	Mode string `json:"mode,omitzero" api:"required"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObjectRecordingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObjectRecordingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObjectRecordingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObjectRecordingPolicy](
		"mode", "off", "flow", "selected_steps",
	)
}

// The property Draft is required.
type WorkflowFlowValidateParamsBodyObject2 struct {
	Draft    WorkflowFlowValidateParamsBodyObject2Draft    `json:"draft,omitzero" api:"required"`
	AgentID  param.Opt[string]                             `json:"agentId,omitzero" format:"uuid"`
	Contract WorkflowFlowValidateParamsBodyObject2Contract `json:"contract,omitzero"`
	Verify   WorkflowFlowValidateParamsBodyObject2Verify   `json:"verify,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Actions, Flow, Trigger are required.
type WorkflowFlowValidateParamsBodyObject2Draft struct {
	Actions []WorkflowFlowValidateParamsBodyObject2DraftAction `json:"actions,omitzero" api:"required"`
	Flow    WorkflowFlowValidateParamsBodyObject2DraftFlow     `json:"flow,omitzero" api:"required"`
	Trigger map[string]any                                     `json:"trigger,omitzero" api:"required"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2Draft) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2Draft
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2Draft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CatalogEntryID, Name, Position, Role are required.
type WorkflowFlowValidateParamsBodyObject2DraftAction struct {
	CatalogEntryID string            `json:"catalogEntryId" api:"required" format:"uuid"`
	Name           string            `json:"name" api:"required"`
	Position       int64             `json:"position" api:"required"`
	Role           string            `json:"role" api:"required"`
	Description    param.Opt[string] `json:"description,omitzero"`
	Params         map[string]any    `json:"params,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2DraftAction) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2DraftAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2DraftAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateParamsBodyObject2DraftFlow struct {
	CooldownSeconds         param.Opt[int64]  `json:"cooldownSeconds,omitzero"`
	Description             param.Opt[string] `json:"description,omitzero"`
	HealthMonitoringEnabled param.Opt[bool]   `json:"healthMonitoringEnabled,omitzero"`
	Name                    param.Opt[string] `json:"name,omitzero"`
	SelfHealingEnabled      param.Opt[bool]   `json:"selfHealingEnabled,omitzero"`
	SelfHealingMaxAttempts  param.Opt[int64]  `json:"selfHealingMaxAttempts,omitzero"`
	// Any of "flow", "device".
	CooldownScope string                                                 `json:"cooldownScope,omitzero"`
	Delivery      WorkflowFlowValidateParamsBodyObject2DraftFlowDelivery `json:"delivery,omitzero"`
	DeviceIDs     []string                                               `json:"deviceIds,omitzero" format:"uuid"`
	// Any of "device", "cloud".
	ExecutionMode   string                                                        `json:"executionMode,omitzero"`
	RecordingPolicy WorkflowFlowValidateParamsBodyObject2DraftFlowRecordingPolicy `json:"recordingPolicy,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2DraftFlow) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2DraftFlow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2DraftFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2DraftFlow](
		"cooldownScope", "flow", "device",
	)
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2DraftFlow](
		"executionMode", "device", "cloud",
	)
}

// The property Destination is required.
type WorkflowFlowValidateParamsBodyObject2DraftFlowDelivery struct {
	// Any of "mobilerun", "one_drive", "google_drive".
	Destination string `json:"destination,omitzero" api:"required"`
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
	Folder      param.Opt[string]                                               `json:"folder,omitzero"`
	Recording   WorkflowFlowValidateParamsBodyObject2DraftFlowDeliveryRecording `json:"recording,omitzero"`
	Screenshots any                                                             `json:"screenshots,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2DraftFlowDelivery) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2DraftFlowDelivery
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2DraftFlowDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2DraftFlowDelivery](
		"destination", "mobilerun", "one_drive", "google_drive",
	)
}

// The property Filename is required.
type WorkflowFlowValidateParamsBodyObject2DraftFlowDeliveryRecording struct {
	Filename string `json:"filename" api:"required"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2DraftFlowDeliveryRecording) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2DraftFlowDeliveryRecording
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2DraftFlowDeliveryRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Mode is required.
type WorkflowFlowValidateParamsBodyObject2DraftFlowRecordingPolicy struct {
	// Any of "off", "flow", "selected_steps".
	Mode string `json:"mode,omitzero" api:"required"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2DraftFlowRecordingPolicy) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2DraftFlowRecordingPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2DraftFlowRecordingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2DraftFlowRecordingPolicy](
		"mode", "off", "flow", "selected_steps",
	)
}

type WorkflowFlowValidateParamsBodyObject2Contract struct {
	// Ignored with agent_routine; derived from the action graph.
	IntegrationRole param.Opt[string] `json:"integrationRole,omitzero"`
	Slug            param.Opt[string] `json:"slug,omitzero"`
	// Ignored with agent_routine; derived from the action graph.
	UploadRole param.Opt[string] `json:"uploadRole,omitzero"`
	// Required without profile. Ignored with agent_routine; derived from
	// meta.sink.type.
	//
	// Any of "none", "file", "chat", "integration".
	Deliverable string                                            `json:"deliverable,omitzero"`
	Meta        WorkflowFlowValidateParamsBodyObject2ContractMeta `json:"meta,omitzero"`
	// Ignored with agent_routine; derived from the action graph.
	ProducerRoles []string `json:"producerRoles,omitzero"`
	// For agent_routine, contract.meta owns trigger activation, eventType,
	// scheduleRule and timezone. Draft timing copies are deprecated and ignored;
	// deploy still checks the types of supplied copies, while public full-draft
	// trigger objects retain their existing loose structure. For agent_routine,
	// contract.meta.flow owns healthMonitoringEnabled, selfHealingEnabled,
	// selfHealingMaxAttempts and explicit recordingMode. Draft setting copies are
	// deprecated and ignored after type checking. When metadata omits recordingMode,
	// the draft recordingPolicy still applies (implicitly off when absent).
	//
	// Any of "agent_routine".
	Profile string `json:"profile,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2Contract) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2Contract
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2Contract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2Contract](
		"deliverable", "none", "file", "chat", "integration",
	)
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2Contract](
		"profile", "agent_routine",
	)
}

// The properties Flow, Sink, Timezone, Trigger, Verification are required.
type WorkflowFlowValidateParamsBodyObject2ContractMeta struct {
	Timezone param.Opt[string]                                        `json:"timezone,omitzero" api:"required"`
	Flow     WorkflowFlowValidateParamsBodyObject2ContractMetaFlow    `json:"flow,omitzero" api:"required"`
	Sink     WorkflowFlowValidateParamsBodyObject2ContractMetaSink    `json:"sink,omitzero" api:"required"`
	Trigger  WorkflowFlowValidateParamsBodyObject2ContractMetaTrigger `json:"trigger,omitzero" api:"required"`
	// Any of "rules", "agent_review", "rules_then_review", "none".
	Verification   string                                                         `json:"verification,omitzero" api:"required"`
	OutputFilename param.Opt[string]                                              `json:"outputFilename,omitzero"`
	CloudActions   []WorkflowFlowValidateParamsBodyObject2ContractMetaCloudAction `json:"cloudActions,omitzero"`
	// Absent means version 1 (internal service.method names); version 2 signs public
	// step names. Signed records are never rewritten.
	//
	// Any of 1, 2.
	Version float64 `json:"version,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2ContractMeta) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2ContractMeta
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2ContractMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2ContractMeta](
		"verification", "rules", "agent_review", "rules_then_review", "none",
	)
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2ContractMeta](
		"version", 1, 2,
	)
}

// The properties HealthMonitoringEnabled, SelfHealingEnabled,
// SelfHealingMaxAttempts are required.
type WorkflowFlowValidateParamsBodyObject2ContractMetaFlow struct {
	HealthMonitoringEnabled bool  `json:"healthMonitoringEnabled" api:"required"`
	SelfHealingEnabled      bool  `json:"selfHealingEnabled" api:"required"`
	SelfHealingMaxAttempts  int64 `json:"selfHealingMaxAttempts" api:"required"`
	// Any of "off", "flow", "selected_steps".
	RecordingMode string `json:"recordingMode,omitzero"`
	// Any of "none", "message", "wake".
	ResultDelivery string `json:"resultDelivery,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2ContractMetaFlow) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2ContractMetaFlow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2ContractMetaFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2ContractMetaFlow](
		"recordingMode", "off", "flow", "selected_steps",
	)
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2ContractMetaFlow](
		"resultDelivery", "none", "message", "wake",
	)
}

// The property Type is required.
type WorkflowFlowValidateParamsBodyObject2ContractMetaSink struct {
	// Any of "chat", "file", "webhook".
	Type       string            `json:"type,omitzero" api:"required"`
	EndpointID param.Opt[string] `json:"endpointId,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2ContractMetaSink) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2ContractMetaSink
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2ContractMetaSink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2ContractMetaSink](
		"type", "chat", "file", "webhook",
	)
}

// The properties Cadence, EventType, ScheduleRule, Timezone, Type are required.
type WorkflowFlowValidateParamsBodyObject2ContractMetaTrigger struct {
	Cadence      param.Opt[string] `json:"cadence,omitzero" api:"required"`
	EventType    param.Opt[string] `json:"eventType,omitzero" api:"required"`
	Timezone     param.Opt[string] `json:"timezone,omitzero" api:"required"`
	ScheduleRule ScheduleRuleParam `json:"scheduleRule,omitzero" api:"required"`
	// Any of "schedule", "app_event", "device_event", "platform_event", "custom".
	Type string `json:"type,omitzero" api:"required"`
	// Optional; derived from trigger.type with agent_routine.
	//
	// Any of "event", "schedule", "custom".
	Activation string `json:"activation,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2ContractMetaTrigger) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2ContractMetaTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2ContractMetaTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2ContractMetaTrigger](
		"type", "schedule", "app_event", "device_event", "platform_event", "custom",
	)
	apijson.RegisterFieldValidator[WorkflowFlowValidateParamsBodyObject2ContractMetaTrigger](
		"activation", "event", "schedule", "custom",
	)
}

// The properties Method, Role are required.
type WorkflowFlowValidateParamsBodyObject2ContractMetaCloudAction struct {
	Method string `json:"method" api:"required"`
	Role   string `json:"role" api:"required"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2ContractMetaCloudAction) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2ContractMetaCloudAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2ContractMetaCloudAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateParamsBodyObject2Verify struct {
	DeviceID param.Opt[string] `json:"deviceId,omitzero" format:"uuid"`
	Payload  map[string]any    `json:"payload,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateParamsBodyObject2Verify) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateParamsBodyObject2Verify
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateParamsBodyObject2Verify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowValidateDocumentParams struct {
	// Routine document; format at GET /flows/document-schema
	Document any `json:"document,omitzero"`
	paramObj
}

func (r WorkflowFlowValidateDocumentParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowValidateDocumentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowValidateDocumentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowFlowVerifyParams struct {
	// Client-supplied idempotency key. A repeat request for the same (flow,
	// invocationId) returns the existing verification run (`deduplicated: true`,
	// HTTP 200) instead of enqueuing another.
	InvocationID string `json:"invocationId" api:"required" format:"uuid"`
	// Device to run the verification on. Must be one the flow is bound to. Optional
	// only when the flow is bound to exactly one device (that device is used);
	// otherwise required.
	DeviceID param.Opt[string] `json:"deviceId,omitzero" format:"uuid"`
	Payload  map[string]any    `json:"payload,omitzero"`
	paramObj
}

func (r WorkflowFlowVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFlowVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFlowVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
