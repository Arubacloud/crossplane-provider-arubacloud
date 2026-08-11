// Package arubacloud contains namespace-scoped resource configuration for the
// ArubaCloud Crossplane provider. Mirrors cluster/arubacloud/configure.go.
package arubacloud

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const uriExtractor = `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("uri",true)`

// Terraform resource type names used as TerraformName in reference configs.
const (
	tfProject       = "arubacloud_project"
	tfVPC           = "arubacloud_vpc"
	tfSubnet        = "arubacloud_subnet"
	tfKeyPair       = "arubacloud_keypair"
	tfElasticIP     = "arubacloud_elasticip"
	tfBlockStorage  = "arubacloud_blockstorage"
	tfSecurityGroup = "arubacloud_securitygroup"
)

// Configure configures ArubaCloud resources for the namespaced provider.
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
	p.AddResourceConfigurator(tfProject, func(r *ujconfig.Resource) {})
}

func configureVPC(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfVPC, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
	})
}

func configureSubnet(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfSubnet, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["vpc_id"] = ujconfig.Reference{TerraformName: tfVPC}
	})
}

func configureKeyPair(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfKeyPair, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
	})
}

func configureElasticIP(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfElasticIP, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
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
	p.AddResourceConfigurator(tfBlockStorage, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
	})
}

func configureCloudServer(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_cloudserver", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["network.vpc_uri_ref"] = ujconfig.Reference{
			TerraformName: tfVPC,
			Extractor:     uriExtractor,
		}
		r.References["network.subnet_uri_refs"] = ujconfig.Reference{
			TerraformName: tfSubnet,
			Extractor:     uriExtractor,
		}
		r.References["network.securitygroup_uri_refs"] = ujconfig.Reference{
			TerraformName: tfSecurityGroup,
			Extractor:     uriExtractor,
		}
		r.References["network.elastic_ip_uri_ref"] = ujconfig.Reference{
			TerraformName: tfElasticIP,
			Extractor:     uriExtractor,
		}
		r.References["settings.key_pair_uri_ref"] = ujconfig.Reference{
			TerraformName: tfKeyPair,
			Extractor:     uriExtractor,
		}
		r.References["storage.boot_volume_uri_ref"] = ujconfig.Reference{
			TerraformName: tfBlockStorage,
			Extractor:     uriExtractor,
		}
	})
}
