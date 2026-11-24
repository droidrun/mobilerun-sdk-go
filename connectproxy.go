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

// Manage your Mobilerun Connect Proxies
//
// ConnectProxyService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectProxyService] method instead.
type ConnectProxyService struct {
	options []option.RequestOption
}

// NewConnectProxyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectProxyService(opts ...option.RequestOption) (r ConnectProxyService) {
	r = ConnectProxyService{}
	r.options = opts
	return
}

// Returns the proxy identified by the path ID. The response includes the proxy's
// password.
func (r *ConnectProxyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ConnectProxyGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connect/proxies/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Renames the proxy identified by the path ID.
func (r *ConnectProxyService) Update(ctx context.Context, id string, body ConnectProxyUpdateParams, opts ...option.RequestOption) (res *ConnectProxyUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connect/proxies/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns proxies owned by the caller. Credentials are omitted from the list.
func (r *ConnectProxyService) List(ctx context.Context, query ConnectProxyListParams, opts ...option.RequestOption) (res *ConnectProxyListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "connect/proxies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Provisions a proxy of the requested type for the caller in the selected country.
func (r *ConnectProxyService) Buy(ctx context.Context, body ConnectProxyBuyParams, opts ...option.RequestOption) (res *ConnectProxyBuyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "connect/proxies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Deletes the proxy identified by the path ID and releases its provisioning.
// Returns 404 if no such proxy exists for the caller.
func (r *ConnectProxyService) Cancel(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("connect/proxies/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Refreshes a pending dedicated proxy's checkout on explicit payment intent.
// Already-paid proxies (provisioning, active, or cancelling) return their current
// status without a payment URL. An expired checkout requires a new purchase.
func (r *ConnectProxyService) Checkout(ctx context.Context, id string, opts ...option.RequestOption) (res *ConnectProxyCheckoutResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connect/proxies/%s/checkout", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Checks right now that the proxy's upstream accepts a connection through its
// current session. If the upstream refuses the session and the proxy is a rotating
// (sticky-session) proxy, the proxy is moved onto a fresh session that has been
// checked to work; its credentials and country stay the same, its exit IP changes.
// Intended for callers that are about to depend on the proxy (e.g. while
// provisioning a phone). After `rotated`, the gateway uses the new session within
// about 30 seconds. A `refused` or `unreachable` result is not an error: the
// response describes the proxy's state.
func (r *ConnectProxyService) EnsureHealthy(ctx context.Context, id string, opts ...option.RequestOption) (res *ConnectProxyEnsureHealthyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connect/proxies/%s/ensure-healthy", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the connection history recorded for this proxy, one item per connection
// (aggregated across the connection's lifetime). Supports filtering on every
// property plus ordering and pagination. Returns 503 when connection history is
// temporarily unavailable.
func (r *ConnectProxyService) ListConnections(ctx context.Context, id string, query ConnectProxyListConnectionsParams, opts ...option.RequestOption) (res *ConnectProxyListConnectionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connect/proxies/%s/connections", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the most recent cached network-latency measurement for the proxy,
// sampled periodically by connecting through the proxy to a fixed target.
// `latency` is null when no measurement is available yet (e.g. the proxy is not
// active, or it has not been sampled since coming online).
func (r *ConnectProxyService) Ping(ctx context.Context, id string, opts ...option.RequestOption) (res *ConnectProxyPingResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connect/proxies/%s/ping", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A proxy including its password. Returned only on create and single-proxy reads.
type ConnectProxyGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// ISO 3166-1 alpha-2 country code (lowercase).
	Country   string    `json:"country" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Host      string    `json:"host" api:"required"`
	// Display name for the proxy. Set it on create or change it with PATCH; when no
	// name has been set, a label generated from the proxy's country, type, and id is
	// returned instead — so this is never empty.
	Name     string `json:"name" api:"required"`
	Password string `json:"password" api:"required"`
	Port     int64  `json:"port" api:"required"`
	// Lifecycle of a proxy. A freshly created proxy is `checking` while its billing
	// identity is being resolved — clients should poll until a `paymentUrl` or a later
	// status appears — then `provisioning` — or `pending_payment` until the customer
	// completes checkout — and becomes `active` once its upstream is assigned.
	// `cancelling` retains full access through the paid period; when the subscription
	// expires the proxy is `ended`. `error` marks a failed provisioning attempt.
	//
	// Any of "checking", "pending_payment", "provisioning", "active", "cancelling",
	// "ended", "error".
	Status ConnectProxyGetResponseStatus `json:"status" api:"required"`
	// Any of "dedicated_residential", "residential", "mobile".
	Type     ConnectProxyGetResponseType `json:"type" api:"required"`
	Username string                      `json:"username" api:"required"`
	// Any of "included", "standalone_paid".
	BillingMode ConnectProxyGetResponseBillingMode `json:"billingMode"`
	// Checkout URL to complete payment while status is `pending_payment`. Null once
	// paid or when no payment was required.
	PaymentURL string `json:"paymentUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Country     respjson.Field
		CreatedAt   respjson.Field
		Host        respjson.Field
		Name        respjson.Field
		Password    respjson.Field
		Port        respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Username    respjson.Field
		BillingMode respjson.Field
		PaymentURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectProxyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle of a proxy. A freshly created proxy is `checking` while its billing
// identity is being resolved — clients should poll until a `paymentUrl` or a later
// status appears — then `provisioning` — or `pending_payment` until the customer
// completes checkout — and becomes `active` once its upstream is assigned.
// `cancelling` retains full access through the paid period; when the subscription
// expires the proxy is `ended`. `error` marks a failed provisioning attempt.
type ConnectProxyGetResponseStatus string

const (
	ConnectProxyGetResponseStatusChecking       ConnectProxyGetResponseStatus = "checking"
	ConnectProxyGetResponseStatusPendingPayment ConnectProxyGetResponseStatus = "pending_payment"
	ConnectProxyGetResponseStatusProvisioning   ConnectProxyGetResponseStatus = "provisioning"
	ConnectProxyGetResponseStatusActive         ConnectProxyGetResponseStatus = "active"
	ConnectProxyGetResponseStatusCancelling     ConnectProxyGetResponseStatus = "cancelling"
	ConnectProxyGetResponseStatusEnded          ConnectProxyGetResponseStatus = "ended"
	ConnectProxyGetResponseStatusError          ConnectProxyGetResponseStatus = "error"
)

type ConnectProxyGetResponseType string

const (
	ConnectProxyGetResponseTypeDedicatedResidential ConnectProxyGetResponseType = "dedicated_residential"
	ConnectProxyGetResponseTypeResidential          ConnectProxyGetResponseType = "residential"
	ConnectProxyGetResponseTypeMobile               ConnectProxyGetResponseType = "mobile"
)

type ConnectProxyGetResponseBillingMode string

const (
	ConnectProxyGetResponseBillingModeIncluded       ConnectProxyGetResponseBillingMode = "included"
	ConnectProxyGetResponseBillingModeStandalonePaid ConnectProxyGetResponseBillingMode = "standalone_paid"
)

// A provisioned proxy without its credentials.
type ConnectProxyUpdateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// ISO 3166-1 alpha-2 country code (lowercase).
	Country   string    `json:"country" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Host      string    `json:"host" api:"required"`
	// Display name for the proxy. Set it on create or change it with PATCH; when no
	// name has been set, a label generated from the proxy's country, type, and id is
	// returned instead — so this is never empty.
	Name string `json:"name" api:"required"`
	Port int64  `json:"port" api:"required"`
	// Lifecycle of a proxy. A freshly created proxy is `checking` while its billing
	// identity is being resolved — clients should poll until a `paymentUrl` or a later
	// status appears — then `provisioning` — or `pending_payment` until the customer
	// completes checkout — and becomes `active` once its upstream is assigned.
	// `cancelling` retains full access through the paid period; when the subscription
	// expires the proxy is `ended`. `error` marks a failed provisioning attempt.
	//
	// Any of "checking", "pending_payment", "provisioning", "active", "cancelling",
	// "ended", "error".
	Status ConnectProxyUpdateResponseStatus `json:"status" api:"required"`
	// Any of "dedicated_residential", "residential", "mobile".
	Type     ConnectProxyUpdateResponseType `json:"type" api:"required"`
	Username string                         `json:"username" api:"required"`
	// Any of "included", "standalone_paid".
	BillingMode ConnectProxyUpdateResponseBillingMode `json:"billingMode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Country     respjson.Field
		CreatedAt   respjson.Field
		Host        respjson.Field
		Name        respjson.Field
		Port        respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Username    respjson.Field
		BillingMode respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectProxyUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle of a proxy. A freshly created proxy is `checking` while its billing
// identity is being resolved — clients should poll until a `paymentUrl` or a later
// status appears — then `provisioning` — or `pending_payment` until the customer
// completes checkout — and becomes `active` once its upstream is assigned.
// `cancelling` retains full access through the paid period; when the subscription
// expires the proxy is `ended`. `error` marks a failed provisioning attempt.
type ConnectProxyUpdateResponseStatus string

const (
	ConnectProxyUpdateResponseStatusChecking       ConnectProxyUpdateResponseStatus = "checking"
	ConnectProxyUpdateResponseStatusPendingPayment ConnectProxyUpdateResponseStatus = "pending_payment"
	ConnectProxyUpdateResponseStatusProvisioning   ConnectProxyUpdateResponseStatus = "provisioning"
	ConnectProxyUpdateResponseStatusActive         ConnectProxyUpdateResponseStatus = "active"
	ConnectProxyUpdateResponseStatusCancelling     ConnectProxyUpdateResponseStatus = "cancelling"
	ConnectProxyUpdateResponseStatusEnded          ConnectProxyUpdateResponseStatus = "ended"
	ConnectProxyUpdateResponseStatusError          ConnectProxyUpdateResponseStatus = "error"
)

type ConnectProxyUpdateResponseType string

const (
	ConnectProxyUpdateResponseTypeDedicatedResidential ConnectProxyUpdateResponseType = "dedicated_residential"
	ConnectProxyUpdateResponseTypeResidential          ConnectProxyUpdateResponseType = "residential"
	ConnectProxyUpdateResponseTypeMobile               ConnectProxyUpdateResponseType = "mobile"
)

type ConnectProxyUpdateResponseBillingMode string

const (
	ConnectProxyUpdateResponseBillingModeIncluded       ConnectProxyUpdateResponseBillingMode = "included"
	ConnectProxyUpdateResponseBillingModeStandalonePaid ConnectProxyUpdateResponseBillingMode = "standalone_paid"
)

// A page of proxies.
type ConnectProxyListResponse struct {
	Items []ConnectProxyListResponseItem `json:"items" api:"required"`
	// Pagination metadata for a list response.
	Pagination ConnectProxyListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectProxyListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A provisioned proxy without its credentials.
type ConnectProxyListResponseItem struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// ISO 3166-1 alpha-2 country code (lowercase).
	Country   string    `json:"country" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Host      string    `json:"host" api:"required"`
	// Display name for the proxy. Set it on create or change it with PATCH; when no
	// name has been set, a label generated from the proxy's country, type, and id is
	// returned instead — so this is never empty.
	Name string `json:"name" api:"required"`
	Port int64  `json:"port" api:"required"`
	// Lifecycle of a proxy. A freshly created proxy is `checking` while its billing
	// identity is being resolved — clients should poll until a `paymentUrl` or a later
	// status appears — then `provisioning` — or `pending_payment` until the customer
	// completes checkout — and becomes `active` once its upstream is assigned.
	// `cancelling` retains full access through the paid period; when the subscription
	// expires the proxy is `ended`. `error` marks a failed provisioning attempt.
	//
	// Any of "checking", "pending_payment", "provisioning", "active", "cancelling",
	// "ended", "error".
	Status string `json:"status" api:"required"`
	// Any of "dedicated_residential", "residential", "mobile".
	Type     string `json:"type" api:"required"`
	Username string `json:"username" api:"required"`
	// Any of "included", "standalone_paid".
	BillingMode string `json:"billingMode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Country     respjson.Field
		CreatedAt   respjson.Field
		Host        respjson.Field
		Name        respjson.Field
		Port        respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Username    respjson.Field
		BillingMode respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectProxyListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata for a list response.
type ConnectProxyListResponsePagination struct {
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
func (r ConnectProxyListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A proxy including its password. Returned only on create and single-proxy reads.
type ConnectProxyBuyResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// ISO 3166-1 alpha-2 country code (lowercase).
	Country   string    `json:"country" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Host      string    `json:"host" api:"required"`
	// Display name for the proxy. Set it on create or change it with PATCH; when no
	// name has been set, a label generated from the proxy's country, type, and id is
	// returned instead — so this is never empty.
	Name     string `json:"name" api:"required"`
	Password string `json:"password" api:"required"`
	Port     int64  `json:"port" api:"required"`
	// Lifecycle of a proxy. A freshly created proxy is `checking` while its billing
	// identity is being resolved — clients should poll until a `paymentUrl` or a later
	// status appears — then `provisioning` — or `pending_payment` until the customer
	// completes checkout — and becomes `active` once its upstream is assigned.
	// `cancelling` retains full access through the paid period; when the subscription
	// expires the proxy is `ended`. `error` marks a failed provisioning attempt.
	//
	// Any of "checking", "pending_payment", "provisioning", "active", "cancelling",
	// "ended", "error".
	Status ConnectProxyBuyResponseStatus `json:"status" api:"required"`
	// Any of "dedicated_residential", "residential", "mobile".
	Type     ConnectProxyBuyResponseType `json:"type" api:"required"`
	Username string                      `json:"username" api:"required"`
	// Any of "included", "standalone_paid".
	BillingMode ConnectProxyBuyResponseBillingMode `json:"billingMode"`
	// Checkout URL to complete payment while status is `pending_payment`. Null once
	// paid or when no payment was required.
	PaymentURL string `json:"paymentUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Country     respjson.Field
		CreatedAt   respjson.Field
		Host        respjson.Field
		Name        respjson.Field
		Password    respjson.Field
		Port        respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Username    respjson.Field
		BillingMode respjson.Field
		PaymentURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectProxyBuyResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyBuyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle of a proxy. A freshly created proxy is `checking` while its billing
// identity is being resolved — clients should poll until a `paymentUrl` or a later
// status appears — then `provisioning` — or `pending_payment` until the customer
// completes checkout — and becomes `active` once its upstream is assigned.
// `cancelling` retains full access through the paid period; when the subscription
// expires the proxy is `ended`. `error` marks a failed provisioning attempt.
type ConnectProxyBuyResponseStatus string

const (
	ConnectProxyBuyResponseStatusChecking       ConnectProxyBuyResponseStatus = "checking"
	ConnectProxyBuyResponseStatusPendingPayment ConnectProxyBuyResponseStatus = "pending_payment"
	ConnectProxyBuyResponseStatusProvisioning   ConnectProxyBuyResponseStatus = "provisioning"
	ConnectProxyBuyResponseStatusActive         ConnectProxyBuyResponseStatus = "active"
	ConnectProxyBuyResponseStatusCancelling     ConnectProxyBuyResponseStatus = "cancelling"
	ConnectProxyBuyResponseStatusEnded          ConnectProxyBuyResponseStatus = "ended"
	ConnectProxyBuyResponseStatusError          ConnectProxyBuyResponseStatus = "error"
)

type ConnectProxyBuyResponseType string

const (
	ConnectProxyBuyResponseTypeDedicatedResidential ConnectProxyBuyResponseType = "dedicated_residential"
	ConnectProxyBuyResponseTypeResidential          ConnectProxyBuyResponseType = "residential"
	ConnectProxyBuyResponseTypeMobile               ConnectProxyBuyResponseType = "mobile"
)

type ConnectProxyBuyResponseBillingMode string

const (
	ConnectProxyBuyResponseBillingModeIncluded       ConnectProxyBuyResponseBillingMode = "included"
	ConnectProxyBuyResponseBillingModeStandalonePaid ConnectProxyBuyResponseBillingMode = "standalone_paid"
)

type ConnectProxyCheckoutResponse struct {
	// Lifecycle of a proxy. A freshly created proxy is `checking` while its billing
	// identity is being resolved — clients should poll until a `paymentUrl` or a later
	// status appears — then `provisioning` — or `pending_payment` until the customer
	// completes checkout — and becomes `active` once its upstream is assigned.
	// `cancelling` retains full access through the paid period; when the subscription
	// expires the proxy is `ended`. `error` marks a failed provisioning attempt.
	//
	// Any of "checking", "pending_payment", "provisioning", "active", "cancelling",
	// "ended", "error".
	Status ConnectProxyCheckoutResponseStatus `json:"status" api:"required"`
	// Payable checkout URL; omitted after payment.
	PaymentURL string `json:"paymentUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		PaymentURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectProxyCheckoutResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyCheckoutResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle of a proxy. A freshly created proxy is `checking` while its billing
// identity is being resolved — clients should poll until a `paymentUrl` or a later
// status appears — then `provisioning` — or `pending_payment` until the customer
// completes checkout — and becomes `active` once its upstream is assigned.
// `cancelling` retains full access through the paid period; when the subscription
// expires the proxy is `ended`. `error` marks a failed provisioning attempt.
type ConnectProxyCheckoutResponseStatus string

const (
	ConnectProxyCheckoutResponseStatusChecking       ConnectProxyCheckoutResponseStatus = "checking"
	ConnectProxyCheckoutResponseStatusPendingPayment ConnectProxyCheckoutResponseStatus = "pending_payment"
	ConnectProxyCheckoutResponseStatusProvisioning   ConnectProxyCheckoutResponseStatus = "provisioning"
	ConnectProxyCheckoutResponseStatusActive         ConnectProxyCheckoutResponseStatus = "active"
	ConnectProxyCheckoutResponseStatusCancelling     ConnectProxyCheckoutResponseStatus = "cancelling"
	ConnectProxyCheckoutResponseStatusEnded          ConnectProxyCheckoutResponseStatus = "ended"
	ConnectProxyCheckoutResponseStatusError          ConnectProxyCheckoutResponseStatus = "error"
)

// The outcome of an ensure-healthy check.
type ConnectProxyEnsureHealthyResponse struct {
	ProxyID string `json:"proxyId" api:"required" format:"uuid"`
	// `healthy`: the upstream accepted the current session. `rotated`: the session was
	// refused and the proxy now uses a fresh, checked session; retry through the
	// proxy. `refused`: the session was refused and could not be replaced.
	// `unreachable`: the upstream could not be reached; the session was left
	// unchanged.
	//
	// Any of "healthy", "rotated", "refused", "unreachable".
	Status ConnectProxyEnsureHealthyResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProxyID     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectProxyEnsureHealthyResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyEnsureHealthyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `healthy`: the upstream accepted the current session. `rotated`: the session was
// refused and the proxy now uses a fresh, checked session; retry through the
// proxy. `refused`: the session was refused and could not be replaced.
// `unreachable`: the upstream could not be reached; the session was left
// unchanged.
type ConnectProxyEnsureHealthyResponseStatus string

const (
	ConnectProxyEnsureHealthyResponseStatusHealthy     ConnectProxyEnsureHealthyResponseStatus = "healthy"
	ConnectProxyEnsureHealthyResponseStatusRotated     ConnectProxyEnsureHealthyResponseStatus = "rotated"
	ConnectProxyEnsureHealthyResponseStatusRefused     ConnectProxyEnsureHealthyResponseStatus = "refused"
	ConnectProxyEnsureHealthyResponseStatusUnreachable ConnectProxyEnsureHealthyResponseStatus = "unreachable"
)

// A page of connections.
type ConnectProxyListConnectionsResponse struct {
	Items []ConnectProxyListConnectionsResponseItem `json:"items" api:"required"`
	// Pagination metadata for a list response.
	Pagination ConnectProxyListConnectionsResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectProxyListConnectionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyListConnectionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single proxy connection, aggregated across its lifetime. Byte counts are
// totals for the whole connection.
type ConnectProxyListConnectionsResponseItem struct {
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
func (r ConnectProxyListConnectionsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyListConnectionsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata for a list response.
type ConnectProxyListConnectionsResponsePagination struct {
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
func (r ConnectProxyListConnectionsResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyListConnectionsResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The latest cached latency reading for a proxy.
type ConnectProxyPingResponse struct {
	// An aggregated latency measurement taken through the proxy.
	Latency ConnectProxyPingResponseLatency `json:"latency" api:"required"`
	ProxyID string                          `json:"proxyId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latency     respjson.Field
		ProxyID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectProxyPingResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyPingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An aggregated latency measurement taken through the proxy.
type ConnectProxyPingResponseLatency struct {
	// Mean round-trip time over successful probes, in milliseconds.
	AvgMs int64 `json:"avgMs" api:"required"`
	// Round-trip time spread (max - min) over successful probes, in milliseconds.
	JitterMs int64 `json:"jitterMs" api:"required"`
	// Maximum round-trip time over successful probes, in milliseconds.
	MaxMs int64 `json:"maxMs" api:"required"`
	// When this measurement was taken.
	MeasuredAt time.Time `json:"measuredAt" api:"required" format:"date-time"`
	// Minimum round-trip time over successful probes, in milliseconds.
	MinMs int64 `json:"minMs" api:"required"`
	// Fraction of probes that failed, 0..1 (1 means the proxy was unreachable; rtt
	// fields are 0).
	PacketLoss float64 `json:"packetLoss" api:"required"`
	// Number of probes taken in this measurement.
	Samples int64 `json:"samples" api:"required"`
	// The host:port the latency was measured against, through the proxy.
	Target string `json:"target" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgMs       respjson.Field
		JitterMs    respjson.Field
		MaxMs       respjson.Field
		MeasuredAt  respjson.Field
		MinMs       respjson.Field
		PacketLoss  respjson.Field
		Samples     respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectProxyPingResponseLatency) RawJSON() string { return r.JSON.raw }
func (r *ConnectProxyPingResponseLatency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectProxyUpdateParams struct {
	// New display name, up to 64 characters excluding surrounding whitespace, and
	// containing no NUL. Send null (or an empty/whitespace-only string) to drop a
	// custom name and go back to the generated label. Omit to leave the name
	// unchanged.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ConnectProxyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConnectProxyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectProxyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectProxyListParams struct {
	// Filter to proxies in this country (ISO 3166-1 alpha-2, lowercase).
	Country param.Opt[string] `query:"country,omitzero" format:"country-code" json:"-"`
	// Page number (1-based).
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Filter to proxies of these types. Repeat the parameter for several types
	// (`?type=a&type=b`).
	//
	// Any of "dedicated_residential", "residential", "mobile".
	Type []string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectProxyListParams]'s query parameters as `url.Values`.
func (r ConnectProxyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectProxyBuyParams struct {
	// ISO 3166-1 alpha-2 country code to provision the proxy in.
	Country string `json:"country" api:"required" format:"country-code"`
	// Any of "dedicated_residential", "residential", "mobile".
	Type ConnectProxyBuyParamsType `json:"type,omitzero" api:"required"`
	// Display name for the proxy, up to 64 characters excluding surrounding
	// whitespace, and containing no NUL. Omit it (or send only whitespace) to get a
	// generated label built from the country, type, and id.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ConnectProxyBuyParams) MarshalJSON() (data []byte, err error) {
	type shadow ConnectProxyBuyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectProxyBuyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectProxyBuyParamsType string

const (
	ConnectProxyBuyParamsTypeDedicatedResidential ConnectProxyBuyParamsType = "dedicated_residential"
	ConnectProxyBuyParamsTypeResidential          ConnectProxyBuyParamsType = "residential"
	ConnectProxyBuyParamsTypeMobile               ConnectProxyBuyParamsType = "mobile"
)

type ConnectProxyListConnectionsParams struct {
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
	// Filter to a single connection by its session id.
	SessionID param.Opt[string] `query:"sessionId,omitzero" format:"uuid" json:"-"`
	// Filter to connections that started at or after this time (inclusive).
	StartedAfter param.Opt[time.Time] `query:"startedAfter,omitzero" format:"date-time" json:"-"`
	// Filter to connections that started at or before this time (inclusive).
	StartedBefore param.Opt[time.Time] `query:"startedBefore,omitzero" format:"date-time" json:"-"`
	// Filter to connections made by this user.
	UserID param.Opt[string] `query:"userId,omitzero" format:"uuid" json:"-"`
	// Sort direction.
	//
	// Any of "asc", "desc".
	Order ConnectProxyListConnectionsParamsOrder `query:"order,omitzero" json:"-"`
	// Property to order the results by.
	//
	// Any of "startedAt", "endedAt", "bytesIn", "bytesOut", "totalBytes",
	// "durationMs".
	OrderBy ConnectProxyListConnectionsParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// Filter to connections of this transport protocol.
	//
	// Any of "tcp", "udp", "unknown".
	Protocol ConnectProxyListConnectionsParamsProtocol `query:"protocol,omitzero" json:"-"`
	// Filter by connection status.
	//
	// Any of "active", "closed".
	Status ConnectProxyListConnectionsParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectProxyListConnectionsParams]'s query parameters as
// `url.Values`.
func (r ConnectProxyListConnectionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction.
type ConnectProxyListConnectionsParamsOrder string

const (
	ConnectProxyListConnectionsParamsOrderAsc  ConnectProxyListConnectionsParamsOrder = "asc"
	ConnectProxyListConnectionsParamsOrderDesc ConnectProxyListConnectionsParamsOrder = "desc"
)

// Property to order the results by.
type ConnectProxyListConnectionsParamsOrderBy string

const (
	ConnectProxyListConnectionsParamsOrderByStartedAt  ConnectProxyListConnectionsParamsOrderBy = "startedAt"
	ConnectProxyListConnectionsParamsOrderByEndedAt    ConnectProxyListConnectionsParamsOrderBy = "endedAt"
	ConnectProxyListConnectionsParamsOrderByBytesIn    ConnectProxyListConnectionsParamsOrderBy = "bytesIn"
	ConnectProxyListConnectionsParamsOrderByBytesOut   ConnectProxyListConnectionsParamsOrderBy = "bytesOut"
	ConnectProxyListConnectionsParamsOrderByTotalBytes ConnectProxyListConnectionsParamsOrderBy = "totalBytes"
	ConnectProxyListConnectionsParamsOrderByDurationMs ConnectProxyListConnectionsParamsOrderBy = "durationMs"
)

// Filter to connections of this transport protocol.
type ConnectProxyListConnectionsParamsProtocol string

const (
	ConnectProxyListConnectionsParamsProtocolTcp     ConnectProxyListConnectionsParamsProtocol = "tcp"
	ConnectProxyListConnectionsParamsProtocolUdp     ConnectProxyListConnectionsParamsProtocol = "udp"
	ConnectProxyListConnectionsParamsProtocolUnknown ConnectProxyListConnectionsParamsProtocol = "unknown"
)

// Filter by connection status.
type ConnectProxyListConnectionsParamsStatus string

const (
	ConnectProxyListConnectionsParamsStatusActive ConnectProxyListConnectionsParamsStatus = "active"
	ConnectProxyListConnectionsParamsStatusClosed ConnectProxyListConnectionsParamsStatus = "closed"
)
