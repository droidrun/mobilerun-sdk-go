// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"net/http"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// WorkflowTimezoneService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowTimezoneService] method instead.
type WorkflowTimezoneService struct {
	options []option.RequestOption
}

// NewWorkflowTimezoneService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowTimezoneService(opts ...option.RequestOption) (r WorkflowTimezoneService) {
	r = WorkflowTimezoneService{}
	r.options = opts
	return
}

// List supported IANA timezones
func (r *WorkflowTimezoneService) List(ctx context.Context, opts ...option.RequestOption) (res *WorkflowTimezoneListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "timezones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WorkflowTimezoneListResponse struct {
	Data []string `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTimezoneListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTimezoneListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
