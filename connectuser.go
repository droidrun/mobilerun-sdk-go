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
func (r *ConnectUserService) List(ctx context.Context, query ConnectUserListParams, opts ...option.RequestOption) (res *pagination.Page[ConnectUserListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "connect/users"
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

// Returns SOCKS5 users owned by the caller. Passwords are omitted from the list.
func (r *ConnectUserService) ListAutoPaging(ctx context.Context, query ConnectUserListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ConnectUserListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
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
func (r *ConnectUserService) ListConnections(ctx context.Context, id string, query ConnectUserListConnectionsParams, opts ...option.RequestOption) (res *pagination.Page[Connection], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connect/users/%s/connections", id)
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

// Returns the connection history recorded for this user, one item per connection
// (aggregated across the connection's lifetime). Supports filtering on every
// property plus ordering and pagination. Returns 503 when connection history is
// temporarily unavailable.
func (r *ConnectUserService) ListConnectionsAutoPaging(ctx context.Context, id string, query ConnectUserListConnectionsParams, opts ...option.RequestOption) *pagination.PageAutoPager[Connection] {
	return pagination.NewPageAutoPager(r.ListConnections(ctx, id, query, opts...))
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

// A SOCKS5 credential without its password.
type ConnectUserListResponse struct {
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
func (r ConnectUserListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectUserListResponse) UnmarshalJSON(data []byte) error {
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
