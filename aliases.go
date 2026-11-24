// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apierror"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type DeviceCarrier = shared.DeviceCarrier

// This is an alias to an internal type.
type DeviceCarrierParam = shared.DeviceCarrierParam

// This is an alias to an internal type.
type DeviceIdentifiers = shared.DeviceIdentifiers

// This is an alias to an internal type.
type DeviceIdentifiersParam = shared.DeviceIdentifiersParam

// This is an alias to an internal type.
type DeviceSpecParam = shared.DeviceSpecParam

// This is an alias to an internal type.
type DeviceSpecProxyParam = shared.DeviceSpecProxyParam

// This is an alias to an internal type.
type DeviceSpecProxyConnectParam = shared.DeviceSpecProxyConnectParam

// This is an alias to an internal type.
type DeviceSpecProxySocks5Param = shared.DeviceSpecProxySocks5Param

// This is an alias to an internal type.
type Location = shared.Location

// This is an alias to an internal type.
type LocationParam = shared.LocationParam

// This is an alias to an internal type.
type Meta = shared.Meta

// This is an alias to an internal type.
type Pagination = shared.Pagination

// Pagination metadata.
//
// This is an alias to an internal type.
type PaginationMeta = shared.PaginationMeta

// This is an alias to an internal type.
type PermissionSet = shared.PermissionSet

// This is an alias to an internal type.
type Socks5ProxyConfig = shared.Socks5ProxyConfig

// This is an alias to an internal type.
type Socks5ProxyConfigProtocol = shared.Socks5ProxyConfigProtocol

// Equals "socks5"
const Socks5ProxyConfigProtocolSocks5 = shared.Socks5ProxyConfigProtocolSocks5
