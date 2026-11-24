// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/internal/requestconfig"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// WorkflowActionServiceService contains methods and other services that help with
// interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowActionServiceService] method instead.
type WorkflowActionServiceService struct {
	options []option.RequestOption
}

// NewWorkflowActionServiceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowActionServiceService(opts ...option.RequestOption) (r WorkflowActionServiceService) {
	r = WorkflowActionServiceService{}
	r.options = opts
	return
}

// Return the names of the services that actions can be built against. Use these
// values to look up each service's allowed methods.
func (r *WorkflowActionServiceService) List(ctx context.Context, opts ...option.RequestOption) (res *WorkflowActionServiceListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "actions/services"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Return the methods allowed for the given service, each with its parameter
// definitions (name, type, whether required, description, and optional
// default/example). Returns 404 if the service is unknown.
func (r *WorkflowActionServiceService) ListMethods(ctx context.Context, service string, opts ...option.RequestOption) (res *WorkflowActionServiceListMethodsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if service == "" {
		err = errors.New("missing required service parameter")
		return nil, err
	}
	path := fmt.Sprintf("actions/services/%s/methods", url.PathEscape(service))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WorkflowActionServiceListResponse struct {
	Data []string `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionServiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionServiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionServiceListMethodsResponse struct {
	Data []WorkflowActionServiceListMethodsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionServiceListMethodsResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionServiceListMethodsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionServiceListMethodsResponseData struct {
	Method string                                              `json:"method" api:"required"`
	Params []WorkflowActionServiceListMethodsResponseDataParam `json:"params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Params      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionServiceListMethodsResponseData) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionServiceListMethodsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionServiceListMethodsResponseDataParam struct {
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	Required    bool   `json:"required" api:"required"`
	// Any of "string", "number", "boolean", "object", "array", "any".
	Type    string `json:"type" api:"required"`
	Default any    `json:"default"`
	Example any    `json:"example"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Default     respjson.Field
		Example     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionServiceListMethodsResponseDataParam) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionServiceListMethodsResponseDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
