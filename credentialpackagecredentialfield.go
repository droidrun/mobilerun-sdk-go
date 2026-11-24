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

// CredentialPackageCredentialFieldService contains methods and other services that
// help with interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialPackageCredentialFieldService] method instead.
type CredentialPackageCredentialFieldService struct {
	options []option.RequestOption
}

// NewCredentialPackageCredentialFieldService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCredentialPackageCredentialFieldService(opts ...option.RequestOption) (r CredentialPackageCredentialFieldService) {
	r = CredentialPackageCredentialFieldService{}
	r.options = opts
	return
}

// Adds a single field to an existing credential. The body specifies a `fieldType`
// (one of the supported field types) and its value. Returns a conflict if a field
// of that type already exists on the credential.
func (r *CredentialPackageCredentialFieldService) New(ctx context.Context, credentialName string, params CredentialPackageCredentialFieldNewParams, opts ...option.RequestOption) (res *CredentialPackageCredentialFieldNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.PackageName == "" {
		err = errors.New("missing required packageName parameter")
		return nil, err
	}
	if credentialName == "" {
		err = errors.New("missing required credentialName parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/packages/%s/credentials/%s/fields", url.PathEscape(params.PackageName), url.PathEscape(credentialName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Updates the value of an existing field on a credential, identified by
// `packageName`, `credentialName`, and `fieldType` in the path. The body carries
// the new value and returns the updated credential.
func (r *CredentialPackageCredentialFieldService) Update(ctx context.Context, fieldType CredentialPackageCredentialFieldUpdateParamsFieldType, params CredentialPackageCredentialFieldUpdateParams, opts ...option.RequestOption) (res *CredentialPackageCredentialFieldUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.PackageName == "" {
		err = errors.New("missing required packageName parameter")
		return nil, err
	}
	if params.CredentialName == "" {
		err = errors.New("missing required credentialName parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/packages/%s/credentials/%s/fields/%v", url.PathEscape(params.PackageName), url.PathEscape(params.CredentialName), fieldType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Removes a single field of the given `fieldType` from the specified credential
// while leaving the credential itself intact. Returns the updated credential.
func (r *CredentialPackageCredentialFieldService) Delete(ctx context.Context, fieldType CredentialPackageCredentialFieldDeleteParamsFieldType, body CredentialPackageCredentialFieldDeleteParams, opts ...option.RequestOption) (res *CredentialPackageCredentialFieldDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.PackageName == "" {
		err = errors.New("missing required packageName parameter")
		return nil, err
	}
	if body.CredentialName == "" {
		err = errors.New("missing required credentialName parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/packages/%s/credentials/%s/fields/%v", url.PathEscape(body.PackageName), url.PathEscape(body.CredentialName), fieldType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type CredentialPackageCredentialFieldNewResponse struct {
	Data    CredentialPackageCredentialFieldNewResponseData `json:"data" api:"required"`
	Message string                                          `json:"message" api:"required"`
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
func (r CredentialPackageCredentialFieldNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldNewResponseData struct {
	CreatedBy      string                                                 `json:"createdBy" api:"required"`
	CredentialName string                                                 `json:"credentialName" api:"required"`
	Fields         []CredentialPackageCredentialFieldNewResponseDataField `json:"fields" api:"required"`
	OwnerID        string                                                 `json:"ownerId" api:"required"`
	PackageName    string                                                 `json:"packageName" api:"required"`
	SecretPath     string                                                 `json:"secretPath" api:"required"`
	// Deprecated: use createdBy (same value — the creating actor). Null for
	// credentials created before rollout.
	//
	// Deprecated: deprecated
	UserID string                                                   `json:"userId" api:"required"`
	Label  string                                                   `json:"label"`
	Link   CredentialPackageCredentialFieldNewResponseDataLinkUnion `json:"link"`
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
func (r CredentialPackageCredentialFieldNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldNewResponseDataField struct {
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
func (r CredentialPackageCredentialFieldNewResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldNewResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageCredentialFieldNewResponseDataLinkUnion contains all possible
// properties and values from
// [CredentialPackageCredentialFieldNewResponseDataLinkObject],
// [CredentialPackageCredentialFieldNewResponseDataLinkObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CredentialPackageCredentialFieldNewResponseDataLinkUnion struct {
	// This field is a union of
	// [[]CredentialPackageCredentialFieldNewResponseDataLinkObjectFieldMap],
	// [[]CredentialPackageCredentialFieldNewResponseDataLinkObject2FieldMap]
	FieldMap CredentialPackageCredentialFieldNewResponseDataLinkUnionFieldMap `json:"fieldMap"`
	// This field is from variant
	// [CredentialPackageCredentialFieldNewResponseDataLinkObject].
	ItemID string `json:"itemId"`
	// This field is from variant
	// [CredentialPackageCredentialFieldNewResponseDataLinkObject].
	ItemTitle string `json:"itemTitle"`
	Provider  string `json:"provider"`
	// This field is from variant
	// [CredentialPackageCredentialFieldNewResponseDataLinkObject].
	VaultID string `json:"vaultId"`
	// This field is from variant
	// [CredentialPackageCredentialFieldNewResponseDataLinkObject].
	VaultName string `json:"vaultName"`
	// This field is from variant
	// [CredentialPackageCredentialFieldNewResponseDataLinkObject2].
	CardID string `json:"cardId"`
	// This field is from variant
	// [CredentialPackageCredentialFieldNewResponseDataLinkObject2].
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

func (u CredentialPackageCredentialFieldNewResponseDataLinkUnion) AsCredentialPackageCredentialFieldNewResponseDataLinkObject() (v CredentialPackageCredentialFieldNewResponseDataLinkObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialPackageCredentialFieldNewResponseDataLinkUnion) AsCredentialPackageCredentialFieldNewResponseDataLinkObject2() (v CredentialPackageCredentialFieldNewResponseDataLinkObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CredentialPackageCredentialFieldNewResponseDataLinkUnion) RawJSON() string { return u.JSON.raw }

func (r *CredentialPackageCredentialFieldNewResponseDataLinkUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageCredentialFieldNewResponseDataLinkUnionFieldMap is an implicit
// subunion of [CredentialPackageCredentialFieldNewResponseDataLinkUnion].
// CredentialPackageCredentialFieldNewResponseDataLinkUnionFieldMap provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [CredentialPackageCredentialFieldNewResponseDataLinkUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfCredentialPackageCredentialFieldNewResponseDataLinkObjectFieldMapArray
// OfCredentialPackageCredentialFieldNewResponseDataLinkObject2FieldMapArray]
type CredentialPackageCredentialFieldNewResponseDataLinkUnionFieldMap struct {
	// This field will be present if the value is a
	// [[]CredentialPackageCredentialFieldNewResponseDataLinkObjectFieldMap] instead of
	// an object.
	OfCredentialPackageCredentialFieldNewResponseDataLinkObjectFieldMapArray []CredentialPackageCredentialFieldNewResponseDataLinkObjectFieldMap `json:",inline"`
	// This field will be present if the value is a
	// [[]CredentialPackageCredentialFieldNewResponseDataLinkObject2FieldMap] instead
	// of an object.
	OfCredentialPackageCredentialFieldNewResponseDataLinkObject2FieldMapArray []CredentialPackageCredentialFieldNewResponseDataLinkObject2FieldMap `json:",inline"`
	JSON                                                                      struct {
		OfCredentialPackageCredentialFieldNewResponseDataLinkObjectFieldMapArray  respjson.Field
		OfCredentialPackageCredentialFieldNewResponseDataLinkObject2FieldMapArray respjson.Field
		raw                                                                       string
	} `json:"-"`
}

func (r *CredentialPackageCredentialFieldNewResponseDataLinkUnionFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldNewResponseDataLinkObject struct {
	FieldMap  []CredentialPackageCredentialFieldNewResponseDataLinkObjectFieldMap `json:"fieldMap" api:"required"`
	ItemID    string                                                              `json:"itemId" api:"required"`
	ItemTitle string                                                              `json:"itemTitle" api:"required"`
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
func (r CredentialPackageCredentialFieldNewResponseDataLinkObject) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldNewResponseDataLinkObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldNewResponseDataLinkObjectFieldMap struct {
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
func (r CredentialPackageCredentialFieldNewResponseDataLinkObjectFieldMap) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldNewResponseDataLinkObjectFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldNewResponseDataLinkObject2 struct {
	CardID   string                                                               `json:"cardId" api:"required" format:"uuid"`
	FieldMap []CredentialPackageCredentialFieldNewResponseDataLinkObject2FieldMap `json:"fieldMap" api:"required"`
	Label    string                                                               `json:"label" api:"required"`
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
func (r CredentialPackageCredentialFieldNewResponseDataLinkObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldNewResponseDataLinkObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldNewResponseDataLinkObject2FieldMap struct {
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
func (r CredentialPackageCredentialFieldNewResponseDataLinkObject2FieldMap) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldNewResponseDataLinkObject2FieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldUpdateResponse struct {
	Data    CredentialPackageCredentialFieldUpdateResponseData `json:"data" api:"required"`
	Message string                                             `json:"message" api:"required"`
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
func (r CredentialPackageCredentialFieldUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldUpdateResponseData struct {
	CreatedBy      string                                                    `json:"createdBy" api:"required"`
	CredentialName string                                                    `json:"credentialName" api:"required"`
	Fields         []CredentialPackageCredentialFieldUpdateResponseDataField `json:"fields" api:"required"`
	OwnerID        string                                                    `json:"ownerId" api:"required"`
	PackageName    string                                                    `json:"packageName" api:"required"`
	SecretPath     string                                                    `json:"secretPath" api:"required"`
	// Deprecated: use createdBy (same value — the creating actor). Null for
	// credentials created before rollout.
	//
	// Deprecated: deprecated
	UserID string                                                      `json:"userId" api:"required"`
	Label  string                                                      `json:"label"`
	Link   CredentialPackageCredentialFieldUpdateResponseDataLinkUnion `json:"link"`
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
func (r CredentialPackageCredentialFieldUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldUpdateResponseDataField struct {
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
func (r CredentialPackageCredentialFieldUpdateResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldUpdateResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageCredentialFieldUpdateResponseDataLinkUnion contains all
// possible properties and values from
// [CredentialPackageCredentialFieldUpdateResponseDataLinkObject],
// [CredentialPackageCredentialFieldUpdateResponseDataLinkObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CredentialPackageCredentialFieldUpdateResponseDataLinkUnion struct {
	// This field is a union of
	// [[]CredentialPackageCredentialFieldUpdateResponseDataLinkObjectFieldMap],
	// [[]CredentialPackageCredentialFieldUpdateResponseDataLinkObject2FieldMap]
	FieldMap CredentialPackageCredentialFieldUpdateResponseDataLinkUnionFieldMap `json:"fieldMap"`
	// This field is from variant
	// [CredentialPackageCredentialFieldUpdateResponseDataLinkObject].
	ItemID string `json:"itemId"`
	// This field is from variant
	// [CredentialPackageCredentialFieldUpdateResponseDataLinkObject].
	ItemTitle string `json:"itemTitle"`
	Provider  string `json:"provider"`
	// This field is from variant
	// [CredentialPackageCredentialFieldUpdateResponseDataLinkObject].
	VaultID string `json:"vaultId"`
	// This field is from variant
	// [CredentialPackageCredentialFieldUpdateResponseDataLinkObject].
	VaultName string `json:"vaultName"`
	// This field is from variant
	// [CredentialPackageCredentialFieldUpdateResponseDataLinkObject2].
	CardID string `json:"cardId"`
	// This field is from variant
	// [CredentialPackageCredentialFieldUpdateResponseDataLinkObject2].
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

func (u CredentialPackageCredentialFieldUpdateResponseDataLinkUnion) AsCredentialPackageCredentialFieldUpdateResponseDataLinkObject() (v CredentialPackageCredentialFieldUpdateResponseDataLinkObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialPackageCredentialFieldUpdateResponseDataLinkUnion) AsCredentialPackageCredentialFieldUpdateResponseDataLinkObject2() (v CredentialPackageCredentialFieldUpdateResponseDataLinkObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CredentialPackageCredentialFieldUpdateResponseDataLinkUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CredentialPackageCredentialFieldUpdateResponseDataLinkUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageCredentialFieldUpdateResponseDataLinkUnionFieldMap is an
// implicit subunion of
// [CredentialPackageCredentialFieldUpdateResponseDataLinkUnion].
// CredentialPackageCredentialFieldUpdateResponseDataLinkUnionFieldMap provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [CredentialPackageCredentialFieldUpdateResponseDataLinkUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfCredentialPackageCredentialFieldUpdateResponseDataLinkObjectFieldMapArray
// OfCredentialPackageCredentialFieldUpdateResponseDataLinkObject2FieldMapArray]
type CredentialPackageCredentialFieldUpdateResponseDataLinkUnionFieldMap struct {
	// This field will be present if the value is a
	// [[]CredentialPackageCredentialFieldUpdateResponseDataLinkObjectFieldMap] instead
	// of an object.
	OfCredentialPackageCredentialFieldUpdateResponseDataLinkObjectFieldMapArray []CredentialPackageCredentialFieldUpdateResponseDataLinkObjectFieldMap `json:",inline"`
	// This field will be present if the value is a
	// [[]CredentialPackageCredentialFieldUpdateResponseDataLinkObject2FieldMap]
	// instead of an object.
	OfCredentialPackageCredentialFieldUpdateResponseDataLinkObject2FieldMapArray []CredentialPackageCredentialFieldUpdateResponseDataLinkObject2FieldMap `json:",inline"`
	JSON                                                                         struct {
		OfCredentialPackageCredentialFieldUpdateResponseDataLinkObjectFieldMapArray  respjson.Field
		OfCredentialPackageCredentialFieldUpdateResponseDataLinkObject2FieldMapArray respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (r *CredentialPackageCredentialFieldUpdateResponseDataLinkUnionFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldUpdateResponseDataLinkObject struct {
	FieldMap  []CredentialPackageCredentialFieldUpdateResponseDataLinkObjectFieldMap `json:"fieldMap" api:"required"`
	ItemID    string                                                                 `json:"itemId" api:"required"`
	ItemTitle string                                                                 `json:"itemTitle" api:"required"`
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
func (r CredentialPackageCredentialFieldUpdateResponseDataLinkObject) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldUpdateResponseDataLinkObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldUpdateResponseDataLinkObjectFieldMap struct {
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
func (r CredentialPackageCredentialFieldUpdateResponseDataLinkObjectFieldMap) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldUpdateResponseDataLinkObjectFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldUpdateResponseDataLinkObject2 struct {
	CardID   string                                                                  `json:"cardId" api:"required" format:"uuid"`
	FieldMap []CredentialPackageCredentialFieldUpdateResponseDataLinkObject2FieldMap `json:"fieldMap" api:"required"`
	Label    string                                                                  `json:"label" api:"required"`
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
func (r CredentialPackageCredentialFieldUpdateResponseDataLinkObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldUpdateResponseDataLinkObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldUpdateResponseDataLinkObject2FieldMap struct {
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
func (r CredentialPackageCredentialFieldUpdateResponseDataLinkObject2FieldMap) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldUpdateResponseDataLinkObject2FieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldDeleteResponse struct {
	Data    CredentialPackageCredentialFieldDeleteResponseData `json:"data" api:"required"`
	Message string                                             `json:"message" api:"required"`
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
func (r CredentialPackageCredentialFieldDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldDeleteResponseData struct {
	CreatedBy      string                                                    `json:"createdBy" api:"required"`
	CredentialName string                                                    `json:"credentialName" api:"required"`
	Fields         []CredentialPackageCredentialFieldDeleteResponseDataField `json:"fields" api:"required"`
	OwnerID        string                                                    `json:"ownerId" api:"required"`
	PackageName    string                                                    `json:"packageName" api:"required"`
	SecretPath     string                                                    `json:"secretPath" api:"required"`
	// Deprecated: use createdBy (same value — the creating actor). Null for
	// credentials created before rollout.
	//
	// Deprecated: deprecated
	UserID string                                                      `json:"userId" api:"required"`
	Label  string                                                      `json:"label"`
	Link   CredentialPackageCredentialFieldDeleteResponseDataLinkUnion `json:"link"`
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
func (r CredentialPackageCredentialFieldDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldDeleteResponseDataField struct {
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
func (r CredentialPackageCredentialFieldDeleteResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldDeleteResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageCredentialFieldDeleteResponseDataLinkUnion contains all
// possible properties and values from
// [CredentialPackageCredentialFieldDeleteResponseDataLinkObject],
// [CredentialPackageCredentialFieldDeleteResponseDataLinkObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CredentialPackageCredentialFieldDeleteResponseDataLinkUnion struct {
	// This field is a union of
	// [[]CredentialPackageCredentialFieldDeleteResponseDataLinkObjectFieldMap],
	// [[]CredentialPackageCredentialFieldDeleteResponseDataLinkObject2FieldMap]
	FieldMap CredentialPackageCredentialFieldDeleteResponseDataLinkUnionFieldMap `json:"fieldMap"`
	// This field is from variant
	// [CredentialPackageCredentialFieldDeleteResponseDataLinkObject].
	ItemID string `json:"itemId"`
	// This field is from variant
	// [CredentialPackageCredentialFieldDeleteResponseDataLinkObject].
	ItemTitle string `json:"itemTitle"`
	Provider  string `json:"provider"`
	// This field is from variant
	// [CredentialPackageCredentialFieldDeleteResponseDataLinkObject].
	VaultID string `json:"vaultId"`
	// This field is from variant
	// [CredentialPackageCredentialFieldDeleteResponseDataLinkObject].
	VaultName string `json:"vaultName"`
	// This field is from variant
	// [CredentialPackageCredentialFieldDeleteResponseDataLinkObject2].
	CardID string `json:"cardId"`
	// This field is from variant
	// [CredentialPackageCredentialFieldDeleteResponseDataLinkObject2].
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

func (u CredentialPackageCredentialFieldDeleteResponseDataLinkUnion) AsCredentialPackageCredentialFieldDeleteResponseDataLinkObject() (v CredentialPackageCredentialFieldDeleteResponseDataLinkObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialPackageCredentialFieldDeleteResponseDataLinkUnion) AsCredentialPackageCredentialFieldDeleteResponseDataLinkObject2() (v CredentialPackageCredentialFieldDeleteResponseDataLinkObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CredentialPackageCredentialFieldDeleteResponseDataLinkUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CredentialPackageCredentialFieldDeleteResponseDataLinkUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageCredentialFieldDeleteResponseDataLinkUnionFieldMap is an
// implicit subunion of
// [CredentialPackageCredentialFieldDeleteResponseDataLinkUnion].
// CredentialPackageCredentialFieldDeleteResponseDataLinkUnionFieldMap provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [CredentialPackageCredentialFieldDeleteResponseDataLinkUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfCredentialPackageCredentialFieldDeleteResponseDataLinkObjectFieldMapArray
// OfCredentialPackageCredentialFieldDeleteResponseDataLinkObject2FieldMapArray]
type CredentialPackageCredentialFieldDeleteResponseDataLinkUnionFieldMap struct {
	// This field will be present if the value is a
	// [[]CredentialPackageCredentialFieldDeleteResponseDataLinkObjectFieldMap] instead
	// of an object.
	OfCredentialPackageCredentialFieldDeleteResponseDataLinkObjectFieldMapArray []CredentialPackageCredentialFieldDeleteResponseDataLinkObjectFieldMap `json:",inline"`
	// This field will be present if the value is a
	// [[]CredentialPackageCredentialFieldDeleteResponseDataLinkObject2FieldMap]
	// instead of an object.
	OfCredentialPackageCredentialFieldDeleteResponseDataLinkObject2FieldMapArray []CredentialPackageCredentialFieldDeleteResponseDataLinkObject2FieldMap `json:",inline"`
	JSON                                                                         struct {
		OfCredentialPackageCredentialFieldDeleteResponseDataLinkObjectFieldMapArray  respjson.Field
		OfCredentialPackageCredentialFieldDeleteResponseDataLinkObject2FieldMapArray respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (r *CredentialPackageCredentialFieldDeleteResponseDataLinkUnionFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldDeleteResponseDataLinkObject struct {
	FieldMap  []CredentialPackageCredentialFieldDeleteResponseDataLinkObjectFieldMap `json:"fieldMap" api:"required"`
	ItemID    string                                                                 `json:"itemId" api:"required"`
	ItemTitle string                                                                 `json:"itemTitle" api:"required"`
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
func (r CredentialPackageCredentialFieldDeleteResponseDataLinkObject) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldDeleteResponseDataLinkObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldDeleteResponseDataLinkObjectFieldMap struct {
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
func (r CredentialPackageCredentialFieldDeleteResponseDataLinkObjectFieldMap) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldDeleteResponseDataLinkObjectFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldDeleteResponseDataLinkObject2 struct {
	CardID   string                                                                  `json:"cardId" api:"required" format:"uuid"`
	FieldMap []CredentialPackageCredentialFieldDeleteResponseDataLinkObject2FieldMap `json:"fieldMap" api:"required"`
	Label    string                                                                  `json:"label" api:"required"`
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
func (r CredentialPackageCredentialFieldDeleteResponseDataLinkObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldDeleteResponseDataLinkObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldDeleteResponseDataLinkObject2FieldMap struct {
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
func (r CredentialPackageCredentialFieldDeleteResponseDataLinkObject2FieldMap) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialPackageCredentialFieldDeleteResponseDataLinkObject2FieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldNewParams struct {
	PackageName string `path:"packageName" api:"required" json:"-"`
	// Any of "email", "username", "password", "api_token", "phone_number",
	// "two_factor_secret", "backup_codes", "cc_number", "cc_exp", "cc_csc", "cc_name".
	FieldType CredentialPackageCredentialFieldNewParamsFieldType `json:"fieldType,omitzero" api:"required"`
	Value     string                                             `json:"value" api:"required"`
	paramObj
}

func (r CredentialPackageCredentialFieldNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CredentialPackageCredentialFieldNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialPackageCredentialFieldNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldNewParamsFieldType string

const (
	CredentialPackageCredentialFieldNewParamsFieldTypeEmail           CredentialPackageCredentialFieldNewParamsFieldType = "email"
	CredentialPackageCredentialFieldNewParamsFieldTypeUsername        CredentialPackageCredentialFieldNewParamsFieldType = "username"
	CredentialPackageCredentialFieldNewParamsFieldTypePassword        CredentialPackageCredentialFieldNewParamsFieldType = "password"
	CredentialPackageCredentialFieldNewParamsFieldTypeAPIToken        CredentialPackageCredentialFieldNewParamsFieldType = "api_token"
	CredentialPackageCredentialFieldNewParamsFieldTypePhoneNumber     CredentialPackageCredentialFieldNewParamsFieldType = "phone_number"
	CredentialPackageCredentialFieldNewParamsFieldTypeTwoFactorSecret CredentialPackageCredentialFieldNewParamsFieldType = "two_factor_secret"
	CredentialPackageCredentialFieldNewParamsFieldTypeBackupCodes     CredentialPackageCredentialFieldNewParamsFieldType = "backup_codes"
	CredentialPackageCredentialFieldNewParamsFieldTypeCcNumber        CredentialPackageCredentialFieldNewParamsFieldType = "cc_number"
	CredentialPackageCredentialFieldNewParamsFieldTypeCcExp           CredentialPackageCredentialFieldNewParamsFieldType = "cc_exp"
	CredentialPackageCredentialFieldNewParamsFieldTypeCcCsc           CredentialPackageCredentialFieldNewParamsFieldType = "cc_csc"
	CredentialPackageCredentialFieldNewParamsFieldTypeCcName          CredentialPackageCredentialFieldNewParamsFieldType = "cc_name"
)

type CredentialPackageCredentialFieldUpdateParams struct {
	PackageName    string `path:"packageName" api:"required" json:"-"`
	CredentialName string `path:"credentialName" api:"required" json:"-"`
	Value          string `json:"value" api:"required"`
	paramObj
}

func (r CredentialPackageCredentialFieldUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CredentialPackageCredentialFieldUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialPackageCredentialFieldUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldUpdateParamsFieldType string

const (
	CredentialPackageCredentialFieldUpdateParamsFieldTypeEmail           CredentialPackageCredentialFieldUpdateParamsFieldType = "email"
	CredentialPackageCredentialFieldUpdateParamsFieldTypeUsername        CredentialPackageCredentialFieldUpdateParamsFieldType = "username"
	CredentialPackageCredentialFieldUpdateParamsFieldTypePassword        CredentialPackageCredentialFieldUpdateParamsFieldType = "password"
	CredentialPackageCredentialFieldUpdateParamsFieldTypeAPIToken        CredentialPackageCredentialFieldUpdateParamsFieldType = "api_token"
	CredentialPackageCredentialFieldUpdateParamsFieldTypePhoneNumber     CredentialPackageCredentialFieldUpdateParamsFieldType = "phone_number"
	CredentialPackageCredentialFieldUpdateParamsFieldTypeTwoFactorSecret CredentialPackageCredentialFieldUpdateParamsFieldType = "two_factor_secret"
	CredentialPackageCredentialFieldUpdateParamsFieldTypeBackupCodes     CredentialPackageCredentialFieldUpdateParamsFieldType = "backup_codes"
	CredentialPackageCredentialFieldUpdateParamsFieldTypeCcNumber        CredentialPackageCredentialFieldUpdateParamsFieldType = "cc_number"
	CredentialPackageCredentialFieldUpdateParamsFieldTypeCcExp           CredentialPackageCredentialFieldUpdateParamsFieldType = "cc_exp"
	CredentialPackageCredentialFieldUpdateParamsFieldTypeCcCsc           CredentialPackageCredentialFieldUpdateParamsFieldType = "cc_csc"
	CredentialPackageCredentialFieldUpdateParamsFieldTypeCcName          CredentialPackageCredentialFieldUpdateParamsFieldType = "cc_name"
)

type CredentialPackageCredentialFieldDeleteParams struct {
	PackageName    string `path:"packageName" api:"required" json:"-"`
	CredentialName string `path:"credentialName" api:"required" json:"-"`
	paramObj
}

type CredentialPackageCredentialFieldDeleteParamsFieldType string

const (
	CredentialPackageCredentialFieldDeleteParamsFieldTypeEmail           CredentialPackageCredentialFieldDeleteParamsFieldType = "email"
	CredentialPackageCredentialFieldDeleteParamsFieldTypeUsername        CredentialPackageCredentialFieldDeleteParamsFieldType = "username"
	CredentialPackageCredentialFieldDeleteParamsFieldTypePassword        CredentialPackageCredentialFieldDeleteParamsFieldType = "password"
	CredentialPackageCredentialFieldDeleteParamsFieldTypeAPIToken        CredentialPackageCredentialFieldDeleteParamsFieldType = "api_token"
	CredentialPackageCredentialFieldDeleteParamsFieldTypePhoneNumber     CredentialPackageCredentialFieldDeleteParamsFieldType = "phone_number"
	CredentialPackageCredentialFieldDeleteParamsFieldTypeTwoFactorSecret CredentialPackageCredentialFieldDeleteParamsFieldType = "two_factor_secret"
	CredentialPackageCredentialFieldDeleteParamsFieldTypeBackupCodes     CredentialPackageCredentialFieldDeleteParamsFieldType = "backup_codes"
	CredentialPackageCredentialFieldDeleteParamsFieldTypeCcNumber        CredentialPackageCredentialFieldDeleteParamsFieldType = "cc_number"
	CredentialPackageCredentialFieldDeleteParamsFieldTypeCcExp           CredentialPackageCredentialFieldDeleteParamsFieldType = "cc_exp"
	CredentialPackageCredentialFieldDeleteParamsFieldTypeCcCsc           CredentialPackageCredentialFieldDeleteParamsFieldType = "cc_csc"
	CredentialPackageCredentialFieldDeleteParamsFieldTypeCcName          CredentialPackageCredentialFieldDeleteParamsFieldType = "cc_name"
)
