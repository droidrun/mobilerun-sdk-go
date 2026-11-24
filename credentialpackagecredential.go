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

// CredentialPackageCredentialService contains methods and other services that help
// with interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialPackageCredentialService] method instead.
type CredentialPackageCredentialService struct {
	options []option.RequestOption
	Fields  CredentialPackageCredentialFieldService
}

// NewCredentialPackageCredentialService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCredentialPackageCredentialService(opts ...option.RequestOption) (r CredentialPackageCredentialService) {
	r = CredentialPackageCredentialService{}
	r.options = opts
	r.Fields = NewCredentialPackageCredentialFieldService(opts...)
	return
}

// Creates a credential under the given package with a `credentialName` and at
// least one field. Field types must be unique. Each field has a `fieldType`
// (email, username, password, api_token, phone_number, two_factor_secret, or
// backup_codes) and a value.
func (r *CredentialPackageCredentialService) New(ctx context.Context, packageName string, body CredentialPackageCredentialNewParams, opts ...option.RequestOption) (res *CredentialPackageCredentialNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/packages/%s", url.PathEscape(packageName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Permanently deletes the credential identified by `packageName` and
// `credentialName`, removing all of its fields. Returns the deleted credential.
func (r *CredentialPackageCredentialService) Delete(ctx context.Context, credentialName string, body CredentialPackageCredentialDeleteParams, opts ...option.RequestOption) (res *CredentialPackageCredentialDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.PackageName == "" {
		err = errors.New("missing required packageName parameter")
		return nil, err
	}
	if credentialName == "" {
		err = errors.New("missing required credentialName parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/packages/%s/credentials/%s", url.PathEscape(body.PackageName), url.PathEscape(credentialName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type CredentialPackageCredentialNewResponse struct {
	Data    CredentialPackageCredentialNewResponseData `json:"data" api:"required"`
	Message string                                     `json:"message" api:"required"`
	// Any of true.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageCredentialNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialNewResponseData struct {
	CreatedBy      string                                            `json:"createdBy" api:"required"`
	CredentialName string                                            `json:"credentialName" api:"required"`
	Fields         []CredentialPackageCredentialNewResponseDataField `json:"fields" api:"required"`
	OwnerID        string                                            `json:"ownerId" api:"required"`
	PackageName    string                                            `json:"packageName" api:"required"`
	SecretPath     string                                            `json:"secretPath" api:"required"`
	// Deprecated: use createdBy (same value — the creating actor). Null for
	// credentials created before rollout.
	//
	// Deprecated: deprecated
	UserID string                                              `json:"userId" api:"required"`
	Label  string                                              `json:"label"`
	Link   CredentialPackageCredentialNewResponseDataLinkUnion `json:"link"`
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
func (r CredentialPackageCredentialNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialNewResponseDataField struct {
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
func (r CredentialPackageCredentialNewResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialNewResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageCredentialNewResponseDataLinkUnion contains all possible
// properties and values from
// [CredentialPackageCredentialNewResponseDataLinkObject],
// [CredentialPackageCredentialNewResponseDataLinkObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CredentialPackageCredentialNewResponseDataLinkUnion struct {
	// This field is a union of
	// [[]CredentialPackageCredentialNewResponseDataLinkObjectFieldMap],
	// [[]CredentialPackageCredentialNewResponseDataLinkObject2FieldMap]
	FieldMap CredentialPackageCredentialNewResponseDataLinkUnionFieldMap `json:"fieldMap"`
	// This field is from variant
	// [CredentialPackageCredentialNewResponseDataLinkObject].
	ItemID string `json:"itemId"`
	// This field is from variant
	// [CredentialPackageCredentialNewResponseDataLinkObject].
	ItemTitle string `json:"itemTitle"`
	Provider  string `json:"provider"`
	// This field is from variant
	// [CredentialPackageCredentialNewResponseDataLinkObject].
	VaultID string `json:"vaultId"`
	// This field is from variant
	// [CredentialPackageCredentialNewResponseDataLinkObject].
	VaultName string `json:"vaultName"`
	// This field is from variant
	// [CredentialPackageCredentialNewResponseDataLinkObject2].
	CardID string `json:"cardId"`
	// This field is from variant
	// [CredentialPackageCredentialNewResponseDataLinkObject2].
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

func (u CredentialPackageCredentialNewResponseDataLinkUnion) AsCredentialPackageCredentialNewResponseDataLinkObject() (v CredentialPackageCredentialNewResponseDataLinkObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialPackageCredentialNewResponseDataLinkUnion) AsCredentialPackageCredentialNewResponseDataLinkObject2() (v CredentialPackageCredentialNewResponseDataLinkObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CredentialPackageCredentialNewResponseDataLinkUnion) RawJSON() string { return u.JSON.raw }

func (r *CredentialPackageCredentialNewResponseDataLinkUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageCredentialNewResponseDataLinkUnionFieldMap is an implicit
// subunion of [CredentialPackageCredentialNewResponseDataLinkUnion].
// CredentialPackageCredentialNewResponseDataLinkUnionFieldMap provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [CredentialPackageCredentialNewResponseDataLinkUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfCredentialPackageCredentialNewResponseDataLinkObjectFieldMapArray
// OfCredentialPackageCredentialNewResponseDataLinkObject2FieldMapArray]
type CredentialPackageCredentialNewResponseDataLinkUnionFieldMap struct {
	// This field will be present if the value is a
	// [[]CredentialPackageCredentialNewResponseDataLinkObjectFieldMap] instead of an
	// object.
	OfCredentialPackageCredentialNewResponseDataLinkObjectFieldMapArray []CredentialPackageCredentialNewResponseDataLinkObjectFieldMap `json:",inline"`
	// This field will be present if the value is a
	// [[]CredentialPackageCredentialNewResponseDataLinkObject2FieldMap] instead of an
	// object.
	OfCredentialPackageCredentialNewResponseDataLinkObject2FieldMapArray []CredentialPackageCredentialNewResponseDataLinkObject2FieldMap `json:",inline"`
	JSON                                                                 struct {
		OfCredentialPackageCredentialNewResponseDataLinkObjectFieldMapArray  respjson.Field
		OfCredentialPackageCredentialNewResponseDataLinkObject2FieldMapArray respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (r *CredentialPackageCredentialNewResponseDataLinkUnionFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialNewResponseDataLinkObject struct {
	FieldMap  []CredentialPackageCredentialNewResponseDataLinkObjectFieldMap `json:"fieldMap" api:"required"`
	ItemID    string                                                         `json:"itemId" api:"required"`
	ItemTitle string                                                         `json:"itemTitle" api:"required"`
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
func (r CredentialPackageCredentialNewResponseDataLinkObject) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialNewResponseDataLinkObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialNewResponseDataLinkObjectFieldMap struct {
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
func (r CredentialPackageCredentialNewResponseDataLinkObjectFieldMap) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialNewResponseDataLinkObjectFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialNewResponseDataLinkObject2 struct {
	CardID   string                                                          `json:"cardId" api:"required" format:"uuid"`
	FieldMap []CredentialPackageCredentialNewResponseDataLinkObject2FieldMap `json:"fieldMap" api:"required"`
	Label    string                                                          `json:"label" api:"required"`
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
func (r CredentialPackageCredentialNewResponseDataLinkObject2) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialNewResponseDataLinkObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialNewResponseDataLinkObject2FieldMap struct {
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
func (r CredentialPackageCredentialNewResponseDataLinkObject2FieldMap) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialNewResponseDataLinkObject2FieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialDeleteResponse struct {
	Data    CredentialPackageCredentialDeleteResponseData `json:"data" api:"required"`
	Message string                                        `json:"message" api:"required"`
	// Any of true.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageCredentialDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialDeleteResponseData struct {
	CreatedBy      string                                               `json:"createdBy" api:"required"`
	CredentialName string                                               `json:"credentialName" api:"required"`
	Fields         []CredentialPackageCredentialDeleteResponseDataField `json:"fields" api:"required"`
	OwnerID        string                                               `json:"ownerId" api:"required"`
	PackageName    string                                               `json:"packageName" api:"required"`
	SecretPath     string                                               `json:"secretPath" api:"required"`
	// Deprecated: use createdBy (same value — the creating actor). Null for
	// credentials created before rollout.
	//
	// Deprecated: deprecated
	UserID string                                                 `json:"userId" api:"required"`
	Label  string                                                 `json:"label"`
	Link   CredentialPackageCredentialDeleteResponseDataLinkUnion `json:"link"`
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
func (r CredentialPackageCredentialDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialDeleteResponseDataField struct {
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
func (r CredentialPackageCredentialDeleteResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialDeleteResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageCredentialDeleteResponseDataLinkUnion contains all possible
// properties and values from
// [CredentialPackageCredentialDeleteResponseDataLinkObject],
// [CredentialPackageCredentialDeleteResponseDataLinkObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CredentialPackageCredentialDeleteResponseDataLinkUnion struct {
	// This field is a union of
	// [[]CredentialPackageCredentialDeleteResponseDataLinkObjectFieldMap],
	// [[]CredentialPackageCredentialDeleteResponseDataLinkObject2FieldMap]
	FieldMap CredentialPackageCredentialDeleteResponseDataLinkUnionFieldMap `json:"fieldMap"`
	// This field is from variant
	// [CredentialPackageCredentialDeleteResponseDataLinkObject].
	ItemID string `json:"itemId"`
	// This field is from variant
	// [CredentialPackageCredentialDeleteResponseDataLinkObject].
	ItemTitle string `json:"itemTitle"`
	Provider  string `json:"provider"`
	// This field is from variant
	// [CredentialPackageCredentialDeleteResponseDataLinkObject].
	VaultID string `json:"vaultId"`
	// This field is from variant
	// [CredentialPackageCredentialDeleteResponseDataLinkObject].
	VaultName string `json:"vaultName"`
	// This field is from variant
	// [CredentialPackageCredentialDeleteResponseDataLinkObject2].
	CardID string `json:"cardId"`
	// This field is from variant
	// [CredentialPackageCredentialDeleteResponseDataLinkObject2].
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

func (u CredentialPackageCredentialDeleteResponseDataLinkUnion) AsCredentialPackageCredentialDeleteResponseDataLinkObject() (v CredentialPackageCredentialDeleteResponseDataLinkObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialPackageCredentialDeleteResponseDataLinkUnion) AsCredentialPackageCredentialDeleteResponseDataLinkObject2() (v CredentialPackageCredentialDeleteResponseDataLinkObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CredentialPackageCredentialDeleteResponseDataLinkUnion) RawJSON() string { return u.JSON.raw }

func (r *CredentialPackageCredentialDeleteResponseDataLinkUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageCredentialDeleteResponseDataLinkUnionFieldMap is an implicit
// subunion of [CredentialPackageCredentialDeleteResponseDataLinkUnion].
// CredentialPackageCredentialDeleteResponseDataLinkUnionFieldMap provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [CredentialPackageCredentialDeleteResponseDataLinkUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfCredentialPackageCredentialDeleteResponseDataLinkObjectFieldMapArray
// OfCredentialPackageCredentialDeleteResponseDataLinkObject2FieldMapArray]
type CredentialPackageCredentialDeleteResponseDataLinkUnionFieldMap struct {
	// This field will be present if the value is a
	// [[]CredentialPackageCredentialDeleteResponseDataLinkObjectFieldMap] instead of
	// an object.
	OfCredentialPackageCredentialDeleteResponseDataLinkObjectFieldMapArray []CredentialPackageCredentialDeleteResponseDataLinkObjectFieldMap `json:",inline"`
	// This field will be present if the value is a
	// [[]CredentialPackageCredentialDeleteResponseDataLinkObject2FieldMap] instead of
	// an object.
	OfCredentialPackageCredentialDeleteResponseDataLinkObject2FieldMapArray []CredentialPackageCredentialDeleteResponseDataLinkObject2FieldMap `json:",inline"`
	JSON                                                                    struct {
		OfCredentialPackageCredentialDeleteResponseDataLinkObjectFieldMapArray  respjson.Field
		OfCredentialPackageCredentialDeleteResponseDataLinkObject2FieldMapArray respjson.Field
		raw                                                                     string
	} `json:"-"`
}

func (r *CredentialPackageCredentialDeleteResponseDataLinkUnionFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialDeleteResponseDataLinkObject struct {
	FieldMap  []CredentialPackageCredentialDeleteResponseDataLinkObjectFieldMap `json:"fieldMap" api:"required"`
	ItemID    string                                                            `json:"itemId" api:"required"`
	ItemTitle string                                                            `json:"itemTitle" api:"required"`
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
func (r CredentialPackageCredentialDeleteResponseDataLinkObject) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialDeleteResponseDataLinkObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialDeleteResponseDataLinkObjectFieldMap struct {
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
func (r CredentialPackageCredentialDeleteResponseDataLinkObjectFieldMap) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialDeleteResponseDataLinkObjectFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialDeleteResponseDataLinkObject2 struct {
	CardID   string                                                             `json:"cardId" api:"required" format:"uuid"`
	FieldMap []CredentialPackageCredentialDeleteResponseDataLinkObject2FieldMap `json:"fieldMap" api:"required"`
	Label    string                                                             `json:"label" api:"required"`
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
func (r CredentialPackageCredentialDeleteResponseDataLinkObject2) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialDeleteResponseDataLinkObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialDeleteResponseDataLinkObject2FieldMap struct {
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
func (r CredentialPackageCredentialDeleteResponseDataLinkObject2FieldMap) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialDeleteResponseDataLinkObject2FieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialNewParams struct {
	CredentialName string                                      `json:"credentialName" api:"required"`
	Fields         []CredentialPackageCredentialNewParamsField `json:"fields,omitzero" api:"required"`
	Label          param.Opt[string]                           `json:"label,omitzero"`
	paramObj
}

func (r CredentialPackageCredentialNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CredentialPackageCredentialNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialPackageCredentialNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FieldType, Value are required.
type CredentialPackageCredentialNewParamsField struct {
	// Any of "email", "username", "password", "api_token", "phone_number",
	// "two_factor_secret", "backup_codes", "cc_number", "cc_exp", "cc_csc", "cc_name".
	FieldType string `json:"fieldType,omitzero" api:"required"`
	Value     string `json:"value" api:"required"`
	paramObj
}

func (r CredentialPackageCredentialNewParamsField) MarshalJSON() (data []byte, err error) {
	type shadow CredentialPackageCredentialNewParamsField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialPackageCredentialNewParamsField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CredentialPackageCredentialNewParamsField](
		"fieldType", "email", "username", "password", "api_token", "phone_number", "two_factor_secret", "backup_codes", "cc_number", "cc_exp", "cc_csc", "cc_name",
	)
}

type CredentialPackageCredentialDeleteParams struct {
	PackageName string `path:"packageName" api:"required" json:"-"`
	paramObj
}
