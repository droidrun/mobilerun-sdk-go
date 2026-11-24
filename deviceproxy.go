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

// DeviceProxyService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceProxyService] method instead.
type DeviceProxyService struct {
	options []option.RequestOption
}

// NewDeviceProxyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceProxyService(opts ...option.RequestOption) (r DeviceProxyService) {
	r = DeviceProxyService{}
	r.options = opts
	return
}

// Routes the device's traffic through a SOCKS5 proxy supplied in the request body,
// replacing any existing connection. A smartIp option can be used to select an IP
// automatically; the legacy flat host/port/user/password fields remain supported.
func (r *DeviceProxyService) Connect(ctx context.Context, deviceID string, params DeviceProxyConnectParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", params.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/proxy", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Disconnects the device's active proxy connection and clears its stored proxy
// state. Returns successfully if no proxy is connected.
func (r *DeviceProxyService) Disconnect(ctx context.Context, deviceID string, body DeviceProxyDisconnectParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", body.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return err
	}
	path := fmt.Sprintf("devices/%s/proxy", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns the device's current proxy connection state, including whether a proxy
// is connected and its protocol and name.
func (r *DeviceProxyService) Status(ctx context.Context, deviceID string, query DeviceProxyStatusParams, opts ...option.RequestOption) (res *DeviceProxyStatusResponse, err error) {
	if !param.IsOmitted(query.XDeviceDisplayID) {
		opts = append(opts, option.WithHeader("X-Device-Display-ID", fmt.Sprintf("%v", query.XDeviceDisplayID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devices/%s/proxy", url.PathEscape(deviceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DeviceProxyStatusResponse struct {
	Connected bool `json:"connected" api:"required"`
	// Active proxy name
	Name string `json:"name" api:"required"`
	// Active proxy protocol (socks5).
	Protocol string `json:"protocol" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connected   respjson.Field
		Name        respjson.Field
		Protocol    respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceProxyStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceProxyStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceProxyConnectParams struct {
	Host param.Opt[string] `json:"host,omitzero" format:"hostname"`
	// Proxy name
	Name             param.Opt[string] `json:"name,omitzero"`
	Password         param.Opt[string] `json:"password,omitzero"`
	Port             param.Opt[int64]  `json:"port,omitzero"`
	SmartIP          param.Opt[bool]   `json:"smartIp,omitzero"`
	User             param.Opt[string] `json:"user,omitzero"`
	XDeviceDisplayID param.Opt[int64]  `header:"X-Device-Display-ID,omitzero" json:"-"`
	// Mobilerun Connect proxy — pass exactly one of id (use an existing proxy's
	// credentials) or country (provision or reuse a rotating residential proxy for the
	// device).
	Connect DeviceProxyConnectParamsConnect `json:"connect,omitzero"`
	// SOCKS5 proxy configuration (required for socks5).
	Socks5 DeviceProxyConnectParamsSocks5 `json:"socks5,omitzero"`
	paramObj
}

func (r DeviceProxyConnectParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceProxyConnectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceProxyConnectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mobilerun Connect proxy — pass exactly one of id (use an existing proxy's
// credentials) or country (provision or reuse a rotating residential proxy for the
// device).
type DeviceProxyConnectParamsConnect struct {
	// Existing Mobilerun Connect proxy id; its credentials are fetched server-side.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// ISO 3166-1 alpha-2 country code; provisions (or reuses) a rotating residential
	// Mobilerun Connect proxy for the device.
	Country param.Opt[string] `json:"country,omitzero"`
	paramObj
}

func (r DeviceProxyConnectParamsConnect) MarshalJSON() (data []byte, err error) {
	type shadow DeviceProxyConnectParamsConnect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceProxyConnectParamsConnect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SOCKS5 proxy configuration (required for socks5).
//
// The properties Host, Port are required.
type DeviceProxyConnectParamsSocks5 struct {
	Host     string            `json:"host" api:"required" format:"hostname"`
	Port     int64             `json:"port" api:"required"`
	Password param.Opt[string] `json:"password,omitzero"`
	User     param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r DeviceProxyConnectParamsSocks5) MarshalJSON() (data []byte, err error) {
	type shadow DeviceProxyConnectParamsSocks5
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceProxyConnectParamsSocks5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceProxyDisconnectParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}

type DeviceProxyStatusParams struct {
	XDeviceDisplayID param.Opt[int64] `header:"X-Device-Display-ID,omitzero" json:"-"`
	paramObj
}
