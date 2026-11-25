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
	Data []Credential `json:"data" api:"required"`
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
