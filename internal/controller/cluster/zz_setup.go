// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backup "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/backup"
	blockstorage "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/blockstorage"
	cloudserver "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/cloudserver"
	containerregistry "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/containerregistry"
	database "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/database"
	databasebackup "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/databasebackup"
	databasegrant "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/databasegrant"
	dbaas "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/dbaas"
	dbaasuser "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/dbaasuser"
	elasticip "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/elasticip"
	kaas "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/kaas"
	keypair "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/keypair"
	kms "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/kms"
	project "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/project"
	restore "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/restore"
	schedulejob "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/schedulejob"
	securitygroup "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/securitygroup"
	securityrule "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/securityrule"
	snapshot "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/snapshot"
	subnet "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/subnet"
	vpc "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/vpc"
	vpcpeering "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/vpcpeering"
	vpcpeeringroute "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/vpcpeeringroute"
	vpnroute "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/vpnroute"
	vpntunnel "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/arubacloud/vpntunnel"
	providerconfig "github.com/arubacloud/crossplane-provider-arubacloud/internal/controller/cluster/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backup.Setup,
		blockstorage.Setup,
		cloudserver.Setup,
		containerregistry.Setup,
		database.Setup,
		databasebackup.Setup,
		databasegrant.Setup,
		dbaas.Setup,
		dbaasuser.Setup,
		elasticip.Setup,
		kaas.Setup,
		keypair.Setup,
		kms.Setup,
		project.Setup,
		restore.Setup,
		schedulejob.Setup,
		securitygroup.Setup,
		securityrule.Setup,
		snapshot.Setup,
		subnet.Setup,
		vpc.Setup,
		vpcpeering.Setup,
		vpcpeeringroute.Setup,
		vpnroute.Setup,
		vpntunnel.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backup.SetupGated,
		blockstorage.SetupGated,
		cloudserver.SetupGated,
		containerregistry.SetupGated,
		database.SetupGated,
		databasebackup.SetupGated,
		databasegrant.SetupGated,
		dbaas.SetupGated,
		dbaasuser.SetupGated,
		elasticip.SetupGated,
		kaas.SetupGated,
		keypair.SetupGated,
		kms.SetupGated,
		project.SetupGated,
		restore.SetupGated,
		schedulejob.SetupGated,
		securitygroup.SetupGated,
		securityrule.SetupGated,
		snapshot.SetupGated,
		subnet.SetupGated,
		vpc.SetupGated,
		vpcpeering.SetupGated,
		vpcpeeringroute.SetupGated,
		vpnroute.SetupGated,
		vpntunnel.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		backup.SetupWebhookWithManager,
		blockstorage.SetupWebhookWithManager,
		cloudserver.SetupWebhookWithManager,
		containerregistry.SetupWebhookWithManager,
		database.SetupWebhookWithManager,
		databasebackup.SetupWebhookWithManager,
		databasegrant.SetupWebhookWithManager,
		dbaas.SetupWebhookWithManager,
		dbaasuser.SetupWebhookWithManager,
		elasticip.SetupWebhookWithManager,
		kaas.SetupWebhookWithManager,
		keypair.SetupWebhookWithManager,
		kms.SetupWebhookWithManager,
		project.SetupWebhookWithManager,
		restore.SetupWebhookWithManager,
		schedulejob.SetupWebhookWithManager,
		securitygroup.SetupWebhookWithManager,
		securityrule.SetupWebhookWithManager,
		snapshot.SetupWebhookWithManager,
		subnet.SetupWebhookWithManager,
		vpc.SetupWebhookWithManager,
		vpcpeering.SetupWebhookWithManager,
		vpcpeeringroute.SetupWebhookWithManager,
		vpnroute.SetupWebhookWithManager,
		vpntunnel.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
