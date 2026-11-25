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

// ComputeFunctionValidationService contains methods and other services that help
// with interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeFunctionValidationService] method instead.
type ComputeFunctionValidationService struct {
	options []option.RequestOption
}

// NewComputeFunctionValidationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewComputeFunctionValidationService(opts ...option.RequestOption) (r ComputeFunctionValidationService) {
	r = ComputeFunctionValidationService{}
	r.options = opts
	return
}

// Validate Python function source without uploading it
func (r *ComputeFunctionValidationService) New(ctx context.Context, params ComputeFunctionValidationNewParams, opts ...option.RequestOption) (res *ComputeFunctionValidationNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "compute/functions/validations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a function validation owned by this organisation
func (r *ComputeFunctionValidationService) Get(ctx context.Context, validationID string, query ComputeFunctionValidationGetParams, opts ...option.RequestOption) (res *ComputeFunctionValidationGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if validationID == "" {
		err = errors.New("missing required validationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/functions/validations/%s", validationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type FunctionValidation struct {
	ID        string                    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time                 `json:"createdAt" api:"required" format:"date-time"`
	Errors    []FunctionValidationError `json:"errors" api:"required"`
	// Any of "pending", "done", "failed".
	Status      FunctionValidationStatus    `json:"status" api:"required"`
	Warnings    []FunctionValidationWarning `json:"warnings" api:"required"`
	DatasetKey  string                      `json:"datasetKey" api:"nullable"`
	FinishedAt  time.Time                   `json:"finishedAt" format:"date-time"`
	InputSchema map[string]any              `json:"inputSchema"`
	// Any of "value", "dataset", "none".
	OutputKind     FunctionValidationOutputKind `json:"outputKind"`
	OutputSchema   map[string]any               `json:"outputSchema"`
	RuntimeVersion string                       `json:"runtimeVersion"`
	Valid          bool                         `json:"valid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Errors         respjson.Field
		Status         respjson.Field
		Warnings       respjson.Field
		DatasetKey     respjson.Field
		FinishedAt     respjson.Field
		InputSchema    respjson.Field
		OutputKind     respjson.Field
		OutputSchema   respjson.Field
		RuntimeVersion respjson.Field
		Valid          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionValidation) RawJSON() string { return r.JSON.raw }
func (r *FunctionValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionValidationError struct {
	Code    string                        `json:"code" api:"required"`
	Message string                        `json:"message" api:"required"`
	Line    int64                         `json:"line"`
	Repair  FunctionValidationErrorRepair `json:"repair"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Line        respjson.Field
		Repair      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionValidationError) RawJSON() string { return r.JSON.raw }
func (r *FunctionValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionValidationErrorRepair struct {
	Action string `json:"action" api:"required"`
	Hint   string `json:"hint" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Hint        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionValidationErrorRepair) RawJSON() string { return r.JSON.raw }
func (r *FunctionValidationErrorRepair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionValidationStatus string

const (
	FunctionValidationStatusPending FunctionValidationStatus = "pending"
	FunctionValidationStatusDone    FunctionValidationStatus = "done"
	FunctionValidationStatusFailed  FunctionValidationStatus = "failed"
)

type FunctionValidationWarning struct {
	Code    string                          `json:"code" api:"required"`
	Message string                          `json:"message" api:"required"`
	Line    int64                           `json:"line"`
	Repair  FunctionValidationWarningRepair `json:"repair"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Line        respjson.Field
		Repair      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionValidationWarning) RawJSON() string { return r.JSON.raw }
func (r *FunctionValidationWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionValidationWarningRepair struct {
	Action string `json:"action" api:"required"`
	Hint   string `json:"hint" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Hint        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionValidationWarningRepair) RawJSON() string { return r.JSON.raw }
func (r *FunctionValidationWarningRepair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionValidationOutputKind string

const (
	FunctionValidationOutputKindValue   FunctionValidationOutputKind = "value"
	FunctionValidationOutputKindDataset FunctionValidationOutputKind = "dataset"
	FunctionValidationOutputKindNone    FunctionValidationOutputKind = "none"
)

type ComputeFunctionValidationNewResponse struct {
	Data FunctionValidation `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionValidationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionValidationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionValidationGetResponse struct {
	Data FunctionValidation `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFunctionValidationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFunctionValidationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFunctionValidationNewParams struct {
	Source string `json:"source" api:"required"`
	// Wait up to 25 seconds for validation to finish. Excess waiters return
	// immediately.
	WaitS    param.Opt[int64]      `query:"waitS,omitzero" json:"-"`
	Manifest FunctionManifestParam `json:"manifest,omitzero"`
	paramObj
}

func (r ComputeFunctionValidationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFunctionValidationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFunctionValidationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ComputeFunctionValidationNewParams]'s query parameters as
// `url.Values`.
func (r ComputeFunctionValidationNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeFunctionValidationGetParams struct {
	// Wait up to 25 seconds for validation to finish. Excess waiters return
	// immediately.
	WaitS param.Opt[int64] `query:"waitS,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeFunctionValidationGetParams]'s query parameters as
// `url.Values`.
func (r ComputeFunctionValidationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
