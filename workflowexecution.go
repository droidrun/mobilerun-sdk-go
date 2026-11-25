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

// WorkflowExecutionService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowExecutionService] method instead.
type WorkflowExecutionService struct {
	options     []option.RequestOption
	Screenshots WorkflowExecutionScreenshotService
}

// NewWorkflowExecutionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowExecutionService(opts ...option.RequestOption) (r WorkflowExecutionService) {
	r = WorkflowExecutionService{}
	r.options = opts
	r.Screenshots = NewWorkflowExecutionScreenshotService(opts...)
	return
}

// Fetch a single flow execution by its ID, including its status, kind, result or
// error, and start/finish timestamps. Add `include=input` to get what each step
// ran with. Returns 404 if no execution matches.
func (r *WorkflowExecutionService) Get(ctx context.Context, executionID string, query WorkflowExecutionGetParams, opts ...option.RequestOption) (res *WorkflowExecutionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if executionID == "" {
		err = errors.New("missing required executionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("executions/%s", executionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return a paginated history of flow executions. Supports filtering by `flowId`,
// `triggerId`, `invocationId` (the client/verify key), `status`, `kind` (one or
// more execution kinds), and a `from`/`to` time range, plus free-text `search` and
// ordering by startedAt, finishedAt, or status.
func (r *WorkflowExecutionService) List(ctx context.Context, query WorkflowExecutionListParams, opts ...option.RequestOption) (res *pagination.Page[WorkflowExecutionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "executions"
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

// Return a paginated history of flow executions. Supports filtering by `flowId`,
// `triggerId`, `invocationId` (the client/verify key), `status`, `kind` (one or
// more execution kinds), and a `from`/`to` time range, plus free-text `search` and
// ordering by startedAt, finishedAt, or status.
func (r *WorkflowExecutionService) ListAutoPaging(ctx context.Context, query WorkflowExecutionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[WorkflowExecutionListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Stops the execution between steps and marks it cancelled. Idempotent-ish:
// already-terminal executions return 409.
func (r *WorkflowExecutionService) Abort(ctx context.Context, executionID string, opts ...option.RequestOption) (res *WorkflowExecutionAbortResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if executionID == "" {
		err = errors.New("missing required executionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("executions/%s/abort", executionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Return aggregate execution metrics — total count, counts by status, average
// duration, and the last execution time. Can be scoped by `flowId`, `triggerId`,
// and a `from`/`to` time range.
func (r *WorkflowExecutionService) GetMetrics(ctx context.Context, query WorkflowExecutionGetMetricsParams, opts ...option.RequestOption) (res *WorkflowExecutionGetMetricsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "executions/metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WorkflowExecutionGetResponse struct {
	Data WorkflowExecutionGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Current agent assignment of the owner-scoped parent flow; null when unassigned.
	AgentID   string `json:"agentId" api:"required" format:"uuid"`
	CreatedBy string `json:"createdBy" api:"required" format:"uuid"`
	// OneDrive/Google Drive delivery lifecycle for this run's recording, file, and
	// screenshot uploads, ordered by startedAt. Empty when the flow has no delivery
	// configured.
	Deliveries []WorkflowExecutionGetResponseDataDelivery `json:"deliveries" api:"required"`
	// Device this execution targets (the job's deviceId). Null for device-less
	// (event-only) runs.
	DeviceID string `json:"deviceId" api:"required" format:"uuid"`
	Error    string `json:"error" api:"required"`
	EventID  string `json:"eventId" api:"required" format:"uuid"`
	// Files produced by files.upload steps, plus files an agent.run step reported on
	// its terminal response (agent-created output or a workflow upload minted during
	// the turn); derived server-side at read time.
	Files      []WorkflowExecutionGetResponseDataFile `json:"files" api:"required"`
	FinishedAt time.Time                              `json:"finishedAt" api:"required" format:"date-time"`
	FlowID     string                                 `json:"flowId" api:"required" format:"uuid"`
	FlowName   string                                 `json:"flowName" api:"required"`
	// Client/verify invocation key this row belongs to. Set on live custom fires (one
	// row per device fan-out) and on verification runs; null for event/schedule live
	// rows.
	InvocationID string `json:"invocationId" api:"required" format:"uuid"`
	// Any of "live", "dry_run", "verification".
	Kind string `json:"kind" api:"required"`
	// Live progress; null for runs started before this feature.
	Progress          WorkflowExecutionGetResponseDataProgress `json:"progress" api:"required"`
	RecordingDeviceID string                                   `json:"recordingDeviceId" api:"required" format:"uuid"`
	// Device recording id for this execution, set once recording starts. Null when the
	// flow has recording disabled, no device is bound, or the recording failed to
	// start.
	RecordingID string `json:"recordingId" api:"required" format:"uuid"`
	// Durable recording segments ordered by step/loop coordinate and retry attempt.
	// Whole-flow recordings use -1 for every coordinate.
	Recordings []WorkflowExecutionGetResponseDataRecording `json:"recordings" api:"required"`
	Repair     WorkflowExecutionGetResponseDataRepair      `json:"repair" api:"required"`
	// Screenshots captured by tasks.run/agent.run steps and saved by scripts into the
	// flow screenshots recording (source recording), ordered by seq. Image bytes are
	// never returned here — fetch a fresh signed URL via GET
	// /executions/{id}/screenshots/{screenshotId}.
	Screenshots []WorkflowExecutionGetResponseDataScreenshot `json:"screenshots" api:"required"`
	StartedAt   time.Time                                    `json:"startedAt" api:"required" format:"date-time"`
	// Any of "pending", "running", "success", "failed", "cancelled", "skipped",
	// "invalid".
	Status      string `json:"status" api:"required"`
	TriggerID   string `json:"triggerId" api:"required" format:"uuid"`
	TriggerName string `json:"triggerName" api:"required"`
	// Opaque result blob ({ steps: [...] } or historical dry-run { actions: [...] }).
	// Steps and nested loop/action children carry `uses` (authored step name,
	// including hidden executable steps; null when unavailable). Admin execution list
	// responses do not carry `uses`. The retained `service` and `method` fields are
	// deprecated; use `uses`. Each step additionally carries a `verdict` field ({
	// outcome, summary, reason? } | null) when it is an agent.run step that opted into
	// a verdict — null otherwise. Steps of current executions also carry a `status`
	// string (e.g. success/failed/stopped); it is optional and absent on older
	// executions, so clients must not assume its presence. With `include=input` on the
	// detail request each step (loop children included) also carries `input`; see that
	// parameter.
	Result any `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AgentID           respjson.Field
		CreatedBy         respjson.Field
		Deliveries        respjson.Field
		DeviceID          respjson.Field
		Error             respjson.Field
		EventID           respjson.Field
		Files             respjson.Field
		FinishedAt        respjson.Field
		FlowID            respjson.Field
		FlowName          respjson.Field
		InvocationID      respjson.Field
		Kind              respjson.Field
		Progress          respjson.Field
		RecordingDeviceID respjson.Field
		RecordingID       respjson.Field
		Recordings        respjson.Field
		Repair            respjson.Field
		Screenshots       respjson.Field
		StartedAt         respjson.Field
		Status            respjson.Field
		TriggerID         respjson.Field
		TriggerName       respjson.Field
		Result            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetResponseDataDelivery struct {
	// Any of "recording", "file", "screenshot".
	Artifact string `json:"artifact" api:"required"`
	// Any of "mobilerun", "one_drive", "google_drive".
	Destination string    `json:"destination" api:"required"`
	ErrorCode   string    `json:"errorCode" api:"required"`
	Filename    string    `json:"filename" api:"required"`
	FinishedAt  time.Time `json:"finishedAt" api:"required" format:"date-time"`
	Folder      string    `json:"folder" api:"required"`
	StartedAt   time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Any of "waiting", "uploading", "succeeded", "failed", "unknown", "cancelled".
	Status    string `json:"status" api:"required"`
	StepIndex int64  `json:"stepIndex" api:"required"`
	WebURL    string `json:"webUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Artifact    respjson.Field
		Destination respjson.Field
		ErrorCode   respjson.Field
		Filename    respjson.Field
		FinishedAt  respjson.Field
		Folder      respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		StepIndex   respjson.Field
		WebURL      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetResponseDataFile struct {
	FileID    string `json:"fileId" api:"required" format:"uuid"`
	Filename  string `json:"filename" api:"required"`
	MimeType  string `json:"mimeType" api:"required"`
	SizeBytes int64  `json:"sizeBytes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Filename    respjson.Field
		MimeType    respjson.Field
		SizeBytes   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetResponseDataFile) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetResponseDataFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Live progress; null for runs started before this feature.
type WorkflowExecutionGetResponseDataProgress struct {
	CurrentIndex int64                                          `json:"currentIndex" api:"required"`
	Steps        []WorkflowExecutionGetResponseDataProgressStep `json:"steps" api:"required"`
	Total        int64                                          `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentIndex respjson.Field
		Steps        respjson.Field
		Total        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetResponseDataProgress) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetResponseDataProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetResponseDataProgressStep struct {
	FinishedAt time.Time `json:"finishedAt" api:"required" format:"date-time"`
	Index      int64     `json:"index" api:"required"`
	// Deprecated: use uses.
	//
	// Deprecated: deprecated
	Method string `json:"method" api:"required"`
	Name   string `json:"name" api:"required"`
	// Deprecated: use uses.
	//
	// Deprecated: deprecated
	Service   string    `json:"service" api:"required"`
	SessionID string    `json:"sessionId" api:"required"`
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Any of "pending", "running", "success", "failed", "skipped", "cancelled".
	Status string `json:"status" api:"required"`
	// Authored step name; null when this step has no name. Includes hidden executable
	// steps.
	Uses string `json:"uses" api:"required"`
	// Zero-based retry attempt of the execution job during which this step last ran (0
	// = initial attempt). Not a per-step retry counter; may undercount after crash
	// recovery.
	Attempt int64 `json:"attempt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishedAt  respjson.Field
		Index       respjson.Field
		Method      respjson.Field
		Name        respjson.Field
		Service     respjson.Field
		SessionID   respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Uses        respjson.Field
		Attempt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetResponseDataProgressStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetResponseDataProgressStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetResponseDataRecording struct {
	ID                string `json:"id" api:"required" format:"uuid"`
	Attempt           int64  `json:"attempt" api:"required"`
	ChildIndex        int64  `json:"childIndex" api:"required"`
	FlowActionID      string `json:"flowActionId" api:"required" format:"uuid"`
	IterationIndex    int64  `json:"iterationIndex" api:"required"`
	LastError         string `json:"lastError" api:"required"`
	ParentIndex       int64  `json:"parentIndex" api:"required"`
	RecordingDeviceID string `json:"recordingDeviceId" api:"required" format:"uuid"`
	RecordingID       string `json:"recordingId" api:"required" format:"uuid"`
	// Any of "flow", "step".
	Scope     string    `json:"scope" api:"required"`
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Any of "starting", "recording", "stopping", "stopped", "failed".
	Status    string    `json:"status" api:"required"`
	StepIndex int64     `json:"stepIndex" api:"required"`
	StoppedAt time.Time `json:"stoppedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Attempt           respjson.Field
		ChildIndex        respjson.Field
		FlowActionID      respjson.Field
		IterationIndex    respjson.Field
		LastError         respjson.Field
		ParentIndex       respjson.Field
		RecordingDeviceID respjson.Field
		RecordingID       respjson.Field
		Scope             respjson.Field
		StartedAt         respjson.Field
		Status            respjson.Field
		StepIndex         respjson.Field
		StoppedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetResponseDataRecording) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetResponseDataRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetResponseDataRepair struct {
	ChatSessionID string `json:"chatSessionId" api:"required" format:"uuid"`
	Episode       int64  `json:"episode" api:"required"`
	EpisodeID     string `json:"episodeId" api:"required" format:"uuid"`
	// Any of "pending", "running", "canary", "promoting", "repaired", "failed",
	// "escalated".
	Status   string                                        `json:"status" api:"required"`
	Terminal bool                                          `json:"terminal" api:"required"`
	Verdict  WorkflowExecutionGetResponseDataRepairVerdict `json:"verdict" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatSessionID respjson.Field
		Episode       respjson.Field
		EpisodeID     respjson.Field
		Status        respjson.Field
		Terminal      respjson.Field
		Verdict       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetResponseDataRepair) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetResponseDataRepair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetResponseDataRepairVerdict struct {
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
func (r WorkflowExecutionGetResponseDataRepairVerdict) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetResponseDataRepairVerdict) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetResponseDataScreenshot struct {
	ID             string    `json:"id" api:"required" format:"uuid"`
	CapturedAt     time.Time `json:"capturedAt" api:"required" format:"date-time"`
	IterationIndex int64     `json:"iterationIndex" api:"required"`
	MimeType       string    `json:"mimeType" api:"required"`
	Seq            int64     `json:"seq" api:"required"`
	// Any of "task", "agent", "recording".
	Source    string `json:"source" api:"required"`
	StepIndex int64  `json:"stepIndex" api:"required"`
	StepName  string `json:"stepName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CapturedAt     respjson.Field
		IterationIndex respjson.Field
		MimeType       respjson.Field
		Seq            respjson.Field
		Source         respjson.Field
		StepIndex      respjson.Field
		StepName       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetResponseDataScreenshot) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetResponseDataScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionListResponse struct {
	ID        string `json:"id" api:"required" format:"uuid"`
	CreatedBy string `json:"createdBy" api:"required" format:"uuid"`
	// Device this execution targets (the job's deviceId). Null for device-less
	// (event-only) runs.
	DeviceID   string    `json:"deviceId" api:"required" format:"uuid"`
	Error      string    `json:"error" api:"required"`
	EventID    string    `json:"eventId" api:"required" format:"uuid"`
	FinishedAt time.Time `json:"finishedAt" api:"required" format:"date-time"`
	FlowID     string    `json:"flowId" api:"required" format:"uuid"`
	FlowName   string    `json:"flowName" api:"required"`
	// Client/verify invocation key this row belongs to. Set on live custom fires (one
	// row per device fan-out) and on verification runs; null for event/schedule live
	// rows.
	InvocationID string `json:"invocationId" api:"required" format:"uuid"`
	// Any of "live", "dry_run", "verification".
	Kind              WorkflowExecutionListResponseKind `json:"kind" api:"required"`
	RecordingDeviceID string                            `json:"recordingDeviceId" api:"required" format:"uuid"`
	// Device recording id for this execution, set once recording starts. Null when the
	// flow has recording disabled, no device is bound, or the recording failed to
	// start.
	RecordingID string `json:"recordingId" api:"required" format:"uuid"`
	// Durable recording segments ordered by step/loop coordinate and retry attempt.
	// Whole-flow recordings use -1 for every coordinate.
	Recordings []WorkflowExecutionListResponseRecording `json:"recordings" api:"required"`
	StartedAt  time.Time                                `json:"startedAt" api:"required" format:"date-time"`
	// Any of "pending", "running", "success", "failed", "cancelled", "skipped",
	// "invalid".
	Status      WorkflowExecutionListResponseStatus `json:"status" api:"required"`
	TriggerID   string                              `json:"triggerId" api:"required" format:"uuid"`
	TriggerName string                              `json:"triggerName" api:"required"`
	// Opaque result blob ({ steps: [...] } or historical dry-run { actions: [...] }).
	// Steps and nested loop/action children carry `uses` (authored step name,
	// including hidden executable steps; null when unavailable). Admin execution list
	// responses do not carry `uses`. The retained `service` and `method` fields are
	// deprecated; use `uses`. Each step additionally carries a `verdict` field ({
	// outcome, summary, reason? } | null) when it is an agent.run step that opted into
	// a verdict — null otherwise. Steps of current executions also carry a `status`
	// string (e.g. success/failed/stopped); it is optional and absent on older
	// executions, so clients must not assume its presence. With `include=input` on the
	// detail request each step (loop children included) also carries `input`; see that
	// parameter.
	Result any `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedBy         respjson.Field
		DeviceID          respjson.Field
		Error             respjson.Field
		EventID           respjson.Field
		FinishedAt        respjson.Field
		FlowID            respjson.Field
		FlowName          respjson.Field
		InvocationID      respjson.Field
		Kind              respjson.Field
		RecordingDeviceID respjson.Field
		RecordingID       respjson.Field
		Recordings        respjson.Field
		StartedAt         respjson.Field
		Status            respjson.Field
		TriggerID         respjson.Field
		TriggerName       respjson.Field
		Result            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionListResponseKind string

const (
	WorkflowExecutionListResponseKindLive         WorkflowExecutionListResponseKind = "live"
	WorkflowExecutionListResponseKindDryRun       WorkflowExecutionListResponseKind = "dry_run"
	WorkflowExecutionListResponseKindVerification WorkflowExecutionListResponseKind = "verification"
)

type WorkflowExecutionListResponseRecording struct {
	ID                string `json:"id" api:"required" format:"uuid"`
	Attempt           int64  `json:"attempt" api:"required"`
	ChildIndex        int64  `json:"childIndex" api:"required"`
	FlowActionID      string `json:"flowActionId" api:"required" format:"uuid"`
	IterationIndex    int64  `json:"iterationIndex" api:"required"`
	LastError         string `json:"lastError" api:"required"`
	ParentIndex       int64  `json:"parentIndex" api:"required"`
	RecordingDeviceID string `json:"recordingDeviceId" api:"required" format:"uuid"`
	RecordingID       string `json:"recordingId" api:"required" format:"uuid"`
	// Any of "flow", "step".
	Scope     string    `json:"scope" api:"required"`
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Any of "starting", "recording", "stopping", "stopped", "failed".
	Status    string    `json:"status" api:"required"`
	StepIndex int64     `json:"stepIndex" api:"required"`
	StoppedAt time.Time `json:"stoppedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Attempt           respjson.Field
		ChildIndex        respjson.Field
		FlowActionID      respjson.Field
		IterationIndex    respjson.Field
		LastError         respjson.Field
		ParentIndex       respjson.Field
		RecordingDeviceID respjson.Field
		RecordingID       respjson.Field
		Scope             respjson.Field
		StartedAt         respjson.Field
		Status            respjson.Field
		StepIndex         respjson.Field
		StoppedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionListResponseRecording) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionListResponseRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionListResponseStatus string

const (
	WorkflowExecutionListResponseStatusPending   WorkflowExecutionListResponseStatus = "pending"
	WorkflowExecutionListResponseStatusRunning   WorkflowExecutionListResponseStatus = "running"
	WorkflowExecutionListResponseStatusSuccess   WorkflowExecutionListResponseStatus = "success"
	WorkflowExecutionListResponseStatusFailed    WorkflowExecutionListResponseStatus = "failed"
	WorkflowExecutionListResponseStatusCancelled WorkflowExecutionListResponseStatus = "cancelled"
	WorkflowExecutionListResponseStatusSkipped   WorkflowExecutionListResponseStatus = "skipped"
	WorkflowExecutionListResponseStatusInvalid   WorkflowExecutionListResponseStatus = "invalid"
)

type WorkflowExecutionAbortResponse struct {
	Data WorkflowExecutionAbortResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionAbortResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionAbortResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionAbortResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Device this execution targets (the job's deviceId). Null for device-less
	// (event-only) runs.
	DeviceID   string    `json:"deviceId" api:"required" format:"uuid"`
	Error      string    `json:"error" api:"required"`
	EventID    string    `json:"eventId" api:"required" format:"uuid"`
	FinishedAt time.Time `json:"finishedAt" api:"required" format:"date-time"`
	FlowID     string    `json:"flowId" api:"required" format:"uuid"`
	FlowName   string    `json:"flowName" api:"required"`
	// Client/verify invocation key this row belongs to. Set on live custom fires (one
	// row per device fan-out) and on verification runs; null for event/schedule live
	// rows.
	InvocationID string `json:"invocationId" api:"required" format:"uuid"`
	// Any of "live", "dry_run", "verification".
	Kind              string `json:"kind" api:"required"`
	RecordingDeviceID string `json:"recordingDeviceId" api:"required" format:"uuid"`
	// Device recording id for this execution, set once recording starts. Null when the
	// flow has recording disabled, no device is bound, or the recording failed to
	// start.
	RecordingID string    `json:"recordingId" api:"required" format:"uuid"`
	StartedAt   time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Any of "pending", "running", "success", "failed", "cancelled", "skipped",
	// "invalid".
	Status      string `json:"status" api:"required"`
	TriggerID   string `json:"triggerId" api:"required" format:"uuid"`
	TriggerName string `json:"triggerName" api:"required"`
	// Opaque result blob ({ steps: [...] } or historical dry-run { actions: [...] }).
	// Steps and nested loop/action children carry `uses` (authored step name,
	// including hidden executable steps; null when unavailable). Admin execution list
	// responses do not carry `uses`. The retained `service` and `method` fields are
	// deprecated; use `uses`. Each step additionally carries a `verdict` field ({
	// outcome, summary, reason? } | null) when it is an agent.run step that opted into
	// a verdict — null otherwise. Steps of current executions also carry a `status`
	// string (e.g. success/failed/stopped); it is optional and absent on older
	// executions, so clients must not assume its presence. With `include=input` on the
	// detail request each step (loop children included) also carries `input`; see that
	// parameter.
	Result any `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		DeviceID          respjson.Field
		Error             respjson.Field
		EventID           respjson.Field
		FinishedAt        respjson.Field
		FlowID            respjson.Field
		FlowName          respjson.Field
		InvocationID      respjson.Field
		Kind              respjson.Field
		RecordingDeviceID respjson.Field
		RecordingID       respjson.Field
		StartedAt         respjson.Field
		Status            respjson.Field
		TriggerID         respjson.Field
		TriggerName       respjson.Field
		Result            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionAbortResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionAbortResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetMetricsResponse struct {
	Data WorkflowExecutionGetMetricsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetMetricsResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetMetricsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetMetricsResponseData struct {
	AvgDurationMs   float64                                         `json:"avgDurationMs" api:"required"`
	ByStatus        WorkflowExecutionGetMetricsResponseDataByStatus `json:"byStatus" api:"required"`
	LastExecutionAt time.Time                                       `json:"lastExecutionAt" api:"required" format:"date-time"`
	Total           int64                                           `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgDurationMs   respjson.Field
		ByStatus        respjson.Field
		LastExecutionAt respjson.Field
		Total           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetMetricsResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetMetricsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetMetricsResponseDataByStatus struct {
	Cancelled int64 `json:"cancelled" api:"required"`
	Failed    int64 `json:"failed" api:"required"`
	Invalid   int64 `json:"invalid" api:"required"`
	Pending   int64 `json:"pending" api:"required"`
	Running   int64 `json:"running" api:"required"`
	Skipped   int64 `json:"skipped" api:"required"`
	Success   int64 `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled   respjson.Field
		Failed      respjson.Field
		Invalid     respjson.Field
		Pending     respjson.Field
		Running     respjson.Field
		Skipped     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowExecutionGetMetricsResponseDataByStatus) RawJSON() string { return r.JSON.raw }
func (r *WorkflowExecutionGetMetricsResponseDataByStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowExecutionGetParams struct {
	// Extra data to add to each step of `result.steps` (loop children included).
	// `input` adds a field `input` per step: what the step ran with, resolved after
	// templating and without values the platform adds (ids, session and idempotency
	// keys, script state). For sandbox steps, nested `input` is the effective input
	// when available, otherwise the submitted input. The value is the input object, a
	// marker `{ "notRecorded": "too_large" | "invalid" | "unavailable" }` (`too_large`
	// adds `bytes`) when it could not be stored or exceeds 64 KiB, or null for a step
	// without a recorded input, such as a run from before inputs were stored. Secrets
	// are redacted before storing, best effort and by key name only: a value is
	// replaced with `[REDACTED]` at any depth when its key, read as words
	// (`clientSecret`, `x-auth-token`), contains `password`, `passwd`, `passphrase`,
	// `secret`, `token`, `authorization`, `cookie`, `credential`, `signature`,
	// `api key`, `private key`, `client key` or `card number`, or is exactly `pwd`,
	// `otp`, `totp`, `pin`, `cvv` or `cvc`. `Bearer`, `Basic` and `Token`
	// authorization values in strings are masked too. Secrets inside other strings
	// (for example agent prompts) or under other key names are not detected. The
	// content of file uploads is never stored, only its size, encoding and type.
	// Without `include` the response has no `input` fields.
	//
	// Any of "input".
	Include WorkflowExecutionGetParamsInclude `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowExecutionGetParams]'s query parameters as
// `url.Values`.
func (r WorkflowExecutionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Extra data to add to each step of `result.steps` (loop children included).
// `input` adds a field `input` per step: what the step ran with, resolved after
// templating and without values the platform adds (ids, session and idempotency
// keys, script state). For sandbox steps, nested `input` is the effective input
// when available, otherwise the submitted input. The value is the input object, a
// marker `{ "notRecorded": "too_large" | "invalid" | "unavailable" }` (`too_large`
// adds `bytes`) when it could not be stored or exceeds 64 KiB, or null for a step
// without a recorded input, such as a run from before inputs were stored. Secrets
// are redacted before storing, best effort and by key name only: a value is
// replaced with `[REDACTED]` at any depth when its key, read as words
// (`clientSecret`, `x-auth-token`), contains `password`, `passwd`, `passphrase`,
// `secret`, `token`, `authorization`, `cookie`, `credential`, `signature`,
// `api key`, `private key`, `client key` or `card number`, or is exactly `pwd`,
// `otp`, `totp`, `pin`, `cvv` or `cvc`. `Bearer`, `Basic` and `Token`
// authorization values in strings are masked too. Secrets inside other strings
// (for example agent prompts) or under other key names are not detected. The
// content of file uploads is never stored, only its size, encoding and type.
// Without `include` the response has no `input` fields.
type WorkflowExecutionGetParamsInclude string

const (
	WorkflowExecutionGetParamsIncludeInput WorkflowExecutionGetParamsInclude = "input"
)

type WorkflowExecutionListParams struct {
	From         param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	To           param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	FlowID       param.Opt[string]    `query:"flowId,omitzero" format:"uuid" json:"-"`
	InvocationID param.Opt[string]    `query:"invocationId,omitzero" format:"uuid" json:"-"`
	Page         param.Opt[int64]     `query:"page,omitzero" json:"-"`
	PageSize     param.Opt[int64]     `query:"pageSize,omitzero" json:"-"`
	Search       param.Opt[string]    `query:"search,omitzero" json:"-"`
	TriggerID    param.Opt[string]    `query:"triggerId,omitzero" format:"uuid" json:"-"`
	// Only include these execution kinds. Omit to include all kinds.
	//
	// Any of "live", "dry_run", "verification".
	Kind []string `query:"kind,omitzero" json:"-"`
	// Any of "startedAt", "finishedAt", "status".
	OrderBy WorkflowExecutionListParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// Any of "asc", "desc".
	OrderByDirection WorkflowExecutionListParamsOrderByDirection `query:"orderByDirection,omitzero" json:"-"`
	// Any of "pending", "running", "success", "failed", "cancelled", "skipped",
	// "invalid".
	Status WorkflowExecutionListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowExecutionListParams]'s query parameters as
// `url.Values`.
func (r WorkflowExecutionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowExecutionListParamsOrderBy string

const (
	WorkflowExecutionListParamsOrderByStartedAt  WorkflowExecutionListParamsOrderBy = "startedAt"
	WorkflowExecutionListParamsOrderByFinishedAt WorkflowExecutionListParamsOrderBy = "finishedAt"
	WorkflowExecutionListParamsOrderByStatus     WorkflowExecutionListParamsOrderBy = "status"
)

type WorkflowExecutionListParamsOrderByDirection string

const (
	WorkflowExecutionListParamsOrderByDirectionAsc  WorkflowExecutionListParamsOrderByDirection = "asc"
	WorkflowExecutionListParamsOrderByDirectionDesc WorkflowExecutionListParamsOrderByDirection = "desc"
)

type WorkflowExecutionListParamsStatus string

const (
	WorkflowExecutionListParamsStatusPending   WorkflowExecutionListParamsStatus = "pending"
	WorkflowExecutionListParamsStatusRunning   WorkflowExecutionListParamsStatus = "running"
	WorkflowExecutionListParamsStatusSuccess   WorkflowExecutionListParamsStatus = "success"
	WorkflowExecutionListParamsStatusFailed    WorkflowExecutionListParamsStatus = "failed"
	WorkflowExecutionListParamsStatusCancelled WorkflowExecutionListParamsStatus = "cancelled"
	WorkflowExecutionListParamsStatusSkipped   WorkflowExecutionListParamsStatus = "skipped"
	WorkflowExecutionListParamsStatusInvalid   WorkflowExecutionListParamsStatus = "invalid"
)

type WorkflowExecutionGetMetricsParams struct {
	From      param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	To        param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	FlowID    param.Opt[string]    `query:"flowId,omitzero" format:"uuid" json:"-"`
	TriggerID param.Opt[string]    `query:"triggerId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowExecutionGetMetricsParams]'s query parameters as
// `url.Values`.
func (r WorkflowExecutionGetMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
