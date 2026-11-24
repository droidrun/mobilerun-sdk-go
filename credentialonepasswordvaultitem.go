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

// CredentialOnepasswordVaultItemService contains methods and other services that
// help with interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialOnepasswordVaultItemService] method instead.
type CredentialOnepasswordVaultItemService struct {
	options []option.RequestOption
	Fields  CredentialOnepasswordVaultItemFieldService
}

// NewCredentialOnepasswordVaultItemService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCredentialOnepasswordVaultItemService(opts ...option.RequestOption) (r CredentialOnepasswordVaultItemService) {
	r = CredentialOnepasswordVaultItemService{}
	r.options = opts
	r.Fields = NewCredentialOnepasswordVaultItemFieldService(opts...)
	return
}

// List 1Password item metadata
func (r *CredentialOnepasswordVaultItemService) List(ctx context.Context, vaultID string, opts ...option.RequestOption) (res *CredentialOnepasswordVaultItemListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if vaultID == "" {
		err = errors.New("missing required vaultId parameter")
		return nil, err
	}
	path := fmt.Sprintf("credentials/onepassword/vaults/%s/items", url.PathEscape(vaultID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CredentialOnepasswordVaultItemListResponse struct {
	Data []CredentialOnepasswordVaultItemListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialOnepasswordVaultItemListResponse) RawJSON() string { return r.JSON.raw }
func (r *CredentialOnepasswordVaultItemListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialOnepasswordVaultItemListResponseData struct {
	ID       string                                                  `json:"id" api:"required"`
	Category string                                                  `json:"category" api:"required"`
	Title    string                                                  `json:"title" api:"required"`
	Websites []CredentialOnepasswordVaultItemListResponseDataWebsite `json:"websites" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Category    respjson.Field
		Title       respjson.Field
		Websites    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialOnepasswordVaultItemListResponseData) RawJSON() string { return r.JSON.raw }
func (r *CredentialOnepasswordVaultItemListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialOnepasswordVaultItemListResponseDataWebsite struct {
	Label string `json:"label" api:"required"`
	URL   string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialOnepasswordVaultItemListResponseDataWebsite) RawJSON() string { return r.JSON.raw }
func (r *CredentialOnepasswordVaultItemListResponseDataWebsite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
