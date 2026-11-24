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

// MailboxMessageService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMailboxMessageService] method instead.
type MailboxMessageService struct {
	options []option.RequestOption
}

// NewMailboxMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMailboxMessageService(opts ...option.RequestOption) (r MailboxMessageService) {
	r = MailboxMessageService{}
	r.options = opts
	return
}

// Get a mailbox message
func (r *MailboxMessageService) Get(ctx context.Context, messageID string, params MailboxMessageGetParams, opts ...option.RequestOption) (res *MailboxMessageGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.MailboxID == "" {
		err = errors.New("missing required mailboxId parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("mailboxes/%s/messages/%s", params.MailboxID, url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Lists mailbox messages with cursor pagination and optional time, sender, and OTP
// filters.
func (r *MailboxMessageService) List(ctx context.Context, mailboxID string, query MailboxMessageListParams, opts ...option.RequestOption) (res *MailboxMessageListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if mailboxID == "" {
		err = errors.New("missing required mailboxId parameter")
		return nil, err
	}
	path := fmt.Sprintf("mailboxes/%s/messages", mailboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MailboxMessageGetResponse struct {
	Data MailboxMessageGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxMessageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxMessageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxMessageGetResponseData struct {
	ID              string                                    `json:"id" api:"required"`
	AttachmentCount int64                                     `json:"attachmentCount" api:"required"`
	Attachments     []MailboxMessageGetResponseDataAttachment `json:"attachments" api:"required"`
	FromAddress     string                                    `json:"fromAddress" api:"required"`
	FromName        string                                    `json:"fromName" api:"required"`
	HasOtp          bool                                      `json:"hasOtp" api:"required"`
	MailboxID       string                                    `json:"mailboxId" api:"required" format:"uuid"`
	ReceivedAt      time.Time                                 `json:"receivedAt" api:"required" format:"date-time"`
	Subject         string                                    `json:"subject" api:"required"`
	TextBody        string                                    `json:"textBody" api:"required"`
	// Raw HTML part as received. Present only with includeHtml=true; null when the
	// mail has no HTML part or it exceeds 1 MiB (UTF-8 bytes). Untrusted content:
	// render only in an isolated sandboxed frame, never inline.
	HTMLBody string `json:"htmlBody" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AttachmentCount respjson.Field
		Attachments     respjson.Field
		FromAddress     respjson.Field
		FromName        respjson.Field
		HasOtp          respjson.Field
		MailboxID       respjson.Field
		ReceivedAt      respjson.Field
		Subject         respjson.Field
		TextBody        respjson.Field
		HTMLBody        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxMessageGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *MailboxMessageGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxMessageGetResponseDataAttachment struct {
	ContentType string `json:"contentType" api:"required"`
	Name        string `json:"name" api:"required"`
	Size        int64  `json:"size" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Name        respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxMessageGetResponseDataAttachment) RawJSON() string { return r.JSON.raw }
func (r *MailboxMessageGetResponseDataAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxMessageListResponse struct {
	Items      []MailboxMessageListResponseItem `json:"items" api:"required"`
	NextCursor string                           `json:"nextCursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxMessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxMessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxMessageListResponseItem struct {
	ID          string    `json:"id" api:"required"`
	FromAddress string    `json:"fromAddress" api:"required"`
	HasOtp      bool      `json:"hasOtp" api:"required"`
	MailboxID   string    `json:"mailboxId" api:"required" format:"uuid"`
	Preview     string    `json:"preview" api:"required"`
	ReceivedAt  time.Time `json:"receivedAt" api:"required" format:"date-time"`
	Subject     string    `json:"subject" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FromAddress respjson.Field
		HasOtp      respjson.Field
		MailboxID   respjson.Field
		Preview     respjson.Field
		ReceivedAt  respjson.Field
		Subject     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxMessageListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *MailboxMessageListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxMessageGetParams struct {
	MailboxID string `path:"mailboxId" api:"required" format:"uuid" json:"-"`
	// Return the raw HTML part as htmlBody (untrusted; render isolated).
	//
	// Any of "true", "false".
	IncludeHTML MailboxMessageGetParamsIncludeHTML `query:"includeHtml,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MailboxMessageGetParams]'s query parameters as
// `url.Values`.
func (r MailboxMessageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Return the raw HTML part as htmlBody (untrusted; render isolated).
type MailboxMessageGetParamsIncludeHTML string

const (
	MailboxMessageGetParamsIncludeHTMLTrue  MailboxMessageGetParamsIncludeHTML = "true"
	MailboxMessageGetParamsIncludeHTMLFalse MailboxMessageGetParamsIncludeHTML = "false"
)

type MailboxMessageListParams struct {
	After  param.Opt[time.Time] `query:"after,omitzero" format:"date-time" json:"-"`
	Cursor param.Opt[string]    `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]     `query:"limit,omitzero" json:"-"`
	Sender param.Opt[string]    `query:"sender,omitzero" json:"-"`
	// Any of "true", "false".
	HasOtp MailboxMessageListParamsHasOtp `query:"hasOtp,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MailboxMessageListParams]'s query parameters as
// `url.Values`.
func (r MailboxMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MailboxMessageListParamsHasOtp string

const (
	MailboxMessageListParamsHasOtpTrue  MailboxMessageListParamsHasOtp = "true"
	MailboxMessageListParamsHasOtpFalse MailboxMessageListParamsHasOtp = "false"
)
