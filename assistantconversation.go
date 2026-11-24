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
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/ssestream"
)

// AssistantConversationService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssistantConversationService] method instead.
type AssistantConversationService struct {
	options []option.RequestOption
}

// NewAssistantConversationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssistantConversationService(opts ...option.RequestOption) (r AssistantConversationService) {
	r = AssistantConversationService{}
	r.options = opts
	return
}

// Creates a titled agent session. Setup may occur on the first prompt. Idempotent
// via the `Idempotency-Key` header — a duplicate submit by the same authenticated
// caller within the 24-hour idempotency window returns the already-created session
// instead of a second one. Reusing the same key with a different request body is a
// 422 `idempotency_key_reused`.
func (r *AssistantConversationService) New(ctx context.Context, body AssistantConversationNewParams, opts ...option.RequestOption) (res *AssistantConversationNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "assistant/chat/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Rename, change status, and/or pin. Title updates apply best-effort. Archiving
// always clears the pinned flag. Archiving a session with an in-flight turn is
// rejected with 409.
func (r *AssistantConversationService) Update(ctx context.Context, id string, body AssistantConversationUpdateParams, opts ...option.RequestOption) (res *AssistantConversationUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("assistant/chat/sessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Default (`kind` absent or `chat`): active named chat sessions, pinned first then
// most recent activity — `workflowId` must be absent, or the request 400s.
// `mine=true` (only valid with `kind=chat`) narrows to sessions the caller
// created. `kind=agent_workflow`: workflow-linked sessions for one workflow
// (`workflowId` required, or the request 400s), no status filter, newest episode
// first.
func (r *AssistantConversationService) List(ctx context.Context, query AssistantConversationListParams, opts ...option.RequestOption) (res *AssistantConversationListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "assistant/chat/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Abort the in-flight chat turn owned by `sessionId`. Idempotent. A turn owned by
// a different session is left untouched (204).
func (r *AssistantConversationService) Abort(ctx context.Context, body AssistantConversationAbortParams, opts ...option.RequestOption) (res *AssistantConversationAbortResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "assistant/chat/abort"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Deliver a HITL approval/rejection for an in-flight turn. Interactive HITL
// clients must start the turn with `Accept: text/event-stream`, wait for a
// `tool-input-available` event with `toolName: "hitl-approval"`, and send its
// `toolCallId` as `permissionId` here at `/assistant/chat/permission`. Do not
// submit approval as free-form user text or as `confirmed: true`. For
// `devices.reset`, only `once` and `reject` are allowed; the generic `always`
// response is rejected with HTTP 400.
func (r *AssistantConversationService) AnswerPermission(ctx context.Context, body AssistantConversationAnswerPermissionParams, opts ...option.RequestOption) (res *AssistantConversationAnswerPermissionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "assistant/chat/permission"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Deliver the user's answers to the agent's pending question for an in-flight
// turn. Idempotent via the `idempotency-key` header.
func (r *AssistantConversationService) AnswerQuestion(ctx context.Context, body AssistantConversationAnswerQuestionParams, opts ...option.RequestOption) (res *AssistantConversationAnswerQuestionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "assistant/chat/question"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Return the user's chat history for the given session. Inactive chat sessions
// return 404; workflow transcripts remain readable after completion.
func (r *AssistantConversationService) History(ctx context.Context, id string, query AssistantConversationHistoryParams, opts ...option.RequestOption) (res *AssistantConversationHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("assistant/chat/sessions/%s/messages", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Dismiss the agent's pending question. Already-resolved questions return 200
// (no-op) so multi-tab dismiss stays idempotent.
func (r *AssistantConversationService) RejectQuestion(ctx context.Context, body AssistantConversationRejectQuestionParams, opts ...option.RequestOption) (res *AssistantConversationRejectQuestionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "assistant/chat/question/reject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send a single user message. The response format follows the Accept header:
// `text/event-stream` for SSE, `application/json` for a buffered assistant reply.
// Interactive HITL requires `Accept: text/event-stream`. An approval request
// arrives as a `tool-input-available` event with `toolName: "hitl-approval"`. Its
// `toolCallId` (equal to `input.permissionID`) is the permission id; deliver the
// decision to `/assistant/chat/permission`. In the message history, the same
// request appears as a `tool-hitl-approval` part. Buffered JSON responses do not
// provide an interactive HITL continuation contract. `sessionId` targets a
// concrete active chat. The resolved chat session ID is returned as
// `chatSessionId` in the JSON body and as the `X-Chat-Session-Id` response header
// on the SSE response.
func (r *AssistantConversationService) Send(ctx context.Context, body AssistantConversationSendParams, opts ...option.RequestOption) (res *AssistantConversationSendResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "assistant/chat/message"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Reconnect to the in-flight turn stream. Replays buffered events from the start
// of the active turn, then continues live until the turn finishes. Responds 204
// when no active turn exists for the requested session. Upstream streaming
// failures return a retryable 503 with `Retry-After`. Resume is best-effort. Does
// not start an inactive session.
func (r *AssistantConversationService) StreamStreaming(ctx context.Context, query AssistantConversationStreamParams, opts ...option.RequestOption) (stream *ssestream.Stream[string]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "assistant/chat/stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &raw, opts...)
	return ssestream.NewStream[string](ssestream.NewDecoder(raw), err)
}

type AssistantConversationNewResponse struct {
	Session AssistantConversationNewResponseSession `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Session     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationNewResponseSession struct {
	ID                 string  `json:"id" api:"required"`
	Agent              string  `json:"agent" api:"required"`
	AgentID            string  `json:"agentId" api:"required" format:"uuid"`
	BasePromptTokens   float64 `json:"basePromptTokens" api:"required"`
	ContextExhaustedAt string  `json:"contextExhaustedAt" api:"required"`
	CostCents          float64 `json:"costCents" api:"required"`
	CostUsd            float64 `json:"costUsd" api:"required"`
	CreatedAt          string  `json:"createdAt" api:"required"`
	CreditsUsed        float64 `json:"creditsUsed" api:"required"`
	Description        string  `json:"description" api:"required"`
	LastActiveAt       string  `json:"lastActiveAt" api:"required"`
	PeakPromptTokens   float64 `json:"peakPromptTokens" api:"required"`
	Pinned             bool    `json:"pinned" api:"required"`
	// Any of "ready", "machine_replaced".
	PromptStatus  string `json:"promptStatus" api:"required"`
	Status        string `json:"status" api:"required"`
	Title         string `json:"title" api:"required"`
	TurnActive    bool   `json:"turnActive" api:"required"`
	TurnStartedAt string `json:"turnStartedAt" api:"required"`
	CreatedBy     string `json:"createdBy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Agent              respjson.Field
		AgentID            respjson.Field
		BasePromptTokens   respjson.Field
		ContextExhaustedAt respjson.Field
		CostCents          respjson.Field
		CostUsd            respjson.Field
		CreatedAt          respjson.Field
		CreditsUsed        respjson.Field
		Description        respjson.Field
		LastActiveAt       respjson.Field
		PeakPromptTokens   respjson.Field
		Pinned             respjson.Field
		PromptStatus       respjson.Field
		Status             respjson.Field
		Title              respjson.Field
		TurnActive         respjson.Field
		TurnStartedAt      respjson.Field
		CreatedBy          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationNewResponseSession) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationNewResponseSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationUpdateResponse struct {
	Session AssistantConversationUpdateResponseSession `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Session     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationUpdateResponseSession struct {
	ID                 string  `json:"id" api:"required"`
	Agent              string  `json:"agent" api:"required"`
	AgentID            string  `json:"agentId" api:"required" format:"uuid"`
	BasePromptTokens   float64 `json:"basePromptTokens" api:"required"`
	ContextExhaustedAt string  `json:"contextExhaustedAt" api:"required"`
	CostCents          float64 `json:"costCents" api:"required"`
	CostUsd            float64 `json:"costUsd" api:"required"`
	CreatedAt          string  `json:"createdAt" api:"required"`
	CreditsUsed        float64 `json:"creditsUsed" api:"required"`
	Description        string  `json:"description" api:"required"`
	LastActiveAt       string  `json:"lastActiveAt" api:"required"`
	PeakPromptTokens   float64 `json:"peakPromptTokens" api:"required"`
	Pinned             bool    `json:"pinned" api:"required"`
	// Any of "ready", "machine_replaced".
	PromptStatus  string `json:"promptStatus" api:"required"`
	Status        string `json:"status" api:"required"`
	Title         string `json:"title" api:"required"`
	TurnActive    bool   `json:"turnActive" api:"required"`
	TurnStartedAt string `json:"turnStartedAt" api:"required"`
	CreatedBy     string `json:"createdBy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Agent              respjson.Field
		AgentID            respjson.Field
		BasePromptTokens   respjson.Field
		ContextExhaustedAt respjson.Field
		CostCents          respjson.Field
		CostUsd            respjson.Field
		CreatedAt          respjson.Field
		CreditsUsed        respjson.Field
		Description        respjson.Field
		LastActiveAt       respjson.Field
		PeakPromptTokens   respjson.Field
		Pinned             respjson.Field
		PromptStatus       respjson.Field
		Status             respjson.Field
		Title              respjson.Field
		TurnActive         respjson.Field
		TurnStartedAt      respjson.Field
		CreatedBy          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationUpdateResponseSession) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationUpdateResponseSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationListResponse struct {
	Sessions []AssistantConversationListResponseSessionUnion `json:"sessions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Sessions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationListResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantConversationListResponseSessionUnion contains all possible properties
// and values from [AssistantConversationListResponseSessionObject],
// [AssistantConversationListResponseSessionObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AssistantConversationListResponseSessionUnion struct {
	ID                 string  `json:"id"`
	Agent              string  `json:"agent"`
	AgentID            string  `json:"agentId"`
	BasePromptTokens   float64 `json:"basePromptTokens"`
	ContextExhaustedAt string  `json:"contextExhaustedAt"`
	CostCents          float64 `json:"costCents"`
	CostUsd            float64 `json:"costUsd"`
	CreatedAt          string  `json:"createdAt"`
	CreditsUsed        float64 `json:"creditsUsed"`
	Description        string  `json:"description"`
	Episode            int64   `json:"episode"`
	Kind               string  `json:"kind"`
	LastActiveAt       string  `json:"lastActiveAt"`
	PeakPromptTokens   float64 `json:"peakPromptTokens"`
	Pinned             bool    `json:"pinned"`
	PromptStatus       string  `json:"promptStatus"`
	// This field is a union of [any], [string]
	SourceExecutionID AssistantConversationListResponseSessionUnionSourceExecutionID `json:"sourceExecutionId"`
	Status            string                                                         `json:"status"`
	Title             string                                                         `json:"title"`
	TurnActive        bool                                                           `json:"turnActive"`
	TurnStartedAt     string                                                         `json:"turnStartedAt"`
	// This field is a union of [any], [string]
	WorkflowID AssistantConversationListResponseSessionUnionWorkflowID `json:"workflowId"`
	CreatedBy  string                                                  `json:"createdBy"`
	JSON       struct {
		ID                 respjson.Field
		Agent              respjson.Field
		AgentID            respjson.Field
		BasePromptTokens   respjson.Field
		ContextExhaustedAt respjson.Field
		CostCents          respjson.Field
		CostUsd            respjson.Field
		CreatedAt          respjson.Field
		CreditsUsed        respjson.Field
		Description        respjson.Field
		Episode            respjson.Field
		Kind               respjson.Field
		LastActiveAt       respjson.Field
		PeakPromptTokens   respjson.Field
		Pinned             respjson.Field
		PromptStatus       respjson.Field
		SourceExecutionID  respjson.Field
		Status             respjson.Field
		Title              respjson.Field
		TurnActive         respjson.Field
		TurnStartedAt      respjson.Field
		WorkflowID         respjson.Field
		CreatedBy          respjson.Field
		raw                string
	} `json:"-"`
}

func (u AssistantConversationListResponseSessionUnion) AsAssistantConversationListResponseSessionObject() (v AssistantConversationListResponseSessionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantConversationListResponseSessionUnion) AsAssistantConversationListResponseSessionObject2() (v AssistantConversationListResponseSessionObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AssistantConversationListResponseSessionUnion) RawJSON() string { return u.JSON.raw }

func (r *AssistantConversationListResponseSessionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantConversationListResponseSessionUnionSourceExecutionID is an implicit
// subunion of [AssistantConversationListResponseSessionUnion].
// AssistantConversationListResponseSessionUnionSourceExecutionID provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AssistantConversationListResponseSessionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAssistantConversationListResponseSessionObjectSourceExecutionID
// OfString]
type AssistantConversationListResponseSessionUnionSourceExecutionID struct {
	// This field will be present if the value is a [any] instead of an object.
	OfAssistantConversationListResponseSessionObjectSourceExecutionID any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfAssistantConversationListResponseSessionObjectSourceExecutionID respjson.Field
		OfString                                                          respjson.Field
		raw                                                               string
	} `json:"-"`
}

func (r *AssistantConversationListResponseSessionUnionSourceExecutionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantConversationListResponseSessionUnionWorkflowID is an implicit subunion
// of [AssistantConversationListResponseSessionUnion].
// AssistantConversationListResponseSessionUnionWorkflowID provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AssistantConversationListResponseSessionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAssistantConversationListResponseSessionObjectWorkflowID
// OfString]
type AssistantConversationListResponseSessionUnionWorkflowID struct {
	// This field will be present if the value is a [any] instead of an object.
	OfAssistantConversationListResponseSessionObjectWorkflowID any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfAssistantConversationListResponseSessionObjectWorkflowID respjson.Field
		OfString                                                   respjson.Field
		raw                                                        string
	} `json:"-"`
}

func (r *AssistantConversationListResponseSessionUnionWorkflowID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationListResponseSessionObject struct {
	ID                 string  `json:"id" api:"required"`
	Agent              string  `json:"agent" api:"required"`
	AgentID            string  `json:"agentId" api:"required" format:"uuid"`
	BasePromptTokens   float64 `json:"basePromptTokens" api:"required"`
	ContextExhaustedAt string  `json:"contextExhaustedAt" api:"required"`
	CostCents          float64 `json:"costCents" api:"required"`
	CostUsd            float64 `json:"costUsd" api:"required"`
	CreatedAt          string  `json:"createdAt" api:"required"`
	CreditsUsed        float64 `json:"creditsUsed" api:"required"`
	Description        string  `json:"description" api:"required"`
	Episode            int64   `json:"episode" api:"required"`
	// Any of "chat".
	Kind             string  `json:"kind" api:"required"`
	LastActiveAt     string  `json:"lastActiveAt" api:"required"`
	PeakPromptTokens float64 `json:"peakPromptTokens" api:"required"`
	Pinned           bool    `json:"pinned" api:"required"`
	// Any of "ready", "machine_replaced".
	PromptStatus      string `json:"promptStatus" api:"required"`
	SourceExecutionID any    `json:"sourceExecutionId" api:"required"`
	Status            string `json:"status" api:"required"`
	Title             string `json:"title" api:"required"`
	TurnActive        bool   `json:"turnActive" api:"required"`
	TurnStartedAt     string `json:"turnStartedAt" api:"required"`
	WorkflowID        any    `json:"workflowId" api:"required"`
	CreatedBy         string `json:"createdBy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Agent              respjson.Field
		AgentID            respjson.Field
		BasePromptTokens   respjson.Field
		ContextExhaustedAt respjson.Field
		CostCents          respjson.Field
		CostUsd            respjson.Field
		CreatedAt          respjson.Field
		CreditsUsed        respjson.Field
		Description        respjson.Field
		Episode            respjson.Field
		Kind               respjson.Field
		LastActiveAt       respjson.Field
		PeakPromptTokens   respjson.Field
		Pinned             respjson.Field
		PromptStatus       respjson.Field
		SourceExecutionID  respjson.Field
		Status             respjson.Field
		Title              respjson.Field
		TurnActive         respjson.Field
		TurnStartedAt      respjson.Field
		WorkflowID         respjson.Field
		CreatedBy          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationListResponseSessionObject) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationListResponseSessionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationListResponseSessionObject2 struct {
	ID                 string  `json:"id" api:"required"`
	Agent              string  `json:"agent" api:"required"`
	AgentID            string  `json:"agentId" api:"required" format:"uuid"`
	BasePromptTokens   float64 `json:"basePromptTokens" api:"required"`
	ContextExhaustedAt string  `json:"contextExhaustedAt" api:"required"`
	CostCents          float64 `json:"costCents" api:"required"`
	CostUsd            float64 `json:"costUsd" api:"required"`
	CreatedAt          string  `json:"createdAt" api:"required"`
	CreditsUsed        float64 `json:"creditsUsed" api:"required"`
	Description        string  `json:"description" api:"required"`
	Episode            int64   `json:"episode" api:"required"`
	// Any of "agent_workflow".
	Kind             string  `json:"kind" api:"required"`
	LastActiveAt     string  `json:"lastActiveAt" api:"required"`
	PeakPromptTokens float64 `json:"peakPromptTokens" api:"required"`
	Pinned           bool    `json:"pinned" api:"required"`
	// Any of "ready", "machine_replaced".
	PromptStatus      string `json:"promptStatus" api:"required"`
	SourceExecutionID string `json:"sourceExecutionId" api:"required"`
	Status            string `json:"status" api:"required"`
	Title             string `json:"title" api:"required"`
	TurnActive        bool   `json:"turnActive" api:"required"`
	TurnStartedAt     string `json:"turnStartedAt" api:"required"`
	WorkflowID        string `json:"workflowId" api:"required"`
	CreatedBy         string `json:"createdBy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Agent              respjson.Field
		AgentID            respjson.Field
		BasePromptTokens   respjson.Field
		ContextExhaustedAt respjson.Field
		CostCents          respjson.Field
		CostUsd            respjson.Field
		CreatedAt          respjson.Field
		CreditsUsed        respjson.Field
		Description        respjson.Field
		Episode            respjson.Field
		Kind               respjson.Field
		LastActiveAt       respjson.Field
		PeakPromptTokens   respjson.Field
		Pinned             respjson.Field
		PromptStatus       respjson.Field
		SourceExecutionID  respjson.Field
		Status             respjson.Field
		Title              respjson.Field
		TurnActive         respjson.Field
		TurnStartedAt      respjson.Field
		WorkflowID         respjson.Field
		CreatedBy          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationListResponseSessionObject2) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationListResponseSessionObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationAbortResponse struct {
	// Any of true.
	Ok bool `json:"ok" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ok          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationAbortResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationAbortResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationAnswerPermissionResponse struct {
	// Any of true.
	Ok bool `json:"ok" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ok          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationAnswerPermissionResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationAnswerPermissionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationAnswerQuestionResponse struct {
	// Any of true.
	Ok bool `json:"ok" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ok          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationAnswerQuestionResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationAnswerQuestionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationHistoryResponse struct {
	// Session Agent identity, present on every page; null only when the session has no
	// Agent.
	Agent AssistantConversationHistoryResponseAgent  `json:"agent" api:"required"`
	Data  []AssistantConversationHistoryResponseData `json:"data" api:"required"`
	// Oldest real row in the page; pass as before to load older turns.
	FirstID string `json:"firstId" api:"required" format:"uuid"`
	// Older whole turns exist.
	HasMore bool `json:"hasMore" api:"required"`
	// Newest real row in the page, before continuation collapse or synthetic rows.
	LastID string `json:"lastId" api:"required" format:"uuid"`
	// Newest-page turn snapshot; null when before is supplied.
	Turn AssistantConversationHistoryResponseTurn `json:"turn" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent       respjson.Field
		Data        respjson.Field
		FirstID     respjson.Field
		HasMore     respjson.Field
		LastID      respjson.Field
		Turn        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Session Agent identity, present on every page; null only when the session has no
// Agent.
type AssistantConversationHistoryResponseAgent struct {
	ID      string                                           `json:"id" api:"required" format:"uuid"`
	Profile AssistantConversationHistoryResponseAgentProfile `json:"profile" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Profile     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationHistoryResponseAgent) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationHistoryResponseAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationHistoryResponseAgentProfile struct {
	// Configuration of a `blob` avatar: its seed and optional user-chosen palette
	// color. A color-only config retains the default seed fallback.
	AvatarConfig AssistantConversationHistoryResponseAgentProfileAvatarConfig `json:"avatarConfig" api:"required"`
	// Any of "blob".
	AvatarKind  string `json:"avatarKind" api:"required"`
	DisplayName string `json:"displayName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvatarConfig respjson.Field
		AvatarKind   respjson.Field
		DisplayName  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationHistoryResponseAgentProfile) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationHistoryResponseAgentProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration of a `blob` avatar: its seed and optional user-chosen palette
// color. A color-only config retains the default seed fallback.
type AssistantConversationHistoryResponseAgentProfileAvatarConfig struct {
	// Any of "slate", "red", "orange", "amber", "lime", "green", "teal", "cyan",
	// "blue", "indigo", "violet", "pink".
	Color string                                                                `json:"color"`
	Seed  AssistantConversationHistoryResponseAgentProfileAvatarConfigSeedUnion `json:"seed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		Seed        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationHistoryResponseAgentProfileAvatarConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AssistantConversationHistoryResponseAgentProfileAvatarConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantConversationHistoryResponseAgentProfileAvatarConfigSeedUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type AssistantConversationHistoryResponseAgentProfileAvatarConfigSeedUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u AssistantConversationHistoryResponseAgentProfileAvatarConfigSeedUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantConversationHistoryResponseAgentProfileAvatarConfigSeedUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AssistantConversationHistoryResponseAgentProfileAvatarConfigSeedUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AssistantConversationHistoryResponseAgentProfileAvatarConfigSeedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationHistoryResponseData struct {
	ID    string                                         `json:"id" api:"required"`
	Parts []AssistantConversationHistoryResponseDataPart `json:"parts" api:"required"`
	// Any of "user", "assistant", "system".
	Role      string                                           `json:"role" api:"required"`
	CreatedAt time.Time                                        `json:"createdAt" format:"date-time"`
	CreatedBy string                                           `json:"createdBy" api:"nullable"`
	Feedback  bool                                             `json:"feedback" api:"nullable"`
	Metadata  AssistantConversationHistoryResponseDataMetadata `json:"metadata"`
	// Any of "cloud", "telegram", "api", "workflow", "notification".
	Source    string `json:"source"`
	Synthetic bool   `json:"synthetic"`
	TurnID    string `json:"turnId" api:"nullable" format:"uuid"`
	// Deprecated: use createdBy.
	//
	// Deprecated: deprecated
	UserID string `json:"userId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Parts       respjson.Field
		Role        respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		Feedback    respjson.Field
		Metadata    respjson.Field
		Source      respjson.Field
		Synthetic   respjson.Field
		TurnID      respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationHistoryResponseData) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationHistoryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationHistoryResponseDataPart struct {
	Type        string         `json:"type" api:"required"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationHistoryResponseDataPart) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationHistoryResponseDataPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationHistoryResponseDataMetadata struct {
	Agent          string `json:"agent"`
	AgentMessageID string `json:"agentMessageId"`
	AgentSessionID string `json:"agentSessionId"`
	// Backend notice classification on synthetic turn notes; user_stop is reserved.
	//
	// Any of "user_stop", "interrupted", "no_progress", "tool_error", "error",
	// "empty_answer".
	Notice string `json:"notice"`
	// Steer user rows appear directly above the collapsed assistant group of their
	// turn, matching live cloud ordering.
	Steer               AssistantConversationHistoryResponseDataMetadataSteer `json:"steer"`
	TurnAnchorMessageID string                                                `json:"turnAnchorMessageId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent               respjson.Field
		AgentMessageID      respjson.Field
		AgentSessionID      respjson.Field
		Notice              respjson.Field
		Steer               respjson.Field
		TurnAnchorMessageID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationHistoryResponseDataMetadata) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationHistoryResponseDataMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Steer user rows appear directly above the collapsed assistant group of their
// turn, matching live cloud ordering.
type AssistantConversationHistoryResponseDataMetadataSteer struct {
	AgentMessageID string `json:"agentMessageId" api:"required"`
	// Any of "accepted", "consumed", "dropped".
	State string `json:"state" api:"required"`
	// Any of "root", "worker".
	TargetRole string    `json:"targetRole" api:"required"`
	ConsumedAt time.Time `json:"consumedAt" format:"date-time"`
	// Any of "turn_closed", "aborted".
	DropReason string `json:"dropReason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentMessageID respjson.Field
		State          respjson.Field
		TargetRole     respjson.Field
		ConsumedAt     respjson.Field
		DropReason     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationHistoryResponseDataMetadataSteer) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationHistoryResponseDataMetadataSteer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Newest-page turn snapshot; null when before is supplied.
type AssistantConversationHistoryResponseTurn struct {
	// Whether this session owns an in-flight turn.
	TurnActive bool                                              `json:"turnActive" api:"required"`
	TurnState  AssistantConversationHistoryResponseTurnTurnState `json:"turnState" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TurnActive  respjson.Field
		TurnState   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationHistoryResponseTurn) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationHistoryResponseTurn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationHistoryResponseTurnTurnState struct {
	ID             string `json:"id" api:"required" format:"uuid"`
	AbortRequested bool   `json:"abortRequested" api:"required"`
	Attachable     bool   `json:"attachable" api:"required"`
	ClaimHeld      bool   `json:"claimHeld" api:"required"`
	// Outcome of the latest turn, or null while it is still open.
	Outcome string `json:"outcome" api:"required"`
	// Any of "created", "running", "closed", "delivered", "failed".
	Phase     string `json:"phase" api:"required"`
	Reaping   bool   `json:"reaping" api:"required"`
	Revision  int64  `json:"revision" api:"required"`
	StartedAt string `json:"startedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AbortRequested respjson.Field
		Attachable     respjson.Field
		ClaimHeld      respjson.Field
		Outcome        respjson.Field
		Phase          respjson.Field
		Reaping        respjson.Field
		Revision       respjson.Field
		StartedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationHistoryResponseTurnTurnState) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationHistoryResponseTurnTurnState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationRejectQuestionResponse struct {
	// Any of true.
	Ok bool `json:"ok" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ok          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationRejectQuestionResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationRejectQuestionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationSendResponse struct {
	AssistantText string `json:"assistantText" api:"required"`
	ChatSessionID string `json:"chatSessionId" api:"required" format:"uuid"`
	ErrorText     string `json:"errorText"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantText respjson.Field
		ChatSessionID respjson.Field
		ErrorText     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantConversationSendResponse) RawJSON() string { return r.JSON.raw }
func (r *AssistantConversationSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationNewParams struct {
	Title       string            `json:"title" api:"required"`
	Agent       param.Opt[string] `json:"agent,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r AssistantConversationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AssistantConversationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantConversationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Pinned      param.Opt[bool]   `json:"pinned,omitzero"`
	Title       param.Opt[string] `json:"title,omitzero"`
	// Any of "active", "archived".
	Status AssistantConversationUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r AssistantConversationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AssistantConversationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantConversationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationUpdateParamsStatus string

const (
	AssistantConversationUpdateParamsStatusActive   AssistantConversationUpdateParamsStatus = "active"
	AssistantConversationUpdateParamsStatusArchived AssistantConversationUpdateParamsStatus = "archived"
)

type AssistantConversationListParams struct {
	WorkflowID param.Opt[string] `query:"workflowId,omitzero" json:"-"`
	// Any of "chat", "agent_workflow".
	Kind AssistantConversationListParamsKind `query:"kind,omitzero" json:"-"`
	// Any of "true", "false".
	Mine AssistantConversationListParamsMine `query:"mine,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssistantConversationListParams]'s query parameters as
// `url.Values`.
func (r AssistantConversationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssistantConversationListParamsKind string

const (
	AssistantConversationListParamsKindChat          AssistantConversationListParamsKind = "chat"
	AssistantConversationListParamsKindAgentWorkflow AssistantConversationListParamsKind = "agent_workflow"
)

type AssistantConversationListParamsMine string

const (
	AssistantConversationListParamsMineTrue  AssistantConversationListParamsMine = "true"
	AssistantConversationListParamsMineFalse AssistantConversationListParamsMine = "false"
)

type AssistantConversationAbortParams struct {
	SessionID      string            `json:"sessionId" api:"required" format:"uuid"`
	ExpectedTurnID param.Opt[string] `json:"expectedTurnId,omitzero" format:"uuid"`
	paramObj
}

func (r AssistantConversationAbortParams) MarshalJSON() (data []byte, err error) {
	type shadow AssistantConversationAbortParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantConversationAbortParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationAnswerPermissionParams struct {
	PermissionID string `json:"permissionId" api:"required"`
	// Any of "once", "always", "reject".
	Response  AssistantConversationAnswerPermissionParamsResponse `json:"response,omitzero" api:"required"`
	SessionID param.Opt[string]                                   `json:"sessionId,omitzero" format:"uuid"`
	paramObj
}

func (r AssistantConversationAnswerPermissionParams) MarshalJSON() (data []byte, err error) {
	type shadow AssistantConversationAnswerPermissionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantConversationAnswerPermissionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationAnswerPermissionParamsResponse string

const (
	AssistantConversationAnswerPermissionParamsResponseOnce   AssistantConversationAnswerPermissionParamsResponse = "once"
	AssistantConversationAnswerPermissionParamsResponseAlways AssistantConversationAnswerPermissionParamsResponse = "always"
	AssistantConversationAnswerPermissionParamsResponseReject AssistantConversationAnswerPermissionParamsResponse = "reject"
)

type AssistantConversationAnswerQuestionParams struct {
	Answers    [][]AssistantConversationAnswerQuestionParamsAnswerUnion `json:"answers,omitzero" api:"required"`
	QuestionID string                                                   `json:"questionId" api:"required"`
	SessionID  param.Opt[string]                                        `json:"sessionId,omitzero" format:"uuid"`
	paramObj
}

func (r AssistantConversationAnswerQuestionParams) MarshalJSON() (data []byte, err error) {
	type shadow AssistantConversationAnswerQuestionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantConversationAnswerQuestionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AssistantConversationAnswerQuestionParamsAnswerUnion struct {
	OfAssistantConversationAnswerQuestionsAnswerLabel  *AssistantConversationAnswerQuestionParamsAnswerLabel  `json:",omitzero,inline"`
	OfAssistantConversationAnswerQuestionsAnswerCustom *AssistantConversationAnswerQuestionParamsAnswerCustom `json:",omitzero,inline"`
	OfAssistantConversationAnswerQuestionsAnswerObject *AssistantConversationAnswerQuestionParamsAnswerObject `json:",omitzero,inline"`
	paramUnion
}

func (u AssistantConversationAnswerQuestionParamsAnswerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAssistantConversationAnswerQuestionsAnswerLabel, u.OfAssistantConversationAnswerQuestionsAnswerCustom, u.OfAssistantConversationAnswerQuestionsAnswerObject)
}
func (u *AssistantConversationAnswerQuestionParamsAnswerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Label is required.
type AssistantConversationAnswerQuestionParamsAnswerLabel struct {
	Label string `json:"label" api:"required"`
	paramObj
}

func (r AssistantConversationAnswerQuestionParamsAnswerLabel) MarshalJSON() (data []byte, err error) {
	type shadow AssistantConversationAnswerQuestionParamsAnswerLabel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantConversationAnswerQuestionParamsAnswerLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Custom is required.
type AssistantConversationAnswerQuestionParamsAnswerCustom struct {
	Custom string `json:"custom" api:"required"`
	paramObj
}

func (r AssistantConversationAnswerQuestionParamsAnswerCustom) MarshalJSON() (data []byte, err error) {
	type shadow AssistantConversationAnswerQuestionParamsAnswerCustom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantConversationAnswerQuestionParamsAnswerCustom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Custom, Label are required.
type AssistantConversationAnswerQuestionParamsAnswerObject struct {
	Custom string `json:"custom" api:"required"`
	Label  string `json:"label" api:"required"`
	paramObj
}

func (r AssistantConversationAnswerQuestionParamsAnswerObject) MarshalJSON() (data []byte, err error) {
	type shadow AssistantConversationAnswerQuestionParamsAnswerObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantConversationAnswerQuestionParamsAnswerObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationHistoryParams struct {
	// Any real row in this session; its containing turn is excluded.
	Before param.Opt[string] `query:"before,omitzero" format:"uuid" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssistantConversationHistoryParams]'s query parameters as
// `url.Values`.
func (r AssistantConversationHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssistantConversationRejectQuestionParams struct {
	QuestionID string            `json:"questionId" api:"required"`
	SessionID  param.Opt[string] `json:"sessionId,omitzero" format:"uuid"`
	paramObj
}

func (r AssistantConversationRejectQuestionParams) MarshalJSON() (data []byte, err error) {
	type shadow AssistantConversationRejectQuestionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantConversationRejectQuestionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationSendParams struct {
	Message   string            `json:"message" api:"required"`
	SessionID string            `json:"sessionId" api:"required" format:"uuid"`
	Agent     param.Opt[string] `json:"agent,omitzero"`
	Model     param.Opt[string] `json:"model,omitzero"`
	paramObj
}

func (r AssistantConversationSendParams) MarshalJSON() (data []byte, err error) {
	type shadow AssistantConversationSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantConversationSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantConversationStreamParams struct {
	SessionID string `query:"sessionId" api:"required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AssistantConversationStreamParams]'s query parameters as
// `url.Values`.
func (r AssistantConversationStreamParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
