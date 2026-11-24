// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"encoding/json"
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

// AgentProfileService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentProfileService] method instead.
type AgentProfileService struct {
	options []option.RequestOption
}

// NewAgentProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentProfileService(opts ...option.RequestOption) (r AgentProfileService) {
	r = AgentProfileService{}
	r.options = opts
	return
}

// Owner-scoped partial update of `displayName`, `instructions`, `modelProfile` and
// the avatar. A null modelProfile uses the environment default; an unknown id
// returns unknown_model_profile with details.validProfiles. Returns the updated
// Agent; an Agent of another owner or an unknown id is a 404.
func (r *AgentProfileService) Update(ctx context.Context, agentID string, body AgentProfileUpdateParams, opts ...option.RequestOption) (res *AgentProfileUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/profile", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// An Agent of the owner: its lifecycle state, profile and runtime.
type AgentProfileUpdateResponse struct {
	ID         string `json:"id" api:"required" format:"uuid"`
	ArchivedAt string `json:"archivedAt" api:"required"`
	CreatedAt  string `json:"createdAt" api:"required"`
	// Why a `disabled` Agent is disabled: `owner` chose it, `over-limit` the plan
	// grants fewer Agent slots than the owner had enabled. Null while the Agent is not
	// disabled.
	//
	// Any of "owner", "over-limit".
	DisabledReason AgentProfileUpdateResponseDisabledReason `json:"disabledReason" api:"required"`
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
	Profile AgentProfileUpdateResponseProfile `json:"profile" api:"required"`
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
	Status AgentProfileUpdateResponseStatus `json:"status" api:"required"`
	// Newest thread excerpt (90 characters), awaits-input and timestamp, from at most
	// 20 newest rows. Null when the thread has no rows.
	ThreadPreview AgentProfileUpdateResponseThreadPreview `json:"threadPreview" api:"required"`
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
	Runtime AgentProfileUpdateResponseRuntime `json:"runtime" api:"nullable"`
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
func (r AgentProfileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentProfileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why a `disabled` Agent is disabled: `owner` chose it, `over-limit` the plan
// grants fewer Agent slots than the owner had enabled. Null while the Agent is not
// disabled.
type AgentProfileUpdateResponseDisabledReason string

const (
	AgentProfileUpdateResponseDisabledReasonOwner     AgentProfileUpdateResponseDisabledReason = "owner"
	AgentProfileUpdateResponseDisabledReasonOverLimit AgentProfileUpdateResponseDisabledReason = "over-limit"
)

// Persona of the Agent plus its individual instructions passage; `instructions` is
// null when none is set.
type AgentProfileUpdateResponseProfile struct {
	// Configuration of a `blob` avatar: its seed and optional user-chosen palette
	// color. A color-only config retains the default seed fallback.
	AvatarConfig AgentProfileUpdateResponseProfileAvatarConfig `json:"avatarConfig" api:"required"`
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
func (r AgentProfileUpdateResponseProfile) RawJSON() string { return r.JSON.raw }
func (r *AgentProfileUpdateResponseProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration of a `blob` avatar: its seed and optional user-chosen palette
// color. A color-only config retains the default seed fallback.
type AgentProfileUpdateResponseProfileAvatarConfig struct {
	// Any of "slate", "red", "orange", "amber", "lime", "green", "teal", "cyan",
	// "blue", "indigo", "violet", "pink".
	Color string                                                 `json:"color"`
	Seed  AgentProfileUpdateResponseProfileAvatarConfigSeedUnion `json:"seed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		Seed        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentProfileUpdateResponseProfileAvatarConfig) RawJSON() string { return r.JSON.raw }
func (r *AgentProfileUpdateResponseProfileAvatarConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentProfileUpdateResponseProfileAvatarConfigSeedUnion contains all possible
// properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type AgentProfileUpdateResponseProfileAvatarConfigSeedUnion struct {
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

func (u AgentProfileUpdateResponseProfileAvatarConfigSeedUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentProfileUpdateResponseProfileAvatarConfigSeedUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentProfileUpdateResponseProfileAvatarConfigSeedUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentProfileUpdateResponseProfileAvatarConfigSeedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `pending`: created but not provisioned yet, it takes no turns. `active`: ready
// for turns. `failed`: provisioning gave up (see `errorCode`); the Agent is
// archived. `archived`: read-only. `disabled`: switched off by the owner or
// because the plan grants fewer Agent slots; it takes no turns and holds no slot,
// and the owner can enable it again while a slot is free.
type AgentProfileUpdateResponseStatus string

const (
	AgentProfileUpdateResponseStatusActive   AgentProfileUpdateResponseStatus = "active"
	AgentProfileUpdateResponseStatusPending  AgentProfileUpdateResponseStatus = "pending"
	AgentProfileUpdateResponseStatusFailed   AgentProfileUpdateResponseStatus = "failed"
	AgentProfileUpdateResponseStatusArchived AgentProfileUpdateResponseStatus = "archived"
	AgentProfileUpdateResponseStatusDisabled AgentProfileUpdateResponseStatus = "disabled"
)

// Newest thread excerpt (90 characters), awaits-input and timestamp, from at most
// 20 newest rows. Null when the thread has no rows.
type AgentProfileUpdateResponseThreadPreview struct {
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
func (r AgentProfileUpdateResponseThreadPreview) RawJSON() string { return r.JSON.raw }
func (r *AgentProfileUpdateResponseThreadPreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the Agent is running right now, derived from stored state only.
// `waking`: an active start, its machine is provisioning or a restart/rehydrate is
// pending and changed within the last 10 minutes (a timed-out or stalled start
// reads `sleeping`). `awake`: the machine is running and this Agent has been
// started in the current run. `sleeping`: everything else (stopped or parked
// machine, an Agent the running harness has not started yet, or an Agent that is
// not active). `awake` can lag a machine stop until the reconciler converges.
type AgentProfileUpdateResponseRuntime string

const (
	AgentProfileUpdateResponseRuntimeAwake    AgentProfileUpdateResponseRuntime = "awake"
	AgentProfileUpdateResponseRuntimeSleeping AgentProfileUpdateResponseRuntime = "sleeping"
	AgentProfileUpdateResponseRuntimeWaking   AgentProfileUpdateResponseRuntime = "waking"
)

type AgentProfileUpdateParams struct {
	// Single-line name shown for the Agent; surrounding whitespace is trimmed and it
	// needs at least one letter or number.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// The Agent's individual instructions passage; an empty or whitespace-only text
	// clears it. Line breaks count three characters toward the rendered size limit.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// Model profile id from GET /agents/model-profiles; null uses the environment
	// default.
	ModelProfile param.Opt[string]                    `json:"modelProfile,omitzero"`
	AvatarConfig AgentProfileUpdateParamsAvatarConfig `json:"avatarConfig,omitzero"`
	// Any of "blob".
	AvatarKind AgentProfileUpdateParamsAvatarKind `json:"avatarKind,omitzero"`
	// Set the avatar palette color without changing its seed; null removes only the
	// color. Omitted color stays unchanged, including during a seed update.
	//
	// Any of "slate", "red", "orange", "amber", "lime", "green", "teal", "cyan",
	// "blue", "indigo", "violet", "pink".
	Color AgentProfileUpdateParamsColor `json:"color,omitzero"`
	paramObj
}

func (r AgentProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Seed is required.
type AgentProfileUpdateParamsAvatarConfig struct {
	Seed AgentProfileUpdateParamsAvatarConfigSeedUnion `json:"seed,omitzero" api:"required"`
	paramObj
}

func (r AgentProfileUpdateParamsAvatarConfig) MarshalJSON() (data []byte, err error) {
	type shadow AgentProfileUpdateParamsAvatarConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentProfileUpdateParamsAvatarConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentProfileUpdateParamsAvatarConfigSeedUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u AgentProfileUpdateParamsAvatarConfigSeedUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *AgentProfileUpdateParamsAvatarConfigSeedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type AgentProfileUpdateParamsAvatarKind string

const (
	AgentProfileUpdateParamsAvatarKindBlob AgentProfileUpdateParamsAvatarKind = "blob"
)

// Set the avatar palette color without changing its seed; null removes only the
// color. Omitted color stays unchanged, including during a seed update.
type AgentProfileUpdateParamsColor string

const (
	AgentProfileUpdateParamsColorSlate  AgentProfileUpdateParamsColor = "slate"
	AgentProfileUpdateParamsColorRed    AgentProfileUpdateParamsColor = "red"
	AgentProfileUpdateParamsColorOrange AgentProfileUpdateParamsColor = "orange"
	AgentProfileUpdateParamsColorAmber  AgentProfileUpdateParamsColor = "amber"
	AgentProfileUpdateParamsColorLime   AgentProfileUpdateParamsColor = "lime"
	AgentProfileUpdateParamsColorGreen  AgentProfileUpdateParamsColor = "green"
	AgentProfileUpdateParamsColorTeal   AgentProfileUpdateParamsColor = "teal"
	AgentProfileUpdateParamsColorCyan   AgentProfileUpdateParamsColor = "cyan"
	AgentProfileUpdateParamsColorBlue   AgentProfileUpdateParamsColor = "blue"
	AgentProfileUpdateParamsColorIndigo AgentProfileUpdateParamsColor = "indigo"
	AgentProfileUpdateParamsColorViolet AgentProfileUpdateParamsColor = "violet"
	AgentProfileUpdateParamsColorPink   AgentProfileUpdateParamsColor = "pink"
)
