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

// DeviceLanguageService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceLanguageService] method instead.
type DeviceLanguageService struct {
	options []option.RequestOption
}

// NewDeviceLanguageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceLanguageService(opts ...option.RequestOption) (r DeviceLanguageService) {
	r = DeviceLanguageService{}
	r.options = opts
	return
}

// Returns the device's current language/locale as a BCP-47 locale string.
func (r *DeviceLanguageService) Get(ctx context.Context, deviceID string, query DeviceLanguageGetParams, opts ...option.RequestOption) (res *DeviceLanguageGetResponse, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/language", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Sets the device language/locale to the BCP-47 locale in the request body. An
// optional restart flag applies the change immediately by restarting the Android
// system framework (open apps are closed) instead of waiting for the next reboot;
// some devices ignore it.
func (r *DeviceLanguageService) Set(ctx context.Context, deviceID string, params DeviceLanguageSetParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/language", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type DeviceLanguageGetResponse struct {
	Locale string `json:"locale" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Locale      respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceLanguageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceLanguageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceLanguageGetParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceLanguageSetParams struct {
	// BCP-47 locale: a 2–3 letter language tag, optionally followed by a 4-letter
	// script and/or a 2-letter region (e.g. en-US, de-DE, ja-JP, zh-Hans-CN).
	Locale string `json:"locale" api:"required"`
	// Restart the Android system framework so the locale change takes full effect
	// immediately; open apps are closed. Without it, the change may not fully apply
	// until the next reboot. Some devices ignore this flag.
	Restart          param.Opt[bool]  `json:"restart,omitzero"`
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

func (r DeviceLanguageSetParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceLanguageSetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceLanguageSetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
