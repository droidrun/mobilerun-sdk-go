// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"net/http"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// IdentityService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityService] method instead.
type IdentityService struct {
	options []option.RequestOption
}

// NewIdentityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIdentityService(opts ...option.RequestOption) (r IdentityService) {
	r = IdentityService{}
	r.options = opts
	return
}

// Return the user, owner and authentication method of the calling credential.
func (r *IdentityService) Me(ctx context.Context, opts ...option.RequestOption) (res *IdentityMeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "identity/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type IdentityMeResponse struct {
	// Any of "api_key", "session".
	AuthMethod IdentityMeResponseAuthMethod `json:"authMethod" api:"required"`
	OwnerID    string                       `json:"ownerId" api:"required"`
	UserID     string                       `json:"userId" api:"required"`
	APIKeyID   string                       `json:"apiKeyId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthMethod  respjson.Field
		OwnerID     respjson.Field
		UserID      respjson.Field
		APIKeyID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityMeResponse) RawJSON() string { return r.JSON.raw }
func (r *IdentityMeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityMeResponseAuthMethod string

const (
	IdentityMeResponseAuthMethodAPIKey  IdentityMeResponseAuthMethod = "api_key"
	IdentityMeResponseAuthMethodSession IdentityMeResponseAuthMethod = "session"
)
