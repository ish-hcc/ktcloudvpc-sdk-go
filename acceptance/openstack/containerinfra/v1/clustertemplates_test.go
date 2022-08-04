//go:build acceptance || containerinfra
// +build acceptance containerinfra

package v1

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/clients"
	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/tools"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/containerinfra/v1/clustertemplates"
	th "github.com/innodreamer/ktvpc-sdk_poc/testhelper"
)

func TestClusterTemplatesCRUD(t *testing.T) {
	client, err := clients.NewContainerInfraV1Client()
	th.AssertNoErr(t, err)

	clusterTemplate, err := CreateKubernetesClusterTemplate(t, client)
	th.AssertNoErr(t, err)
	t.Log(clusterTemplate.Name)

	defer DeleteClusterTemplate(t, client, clusterTemplate.UUID)

	// Test clusters list
	allPages, err := clustertemplates.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allClusterTemplates, err := clustertemplates.ExtractClusterTemplates(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, v := range allClusterTemplates {
		if v.UUID == clusterTemplate.UUID {
			found = true
		}
	}

	th.AssertEquals(t, found, true)

	template, err := clustertemplates.Get(client, clusterTemplate.UUID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, clusterTemplate.UUID, template.UUID)

	// Test cluster update
	updateOpts := []clustertemplates.UpdateOptsBuilder{
		clustertemplates.UpdateOpts{
			Op:    clustertemplates.ReplaceOp,
			Path:  "/master_lb_enabled",
			Value: "false",
		},
		clustertemplates.UpdateOpts{
			Op:    clustertemplates.ReplaceOp,
			Path:  "/registry_enabled",
			Value: "false",
		},
		clustertemplates.UpdateOpts{
			Op:    clustertemplates.AddOp,
			Path:  "/labels/test",
			Value: "test",
		},
	}

	updateClusterTemplate, err := clustertemplates.Update(client, clusterTemplate.UUID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, false, updateClusterTemplate.MasterLBEnabled)
	th.AssertEquals(t, false, updateClusterTemplate.RegistryEnabled)
	th.AssertEquals(t, "test", updateClusterTemplate.Labels["test"])
	tools.PrintResource(t, updateClusterTemplate)

}