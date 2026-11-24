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
	Data CredentialPackageLinkNewResponseData `json:"data" api:"required"`
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

type CredentialPackageLinkNewResponseData struct {
	CreatedBy      string                                      `json:"createdBy" api:"required"`
	CredentialName string                                      `json:"credentialName" api:"required"`
	Fields         []CredentialPackageLinkNewResponseDataField `json:"fields" api:"required"`
	OwnerID        string                                      `json:"ownerId" api:"required"`
	PackageName    string                                      `json:"packageName" api:"required"`
	SecretPath     string                                      `json:"secretPath" api:"required"`
	// Deprecated: use createdBy (same value — the creating actor). Null for
	// credentials created before rollout.
	//
	// Deprecated: deprecated
	UserID string                                        `json:"userId" api:"required"`
	Label  string                                        `json:"label"`
	Link   CredentialPackageLinkNewResponseDataLinkUnion `json:"link"`
	// Any of "mobilerun", "onepassword", "wallet".
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy      respjson.Field
		CredentialName respjson.Field
		Fields         respjson.Field
		OwnerID        respjson.Field
		PackageName    respjson.Field
		SecretPath     respjson.Field
		UserID         respjson.Field
		Label          respjson.Field
		Link           respjson.Field
		Source         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageLinkNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageLinkNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageLinkNewResponseDataField struct {
	// Any of "email", "username", "password", "api_token", "phone_number",
	// "two_factor_secret", "backup_codes", "cc_number", "cc_exp", "cc_csc", "cc_name".
	FieldType string `json:"fieldType" api:"required"`
	Value     string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldType   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageLinkNewResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageLinkNewResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageLinkNewResponseDataLinkUnion contains all possible properties
// and values from [CredentialPackageLinkNewResponseDataLinkObject],
// [CredentialPackageLinkNewResponseDataLinkObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CredentialPackageLinkNewResponseDataLinkUnion struct {
	// This field is a union of
	// [[]CredentialPackageLinkNewResponseDataLinkObjectFieldMap],
	// [[]CredentialPackageLinkNewResponseDataLinkObject2FieldMap]
	FieldMap CredentialPackageLinkNewResponseDataLinkUnionFieldMap `json:"fieldMap"`
	// This field is from variant [CredentialPackageLinkNewResponseDataLinkObject].
	ItemID string `json:"itemId"`
	// This field is from variant [CredentialPackageLinkNewResponseDataLinkObject].
	ItemTitle string `json:"itemTitle"`
	Provider  string `json:"provider"`
	// This field is from variant [CredentialPackageLinkNewResponseDataLinkObject].
	VaultID string `json:"vaultId"`
	// This field is from variant [CredentialPackageLinkNewResponseDataLinkObject].
	VaultName string `json:"vaultName"`
	// This field is from variant [CredentialPackageLinkNewResponseDataLinkObject2].
	CardID string `json:"cardId"`
	// This field is from variant [CredentialPackageLinkNewResponseDataLinkObject2].
	Label string `json:"label"`
	JSON  struct {
		FieldMap  respjson.Field
		ItemID    respjson.Field
		ItemTitle respjson.Field
		Provider  respjson.Field
		VaultID   respjson.Field
		VaultName respjson.Field
		CardID    respjson.Field
		Label     respjson.Field
		raw       string
	} `json:"-"`
}

func (u CredentialPackageLinkNewResponseDataLinkUnion) AsCredentialPackageLinkNewResponseDataLinkObject() (v CredentialPackageLinkNewResponseDataLinkObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialPackageLinkNewResponseDataLinkUnion) AsCredentialPackageLinkNewResponseDataLinkObject2() (v CredentialPackageLinkNewResponseDataLinkObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CredentialPackageLinkNewResponseDataLinkUnion) RawJSON() string { return u.JSON.raw }

func (r *CredentialPackageLinkNewResponseDataLinkUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageLinkNewResponseDataLinkUnionFieldMap is an implicit subunion of
// [CredentialPackageLinkNewResponseDataLinkUnion].
// CredentialPackageLinkNewResponseDataLinkUnionFieldMap provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [CredentialPackageLinkNewResponseDataLinkUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCredentialPackageLinkNewResponseDataLinkObjectFieldMapArray
// OfCredentialPackageLinkNewResponseDataLinkObject2FieldMapArray]
type CredentialPackageLinkNewResponseDataLinkUnionFieldMap struct {
	// This field will be present if the value is a
	// [[]CredentialPackageLinkNewResponseDataLinkObjectFieldMap] instead of an object.
	OfCredentialPackageLinkNewResponseDataLinkObjectFieldMapArray []CredentialPackageLinkNewResponseDataLinkObjectFieldMap `json:",inline"`
	// This field will be present if the value is a
	// [[]CredentialPackageLinkNewResponseDataLinkObject2FieldMap] instead of an
	// object.
	OfCredentialPackageLinkNewResponseDataLinkObject2FieldMapArray []CredentialPackageLinkNewResponseDataLinkObject2FieldMap `json:",inline"`
	JSON                                                           struct {
		OfCredentialPackageLinkNewResponseDataLinkObjectFieldMapArray  respjson.Field
		OfCredentialPackageLinkNewResponseDataLinkObject2FieldMapArray respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (r *CredentialPackageLinkNewResponseDataLinkUnionFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageLinkNewResponseDataLinkObject struct {
	FieldMap  []CredentialPackageLinkNewResponseDataLinkObjectFieldMap `json:"fieldMap" api:"required"`
	ItemID    string                                                   `json:"itemId" api:"required"`
	ItemTitle string                                                   `json:"itemTitle" api:"required"`
	// Any of "onepassword".
	Provider  string `json:"provider" api:"required"`
	VaultID   string `json:"vaultId" api:"required"`
	VaultName string `json:"vaultName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldMap    respjson.Field
		ItemID      respjson.Field
		ItemTitle   respjson.Field
		Provider    respjson.Field
		VaultID     respjson.Field
		VaultName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageLinkNewResponseDataLinkObject) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageLinkNewResponseDataLinkObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageLinkNewResponseDataLinkObjectFieldMap struct {
	FieldID string `json:"fieldId" api:"required"`
	// Any of "email", "username", "password", "api_token", "phone_number",
	// "two_factor_secret", "backup_codes", "cc_number", "cc_exp", "cc_csc", "cc_name".
	FieldType string `json:"fieldType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldID     respjson.Field
		FieldType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageLinkNewResponseDataLinkObjectFieldMap) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageLinkNewResponseDataLinkObjectFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageLinkNewResponseDataLinkObject2 struct {
	CardID   string                                                    `json:"cardId" api:"required" format:"uuid"`
	FieldMap []CredentialPackageLinkNewResponseDataLinkObject2FieldMap `json:"fieldMap" api:"required"`
	Label    string                                                    `json:"label" api:"required"`
	// Any of "wallet".
	Provider string `json:"provider" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CardID      respjson.Field
		FieldMap    respjson.Field
		Label       respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageLinkNewResponseDataLinkObject2) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageLinkNewResponseDataLinkObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageLinkNewResponseDataLinkObject2FieldMap struct {
	FieldID string `json:"fieldId" api:"required"`
	// Any of "email", "username", "password", "api_token", "phone_number",
	// "two_factor_secret", "backup_codes", "cc_number", "cc_exp", "cc_csc", "cc_name".
	FieldType string `json:"fieldType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldID     respjson.Field
		FieldType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageLinkNewResponseDataLinkObject2FieldMap) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageLinkNewResponseDataLinkObject2FieldMap) UnmarshalJSON(data []byte) error {
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
