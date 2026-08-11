// Package arubacloud contains namespace-scoped resource configuration for the
// ArubaCloud Crossplane provider.
package arubacloud

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures ArubaCloud resources for the namespaced provider.
func Configure(p *ujconfig.Provider) {
	_ = p
}
