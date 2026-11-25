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

// Sets or clears the optional `label` (human-readable name, free text, max 200
// characters) of an existing credential. `null` or an empty text clears it. Works
// for linked credentials too: the label is metadata of the credential, not a field
// of the linked item. Returns the updated credential.
func (r *CredentialPackageCredentialService) Update(ctx context.Context, credentialName string, params CredentialPackageCredentialUpdateParams, opts ...option.RequestOption) (res *CredentialPackageCredentialUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.PackageName == "" {
		err = errors.New("missing required packageName parameter")
		return nil, err
	}
	if credentialName == "" {
		err = errors.New("missing required credentialName parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/packages/%s/credentials/%s", url.PathEscape(params.PackageName), url.PathEscape(credentialName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
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
	Data    Credential `json:"data" api:"required"`
	Message string     `json:"message" api:"required"`
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

type CredentialPackageCredentialUpdateResponse struct {
	Data    Credential `json:"data" api:"required"`
	Message string     `json:"message" api:"required"`
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
func (r CredentialPackageCredentialUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialDeleteResponse struct {
	Data    Credential `json:"data" api:"required"`
	Message string     `json:"message" api:"required"`
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

type CredentialPackageCredentialUpdateParams struct {
	Label       param.Opt[string] `json:"label,omitzero" api:"required"`
	PackageName string            `path:"packageName" api:"required" json:"-"`
	paramObj
}

func (r CredentialPackageCredentialUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CredentialPackageCredentialUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CredentialPackageCredentialUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialDeleteParams struct {
	PackageName string `path:"packageName" api:"required" json:"-"`
	paramObj
}
