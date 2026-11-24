// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// AgentService contains methods and other services that help with interacting with
// the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	options       []option.RequestOption
	Profile       AgentProfileService
	Files         AgentFileService
	Conversations AgentConversationService
	Requests      AgentRequestService
	Routines      AgentRoutineService
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.options = opts
	r.Profile = NewAgentProfileService(opts...)
	r.Files = NewAgentFileService(opts...)
	r.Conversations = NewAgentConversationService(opts...)
	r.Requests = NewAgentRequestService(opts...)
	r.Routines = NewAgentRoutineService(opts...)
	return
}

// Creates an Agent directly on the owner runtime, without a chat message or
// approval card. Provisioning runs in the background; the new Agent has no chat
// history. Requires Idempotency-Key; a replay returns the current Agent with the
// original 201 status.
func (r *AgentService) New(ctx context.Context, body AgentNewParams, opts ...option.RequestOption) (res *AgentNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Owner-scoped. Returns every Agent of the owner that is not on a retired machine,
// archived and failed ones included, oldest first.
func (r *AgentService) List(ctx context.Context, opts ...option.RequestOption) (res *AgentListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the environment model profiles and their default. Model routing is managed
// by the server.
func (r *AgentService) ModelProfiles(ctx context.Context, opts ...option.RequestOption) (res *AgentModelProfilesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "agents/model-profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AgentNewResponse struct {
	// An Agent of the owner: its lifecycle state, profile and runtime.
	Agent AgentNewResponseAgent `json:"agent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Agent of the owner: its lifecycle state, profile and runtime.
type AgentNewResponseAgent struct {
	ID         string `json:"id" api:"required" format:"uuid"`
	ArchivedAt string `json:"archivedAt" api:"required"`
	CreatedAt  string `json:"createdAt" api:"required"`
	// Why a `disabled` Agent is disabled: `owner` chose it, `over-limit` the plan
	// grants fewer Agent slots than the owner had enabled. Null while the Agent is not
	// disabled.
	//
	// Any of "owner", "over-limit".
	DisabledReason string `json:"disabledReason" api:"required"`
	// Why provisioning failed, null otherwise. An open set: new codes can appear.
	// Codes written today: `provision_timeout`, `agent_limit_reached`,
	// `provision_failed`; a code the harness reports for a failed start is passed
	// through.
	ErrorCode string `json:"errorCode" api:"required"`
	// Whether this is the owner's primary Agent, the one a chat session uses when it
	// names no `agentId`.
	Primary bool `json:"primary" api:"required"`
	// Persona of the Agent plus its individual instructions passage; `instructions` is
	// null when none is set.
	Profile AgentNewResponseAgentProfile `json:"profile" api:"required"`
	// When provisioning of the Agent completed (ISO 8601); null while a newly created
	// Agent is still pending and after its provisioning failed.
	ProvisionedAt string `json:"provisionedAt" api:"required"`
	// `pending`: created but not provisioned yet, it takes no turns. `active`: ready
	// for turns. `failed`: provisioning gave up (see `errorCode`); the Agent is
	// archived. `archived`: read-only. `disabled`: switched off by the owner or
	// because the plan grants fewer Agent slots; it takes no turns and holds no slot,
	// and the owner can enable it again while a slot is free.
	//
	// Any of "active", "pending", "failed", "archived", "disabled".
	Status string `json:"status" api:"required"`
	// Newest thread excerpt (90 characters), awaits-input and timestamp, from at most
	// 20 newest rows. Null when the thread has no rows.
	ThreadPreview AgentNewResponseAgentThreadPreview `json:"threadPreview" api:"required"`
	// Active chat on the active machine, pinned first, then session lastActiveAt DESC,
	// id DESC. Null for pending/archived Agents or no thread.
	ThreadSessionID string `json:"threadSessionId" api:"required" format:"uuid"`
	// Whether this active primary Agent is still due its one-time greeting: the
	// welcome-back for an adopted primary, or the onboarding greeting for a primary
	// whose chat never had a message.
	WelcomePending bool `json:"welcomePending" api:"required"`
	// Latest activity of the Agent (newest chat session activity or live turn
	// heartbeat), ISO 8601; null when it has never been active.
	LastActiveAt string `json:"lastActiveAt" api:"nullable"`
	// Whether the Agent is running right now, derived from stored state only.
	// `waking`: an active start, its machine is provisioning or a restart/rehydrate is
	// pending and changed within the last 10 minutes (a timed-out or stalled start
	// reads `sleeping`). `awake`: the machine is running and this Agent has been
	// started in the current run. `sleeping`: everything else (stopped or parked
	// machine, an Agent the running harness has not started yet, or an Agent that is
	// not active). `awake` can lag a machine stop until the reconciler converges.
	//
	// Any of "awake", "sleeping", "waking".
	Runtime string `json:"runtime" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ArchivedAt      respjson.Field
		CreatedAt       respjson.Field
		DisabledReason  respjson.Field
		ErrorCode       respjson.Field
		Primary         respjson.Field
		Profile         respjson.Field
		ProvisionedAt   respjson.Field
		Status          respjson.Field
		ThreadPreview   respjson.Field
		ThreadSessionID respjson.Field
		WelcomePending  respjson.Field
		LastActiveAt    respjson.Field
		Runtime         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentNewResponseAgent) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponseAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Persona of the Agent plus its individual instructions passage; `instructions` is
// null when none is set.
type AgentNewResponseAgentProfile struct {
	// Configuration of a `blob` avatar: its seed and optional user-chosen palette
	// color. A color-only config retains the default seed fallback.
	AvatarConfig AgentNewResponseAgentProfileAvatarConfig `json:"avatarConfig" api:"required"`
	// Any of "blob".
	AvatarKind   string `json:"avatarKind" api:"required"`
	DisplayName  string `json:"displayName" api:"required"`
	Instructions string `json:"instructions" api:"required"`
	ModelProfile string `json:"modelProfile" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvatarConfig respjson.Field
		AvatarKind   respjson.Field
		DisplayName  respjson.Field
		Instructions respjson.Field
		ModelProfile respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentNewResponseAgentProfile) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponseAgentProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration of a `blob` avatar: its seed and optional user-chosen palette
// color. A color-only config retains the default seed fallback.
type AgentNewResponseAgentProfileAvatarConfig struct {
	// Any of "slate", "red", "orange", "amber", "lime", "green", "teal", "cyan",
	// "blue", "indigo", "violet", "pink".
	Color string                                            `json:"color"`
	Seed  AgentNewResponseAgentProfileAvatarConfigSeedUnion `json:"seed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		Seed        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentNewResponseAgentProfileAvatarConfig) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponseAgentProfileAvatarConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentNewResponseAgentProfileAvatarConfigSeedUnion contains all possible
// properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type AgentNewResponseAgentProfileAvatarConfigSeedUnion struct {
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

func (u AgentNewResponseAgentProfileAvatarConfigSeedUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentNewResponseAgentProfileAvatarConfigSeedUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentNewResponseAgentProfileAvatarConfigSeedUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentNewResponseAgentProfileAvatarConfigSeedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Newest thread excerpt (90 characters), awaits-input and timestamp, from at most
// 20 newest rows. Null when the thread has no rows.
type AgentNewResponseAgentThreadPreview struct {
	AwaitsInput bool      `json:"awaitsInput" api:"required"`
	Excerpt     string    `json:"excerpt" api:"required"`
	Timestamp   time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwaitsInput respjson.Field
		Excerpt     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentNewResponseAgentThreadPreview) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponseAgentThreadPreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListResponse struct {
	Agents []AgentListResponseAgent `json:"agents" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agents      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Agent of the owner: its lifecycle state, profile and runtime.
type AgentListResponseAgent struct {
	ID         string `json:"id" api:"required" format:"uuid"`
	ArchivedAt string `json:"archivedAt" api:"required"`
	CreatedAt  string `json:"createdAt" api:"required"`
	// Why a `disabled` Agent is disabled: `owner` chose it, `over-limit` the plan
	// grants fewer Agent slots than the owner had enabled. Null while the Agent is not
	// disabled.
	//
	// Any of "owner", "over-limit".
	DisabledReason string `json:"disabledReason" api:"required"`
	// Why provisioning failed, null otherwise. An open set: new codes can appear.
	// Codes written today: `provision_timeout`, `agent_limit_reached`,
	// `provision_failed`; a code the harness reports for a failed start is passed
	// through.
	ErrorCode string `json:"errorCode" api:"required"`
	// Whether this is the owner's primary Agent, the one a chat session uses when it
	// names no `agentId`.
	Primary bool `json:"primary" api:"required"`
	// Persona of the Agent plus its individual instructions passage; `instructions` is
	// null when none is set.
	Profile AgentListResponseAgentProfile `json:"profile" api:"required"`
	// When provisioning of the Agent completed (ISO 8601); null while a newly created
	// Agent is still pending and after its provisioning failed.
	ProvisionedAt string `json:"provisionedAt" api:"required"`
	// `pending`: created but not provisioned yet, it takes no turns. `active`: ready
	// for turns. `failed`: provisioning gave up (see `errorCode`); the Agent is
	// archived. `archived`: read-only. `disabled`: switched off by the owner or
	// because the plan grants fewer Agent slots; it takes no turns and holds no slot,
	// and the owner can enable it again while a slot is free.
	//
	// Any of "active", "pending", "failed", "archived", "disabled".
	Status string `json:"status" api:"required"`
	// Newest thread excerpt (90 characters), awaits-input and timestamp, from at most
	// 20 newest rows. Null when the thread has no rows.
	ThreadPreview AgentListResponseAgentThreadPreview `json:"threadPreview" api:"required"`
	// Active chat on the active machine, pinned first, then session lastActiveAt DESC,
	// id DESC. Null for pending/archived Agents or no thread.
	ThreadSessionID string `json:"threadSessionId" api:"required" format:"uuid"`
	// Whether this active primary Agent is still due its one-time greeting: the
	// welcome-back for an adopted primary, or the onboarding greeting for a primary
	// whose chat never had a message.
	WelcomePending bool `json:"welcomePending" api:"required"`
	// Latest activity of the Agent (newest chat session activity or live turn
	// heartbeat), ISO 8601; null when it has never been active.
	LastActiveAt string `json:"lastActiveAt" api:"nullable"`
	// Whether the Agent is running right now, derived from stored state only.
	// `waking`: an active start, its machine is provisioning or a restart/rehydrate is
	// pending and changed within the last 10 minutes (a timed-out or stalled start
	// reads `sleeping`). `awake`: the machine is running and this Agent has been
	// started in the current run. `sleeping`: everything else (stopped or parked
	// machine, an Agent the running harness has not started yet, or an Agent that is
	// not active). `awake` can lag a machine stop until the reconciler converges.
	//
	// Any of "awake", "sleeping", "waking".
	Runtime string `json:"runtime" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ArchivedAt      respjson.Field
		CreatedAt       respjson.Field
		DisabledReason  respjson.Field
		ErrorCode       respjson.Field
		Primary         respjson.Field
		Profile         respjson.Field
		ProvisionedAt   respjson.Field
		Status          respjson.Field
		ThreadPreview   respjson.Field
		ThreadSessionID respjson.Field
		WelcomePending  respjson.Field
		LastActiveAt    respjson.Field
		Runtime         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentListResponseAgent) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponseAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Persona of the Agent plus its individual instructions passage; `instructions` is
// null when none is set.
type AgentListResponseAgentProfile struct {
	// Configuration of a `blob` avatar: its seed and optional user-chosen palette
	// color. A color-only config retains the default seed fallback.
	AvatarConfig AgentListResponseAgentProfileAvatarConfig `json:"avatarConfig" api:"required"`
	// Any of "blob".
	AvatarKind   string `json:"avatarKind" api:"required"`
	DisplayName  string `json:"displayName" api:"required"`
	Instructions string `json:"instructions" api:"required"`
	ModelProfile string `json:"modelProfile" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvatarConfig respjson.Field
		AvatarKind   respjson.Field
		DisplayName  respjson.Field
		Instructions respjson.Field
		ModelProfile respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentListResponseAgentProfile) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponseAgentProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration of a `blob` avatar: its seed and optional user-chosen palette
// color. A color-only config retains the default seed fallback.
type AgentListResponseAgentProfileAvatarConfig struct {
	// Any of "slate", "red", "orange", "amber", "lime", "green", "teal", "cyan",
	// "blue", "indigo", "violet", "pink".
	Color string                                             `json:"color"`
	Seed  AgentListResponseAgentProfileAvatarConfigSeedUnion `json:"seed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		Seed        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentListResponseAgentProfileAvatarConfig) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponseAgentProfileAvatarConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentListResponseAgentProfileAvatarConfigSeedUnion contains all possible
// properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type AgentListResponseAgentProfileAvatarConfigSeedUnion struct {
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

func (u AgentListResponseAgentProfileAvatarConfigSeedUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentListResponseAgentProfileAvatarConfigSeedUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentListResponseAgentProfileAvatarConfigSeedUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentListResponseAgentProfileAvatarConfigSeedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Newest thread excerpt (90 characters), awaits-input and timestamp, from at most
// 20 newest rows. Null when the thread has no rows.
type AgentListResponseAgentThreadPreview struct {
	AwaitsInput bool      `json:"awaitsInput" api:"required"`
	Excerpt     string    `json:"excerpt" api:"required"`
	Timestamp   time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwaitsInput respjson.Field
		Excerpt     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentListResponseAgentThreadPreview) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponseAgentThreadPreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentModelProfilesResponse struct {
	DefaultProfileID string                              `json:"defaultProfileId" api:"required"`
	Profiles         []AgentModelProfilesResponseProfile `json:"profiles" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultProfileID respjson.Field
		Profiles         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentModelProfilesResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentModelProfilesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentModelProfilesResponseProfile struct {
	ID          string `json:"id" api:"required"`
	CostHint    string `json:"costHint" api:"required"`
	Description string `json:"description" api:"required"`
	Label       string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CostHint    respjson.Field
		Description respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentModelProfilesResponseProfile) RawJSON() string { return r.JSON.raw }
func (r *AgentModelProfilesResponseProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentNewParams struct {
	// Single-line name shown for the Agent; surrounding whitespace is trimmed and it
	// needs at least one letter or number.
	DisplayName string `json:"displayName" api:"required"`
	// The Agent's individual instructions passage; an empty or whitespace-only text
	// clears it. Line breaks count three characters toward the rendered size limit.
	Instructions param.Opt[string]             `json:"instructions,omitzero"`
	AvatarSeed   AgentNewParamsAvatarSeedUnion `json:"avatarSeed,omitzero"`
	// Any of "slate", "red", "orange", "amber", "lime", "green", "teal", "cyan",
	// "blue", "indigo", "violet", "pink".
	Color AgentNewParamsColor `json:"color,omitzero"`
	paramObj
}

func (r AgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentNewParamsAvatarSeedUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u AgentNewParamsAvatarSeedUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *AgentNewParamsAvatarSeedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type AgentNewParamsColor string

const (
	AgentNewParamsColorSlate  AgentNewParamsColor = "slate"
	AgentNewParamsColorRed    AgentNewParamsColor = "red"
	AgentNewParamsColorOrange AgentNewParamsColor = "orange"
	AgentNewParamsColorAmber  AgentNewParamsColor = "amber"
	AgentNewParamsColorLime   AgentNewParamsColor = "lime"
	AgentNewParamsColorGreen  AgentNewParamsColor = "green"
	AgentNewParamsColorTeal   AgentNewParamsColor = "teal"
	AgentNewParamsColorCyan   AgentNewParamsColor = "cyan"
	AgentNewParamsColorBlue   AgentNewParamsColor = "blue"
	AgentNewParamsColorIndigo AgentNewParamsColor = "indigo"
	AgentNewParamsColorViolet AgentNewParamsColor = "violet"
	AgentNewParamsColorPink   AgentNewParamsColor = "pink"
)
