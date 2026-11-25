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
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/pagination"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
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
func (r *CredentialService) List(ctx context.Context, query CredentialListParams, opts ...option.RequestOption) (res *pagination.Page[Credential], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "credentials"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a paginated list of all credentials belonging to the authenticated user
// across every package. Accepts standard pagination query parameters and responds
// with the credential items plus pagination metadata.
func (r *CredentialService) ListAutoPaging(ctx context.Context, query CredentialListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Credential] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type Credential struct {
	CreatedBy      string            `json:"createdBy" api:"required"`
	CredentialName string            `json:"credentialName" api:"required"`
	Fields         []CredentialField `json:"fields" api:"required"`
	OwnerID        string            `json:"ownerId" api:"required"`
	PackageName    string            `json:"packageName" api:"required"`
	SecretPath     string            `json:"secretPath" api:"required"`
	// Deprecated: use createdBy (same value — the creating actor). Null for
	// credentials created before rollout.
	//
	// Deprecated: deprecated
	UserID string              `json:"userId" api:"required"`
	Label  string              `json:"label"`
	Link   CredentialLinkUnion `json:"link"`
	// Any of "mobilerun", "onepassword", "wallet".
	Source CredentialSource `json:"source"`
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
func (r Credential) RawJSON() string { return r.JSON.raw }
func (r *Credential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialField struct {
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
func (r CredentialField) RawJSON() string { return r.JSON.raw }
func (r *CredentialField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialLinkUnion contains all possible properties and values from
// [CredentialLinkObject], [CredentialLinkObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CredentialLinkUnion struct {
	// This field is a union of [[]CredentialLinkObjectFieldMap],
	// [[]CredentialLinkObject2FieldMap]
	FieldMap CredentialLinkUnionFieldMap `json:"fieldMap"`
	// This field is from variant [CredentialLinkObject].
	ItemID string `json:"itemId"`
	// This field is from variant [CredentialLinkObject].
	ItemTitle string `json:"itemTitle"`
	Provider  string `json:"provider"`
	// This field is from variant [CredentialLinkObject].
	VaultID string `json:"vaultId"`
	// This field is from variant [CredentialLinkObject].
	VaultName string `json:"vaultName"`
	// This field is from variant [CredentialLinkObject2].
	CardID string `json:"cardId"`
	// This field is from variant [CredentialLinkObject2].
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

func (u CredentialLinkUnion) AsCredentialLinkObject() (v CredentialLinkObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialLinkUnion) AsCredentialLinkObject2() (v CredentialLinkObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CredentialLinkUnion) RawJSON() string { return u.JSON.raw }

func (r *CredentialLinkUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialLinkUnionFieldMap is an implicit subunion of [CredentialLinkUnion].
// CredentialLinkUnionFieldMap provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [CredentialLinkUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCredentialLinkObjectFieldMapArray
// OfCredentialLinkObject2FieldMapArray]
type CredentialLinkUnionFieldMap struct {
	// This field will be present if the value is a [[]CredentialLinkObjectFieldMap]
	// instead of an object.
	OfCredentialLinkObjectFieldMapArray []CredentialLinkObjectFieldMap `json:",inline"`
	// This field will be present if the value is a [[]CredentialLinkObject2FieldMap]
	// instead of an object.
	OfCredentialLinkObject2FieldMapArray []CredentialLinkObject2FieldMap `json:",inline"`
	JSON                                 struct {
		OfCredentialLinkObjectFieldMapArray  respjson.Field
		OfCredentialLinkObject2FieldMapArray respjson.Field
		raw                                  string
	} `json:"-"`
}

func (r *CredentialLinkUnionFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialLinkObject struct {
	FieldMap  []CredentialLinkObjectFieldMap `json:"fieldMap" api:"required"`
	ItemID    string                         `json:"itemId" api:"required"`
	ItemTitle string                         `json:"itemTitle" api:"required"`
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
func (r CredentialLinkObject) RawJSON() string { return r.JSON.raw }
func (r *CredentialLinkObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialLinkObjectFieldMap struct {
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
func (r CredentialLinkObjectFieldMap) RawJSON() string { return r.JSON.raw }
func (r *CredentialLinkObjectFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialLinkObject2 struct {
	CardID   string                          `json:"cardId" api:"required" format:"uuid"`
	FieldMap []CredentialLinkObject2FieldMap `json:"fieldMap" api:"required"`
	Label    string                          `json:"label" api:"required"`
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
func (r CredentialLinkObject2) RawJSON() string { return r.JSON.raw }
func (r *CredentialLinkObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialLinkObject2FieldMap struct {
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
func (r CredentialLinkObject2FieldMap) RawJSON() string { return r.JSON.raw }
func (r *CredentialLinkObject2FieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialSource string

const (
	CredentialSourceMobilerun   CredentialSource = "mobilerun"
	CredentialSourceOnepassword CredentialSource = "onepassword"
	CredentialSourceWallet      CredentialSource = "wallet"
)

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
