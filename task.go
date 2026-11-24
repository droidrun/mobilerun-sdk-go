// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"encoding/json"
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
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared/constant"
)

// Tasks API
//
// TaskService contains methods and other services that help with interacting with
// the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskService] method instead.
type TaskService struct {
	options []option.RequestOption
	// Tasks API
	Screenshots TaskScreenshotService
	// Tasks API
	UiStates TaskUiStateService
}

// NewTaskService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaskService(opts ...option.RequestOption) (r TaskService) {
	r = TaskService{}
	r.options = opts
	r.Screenshots = NewTaskScreenshotService(opts...)
	r.UiStates = NewTaskUiStateService(opts...)
	return
}

// Get full details of a task by ID.
func (r *TaskService) Get(ctx context.Context, taskID string, opts ...option.RequestOption) (res *TaskGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tasks/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List tasks with optional filtering, sorting, and pagination.
func (r *TaskService) List(ctx context.Context, query TaskListParams, opts ...option.RequestOption) (res *TaskListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Attach to a running task and receive its events as an SSE stream.
func (r *TaskService) Attach(ctx context.Context, taskID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return err
	}
	path := fmt.Sprintf("tasks/%s/attach", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Get the status of a task.
func (r *TaskService) GetStatus(ctx context.Context, taskID string, opts ...option.RequestOption) (res *TaskGetStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tasks/%s/status", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the trajectory of a task.
func (r *TaskService) GetTrajectory(ctx context.Context, taskID string, opts ...option.RequestOption) (res *TaskGetTrajectoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tasks/%s/trajectory", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create and dispatch a new agent task. Returns the task ID and device stream
// details.
func (r *TaskService) Run(ctx context.Context, body TaskRunParams, opts ...option.RequestOption) (res *TaskRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create and dispatch a new agent task, returning an SSE stream of task events.
// Cancels the task if the client disconnects.
func (r *TaskService) RunStreamed(ctx context.Context, body TaskRunStreamedParams, opts ...option.RequestOption) (res *TaskRunStreamedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "tasks/stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send a message to a running agent task. The message ID is delivered via SSE
// (UserMessageEvent with action=queued).
func (r *TaskService) SendMessage(ctx context.Context, taskID string, body TaskSendMessageParams, opts ...option.RequestOption) (res *TaskSendMessageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tasks/%s/message", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Cancel a running task. Returns an error if the task is already in a terminal
// state.
func (r *TaskService) Stop(ctx context.Context, taskID string, opts ...option.RequestOption) (res *TaskStopResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tasks/%s/cancel", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type TaskGetResponse struct {
	// The task
	Task TaskGetResponseTask `json:"task" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Task        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The task
type TaskGetResponseTask struct {
	ID        string `json:"id" api:"required" format:"uuid"`
	DeviceID  string `json:"deviceId" api:"required" format:"uuid"`
	DisplayID int64  `json:"displayId" api:"required"`
	// The LLM model identifier to use for the task (e.g. 'gemini/gemini-2.5-flash')
	LlmModel string `json:"llmModel" api:"required"`
	OwnerID  string `json:"ownerId" api:"required"`
	// Any of "prepared", "queued", "created", "running", "cancelling", "completed",
	// "failed", "cancelled".
	Status    string `json:"status" api:"required"`
	Task      string `json:"task" api:"required"`
	TmpDevice bool   `json:"tmpDevice" api:"required"`
	// Deprecated: use ownerId (tenancy) / createdBy (actor).
	//
	// Deprecated: deprecated
	UserID            string                          `json:"userId" api:"required"`
	Accessibility     bool                            `json:"accessibility"`
	Apps              []string                        `json:"apps"`
	CancelRequestedAt time.Time                       `json:"cancelRequestedAt" api:"nullable" format:"date-time"`
	ClaimedAt         time.Time                       `json:"claimedAt" api:"nullable" format:"date-time"`
	ContinueOnFailure bool                            `json:"continueOnFailure"`
	CreatedAt         time.Time                       `json:"createdAt" format:"date-time"`
	CreatedBy         string                          `json:"createdBy" api:"nullable"`
	Credentials       []TaskGetResponseTaskCredential `json:"credentials"`
	CreditsUsed       float64                         `json:"creditsUsed" api:"nullable"`
	DispatchedAt      time.Time                       `json:"dispatchedAt" api:"nullable" format:"date-time"`
	ExecutionTimeout  int64                           `json:"executionTimeout"`
	Files             []string                        `json:"files"`
	FinishedAt        time.Time                       `json:"finishedAt" api:"nullable" format:"date-time"`
	// Whether human-like input was requested. Always equal to the deprecated stealth
	// field.
	Humanize     bool           `json:"humanize"`
	MaxSteps     int64          `json:"maxSteps"`
	Message      string         `json:"message" api:"nullable"`
	Output       map[string]any `json:"output" api:"nullable"`
	OutputSchema map[string]any `json:"outputSchema" api:"nullable"`
	// A human approval the running task is waiting for.
	PendingApproval   TaskGetResponseTaskPendingApproval `json:"pendingApproval" api:"nullable"`
	Reasoning         bool                               `json:"reasoning"`
	RecordingDeviceID string                             `json:"recordingDeviceId" api:"nullable" format:"uuid"`
	// Record device video for the whole task and persist a retrievable reference
	RecordingEnabled bool   `json:"recordingEnabled"`
	RecordingID      string `json:"recordingId" api:"nullable" format:"uuid"`
	// Where the task came from: 'api' for tasks created via POST /tasks, 'agent' for
	// tasks spawned by an agent step or created with internal=true. 'source'
	// classifies visibility only: tasks created with internal=true stay controllable
	// via this API, while tasks spawned by an agent step are readable (status,
	// trajectory, media) but not controllable.
	//
	// Any of "api", "agent".
	Source string `json:"source"`
	// Deprecated: use humanize. Both names must agree when supplied together.
	//
	// Deprecated: deprecated
	Stealth   bool   `json:"stealth"`
	Steps     int64  `json:"steps" api:"nullable"`
	StreamURL string `json:"streamUrl" api:"nullable"`
	// LLM model used by sub-agent roles: executor, app_opener, structured_output
	SubagentModel string `json:"subagentModel"`
	Succeeded     bool   `json:"succeeded" api:"nullable"`
	// Deprecated and ignored. Sampling behavior is controlled by the model provider.
	//
	// Deprecated: deprecated
	Temperature float64   `json:"temperature"`
	UpdatedAt   time.Time `json:"updatedAt" format:"date-time"`
	Vision      bool      `json:"vision"`
	// Any of "US", "BR", "FR", "DE", "IN", "JP", "KR", "ZA".
	VpnCountry string `json:"vpnCountry" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		DeviceID          respjson.Field
		DisplayID         respjson.Field
		LlmModel          respjson.Field
		OwnerID           respjson.Field
		Status            respjson.Field
		Task              respjson.Field
		TmpDevice         respjson.Field
		UserID            respjson.Field
		Accessibility     respjson.Field
		Apps              respjson.Field
		CancelRequestedAt respjson.Field
		ClaimedAt         respjson.Field
		ContinueOnFailure respjson.Field
		CreatedAt         respjson.Field
		CreatedBy         respjson.Field
		Credentials       respjson.Field
		CreditsUsed       respjson.Field
		DispatchedAt      respjson.Field
		ExecutionTimeout  respjson.Field
		Files             respjson.Field
		FinishedAt        respjson.Field
		Humanize          respjson.Field
		MaxSteps          respjson.Field
		Message           respjson.Field
		Output            respjson.Field
		OutputSchema      respjson.Field
		PendingApproval   respjson.Field
		Reasoning         respjson.Field
		RecordingDeviceID respjson.Field
		RecordingEnabled  respjson.Field
		RecordingID       respjson.Field
		Source            respjson.Field
		Stealth           respjson.Field
		Steps             respjson.Field
		StreamURL         respjson.Field
		SubagentModel     respjson.Field
		Succeeded         respjson.Field
		Temperature       respjson.Field
		UpdatedAt         respjson.Field
		Vision            respjson.Field
		VpnCountry        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetResponseTask) RawJSON() string { return r.JSON.raw }
func (r *TaskGetResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetResponseTaskCredential struct {
	CredentialNames []string `json:"credentialNames" api:"required"`
	PackageName     string   `json:"packageName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CredentialNames respjson.Field
		PackageName     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetResponseTaskCredential) RawJSON() string { return r.JSON.raw }
func (r *TaskGetResponseTaskCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A human approval the running task is waiting for.
type TaskGetResponseTaskPendingApproval struct {
	// Identifies the approval
	ApprovalID string `json:"approvalId" api:"required" format:"uuid"`
	// Kind-specific, non-secret context (payment_card: label, secretGroup,
	// targetPackage, targetHost?)
	Details map[string]any `json:"details" api:"required"`
	// Without an answer by this time the task fails
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// What is being approved, e.g. payment_card; new kinds may be added
	Kind string `json:"kind" api:"required"`
	// Device the task runs on, so a client can show it next to the approval
	DeviceID string `json:"deviceId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalID  respjson.Field
		Details     respjson.Field
		ExpiresAt   respjson.Field
		Kind        respjson.Field
		DeviceID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetResponseTaskPendingApproval) RawJSON() string { return r.JSON.raw }
func (r *TaskGetResponseTaskPendingApproval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListResponse struct {
	// The paginated items
	Items []TaskListResponseItem `json:"items" api:"required"`
	// Pagination metadata
	Pagination shared.PaginationMeta `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Task representation for list endpoints — omits the large trajectory field.
type TaskListResponseItem struct {
	ID        string `json:"id" api:"required" format:"uuid"`
	DeviceID  string `json:"deviceId" api:"required" format:"uuid"`
	DisplayID int64  `json:"displayId" api:"required"`
	// The LLM model identifier to use for the task (e.g. 'gemini/gemini-2.5-flash')
	LlmModel string `json:"llmModel" api:"required"`
	OwnerID  string `json:"ownerId" api:"required"`
	// Any of "prepared", "queued", "created", "running", "cancelling", "completed",
	// "failed", "cancelled".
	Status    string `json:"status" api:"required"`
	Task      string `json:"task" api:"required"`
	TmpDevice bool   `json:"tmpDevice" api:"required"`
	// Deprecated: use ownerId (tenancy) / createdBy (actor).
	//
	// Deprecated: deprecated
	UserID            string                           `json:"userId" api:"required"`
	Accessibility     bool                             `json:"accessibility"`
	Apps              []string                         `json:"apps"`
	CancelRequestedAt time.Time                        `json:"cancelRequestedAt" api:"nullable" format:"date-time"`
	ClaimedAt         time.Time                        `json:"claimedAt" api:"nullable" format:"date-time"`
	ContinueOnFailure bool                             `json:"continueOnFailure"`
	CreatedAt         time.Time                        `json:"createdAt" format:"date-time"`
	CreatedBy         string                           `json:"createdBy" api:"nullable"`
	Credentials       []TaskListResponseItemCredential `json:"credentials"`
	CreditsUsed       float64                          `json:"creditsUsed" api:"nullable"`
	DispatchedAt      time.Time                        `json:"dispatchedAt" api:"nullable" format:"date-time"`
	ExecutionTimeout  int64                            `json:"executionTimeout"`
	Files             []string                         `json:"files"`
	FinishedAt        time.Time                        `json:"finishedAt" api:"nullable" format:"date-time"`
	// Whether human-like input was requested. Always equal to the deprecated stealth
	// field.
	Humanize     bool           `json:"humanize"`
	MaxSteps     int64          `json:"maxSteps"`
	Message      string         `json:"message" api:"nullable"`
	Output       map[string]any `json:"output" api:"nullable"`
	OutputSchema map[string]any `json:"outputSchema" api:"nullable"`
	// A human approval the running task is waiting for.
	PendingApproval   TaskListResponseItemPendingApproval `json:"pendingApproval" api:"nullable"`
	Reasoning         bool                                `json:"reasoning"`
	RecordingDeviceID string                              `json:"recordingDeviceId" api:"nullable" format:"uuid"`
	// Record device video for the whole task and persist a retrievable reference
	RecordingEnabled bool   `json:"recordingEnabled"`
	RecordingID      string `json:"recordingId" api:"nullable" format:"uuid"`
	// Where the task came from: 'api' for tasks created via POST /tasks, 'agent' for
	// tasks spawned by an agent step or created with internal=true. 'source'
	// classifies visibility only: tasks created with internal=true stay controllable
	// via this API, while tasks spawned by an agent step are readable (status,
	// trajectory, media) but not controllable.
	//
	// Any of "api", "agent".
	Source string `json:"source"`
	// Deprecated: use humanize. Both names must agree when supplied together.
	//
	// Deprecated: deprecated
	Stealth   bool   `json:"stealth"`
	Steps     int64  `json:"steps" api:"nullable"`
	StreamURL string `json:"streamUrl" api:"nullable"`
	// LLM model used by sub-agent roles: executor, app_opener, structured_output
	SubagentModel string `json:"subagentModel"`
	Succeeded     bool   `json:"succeeded" api:"nullable"`
	// Deprecated and ignored. Sampling behavior is controlled by the model provider.
	//
	// Deprecated: deprecated
	Temperature float64   `json:"temperature"`
	UpdatedAt   time.Time `json:"updatedAt" format:"date-time"`
	Vision      bool      `json:"vision"`
	// Any of "US", "BR", "FR", "DE", "IN", "JP", "KR", "ZA".
	VpnCountry string `json:"vpnCountry" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		DeviceID          respjson.Field
		DisplayID         respjson.Field
		LlmModel          respjson.Field
		OwnerID           respjson.Field
		Status            respjson.Field
		Task              respjson.Field
		TmpDevice         respjson.Field
		UserID            respjson.Field
		Accessibility     respjson.Field
		Apps              respjson.Field
		CancelRequestedAt respjson.Field
		ClaimedAt         respjson.Field
		ContinueOnFailure respjson.Field
		CreatedAt         respjson.Field
		CreatedBy         respjson.Field
		Credentials       respjson.Field
		CreditsUsed       respjson.Field
		DispatchedAt      respjson.Field
		ExecutionTimeout  respjson.Field
		Files             respjson.Field
		FinishedAt        respjson.Field
		Humanize          respjson.Field
		MaxSteps          respjson.Field
		Message           respjson.Field
		Output            respjson.Field
		OutputSchema      respjson.Field
		PendingApproval   respjson.Field
		Reasoning         respjson.Field
		RecordingDeviceID respjson.Field
		RecordingEnabled  respjson.Field
		RecordingID       respjson.Field
		Source            respjson.Field
		Stealth           respjson.Field
		Steps             respjson.Field
		StreamURL         respjson.Field
		SubagentModel     respjson.Field
		Succeeded         respjson.Field
		Temperature       respjson.Field
		UpdatedAt         respjson.Field
		Vision            respjson.Field
		VpnCountry        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListResponseItemCredential struct {
	CredentialNames []string `json:"credentialNames" api:"required"`
	PackageName     string   `json:"packageName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CredentialNames respjson.Field
		PackageName     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskListResponseItemCredential) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponseItemCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A human approval the running task is waiting for.
type TaskListResponseItemPendingApproval struct {
	// Identifies the approval
	ApprovalID string `json:"approvalId" api:"required" format:"uuid"`
	// Kind-specific, non-secret context (payment_card: label, secretGroup,
	// targetPackage, targetHost?)
	Details map[string]any `json:"details" api:"required"`
	// Without an answer by this time the task fails
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// What is being approved, e.g. payment_card; new kinds may be added
	Kind string `json:"kind" api:"required"`
	// Device the task runs on, so a client can show it next to the approval
	DeviceID string `json:"deviceId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalID  respjson.Field
		Details     respjson.Field
		ExpiresAt   respjson.Field
		Kind        respjson.Field
		DeviceID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskListResponseItemPendingApproval) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponseItemPendingApproval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetStatusResponse struct {
	// The status of the task
	//
	// Any of "prepared", "queued", "created", "running", "cancelling", "completed",
	// "failed", "cancelled".
	Status TaskGetStatusResponseStatus `json:"status" api:"required"`
	// True when the task has reached a final status and will not change again
	Terminal bool `json:"terminal" api:"required"`
	// Execution metadata for abnormal terminal outcomes
	Execution TaskGetStatusResponseExecution `json:"execution" api:"nullable"`
	// The last agent response (FastAgentResponseEvent or ManagerPlanEvent)
	LastResponse map[string]any `json:"lastResponse" api:"nullable"`
	// The agent's final answer or failure reason
	Message string `json:"message" api:"nullable"`
	// Structured output if outputSchema was set
	Output map[string]any `json:"output" api:"nullable"`
	// A human approval the running task is waiting for.
	PendingApproval TaskGetStatusResponsePendingApproval `json:"pendingApproval" api:"nullable"`
	// Device ID associated with recordingId
	RecordingDeviceID string `json:"recordingDeviceId" api:"nullable" format:"uuid"`
	// ID of the task's whole-task video recording, if recordingEnabled was set
	RecordingID string `json:"recordingId" api:"nullable" format:"uuid"`
	// Number of steps taken
	Steps int64 `json:"steps" api:"nullable"`
	// Whether the task succeeded
	Succeeded bool `json:"succeeded" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status            respjson.Field
		Terminal          respjson.Field
		Execution         respjson.Field
		LastResponse      respjson.Field
		Message           respjson.Field
		Output            respjson.Field
		PendingApproval   respjson.Field
		RecordingDeviceID respjson.Field
		RecordingID       respjson.Field
		Steps             respjson.Field
		Succeeded         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the task
type TaskGetStatusResponseStatus string

const (
	TaskGetStatusResponseStatusPrepared   TaskGetStatusResponseStatus = "prepared"
	TaskGetStatusResponseStatusQueued     TaskGetStatusResponseStatus = "queued"
	TaskGetStatusResponseStatusCreated    TaskGetStatusResponseStatus = "created"
	TaskGetStatusResponseStatusRunning    TaskGetStatusResponseStatus = "running"
	TaskGetStatusResponseStatusCancelling TaskGetStatusResponseStatus = "cancelling"
	TaskGetStatusResponseStatusCompleted  TaskGetStatusResponseStatus = "completed"
	TaskGetStatusResponseStatusFailed     TaskGetStatusResponseStatus = "failed"
	TaskGetStatusResponseStatusCancelled  TaskGetStatusResponseStatus = "cancelled"
)

// Execution metadata for abnormal terminal outcomes
type TaskGetStatusResponseExecution struct {
	Checkpoint TaskGetStatusResponseExecutionCheckpoint `json:"checkpoint" api:"required"`
	// Any of "execution_timeout", "worker_lost", "cancelled", "agent_failed".
	TerminationReason string `json:"terminationReason" api:"required"`
	// Any of "unknown".
	RetrySafety string `json:"retrySafety"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checkpoint        respjson.Field
		TerminationReason respjson.Field
		RetrySafety       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetStatusResponseExecution) RawJSON() string { return r.JSON.raw }
func (r *TaskGetStatusResponseExecution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetStatusResponseExecutionCheckpoint struct {
	CurrentSubgoal  string `json:"currentSubgoal" api:"nullable"`
	Plan            string `json:"plan" api:"nullable"`
	ProgressSummary string `json:"progressSummary" api:"nullable"`
	Steps           int64  `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentSubgoal  respjson.Field
		Plan            respjson.Field
		ProgressSummary respjson.Field
		Steps           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetStatusResponseExecutionCheckpoint) RawJSON() string { return r.JSON.raw }
func (r *TaskGetStatusResponseExecutionCheckpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A human approval the running task is waiting for.
type TaskGetStatusResponsePendingApproval struct {
	// Identifies the approval
	ApprovalID string `json:"approvalId" api:"required" format:"uuid"`
	// Kind-specific, non-secret context (payment_card: label, secretGroup,
	// targetPackage, targetHost?)
	Details map[string]any `json:"details" api:"required"`
	// Without an answer by this time the task fails
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// What is being approved, e.g. payment_card; new kinds may be added
	Kind string `json:"kind" api:"required"`
	// Device the task runs on, so a client can show it next to the approval
	DeviceID string `json:"deviceId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalID  respjson.Field
		Details     respjson.Field
		ExpiresAt   respjson.Field
		Kind        respjson.Field
		DeviceID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetStatusResponsePendingApproval) RawJSON() string { return r.JSON.raw }
func (r *TaskGetStatusResponsePendingApproval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponse struct {
	// The trajectory of the task
	Trajectory []TaskGetTrajectoryResponseTrajectoryUnion `json:"trajectory" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Trajectory  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskGetTrajectoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskGetTrajectoryResponseTrajectoryUnion contains all possible properties and
// values from [TaskGetTrajectoryResponseTrajectoryTrajectoryQueuedEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryCreatedEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryExceptionEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryCancelEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryScreenshotEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryStartEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryFinalizeEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryStopEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryResultEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryManagerInputEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorInputEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResultEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentInputEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentToolCallEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentOutputEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentEndEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentExecuteEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResultEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryRecordUiStateEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryManagerContextEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorContextEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEvent],
// [TaskGetTrajectoryResponseTrajectoryTrajectoryUnknownEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TaskGetTrajectoryResponseTrajectoryUnion struct {
	// This field is a union of
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryQueuedEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryCreatedEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExceptionEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryCancelEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryScreenshotEventData], [any],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFinalizeEventData], [any],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryResultEventData], [any],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorInputEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResultEventData], [any],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentToolCallEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentOutputEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentEndEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentExecuteEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResultEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryRecordUiStateEventData], [any],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorContextEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEventData],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEventData],
	// [map[string]any]
	Data  TaskGetTrajectoryResponseTrajectoryUnionData `json:"data"`
	Event string                                       `json:"event"`
	JSON  struct {
		Data  respjson.Field
		Event respjson.Field
		raw   string
	} `json:"-"`
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryQueuedEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryQueuedEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryCreatedEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryCreatedEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryExceptionEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryExceptionEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryCancelEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryCancelEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryScreenshotEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryScreenshotEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryStartEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryStartEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryFinalizeEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryFinalizeEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryStopEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryStopEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryResultEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryResultEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryManagerInputEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryManagerInputEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryManagerPlanEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryExecutorInputEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorInputEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryExecutorResultEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResultEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryFastAgentInputEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentInputEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryFastAgentResponseEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryFastAgentToolCallEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentToolCallEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryFastAgentOutputEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentOutputEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryFastAgentEndEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentEndEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryFastAgentExecuteEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentExecuteEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryFastAgentResultEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResultEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryToolExecutionEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryRecordUiStateEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryRecordUiStateEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryManagerContextEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryManagerContextEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryManagerResponseEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryManagerPlanDetailsEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryExecutorContextEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorContextEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryExecutorResponseEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryExecutorActionEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryExecutorActionResultEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryUserMessageEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskGetTrajectoryResponseTrajectoryUnion) AsTrajectoryUnknownEvent() (v TaskGetTrajectoryResponseTrajectoryTrajectoryUnknownEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskGetTrajectoryResponseTrajectoryUnion) RawJSON() string { return u.JSON.raw }

func (r *TaskGetTrajectoryResponseTrajectoryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskGetTrajectoryResponseTrajectoryUnionData is an implicit subunion of
// [TaskGetTrajectoryResponseTrajectoryUnion].
// TaskGetTrajectoryResponseTrajectoryUnionData provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskGetTrajectoryResponseTrajectoryUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTaskGetTrajectoryResponseTrajectoryTrajectoryUnknownEventData]
type TaskGetTrajectoryResponseTrajectoryUnionData struct {
	// This field will be present if the value is a [any] instead of an object.
	OfTaskGetTrajectoryResponseTrajectoryTrajectoryUnknownEventData any    `json:",inline"`
	ID                                                              string `json:"id"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryQueuedEventData].
	Status string `json:"status"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryCreatedEventData].
	StreamURL string `json:"streamUrl"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExceptionEventData].
	Exception string `json:"exception"`
	Reason    string `json:"reason"`
	Index     int64  `json:"index"`
	URL       string `json:"url"`
	Success   bool   `json:"success"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryResultEventData].
	Message string `json:"message"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryResultEventData].
	Steps int64 `json:"steps"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryResultEventData].
	StructuredOutput map[string]any `json:"structured_output"`
	CurrentSubgoal   string         `json:"current_subgoal"`
	Plan             string         `json:"plan"`
	Thought          string         `json:"thought"`
	Answer           string         `json:"answer"`
	// This field is a union of [map[string]any], [map[string]any], [string]
	Action TaskGetTrajectoryResponseTrajectoryUnionDataAction `json:"action"`
	Error  string                                             `json:"error"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResultEventData].
	Outcome bool   `json:"outcome"`
	Summary string `json:"summary"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventData].
	Code string `json:"code"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventData].
	ToolCallStatus string `json:"tool_call_status"`
	// This field is a union of
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventDataUsage],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEventDataUsage],
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEventDataUsage]
	Usage TaskGetTrajectoryResponseTrajectoryUnionDataUsage `json:"usage"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentToolCallEventData].
	ToolCallsRepr string `json:"tool_calls_repr"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentOutputEventData].
	Output string `json:"output"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentEndEventData].
	ToolCallCount int64  `json:"tool_call_count"`
	Instruction   string `json:"instruction"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEventData].
	ToolArgs map[string]any `json:"tool_args"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEventData].
	ToolName     string `json:"tool_name"`
	Response     string `json:"response"`
	Subgoal      string `json:"subgoal"`
	FullResponse string `json:"full_response"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEventData].
	MemoryUpdate string `json:"memory_update"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEventData].
	ProgressSummary string `json:"progress_summary"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEventData].
	ActionJson string `json:"action_json"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEventData].
	Description string `json:"description"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEventData].
	MessageIDs []string `json:"message_ids"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEventData].
	Consumer string `json:"consumer"`
	// This field is from variant
	// [TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEventData].
	StepNumber int64 `json:"step_number"`
	JSON       struct {
		OfTaskGetTrajectoryResponseTrajectoryTrajectoryUnknownEventData respjson.Field
		ID                                                              respjson.Field
		Status                                                          respjson.Field
		StreamURL                                                       respjson.Field
		Exception                                                       respjson.Field
		Reason                                                          respjson.Field
		Index                                                           respjson.Field
		URL                                                             respjson.Field
		Success                                                         respjson.Field
		Message                                                         respjson.Field
		Steps                                                           respjson.Field
		StructuredOutput                                                respjson.Field
		CurrentSubgoal                                                  respjson.Field
		Plan                                                            respjson.Field
		Thought                                                         respjson.Field
		Answer                                                          respjson.Field
		Action                                                          respjson.Field
		Error                                                           respjson.Field
		Outcome                                                         respjson.Field
		Summary                                                         respjson.Field
		Code                                                            respjson.Field
		ToolCallStatus                                                  respjson.Field
		Usage                                                           respjson.Field
		ToolCallsRepr                                                   respjson.Field
		Output                                                          respjson.Field
		ToolCallCount                                                   respjson.Field
		Instruction                                                     respjson.Field
		ToolArgs                                                        respjson.Field
		ToolName                                                        respjson.Field
		Response                                                        respjson.Field
		Subgoal                                                         respjson.Field
		FullResponse                                                    respjson.Field
		MemoryUpdate                                                    respjson.Field
		ProgressSummary                                                 respjson.Field
		ActionJson                                                      respjson.Field
		Description                                                     respjson.Field
		MessageIDs                                                      respjson.Field
		Consumer                                                        respjson.Field
		StepNumber                                                      respjson.Field
		raw                                                             string
	} `json:"-"`
}

func (r *TaskGetTrajectoryResponseTrajectoryUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskGetTrajectoryResponseTrajectoryUnionDataAction is an implicit subunion of
// [TaskGetTrajectoryResponseTrajectoryUnion].
// TaskGetTrajectoryResponseTrajectoryUnionDataAction provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskGetTrajectoryResponseTrajectoryUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfTaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEventDataAction
// OfString]
type TaskGetTrajectoryResponseTrajectoryUnionDataAction struct {
	// This field will be present if the value is a [any] instead of an object.
	OfTaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEventDataAction any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfTaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEventDataAction respjson.Field
		OfString                                                                           respjson.Field
		raw                                                                                string
	} `json:"-"`
}

func (r *TaskGetTrajectoryResponseTrajectoryUnionDataAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskGetTrajectoryResponseTrajectoryUnionDataUsage is an implicit subunion of
// [TaskGetTrajectoryResponseTrajectoryUnion].
// TaskGetTrajectoryResponseTrajectoryUnionDataUsage provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskGetTrajectoryResponseTrajectoryUnion].
type TaskGetTrajectoryResponseTrajectoryUnionDataUsage struct {
	RequestTokens  int64 `json:"request_tokens"`
	Requests       int64 `json:"requests"`
	ResponseTokens int64 `json:"response_tokens"`
	TotalTokens    int64 `json:"total_tokens"`
	JSON           struct {
		RequestTokens  respjson.Field
		Requests       respjson.Field
		ResponseTokens respjson.Field
		TotalTokens    respjson.Field
		raw            string
	} `json:"-"`
}

func (r *TaskGetTrajectoryResponseTrajectoryUnionDataUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryQueuedEvent struct {
	// Emitted to SSE clients when the task is waiting in the device queue.
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryQueuedEventData `json:"data" api:"required"`
	Event constant.QueuedEvent                                         `json:"event" default:"QueuedEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryQueuedEvent) RawJSON() string { return r.JSON.raw }
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryQueuedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emitted to SSE clients when the task is waiting in the device queue.
type TaskGetTrajectoryResponseTrajectoryTrajectoryQueuedEventData struct {
	ID     string `json:"id" api:"required"`
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryQueuedEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryQueuedEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryCreatedEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryCreatedEventData `json:"data" api:"required"`
	Event constant.CreatedEvent                                         `json:"event" default:"CreatedEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryCreatedEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryCreatedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryCreatedEventData struct {
	ID        string `json:"id" api:"required"`
	StreamURL string `json:"streamUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		StreamURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryCreatedEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryCreatedEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExceptionEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryExceptionEventData `json:"data" api:"required"`
	Event constant.ExceptionEvent                                         `json:"event" default:"ExceptionEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExceptionEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExceptionEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExceptionEventData struct {
	Exception string `json:"exception" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exception   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExceptionEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExceptionEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryCancelEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryCancelEventData `json:"data" api:"required"`
	Event constant.CancelEvent                                         `json:"event" default:"CancelEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryCancelEvent) RawJSON() string { return r.JSON.raw }
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryCancelEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryCancelEventData struct {
	Reason string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryCancelEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryCancelEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryScreenshotEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryScreenshotEventData `json:"data" api:"required"`
	Event constant.ScreenshotEvent                                         `json:"event" default:"ScreenshotEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryScreenshotEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryScreenshotEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryScreenshotEventData struct {
	Index int64  `json:"index" api:"required"`
	URL   string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryScreenshotEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryScreenshotEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryStartEvent struct {
	Data  any                 `json:"data" api:"required"`
	Event constant.StartEvent `json:"event" default:"StartEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryStartEvent) RawJSON() string { return r.JSON.raw }
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryStartEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFinalizeEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryFinalizeEventData `json:"data" api:"required"`
	Event constant.FinalizeEvent                                         `json:"event" default:"FinalizeEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFinalizeEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFinalizeEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFinalizeEventData struct {
	Reason  string `json:"reason" api:"required"`
	Success bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFinalizeEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFinalizeEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryStopEvent struct {
	Data  any                `json:"data" api:"required"`
	Event constant.StopEvent `json:"event" default:"StopEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryStopEvent) RawJSON() string { return r.JSON.raw }
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryStopEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryResultEvent struct {
	// Final result event of a task.
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryResultEventData `json:"data" api:"required"`
	Event constant.ResultEvent                                         `json:"event" default:"ResultEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryResultEvent) RawJSON() string { return r.JSON.raw }
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryResultEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Final result event of a task.
type TaskGetTrajectoryResponseTrajectoryTrajectoryResultEventData struct {
	Message          string         `json:"message" api:"nullable"`
	Steps            int64          `json:"steps" api:"nullable"`
	StructuredOutput map[string]any `json:"structured_output" api:"nullable"`
	Success          bool           `json:"success" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message          respjson.Field
		Steps            respjson.Field
		StructuredOutput respjson.Field
		Success          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryResultEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryResultEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryManagerInputEvent struct {
	Data  any                        `json:"data" api:"required"`
	Event constant.ManagerInputEvent `json:"event" default:"ManagerInputEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryManagerInputEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryManagerInputEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanEventData `json:"data" api:"required"`
	Event constant.ManagerPlanEvent                                         `json:"event" default:"ManagerPlanEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanEventData struct {
	CurrentSubgoal string `json:"current_subgoal" api:"required"`
	Plan           string `json:"plan" api:"required"`
	Thought        string `json:"thought" api:"required"`
	Answer         string `json:"answer"`
	Success        bool   `json:"success" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentSubgoal respjson.Field
		Plan           respjson.Field
		Thought        respjson.Field
		Answer         respjson.Field
		Success        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorInputEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorInputEventData `json:"data" api:"required"`
	Event constant.ExecutorInputEvent                                         `json:"event" default:"ExecutorInputEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorInputEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorInputEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorInputEventData struct {
	CurrentSubgoal string `json:"current_subgoal" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentSubgoal respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorInputEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorInputEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResultEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResultEventData `json:"data" api:"required"`
	Event constant.ExecutorResultEvent                                         `json:"event" default:"ExecutorResultEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResultEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResultEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResultEventData struct {
	Action  map[string]any `json:"action" api:"required"`
	Error   string         `json:"error" api:"required"`
	Outcome bool           `json:"outcome" api:"required"`
	Summary string         `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Error       respjson.Field
		Outcome     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResultEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResultEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentInputEvent struct {
	Data  any                          `json:"data" api:"required"`
	Event constant.FastAgentInputEvent `json:"event" default:"FastAgentInputEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentInputEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentInputEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventData `json:"data" api:"required"`
	Event constant.FastAgentResponseEvent                                         `json:"event" default:"FastAgentResponseEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventData struct {
	Thought string `json:"thought" api:"required"`
	Code    string `json:"code" api:"nullable"`
	// Classification of tool-call markup in an LLM response.
	//
	// Any of "valid", "no_markup", "malformed".
	ToolCallStatus string                                                                       `json:"tool_call_status"`
	Usage          TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventDataUsage `json:"usage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Thought        respjson.Field
		Code           respjson.Field
		ToolCallStatus respjson.Field
		Usage          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventDataUsage struct {
	RequestTokens  int64 `json:"request_tokens" api:"required"`
	Requests       int64 `json:"requests" api:"required"`
	ResponseTokens int64 `json:"response_tokens" api:"required"`
	TotalTokens    int64 `json:"total_tokens" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestTokens  respjson.Field
		Requests       respjson.Field
		ResponseTokens respjson.Field
		TotalTokens    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventDataUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResponseEventDataUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentToolCallEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentToolCallEventData `json:"data" api:"required"`
	Event constant.FastAgentToolCallEvent                                         `json:"event" default:"FastAgentToolCallEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentToolCallEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentToolCallEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentToolCallEventData struct {
	ToolCallsRepr string `json:"tool_calls_repr" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolCallsRepr respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentToolCallEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentToolCallEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentOutputEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentOutputEventData `json:"data" api:"required"`
	Event constant.FastAgentOutputEvent                                         `json:"event" default:"FastAgentOutputEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentOutputEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentOutputEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentOutputEventData struct {
	Output string `json:"output" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentOutputEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentOutputEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentEndEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentEndEventData `json:"data" api:"required"`
	Event constant.FastAgentEndEvent                                         `json:"event" default:"FastAgentEndEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentEndEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentEndEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentEndEventData struct {
	Reason        string `json:"reason" api:"required"`
	Success       bool   `json:"success" api:"required"`
	ToolCallCount int64  `json:"tool_call_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason        respjson.Field
		Success       respjson.Field
		ToolCallCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentEndEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentEndEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentExecuteEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentExecuteEventData `json:"data" api:"required"`
	Event constant.FastAgentExecuteEvent                                         `json:"event" default:"FastAgentExecuteEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentExecuteEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentExecuteEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentExecuteEventData struct {
	Instruction string `json:"instruction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instruction respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentExecuteEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentExecuteEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResultEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResultEventData `json:"data" api:"required"`
	Event constant.FastAgentResultEvent                                         `json:"event" default:"FastAgentResultEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResultEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResultEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResultEventData struct {
	Instruction string `json:"instruction" api:"required"`
	Reason      string `json:"reason" api:"required"`
	Success     bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instruction respjson.Field
		Reason      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResultEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryFastAgentResultEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEventData `json:"data" api:"required"`
	Event constant.ToolExecutionEvent                                         `json:"event" default:"ToolExecutionEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEventData struct {
	Success  bool           `json:"success" api:"required"`
	Summary  string         `json:"summary" api:"required"`
	ToolArgs map[string]any `json:"tool_args" api:"required"`
	ToolName string         `json:"tool_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		Summary     respjson.Field
		ToolArgs    respjson.Field
		ToolName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryToolExecutionEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryRecordUiStateEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryRecordUiStateEventData `json:"data" api:"required"`
	Event constant.RecordUiStateEvent                                         `json:"event" default:"RecordUIStateEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryRecordUiStateEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryRecordUiStateEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryRecordUiStateEventData struct {
	Index int64  `json:"index" api:"required"`
	URL   string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryRecordUiStateEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryRecordUiStateEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryManagerContextEvent struct {
	Data  any                          `json:"data" api:"required"`
	Event constant.ManagerContextEvent `json:"event" default:"ManagerContextEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryManagerContextEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryManagerContextEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEventData `json:"data" api:"required"`
	Event constant.ManagerResponseEvent                                         `json:"event" default:"ManagerResponseEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEventData struct {
	Response string                                                                     `json:"response" api:"required"`
	Usage    TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEventDataUsage `json:"usage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response    respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEventDataUsage struct {
	RequestTokens  int64 `json:"request_tokens" api:"required"`
	Requests       int64 `json:"requests" api:"required"`
	ResponseTokens int64 `json:"response_tokens" api:"required"`
	TotalTokens    int64 `json:"total_tokens" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestTokens  respjson.Field
		Requests       respjson.Field
		ResponseTokens respjson.Field
		TotalTokens    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEventDataUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryManagerResponseEventDataUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEventData `json:"data" api:"required"`
	Event constant.ManagerPlanDetailsEvent                                         `json:"event" default:"ManagerPlanDetailsEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEventData struct {
	Plan            string `json:"plan" api:"required"`
	Subgoal         string `json:"subgoal" api:"required"`
	Thought         string `json:"thought" api:"required"`
	Answer          string `json:"answer"`
	FullResponse    string `json:"full_response"`
	MemoryUpdate    string `json:"memory_update"`
	ProgressSummary string `json:"progress_summary"`
	Success         bool   `json:"success" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Plan            respjson.Field
		Subgoal         respjson.Field
		Thought         respjson.Field
		Answer          respjson.Field
		FullResponse    respjson.Field
		MemoryUpdate    respjson.Field
		ProgressSummary respjson.Field
		Success         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryManagerPlanDetailsEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorContextEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorContextEventData `json:"data" api:"required"`
	Event constant.ExecutorContextEvent                                         `json:"event" default:"ExecutorContextEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorContextEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorContextEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorContextEventData struct {
	Subgoal string `json:"subgoal" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Subgoal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorContextEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorContextEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEventData `json:"data" api:"required"`
	Event constant.ExecutorResponseEvent                                         `json:"event" default:"ExecutorResponseEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEventData struct {
	Response string                                                                      `json:"response" api:"required"`
	Usage    TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEventDataUsage `json:"usage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response    respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEventDataUsage struct {
	RequestTokens  int64 `json:"request_tokens" api:"required"`
	Requests       int64 `json:"requests" api:"required"`
	ResponseTokens int64 `json:"response_tokens" api:"required"`
	TotalTokens    int64 `json:"total_tokens" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestTokens  respjson.Field
		Requests       respjson.Field
		ResponseTokens respjson.Field
		TotalTokens    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEventDataUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorResponseEventDataUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEventData `json:"data" api:"required"`
	Event constant.ExecutorActionEvent                                         `json:"event" default:"ExecutorActionEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEventData struct {
	ActionJson   string `json:"action_json" api:"required"`
	Description  string `json:"description" api:"required"`
	Thought      string `json:"thought" api:"required"`
	FullResponse string `json:"full_response"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionJson   respjson.Field
		Description  respjson.Field
		Thought      respjson.Field
		FullResponse respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEvent struct {
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEventData `json:"data" api:"required"`
	Event constant.ExecutorActionResultEvent                                         `json:"event" default:"ExecutorActionResultEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEventData struct {
	Action       map[string]any `json:"action" api:"required"`
	Error        string         `json:"error" api:"required"`
	Success      bool           `json:"success" api:"required"`
	Summary      string         `json:"summary" api:"required"`
	FullResponse string         `json:"full_response"`
	Thought      string         `json:"thought"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action       respjson.Field
		Error        respjson.Field
		Success      respjson.Field
		Summary      respjson.Field
		FullResponse respjson.Field
		Thought      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryExecutorActionResultEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEvent struct {
	// Tracks the lifecycle of an external user message: queued → applied | dropped.
	Data  TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEventData `json:"data" api:"required"`
	Event constant.UserMessageEvent                                         `json:"event" default:"UserMessageEvent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tracks the lifecycle of an external user message: queued → applied | dropped.
type TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEventData struct {
	Action     string   `json:"action" api:"required"`
	MessageIDs []string `json:"message_ids" api:"required"`
	Consumer   string   `json:"consumer" api:"nullable"`
	Reason     string   `json:"reason" api:"nullable"`
	StepNumber int64    `json:"step_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		MessageIDs  respjson.Field
		Consumer    respjson.Field
		Reason      respjson.Field
		StepNumber  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEventData) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryUserMessageEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetTrajectoryResponseTrajectoryTrajectoryUnknownEvent struct {
	Event string         `json:"event" api:"required"`
	Data  map[string]any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetTrajectoryResponseTrajectoryTrajectoryUnknownEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskGetTrajectoryResponseTrajectoryTrajectoryUnknownEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskRunResponse struct {
	// The ID of the task
	ID string `json:"id" api:"required" format:"uuid"`
	// The status of the task (queued or created)
	//
	// Any of "prepared", "queued", "created", "running", "cancelling", "completed",
	// "failed", "cancelled".
	Status TaskRunResponseStatus `json:"status" api:"required"`
	// The URL of the stream (null when queued)
	StreamURL string `json:"streamUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		StreamURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskRunResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the task (queued or created)
type TaskRunResponseStatus string

const (
	TaskRunResponseStatusPrepared   TaskRunResponseStatus = "prepared"
	TaskRunResponseStatusQueued     TaskRunResponseStatus = "queued"
	TaskRunResponseStatusCreated    TaskRunResponseStatus = "created"
	TaskRunResponseStatusRunning    TaskRunResponseStatus = "running"
	TaskRunResponseStatusCancelling TaskRunResponseStatus = "cancelling"
	TaskRunResponseStatusCompleted  TaskRunResponseStatus = "completed"
	TaskRunResponseStatusFailed     TaskRunResponseStatus = "failed"
	TaskRunResponseStatusCancelled  TaskRunResponseStatus = "cancelled"
)

type TaskRunStreamedResponse = any

type TaskSendMessageResponse struct {
	// Whether the message was queued for delivery
	Sent bool `json:"sent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Sent        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskSendMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskSendMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskStopResponse struct {
	// Whether the task was cancelled
	Cancelled bool `json:"cancelled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskStopResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListParams struct {
	// Only tasks created by this user id.
	CreatedBy param.Opt[string] `query:"createdBy,omitzero" format:"uuid" json:"-"`
	// Search in task description.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Only tasks created by the calling user.
	Mine     param.Opt[bool]  `query:"mine,omitzero" json:"-"`
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Any of "id", "createdAt", "finishedAt", "status".
	OrderBy TaskListParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// Only tasks created via the API ('api') or tasks spawned by an agent step or
	// created with internal=true ('agent').
	//
	// Any of "api", "agent".
	Source TaskListParamsSource `query:"source,omitzero" json:"-"`
	// Any of "prepared", "queued", "created", "running", "cancelling", "completed",
	// "failed", "cancelled".
	Status TaskListParamsStatus `query:"status,omitzero" json:"-"`
	// Any of "asc", "desc".
	OrderByDirection TaskListParamsOrderByDirection `query:"orderByDirection,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskListParams]'s query parameters as `url.Values`.
func (r TaskListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TaskListParamsOrderBy string

const (
	TaskListParamsOrderByID         TaskListParamsOrderBy = "id"
	TaskListParamsOrderByCreatedAt  TaskListParamsOrderBy = "createdAt"
	TaskListParamsOrderByFinishedAt TaskListParamsOrderBy = "finishedAt"
	TaskListParamsOrderByStatus     TaskListParamsOrderBy = "status"
)

type TaskListParamsOrderByDirection string

const (
	TaskListParamsOrderByDirectionAsc  TaskListParamsOrderByDirection = "asc"
	TaskListParamsOrderByDirectionDesc TaskListParamsOrderByDirection = "desc"
)

// Only tasks created via the API ('api') or tasks spawned by an agent step or
// created with internal=true ('agent').
type TaskListParamsSource string

const (
	TaskListParamsSourceAPI   TaskListParamsSource = "api"
	TaskListParamsSourceAgent TaskListParamsSource = "agent"
)

type TaskListParamsStatus string

const (
	TaskListParamsStatusPrepared   TaskListParamsStatus = "prepared"
	TaskListParamsStatusQueued     TaskListParamsStatus = "queued"
	TaskListParamsStatusCreated    TaskListParamsStatus = "created"
	TaskListParamsStatusRunning    TaskListParamsStatus = "running"
	TaskListParamsStatusCancelling TaskListParamsStatus = "cancelling"
	TaskListParamsStatusCompleted  TaskListParamsStatus = "completed"
	TaskListParamsStatusFailed     TaskListParamsStatus = "failed"
	TaskListParamsStatusCancelled  TaskListParamsStatus = "cancelled"
)

type TaskRunParams struct {
	// The ID of the device to run the task on.
	DeviceID string `json:"deviceId" api:"required" format:"uuid"`
	Task     string `json:"task" api:"required"`
	// Optional custom behavioral overlay applied on top of the agent's default system
	// prompts. Never echoed back in responses or errors.
	SystemPrompt      param.Opt[string] `json:"systemPrompt,omitzero"`
	Accessibility     param.Opt[bool]   `json:"accessibility,omitzero"`
	ContinueOnFailure param.Opt[bool]   `json:"continueOnFailure,omitzero"`
	// The display ID of the device to run the task on.
	DisplayID param.Opt[int64] `json:"displayId,omitzero"`
	// Maximum agent execution time in seconds (1–2700).
	ExecutionTimeout param.Opt[int64] `json:"executionTimeout,omitzero"`
	// Request human-like input where supported by the device. Replaces the deprecated
	// stealth field; both must agree when supplied together.
	Humanize param.Opt[bool] `json:"humanize,omitzero"`
	// Mark a task started by an agent or worker on the user's behalf. It is listed as
	// source='agent' (excluded from source='api' lists) and its events carry
	// internal=true so user-facing notifications and webhooks skip it, except task
	// approval events, which always reach the owner. It is still billed, controllable
	// and retrievable by id.
	Internal param.Opt[bool] `json:"internal,omitzero"`
	// The LLM model identifier to use for the task (e.g.
	// 'mobilerun/mobile-agent-fast'). When omitted with reasoning=false, the fast-mode
	// default model is used.
	LlmModel  param.Opt[string] `json:"llmModel,omitzero"`
	MaxSteps  param.Opt[int64]  `json:"maxSteps,omitzero"`
	Reasoning param.Opt[bool]   `json:"reasoning,omitzero"`
	// Record device video for the whole task and persist a retrievable reference
	RecordingEnabled param.Opt[bool] `json:"recordingEnabled,omitzero"`
	// Deprecated: use humanize. Both names must agree when supplied together.
	Stealth param.Opt[bool] `json:"stealth,omitzero"`
	// LLM model used by sub-agent roles: executor, app_opener, structured_output
	SubagentModel param.Opt[string] `json:"subagentModel,omitzero"`
	// Deprecated and ignored. Sampling behavior is controlled by the model provider.
	Temperature  param.Opt[float64] `json:"temperature,omitzero"`
	Vision       param.Opt[bool]    `json:"vision,omitzero"`
	OutputSchema map[string]any     `json:"outputSchema,omitzero"`
	// Any of "US", "BR", "FR", "DE", "IN", "JP", "KR", "ZA".
	VpnCountry  TaskRunParamsVpnCountry   `json:"vpnCountry,omitzero"`
	Apps        []string                  `json:"apps,omitzero"`
	Credentials []TaskRunParamsCredential `json:"credentials,omitzero"`
	Files       []string                  `json:"files,omitzero"`
	paramObj
}

func (r TaskRunParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CredentialNames, PackageName are required.
type TaskRunParamsCredential struct {
	CredentialNames []string `json:"credentialNames,omitzero" api:"required"`
	PackageName     string   `json:"packageName" api:"required"`
	paramObj
}

func (r TaskRunParamsCredential) MarshalJSON() (data []byte, err error) {
	type shadow TaskRunParamsCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskRunParamsCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskRunParamsVpnCountry string

const (
	TaskRunParamsVpnCountryUs TaskRunParamsVpnCountry = "US"
	TaskRunParamsVpnCountryBr TaskRunParamsVpnCountry = "BR"
	TaskRunParamsVpnCountryFr TaskRunParamsVpnCountry = "FR"
	TaskRunParamsVpnCountryDe TaskRunParamsVpnCountry = "DE"
	TaskRunParamsVpnCountryIn TaskRunParamsVpnCountry = "IN"
	TaskRunParamsVpnCountryJp TaskRunParamsVpnCountry = "JP"
	TaskRunParamsVpnCountryKr TaskRunParamsVpnCountry = "KR"
	TaskRunParamsVpnCountryZa TaskRunParamsVpnCountry = "ZA"
)

type TaskRunStreamedParams struct {
	// The ID of the device to run the task on.
	DeviceID string `json:"deviceId" api:"required" format:"uuid"`
	Task     string `json:"task" api:"required"`
	// Optional custom behavioral overlay applied on top of the agent's default system
	// prompts. Never echoed back in responses or errors.
	SystemPrompt      param.Opt[string] `json:"systemPrompt,omitzero"`
	Accessibility     param.Opt[bool]   `json:"accessibility,omitzero"`
	ContinueOnFailure param.Opt[bool]   `json:"continueOnFailure,omitzero"`
	// The display ID of the device to run the task on.
	DisplayID param.Opt[int64] `json:"displayId,omitzero"`
	// Maximum agent execution time in seconds (1–2700).
	ExecutionTimeout param.Opt[int64] `json:"executionTimeout,omitzero"`
	// Request human-like input where supported by the device. Replaces the deprecated
	// stealth field; both must agree when supplied together.
	Humanize param.Opt[bool] `json:"humanize,omitzero"`
	// Mark a task started by an agent or worker on the user's behalf. It is listed as
	// source='agent' (excluded from source='api' lists) and its events carry
	// internal=true so user-facing notifications and webhooks skip it, except task
	// approval events, which always reach the owner. It is still billed, controllable
	// and retrievable by id.
	Internal param.Opt[bool] `json:"internal,omitzero"`
	// The LLM model identifier to use for the task (e.g.
	// 'mobilerun/mobile-agent-fast'). When omitted with reasoning=false, the fast-mode
	// default model is used.
	LlmModel  param.Opt[string] `json:"llmModel,omitzero"`
	MaxSteps  param.Opt[int64]  `json:"maxSteps,omitzero"`
	Reasoning param.Opt[bool]   `json:"reasoning,omitzero"`
	// Record device video for the whole task and persist a retrievable reference
	RecordingEnabled param.Opt[bool] `json:"recordingEnabled,omitzero"`
	// Deprecated: use humanize. Both names must agree when supplied together.
	Stealth param.Opt[bool] `json:"stealth,omitzero"`
	// LLM model used by sub-agent roles: executor, app_opener, structured_output
	SubagentModel param.Opt[string] `json:"subagentModel,omitzero"`
	// Deprecated and ignored. Sampling behavior is controlled by the model provider.
	Temperature  param.Opt[float64] `json:"temperature,omitzero"`
	Vision       param.Opt[bool]    `json:"vision,omitzero"`
	OutputSchema map[string]any     `json:"outputSchema,omitzero"`
	// Any of "US", "BR", "FR", "DE", "IN", "JP", "KR", "ZA".
	VpnCountry  TaskRunStreamedParamsVpnCountry   `json:"vpnCountry,omitzero"`
	Apps        []string                          `json:"apps,omitzero"`
	Credentials []TaskRunStreamedParamsCredential `json:"credentials,omitzero"`
	Files       []string                          `json:"files,omitzero"`
	paramObj
}

func (r TaskRunStreamedParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskRunStreamedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskRunStreamedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CredentialNames, PackageName are required.
type TaskRunStreamedParamsCredential struct {
	CredentialNames []string `json:"credentialNames,omitzero" api:"required"`
	PackageName     string   `json:"packageName" api:"required"`
	paramObj
}

func (r TaskRunStreamedParamsCredential) MarshalJSON() (data []byte, err error) {
	type shadow TaskRunStreamedParamsCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskRunStreamedParamsCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskRunStreamedParamsVpnCountry string

const (
	TaskRunStreamedParamsVpnCountryUs TaskRunStreamedParamsVpnCountry = "US"
	TaskRunStreamedParamsVpnCountryBr TaskRunStreamedParamsVpnCountry = "BR"
	TaskRunStreamedParamsVpnCountryFr TaskRunStreamedParamsVpnCountry = "FR"
	TaskRunStreamedParamsVpnCountryDe TaskRunStreamedParamsVpnCountry = "DE"
	TaskRunStreamedParamsVpnCountryIn TaskRunStreamedParamsVpnCountry = "IN"
	TaskRunStreamedParamsVpnCountryJp TaskRunStreamedParamsVpnCountry = "JP"
	TaskRunStreamedParamsVpnCountryKr TaskRunStreamedParamsVpnCountry = "KR"
	TaskRunStreamedParamsVpnCountryZa TaskRunStreamedParamsVpnCountry = "ZA"
)

type TaskSendMessageParams struct {
	// Message to send to the running agent
	Message string `json:"message" api:"required"`
	paramObj
}

func (r TaskSendMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskSendMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskSendMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
