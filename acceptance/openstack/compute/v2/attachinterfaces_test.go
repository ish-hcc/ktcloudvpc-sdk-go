//go:build acceptance || compute || servers
// +build acceptance compute servers

package v2

import (
	"testing"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/acceptance/clients"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/acceptance/tools"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/compute/v2/servers"
	th "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
)

func TestAttachDetachInterface(t *testing.T) {
	clients.RequireLong(t)

	choices, err := clients.AcceptanceTestChoicesFromEnv()
	th.AssertNoErr(t, err)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	server, err := CreateServer(t, client)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	iface, err := AttachInterface(t, client, server.ID)
	th.AssertNoErr(t, err)
	defer DetachInterface(t, client, server.ID, iface.PortID)

	tools.PrintResource(t, iface)

	server, err = servers.Get(client, server.ID).Extract()
	th.AssertNoErr(t, err)

	var found bool
	for _, networkAddresses := range server.Addresses[choices.NetworkName].([]interface{}) {
		address := networkAddresses.(map[string]interface{})
		if address["OS-EXT-IPS:type"] == "fixed" {
			fixedIP := address["addr"].(string)

			for _, v := range iface.FixedIPs {
				if fixedIP == v.IPAddress {
					found = true
				}
			}
		}
	}

	th.AssertEquals(t, found, true)
}
