// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mobilerun

import (
	"github.com/droidrun/mobilerun-sdk-go/v6/option"
)

// ConnectService contains methods and other services that help with interacting
// with the mobilerun API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectService] method instead.
type ConnectService struct {
	options []option.RequestOption
	// Mobilerun Connect country coverage information
	Countries ConnectCountryService
	// Manage your Mobilerun Connect Proxies
	Proxies ConnectProxyService
	// Manage your Mobilerun Connect Socks Users
	Users ConnectUserService
}

// NewConnectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConnectService(opts ...option.RequestOption) (r ConnectService) {
	r = ConnectService{}
	r.options = opts
	r.Countries = NewConnectCountryService(opts...)
	r.Proxies = NewConnectProxyService(opts...)
	r.Users = NewConnectUserService(opts...)
	return
}
