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
func (r CredentialPackageCredentialFieldNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldUpdateResponse struct {
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
func (r CredentialPackageCredentialFieldUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialPackageCredentialFieldDeleteResponse struct {
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
func (r CredentialPackageCredentialFieldDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialPackageCredentialFieldDeleteResponse) UnmarshalJSON(data []byte) error {
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
