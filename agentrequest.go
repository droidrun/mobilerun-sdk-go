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

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// AgentRequestService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentRequestService] method instead.
type AgentRequestService struct {
	options []option.RequestOption
}

// NewAgentRequestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentRequestService(opts ...option.RequestOption) (r AgentRequestService) {
	r = AgentRequestService{}
	r.options = opts
	return
}

// Owner-scoped. Returns the request with its payloads and event timeline; an
// unknown id or another owner's request is a 404.
func (r *AgentRequestService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *AgentRequestGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if requestID == "" {
		err = errors.New("missing required requestId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/requests/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Owner-scoped. Returns newest requests first, ordered by creation time and id.
// Use nextCursor as before to read older pages. Invalid cursors return 400
// invalid_cursor. Use the single read for payloads and events.
func (r *AgentRequestService) List(ctx context.Context, query AgentRequestListParams, opts ...option.RequestOption) (res *AgentRequestListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "agents/requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Owner-scoped. Returns newest requests where the Agent is source or target. Use
// nextCursor as before to read older pages. Invalid cursors return 400
// invalid_cursor.
func (r *AgentRequestService) ListForAgent(ctx context.Context, agentID string, query AgentRequestListForAgentParams, opts ...option.RequestOption) (res *AgentRequestListForAgentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/requests", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Owner-scoped. Returns the Agent-to-Agent request whose target session is
// {sessionId}, including its payloads and event timeline. Unknown sessions,
// source-only sessions, sessions without a request, and another owner's sessions
// return 404 agent_request_not_found.
func (r *AgentRequestService) GetBySession(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *AgentRequestGetBySessionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/sessions/%s/request", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AgentRequestGetResponse struct {
	ID               string                             `json:"id" api:"required" format:"uuid"`
	CreatedAt        string                             `json:"createdAt" api:"required"`
	DeadlineAt       string                             `json:"deadlineAt" api:"required"`
	Depth            int64                              `json:"depth" api:"required"`
	DispatchAttempts int64                              `json:"dispatchAttempts" api:"required"`
	DispatchedAt     string                             `json:"dispatchedAt" api:"required"`
	Error            AgentRequestGetResponseError       `json:"error" api:"required"`
	Events           []AgentRequestGetResponseEvent     `json:"events" api:"required"`
	FinishedAt       string                             `json:"finishedAt" api:"required"`
	Request          AgentRequestGetResponseRequest     `json:"request" api:"required"`
	Result           AgentRequestGetResponseResultUnion `json:"result" api:"required"`
	Room             AgentRequestGetResponseRoom        `json:"room" api:"required"`
	SourceAgentID    string                             `json:"sourceAgentId" api:"required" format:"uuid"`
	SourceSessionID  string                             `json:"sourceSessionId" api:"required" format:"uuid"`
	// Any of "queued", "running", "answered", "failed", "timed_out", "cancelled",
	// "late_answer", "needs_user_input".
	Status          AgentRequestGetResponseStatus `json:"status" api:"required"`
	TargetAgentID   string                        `json:"targetAgentId" api:"required" format:"uuid"`
	TargetSessionID string                        `json:"targetSessionId" api:"required" format:"uuid"`
	UpdatedAt       string                        `json:"updatedAt" api:"required"`
	Visited         []string                      `json:"visited" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		DeadlineAt       respjson.Field
		Depth            respjson.Field
		DispatchAttempts respjson.Field
		DispatchedAt     respjson.Field
		Error            respjson.Field
		Events           respjson.Field
		FinishedAt       respjson.Field
		Request          respjson.Field
		Result           respjson.Field
		Room             respjson.Field
		SourceAgentID    respjson.Field
		SourceSessionID  respjson.Field
		Status           respjson.Field
		TargetAgentID    respjson.Field
		TargetSessionID  respjson.Field
		UpdatedAt        respjson.Field
		Visited          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetResponseError struct {
	// Why the request did not produce an answer. An open set: new codes can appear.
	// Codes written today: `agent_archived`, `agent_disabled`, `dispatch_failed`,
	// `timed_out`, `cancelled`, `result_too_large`, `no_answer`, `target_turn_failed`,
	// `target_error`, `target_aborted`, `target_crash_recovered`,
	// `target_unknown_outcome`.
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	Preview string `json:"preview"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Preview     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetResponseEvent struct {
	CreatedAt string         `json:"createdAt" api:"required"`
	Kind      string         `json:"kind" api:"required"`
	Payload   map[string]any `json:"payload" api:"required"`
	Seq       int64          `json:"seq" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Kind        respjson.Field
		Payload     respjson.Field
		Seq         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetResponseEvent) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetResponseEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetResponseRequest struct {
	Message           string   `json:"message" api:"required"`
	SourceDisplayName string   `json:"sourceDisplayName" api:"required"`
	TimeoutSeconds    int64    `json:"timeoutSeconds" api:"required"`
	FileIDs           []string `json:"fileIds" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message           respjson.Field
		SourceDisplayName respjson.Field
		TimeoutSeconds    respjson.Field
		FileIDs           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRequestGetResponseResultUnion contains all possible properties and values
// from [AgentRequestGetResponseResultText],
// [AgentRequestGetResponseResultQuestion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AgentRequestGetResponseResultUnion struct {
	// This field is from variant [AgentRequestGetResponseResultText].
	Text string `json:"text"`
	// This field is from variant [AgentRequestGetResponseResultQuestion].
	Question string `json:"question"`
	JSON     struct {
		Text     respjson.Field
		Question respjson.Field
		raw      string
	} `json:"-"`
}

func (u AgentRequestGetResponseResultUnion) AsAgentRequestGetResponseResultText() (v AgentRequestGetResponseResultText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentRequestGetResponseResultUnion) AsAgentRequestGetResponseResultQuestion() (v AgentRequestGetResponseResultQuestion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentRequestGetResponseResultUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentRequestGetResponseResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetResponseResultText struct {
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetResponseResultText) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetResponseResultText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetResponseResultQuestion struct {
	Question string `json:"question" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Question    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetResponseResultQuestion) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetResponseResultQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetResponseRoom struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "direct", "group".
	Kind    string                              `json:"kind" api:"required"`
	Members []AgentRequestGetResponseRoomMember `json:"members" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Kind        respjson.Field
		Members     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetResponseRoom) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetResponseRoom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetResponseRoomMember struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	AgentID  string `json:"agentId" api:"required" format:"uuid"`
	JoinedAt string `json:"joinedAt" api:"required"`
	LeftAt   string `json:"leftAt" api:"required"`
	// Any of "agent", "user".
	ParticipantType string `json:"participantType" api:"required"`
	// Any of "owner", "member".
	Role   string `json:"role" api:"required"`
	UserID string `json:"userId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AgentID         respjson.Field
		JoinedAt        respjson.Field
		LeftAt          respjson.Field
		ParticipantType respjson.Field
		Role            respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetResponseRoomMember) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetResponseRoomMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetResponseStatus string

const (
	AgentRequestGetResponseStatusQueued         AgentRequestGetResponseStatus = "queued"
	AgentRequestGetResponseStatusRunning        AgentRequestGetResponseStatus = "running"
	AgentRequestGetResponseStatusAnswered       AgentRequestGetResponseStatus = "answered"
	AgentRequestGetResponseStatusFailed         AgentRequestGetResponseStatus = "failed"
	AgentRequestGetResponseStatusTimedOut       AgentRequestGetResponseStatus = "timed_out"
	AgentRequestGetResponseStatusCancelled      AgentRequestGetResponseStatus = "cancelled"
	AgentRequestGetResponseStatusLateAnswer     AgentRequestGetResponseStatus = "late_answer"
	AgentRequestGetResponseStatusNeedsUserInput AgentRequestGetResponseStatus = "needs_user_input"
)

type AgentRequestListResponse struct {
	// Cursor for the next page; null when there are no older requests.
	NextCursor string                            `json:"nextCursor" api:"required"`
	Requests   []AgentRequestListResponseRequest `json:"requests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		Requests    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestListResponseRequest struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	CreatedAt    string `json:"createdAt" api:"required"`
	DeadlineAt   string `json:"deadlineAt" api:"required"`
	Depth        int64  `json:"depth" api:"required"`
	DispatchedAt string `json:"dispatchedAt" api:"required"`
	// `error.code` of the full request, null when it has none. An open set: new codes
	// can appear. Codes written today: `agent_archived`, `agent_disabled`,
	// `dispatch_failed`, `timed_out`, `cancelled`, `result_too_large`, `no_answer`,
	// `target_turn_failed`, `target_error`, `target_aborted`,
	// `target_crash_recovered`, `target_unknown_outcome`.
	ErrorCode       string                              `json:"errorCode" api:"required"`
	FinishedAt      string                              `json:"finishedAt" api:"required"`
	Room            AgentRequestListResponseRequestRoom `json:"room" api:"required"`
	SourceAgentID   string                              `json:"sourceAgentId" api:"required" format:"uuid"`
	SourceSessionID string                              `json:"sourceSessionId" api:"required" format:"uuid"`
	// Any of "queued", "running", "answered", "failed", "timed_out", "cancelled",
	// "late_answer", "needs_user_input".
	Status          string `json:"status" api:"required"`
	TargetAgentID   string `json:"targetAgentId" api:"required" format:"uuid"`
	TargetSessionID string `json:"targetSessionId" api:"required" format:"uuid"`
	UpdatedAt       string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		DeadlineAt      respjson.Field
		Depth           respjson.Field
		DispatchedAt    respjson.Field
		ErrorCode       respjson.Field
		FinishedAt      respjson.Field
		Room            respjson.Field
		SourceAgentID   respjson.Field
		SourceSessionID respjson.Field
		Status          respjson.Field
		TargetAgentID   respjson.Field
		TargetSessionID respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestListResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestListResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestListResponseRequestRoom struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "direct", "group".
	Kind    string                                      `json:"kind" api:"required"`
	Members []AgentRequestListResponseRequestRoomMember `json:"members" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Kind        respjson.Field
		Members     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestListResponseRequestRoom) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestListResponseRequestRoom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestListResponseRequestRoomMember struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	AgentID  string `json:"agentId" api:"required" format:"uuid"`
	JoinedAt string `json:"joinedAt" api:"required"`
	LeftAt   string `json:"leftAt" api:"required"`
	// Any of "agent", "user".
	ParticipantType string `json:"participantType" api:"required"`
	// Any of "owner", "member".
	Role   string `json:"role" api:"required"`
	UserID string `json:"userId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AgentID         respjson.Field
		JoinedAt        respjson.Field
		LeftAt          respjson.Field
		ParticipantType respjson.Field
		Role            respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestListResponseRequestRoomMember) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestListResponseRequestRoomMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestListForAgentResponse struct {
	// Cursor for the next page; null when there are no older requests.
	NextCursor string                                    `json:"nextCursor" api:"required"`
	Requests   []AgentRequestListForAgentResponseRequest `json:"requests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		Requests    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestListForAgentResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestListForAgentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestListForAgentResponseRequest struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	CreatedAt    string `json:"createdAt" api:"required"`
	DeadlineAt   string `json:"deadlineAt" api:"required"`
	Depth        int64  `json:"depth" api:"required"`
	DispatchedAt string `json:"dispatchedAt" api:"required"`
	// `error.code` of the full request, null when it has none. An open set: new codes
	// can appear. Codes written today: `agent_archived`, `agent_disabled`,
	// `dispatch_failed`, `timed_out`, `cancelled`, `result_too_large`, `no_answer`,
	// `target_turn_failed`, `target_error`, `target_aborted`,
	// `target_crash_recovered`, `target_unknown_outcome`.
	ErrorCode       string                                      `json:"errorCode" api:"required"`
	FinishedAt      string                                      `json:"finishedAt" api:"required"`
	Room            AgentRequestListForAgentResponseRequestRoom `json:"room" api:"required"`
	SourceAgentID   string                                      `json:"sourceAgentId" api:"required" format:"uuid"`
	SourceSessionID string                                      `json:"sourceSessionId" api:"required" format:"uuid"`
	// Any of "queued", "running", "answered", "failed", "timed_out", "cancelled",
	// "late_answer", "needs_user_input".
	Status          string `json:"status" api:"required"`
	TargetAgentID   string `json:"targetAgentId" api:"required" format:"uuid"`
	TargetSessionID string `json:"targetSessionId" api:"required" format:"uuid"`
	UpdatedAt       string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		DeadlineAt      respjson.Field
		Depth           respjson.Field
		DispatchedAt    respjson.Field
		ErrorCode       respjson.Field
		FinishedAt      respjson.Field
		Room            respjson.Field
		SourceAgentID   respjson.Field
		SourceSessionID respjson.Field
		Status          respjson.Field
		TargetAgentID   respjson.Field
		TargetSessionID respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestListForAgentResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestListForAgentResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestListForAgentResponseRequestRoom struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "direct", "group".
	Kind    string                                              `json:"kind" api:"required"`
	Members []AgentRequestListForAgentResponseRequestRoomMember `json:"members" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Kind        respjson.Field
		Members     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestListForAgentResponseRequestRoom) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestListForAgentResponseRequestRoom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestListForAgentResponseRequestRoomMember struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	AgentID  string `json:"agentId" api:"required" format:"uuid"`
	JoinedAt string `json:"joinedAt" api:"required"`
	LeftAt   string `json:"leftAt" api:"required"`
	// Any of "agent", "user".
	ParticipantType string `json:"participantType" api:"required"`
	// Any of "owner", "member".
	Role   string `json:"role" api:"required"`
	UserID string `json:"userId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AgentID         respjson.Field
		JoinedAt        respjson.Field
		LeftAt          respjson.Field
		ParticipantType respjson.Field
		Role            respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestListForAgentResponseRequestRoomMember) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestListForAgentResponseRequestRoomMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetBySessionResponse struct {
	ID               string                                      `json:"id" api:"required" format:"uuid"`
	CreatedAt        string                                      `json:"createdAt" api:"required"`
	DeadlineAt       string                                      `json:"deadlineAt" api:"required"`
	Depth            int64                                       `json:"depth" api:"required"`
	DispatchAttempts int64                                       `json:"dispatchAttempts" api:"required"`
	DispatchedAt     string                                      `json:"dispatchedAt" api:"required"`
	Error            AgentRequestGetBySessionResponseError       `json:"error" api:"required"`
	Events           []AgentRequestGetBySessionResponseEvent     `json:"events" api:"required"`
	FinishedAt       string                                      `json:"finishedAt" api:"required"`
	Request          AgentRequestGetBySessionResponseRequest     `json:"request" api:"required"`
	Result           AgentRequestGetBySessionResponseResultUnion `json:"result" api:"required"`
	Room             AgentRequestGetBySessionResponseRoom        `json:"room" api:"required"`
	SourceAgentID    string                                      `json:"sourceAgentId" api:"required" format:"uuid"`
	SourceSessionID  string                                      `json:"sourceSessionId" api:"required" format:"uuid"`
	// Any of "queued", "running", "answered", "failed", "timed_out", "cancelled",
	// "late_answer", "needs_user_input".
	Status          AgentRequestGetBySessionResponseStatus `json:"status" api:"required"`
	TargetAgentID   string                                 `json:"targetAgentId" api:"required" format:"uuid"`
	TargetSessionID string                                 `json:"targetSessionId" api:"required" format:"uuid"`
	UpdatedAt       string                                 `json:"updatedAt" api:"required"`
	Visited         []string                               `json:"visited" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		DeadlineAt       respjson.Field
		Depth            respjson.Field
		DispatchAttempts respjson.Field
		DispatchedAt     respjson.Field
		Error            respjson.Field
		Events           respjson.Field
		FinishedAt       respjson.Field
		Request          respjson.Field
		Result           respjson.Field
		Room             respjson.Field
		SourceAgentID    respjson.Field
		SourceSessionID  respjson.Field
		Status           respjson.Field
		TargetAgentID    respjson.Field
		TargetSessionID  respjson.Field
		UpdatedAt        respjson.Field
		Visited          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetBySessionResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetBySessionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetBySessionResponseError struct {
	// Why the request did not produce an answer. An open set: new codes can appear.
	// Codes written today: `agent_archived`, `agent_disabled`, `dispatch_failed`,
	// `timed_out`, `cancelled`, `result_too_large`, `no_answer`, `target_turn_failed`,
	// `target_error`, `target_aborted`, `target_crash_recovered`,
	// `target_unknown_outcome`.
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	Preview string `json:"preview"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Preview     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetBySessionResponseError) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetBySessionResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetBySessionResponseEvent struct {
	CreatedAt string         `json:"createdAt" api:"required"`
	Kind      string         `json:"kind" api:"required"`
	Payload   map[string]any `json:"payload" api:"required"`
	Seq       int64          `json:"seq" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Kind        respjson.Field
		Payload     respjson.Field
		Seq         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetBySessionResponseEvent) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetBySessionResponseEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetBySessionResponseRequest struct {
	Message           string   `json:"message" api:"required"`
	SourceDisplayName string   `json:"sourceDisplayName" api:"required"`
	TimeoutSeconds    int64    `json:"timeoutSeconds" api:"required"`
	FileIDs           []string `json:"fileIds" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message           respjson.Field
		SourceDisplayName respjson.Field
		TimeoutSeconds    respjson.Field
		FileIDs           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetBySessionResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetBySessionResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRequestGetBySessionResponseResultUnion contains all possible properties and
// values from [AgentRequestGetBySessionResponseResultText],
// [AgentRequestGetBySessionResponseResultQuestion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AgentRequestGetBySessionResponseResultUnion struct {
	// This field is from variant [AgentRequestGetBySessionResponseResultText].
	Text string `json:"text"`
	// This field is from variant [AgentRequestGetBySessionResponseResultQuestion].
	Question string `json:"question"`
	JSON     struct {
		Text     respjson.Field
		Question respjson.Field
		raw      string
	} `json:"-"`
}

func (u AgentRequestGetBySessionResponseResultUnion) AsAgentRequestGetBySessionResponseResultText() (v AgentRequestGetBySessionResponseResultText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentRequestGetBySessionResponseResultUnion) AsAgentRequestGetBySessionResponseResultQuestion() (v AgentRequestGetBySessionResponseResultQuestion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentRequestGetBySessionResponseResultUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentRequestGetBySessionResponseResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetBySessionResponseResultText struct {
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetBySessionResponseResultText) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetBySessionResponseResultText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetBySessionResponseResultQuestion struct {
	Question string `json:"question" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Question    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetBySessionResponseResultQuestion) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetBySessionResponseResultQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetBySessionResponseRoom struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "direct", "group".
	Kind    string                                       `json:"kind" api:"required"`
	Members []AgentRequestGetBySessionResponseRoomMember `json:"members" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Kind        respjson.Field
		Members     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetBySessionResponseRoom) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetBySessionResponseRoom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetBySessionResponseRoomMember struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	AgentID  string `json:"agentId" api:"required" format:"uuid"`
	JoinedAt string `json:"joinedAt" api:"required"`
	LeftAt   string `json:"leftAt" api:"required"`
	// Any of "agent", "user".
	ParticipantType string `json:"participantType" api:"required"`
	// Any of "owner", "member".
	Role   string `json:"role" api:"required"`
	UserID string `json:"userId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AgentID         respjson.Field
		JoinedAt        respjson.Field
		LeftAt          respjson.Field
		ParticipantType respjson.Field
		Role            respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRequestGetBySessionResponseRoomMember) RawJSON() string { return r.JSON.raw }
func (r *AgentRequestGetBySessionResponseRoomMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRequestGetBySessionResponseStatus string

const (
	AgentRequestGetBySessionResponseStatusQueued         AgentRequestGetBySessionResponseStatus = "queued"
	AgentRequestGetBySessionResponseStatusRunning        AgentRequestGetBySessionResponseStatus = "running"
	AgentRequestGetBySessionResponseStatusAnswered       AgentRequestGetBySessionResponseStatus = "answered"
	AgentRequestGetBySessionResponseStatusFailed         AgentRequestGetBySessionResponseStatus = "failed"
	AgentRequestGetBySessionResponseStatusTimedOut       AgentRequestGetBySessionResponseStatus = "timed_out"
	AgentRequestGetBySessionResponseStatusCancelled      AgentRequestGetBySessionResponseStatus = "cancelled"
	AgentRequestGetBySessionResponseStatusLateAnswer     AgentRequestGetBySessionResponseStatus = "late_answer"
	AgentRequestGetBySessionResponseStatusNeedsUserInput AgentRequestGetBySessionResponseStatus = "needs_user_input"
)

type AgentRequestListParams struct {
	// Opaque nextCursor from the previous page.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of requests to return (1–100; default 50).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentRequestListParams]'s query parameters as `url.Values`.
func (r AgentRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentRequestListForAgentParams struct {
	// Opaque nextCursor from the previous page.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of requests to return (1–100; default 50).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentRequestListForAgentParams]'s query parameters as
// `url.Values`.
func (r AgentRequestListForAgentParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
