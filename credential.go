// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

// CredentialService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialService] method instead.
type CredentialService struct {
	options     []option.RequestOption
	Packages    CredentialPackageService
	Onepassword CredentialOnepasswordService
}

// NewCredentialService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCredentialService(opts ...option.RequestOption) (r CredentialService) {
	r = CredentialService{}
	r.options = opts
	r.Packages = NewCredentialPackageService(opts...)
	r.Onepassword = NewCredentialOnepasswordService(opts...)
	return
}

// Returns a paginated list of all credentials belonging to the authenticated user
// across every package. Accepts standard pagination query parameters and responds
// with the credential items plus pagination metadata.
func (r *CredentialService) List(ctx context.Context, query CredentialListParams, opts ...option.RequestOption) (res *CredentialListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credentials"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type CredentialListResponse struct {
	Items      []CredentialListResponseItem `json:"items" api:"required"`
	Pagination shared.Pagination            `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialListResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialListResponseItem struct {
	CreatedBy      string                            `json:"createdBy" api:"required"`
	CredentialName string                            `json:"credentialName" api:"required"`
	Fields         []CredentialListResponseItemField `json:"fields" api:"required"`
	OwnerID        string                            `json:"ownerId" api:"required"`
	PackageName    string                            `json:"packageName" api:"required"`
	SecretPath     string                            `json:"secretPath" api:"required"`
	// Deprecated: use createdBy (same value — the creating actor). Null for
	// credentials created before rollout.
	//
	// Deprecated: deprecated
	UserID string                              `json:"userId" api:"required"`
	Label  string                              `json:"label"`
	Link   CredentialListResponseItemLinkUnion `json:"link"`
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
func (r CredentialListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *CredentialListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialListResponseItemField struct {
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
func (r CredentialListResponseItemField) RawJSON() string { return r.JSON.raw }
func (r *CredentialListResponseItemField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialListResponseItemLinkUnion contains all possible properties and values
// from [CredentialListResponseItemLinkObject],
// [CredentialListResponseItemLinkObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CredentialListResponseItemLinkUnion struct {
	// This field is a union of [[]CredentialListResponseItemLinkObjectFieldMap],
	// [[]CredentialListResponseItemLinkObject2FieldMap]
	FieldMap CredentialListResponseItemLinkUnionFieldMap `json:"fieldMap"`
	// This field is from variant [CredentialListResponseItemLinkObject].
	ItemID string `json:"itemId"`
	// This field is from variant [CredentialListResponseItemLinkObject].
	ItemTitle string `json:"itemTitle"`
	Provider  string `json:"provider"`
	// This field is from variant [CredentialListResponseItemLinkObject].
	VaultID string `json:"vaultId"`
	// This field is from variant [CredentialListResponseItemLinkObject].
	VaultName string `json:"vaultName"`
	// This field is from variant [CredentialListResponseItemLinkObject2].
	CardID string `json:"cardId"`
	// This field is from variant [CredentialListResponseItemLinkObject2].
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

func (u CredentialListResponseItemLinkUnion) AsCredentialListResponseItemLinkObject() (v CredentialListResponseItemLinkObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialListResponseItemLinkUnion) AsCredentialListResponseItemLinkObject2() (v CredentialListResponseItemLinkObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CredentialListResponseItemLinkUnion) RawJSON() string { return u.JSON.raw }

func (r *CredentialListResponseItemLinkUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialListResponseItemLinkUnionFieldMap is an implicit subunion of
// [CredentialListResponseItemLinkUnion].
// CredentialListResponseItemLinkUnionFieldMap provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [CredentialListResponseItemLinkUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCredentialListResponseItemLinkObjectFieldMapArray
// OfCredentialListResponseItemLinkObject2FieldMapArray]
type CredentialListResponseItemLinkUnionFieldMap struct {
	// This field will be present if the value is a
	// [[]CredentialListResponseItemLinkObjectFieldMap] instead of an object.
	OfCredentialListResponseItemLinkObjectFieldMapArray []CredentialListResponseItemLinkObjectFieldMap `json:",inline"`
	// This field will be present if the value is a
	// [[]CredentialListResponseItemLinkObject2FieldMap] instead of an object.
	OfCredentialListResponseItemLinkObject2FieldMapArray []CredentialListResponseItemLinkObject2FieldMap `json:",inline"`
	JSON                                                 struct {
		OfCredentialListResponseItemLinkObjectFieldMapArray  respjson.Field
		OfCredentialListResponseItemLinkObject2FieldMapArray respjson.Field
		raw                                                  string
	} `json:"-"`
}

func (r *CredentialListResponseItemLinkUnionFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialListResponseItemLinkObject struct {
	FieldMap  []CredentialListResponseItemLinkObjectFieldMap `json:"fieldMap" api:"required"`
	ItemID    string                                         `json:"itemId" api:"required"`
	ItemTitle string                                         `json:"itemTitle" api:"required"`
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
func (r CredentialListResponseItemLinkObject) RawJSON() string { return r.JSON.raw }
func (r *CredentialListResponseItemLinkObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialListResponseItemLinkObjectFieldMap struct {
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
func (r CredentialListResponseItemLinkObjectFieldMap) RawJSON() string { return r.JSON.raw }
func (r *CredentialListResponseItemLinkObjectFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialListResponseItemLinkObject2 struct {
	CardID   string                                          `json:"cardId" api:"required" format:"uuid"`
	FieldMap []CredentialListResponseItemLinkObject2FieldMap `json:"fieldMap" api:"required"`
	Label    string                                          `json:"label" api:"required"`
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
func (r CredentialListResponseItemLinkObject2) RawJSON() string { return r.JSON.raw }
func (r *CredentialListResponseItemLinkObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialListResponseItemLinkObject2FieldMap struct {
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
func (r CredentialListResponseItemLinkObject2FieldMap) RawJSON() string { return r.JSON.raw }
func (r *CredentialListResponseItemLinkObject2FieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CredentialListParams]'s query parameters as `url.Values`.
func (r CredentialListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
