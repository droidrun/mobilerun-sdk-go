// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"encoding/json"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6/internal/apijson"
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
	"github.com/droidrun/mobilerun-sdk-go/v6/packages/respjson"
)

// ComputeService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeService] method instead.
type ComputeService struct {
	options   []option.RequestOption
	Functions ComputeFunctionService
	Runs      ComputeRunService
	Sandboxes ComputeSandboxService
}

// NewComputeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewComputeService(opts ...option.RequestOption) (r ComputeService) {
	r = ComputeService{}
	r.options = opts
	r.Functions = NewComputeFunctionService(opts...)
	r.Runs = NewComputeRunService(opts...)
	r.Sandboxes = NewComputeSandboxService(opts...)
	return
}

type FunctionRun struct {
	ID         string    `json:"id" api:"required" format:"uuid"`
	DurationMs int64     `json:"durationMs" api:"required"`
	EndedAt    time.Time `json:"endedAt" api:"required" format:"date-time"`
	// Any of "function_error", "input_invalid", "timeout", "cancelled",
	// "sandbox_lost", "sandbox_start_failed", "function_not_found", "internal_error",
	// "output_invalid", "result_too_large", "state_too_large", "dataset_too_large",
	// "executor_unavailable".
	ErrorCode  FunctionRunErrorCode `json:"errorCode" api:"required"`
	FunctionID string               `json:"functionId" api:"required" format:"uuid"`
	Input      map[string]any       `json:"input" api:"required"`
	// Any of "none", "pending", "available", "gone".
	Logs FunctionRunLogs `json:"logs" api:"required"`
	// Any of "succeeded", "partial", "failed".
	Outcome   FunctionRunOutcome     `json:"outcome" api:"required"`
	Result    FunctionRunResultUnion `json:"result" api:"required"`
	StartedAt time.Time              `json:"startedAt" api:"required" format:"date-time"`
	// Any of "running", "done", "failed".
	Status      FunctionRunStatus      `json:"status" api:"required"`
	Stderr      string                 `json:"stderr" api:"required"`
	Stdout      string                 `json:"stdout" api:"required"`
	Truncated   bool                   `json:"truncated" api:"required"`
	UsedSeconds int64                  `json:"usedSeconds" api:"required"`
	Version     int64                  `json:"version" api:"required"`
	APIKey      FunctionRunAPIKeyUnion `json:"apiKey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DurationMs  respjson.Field
		EndedAt     respjson.Field
		ErrorCode   respjson.Field
		FunctionID  respjson.Field
		Input       respjson.Field
		Logs        respjson.Field
		Outcome     respjson.Field
		Result      respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Stderr      respjson.Field
		Stdout      respjson.Field
		Truncated   respjson.Field
		UsedSeconds respjson.Field
		Version     respjson.Field
		APIKey      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionRun) RawJSON() string { return r.JSON.raw }
func (r *FunctionRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionRunErrorCode string

const (
	FunctionRunErrorCodeFunctionError       FunctionRunErrorCode = "function_error"
	FunctionRunErrorCodeInputInvalid        FunctionRunErrorCode = "input_invalid"
	FunctionRunErrorCodeTimeout             FunctionRunErrorCode = "timeout"
	FunctionRunErrorCodeCancelled           FunctionRunErrorCode = "cancelled"
	FunctionRunErrorCodeSandboxLost         FunctionRunErrorCode = "sandbox_lost"
	FunctionRunErrorCodeSandboxStartFailed  FunctionRunErrorCode = "sandbox_start_failed"
	FunctionRunErrorCodeFunctionNotFound    FunctionRunErrorCode = "function_not_found"
	FunctionRunErrorCodeInternalError       FunctionRunErrorCode = "internal_error"
	FunctionRunErrorCodeOutputInvalid       FunctionRunErrorCode = "output_invalid"
	FunctionRunErrorCodeResultTooLarge      FunctionRunErrorCode = "result_too_large"
	FunctionRunErrorCodeStateTooLarge       FunctionRunErrorCode = "state_too_large"
	FunctionRunErrorCodeDatasetTooLarge     FunctionRunErrorCode = "dataset_too_large"
	FunctionRunErrorCodeExecutorUnavailable FunctionRunErrorCode = "executor_unavailable"
)

type FunctionRunLogs string

const (
	FunctionRunLogsNone      FunctionRunLogs = "none"
	FunctionRunLogsPending   FunctionRunLogs = "pending"
	FunctionRunLogsAvailable FunctionRunLogs = "available"
	FunctionRunLogsGone      FunctionRunLogs = "gone"
)

type FunctionRunOutcome string

const (
	FunctionRunOutcomeSucceeded FunctionRunOutcome = "succeeded"
	FunctionRunOutcomePartial   FunctionRunOutcome = "partial"
	FunctionRunOutcomeFailed    FunctionRunOutcome = "failed"
)

// FunctionRunResultUnion contains all possible properties and values from
// [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFunctionRunResultMapItem OfAnyArray OfString OfFloat OfBool]
type FunctionRunResultUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfFunctionRunResultMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfFunctionRunResultMapItem respjson.Field
		OfAnyArray                 respjson.Field
		OfString                   respjson.Field
		OfFloat                    respjson.Field
		OfBool                     respjson.Field
		raw                        string
	} `json:"-"`
}

func (u FunctionRunResultUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionRunResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionRunResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionRunResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionRunResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunctionRunResultUnion) RawJSON() string { return u.JSON.raw }

func (r *FunctionRunResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionRunStatus string

const (
	FunctionRunStatusRunning FunctionRunStatus = "running"
	FunctionRunStatusDone    FunctionRunStatus = "done"
	FunctionRunStatusFailed  FunctionRunStatus = "failed"
)

// FunctionRunAPIKeyUnion contains all possible properties and values from
// [FunctionRunAPIKeyObject], [FunctionRunAPIKeyObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunctionRunAPIKeyUnion struct {
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant [FunctionRunAPIKeyObject].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [FunctionRunAPIKeyObject2].
	Full bool `json:"full"`
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u FunctionRunAPIKeyUnion) AsFunctionRunAPIKeyObject() (v FunctionRunAPIKeyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionRunAPIKeyUnion) AsFunctionRunAPIKeyObject2() (v FunctionRunAPIKeyObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunctionRunAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *FunctionRunAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionRunAPIKeyObject struct {
	ExpiresAt   time.Time           `json:"expiresAt" api:"required" format:"date-time"`
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionRunAPIKeyObject) RawJSON() string { return r.JSON.raw }
func (r *FunctionRunAPIKeyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionRunAPIKeyObject2 struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Any of true.
	Full bool `json:"full" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Full        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionRunAPIKeyObject2) RawJSON() string { return r.JSON.raw }
func (r *FunctionRunAPIKeyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionRunSummary struct {
	ID         string    `json:"id" api:"required" format:"uuid"`
	DurationMs int64     `json:"durationMs" api:"required"`
	EndedAt    time.Time `json:"endedAt" api:"required" format:"date-time"`
	// Any of "function_error", "input_invalid", "timeout", "cancelled",
	// "sandbox_lost", "sandbox_start_failed", "function_not_found", "internal_error",
	// "output_invalid", "result_too_large", "state_too_large", "dataset_too_large",
	// "executor_unavailable".
	ErrorCode  FunctionRunSummaryErrorCode `json:"errorCode" api:"required"`
	FunctionID string                      `json:"functionId" api:"required" format:"uuid"`
	// Any of "none", "pending", "available", "gone".
	Logs FunctionRunSummaryLogs `json:"logs" api:"required"`
	// Any of "succeeded", "partial", "failed".
	Outcome   FunctionRunSummaryOutcome `json:"outcome" api:"required"`
	StartedAt time.Time                 `json:"startedAt" api:"required" format:"date-time"`
	// Any of "running", "done", "failed".
	Status      FunctionRunSummaryStatus      `json:"status" api:"required"`
	UsedSeconds int64                         `json:"usedSeconds" api:"required"`
	Version     int64                         `json:"version" api:"required"`
	APIKey      FunctionRunSummaryAPIKeyUnion `json:"apiKey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DurationMs  respjson.Field
		EndedAt     respjson.Field
		ErrorCode   respjson.Field
		FunctionID  respjson.Field
		Logs        respjson.Field
		Outcome     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		UsedSeconds respjson.Field
		Version     respjson.Field
		APIKey      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionRunSummary) RawJSON() string { return r.JSON.raw }
func (r *FunctionRunSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionRunSummaryErrorCode string

const (
	FunctionRunSummaryErrorCodeFunctionError       FunctionRunSummaryErrorCode = "function_error"
	FunctionRunSummaryErrorCodeInputInvalid        FunctionRunSummaryErrorCode = "input_invalid"
	FunctionRunSummaryErrorCodeTimeout             FunctionRunSummaryErrorCode = "timeout"
	FunctionRunSummaryErrorCodeCancelled           FunctionRunSummaryErrorCode = "cancelled"
	FunctionRunSummaryErrorCodeSandboxLost         FunctionRunSummaryErrorCode = "sandbox_lost"
	FunctionRunSummaryErrorCodeSandboxStartFailed  FunctionRunSummaryErrorCode = "sandbox_start_failed"
	FunctionRunSummaryErrorCodeFunctionNotFound    FunctionRunSummaryErrorCode = "function_not_found"
	FunctionRunSummaryErrorCodeInternalError       FunctionRunSummaryErrorCode = "internal_error"
	FunctionRunSummaryErrorCodeOutputInvalid       FunctionRunSummaryErrorCode = "output_invalid"
	FunctionRunSummaryErrorCodeResultTooLarge      FunctionRunSummaryErrorCode = "result_too_large"
	FunctionRunSummaryErrorCodeStateTooLarge       FunctionRunSummaryErrorCode = "state_too_large"
	FunctionRunSummaryErrorCodeDatasetTooLarge     FunctionRunSummaryErrorCode = "dataset_too_large"
	FunctionRunSummaryErrorCodeExecutorUnavailable FunctionRunSummaryErrorCode = "executor_unavailable"
)

type FunctionRunSummaryLogs string

const (
	FunctionRunSummaryLogsNone      FunctionRunSummaryLogs = "none"
	FunctionRunSummaryLogsPending   FunctionRunSummaryLogs = "pending"
	FunctionRunSummaryLogsAvailable FunctionRunSummaryLogs = "available"
	FunctionRunSummaryLogsGone      FunctionRunSummaryLogs = "gone"
)

type FunctionRunSummaryOutcome string

const (
	FunctionRunSummaryOutcomeSucceeded FunctionRunSummaryOutcome = "succeeded"
	FunctionRunSummaryOutcomePartial   FunctionRunSummaryOutcome = "partial"
	FunctionRunSummaryOutcomeFailed    FunctionRunSummaryOutcome = "failed"
)

type FunctionRunSummaryStatus string

const (
	FunctionRunSummaryStatusRunning FunctionRunSummaryStatus = "running"
	FunctionRunSummaryStatusDone    FunctionRunSummaryStatus = "done"
	FunctionRunSummaryStatusFailed  FunctionRunSummaryStatus = "failed"
)

// FunctionRunSummaryAPIKeyUnion contains all possible properties and values from
// [FunctionRunSummaryAPIKeyObject], [FunctionRunSummaryAPIKeyObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunctionRunSummaryAPIKeyUnion struct {
	ExpiresAt time.Time `json:"expiresAt"`
	// This field is from variant [FunctionRunSummaryAPIKeyObject].
	Permissions map[string][]string `json:"permissions"`
	// This field is from variant [FunctionRunSummaryAPIKeyObject2].
	Full bool `json:"full"`
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		Full        respjson.Field
		raw         string
	} `json:"-"`
}

func (u FunctionRunSummaryAPIKeyUnion) AsFunctionRunSummaryAPIKeyObject() (v FunctionRunSummaryAPIKeyObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionRunSummaryAPIKeyUnion) AsFunctionRunSummaryAPIKeyObject2() (v FunctionRunSummaryAPIKeyObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunctionRunSummaryAPIKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *FunctionRunSummaryAPIKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionRunSummaryAPIKeyObject struct {
	ExpiresAt   time.Time           `json:"expiresAt" api:"required" format:"date-time"`
	Permissions map[string][]string `json:"permissions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionRunSummaryAPIKeyObject) RawJSON() string { return r.JSON.raw }
func (r *FunctionRunSummaryAPIKeyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionRunSummaryAPIKeyObject2 struct {
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Any of true.
	Full bool `json:"full" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Full        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionRunSummaryAPIKeyObject2) RawJSON() string { return r.JSON.raw }
func (r *FunctionRunSummaryAPIKeyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
