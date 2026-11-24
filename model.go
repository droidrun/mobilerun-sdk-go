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
	"github.com/droidrun/mobilerun-sdk-go/v6/shared/constant"
)

// ModelService contains methods and other services that help with interacting with
// the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.options = opts
	return
}

// Returns the available LLM models with their identifiers, providers, and optional
// display metadata. Requires only your normal API authentication. Token prices are
// not included.
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *ModelListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ModelListResponse struct {
	Data   []ModelListResponseData `json:"data" api:"required"`
	Object constant.List           `json:"object" default:"list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelListResponseData struct {
	ID      string         `json:"id" api:"required"`
	Created int64          `json:"created" api:"required"`
	Object  constant.Model `json:"object" default:"model"`
	OwnedBy string         `json:"owned_by" api:"required"`
	// Any of "recommended", "fast", "more".
	Group string `json:"group"`
	Label string `json:"label"`
	Order int64  `json:"order"`
	// Model pickers in which this model is available. When omitted, the model is
	// available in every picker.
	//
	// Any of "assistant", "tasks".
	Surfaces []string `json:"surfaces"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Object      respjson.Field
		OwnedBy     respjson.Field
		Group       respjson.Field
		Label       respjson.Field
		Order       respjson.Field
		Surfaces    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
