// Package arubacloud contains cluster-scoped resource configuration for the
// ArubaCloud Crossplane provider.
package arubacloud

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// uriExtractor is the Upjet extractor expression that reads the `uri` field
// from a referenced resource's status.atProvider. All ArubaCloud cross-resource
// references use full URIs rather than bare IDs.
const uriExtractor = `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("uri",true)`

// Configure configures ArubaCloud resources for the cluster-scoped provider.
func Configure(p *ujconfig.Provider) {
	configureProject(p)
	configureVPC(p)
	configureSubnet(p)
	configureKeyPair(p)
	configureElasticIP(p)
	configureBlockStorage(p)
	configureCloudServer(p)
}

func configureProject(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_project", func(r *ujconfig.Resource) {
		// Project is the root resource — no incoming references.
	})
}

func configureVPC(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_vpc", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "arubacloud_project",
		}
	})
}

func configureSubnet(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_subnet", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "arubacloud_project",
		}
		r.References["vpc_id"] = ujconfig.Reference{
			TerraformName: "arubacloud_vpc",
		}
	})
}

func configureKeyPair(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_keypair", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "arubacloud_project",
		}
	})
}

func configureElasticIP(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_elasticip", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "arubacloud_project",
		}

		// Expose the assigned public IP as a connection detail.
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if addr, ok := attr["address"].(string); ok && addr != "" {
				conn["address"] = []byte(addr)
			}
			return conn, nil
		}
	})
}

func configureBlockStorage(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_blockstorage", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "arubacloud_project",
		}
	})
}

func configureCloudServer(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_cloudserver", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "arubacloud_project",
		}
		// URI-based references: extract status.atProvider.uri from the referenced resource.
		r.References["network.vpc_uri_ref"] = ujconfig.Reference{
			TerraformName: "arubacloud_vpc",
			Extractor:     uriExtractor,
		}
		r.References["network.subnet_uri_refs"] = ujconfig.Reference{
			TerraformName: "arubacloud_subnet",
			Extractor:     uriExtractor,
		}
		r.References["network.securitygroup_uri_refs"] = ujconfig.Reference{
			TerraformName: "arubacloud_securitygroup",
			Extractor:     uriExtractor,
		}
		r.References["network.elastic_ip_uri_ref"] = ujconfig.Reference{
			TerraformName: "arubacloud_elasticip",
			Extractor:     uriExtractor,
		}
		r.References["settings.key_pair_uri_ref"] = ujconfig.Reference{
			TerraformName: "arubacloud_keypair",
			Extractor:     uriExtractor,
		}
		r.References["storage.boot_volume_uri_ref"] = ujconfig.Reference{
			TerraformName: "arubacloud_blockstorage",
			Extractor:     uriExtractor,
		}
	})
}
