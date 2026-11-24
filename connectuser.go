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

// Manage your Mobilerun Connect Socks Users
//
// ConnectUserService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectUserService] method instead.
type ConnectUserService struct {
	options []option.RequestOption
}

// NewConnectUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectUserService(opts ...option.RequestOption) (r ConnectUserService) {
	r = ConnectUserService{}
	r.options = opts
	return
}

// Creates a SOCKS5 credential, optionally bound to a proxy for dedicated routing.
// Username and password are generated when omitted.
func (r *ConnectUserService) New(ctx context.Context, body ConnectUserNewParams, opts ...option.RequestOption) (res *ConnectUserNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "connect/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the SOCKS5 user identified by the path ID. The response includes the
// user's password.
func (r *ConnectUserService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ConnectUserGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connect/users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rebind the user to a different proxy (or detach it by passing null).
func (r *ConnectUserService) Update(ctx context.Context, id string, body ConnectUserUpdateParams, opts ...option.RequestOption) (res *ConnectUserUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connect/users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns SOCKS5 users owned by the caller. Passwords are omitted from the list.
func (r *ConnectUserService) List(ctx context.Context, query ConnectUserListParams, opts ...option.RequestOption) (res *ConnectUserListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "connect/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes the SOCKS5 user identified by the path ID, revoking its credentials and
// any proxy binding. Returns 404 if no such user exists for the caller.
func (r *ConnectUserService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("connect/users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns the connection history recorded for this user, one item per connection
// (aggregated across the connection's lifetime). Supports filtering on every
// property plus ordering and pagination. Returns 503 when connection history is
// temporarily unavailable.
func (r *ConnectUserService) ListConnections(ctx context.Context, id string, query ConnectUserListConnectionsParams, opts ...option.RequestOption) (res *ConnectUserListConnectionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connect/users/%s/connections", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A SOCKS5 user including its password. Returned only on create and single-user
// reads.
type ConnectUserNewResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Password  string    `json:"password" api:"required"`
	Username  string    `json:"username" api:"required"`
	// The proxy this user routes to (dedicated routing), or null if unbound.
	ProxyID string `json:"proxyId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Password    respjson.Field
		Username    respjson.Field
		ProxyID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectUserNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectUserNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A SOCKS5 user including its password. Returned only on create and single-user
// reads.
type ConnectUserGetResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Password  string    `json:"password" api:"required"`
	Username  string    `json:"username" api:"required"`
	// The proxy this user routes to (dedicated routing), or null if unbound.
	ProxyID string `json:"proxyId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Password    respjson.Field
		Username    respjson.Field
		ProxyID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectUserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectUserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A SOCKS5 credential without its password.
type ConnectUserUpdateResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Username  string    `json:"username" api:"required"`
	// The proxy this user routes to (dedicated routing), or null if unbound.
	ProxyID string `json:"proxyId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Username    respjson.Field
		ProxyID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectUserUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectUserUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page of SOCKS5 users.
type ConnectUserListResponse struct {
	Items []ConnectUserListResponseItem `json:"items" api:"required"`
	// Pagination metadata for a list response.
	Pagination ConnectUserListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectUserListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectUserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A SOCKS5 credential without its password.
type ConnectUserListResponseItem struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Username  string    `json:"username" api:"required"`
	// The proxy this user routes to (dedicated routing), or null if unbound.
	ProxyID string `json:"proxyId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Username    respjson.Field
		ProxyID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectUserListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ConnectUserListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata for a list response.
type ConnectUserListResponsePagination struct {
	// Whether a next page exists.
	HasNext bool `json:"hasNext" api:"required"`
	// Whether a previous page exists.
	HasPrev bool `json:"hasPrev" api:"required"`
	// Current page number (1-based).
	Page int64 `json:"page" api:"required"`
	// Total number of pages.
	Pages int64 `json:"pages" api:"required"`
	// Number of items per page.
	PageSize int64 `json:"pageSize" api:"required"`
	// Total number of items across all pages.
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
func (r ConnectUserListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ConnectUserListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page of connections.
type ConnectUserListConnectionsResponse struct {
	Items []ConnectUserListConnectionsResponseItem `json:"items" api:"required"`
	// Pagination metadata for a list response.
	Pagination ConnectUserListConnectionsResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectUserListConnectionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectUserListConnectionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single proxy connection, aggregated across its lifetime. Byte counts are
// totals for the whole connection.
type ConnectUserListConnectionsResponseItem struct {
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
	Protocol string `json:"protocol" api:"required"`
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
	Status string `json:"status" api:"required"`
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
func (r ConnectUserListConnectionsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ConnectUserListConnectionsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata for a list response.
type ConnectUserListConnectionsResponsePagination struct {
	// Whether a next page exists.
	HasNext bool `json:"hasNext" api:"required"`
	// Whether a previous page exists.
	HasPrev bool `json:"hasPrev" api:"required"`
	// Current page number (1-based).
	Page int64 `json:"page" api:"required"`
	// Total number of pages.
	Pages int64 `json:"pages" api:"required"`
	// Number of items per page.
	PageSize int64 `json:"pageSize" api:"required"`
	// Total number of items across all pages.
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
func (r ConnectUserListConnectionsResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ConnectUserListConnectionsResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectUserNewParams struct {
	// Desired SOCKS5 password, 1-255 bytes (RFC 1929). Generated when omitted.
	Password param.Opt[string] `json:"password,omitzero"`
	// Proxy to bind the user to for dedicated routing.
	ProxyID param.Opt[string] `json:"proxyId,omitzero" format:"uuid"`
	// Desired SOCKS5 username, 1-255 bytes (RFC 1929). Generated when omitted.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r ConnectUserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConnectUserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectUserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectUserUpdateParams struct {
	// Proxy to rebind to, or null to detach. Omit to leave the user's current binding
	// unchanged — only an explicit null detaches.
	ProxyID param.Opt[string] `json:"proxyId,omitzero" format:"uuid"`
	paramObj
}

func (r ConnectUserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConnectUserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectUserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectUserListParams struct {
	// Page number (1-based).
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Filter to users bound to this proxy. Users not bound to it (including unbound
	// users) are excluded.
	ProxyID param.Opt[string] `query:"proxyId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectUserListParams]'s query parameters as `url.Values`.
func (r ConnectUserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectUserListConnectionsParams struct {
	// Filter to connections that closed with this reason (closed connections only).
	CloseReason param.Opt[string] `query:"closeReason,omitzero" json:"-"`
	// Filter to connections served from this upstream country (ISO 3166-1 alpha-2).
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Filter to connections to this destination host (exact match).
	DstHost param.Opt[string] `query:"dstHost,omitzero" json:"-"`
	// Filter to connections to this destination port.
	DstPort param.Opt[int64] `query:"dstPort,omitzero" json:"-"`
	// Filter to connections whose last activity was at or after this time (inclusive).
	EndedAfter param.Opt[time.Time] `query:"endedAfter,omitzero" format:"date-time" json:"-"`
	// Filter to connections whose last activity was at or before this time
	// (inclusive).
	EndedBefore param.Opt[time.Time] `query:"endedBefore,omitzero" format:"date-time" json:"-"`
	// Filter to connections with at most this many bytes received from upstream.
	MaxBytesIn param.Opt[int64] `query:"maxBytesIn,omitzero" json:"-"`
	// Filter to connections with at most this many bytes sent to upstream.
	MaxBytesOut param.Opt[int64] `query:"maxBytesOut,omitzero" json:"-"`
	// Filter to connections lasting at most this many milliseconds.
	MaxDurationMs param.Opt[int64] `query:"maxDurationMs,omitzero" json:"-"`
	// Filter to connections with at most this much total traffic (bytesIn + bytesOut).
	MaxTotalBytes param.Opt[int64] `query:"maxTotalBytes,omitzero" json:"-"`
	// Filter to connections with at least this many bytes received from upstream.
	MinBytesIn param.Opt[int64] `query:"minBytesIn,omitzero" json:"-"`
	// Filter to connections with at least this many bytes sent to upstream.
	MinBytesOut param.Opt[int64] `query:"minBytesOut,omitzero" json:"-"`
	// Filter to connections lasting at least this many milliseconds.
	MinDurationMs param.Opt[int64] `query:"minDurationMs,omitzero" json:"-"`
	// Filter to connections with at least this much total traffic (bytesIn +
	// bytesOut).
	MinTotalBytes param.Opt[int64] `query:"minTotalBytes,omitzero" json:"-"`
	// Page number (1-based).
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Filter to connections served by this upstream provider identifier.
	Provider param.Opt[string] `query:"provider,omitzero" json:"-"`
	// Filter to connections routed through this proxy.
	ProxyID param.Opt[string] `query:"proxyId,omitzero" format:"uuid" json:"-"`
	// Filter to a single connection by its session id.
	SessionID param.Opt[string] `query:"sessionId,omitzero" format:"uuid" json:"-"`
	// Filter to connections that started at or after this time (inclusive).
	StartedAfter param.Opt[time.Time] `query:"startedAfter,omitzero" format:"date-time" json:"-"`
	// Filter to connections that started at or before this time (inclusive).
	StartedBefore param.Opt[time.Time] `query:"startedBefore,omitzero" format:"date-time" json:"-"`
	// Sort direction.
	//
	// Any of "asc", "desc".
	Order ConnectUserListConnectionsParamsOrder `query:"order,omitzero" json:"-"`
	// Property to order the results by.
	//
	// Any of "startedAt", "endedAt", "bytesIn", "bytesOut", "totalBytes",
	// "durationMs".
	OrderBy ConnectUserListConnectionsParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// Filter to connections of this transport protocol.
	//
	// Any of "tcp", "udp", "unknown".
	Protocol ConnectUserListConnectionsParamsProtocol `query:"protocol,omitzero" json:"-"`
	// Filter by connection status.
	//
	// Any of "active", "closed".
	Status ConnectUserListConnectionsParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectUserListConnectionsParams]'s query parameters as
// `url.Values`.
func (r ConnectUserListConnectionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction.
type ConnectUserListConnectionsParamsOrder string

const (
	ConnectUserListConnectionsParamsOrderAsc  ConnectUserListConnectionsParamsOrder = "asc"
	ConnectUserListConnectionsParamsOrderDesc ConnectUserListConnectionsParamsOrder = "desc"
)

// Property to order the results by.
type ConnectUserListConnectionsParamsOrderBy string

const (
	ConnectUserListConnectionsParamsOrderByStartedAt  ConnectUserListConnectionsParamsOrderBy = "startedAt"
	ConnectUserListConnectionsParamsOrderByEndedAt    ConnectUserListConnectionsParamsOrderBy = "endedAt"
	ConnectUserListConnectionsParamsOrderByBytesIn    ConnectUserListConnectionsParamsOrderBy = "bytesIn"
	ConnectUserListConnectionsParamsOrderByBytesOut   ConnectUserListConnectionsParamsOrderBy = "bytesOut"
	ConnectUserListConnectionsParamsOrderByTotalBytes ConnectUserListConnectionsParamsOrderBy = "totalBytes"
	ConnectUserListConnectionsParamsOrderByDurationMs ConnectUserListConnectionsParamsOrderBy = "durationMs"
)

// Filter to connections of this transport protocol.
type ConnectUserListConnectionsParamsProtocol string

const (
	ConnectUserListConnectionsParamsProtocolTcp     ConnectUserListConnectionsParamsProtocol = "tcp"
	ConnectUserListConnectionsParamsProtocolUdp     ConnectUserListConnectionsParamsProtocol = "udp"
	ConnectUserListConnectionsParamsProtocolUnknown ConnectUserListConnectionsParamsProtocol = "unknown"
)

// Filter by connection status.
type ConnectUserListConnectionsParamsStatus string

const (
	ConnectUserListConnectionsParamsStatusActive ConnectUserListConnectionsParamsStatus = "active"
	ConnectUserListConnectionsParamsStatusClosed ConnectUserListConnectionsParamsStatus = "closed"
)
