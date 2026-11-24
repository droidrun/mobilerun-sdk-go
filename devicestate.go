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
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apiquery"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// DeviceStateService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceStateService] method instead.
type DeviceStateService struct {
	options []option.RequestOption
}

// NewDeviceStateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceStateService(opts ...option.RequestOption) (r DeviceStateService) {
	r = DeviceStateService{}
	r.options = opts
	return
}

// Deprecated: use POST /devices/{deviceId}/screenshots. Captures the device screen
// and returns it as a PNG image. An optional hideOverlay query parameter excludes
// the accessibility overlay from the capture.
//
// Deprecated: deprecated
func (r *DeviceStateService) Screenshot(ctx context.Context, deviceID string, params DeviceStateScreenshotParams, opts ...option.RequestOption) (res *string, err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/screenshot", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns the device's current wall-clock time as an RFC 3339 timestamp.
func (r *DeviceStateService) Time(ctx context.Context, deviceID string, query DeviceStateTimeParams, opts ...option.RequestOption) (res *string, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/time", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the current accessibility UI state of the device as a structured tree of
// on-screen elements. An optional filter query reduces the result to interactive
// elements.
func (r *DeviceStateService) Ui(ctx context.Context, deviceID string, params DeviceStateUiParams, opts ...option.RequestOption) (res *DeviceStateUiResponse, err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/ui-state", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type A11YNode struct {
	BoundsInScreen     A11YNodeBoundsInScreen `json:"boundsInScreen" api:"required"`
	Children           []A11YNode             `json:"children" api:"required"`
	ClassName          string                 `json:"className" api:"required"`
	ContentDescription string                 `json:"contentDescription" api:"required"`
	IsCheckable        bool                   `json:"isCheckable" api:"required"`
	IsChecked          bool                   `json:"isChecked" api:"required"`
	IsClickable        bool                   `json:"isClickable" api:"required"`
	IsEnabled          bool                   `json:"isEnabled" api:"required"`
	IsFocusable        bool                   `json:"isFocusable" api:"required"`
	IsFocused          bool                   `json:"isFocused" api:"required"`
	IsLongClickable    bool                   `json:"isLongClickable" api:"required"`
	IsPassword         bool                   `json:"isPassword" api:"required"`
	IsScrollable       bool                   `json:"isScrollable" api:"required"`
	IsSelected         bool                   `json:"isSelected" api:"required"`
	PackageName        string                 `json:"packageName" api:"required"`
	ResourceID         string                 `json:"resourceId" api:"required"`
	Text               string                 `json:"text" api:"required"`
	Hint               string                 `json:"hint"`
	IsVisibleToUser    bool                   `json:"isVisibleToUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoundsInScreen     respjson.Field
		Children           respjson.Field
		ClassName          respjson.Field
		ContentDescription respjson.Field
		IsCheckable        respjson.Field
		IsChecked          respjson.Field
		IsClickable        respjson.Field
		IsEnabled          respjson.Field
		IsFocusable        respjson.Field
		IsFocused          respjson.Field
		IsLongClickable    respjson.Field
		IsPassword         respjson.Field
		IsScrollable       respjson.Field
		IsSelected         respjson.Field
		PackageName        respjson.Field
		ResourceID         respjson.Field
		Text               respjson.Field
		Hint               respjson.Field
		IsVisibleToUser    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r A11YNode) RawJSON() string { return r.JSON.raw }
func (r *A11YNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type A11YNodeBoundsInScreen struct {
	Bottom int64 `json:"bottom" api:"required"`
	Left   int64 `json:"left" api:"required"`
	Right  int64 `json:"right" api:"required"`
	Top    int64 `json:"top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bottom      respjson.Field
		Left        respjson.Field
		Right       respjson.Field
		Top         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r A11YNodeBoundsInScreen) RawJSON() string { return r.JSON.raw }
func (r *A11YNodeBoundsInScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceStateUiResponse struct {
	A11yTree      A11YNode                           `json:"a11y_tree" api:"required"`
	DeviceContext DeviceStateUiResponseDeviceContext `json:"device_context" api:"required"`
	ImeTree       A11YNode                           `json:"ime_tree" api:"required"`
	PhoneState    DeviceStateUiResponsePhoneState    `json:"phone_state" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		A11yTree      respjson.Field
		DeviceContext respjson.Field
		ImeTree       respjson.Field
		PhoneState    respjson.Field
		Schema        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceStateUiResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceStateUiResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceStateUiResponseDeviceContext struct {
	DisplayMetrics  DeviceStateUiResponseDeviceContextDisplayMetrics  `json:"display_metrics" api:"required"`
	FilteringParams DeviceStateUiResponseDeviceContextFilteringParams `json:"filtering_params" api:"required"`
	ScreenBounds    DeviceStateUiResponseDeviceContextScreenBounds    `json:"screen_bounds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayMetrics  respjson.Field
		FilteringParams respjson.Field
		ScreenBounds    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceStateUiResponseDeviceContext) RawJSON() string { return r.JSON.raw }
func (r *DeviceStateUiResponseDeviceContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceStateUiResponseDeviceContextDisplayMetrics struct {
	Density       float64 `json:"density" api:"required"`
	DensityDpi    int64   `json:"densityDpi" api:"required"`
	HeightPixels  int64   `json:"heightPixels" api:"required"`
	ScaledDensity float64 `json:"scaledDensity" api:"required"`
	WidthPixels   int64   `json:"widthPixels" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Density       respjson.Field
		DensityDpi    respjson.Field
		HeightPixels  respjson.Field
		ScaledDensity respjson.Field
		WidthPixels   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceStateUiResponseDeviceContextDisplayMetrics) RawJSON() string { return r.JSON.raw }
func (r *DeviceStateUiResponseDeviceContextDisplayMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceStateUiResponseDeviceContextFilteringParams struct {
	MinElementSize int64 `json:"min_element_size" api:"required"`
	OverlayOffset  int64 `json:"overlay_offset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MinElementSize respjson.Field
		OverlayOffset  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceStateUiResponseDeviceContextFilteringParams) RawJSON() string { return r.JSON.raw }
func (r *DeviceStateUiResponseDeviceContextFilteringParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceStateUiResponseDeviceContextScreenBounds struct {
	Height int64 `json:"height" api:"required"`
	Width  int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceStateUiResponseDeviceContextScreenBounds) RawJSON() string { return r.JSON.raw }
func (r *DeviceStateUiResponseDeviceContextScreenBounds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceStateUiResponsePhoneState struct {
	IsEditable      bool                                          `json:"isEditable" api:"required"`
	KeyboardVisible bool                                          `json:"keyboardVisible" api:"required"`
	ActivityName    string                                        `json:"activityName"`
	CurrentApp      string                                        `json:"currentApp"`
	FocusedElement  DeviceStateUiResponsePhoneStateFocusedElement `json:"focusedElement"`
	PackageName     string                                        `json:"packageName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsEditable      respjson.Field
		KeyboardVisible respjson.Field
		ActivityName    respjson.Field
		CurrentApp      respjson.Field
		FocusedElement  respjson.Field
		PackageName     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceStateUiResponsePhoneState) RawJSON() string { return r.JSON.raw }
func (r *DeviceStateUiResponsePhoneState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceStateUiResponsePhoneStateFocusedElement struct {
	ClassName  string `json:"className"`
	ResourceID string `json:"resourceId"`
	Text       string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassName   respjson.Field
		ResourceID  respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceStateUiResponsePhoneStateFocusedElement) RawJSON() string { return r.JSON.raw }
func (r *DeviceStateUiResponsePhoneStateFocusedElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceStateScreenshotParams struct {
	HideOverlay param.Opt[bool] `query:"hideOverlay,omitzero" json:"-"`
	// Store the screenshot as device media: in the owner's active recording, or
	// standalone for 7 days. Without save=true nothing is stored. The PNG response is
	// unchanged; X-Screenshot-Saved reports whether it was stored.
	Save             param.Opt[bool]  `query:"save,omitzero" json:"-"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceStateScreenshotParams]'s query parameters as
// `url.Values`.
func (r DeviceStateScreenshotParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeviceStateTimeParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceStateUiParams struct {
	Filter           param.Opt[bool]  `query:"filter,omitzero" json:"-"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeviceStateUiParams]'s query parameters as `url.Values`.
func (r DeviceStateUiParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
