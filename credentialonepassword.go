// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"net/http"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// CredentialOnepasswordService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialOnepasswordService] method instead.
type CredentialOnepasswordService struct {
	options []option.RequestOption
	Vaults  CredentialOnepasswordVaultService
}

// NewCredentialOnepasswordService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCredentialOnepasswordService(opts ...option.RequestOption) (r CredentialOnepasswordService) {
	r = CredentialOnepasswordService{}
	r.options = opts
	r.Vaults = NewCredentialOnepasswordVaultService(opts...)
	return
}

// Get 1Password connection
func (r *CredentialOnepasswordService) Get(ctx context.Context, opts ...option.RequestOption) (res *CredentialOnepasswordGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credentials/onepassword"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Connect a 1Password service account
func (r *CredentialOnepasswordService) Connect(ctx context.Context, body CredentialOnepasswordConnectParams, opts ...option.RequestOption) (res *CredentialOnepasswordConnectResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credentials/onepassword"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Disconnect 1Password
func (r *CredentialOnepasswordService) Disconnect(ctx context.Context, opts ...option.RequestOption) (res *CredentialOnepasswordDisconnectResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credentials/onepassword"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type OnePasswordConnection struct {
	Connected   bool                         `json:"connected" api:"required"`
	ConnectedAt string                       `json:"connectedAt"`
	ConnectedBy string                       `json:"connectedBy"`
	Vaults      []OnePasswordConnectionVault `json:"vaults"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connected   respjson.Field
		ConnectedAt respjson.Field
		ConnectedBy respjson.Field
		Vaults      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnePasswordConnection) RawJSON() string { return r.JSON.raw }
func (r *OnePasswordConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnePasswordConnectionVault struct {
	ID    string `json:"id" api:"required"`
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnePasswordConnectionVault) RawJSON() string { return r.JSON.raw }
func (r *OnePasswordConnectionVault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialOnepasswordGetResponse struct {
	Data OnePasswordConnection `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialOnepasswordGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialOnepasswordGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialOnepasswordConnectResponse struct {
	Data OnePasswordConnection `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialOnepasswordConnectResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialOnepasswordConnectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialOnepasswordDisconnectResponse struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialOnepasswordDisconnectResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialOnepasswordDisconnectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialOnepasswordConnectParams struct {
	ServiceAccountToken string `json:"serviceAccountToken" api:"required"`
	paramObj
}

func (r CredentialOnepasswordConnectParams) MarshalJSON() (data []byte, err error) {
	type shadow CredentialOnepasswordConnectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialOnepasswordConnectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
