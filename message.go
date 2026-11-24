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
	"github.com/droidrun/mobilerun-sdk-go/v6/shared"
)

// MessageService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	options       []option.RequestOption
	Conversations MessageConversationService
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.options = opts
	r.Conversations = NewMessageConversationService(opts...)
	return
}

// Returns one message. For an outbound message that is not yet final, the status
// may be refreshed once (at most every 30 s) before it is returned; a failed
// refresh returns the stored message.
func (r *MessageService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MessageGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("numbers/messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists SMS messages newest first, with filters for direction, eSIM, phone number,
// status, and conversation.
func (r *MessageService) List(ctx context.Context, query MessageListParams, opts ...option.RequestOption) (res *MessageListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "numbers/messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MessageGetResponse struct {
	Data MessageGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageGetResponseData struct {
	ID             string                               `json:"id" api:"required" format:"uuid"`
	Body           string                               `json:"body" api:"required"`
	Classification MessageGetResponseDataClassification `json:"classification" api:"required"`
	CreatedAt      time.Time                            `json:"createdAt" api:"required" format:"date-time"`
	DeliveryStatus string                               `json:"deliveryStatus" api:"required"`
	DetectedSender string                               `json:"detectedSender" api:"required"`
	// Any of "inbound", "outbound".
	Direction    string    `json:"direction" api:"required"`
	EsimID       string    `json:"esimId" api:"required" format:"uuid"`
	OccurredAt   time.Time `json:"occurredAt" api:"required" format:"date-time"`
	PeerKey      string    `json:"peerKey" api:"required"`
	PeerNumber   string    `json:"peerNumber" api:"required"`
	ProviderCode string    `json:"providerCode" api:"required"`
	// Message status. Outbound messages move queued → sending → sent, or end at
	// failed. `sent` means the mobile network accepted the message and is the final
	// success status. `delivered` and `sent_unconfirmed` are kept for compatibility;
	// current lines return no carrier delivery report, so neither is emitted.
	//
	// Any of "received", "queued", "claimed", "sending", "sent", "sent_unconfirmed",
	// "delivered", "failed".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Body           respjson.Field
		Classification respjson.Field
		CreatedAt      respjson.Field
		DeliveryStatus respjson.Field
		DetectedSender respjson.Field
		Direction      respjson.Field
		EsimID         respjson.Field
		OccurredAt     respjson.Field
		PeerKey        respjson.Field
		PeerNumber     respjson.Field
		ProviderCode   respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessageGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageGetResponseDataClassification struct {
	Spam bool `json:"spam" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Spam        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetResponseDataClassification) RawJSON() string { return r.JSON.raw }
func (r *MessageGetResponseDataClassification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageListResponse struct {
	Items      []MessageListResponseItem `json:"items" api:"required"`
	Pagination shared.Pagination         `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageListResponseItem struct {
	ID             string                                `json:"id" api:"required" format:"uuid"`
	Body           string                                `json:"body" api:"required"`
	Classification MessageListResponseItemClassification `json:"classification" api:"required"`
	CreatedAt      time.Time                             `json:"createdAt" api:"required" format:"date-time"`
	DeliveryStatus string                                `json:"deliveryStatus" api:"required"`
	DetectedSender string                                `json:"detectedSender" api:"required"`
	// Any of "inbound", "outbound".
	Direction    string    `json:"direction" api:"required"`
	EsimID       string    `json:"esimId" api:"required" format:"uuid"`
	OccurredAt   time.Time `json:"occurredAt" api:"required" format:"date-time"`
	PeerKey      string    `json:"peerKey" api:"required"`
	PeerNumber   string    `json:"peerNumber" api:"required"`
	ProviderCode string    `json:"providerCode" api:"required"`
	// Message status. Outbound messages move queued → sending → sent, or end at
	// failed. `sent` means the mobile network accepted the message and is the final
	// success status. `delivered` and `sent_unconfirmed` are kept for compatibility;
	// current lines return no carrier delivery report, so neither is emitted.
	//
	// Any of "received", "queued", "claimed", "sending", "sent", "sent_unconfirmed",
	// "delivered", "failed".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Body           respjson.Field
		Classification respjson.Field
		CreatedAt      respjson.Field
		DeliveryStatus respjson.Field
		DetectedSender respjson.Field
		Direction      respjson.Field
		EsimID         respjson.Field
		OccurredAt     respjson.Field
		PeerKey        respjson.Field
		PeerNumber     respjson.Field
		ProviderCode   respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *MessageListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageListResponseItemClassification struct {
	Spam bool `json:"spam" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Spam        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageListResponseItemClassification) RawJSON() string { return r.JSON.raw }
func (r *MessageListResponseItemClassification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageListParams struct {
	EsimID     param.Opt[string] `query:"esimId,omitzero" format:"uuid" json:"-"`
	NumberID   param.Opt[string] `query:"numberId,omitzero" format:"uuid" json:"-"`
	Page       param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	PeerKey    param.Opt[string] `query:"peerKey,omitzero" json:"-"`
	PeerNumber param.Opt[string] `query:"peerNumber,omitzero" json:"-"`
	// Any of "all", "inbound", "outbound".
	Direction MessageListParamsDirection `query:"direction,omitzero" json:"-"`
	// Filter by message status. `delivered` and `sent_unconfirmed` are kept for
	// compatibility and are not emitted on current lines.
	//
	// Any of "all", "received", "queued", "claimed", "sending", "sent",
	// "sent_unconfirmed", "delivered", "failed".
	Status MessageListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageListParams]'s query parameters as `url.Values`.
func (r MessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageListParamsDirection string

const (
	MessageListParamsDirectionAll      MessageListParamsDirection = "all"
	MessageListParamsDirectionInbound  MessageListParamsDirection = "inbound"
	MessageListParamsDirectionOutbound MessageListParamsDirection = "outbound"
)

// Filter by message status. `delivered` and `sent_unconfirmed` are kept for
// compatibility and are not emitted on current lines.
type MessageListParamsStatus string

const (
	MessageListParamsStatusAll             MessageListParamsStatus = "all"
	MessageListParamsStatusReceived        MessageListParamsStatus = "received"
	MessageListParamsStatusQueued          MessageListParamsStatus = "queued"
	MessageListParamsStatusClaimed         MessageListParamsStatus = "claimed"
	MessageListParamsStatusSending         MessageListParamsStatus = "sending"
	MessageListParamsStatusSent            MessageListParamsStatus = "sent"
	MessageListParamsStatusSentUnconfirmed MessageListParamsStatus = "sent_unconfirmed"
	MessageListParamsStatusDelivered       MessageListParamsStatus = "delivered"
	MessageListParamsStatusFailed          MessageListParamsStatus = "failed"
)
