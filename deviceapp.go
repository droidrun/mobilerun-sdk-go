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
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// DeviceAppService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceAppService] method instead.
type DeviceAppService struct {
	options []option.RequestOption
}

// NewDeviceAppService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceAppService(opts ...option.RequestOption) (r DeviceAppService) {
	r = DeviceAppService{}
	r.options = opts
	return
}

// Returns detailed information about apps installed on the device, including
// package name and label. System and protected apps are excluded unless the
// corresponding query parameters are set.
func (r *DeviceAppService) List(ctx context.Context, deviceID string, params DeviceAppListParams, opts ...option.RequestOption) (res *[]DeviceAppListResponse, err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/apps", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Uninstalls the app identified by the path package name from the device.
// Protected packages cannot be deleted.
func (r *DeviceAppService) Delete(ctx context.Context, packageName string, params DeviceAppDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/apps/%s", url.PathEscape(params.DeviceID), url.PathEscape(packageName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Grants an Android runtime permission to the package named in the path. The
// permission is given by its short name (e.g. POST_NOTIFICATIONS).
func (r *DeviceAppService) GrantPermission(ctx context.Context, permission DeviceAppGrantPermissionParamsPermission, params DeviceAppGrantPermissionParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	if params.PackageName == "" {
		err = errors.New("missing required packageName parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/apps/%s/permissions/%v", url.PathEscape(params.DeviceID), url.PathEscape(params.PackageName), permission)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return err
}

// Requests an app install on the device. The request body must supply exactly one
// of an Android packageName or an iOS bundleId. For packageName installs, optional
// country and versionCode select an exact regional uploaded Android version, and
// protected packages are rejected. background (default false) selects the response
// contract: false installs inline and returns the outcome directly (200 on
// success, an error status on failure); true accepts the request and runs the
// download + install in the background, returning 202 immediately — poll
// list-app-installs for the backend's view of that attempt's status. Refuses with
// 409 once 2 other installs are already running on the device, in either mode; a
// repeat request for an app that already has an install running is also refused
// with 409 rather than superseding it — retry once that attempt reaches a terminal
// state. Synchronous failures report APP_DOWNLOAD_FAILED (502) when the app file
// cannot be downloaded or APP_INSTALL_TIMEOUT (504) when the install times out.
func (r *DeviceAppService) Install(ctx context.Context, deviceID string, params DeviceAppInstallParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/apps", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Reports the backend's durable view of background app-install attempts on this
// device — status reflects the install ATTEMPT, not device ground truth; list-apps
// remains authoritative for what is actually installed. Terminal and projected
// timeout records are kept ~15 minutes. Not gated on device readiness, so it also
// answers while the device is offline or crashed.
func (r *DeviceAppService) ListInstalls(ctx context.Context, deviceID string, query DeviceAppListInstallsParams, opts ...option.RequestOption) (res *DeviceAppListInstallsResponse, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/apps/installs", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revokes an Android runtime permission from the package named in the path. The
// permission is given by its short name (e.g. POST_NOTIFICATIONS).
func (r *DeviceAppService) RevokePermission(ctx context.Context, permission DeviceAppRevokePermissionParamsPermission, params DeviceAppRevokePermissionParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	if params.PackageName == "" {
		err = errors.New("missing required packageName parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/apps/%s/permissions/%v", url.PathEscape(params.DeviceID), url.PathEscape(params.PackageName), permission)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Launches the app identified by the path package name, optionally starting a
// specific activity given in the request body. Protected packages cannot be
// started.
func (r *DeviceAppService) Start(ctx context.Context, packageName string, params DeviceAppStartParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/apps/%s", url.PathEscape(params.DeviceID), url.PathEscape(packageName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

// Force-stops the app identified by the path package name. When clearData is set
// in the request body, the app's data is also cleared. Protected packages cannot
// be stopped.
func (r *DeviceAppService) Stop(ctx context.Context, packageName string, params DeviceAppStopParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.DeviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/apps/%s", url.PathEscape(params.DeviceID), url.PathEscape(packageName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return err
}

type DeviceAppListResponse struct {
	IsSystemApp     bool                                 `json:"isSystemApp" api:"required"`
	Label           string                               `json:"label" api:"required"`
	PackageName     string                               `json:"packageName" api:"required"`
	VersionCode     int64                                `json:"versionCode" api:"required"`
	VersionName     string                               `json:"versionName" api:"required"`
	AvailableUpdate DeviceAppListResponseAvailableUpdate `json:"availableUpdate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSystemApp     respjson.Field
		Label           respjson.Field
		PackageName     respjson.Field
		VersionCode     respjson.Field
		VersionName     respjson.Field
		AvailableUpdate respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceAppListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceAppListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceAppListResponseAvailableUpdate struct {
	Country     string `json:"country" api:"required"`
	VersionCode int64  `json:"versionCode" api:"required"`
	VersionName string `json:"versionName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		VersionCode respjson.Field
		VersionName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceAppListResponseAvailableUpdate) RawJSON() string { return r.JSON.raw }
func (r *DeviceAppListResponseAvailableUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceAppListInstallsResponse struct {
	Installs []DeviceAppListInstallsResponseInstall `json:"installs" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Installs    respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceAppListInstallsResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceAppListInstallsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceAppListInstallsResponseInstall struct {
	// Android package name or iOS bundle id
	AppID string `json:"appId" api:"required"`
	// android or ios
	Platform  string    `json:"platform" api:"required"`
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// On some iOS devices, succeeded means the device accepted the install command;
	// the install may still be running on the device.
	//
	// Any of "running", "succeeded", "failed".
	Status    string    `json:"status" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Only present when status is failed.
	//
	// Any of "download_failed", "adb_install_failed", "panic", "timeout", "failed".
	ErrorClass string `json:"errorClass"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppID       respjson.Field
		Platform    respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ErrorClass  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceAppListInstallsResponseInstall) RawJSON() string { return r.JSON.raw }
func (r *DeviceAppListInstallsResponseInstall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceAppListParams struct {
	IncludeProtectedApps param.Opt[bool] `query:"includeProtectedApps,omitzero" json:"-"`
	IncludeSystemApps    param.Opt[bool] `query:"includeSystemApps,omitzero" json:"-"`
	// Resolve availableUpdate for Android user apps (one library lookup per app). Off
	// by default.
	IncludeUpdates   param.Opt[bool]  `query:"includeUpdates,omitzero" json:"-"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceAppListParams]'s query parameters as `url.Values`.
func (r DeviceAppListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceAppDeleteParams struct {
	DeviceID         string           `path:"deviceId" api:"required" json:"-"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceAppDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceAppDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceAppDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceAppGrantPermissionParams struct {
	DeviceID         string           `path:"deviceId" api:"required" json:"-"`
	PackageName      string           `path:"packageName" api:"required" json:"-"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

// Android runtime permission, short name (e.g. POST_NOTIFICATIONS).
type DeviceAppGrantPermissionParamsPermission string

const (
	DeviceAppGrantPermissionParamsPermissionPostNotifications DeviceAppGrantPermissionParamsPermission = "POST_NOTIFICATIONS"
)

type DeviceAppInstallParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfObject *DeviceAppInstallParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfDeviceAppInstallsBodyObject2 *DeviceAppInstallParamsBodyObject2 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfDeviceAppInstallsBodyObject3 *DeviceAppInstallParamsBodyObject3 `json:",inline"`

	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (u DeviceAppInstallParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfObject, u.OfDeviceAppInstallsBodyObject2, u.OfDeviceAppInstallsBodyObject3)
}
func (r *DeviceAppInstallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BundleID is required.
type DeviceAppInstallParamsBodyObject struct {
	// iOS bundle identifier (e.g. com.example.app)
	BundleID string `json:"bundleId" api:"required"`
	// UUID of a published Android or iOS store app.
	AppID param.Opt[string] `json:"appId,omitzero" format:"uuid"`
	// HTTPS APK URL downloaded and installed by the device; not stored in your app
	// library.
	AppURL param.Opt[string] `json:"appUrl,omitzero"`
	// true: return 202 immediately and install in the background (poll
	// list-app-installs). false/omitted: install inline and return the outcome
	// directly (200 on success, an error status on failure).
	Background param.Opt[bool] `json:"background,omitzero"`
	// Optional ISO 3166-1 alpha-2 country of the library or store app version (e.g. MY
	// or SG).
	Country param.Opt[string] `json:"country,omitzero"`
	// Android package name (e.g. com.example.app)
	PackageName param.Opt[string] `json:"packageName,omitzero"`
	// Optional exact library or store version code. Use with country when multiple
	// regional versions share an identifier.
	VersionCode param.Opt[int64] `json:"versionCode,omitzero"`
	paramObj
}

func (r DeviceAppInstallParamsBodyObject) MarshalJSON() (data []byte, err error) {
	type shadow DeviceAppInstallParamsBodyObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceAppInstallParamsBodyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PackageName is required.
type DeviceAppInstallParamsBodyObject2 struct {
	// Android package name (e.g. com.example.app)
	PackageName string `json:"packageName" api:"required"`
	// UUID of a published Android or iOS store app.
	AppID param.Opt[string] `json:"appId,omitzero" format:"uuid"`
	// HTTPS APK URL downloaded and installed by the device; not stored in your app
	// library.
	AppURL param.Opt[string] `json:"appUrl,omitzero"`
	// true: return 202 immediately and install in the background (poll
	// list-app-installs). false/omitted: install inline and return the outcome
	// directly (200 on success, an error status on failure).
	Background param.Opt[bool] `json:"background,omitzero"`
	// iOS bundle identifier (e.g. com.example.app)
	BundleID param.Opt[string] `json:"bundleId,omitzero"`
	// Optional ISO 3166-1 alpha-2 country of the library or store app version (e.g. MY
	// or SG).
	Country param.Opt[string] `json:"country,omitzero"`
	// Optional exact library or store version code. Use with country when multiple
	// regional versions share an identifier.
	VersionCode param.Opt[int64] `json:"versionCode,omitzero"`
	paramObj
}

func (r DeviceAppInstallParamsBodyObject2) MarshalJSON() (data []byte, err error) {
	type shadow DeviceAppInstallParamsBodyObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceAppInstallParamsBodyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property AppID is required.
type DeviceAppInstallParamsBodyObject3 struct {
	// UUID of a published Android or iOS store app.
	AppID string `json:"appId" api:"required" format:"uuid"`
	// HTTPS APK URL downloaded and installed by the device; not stored in your app
	// library.
	AppURL param.Opt[string] `json:"appUrl,omitzero"`
	// true: return 202 immediately and install in the background (poll
	// list-app-installs). false/omitted: install inline and return the outcome
	// directly (200 on success, an error status on failure).
	Background param.Opt[bool] `json:"background,omitzero"`
	// iOS bundle identifier (e.g. com.example.app)
	BundleID param.Opt[string] `json:"bundleId,omitzero"`
	// Optional ISO 3166-1 alpha-2 country of the library or store app version (e.g. MY
	// or SG).
	Country param.Opt[string] `json:"country,omitzero"`
	// Android package name (e.g. com.example.app)
	PackageName param.Opt[string] `json:"packageName,omitzero"`
	// Optional exact library or store version code. Use with country when multiple
	// regional versions share an identifier.
	VersionCode param.Opt[int64] `json:"versionCode,omitzero"`
	paramObj
}

func (r DeviceAppInstallParamsBodyObject3) MarshalJSON() (data []byte, err error) {
	type shadow DeviceAppInstallParamsBodyObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceAppInstallParamsBodyObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceAppListInstallsParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceAppRevokePermissionParams struct {
	DeviceID         string           `path:"deviceId" api:"required" json:"-"`
	PackageName      string           `path:"packageName" api:"required" json:"-"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

// Android runtime permission, short name (e.g. POST_NOTIFICATIONS).
type DeviceAppRevokePermissionParamsPermission string

const (
	DeviceAppRevokePermissionParamsPermissionPostNotifications DeviceAppRevokePermissionParamsPermission = "POST_NOTIFICATIONS"
)

type DeviceAppStartParams struct {
	DeviceID         string            `path:"deviceId" api:"required" json:"-"`
	Activity         param.Opt[string] `json:"activity,omitzero"`
	XDeviceDisplayID param.Opt[int64]  `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceAppStartParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceAppStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceAppStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceAppStopParams struct {
	DeviceID string `path:"deviceId" api:"required" json:"-"`
	// If true, clears all app data (pm clear) in addition to stopping the app.
	ClearData        param.Opt[bool]  `json:"clearData,omitzero"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceAppStopParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceAppStopParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceAppStopParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
