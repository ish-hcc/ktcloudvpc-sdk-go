package testing

import (
	"testing"

	az "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/blockstorage/extensions/availabilityzones"
	th "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper/client"
)

// Verifies that availability zones can be listed correctly
func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleGetSuccessfully(t)

	allPages, err := az.List(client.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)

	actual, err := az.ExtractAvailabilityZones(allPages)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, AZResult, actual)
}
