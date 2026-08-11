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

// Terraform resource type names used as TerraformName in reference configs.
const (
	tfProject       = "arubacloud_project"
	tfVPC           = "arubacloud_vpc"
	tfSubnet        = "arubacloud_subnet"
	tfKeyPair       = "arubacloud_keypair"
	tfElasticIP     = "arubacloud_elasticip"
	tfBlockStorage  = "arubacloud_blockstorage"
	tfSecurityGroup = "arubacloud_securitygroup"
	tfBackup        = "arubacloud_backup"
	tfVPCPeering    = "arubacloud_vpcpeering"
	tfVPNTunnel     = "arubacloud_vpntunnel"
	tfDBaaS         = "arubacloud_dbaas"
	tfDatabase      = "arubacloud_database"
	tfDBaaSUser     = "arubacloud_dbaasuser"
)

// nameExtractor reads spec.forProvider.name from a referenced resource.
// Used for KaaS security_group_name which references a SecurityGroup by display name.
const nameExtractor = `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("name",false)`

// Configure configures ArubaCloud resources for the cluster-scoped provider.
func Configure(p *ujconfig.Provider) {
	configureKaaS(p)
	configureContainerRegistry(p)
	configureDBaaS(p)
	configureDatabase(p)
	configureDBaaSUser(p)
	configureDatabaseGrant(p)
	configureDatabaseBackup(p)
	configureKMS(p)
	configureScheduleJob(p)
	configureProject(p)
	configureVPC(p)
	configureSubnet(p)
	configureKeyPair(p)
	configureElasticIP(p)
	configureBlockStorage(p)
	configureSnapshot(p)
	configureBackup(p)
	configureRestore(p)
	configureCloudServer(p)
	configureSecurityGroup(p)
	configureSecurityRule(p)
	configureVPCPeering(p)
	configureVPCPeeringRoute(p)
	configureVPNTunnel(p)
	configureVPNRoute(p)
}

func configureProject(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfProject, func(r *ujconfig.Resource) {
		// Project is the root resource — no incoming references.
	})
}

func configureVPC(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfVPC, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: tfProject,
		}
	})
}

func configureSubnet(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfSubnet, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: tfProject,
		}
		r.References["vpc_id"] = ujconfig.Reference{
			TerraformName: tfVPC,
		}
	})
}

func configureKeyPair(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfKeyPair, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: tfProject,
		}
	})
}

func configureElasticIP(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfElasticIP, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: tfProject,
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
	p.AddResourceConfigurator(tfBlockStorage, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: tfProject,
		}
	})
}

func configureCloudServer(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_cloudserver", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: tfProject,
		}
		// URI-based references: extract status.atProvider.uri from the referenced resource.
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

func configureSnapshot(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_snapshot", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		// volume_uri references the full URI of an existing BlockStorage volume.
		r.References["volume_uri"] = ujconfig.Reference{
			TerraformName: tfBlockStorage,
			Extractor:     uriExtractor,
		}
	})
}

func configureBackup(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfBackup, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["volume_id"] = ujconfig.Reference{TerraformName: tfBlockStorage}
	})
}

func configureRestore(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_restore", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["backup_id"] = ujconfig.Reference{TerraformName: tfBackup}
		r.References["volume_id"] = ujconfig.Reference{TerraformName: tfBlockStorage}
	})
}

func configureSecurityGroup(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfSecurityGroup, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["vpc_id"] = ujconfig.Reference{TerraformName: tfVPC}
	})
}

func configureSecurityRule(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_securityrule", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["vpc_id"] = ujconfig.Reference{TerraformName: tfVPC}
		r.References["security_group_id"] = ujconfig.Reference{TerraformName: tfSecurityGroup}
	})
}

func configureVPCPeering(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfVPCPeering, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["vpc_id"] = ujconfig.Reference{TerraformName: tfVPC}
		// peer_vpc is optional and may reference a VPC in another project;
		// configure as a reference but keep optional.
		r.References["peer_vpc"] = ujconfig.Reference{TerraformName: tfVPC}
	})
}

func configureVPCPeeringRoute(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_vpcpeeringroute", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["vpc_id"] = ujconfig.Reference{TerraformName: tfVPC}
		r.References["vpc_peering_id"] = ujconfig.Reference{TerraformName: tfVPCPeering}
	})
}

func configureVPNTunnel(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfVPNTunnel, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		// VPN tunnel references VPC, Subnet, and ElasticIP by ID (not URI)
		// inside the nested properties.ip_configurations structure.
		r.References["properties.ip_configurations.vpc.id"] = ujconfig.Reference{
			TerraformName: tfVPC,
		}
		r.References["properties.ip_configurations.subnet.id"] = ujconfig.Reference{
			TerraformName: tfSubnet,
		}
		r.References["properties.ip_configurations.public_ip.id"] = ujconfig.Reference{
			TerraformName: tfElasticIP,
		}
	})
}

func configureVPNRoute(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_vpnroute", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["vpn_tunnel_id"] = ujconfig.Reference{TerraformName: tfVPNTunnel}
	})
}

func configureKaaS(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_kaas", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["network.vpc_uri_ref"] = ujconfig.Reference{
			TerraformName: tfVPC,
			Extractor:     uriExtractor,
		}
		r.References["network.subnet_uri_ref"] = ujconfig.Reference{
			TerraformName: tfSubnet,
			Extractor:     uriExtractor,
		}
		// security_group_name is a display-name reference (not URI).
		// nameExtractor reads spec.forProvider.name from the SecurityGroup resource.
		r.References["network.security_group_name"] = ujconfig.Reference{
			TerraformName: tfSecurityGroup,
			Extractor:     nameExtractor,
		}
		// pod_cidr is user-controlled; the Terraform provider never overwrites it
		// from the API response, so we must not late-initialize it either.
		r.LateInitializer = ujconfig.LateInitializer{
			IgnoredFields: []string{"network.pod_cidr"},
		}
		// Expose kubeconfig and management_ip as connection details.
		// kubeconfig may already be auto-mapped by Upjet if it is Sensitive in the
		// Terraform schema; adding it here ensures it is always present.
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if kc, ok := attr["kubeconfig"].(string); ok && kc != "" {
				conn["kubeconfig"] = []byte(kc)
			}
			if ip, ok := attr["management_ip"].(string); ok && ip != "" {
				conn["management_ip"] = []byte(ip)
			}
			return conn, nil
		}
	})
}

func configureContainerRegistry(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_containerregistry", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["network.public_ip_uri_ref"] = ujconfig.Reference{
			TerraformName: tfElasticIP,
			Extractor:     uriExtractor,
		}
		r.References["network.vpc_uri_ref"] = ujconfig.Reference{
			TerraformName: tfVPC,
			Extractor:     uriExtractor,
		}
		r.References["network.subnet_uri_ref"] = ujconfig.Reference{
			TerraformName: tfSubnet,
			Extractor:     uriExtractor,
		}
		r.References["network.security_group_uri_ref"] = ujconfig.Reference{
			TerraformName: tfSecurityGroup,
			Extractor:     uriExtractor,
		}
		r.References["storage.block_storage_uri_ref"] = ujconfig.Reference{
			TerraformName: tfBlockStorage,
			Extractor:     uriExtractor,
		}
	})
}

// ── Phase 6 — Databases ───────────────────────────────────────────────────────

func configureDBaaS(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfDBaaS, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["network.vpc_uri_ref"] = ujconfig.Reference{
			TerraformName: tfVPC,
			Extractor:     uriExtractor,
		}
		r.References["network.subnet_uri_ref"] = ujconfig.Reference{
			TerraformName: tfSubnet,
			Extractor:     uriExtractor,
		}
		r.References["network.security_group_uri_ref"] = ujconfig.Reference{
			TerraformName: tfSecurityGroup,
			Extractor:     uriExtractor,
		}
		r.References["network.elastic_ip_uri_ref"] = ujconfig.Reference{
			TerraformName: tfElasticIP,
			Extractor:     uriExtractor,
		}
	})
}

func configureDatabase(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfDatabase, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["dbaas_id"] = ujconfig.Reference{TerraformName: tfDBaaS}
	})
}

func configureDBaaSUser(p *ujconfig.Provider) {
	p.AddResourceConfigurator(tfDBaaSUser, func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["dbaas_id"] = ujconfig.Reference{TerraformName: tfDBaaS}
		// password is already handled as a SecretKeySelector (PasswordSecretRef)
		// and GetConnectionDetailsMapping is auto-generated by Upjet.
	})
}

func configureDatabaseGrant(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_databasegrant", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["dbaas_id"] = ujconfig.Reference{TerraformName: tfDBaaS}
		// database holds the database name; Database.id == Database.name,
		// so the default (external name) extractor resolves correctly.
		r.References["database"] = ujconfig.Reference{TerraformName: tfDatabase}
		r.References["user_id"] = ujconfig.Reference{TerraformName: tfDBaaSUser}
	})
}

func configureDatabaseBackup(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_databasebackup", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
		r.References["dbaas_id"] = ujconfig.Reference{TerraformName: tfDBaaS}
		r.References["database"] = ujconfig.Reference{TerraformName: tfDatabase}
	})
}

// ── Phase 7 — Security & Scheduling ──────────────────────────────────────────

func configureKMS(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_kms", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
	})
}

func configureScheduleJob(p *ujconfig.Provider) {
	p.AddResourceConfigurator("arubacloud_schedulejob", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{TerraformName: tfProject}
	})
}
