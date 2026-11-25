// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
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
func (r *AgentProfileService) Update(ctx context.Context, agentID string, body AgentProfileUpdateParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("agents/%s/profile", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

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
