// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
)

// CredentialOnepasswordVaultService contains methods and other services that help
// with interacting with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialOnepasswordVaultService] method instead.
type CredentialOnepasswordVaultService struct {
	options []option.RequestOption
	Items   CredentialOnepasswordVaultItemService
}

// NewCredentialOnepasswordVaultService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCredentialOnepasswordVaultService(opts ...option.RequestOption) (r CredentialOnepasswordVaultService) {
	r = CredentialOnepasswordVaultService{}
	r.options = opts
	r.Items = NewCredentialOnepasswordVaultItemService(opts...)
	return
}
