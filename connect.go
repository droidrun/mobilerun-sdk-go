// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// ConnectService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectService] method instead.
type ConnectService struct {
	options []option.RequestOption
	// Mobilerun Connect country coverage information
	Countries ConnectCountryService
	// Manage your Mobilerun Connect Proxies
	Proxies ConnectProxyService
	// Manage your Mobilerun Connect Socks Users
	Users ConnectUserService
}

// NewConnectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConnectService(opts ...option.RequestOption) (r ConnectService) {
	r = ConnectService{}
	r.options = opts
	r.Countries = NewConnectCountryService(opts...)
	r.Proxies = NewConnectProxyService(opts...)
	r.Users = NewConnectUserService(opts...)
	return
}

// A single proxy connection, aggregated across its lifetime. Byte counts are
// totals for the whole connection.
type Connection struct {
	// Total bytes received from upstream over the connection's lifetime.
	BytesIn int64 `json:"bytesIn" api:"required"`
	// Total bytes sent to upstream over the connection's lifetime.
	BytesOut int64 `json:"bytesOut" api:"required"`
	// Upstream country code (ISO 3166-1 alpha-2), or empty if unknown.
	Country string `json:"country" api:"required"`
	// Destination host the client connected to.
	DstHost string `json:"dstHost" api:"required"`
	// Destination port the client connected to.
	DstPort int64 `json:"dstPort" api:"required"`
	// Elapsed time between startedAt and endedAt, in milliseconds.
	DurationMs int64 `json:"durationMs" api:"required"`
	// Time of the connection's last recorded activity (close time once closed).
	EndedAt time.Time `json:"endedAt" api:"required" format:"date-time"`
	// Transport protocol of a connection.
	//
	// Any of "tcp", "udp", "unknown".
	Protocol ConnectionProtocol `json:"protocol" api:"required"`
	// Identifier of the upstream provider that served the connection.
	Provider string `json:"provider" api:"required"`
	// The proxy the connection was routed through. All-zero when the upstream was
	// unresolved at capture time.
	ProxyID string `json:"proxyId" api:"required" format:"uuid"`
	// Unique id of this connection.
	SessionID string `json:"sessionId" api:"required" format:"uuid"`
	// Client source IP address.
	SrcIP string `json:"srcIp" api:"required"`
	// When the connection started.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// `active` while the connection is still open (no terminal record yet); `closed`
	// once it has ended.
	//
	// Any of "active", "closed".
	Status ConnectionStatus `json:"status" api:"required"`
	// bytesIn + bytesOut.
	TotalBytes int64 `json:"totalBytes" api:"required"`
	// The user that made the connection.
	UserID string `json:"userId" api:"required" format:"uuid"`
	// Why the connection closed; null while still active.
	CloseReason string `json:"closeReason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BytesIn     respjson.Field
		BytesOut    respjson.Field
		Country     respjson.Field
		DstHost     respjson.Field
		DstPort     respjson.Field
		DurationMs  respjson.Field
		EndedAt     respjson.Field
		Protocol    respjson.Field
		Provider    respjson.Field
		ProxyID     respjson.Field
		SessionID   respjson.Field
		SrcIP       respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		TotalBytes  respjson.Field
		UserID      respjson.Field
		CloseReason respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Connection) RawJSON() string { return r.JSON.raw }
func (r *Connection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transport protocol of a connection.
type ConnectionProtocol string

const (
	ConnectionProtocolTcp     ConnectionProtocol = "tcp"
	ConnectionProtocolUdp     ConnectionProtocol = "udp"
	ConnectionProtocolUnknown ConnectionProtocol = "unknown"
)

// `active` while the connection is still open (no terminal record yet); `closed`
// once it has ended.
type ConnectionStatus string

const (
	ConnectionStatusActive ConnectionStatus = "active"
	ConnectionStatusClosed ConnectionStatus = "closed"
)
