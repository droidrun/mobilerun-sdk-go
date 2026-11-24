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
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// CredentialOnepasswordVaultItemFieldService contains methods and other services
// that help with interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialOnepasswordVaultItemFieldService] method instead.
type CredentialOnepasswordVaultItemFieldService struct {
	options []option.RequestOption
}

// NewCredentialOnepasswordVaultItemFieldService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewCredentialOnepasswordVaultItemFieldService(opts ...option.RequestOption) (r CredentialOnepasswordVaultItemFieldService) {
	r = CredentialOnepasswordVaultItemFieldService{}
	r.options = opts
	return
}

// List 1Password field metadata
func (r *CredentialOnepasswordVaultItemFieldService) List(ctx context.Context, itemID string, query CredentialOnepasswordVaultItemFieldListParams, opts ...option.RequestOption) (res *CredentialOnepasswordVaultItemFieldListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.VaultID == "" {
		err = errors.New("missing required vaultId parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required itemId parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/onepassword/vaults/%s/items/%s/fields", url.PathEscape(query.VaultID), url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CredentialOnepasswordVaultItemFieldListResponse struct {
	Data []CredentialOnepasswordVaultItemFieldListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialOnepasswordVaultItemFieldListResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialOnepasswordVaultItemFieldListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialOnepasswordVaultItemFieldListResponseData struct {
	ID        string `json:"id" api:"required"`
	FieldType string `json:"fieldType" api:"required"`
	Title     string `json:"title" api:"required"`
	SectionID string `json:"sectionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FieldType   respjson.Field
		Title       respjson.Field
		SectionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialOnepasswordVaultItemFieldListResponseData) RawJSON() string { return r.JSON.raw }
func (r *CredentialOnepasswordVaultItemFieldListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialOnepasswordVaultItemFieldListParams struct {
	VaultID string `path:"vaultId" api:"required" json:"-"`
	paramObj
}
