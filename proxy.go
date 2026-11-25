// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"encoding/json"
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
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared/constant"
)

// ProxyService contains methods and other services that help with interacting with
// the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProxyService] method instead.
type ProxyService struct {
	options []option.RequestOption
}

// NewProxyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProxyService(opts ...option.RequestOption) (r ProxyService) {
	r = ProxyService{}
	r.options = opts
	return
}

// Creates a proxy config. The body is a discriminated union on `protocol`:
// `socks5` requires name, host, port, user, and password, while `wireguard`
// requires name and a config string. Returns the created config with its generated
// `proxyId`.
func (r *ProxyService) New(ctx context.Context, body ProxyNewParams, opts ...option.RequestOption) (res *ProxyNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "proxies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetches a single proxy config by its `proxyId`. The response shape depends on
// the proxy's `protocol` (socks5 or wireguard). Returns not found if no matching
// config exists.
func (r *ProxyService) Get(ctx context.Context, proxyID string, opts ...option.RequestOption) (res *ProxyGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if proxyID == "" {
		err = errors.New("missing required proxyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("proxies/%s", proxyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replaces the proxy config identified by `proxyId` with the provided body. As
// with creation, the body is a `protocol`-discriminated union of socks5 or
// wireguard fields. Returns the updated config.
func (r *ProxyService) Update(ctx context.Context, proxyID string, body ProxyUpdateParams, opts ...option.RequestOption) (res *ProxyUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if proxyID == "" {
		err = errors.New("missing required proxyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("proxies/%s", proxyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns all proxy configs for the authenticated user. An optional `protocol`
// query parameter (`socks5` or `wireguard`) filters the results by proxy type.
func (r *ProxyService) List(ctx context.Context, query ProxyListParams, opts ...option.RequestOption) (res *ProxyListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "proxies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes the proxy config identified by `proxyId` and returns the
// deleted config. Returns not found if no matching config exists.
func (r *ProxyService) Delete(ctx context.Context, proxyID string, opts ...option.RequestOption) (res *ProxyDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if proxyID == "" {
		err = errors.New("missing required proxyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("proxies/%s", proxyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type WireguardProxyConfig struct {
	Config string `json:"config" api:"required"`
	Name   string `json:"name" api:"required"`
	// Any of "wireguard".
	Protocol WireguardProxyConfigProtocol `json:"protocol" api:"required"`
	ProxyID  string                       `json:"proxyId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Name        respjson.Field
		Protocol    respjson.Field
		ProxyID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireguardProxyConfig) RawJSON() string { return r.JSON.raw }
func (r *WireguardProxyConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WireguardProxyConfigProtocol string

const (
	WireguardProxyConfigProtocolWireguard WireguardProxyConfigProtocol = "wireguard"
)

type ProxyNewResponse struct {
	Data    ProxyNewResponseDataUnion `json:"data" api:"required"`
	Message string                    `json:"message" api:"required"`
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
func (r ProxyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProxyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProxyNewResponseDataUnion contains all possible properties and values from
// [shared.Socks5ProxyConfig], [WireguardProxyConfig].
//
// Use the [ProxyNewResponseDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProxyNewResponseDataUnion struct {
	// This field is from variant [shared.Socks5ProxyConfig].
	Host string `json:"host"`
	Name string `json:"name"`
	// This field is from variant [shared.Socks5ProxyConfig].
	Password string `json:"password"`
	// This field is from variant [shared.Socks5ProxyConfig].
	Port int64 `json:"port"`
	// Any of "socks5", "wireguard".
	Protocol string `json:"protocol"`
	ProxyID  string `json:"proxyId"`
	// This field is from variant [shared.Socks5ProxyConfig].
	User string `json:"user"`
	// This field is from variant [WireguardProxyConfig].
	Config string `json:"config"`
	JSON   struct {
		Host     respjson.Field
		Name     respjson.Field
		Password respjson.Field
		Port     respjson.Field
		Protocol respjson.Field
		ProxyID  respjson.Field
		User     respjson.Field
		Config   respjson.Field
		raw      string
	} `json:"-"`
}

// anyProxyNewResponseData is implemented by each variant of
// [ProxyNewResponseDataUnion] to add type safety for the return type of
// [ProxyNewResponseDataUnion.AsAny]
type anyProxyNewResponseData interface {
	ImplProxyNewResponseDataUnion()
}

func (WireguardProxyConfig) ImplProxyNewResponseDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ProxyNewResponseDataUnion.AsAny().(type) {
//	case shared.Socks5ProxyConfig:
//	case mobilerun.WireguardProxyConfig:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ProxyNewResponseDataUnion) AsAny() anyProxyNewResponseData {
	switch u.Protocol {
	case "socks5":
		return u.AsSocks5()
	case "wireguard":
		return u.AsWireguard()
	}
	return nil
}

func (u ProxyNewResponseDataUnion) AsSocks5() (v shared.Socks5ProxyConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProxyNewResponseDataUnion) AsWireguard() (v WireguardProxyConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProxyNewResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ProxyNewResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyGetResponse struct {
	Data ProxyGetResponseDataUnion `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProxyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProxyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProxyGetResponseDataUnion contains all possible properties and values from
// [shared.Socks5ProxyConfig], [WireguardProxyConfig].
//
// Use the [ProxyGetResponseDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProxyGetResponseDataUnion struct {
	// This field is from variant [shared.Socks5ProxyConfig].
	Host string `json:"host"`
	Name string `json:"name"`
	// This field is from variant [shared.Socks5ProxyConfig].
	Password string `json:"password"`
	// This field is from variant [shared.Socks5ProxyConfig].
	Port int64 `json:"port"`
	// Any of "socks5", "wireguard".
	Protocol string `json:"protocol"`
	ProxyID  string `json:"proxyId"`
	// This field is from variant [shared.Socks5ProxyConfig].
	User string `json:"user"`
	// This field is from variant [WireguardProxyConfig].
	Config string `json:"config"`
	JSON   struct {
		Host     respjson.Field
		Name     respjson.Field
		Password respjson.Field
		Port     respjson.Field
		Protocol respjson.Field
		ProxyID  respjson.Field
		User     respjson.Field
		Config   respjson.Field
		raw      string
	} `json:"-"`
}

// anyProxyGetResponseData is implemented by each variant of
// [ProxyGetResponseDataUnion] to add type safety for the return type of
// [ProxyGetResponseDataUnion.AsAny]
type anyProxyGetResponseData interface {
	ImplProxyGetResponseDataUnion()
}

func (WireguardProxyConfig) ImplProxyGetResponseDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ProxyGetResponseDataUnion.AsAny().(type) {
//	case shared.Socks5ProxyConfig:
//	case mobilerun.WireguardProxyConfig:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ProxyGetResponseDataUnion) AsAny() anyProxyGetResponseData {
	switch u.Protocol {
	case "socks5":
		return u.AsSocks5()
	case "wireguard":
		return u.AsWireguard()
	}
	return nil
}

func (u ProxyGetResponseDataUnion) AsSocks5() (v shared.Socks5ProxyConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProxyGetResponseDataUnion) AsWireguard() (v WireguardProxyConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProxyGetResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ProxyGetResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyUpdateResponse struct {
	Data    ProxyUpdateResponseDataUnion `json:"data" api:"required"`
	Message string                       `json:"message" api:"required"`
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
func (r ProxyUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ProxyUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProxyUpdateResponseDataUnion contains all possible properties and values from
// [shared.Socks5ProxyConfig], [WireguardProxyConfig].
//
// Use the [ProxyUpdateResponseDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProxyUpdateResponseDataUnion struct {
	// This field is from variant [shared.Socks5ProxyConfig].
	Host string `json:"host"`
	Name string `json:"name"`
	// This field is from variant [shared.Socks5ProxyConfig].
	Password string `json:"password"`
	// This field is from variant [shared.Socks5ProxyConfig].
	Port int64 `json:"port"`
	// Any of "socks5", "wireguard".
	Protocol string `json:"protocol"`
	ProxyID  string `json:"proxyId"`
	// This field is from variant [shared.Socks5ProxyConfig].
	User string `json:"user"`
	// This field is from variant [WireguardProxyConfig].
	Config string `json:"config"`
	JSON   struct {
		Host     respjson.Field
		Name     respjson.Field
		Password respjson.Field
		Port     respjson.Field
		Protocol respjson.Field
		ProxyID  respjson.Field
		User     respjson.Field
		Config   respjson.Field
		raw      string
	} `json:"-"`
}

// anyProxyUpdateResponseData is implemented by each variant of
// [ProxyUpdateResponseDataUnion] to add type safety for the return type of
// [ProxyUpdateResponseDataUnion.AsAny]
type anyProxyUpdateResponseData interface {
	ImplProxyUpdateResponseDataUnion()
}

func (WireguardProxyConfig) ImplProxyUpdateResponseDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ProxyUpdateResponseDataUnion.AsAny().(type) {
//	case shared.Socks5ProxyConfig:
//	case mobilerun.WireguardProxyConfig:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ProxyUpdateResponseDataUnion) AsAny() anyProxyUpdateResponseData {
	switch u.Protocol {
	case "socks5":
		return u.AsSocks5()
	case "wireguard":
		return u.AsWireguard()
	}
	return nil
}

func (u ProxyUpdateResponseDataUnion) AsSocks5() (v shared.Socks5ProxyConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProxyUpdateResponseDataUnion) AsWireguard() (v WireguardProxyConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProxyUpdateResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ProxyUpdateResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyListResponse struct {
	Data []ProxyListResponseDataUnion `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProxyListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProxyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProxyListResponseDataUnion contains all possible properties and values from
// [shared.Socks5ProxyConfig], [WireguardProxyConfig].
//
// Use the [ProxyListResponseDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProxyListResponseDataUnion struct {
	// This field is from variant [shared.Socks5ProxyConfig].
	Host string `json:"host"`
	Name string `json:"name"`
	// This field is from variant [shared.Socks5ProxyConfig].
	Password string `json:"password"`
	// This field is from variant [shared.Socks5ProxyConfig].
	Port int64 `json:"port"`
	// Any of "socks5", "wireguard".
	Protocol string `json:"protocol"`
	ProxyID  string `json:"proxyId"`
	// This field is from variant [shared.Socks5ProxyConfig].
	User string `json:"user"`
	// This field is from variant [WireguardProxyConfig].
	Config string `json:"config"`
	JSON   struct {
		Host     respjson.Field
		Name     respjson.Field
		Password respjson.Field
		Port     respjson.Field
		Protocol respjson.Field
		ProxyID  respjson.Field
		User     respjson.Field
		Config   respjson.Field
		raw      string
	} `json:"-"`
}

// anyProxyListResponseData is implemented by each variant of
// [ProxyListResponseDataUnion] to add type safety for the return type of
// [ProxyListResponseDataUnion.AsAny]
type anyProxyListResponseData interface {
	ImplProxyListResponseDataUnion()
}

func (WireguardProxyConfig) ImplProxyListResponseDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ProxyListResponseDataUnion.AsAny().(type) {
//	case shared.Socks5ProxyConfig:
//	case mobilerun.WireguardProxyConfig:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ProxyListResponseDataUnion) AsAny() anyProxyListResponseData {
	switch u.Protocol {
	case "socks5":
		return u.AsSocks5()
	case "wireguard":
		return u.AsWireguard()
	}
	return nil
}

func (u ProxyListResponseDataUnion) AsSocks5() (v shared.Socks5ProxyConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProxyListResponseDataUnion) AsWireguard() (v WireguardProxyConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProxyListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ProxyListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyDeleteResponse struct {
	Data    ProxyDeleteResponseDataUnion `json:"data" api:"required"`
	Message string                       `json:"message" api:"required"`
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
func (r ProxyDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ProxyDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProxyDeleteResponseDataUnion contains all possible properties and values from
// [shared.Socks5ProxyConfig], [WireguardProxyConfig].
//
// Use the [ProxyDeleteResponseDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProxyDeleteResponseDataUnion struct {
	// This field is from variant [shared.Socks5ProxyConfig].
	Host string `json:"host"`
	Name string `json:"name"`
	// This field is from variant [shared.Socks5ProxyConfig].
	Password string `json:"password"`
	// This field is from variant [shared.Socks5ProxyConfig].
	Port int64 `json:"port"`
	// Any of "socks5", "wireguard".
	Protocol string `json:"protocol"`
	ProxyID  string `json:"proxyId"`
	// This field is from variant [shared.Socks5ProxyConfig].
	User string `json:"user"`
	// This field is from variant [WireguardProxyConfig].
	Config string `json:"config"`
	JSON   struct {
		Host     respjson.Field
		Name     respjson.Field
		Password respjson.Field
		Port     respjson.Field
		Protocol respjson.Field
		ProxyID  respjson.Field
		User     respjson.Field
		Config   respjson.Field
		raw      string
	} `json:"-"`
}

// anyProxyDeleteResponseData is implemented by each variant of
// [ProxyDeleteResponseDataUnion] to add type safety for the return type of
// [ProxyDeleteResponseDataUnion.AsAny]
type anyProxyDeleteResponseData interface {
	ImplProxyDeleteResponseDataUnion()
}

func (WireguardProxyConfig) ImplProxyDeleteResponseDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ProxyDeleteResponseDataUnion.AsAny().(type) {
//	case shared.Socks5ProxyConfig:
//	case mobilerun.WireguardProxyConfig:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ProxyDeleteResponseDataUnion) AsAny() anyProxyDeleteResponseData {
	switch u.Protocol {
	case "socks5":
		return u.AsSocks5()
	case "wireguard":
		return u.AsWireguard()
	}
	return nil
}

func (u ProxyDeleteResponseDataUnion) AsSocks5() (v shared.Socks5ProxyConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProxyDeleteResponseDataUnion) AsWireguard() (v WireguardProxyConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProxyDeleteResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ProxyDeleteResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfSocks5 *ProxyNewParamsBodySocks5 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWireguard *ProxyNewParamsBodyWireguard `json:",inline"`

	paramObj
}

func (u ProxyNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSocks5, u.OfWireguard)
}
func (r *ProxyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Host, Name, Password, Port, Protocol, User are required.
type ProxyNewParamsBodySocks5 struct {
	Host     string `json:"host" api:"required"`
	Name     string `json:"name" api:"required"`
	Password string `json:"password" api:"required"`
	Port     int64  `json:"port" api:"required"`
	User     string `json:"user" api:"required"`
	// This field can be elided, and will marshal its zero value as "socks5".
	Protocol constant.Socks5 `json:"protocol" default:"socks5"`
	paramObj
}

func (r ProxyNewParamsBodySocks5) MarshalJSON() (data []byte, err error) {
	type shadow ProxyNewParamsBodySocks5
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProxyNewParamsBodySocks5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, Name, Protocol are required.
type ProxyNewParamsBodyWireguard struct {
	Config string `json:"config" api:"required"`
	Name   string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "wireguard".
	Protocol constant.Wireguard `json:"protocol" default:"wireguard"`
	paramObj
}

func (r ProxyNewParamsBodyWireguard) MarshalJSON() (data []byte, err error) {
	type shadow ProxyNewParamsBodyWireguard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProxyNewParamsBodyWireguard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfSocks5 *ProxyUpdateParamsBodySocks5 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWireguard *ProxyUpdateParamsBodyWireguard `json:",inline"`

	paramObj
}

func (u ProxyUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSocks5, u.OfWireguard)
}
func (r *ProxyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Host, Name, Password, Port, Protocol, User are required.
type ProxyUpdateParamsBodySocks5 struct {
	Host     string `json:"host" api:"required"`
	Name     string `json:"name" api:"required"`
	Password string `json:"password" api:"required"`
	Port     int64  `json:"port" api:"required"`
	User     string `json:"user" api:"required"`
	// This field can be elided, and will marshal its zero value as "socks5".
	Protocol constant.Socks5 `json:"protocol" default:"socks5"`
	paramObj
}

func (r ProxyUpdateParamsBodySocks5) MarshalJSON() (data []byte, err error) {
	type shadow ProxyUpdateParamsBodySocks5
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProxyUpdateParamsBodySocks5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Config, Name, Protocol are required.
type ProxyUpdateParamsBodyWireguard struct {
	Config string `json:"config" api:"required"`
	Name   string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "wireguard".
	Protocol constant.Wireguard `json:"protocol" default:"wireguard"`
	paramObj
}

func (r ProxyUpdateParamsBodyWireguard) MarshalJSON() (data []byte, err error) {
	type shadow ProxyUpdateParamsBodyWireguard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProxyUpdateParamsBodyWireguard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyListParams struct {
	// Any of "socks5", "wireguard".
	Protocol ProxyListParamsProtocol `query:"protocol,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProxyListParams]'s query parameters as `url.Values`.
func (r ProxyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProxyListParamsProtocol string

const (
	ProxyListParamsProtocolSocks5    ProxyListParamsProtocol = "socks5"
	ProxyListParamsProtocolWireguard ProxyListParamsProtocol = "wireguard"
)
