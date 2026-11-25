// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
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

// AppService contains methods and other services that help with interacting with
// the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r AppService) {
	r = AppService{}
	r.options = opts
	return
}

// Retrieves an app by its ID
func (r *AppService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AppGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("apps/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves a paginated list of apps with filtering and search capabilities
func (r *AppService) List(ctx context.Context, query AppListParams, opts ...option.RequestOption) (res *pagination.Page[AppListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "apps"
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

// Retrieves a paginated list of apps with filtering and search capabilities
func (r *AppService) ListAutoPaging(ctx context.Context, query AppListParams, opts ...option.RequestOption) *pagination.PageAutoPager[AppListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an uploaded app by ID. Removes its files and metadata.
func (r *AppService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AppDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("apps/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Verifies the uploaded files and sets the app version status to available.
// Idempotent: replaying confirmation for an already-available version returns the
// same successful response without re-verifying the files.
func (r *AppService) ConfirmUpload(ctx context.Context, id string, opts ...option.RequestOption) (res *AppConfirmUploadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("apps/%s/confirm-upload", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Creates or updates an app and returns pre-signed upload URLs for each file.
// Honours an optional `Idempotency-Key` header: a retry with the same key and body
// returns the same app and version with freshly signed URLs and
// `Idempotent-Replayed: true`, and never creates a second version. An
// Idempotency-Key is honoured for 7 days; after that it behaves like a fresh
// request. A replay whose version was deleted is a 404; a replay whose version
// already left the queued state is a 409 `idempotency_key_consumed`; a replay that
// cannot get the version lock within 5 s is a 409
// `idempotency_request_in_progress`. Every response produced by a replay,
// including those 404/409 errors, carries `Idempotent-Replayed: true`. A same-key
// request still in flight after a 5 s wait is a 409
// `idempotency_request_in_progress` (retry after `Retry-After`). If signing fails,
// nothing is persisted (with or without a key) and the same request, key included,
// can be retried.
func (r *AppService) NewSignedUploadURL(ctx context.Context, body AppNewSignedUploadURLParams, opts ...option.RequestOption) (res *AppNewSignedUploadURLResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "apps/create-signed-upload-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves all versions of an app visible to the user (own uploads + system
// versions)
func (r *AppService) ListVersions(ctx context.Context, id string, opts ...option.RequestOption) (res *AppListVersionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("apps/%s/versions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Sets the app status to failed.
func (r *AppService) MarkFailed(ctx context.Context, id string, opts ...option.RequestOption) (res *AppMarkFailedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("apps/%s/mark-failed", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the user’s storage allowance in bytes: bytes used, bytes included in the
// plan, the hard maximum (included + max purchasable overage; null = unlimited)
// and whether usage above the included amount is billed as overage. With storage
// billing off there is no limit: only usedBytes is set.
func (r *AppService) StorageUsage(ctx context.Context, opts ...option.RequestOption) (res *AppStorageUsageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "apps/storage-usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type App struct {
	ID            string    `json:"id" api:"required" format:"uuid"`
	AppStoreID    int64     `json:"appStoreId" api:"required"`
	BundleID      string    `json:"bundleId" api:"required"`
	CreatedAt     time.Time `json:"createdAt" api:"required" format:"date-time"`
	Description   string    `json:"description" api:"required"`
	DeveloperName string    `json:"developerName" api:"required"`
	DisplayName   string    `json:"displayName" api:"required"`
	IconURL       string    `json:"iconURL" api:"required"`
	// Any of "android", "ios".
	Platform  AppPlatform `json:"platform" api:"required"`
	UpdatedAt time.Time   `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AppStoreID    respjson.Field
		BundleID      respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		DeveloperName respjson.Field
		DisplayName   respjson.Field
		IconURL       respjson.Field
		Platform      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r App) RawJSON() string { return r.JSON.raw }
func (r *App) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppPlatform string

const (
	AppPlatformAndroid AppPlatform = "android"
	AppPlatformIos     AppPlatform = "ios"
)

type PublicAppVersion struct {
	ID    string `json:"id" api:"required" format:"uuid"`
	AppID string `json:"appId" api:"required" format:"uuid"`
	// Any of "AF", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW",
	// "AP", "AU", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM",
	// "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF", "BI", "KH",
	// "CM", "CA", "CV", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG",
	// "CD", "CK", "CR", "HR", "CU", "CW", "CY", "CZ", "CI", "DK", "DJ", "DM", "DO",
	// "EC", "EG", "SV", "GQ", "ER", "EE", "ET", "FK", "FO", "FJ", "FI", "FR", "GF",
	// "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU",
	// "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN",
	// "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE",
	// "KI", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU",
	// "MO", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX",
	// "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR", "NP", "NL",
	// "AN", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "KP", "MK", "MP", "NO", "OM",
	// "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA",
	// "RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS", "SM",
	// "ST", "SA", "SN", "RS", "CS", "SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO",
	// "ZA", "GS", "SS", "ES", "LK", "SD", "SR", "SJ", "SZ", "SE", "CH", "SY", "TW",
	// "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV",
	// "UG", "UA", "AE", "GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG", "VI",
	// "WF", "EH", "YE", "ZM", "ZW", "AX".
	Country   PublicAppVersionCountry `json:"country" api:"required"`
	CreatedAt time.Time               `json:"createdAt" api:"required" format:"date-time"`
	CreatedBy string                  `json:"createdBy" api:"required"`
	OwnerID   string                  `json:"ownerId" api:"required"`
	QueuedAt  time.Time               `json:"queuedAt" api:"required" format:"date-time"`
	SizeBytes int64                   `json:"sizeBytes" api:"required"`
	// Any of "user", "system", "portal", "store".
	Source PublicAppVersionSource `json:"source" api:"required"`
	// Any of "queued", "available", "failed", "pruned".
	Status    PublicAppVersionStatus `json:"status" api:"required"`
	TargetSDK int64                  `json:"targetSdk" api:"required"`
	UpdatedAt time.Time              `json:"updatedAt" api:"required" format:"date-time"`
	// Deprecated: use ownerId (tenancy) / createdBy (actor).
	//
	// Deprecated: deprecated
	UserID      string `json:"userId" api:"required"`
	VersionCode int64  `json:"versionCode" api:"required"`
	VersionName string `json:"versionName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppID       respjson.Field
		Country     respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		OwnerID     respjson.Field
		QueuedAt    respjson.Field
		SizeBytes   respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		TargetSDK   respjson.Field
		UpdatedAt   respjson.Field
		UserID      respjson.Field
		VersionCode respjson.Field
		VersionName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAppVersion) RawJSON() string { return r.JSON.raw }
func (r *PublicAppVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAppVersionCountry string

const (
	PublicAppVersionCountryAf PublicAppVersionCountry = "AF"
	PublicAppVersionCountryAl PublicAppVersionCountry = "AL"
	PublicAppVersionCountryDz PublicAppVersionCountry = "DZ"
	PublicAppVersionCountryAs PublicAppVersionCountry = "AS"
	PublicAppVersionCountryAd PublicAppVersionCountry = "AD"
	PublicAppVersionCountryAo PublicAppVersionCountry = "AO"
	PublicAppVersionCountryAI PublicAppVersionCountry = "AI"
	PublicAppVersionCountryAq PublicAppVersionCountry = "AQ"
	PublicAppVersionCountryAg PublicAppVersionCountry = "AG"
	PublicAppVersionCountryAr PublicAppVersionCountry = "AR"
	PublicAppVersionCountryAm PublicAppVersionCountry = "AM"
	PublicAppVersionCountryAw PublicAppVersionCountry = "AW"
	PublicAppVersionCountryAp PublicAppVersionCountry = "AP"
	PublicAppVersionCountryAu PublicAppVersionCountry = "AU"
	PublicAppVersionCountryAt PublicAppVersionCountry = "AT"
	PublicAppVersionCountryAz PublicAppVersionCountry = "AZ"
	PublicAppVersionCountryBs PublicAppVersionCountry = "BS"
	PublicAppVersionCountryBh PublicAppVersionCountry = "BH"
	PublicAppVersionCountryBd PublicAppVersionCountry = "BD"
	PublicAppVersionCountryBb PublicAppVersionCountry = "BB"
	PublicAppVersionCountryBy PublicAppVersionCountry = "BY"
	PublicAppVersionCountryBe PublicAppVersionCountry = "BE"
	PublicAppVersionCountryBz PublicAppVersionCountry = "BZ"
	PublicAppVersionCountryBj PublicAppVersionCountry = "BJ"
	PublicAppVersionCountryBm PublicAppVersionCountry = "BM"
	PublicAppVersionCountryBt PublicAppVersionCountry = "BT"
	PublicAppVersionCountryBo PublicAppVersionCountry = "BO"
	PublicAppVersionCountryBq PublicAppVersionCountry = "BQ"
	PublicAppVersionCountryBa PublicAppVersionCountry = "BA"
	PublicAppVersionCountryBw PublicAppVersionCountry = "BW"
	PublicAppVersionCountryBv PublicAppVersionCountry = "BV"
	PublicAppVersionCountryBr PublicAppVersionCountry = "BR"
	PublicAppVersionCountryIo PublicAppVersionCountry = "IO"
	PublicAppVersionCountryBn PublicAppVersionCountry = "BN"
	PublicAppVersionCountryBg PublicAppVersionCountry = "BG"
	PublicAppVersionCountryBf PublicAppVersionCountry = "BF"
	PublicAppVersionCountryBi PublicAppVersionCountry = "BI"
	PublicAppVersionCountryKh PublicAppVersionCountry = "KH"
	PublicAppVersionCountryCm PublicAppVersionCountry = "CM"
	PublicAppVersionCountryCa PublicAppVersionCountry = "CA"
	PublicAppVersionCountryCv PublicAppVersionCountry = "CV"
	PublicAppVersionCountryKy PublicAppVersionCountry = "KY"
	PublicAppVersionCountryCf PublicAppVersionCountry = "CF"
	PublicAppVersionCountryTd PublicAppVersionCountry = "TD"
	PublicAppVersionCountryCl PublicAppVersionCountry = "CL"
	PublicAppVersionCountryCn PublicAppVersionCountry = "CN"
	PublicAppVersionCountryCx PublicAppVersionCountry = "CX"
	PublicAppVersionCountryCc PublicAppVersionCountry = "CC"
	PublicAppVersionCountryCo PublicAppVersionCountry = "CO"
	PublicAppVersionCountryKm PublicAppVersionCountry = "KM"
	PublicAppVersionCountryCg PublicAppVersionCountry = "CG"
	PublicAppVersionCountryCd PublicAppVersionCountry = "CD"
	PublicAppVersionCountryCk PublicAppVersionCountry = "CK"
	PublicAppVersionCountryCr PublicAppVersionCountry = "CR"
	PublicAppVersionCountryHr PublicAppVersionCountry = "HR"
	PublicAppVersionCountryCu PublicAppVersionCountry = "CU"
	PublicAppVersionCountryCw PublicAppVersionCountry = "CW"
	PublicAppVersionCountryCy PublicAppVersionCountry = "CY"
	PublicAppVersionCountryCz PublicAppVersionCountry = "CZ"
	PublicAppVersionCountryCi PublicAppVersionCountry = "CI"
	PublicAppVersionCountryDk PublicAppVersionCountry = "DK"
	PublicAppVersionCountryDj PublicAppVersionCountry = "DJ"
	PublicAppVersionCountryDm PublicAppVersionCountry = "DM"
	PublicAppVersionCountryDo PublicAppVersionCountry = "DO"
	PublicAppVersionCountryEc PublicAppVersionCountry = "EC"
	PublicAppVersionCountryEg PublicAppVersionCountry = "EG"
	PublicAppVersionCountrySv PublicAppVersionCountry = "SV"
	PublicAppVersionCountryGq PublicAppVersionCountry = "GQ"
	PublicAppVersionCountryEr PublicAppVersionCountry = "ER"
	PublicAppVersionCountryEe PublicAppVersionCountry = "EE"
	PublicAppVersionCountryEt PublicAppVersionCountry = "ET"
	PublicAppVersionCountryFk PublicAppVersionCountry = "FK"
	PublicAppVersionCountryFo PublicAppVersionCountry = "FO"
	PublicAppVersionCountryFj PublicAppVersionCountry = "FJ"
	PublicAppVersionCountryFi PublicAppVersionCountry = "FI"
	PublicAppVersionCountryFr PublicAppVersionCountry = "FR"
	PublicAppVersionCountryGf PublicAppVersionCountry = "GF"
	PublicAppVersionCountryPf PublicAppVersionCountry = "PF"
	PublicAppVersionCountryTf PublicAppVersionCountry = "TF"
	PublicAppVersionCountryGa PublicAppVersionCountry = "GA"
	PublicAppVersionCountryGm PublicAppVersionCountry = "GM"
	PublicAppVersionCountryGe PublicAppVersionCountry = "GE"
	PublicAppVersionCountryDe PublicAppVersionCountry = "DE"
	PublicAppVersionCountryGh PublicAppVersionCountry = "GH"
	PublicAppVersionCountryGi PublicAppVersionCountry = "GI"
	PublicAppVersionCountryGr PublicAppVersionCountry = "GR"
	PublicAppVersionCountryGl PublicAppVersionCountry = "GL"
	PublicAppVersionCountryGd PublicAppVersionCountry = "GD"
	PublicAppVersionCountryGp PublicAppVersionCountry = "GP"
	PublicAppVersionCountryGu PublicAppVersionCountry = "GU"
	PublicAppVersionCountryGt PublicAppVersionCountry = "GT"
	PublicAppVersionCountryGg PublicAppVersionCountry = "GG"
	PublicAppVersionCountryGn PublicAppVersionCountry = "GN"
	PublicAppVersionCountryGw PublicAppVersionCountry = "GW"
	PublicAppVersionCountryGy PublicAppVersionCountry = "GY"
	PublicAppVersionCountryHt PublicAppVersionCountry = "HT"
	PublicAppVersionCountryHm PublicAppVersionCountry = "HM"
	PublicAppVersionCountryVa PublicAppVersionCountry = "VA"
	PublicAppVersionCountryHn PublicAppVersionCountry = "HN"
	PublicAppVersionCountryHk PublicAppVersionCountry = "HK"
	PublicAppVersionCountryHu PublicAppVersionCountry = "HU"
	PublicAppVersionCountryIs PublicAppVersionCountry = "IS"
	PublicAppVersionCountryIn PublicAppVersionCountry = "IN"
	PublicAppVersionCountryID PublicAppVersionCountry = "ID"
	PublicAppVersionCountryIr PublicAppVersionCountry = "IR"
	PublicAppVersionCountryIq PublicAppVersionCountry = "IQ"
	PublicAppVersionCountryIe PublicAppVersionCountry = "IE"
	PublicAppVersionCountryIm PublicAppVersionCountry = "IM"
	PublicAppVersionCountryIl PublicAppVersionCountry = "IL"
	PublicAppVersionCountryIt PublicAppVersionCountry = "IT"
	PublicAppVersionCountryJm PublicAppVersionCountry = "JM"
	PublicAppVersionCountryJp PublicAppVersionCountry = "JP"
	PublicAppVersionCountryJe PublicAppVersionCountry = "JE"
	PublicAppVersionCountryJo PublicAppVersionCountry = "JO"
	PublicAppVersionCountryKz PublicAppVersionCountry = "KZ"
	PublicAppVersionCountryKe PublicAppVersionCountry = "KE"
	PublicAppVersionCountryKi PublicAppVersionCountry = "KI"
	PublicAppVersionCountryKr PublicAppVersionCountry = "KR"
	PublicAppVersionCountryKw PublicAppVersionCountry = "KW"
	PublicAppVersionCountryKg PublicAppVersionCountry = "KG"
	PublicAppVersionCountryLa PublicAppVersionCountry = "LA"
	PublicAppVersionCountryLv PublicAppVersionCountry = "LV"
	PublicAppVersionCountryLb PublicAppVersionCountry = "LB"
	PublicAppVersionCountryLs PublicAppVersionCountry = "LS"
	PublicAppVersionCountryLr PublicAppVersionCountry = "LR"
	PublicAppVersionCountryLy PublicAppVersionCountry = "LY"
	PublicAppVersionCountryLi PublicAppVersionCountry = "LI"
	PublicAppVersionCountryLt PublicAppVersionCountry = "LT"
	PublicAppVersionCountryLu PublicAppVersionCountry = "LU"
	PublicAppVersionCountryMo PublicAppVersionCountry = "MO"
	PublicAppVersionCountryMg PublicAppVersionCountry = "MG"
	PublicAppVersionCountryMw PublicAppVersionCountry = "MW"
	PublicAppVersionCountryMy PublicAppVersionCountry = "MY"
	PublicAppVersionCountryMv PublicAppVersionCountry = "MV"
	PublicAppVersionCountryMl PublicAppVersionCountry = "ML"
	PublicAppVersionCountryMt PublicAppVersionCountry = "MT"
	PublicAppVersionCountryMh PublicAppVersionCountry = "MH"
	PublicAppVersionCountryMq PublicAppVersionCountry = "MQ"
	PublicAppVersionCountryMr PublicAppVersionCountry = "MR"
	PublicAppVersionCountryMu PublicAppVersionCountry = "MU"
	PublicAppVersionCountryYt PublicAppVersionCountry = "YT"
	PublicAppVersionCountryMx PublicAppVersionCountry = "MX"
	PublicAppVersionCountryFm PublicAppVersionCountry = "FM"
	PublicAppVersionCountryMd PublicAppVersionCountry = "MD"
	PublicAppVersionCountryMc PublicAppVersionCountry = "MC"
	PublicAppVersionCountryMn PublicAppVersionCountry = "MN"
	PublicAppVersionCountryMe PublicAppVersionCountry = "ME"
	PublicAppVersionCountryMs PublicAppVersionCountry = "MS"
	PublicAppVersionCountryMa PublicAppVersionCountry = "MA"
	PublicAppVersionCountryMz PublicAppVersionCountry = "MZ"
	PublicAppVersionCountryMm PublicAppVersionCountry = "MM"
	PublicAppVersionCountryNa PublicAppVersionCountry = "NA"
	PublicAppVersionCountryNr PublicAppVersionCountry = "NR"
	PublicAppVersionCountryNp PublicAppVersionCountry = "NP"
	PublicAppVersionCountryNl PublicAppVersionCountry = "NL"
	PublicAppVersionCountryAn PublicAppVersionCountry = "AN"
	PublicAppVersionCountryNc PublicAppVersionCountry = "NC"
	PublicAppVersionCountryNz PublicAppVersionCountry = "NZ"
	PublicAppVersionCountryNi PublicAppVersionCountry = "NI"
	PublicAppVersionCountryNe PublicAppVersionCountry = "NE"
	PublicAppVersionCountryNg PublicAppVersionCountry = "NG"
	PublicAppVersionCountryNu PublicAppVersionCountry = "NU"
	PublicAppVersionCountryNf PublicAppVersionCountry = "NF"
	PublicAppVersionCountryKp PublicAppVersionCountry = "KP"
	PublicAppVersionCountryMk PublicAppVersionCountry = "MK"
	PublicAppVersionCountryMp PublicAppVersionCountry = "MP"
	PublicAppVersionCountryNo PublicAppVersionCountry = "NO"
	PublicAppVersionCountryOm PublicAppVersionCountry = "OM"
	PublicAppVersionCountryPk PublicAppVersionCountry = "PK"
	PublicAppVersionCountryPw PublicAppVersionCountry = "PW"
	PublicAppVersionCountryPs PublicAppVersionCountry = "PS"
	PublicAppVersionCountryPa PublicAppVersionCountry = "PA"
	PublicAppVersionCountryPg PublicAppVersionCountry = "PG"
	PublicAppVersionCountryPy PublicAppVersionCountry = "PY"
	PublicAppVersionCountryPe PublicAppVersionCountry = "PE"
	PublicAppVersionCountryPh PublicAppVersionCountry = "PH"
	PublicAppVersionCountryPn PublicAppVersionCountry = "PN"
	PublicAppVersionCountryPl PublicAppVersionCountry = "PL"
	PublicAppVersionCountryPt PublicAppVersionCountry = "PT"
	PublicAppVersionCountryPr PublicAppVersionCountry = "PR"
	PublicAppVersionCountryQa PublicAppVersionCountry = "QA"
	PublicAppVersionCountryRe PublicAppVersionCountry = "RE"
	PublicAppVersionCountryRo PublicAppVersionCountry = "RO"
	PublicAppVersionCountryRu PublicAppVersionCountry = "RU"
	PublicAppVersionCountryRw PublicAppVersionCountry = "RW"
	PublicAppVersionCountryBl PublicAppVersionCountry = "BL"
	PublicAppVersionCountrySh PublicAppVersionCountry = "SH"
	PublicAppVersionCountryKn PublicAppVersionCountry = "KN"
	PublicAppVersionCountryLc PublicAppVersionCountry = "LC"
	PublicAppVersionCountryMf PublicAppVersionCountry = "MF"
	PublicAppVersionCountryPm PublicAppVersionCountry = "PM"
	PublicAppVersionCountryVc PublicAppVersionCountry = "VC"
	PublicAppVersionCountryWs PublicAppVersionCountry = "WS"
	PublicAppVersionCountrySm PublicAppVersionCountry = "SM"
	PublicAppVersionCountrySt PublicAppVersionCountry = "ST"
	PublicAppVersionCountrySa PublicAppVersionCountry = "SA"
	PublicAppVersionCountrySn PublicAppVersionCountry = "SN"
	PublicAppVersionCountryRs PublicAppVersionCountry = "RS"
	PublicAppVersionCountryCs PublicAppVersionCountry = "CS"
	PublicAppVersionCountrySc PublicAppVersionCountry = "SC"
	PublicAppVersionCountrySl PublicAppVersionCountry = "SL"
	PublicAppVersionCountrySg PublicAppVersionCountry = "SG"
	PublicAppVersionCountrySx PublicAppVersionCountry = "SX"
	PublicAppVersionCountrySk PublicAppVersionCountry = "SK"
	PublicAppVersionCountrySi PublicAppVersionCountry = "SI"
	PublicAppVersionCountrySb PublicAppVersionCountry = "SB"
	PublicAppVersionCountrySo PublicAppVersionCountry = "SO"
	PublicAppVersionCountryZa PublicAppVersionCountry = "ZA"
	PublicAppVersionCountryGs PublicAppVersionCountry = "GS"
	PublicAppVersionCountrySS PublicAppVersionCountry = "SS"
	PublicAppVersionCountryEs PublicAppVersionCountry = "ES"
	PublicAppVersionCountryLk PublicAppVersionCountry = "LK"
	PublicAppVersionCountrySd PublicAppVersionCountry = "SD"
	PublicAppVersionCountrySr PublicAppVersionCountry = "SR"
	PublicAppVersionCountrySj PublicAppVersionCountry = "SJ"
	PublicAppVersionCountrySz PublicAppVersionCountry = "SZ"
	PublicAppVersionCountrySe PublicAppVersionCountry = "SE"
	PublicAppVersionCountryCh PublicAppVersionCountry = "CH"
	PublicAppVersionCountrySy PublicAppVersionCountry = "SY"
	PublicAppVersionCountryTw PublicAppVersionCountry = "TW"
	PublicAppVersionCountryTj PublicAppVersionCountry = "TJ"
	PublicAppVersionCountryTz PublicAppVersionCountry = "TZ"
	PublicAppVersionCountryTh PublicAppVersionCountry = "TH"
	PublicAppVersionCountryTl PublicAppVersionCountry = "TL"
	PublicAppVersionCountryTg PublicAppVersionCountry = "TG"
	PublicAppVersionCountryTk PublicAppVersionCountry = "TK"
	PublicAppVersionCountryTo PublicAppVersionCountry = "TO"
	PublicAppVersionCountryTt PublicAppVersionCountry = "TT"
	PublicAppVersionCountryTn PublicAppVersionCountry = "TN"
	PublicAppVersionCountryTr PublicAppVersionCountry = "TR"
	PublicAppVersionCountryTm PublicAppVersionCountry = "TM"
	PublicAppVersionCountryTc PublicAppVersionCountry = "TC"
	PublicAppVersionCountryTv PublicAppVersionCountry = "TV"
	PublicAppVersionCountryUg PublicAppVersionCountry = "UG"
	PublicAppVersionCountryUa PublicAppVersionCountry = "UA"
	PublicAppVersionCountryAe PublicAppVersionCountry = "AE"
	PublicAppVersionCountryGB PublicAppVersionCountry = "GB"
	PublicAppVersionCountryUs PublicAppVersionCountry = "US"
	PublicAppVersionCountryUm PublicAppVersionCountry = "UM"
	PublicAppVersionCountryUy PublicAppVersionCountry = "UY"
	PublicAppVersionCountryUz PublicAppVersionCountry = "UZ"
	PublicAppVersionCountryVu PublicAppVersionCountry = "VU"
	PublicAppVersionCountryVe PublicAppVersionCountry = "VE"
	PublicAppVersionCountryVn PublicAppVersionCountry = "VN"
	PublicAppVersionCountryVg PublicAppVersionCountry = "VG"
	PublicAppVersionCountryVi PublicAppVersionCountry = "VI"
	PublicAppVersionCountryWf PublicAppVersionCountry = "WF"
	PublicAppVersionCountryEh PublicAppVersionCountry = "EH"
	PublicAppVersionCountryYe PublicAppVersionCountry = "YE"
	PublicAppVersionCountryZm PublicAppVersionCountry = "ZM"
	PublicAppVersionCountryZw PublicAppVersionCountry = "ZW"
	PublicAppVersionCountryAx PublicAppVersionCountry = "AX"
)

type PublicAppVersionSource string

const (
	PublicAppVersionSourceUser   PublicAppVersionSource = "user"
	PublicAppVersionSourceSystem PublicAppVersionSource = "system"
	PublicAppVersionSourcePortal PublicAppVersionSource = "portal"
	PublicAppVersionSourceStore  PublicAppVersionSource = "store"
)

type PublicAppVersionStatus string

const (
	PublicAppVersionStatusQueued    PublicAppVersionStatus = "queued"
	PublicAppVersionStatusAvailable PublicAppVersionStatus = "available"
	PublicAppVersionStatusFailed    PublicAppVersionStatus = "failed"
	PublicAppVersionStatusPruned    PublicAppVersionStatus = "pruned"
)

type AppGetResponse struct {
	Data App `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AppGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppListResponse struct {
	ID            string    `json:"id" api:"required" format:"uuid"`
	AppStoreID    int64     `json:"appStoreId" api:"required"`
	BundleID      string    `json:"bundleId" api:"required"`
	CreatedAt     time.Time `json:"createdAt" api:"required" format:"date-time"`
	Description   string    `json:"description" api:"required"`
	DeveloperName string    `json:"developerName" api:"required"`
	DisplayName   string    `json:"displayName" api:"required"`
	IconURL       string    `json:"iconURL" api:"required"`
	// Any of "android", "ios".
	Platform  AppListResponsePlatform `json:"platform" api:"required"`
	UpdatedAt time.Time               `json:"updatedAt" api:"required" format:"date-time"`
	Version   PublicAppVersion        `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AppStoreID    respjson.Field
		BundleID      respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		DeveloperName respjson.Field
		DisplayName   respjson.Field
		IconURL       respjson.Field
		Platform      respjson.Field
		UpdatedAt     respjson.Field
		Version       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppListResponse) RawJSON() string { return r.JSON.raw }
func (r *AppListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppListResponsePlatform string

const (
	AppListResponsePlatformAndroid AppListResponsePlatform = "android"
	AppListResponsePlatformIos     AppListResponsePlatform = "ios"
)

type AppDeleteResponse struct {
	Message string `json:"message" api:"required"`
	// Any of true.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AppDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppConfirmUploadResponse struct {
	Message string `json:"message" api:"required"`
	// Any of true.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppConfirmUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *AppConfirmUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppNewSignedUploadURLResponse struct {
	// App ID
	AppID string `json:"appId" api:"required" format:"uuid"`
	// Pre-signed upload URLs for uploading app files
	R2UploadURLs []AppNewSignedUploadURLResponseR2UploadURL `json:"r2UploadUrls" api:"required"`
	// App version ID
	VersionID string `json:"versionId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID        respjson.Field
		R2UploadURLs respjson.Field
		VersionID    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppNewSignedUploadURLResponse) RawJSON() string { return r.JSON.raw }
func (r *AppNewSignedUploadURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppNewSignedUploadURLResponseR2UploadURL struct {
	FileName    string `json:"fileName" api:"required"`
	R2UploadURL string `json:"r2UploadUrl" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileName    respjson.Field
		R2UploadURL respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppNewSignedUploadURLResponseR2UploadURL) RawJSON() string { return r.JSON.raw }
func (r *AppNewSignedUploadURLResponseR2UploadURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppListVersionsResponse struct {
	Data []PublicAppVersion `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppListVersionsResponse) RawJSON() string { return r.JSON.raw }
func (r *AppListVersionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppMarkFailedResponse struct {
	Message string `json:"message" api:"required"`
	// Any of true.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppMarkFailedResponse) RawJSON() string { return r.JSON.raw }
func (r *AppMarkFailedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppStorageUsageResponse struct {
	Data AppStorageUsageResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppStorageUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *AppStorageUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppStorageUsageResponseData struct {
	// Bytes included in the plan. Null when unlimited, unknown, or storage billing is
	// off.
	IncludedBytes float64 `json:"includedBytes" api:"required"`
	// Hard admission bound in bytes: included + max purchasable overage, or included
	// when overage is not allowed. Null = unlimited, unknown, or storage billing is
	// off.
	MaxBytes float64 `json:"maxBytes" api:"required"`
	// True when usage above includedBytes is billed as overage instead of blocked
	OverageAllowed bool `json:"overageAllowed" api:"required"`
	// Bytes currently used (the sum of the user’s app versions when storage billing is
	// off)
	UsedBytes float64 `json:"usedBytes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludedBytes  respjson.Field
		MaxBytes       respjson.Field
		OverageAllowed respjson.Field
		UsedBytes      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppStorageUsageResponseData) RawJSON() string { return r.JSON.raw }
func (r *AppStorageUsageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppListParams struct {
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	// Any of "asc", "desc".
	Order AppListParamsOrder `query:"order,omitzero" json:"-"`
	// Any of "all", "android", "ios".
	Platform AppListParamsPlatform `query:"platform,omitzero" json:"-"`
	// Any of "createdAt", "name".
	SortBy AppListParamsSortBy `query:"sortBy,omitzero" json:"-"`
	// Only list apps that have a version of this source (for example `user` for the
	// apps you uploaded).
	//
	// Any of "user", "system", "portal", "store".
	Source AppListParamsSource `query:"source,omitzero" json:"-"`
	// Any of "all", "queued", "available", "failed", "pruned".
	Status AppListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AppListParams]'s query parameters as `url.Values`.
func (r AppListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AppListParamsOrder string

const (
	AppListParamsOrderAsc  AppListParamsOrder = "asc"
	AppListParamsOrderDesc AppListParamsOrder = "desc"
)

type AppListParamsPlatform string

const (
	AppListParamsPlatformAll     AppListParamsPlatform = "all"
	AppListParamsPlatformAndroid AppListParamsPlatform = "android"
	AppListParamsPlatformIos     AppListParamsPlatform = "ios"
)

type AppListParamsSortBy string

const (
	AppListParamsSortByCreatedAt AppListParamsSortBy = "createdAt"
	AppListParamsSortByName      AppListParamsSortBy = "name"
)

// Only list apps that have a version of this source (for example `user` for the
// apps you uploaded).
type AppListParamsSource string

const (
	AppListParamsSourceUser   AppListParamsSource = "user"
	AppListParamsSourceSystem AppListParamsSource = "system"
	AppListParamsSourcePortal AppListParamsSource = "portal"
	AppListParamsSourceStore  AppListParamsSource = "store"
)

type AppListParamsStatus string

const (
	AppListParamsStatusAll       AppListParamsStatus = "all"
	AppListParamsStatusQueued    AppListParamsStatus = "queued"
	AppListParamsStatusAvailable AppListParamsStatus = "available"
	AppListParamsStatusFailed    AppListParamsStatus = "failed"
	AppListParamsStatusPruned    AppListParamsStatus = "pruned"
)

type AppNewSignedUploadURLParams struct {
	BundleID    string                            `json:"bundleId" api:"required"`
	DisplayName string                            `json:"displayName" api:"required"`
	Files       []AppNewSignedUploadURLParamsFile `json:"files,omitzero" api:"required"`
	VersionCode float64                           `json:"versionCode" api:"required"`
	VersionName string                            `json:"versionName" api:"required"`
	// Country code for Search Results
	Country       param.Opt[string]  `json:"country,omitzero"`
	Description   param.Opt[string]  `json:"description,omitzero"`
	DeveloperName param.Opt[string]  `json:"developerName,omitzero"`
	IconURL       param.Opt[string]  `json:"iconURL,omitzero"`
	TargetSDK     param.Opt[float64] `json:"targetSdk,omitzero"`
	// Any of "android", "ios".
	Platform AppNewSignedUploadURLParamsPlatform `json:"platform,omitzero"`
	paramObj
}

func (r AppNewSignedUploadURLParams) MarshalJSON() (data []byte, err error) {
	type shadow AppNewSignedUploadURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppNewSignedUploadURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContentType, FileName are required.
type AppNewSignedUploadURLParamsFile struct {
	// Any of "application/vnd.android.package-archive", "application/octet-stream",
	// "application/zip".
	ContentType string            `json:"contentType,omitzero" api:"required"`
	FileName    string            `json:"fileName" api:"required"`
	Sha256      param.Opt[string] `json:"sha256,omitzero"`
	paramObj
}

func (r AppNewSignedUploadURLParamsFile) MarshalJSON() (data []byte, err error) {
	type shadow AppNewSignedUploadURLParamsFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppNewSignedUploadURLParamsFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AppNewSignedUploadURLParamsFile](
		"contentType", "application/vnd.android.package-archive", "application/octet-stream", "application/zip",
	)
}

type AppNewSignedUploadURLParamsPlatform string

const (
	AppNewSignedUploadURLParamsPlatformAndroid AppNewSignedUploadURLParamsPlatform = "android"
	AppNewSignedUploadURLParamsPlatformIos     AppNewSignedUploadURLParamsPlatform = "ios"
)
