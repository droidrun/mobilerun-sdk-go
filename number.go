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

// NumberService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumberService] method instead.
type NumberService struct {
	options  []option.RequestOption
	Messages NumberMessageService
}

// NewNumberService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNumberService(opts ...option.RequestOption) (r NumberService) {
	r = NumberService{}
	r.options = opts
	r.Messages = NewNumberMessageService(opts...)
	return
}

// Starts a phone-number purchase. Poll the returned phone number for status
// updates.
func (r *NumberService) New(ctx context.Context, body NumberNewParams, opts ...option.RequestOption) (res *NumberNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "numbers/phones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single phone number.
func (r *NumberService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *NumberGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("numbers/phones/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the display label. Omitting `label` leaves it unchanged; null or an
// empty string clears it.
func (r *NumberService) Update(ctx context.Context, id string, body NumberUpdateParams, opts ...option.RequestOption) (res *NumberUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("numbers/phones/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists the caller's phone numbers.
func (r *NumberService) List(ctx context.Context, query NumberListParams, opts ...option.RequestOption) (res *pagination.Page[PublicNumber], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "numbers/phones"
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

// Lists the caller's phone numbers.
func (r *NumberService) ListAutoPaging(ctx context.Context, query NumberListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PublicNumber] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Cancels a pending purchase or schedules cancellation of an active paid phone
// number. Repeating a scheduled cancellation is safe. Returns 409 when
// cancellation is not available.
func (r *NumberService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *NumberDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("numbers/phones/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns included phone capacity for a country. Creating a phone is
// authoritative.
func (r *NumberService) Capacity(ctx context.Context, query NumberCapacityParams, opts ...option.RequestOption) (res *NumberCapacityResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "numbers/phones/capacity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lists available countries and current phone-number availability.
func (r *NumberService) Countries(ctx context.Context, opts ...option.RequestOption) (res *NumberCountriesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "numbers/phones/countries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The device this number is bound to, or null.
type PublicDeviceBinding struct {
	BoundAt    time.Time `json:"boundAt" api:"required" format:"date-time"`
	DeviceID   string    `json:"deviceId" api:"required"`
	DeviceType string    `json:"deviceType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoundAt     respjson.Field
		DeviceID    respjson.Field
		DeviceType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicDeviceBinding) RawJSON() string { return r.JSON.raw }
func (r *PublicDeviceBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicNumber struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Actions currently available for this phone number.
	Actions           PublicNumberActions `json:"actions" api:"required"`
	CancelAtPeriodEnd bool                `json:"cancelAtPeriodEnd" api:"required"`
	// Deprecated: use `actions.cancel`
	//
	// Deprecated: deprecated
	Cancellable bool `json:"cancellable" api:"required"`
	// Deprecated: use `actions.send`
	//
	// Deprecated: deprecated
	CanSend bool `json:"canSend" api:"required"`
	// Deprecated: use `lineCapabilities`
	//
	// Any of "sms", "voice".
	//
	// Deprecated: deprecated
	Capabilities     []string  `json:"capabilities" api:"required"`
	CheckoutURL      string    `json:"checkoutUrl" api:"required"`
	CountryCode      string    `json:"countryCode" api:"required"`
	CreatedAt        time.Time `json:"createdAt" api:"required" format:"date-time"`
	CurrentPeriodEnd time.Time `json:"currentPeriodEnd" api:"required" format:"date-time"`
	// The device this number is bound to, or null.
	DeviceBinding PublicDeviceBinding `json:"deviceBinding" api:"required"`
	// Why setup failed. Null unless state is failed, and null when the failure has no
	// classified reason. out_of_stock: no number is currently available for the
	// requested country.
	//
	// Any of "out_of_stock", "unknown".
	FailureReason PublicNumberFailureReason `json:"failureReason" api:"required"`
	// ICCID of the SIM behind this line when a trustworthy value is known, else null.
	Iccid string `json:"iccid" api:"required"`
	Label string `json:"label" api:"required"`
	// What this line supports in principle, per channel. Use `actions` for what you
	// can do right now.
	LineCapabilities PublicNumberLineCapabilities `json:"lineCapabilities" api:"required"`
	PhoneNumber      string                       `json:"phoneNumber" api:"required"`
	Purpose          string                       `json:"purpose" api:"required"`
	// SIM of a phone connected through the Mobilerun Portal app; null for other phones
	// or before the first sync.
	Sim PublicNumberSim `json:"sim" api:"required"`
	// Any of "awaiting_payment", "provisioning", "active", "cancel_scheduled",
	// "expired", "failed".
	State     PublicNumberState `json:"state" api:"required"`
	UpdatedAt time.Time         `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Actions           respjson.Field
		CancelAtPeriodEnd respjson.Field
		Cancellable       respjson.Field
		CanSend           respjson.Field
		Capabilities      respjson.Field
		CheckoutURL       respjson.Field
		CountryCode       respjson.Field
		CreatedAt         respjson.Field
		CurrentPeriodEnd  respjson.Field
		DeviceBinding     respjson.Field
		FailureReason     respjson.Field
		Iccid             respjson.Field
		Label             respjson.Field
		LineCapabilities  respjson.Field
		PhoneNumber       respjson.Field
		Purpose           respjson.Field
		Sim               respjson.Field
		State             respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNumber) RawJSON() string { return r.JSON.raw }
func (r *PublicNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Actions currently available for this phone number.
type PublicNumberActions struct {
	// True when this number can be bound to a cloud device (active, has a phone
	// number, receives SMS, not a BYO device line). When false,
	// `bindDeviceBlockedReason` says why.
	BindDevice bool `json:"bindDevice" api:"required"`
	// Why `bindDevice` is false; null exactly when it is true.
	//
	// Any of "number_not_active", "no_phone_number", "sms_inbound_unsupported",
	// "byo_device_line".
	BindDeviceBlockedReason string `json:"bindDeviceBlockedReason" api:"required"`
	// True when DELETE /numbers/phones/{id} would cancel or abort this number now.
	// False once a cancellation is already scheduled; the DELETE stays idempotent.
	Cancel bool `json:"cancel" api:"required"`
	// True when this number passes the send gate of POST
	// /numbers/phones/{id}/messages: self-service sending is switched on and the
	// number is active and able to send SMS. Daily and burst limits and recipient/body
	// checks still apply; the send endpoint stays the final judge.
	Send bool `json:"send" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BindDevice              respjson.Field
		BindDeviceBlockedReason respjson.Field
		Cancel                  respjson.Field
		Send                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNumberActions) RawJSON() string { return r.JSON.raw }
func (r *PublicNumberActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why setup failed. Null unless state is failed, and null when the failure has no
// classified reason. out_of_stock: no number is currently available for the
// requested country.
type PublicNumberFailureReason string

const (
	PublicNumberFailureReasonOutOfStock PublicNumberFailureReason = "out_of_stock"
	PublicNumberFailureReasonUnknown    PublicNumberFailureReason = "unknown"
)

// What this line supports in principle, per channel. Use `actions` for what you
// can do right now.
type PublicNumberLineCapabilities struct {
	SMS PublicNumberLineCapabilitiesSMS `json:"sms" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SMS         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNumberLineCapabilities) RawJSON() string { return r.JSON.raw }
func (r *PublicNumberLineCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicNumberLineCapabilitiesSMS struct {
	// This line can receive on this channel.
	Inbound bool `json:"inbound" api:"required"`
	// This line can send on this channel.
	Outbound bool `json:"outbound" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inbound     respjson.Field
		Outbound    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNumberLineCapabilitiesSMS) RawJSON() string { return r.JSON.raw }
func (r *PublicNumberLineCapabilitiesSMS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SIM of a phone connected through the Mobilerun Portal app; null for other phones
// or before the first sync.
type PublicNumberSim struct {
	CarrierName string `json:"carrierName" api:"required"`
	DisplayName string `json:"displayName" api:"required"`
	SimSlot     int64  `json:"simSlot" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CarrierName respjson.Field
		DisplayName respjson.Field
		SimSlot     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNumberSim) RawJSON() string { return r.JSON.raw }
func (r *PublicNumberSim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicNumberState string

const (
	PublicNumberStateAwaitingPayment PublicNumberState = "awaiting_payment"
	PublicNumberStateProvisioning    PublicNumberState = "provisioning"
	PublicNumberStateActive          PublicNumberState = "active"
	PublicNumberStateCancelScheduled PublicNumberState = "cancel_scheduled"
	PublicNumberStateExpired         PublicNumberState = "expired"
	PublicNumberStateFailed          PublicNumberState = "failed"
)

type NumberNewResponse struct {
	Data NumberNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberNewResponseData struct {
	CheckoutURL string `json:"checkoutUrl" api:"required"`
	NumberID    string `json:"numberId" api:"required" format:"uuid"`
	// Any of "awaiting_payment", "provisioning".
	State string `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckoutURL respjson.Field
		NumberID    respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *NumberNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberGetResponse struct {
	Data PublicNumber `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberUpdateResponse struct {
	Data PublicNumber `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberDeleteResponse struct {
	Data PublicNumber `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberCapacityResponse struct {
	Data NumberCapacityResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberCapacityResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberCapacityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberCapacityResponseData struct {
	Included int64 `json:"included" api:"required"`
	// Deprecated — always equal to `remaining`. Migrate to `remaining`; this field
	// will be removed in a future revision.
	//
	// Deprecated: deprecated
	IncludedRemaining int64 `json:"includedRemaining" api:"required"`
	Remaining         int64 `json:"remaining" api:"required"`
	// Any of "available", "exhausted", "not_included".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Included          respjson.Field
		IncludedRemaining respjson.Field
		Remaining         respjson.Field
		Status            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberCapacityResponseData) RawJSON() string { return r.JSON.raw }
func (r *NumberCapacityResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberCountriesResponse struct {
	Data NumberCountriesResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberCountriesResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberCountriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberCountriesResponseData struct {
	Items []NumberCountriesResponseDataItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberCountriesResponseData) RawJSON() string { return r.JSON.raw }
func (r *NumberCountriesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberCountriesResponseDataItem struct {
	Country string `json:"country" api:"required"`
	InStock bool   `json:"inStock" api:"required"`
	Name    string `json:"name" api:"required"`
	PlanID  string `json:"planId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		InStock     respjson.Field
		Name        respjson.Field
		PlanID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberCountriesResponseDataItem) RawJSON() string { return r.JSON.raw }
func (r *NumberCountriesResponseDataItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberNewParams struct {
	// User-defined display label — NFC-normalized, up to 100 GRAPHEMES (not UTF-16
	// code units; an emoji/flag may span several). Display-only, never used for
	// routing.
	Label param.Opt[string] `json:"label,omitzero"`
	// Optional ISO 3166-1 alpha-2 country code from GET /numbers/phones/countries.
	Country param.Opt[string] `json:"country,omitzero"`
	// Use included capacity when available, require included capacity without paid
	// fallback (included_only), or start a paid checkout (rent).
	//
	// Any of "included", "included_only", "rent".
	BillingPreference NumberNewParamsBillingPreference `json:"billingPreference,omitzero"`
	paramObj
}

func (r NumberNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NumberNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use included capacity when available, require included capacity without paid
// fallback (included_only), or start a paid checkout (rent).
type NumberNewParamsBillingPreference string

const (
	NumberNewParamsBillingPreferenceIncluded     NumberNewParamsBillingPreference = "included"
	NumberNewParamsBillingPreferenceIncludedOnly NumberNewParamsBillingPreference = "included_only"
	NumberNewParamsBillingPreferenceRent         NumberNewParamsBillingPreference = "rent"
)

type NumberUpdateParams struct {
	// User-defined display label — NFC-normalized, up to 100 GRAPHEMES (not UTF-16
	// code units; an emoji/flag may span several). Display-only, never used for
	// routing.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r NumberUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NumberUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberListParams]'s query parameters as `url.Values`.
func (r NumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NumberCapacityParams struct {
	// ISO 3166-1 alpha-2 country code from GET /numbers/phones/countries.
	Country string `query:"country" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [NumberCapacityParams]'s query parameters as `url.Values`.
func (r NumberCapacityParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
