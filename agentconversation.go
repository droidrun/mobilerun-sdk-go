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

// AgentConversationService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentConversationService] method instead.
type AgentConversationService struct {
	options []option.RequestOption
}

// NewAgentConversationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgentConversationService(opts ...option.RequestOption) (r AgentConversationService) {
	r = AgentConversationService{}
	r.options = opts
	return
}

// Create a named chat session bound to one agent. Idempotent via the
// `Idempotency-Key` header, like the unscoped create.
func (r *AgentConversationService) New(ctx context.Context, agentID string, body AgentConversationNewParams, opts ...option.RequestOption) (res *AgentConversationNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/chat/sessions", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List the user's named chat sessions for one agent. Supports the same `kind`,
// `workflowId` and `mine` filters as the unscoped list.
func (r *AgentConversationService) List(ctx context.Context, agentID string, query AgentConversationListParams, opts ...option.RequestOption) (res *AgentConversationListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/chat/sessions", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AgentConversationNewResponse struct {
	Session AgentConversationNewResponseSession `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Session     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConversationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentConversationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConversationNewResponseSession struct {
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
func (r AgentConversationNewResponseSession) RawJSON() string { return r.JSON.raw }
func (r *AgentConversationNewResponseSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConversationListResponse struct {
	Sessions []AgentConversationListResponseSessionUnion `json:"sessions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Sessions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConversationListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentConversationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentConversationListResponseSessionUnion contains all possible properties and
// values from [AgentConversationListResponseSessionObject],
// [AgentConversationListResponseSessionObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AgentConversationListResponseSessionUnion struct {
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
	SourceExecutionID AgentConversationListResponseSessionUnionSourceExecutionID `json:"sourceExecutionId"`
	Status            string                                                     `json:"status"`
	Title             string                                                     `json:"title"`
	TurnActive        bool                                                       `json:"turnActive"`
	TurnStartedAt     string                                                     `json:"turnStartedAt"`
	// This field is a union of [any], [string]
	WorkflowID AgentConversationListResponseSessionUnionWorkflowID `json:"workflowId"`
	CreatedBy  string                                              `json:"createdBy"`
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

func (u AgentConversationListResponseSessionUnion) AsAgentConversationListResponseSessionObject() (v AgentConversationListResponseSessionObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentConversationListResponseSessionUnion) AsAgentConversationListResponseSessionObject2() (v AgentConversationListResponseSessionObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentConversationListResponseSessionUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentConversationListResponseSessionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentConversationListResponseSessionUnionSourceExecutionID is an implicit
// subunion of [AgentConversationListResponseSessionUnion].
// AgentConversationListResponseSessionUnionSourceExecutionID provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AgentConversationListResponseSessionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAgentConversationListResponseSessionObjectSourceExecutionID
// OfString]
type AgentConversationListResponseSessionUnionSourceExecutionID struct {
	// This field will be present if the value is a [any] instead of an object.
	OfAgentConversationListResponseSessionObjectSourceExecutionID any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfAgentConversationListResponseSessionObjectSourceExecutionID respjson.Field
		OfString                                                      respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (r *AgentConversationListResponseSessionUnionSourceExecutionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentConversationListResponseSessionUnionWorkflowID is an implicit subunion of
// [AgentConversationListResponseSessionUnion].
// AgentConversationListResponseSessionUnionWorkflowID provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AgentConversationListResponseSessionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAgentConversationListResponseSessionObjectWorkflowID OfString]
type AgentConversationListResponseSessionUnionWorkflowID struct {
	// This field will be present if the value is a [any] instead of an object.
	OfAgentConversationListResponseSessionObjectWorkflowID any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfAgentConversationListResponseSessionObjectWorkflowID respjson.Field
		OfString                                               respjson.Field
		raw                                                    string
	} `json:"-"`
}

func (r *AgentConversationListResponseSessionUnionWorkflowID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConversationListResponseSessionObject struct {
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
func (r AgentConversationListResponseSessionObject) RawJSON() string { return r.JSON.raw }
func (r *AgentConversationListResponseSessionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConversationListResponseSessionObject2 struct {
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
func (r AgentConversationListResponseSessionObject2) RawJSON() string { return r.JSON.raw }
func (r *AgentConversationListResponseSessionObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConversationNewParams struct {
	Title       string            `json:"title" api:"required"`
	Agent       param.Opt[string] `json:"agent,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r AgentConversationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentConversationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentConversationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConversationListParams struct {
	WorkflowID param.Opt[string] `query:"workflowId,omitzero" json:"-"`
	// Any of "chat", "agent_workflow".
	Kind AgentConversationListParamsKind `query:"kind,omitzero" json:"-"`
	// Any of "true", "false".
	Mine AgentConversationListParamsMine `query:"mine,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentConversationListParams]'s query parameters as
// `url.Values`.
func (r AgentConversationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentConversationListParamsKind string

const (
	AgentConversationListParamsKindChat          AgentConversationListParamsKind = "chat"
	AgentConversationListParamsKindAgentWorkflow AgentConversationListParamsKind = "agent_workflow"
)

type AgentConversationListParamsMine string

const (
	AgentConversationListParamsMineTrue  AgentConversationListParamsMine = "true"
	AgentConversationListParamsMineFalse AgentConversationListParamsMine = "false"
)
