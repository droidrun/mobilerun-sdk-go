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

// CredentialPackageService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialPackageService] method instead.
type CredentialPackageService struct {
	options     []option.RequestOption
	Credentials CredentialPackageCredentialService
	Links       CredentialPackageLinkService
}

// NewCredentialPackageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCredentialPackageService(opts ...option.RequestOption) (r CredentialPackageService) {
	r = CredentialPackageService{}
	r.options = opts
	r.Credentials = NewCredentialPackageCredentialService(opts...)
	r.Links = NewCredentialPackageLinkService(opts...)
	return
}

// Creates a new package (identified by `packageName`) under which credentials can
// be grouped. Returns a conflict if a package with the same name already exists
// for the user.
func (r *CredentialPackageService) New(ctx context.Context, body CredentialPackageNewParams, opts ...option.RequestOption) (res *CredentialPackageNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credentials/packages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns all credentials stored under the given `packageName`. Each credential
// includes its name, secret path, and the list of fields it holds.
func (r *CredentialPackageService) List(ctx context.Context, packageName string, opts ...option.RequestOption) (res *CredentialPackageListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/packages/%s", url.PathEscape(packageName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the names of all packages (apps) the authenticated owner has credentials
// grouped under. Use this to discover which `packageName` values are valid for the
// per-package credential routes.
func (r *CredentialPackageService) ListAll(ctx context.Context, opts ...option.RequestOption) (res *CredentialPackageListAllResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credentials/packages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CredentialPackageNewResponse struct {
	Data    CredentialPackageNewResponseData `json:"data" api:"required"`
	Message string                           `json:"message" api:"required"`
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
func (r CredentialPackageNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageNewResponseData struct {
	PackageName string `json:"packageName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PackageName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageListResponse struct {
	Data []CredentialPackageListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageListResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageListResponseData struct {
	CreatedBy      string                                   `json:"createdBy" api:"required"`
	CredentialName string                                   `json:"credentialName" api:"required"`
	Fields         []CredentialPackageListResponseDataField `json:"fields" api:"required"`
	OwnerID        string                                   `json:"ownerId" api:"required"`
	PackageName    string                                   `json:"packageName" api:"required"`
	SecretPath     string                                   `json:"secretPath" api:"required"`
	// Deprecated: use createdBy (same value — the creating actor). Null for
	// credentials created before rollout.
	//
	// Deprecated: deprecated
	UserID string                                     `json:"userId" api:"required"`
	Label  string                                     `json:"label"`
	Link   CredentialPackageListResponseDataLinkUnion `json:"link"`
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
func (r CredentialPackageListResponseData) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageListResponseDataField struct {
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
func (r CredentialPackageListResponseDataField) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageListResponseDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageListResponseDataLinkUnion contains all possible properties and
// values from [CredentialPackageListResponseDataLinkObject],
// [CredentialPackageListResponseDataLinkObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CredentialPackageListResponseDataLinkUnion struct {
	// This field is a union of
	// [[]CredentialPackageListResponseDataLinkObjectFieldMap],
	// [[]CredentialPackageListResponseDataLinkObject2FieldMap]
	FieldMap CredentialPackageListResponseDataLinkUnionFieldMap `json:"fieldMap"`
	// This field is from variant [CredentialPackageListResponseDataLinkObject].
	ItemID string `json:"itemId"`
	// This field is from variant [CredentialPackageListResponseDataLinkObject].
	ItemTitle string `json:"itemTitle"`
	Provider  string `json:"provider"`
	// This field is from variant [CredentialPackageListResponseDataLinkObject].
	VaultID string `json:"vaultId"`
	// This field is from variant [CredentialPackageListResponseDataLinkObject].
	VaultName string `json:"vaultName"`
	// This field is from variant [CredentialPackageListResponseDataLinkObject2].
	CardID string `json:"cardId"`
	// This field is from variant [CredentialPackageListResponseDataLinkObject2].
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

func (u CredentialPackageListResponseDataLinkUnion) AsCredentialPackageListResponseDataLinkObject() (v CredentialPackageListResponseDataLinkObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CredentialPackageListResponseDataLinkUnion) AsCredentialPackageListResponseDataLinkObject2() (v CredentialPackageListResponseDataLinkObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CredentialPackageListResponseDataLinkUnion) RawJSON() string { return u.JSON.raw }

func (r *CredentialPackageListResponseDataLinkUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialPackageListResponseDataLinkUnionFieldMap is an implicit subunion of
// [CredentialPackageListResponseDataLinkUnion].
// CredentialPackageListResponseDataLinkUnionFieldMap provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [CredentialPackageListResponseDataLinkUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCredentialPackageListResponseDataLinkObjectFieldMapArray
// OfCredentialPackageListResponseDataLinkObject2FieldMapArray]
type CredentialPackageListResponseDataLinkUnionFieldMap struct {
	// This field will be present if the value is a
	// [[]CredentialPackageListResponseDataLinkObjectFieldMap] instead of an object.
	OfCredentialPackageListResponseDataLinkObjectFieldMapArray []CredentialPackageListResponseDataLinkObjectFieldMap `json:",inline"`
	// This field will be present if the value is a
	// [[]CredentialPackageListResponseDataLinkObject2FieldMap] instead of an object.
	OfCredentialPackageListResponseDataLinkObject2FieldMapArray []CredentialPackageListResponseDataLinkObject2FieldMap `json:",inline"`
	JSON                                                        struct {
		OfCredentialPackageListResponseDataLinkObjectFieldMapArray  respjson.Field
		OfCredentialPackageListResponseDataLinkObject2FieldMapArray respjson.Field
		raw                                                         string
	} `json:"-"`
}

func (r *CredentialPackageListResponseDataLinkUnionFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageListResponseDataLinkObject struct {
	FieldMap  []CredentialPackageListResponseDataLinkObjectFieldMap `json:"fieldMap" api:"required"`
	ItemID    string                                                `json:"itemId" api:"required"`
	ItemTitle string                                                `json:"itemTitle" api:"required"`
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
func (r CredentialPackageListResponseDataLinkObject) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageListResponseDataLinkObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageListResponseDataLinkObjectFieldMap struct {
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
func (r CredentialPackageListResponseDataLinkObjectFieldMap) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageListResponseDataLinkObjectFieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageListResponseDataLinkObject2 struct {
	CardID   string                                                 `json:"cardId" api:"required" format:"uuid"`
	FieldMap []CredentialPackageListResponseDataLinkObject2FieldMap `json:"fieldMap" api:"required"`
	Label    string                                                 `json:"label" api:"required"`
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
func (r CredentialPackageListResponseDataLinkObject2) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageListResponseDataLinkObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageListResponseDataLinkObject2FieldMap struct {
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
func (r CredentialPackageListResponseDataLinkObject2FieldMap) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageListResponseDataLinkObject2FieldMap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageListAllResponse struct {
	Data []CredentialPackageListAllResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageListAllResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageListAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageListAllResponseData struct {
	PackageName string `json:"packageName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PackageName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialPackageListAllResponseData) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageListAllResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageNewParams struct {
	PackageName string `json:"packageName" api:"required"`
	paramObj
}

func (r CredentialPackageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CredentialPackageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialPackageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
