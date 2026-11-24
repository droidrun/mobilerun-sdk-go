// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
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
