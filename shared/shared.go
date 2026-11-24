// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type DeviceCarrier struct {
	GsmOperatorAlpha         string `json:"GsmOperatorAlpha" api:"required"`
	GsmOperatorNumeric       int64  `json:"GsmOperatorNumeric" api:"required"`
	GsmSimOperatorAlpha      string `json:"GsmSimOperatorAlpha" api:"required"`
	GsmSimOperatorISOCountry string `json:"GsmSimOperatorIsoCountry" api:"required"`
	GsmSimOperatorNumeric    int64  `json:"GsmSimOperatorNumeric" api:"required"`
	PersistSysTimezone       string `json:"PersistSysTimezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GsmOperatorAlpha         respjson.Field
		GsmOperatorNumeric       respjson.Field
		GsmSimOperatorAlpha      respjson.Field
		GsmSimOperatorISOCountry respjson.Field
		GsmSimOperatorNumeric    respjson.Field
		PersistSysTimezone       respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceCarrier) RawJSON() string { return r.JSON.raw }
func (r *DeviceCarrier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DeviceCarrier to a DeviceCarrierParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DeviceCarrierParam.Overrides()
func (r DeviceCarrier) ToParam() DeviceCarrierParam {
	return param.Override[DeviceCarrierParam](json.RawMessage(r.RawJSON()))
}

// The properties GsmOperatorAlpha, GsmOperatorNumeric, GsmSimOperatorAlpha,
// GsmSimOperatorISOCountry, GsmSimOperatorNumeric, PersistSysTimezone are
// required.
type DeviceCarrierParam struct {
	GsmOperatorAlpha         string `json:"GsmOperatorAlpha" api:"required"`
	GsmOperatorNumeric       int64  `json:"GsmOperatorNumeric" api:"required"`
	GsmSimOperatorAlpha      string `json:"GsmSimOperatorAlpha" api:"required"`
	GsmSimOperatorISOCountry string `json:"GsmSimOperatorIsoCountry" api:"required"`
	GsmSimOperatorNumeric    int64  `json:"GsmSimOperatorNumeric" api:"required"`
	PersistSysTimezone       string `json:"PersistSysTimezone" api:"required"`
	paramObj
}

func (r DeviceCarrierParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceCarrierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceCarrierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceIdentifiers struct {
	BootloaderSerialNumber string `json:"BootloaderSerialNumber" api:"required"`
	IdentifierAndroidID    string `json:"IdentifierAndroidID" api:"required"`
	IdentifierAppSetID     string `json:"IdentifierAppSetID" api:"required"`
	IdentifierBluetoothMac string `json:"IdentifierBluetoothMAC" api:"required"`
	IdentifierGaid         string `json:"IdentifierGAID" api:"required"`
	IdentifierGsfid        string `json:"IdentifierGSFID" api:"required"`
	IdentifierIccid        string `json:"IdentifierICCID" api:"required"`
	IdentifierImei         string `json:"IdentifierIMEI" api:"required"`
	IdentifierImsi         string `json:"IdentifierIMSI" api:"required"`
	IdentifierMediaDrmid   string `json:"IdentifierMediaDRMID" api:"required"`
	IdentifierMeid         string `json:"IdentifierMEID" api:"required"`
	IdentifierPhoneNumber  string `json:"IdentifierPhoneNumber" api:"required"`
	IdentifierSerial       string `json:"IdentifierSerial" api:"required"`
	IdentifierWifiMac      string `json:"IdentifierWifiMAC" api:"required"`
	SerialNumber           string `json:"SerialNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BootloaderSerialNumber respjson.Field
		IdentifierAndroidID    respjson.Field
		IdentifierAppSetID     respjson.Field
		IdentifierBluetoothMac respjson.Field
		IdentifierGaid         respjson.Field
		IdentifierGsfid        respjson.Field
		IdentifierIccid        respjson.Field
		IdentifierImei         respjson.Field
		IdentifierImsi         respjson.Field
		IdentifierMediaDrmid   respjson.Field
		IdentifierMeid         respjson.Field
		IdentifierPhoneNumber  respjson.Field
		IdentifierSerial       respjson.Field
		IdentifierWifiMac      respjson.Field
		SerialNumber           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceIdentifiers) RawJSON() string { return r.JSON.raw }
func (r *DeviceIdentifiers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DeviceIdentifiers to a DeviceIdentifiersParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DeviceIdentifiersParam.Overrides()
func (r DeviceIdentifiers) ToParam() DeviceIdentifiersParam {
	return param.Override[DeviceIdentifiersParam](json.RawMessage(r.RawJSON()))
}

// The properties BootloaderSerialNumber, IdentifierAndroidID, IdentifierAppSetID,
// IdentifierBluetoothMac, IdentifierGaid, IdentifierGsfid, IdentifierIccid,
// IdentifierImei, IdentifierImsi, IdentifierMediaDrmid, IdentifierMeid,
// IdentifierPhoneNumber, IdentifierSerial, IdentifierWifiMac, SerialNumber are
// required.
type DeviceIdentifiersParam struct {
	BootloaderSerialNumber string `json:"BootloaderSerialNumber" api:"required"`
	IdentifierAndroidID    string `json:"IdentifierAndroidID" api:"required"`
	IdentifierAppSetID     string `json:"IdentifierAppSetID" api:"required"`
	IdentifierBluetoothMac string `json:"IdentifierBluetoothMAC" api:"required"`
	IdentifierGaid         string `json:"IdentifierGAID" api:"required"`
	IdentifierGsfid        string `json:"IdentifierGSFID" api:"required"`
	IdentifierIccid        string `json:"IdentifierICCID" api:"required"`
	IdentifierImei         string `json:"IdentifierIMEI" api:"required"`
	IdentifierImsi         string `json:"IdentifierIMSI" api:"required"`
	IdentifierMediaDrmid   string `json:"IdentifierMediaDRMID" api:"required"`
	IdentifierMeid         string `json:"IdentifierMEID" api:"required"`
	IdentifierPhoneNumber  string `json:"IdentifierPhoneNumber" api:"required"`
	IdentifierSerial       string `json:"IdentifierSerial" api:"required"`
	IdentifierWifiMac      string `json:"IdentifierWifiMAC" api:"required"`
	SerialNumber           string `json:"SerialNumber" api:"required"`
	paramObj
}

func (r DeviceIdentifiersParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceIdentifiersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceIdentifiersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSpecParam struct {
	AndroidVersion param.Opt[int64]       `json:"androidVersion,omitzero"`
	Country        param.Opt[string]      `json:"country,omitzero"`
	Locale         param.Opt[string]      `json:"locale,omitzero"`
	Name           param.Opt[string]      `json:"name,omitzero"`
	Timezone       param.Opt[string]      `json:"timezone,omitzero"`
	Apps           []string               `json:"apps,omitzero"`
	Files          []string               `json:"files,omitzero"`
	Carrier        DeviceCarrierParam     `json:"carrier,omitzero"`
	Identifiers    DeviceIdentifiersParam `json:"identifiers,omitzero"`
	Location       LocationParam          `json:"location,omitzero"`
	Proxy          DeviceSpecProxyParam   `json:"proxy,omitzero"`
	paramObj
}

func (r DeviceSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSpecProxyParam struct {
	Name    param.Opt[string]           `json:"name,omitzero"`
	SmartIP param.Opt[bool]             `json:"smartIp,omitzero"`
	Connect DeviceSpecProxyConnectParam `json:"connect,omitzero"`
	Socks5  DeviceSpecProxySocks5Param  `json:"socks5,omitzero"`
	paramObj
}

func (r DeviceSpecProxyParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceSpecProxyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceSpecProxyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceSpecProxyConnectParam struct {
	// Existing Mobilerun Connect proxy id; its credentials are fetched server-side.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// ISO 3166-1 alpha-2 country code; provisions (or reuses) a rotating residential
	// Mobilerun Connect proxy for the device.
	Country param.Opt[string] `json:"country,omitzero"`
	paramObj
}

func (r DeviceSpecProxyConnectParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceSpecProxyConnectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceSpecProxyConnectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Host, Password, Port, User are required.
type DeviceSpecProxySocks5Param struct {
	Host     string `json:"host" api:"required"`
	Password string `json:"password" api:"required"`
	Port     int64  `json:"port" api:"required"`
	User     string `json:"user" api:"required"`
	paramObj
}

func (r DeviceSpecProxySocks5Param) MarshalJSON() (data []byte, err error) {
	type shadow DeviceSpecProxySocks5Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceSpecProxySocks5Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Location struct {
	Latitude  float64 `json:"latitude" api:"required"`
	Longitude float64 `json:"longitude" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Location) RawJSON() string { return r.JSON.raw }
func (r *Location) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Location to a LocationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LocationParam.Overrides()
func (r Location) ToParam() LocationParam {
	return param.Override[LocationParam](json.RawMessage(r.RawJSON()))
}

// The properties Latitude, Longitude are required.
type LocationParam struct {
	Latitude  float64 `json:"latitude" api:"required"`
	Longitude float64 `json:"longitude" api:"required"`
	paramObj
}

func (r LocationParam) MarshalJSON() (data []byte, err error) {
	type shadow LocationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LocationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Meta struct {
	HasNext  bool  `json:"hasNext" api:"required"`
	HasPrev  bool  `json:"hasPrev" api:"required"`
	Page     int64 `json:"page" api:"required"`
	Pages    int64 `json:"pages" api:"required"`
	PageSize int64 `json:"pageSize" api:"required"`
	Total    int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNext     respjson.Field
		HasPrev     respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		PageSize    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Meta) RawJSON() string { return r.JSON.raw }
func (r *Meta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Pagination struct {
	HasNext  bool  `json:"hasNext" api:"required"`
	HasPrev  bool  `json:"hasPrev" api:"required"`
	Page     int64 `json:"page" api:"required"`
	Pages    int64 `json:"pages" api:"required"`
	PageSize int64 `json:"pageSize" api:"required"`
	Total    int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNext     respjson.Field
		HasPrev     respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		PageSize    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pagination) RawJSON() string { return r.JSON.raw }
func (r *Pagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata.
type PaginationMeta struct {
	// Whether there is a next page
	HasNext bool `json:"hasNext" api:"required"`
	// Whether there is a previous page
	HasPrev bool `json:"hasPrev" api:"required"`
	// Current page number (1-based)
	Page int64 `json:"page" api:"required"`
	// Total number of pages
	Pages int64 `json:"pages" api:"required"`
	// Number of items per page
	PageSize int64 `json:"pageSize" api:"required"`
	// Total number of items
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNext     respjson.Field
		HasPrev     respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		PageSize    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *PaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PermissionSet struct {
	Execute bool `json:"execute" api:"required"`
	Read    bool `json:"read" api:"required"`
	Write   bool `json:"write" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Execute     respjson.Field
		Read        respjson.Field
		Write       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PermissionSet) RawJSON() string { return r.JSON.raw }
func (r *PermissionSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Socks5ProxyConfig struct {
	Host     string `json:"host" api:"required"`
	Name     string `json:"name" api:"required"`
	Password string `json:"password" api:"required"`
	Port     int64  `json:"port" api:"required"`
	// Any of "socks5".
	Protocol Socks5ProxyConfigProtocol `json:"protocol" api:"required"`
	ProxyID  string                    `json:"proxyId" api:"required" format:"uuid"`
	User     string                    `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Host        respjson.Field
		Name        respjson.Field
		Password    respjson.Field
		Port        respjson.Field
		Protocol    respjson.Field
		ProxyID     respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Socks5ProxyConfig) RawJSON() string { return r.JSON.raw }
func (r *Socks5ProxyConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Socks5ProxyConfig) ImplProxyNewResponseDataUnion()    {}
func (Socks5ProxyConfig) ImplProxyGetResponseDataUnion()    {}
func (Socks5ProxyConfig) ImplProxyUpdateResponseDataUnion() {}
func (Socks5ProxyConfig) ImplProxyListResponseDataUnion()   {}
func (Socks5ProxyConfig) ImplProxyDeleteResponseDataUnion() {}

type Socks5ProxyConfigProtocol string

const (
	Socks5ProxyConfigProtocolSocks5 Socks5ProxyConfigProtocol = "socks5"
)
