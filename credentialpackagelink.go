// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// CredentialPackageLinkService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialPackageLinkService] method instead.
type CredentialPackageLinkService struct {
	options []option.RequestOption
}

// NewCredentialPackageLinkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCredentialPackageLinkService(opts ...option.RequestOption) (r CredentialPackageLinkService) {
	r = CredentialPackageLinkService{}
	r.options = opts
	return
}

// Creates a read-only credential that resolves its values live from the provider.
// `provider` defaults to `onepassword` (vaultId, itemId, fieldMap). `wallet` takes
// a cardId, is only allowed in the reserved payment_cards package, and only the
// cardholder may link the card (404 otherwise). Never replaces an existing
// credential of the same name.
func (r *CredentialPackageLinkService) New(ctx context.Context, packageName string, body CredentialPackageLinkNewParams, opts ...option.RequestOption) (res *CredentialPackageLinkNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/packages/%s/links", url.PathEscape(packageName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CredentialPackageLinkNewResponse struct {
	Data Credential `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageLinkNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageLinkNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageLinkNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfLinkOnePasswordCredentialBody *CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBody `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfLinkWalletCredentialBody *CredentialPackageLinkNewParamsBodyLinkWalletCredentialBody `json:",inline"`

	paramObj
}

func (u CredentialPackageLinkNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLinkOnePasswordCredentialBody, u.OfLinkWalletCredentialBody)
}
func (r *CredentialPackageLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CredentialName, FieldMap, ItemID, VaultID are required.
type CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBody struct {
	CredentialName string                                                                    `json:"credentialName" api:"required"`
	FieldMap       []CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBodyFieldMap `json:"fieldMap,omitzero" api:"required"`
	ItemID         string                                                                    `json:"itemId" api:"required"`
	VaultID        string                                                                    `json:"vaultId" api:"required"`
	// Any of "onepassword".
	Provider string `json:"provider,omitzero"`
	paramObj
}

func (r CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBody) MarshalJSON() (data []byte, err error) {
	type shadow CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBody](
		"provider", "onepassword",
	)
}

// The properties FieldID, FieldType are required.
type CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBodyFieldMap struct {
	FieldID string `json:"fieldId" api:"required"`
	// Any of "email", "username", "password", "api_token", "phone_number",
	// "two_factor_secret", "backup_codes", "cc_number", "cc_exp", "cc_csc", "cc_name".
	FieldType string `json:"fieldType,omitzero" api:"required"`
	paramObj
}

func (r CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBodyFieldMap) MarshalJSON() (data []byte, err error) {
	type shadow CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBodyFieldMap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBodyFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CredentialPackageLinkNewParamsBodyLinkOnePasswordCredentialBodyFieldMap](
		"fieldType", "email", "username", "password", "api_token", "phone_number", "two_factor_secret", "backup_codes", "cc_number", "cc_exp", "cc_csc", "cc_name",
	)
}

// The properties CardID, CredentialName, Provider are required.
type CredentialPackageLinkNewParamsBodyLinkWalletCredentialBody struct {
	CardID         string `json:"cardId" api:"required" format:"uuid"`
	CredentialName string `json:"credentialName" api:"required"`
	// Any of "wallet".
	Provider string `json:"provider,omitzero" api:"required"`
	paramObj
}

func (r CredentialPackageLinkNewParamsBodyLinkWalletCredentialBody) MarshalJSON() (data []byte, err error) {
	type shadow CredentialPackageLinkNewParamsBodyLinkWalletCredentialBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialPackageLinkNewParamsBodyLinkWalletCredentialBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CredentialPackageLinkNewParamsBodyLinkWalletCredentialBody](
		"provider", "wallet",
	)
}
