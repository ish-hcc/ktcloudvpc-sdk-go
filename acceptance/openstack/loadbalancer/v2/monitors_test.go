//go:build acceptance || networking || loadbalancer || monitors
// +build acceptance networking loadbalancer monitors

package v2

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/clients"
	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/tools"
	"github.com/innodreamer/ktvpc-sdk_poc/openstack/loadbalancer/v2/monitors"
)

func TestMonitorsList(t *testing.T) {
	client, err := clients.NewLoadBalancerV2Client()
	if err != nil {
		t.Fatalf("Unable to create a loadbalancer client: %v", err)
	}

	allPages, err := monitors.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list monitors: %v", err)
	}

	allMonitors, err := monitors.ExtractMonitors(allPages)
	if err != nil {
		t.Fatalf("Unable to extract monitors: %v", err)
	}

	for _, monitor := range allMonitors {
		tools.PrintResource(t, monitor)
	}
}