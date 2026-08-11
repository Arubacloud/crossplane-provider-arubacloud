// Package arubacloud contains cluster-scoped resource configuration for the
// ArubaCloud Crossplane provider. Add per-resource configurators here as
// Phase 2+ resources are implemented.
package arubacloud

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures ArubaCloud resources for the cluster-scoped provider.
func Configure(p *ujconfig.Provider) {
	// Phase 1: external names are configured centrally in config/external_name.go.
	// Per-resource reference configuration will be added in Phase 2.
	_ = p
}
