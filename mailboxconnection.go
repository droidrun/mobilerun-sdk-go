// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/param"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// MailboxConnectionService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMailboxConnectionService] method instead.
type MailboxConnectionService struct {
	options []option.RequestOption
}

// NewMailboxConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMailboxConnectionService(opts ...option.RequestOption) (r MailboxConnectionService) {
	r = MailboxConnectionService{}
	r.options = opts
	return
}

// Links the caller’s Gmail account. Replaying the same Idempotency-Key returns the
// same mailbox. When Gmail is already active, redirectUrl is null and the mailbox
// is usable immediately.
func (r *MailboxConnectionService) New(ctx context.Context, body MailboxConnectionNewParams, opts ...option.RequestOption) (res *MailboxConnectionNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "mailboxes/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Reports pending, active (with the Gmail address), expired, or removed.
// connectionId is the mailbox id.
func (r *MailboxConnectionService) Get(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *MailboxConnectionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("mailboxes/connections/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MailboxConnectionNewResponse struct {
	Address      string `json:"address" api:"required"`
	ConnectionID string `json:"connectionId" api:"required" format:"uuid"`
	MailboxID    string `json:"mailboxId" api:"required" format:"uuid"`
	RedirectURL  string `json:"redirectUrl" api:"required"`
	// Any of "pending", "active", "expired", "removed".
	Status MailboxConnectionNewResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address      respjson.Field
		ConnectionID respjson.Field
		MailboxID    respjson.Field
		RedirectURL  respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxConnectionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxConnectionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxConnectionNewResponseStatus string

const (
	MailboxConnectionNewResponseStatusPending MailboxConnectionNewResponseStatus = "pending"
	MailboxConnectionNewResponseStatusActive  MailboxConnectionNewResponseStatus = "active"
	MailboxConnectionNewResponseStatusExpired MailboxConnectionNewResponseStatus = "expired"
	MailboxConnectionNewResponseStatusRemoved MailboxConnectionNewResponseStatus = "removed"
)

type MailboxConnectionGetResponse struct {
	Address      string `json:"address" api:"required"`
	ConnectionID string `json:"connectionId" api:"required" format:"uuid"`
	MailboxID    string `json:"mailboxId" api:"required" format:"uuid"`
	RedirectURL  string `json:"redirectUrl" api:"required"`
	// Any of "pending", "active", "expired", "removed".
	Status MailboxConnectionGetResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address      respjson.Field
		ConnectionID respjson.Field
		MailboxID    respjson.Field
		RedirectURL  respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxConnectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxConnectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxConnectionGetResponseStatus string

const (
	MailboxConnectionGetResponseStatusPending MailboxConnectionGetResponseStatus = "pending"
	MailboxConnectionGetResponseStatusActive  MailboxConnectionGetResponseStatus = "active"
	MailboxConnectionGetResponseStatusExpired MailboxConnectionGetResponseStatus = "expired"
	MailboxConnectionGetResponseStatusRemoved MailboxConnectionGetResponseStatus = "removed"
)

type MailboxConnectionNewParams struct {
	// Any of "gmail".
	Provider MailboxConnectionNewParamsProvider `json:"provider,omitzero" api:"required"`
	// sha256 hex of the browser-held connect binding secret.
	BindingHash param.Opt[string] `json:"bindingHash,omitzero"`
	Label       param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r MailboxConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MailboxConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailboxConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxConnectionNewParamsProvider string

const (
	MailboxConnectionNewParamsProviderGmail MailboxConnectionNewParamsProvider = "gmail"
)
