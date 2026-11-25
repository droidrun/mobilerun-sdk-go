// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/pagination"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// StoreAppService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreAppService] method instead.
type StoreAppService struct {
	options []option.RequestOption
}

// NewStoreAppService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStoreAppService(opts ...option.RequestOption) (r StoreAppService) {
	r = StoreAppService{}
	r.options = opts
	return
}

// Paginated list of published store listings, ordered featured desc, sortOrder
// asc, displayName asc.
func (r *StoreAppService) List(ctx context.Context, query StoreAppListParams, opts ...option.RequestOption) (res *pagination.Page[PublicStoreApp], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "store/apps"
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

// Paginated list of published store listings, ordered featured desc, sortOrder
// asc, displayName asc.
func (r *StoreAppService) ListAutoPaging(ctx context.Context, query StoreAppListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicStoreApp] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type PublicStoreApp struct {
	ID            string    `json:"id" api:"required" format:"uuid"`
	AppStoreID    int64     `json:"appStoreId" api:"required"`
	BundleID      string    `json:"bundleId" api:"required"`
	Category      string    `json:"category" api:"required"`
	CreatedAt     time.Time `json:"createdAt" api:"required" format:"date-time"`
	Description   string    `json:"description" api:"required"`
	DeveloperName string    `json:"developerName" api:"required"`
	DisplayName   string    `json:"displayName" api:"required"`
	Featured      bool      `json:"featured" api:"required"`
	IconURL       string    `json:"iconURL" api:"required"`
	// Any of "android", "ios".
	Platform  PublicStoreAppPlatform `json:"platform" api:"required"`
	UpdatedAt time.Time              `json:"updatedAt" api:"required" format:"date-time"`
	Version   PublicStoreAppVersion  `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AppStoreID    respjson.Field
		BundleID      respjson.Field
		Category      respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		DeveloperName respjson.Field
		DisplayName   respjson.Field
		Featured      respjson.Field
		IconURL       respjson.Field
		Platform      respjson.Field
		UpdatedAt     respjson.Field
		Version       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicStoreApp) RawJSON() string { return r.JSON.raw }
func (r *PublicStoreApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicStoreAppPlatform string

const (
	PublicStoreAppPlatformAndroid PublicStoreAppPlatform = "android"
	PublicStoreAppPlatformIos     PublicStoreAppPlatform = "ios"
)

type PublicStoreAppVersion struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PublicAppVersion
}

// Returns the unmodified JSON received from the API
func (r PublicStoreAppVersion) RawJSON() string { return r.JSON.raw }
func (r *PublicStoreAppVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoreAppListParams struct {
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	// Only listings of this app platform. All or omitted: every platform.
	//
	// Any of "all", "android", "ios".
	Platform StoreAppListParamsPlatform `query:"platform,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StoreAppListParams]'s query parameters as `url.Values`.
func (r StoreAppListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only listings of this app platform. All or omitted: every platform.
type StoreAppListParamsPlatform string

const (
	StoreAppListParamsPlatformAll     StoreAppListParamsPlatform = "all"
	StoreAppListParamsPlatformAndroid StoreAppListParamsPlatform = "android"
	StoreAppListParamsPlatformIos     StoreAppListParamsPlatform = "ios"
)
