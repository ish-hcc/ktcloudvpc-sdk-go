package testing

import (
	"testing"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/clustering/v1/events"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/pagination"
	th "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
	fake "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper/client"
)

func TestListEvents(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleListSuccessfully(t)

	pageCount := 0
	err := events.List(fake.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		pageCount++
		actual, err := events.ExtractEvents(page)
		th.AssertNoErr(t, err)

		th.AssertDeepEquals(t, ExpectedEvents, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)

	if pageCount != 1 {
		t.Errorf("Expected 1 page, got %d", pageCount)
	}
}

func TestGetEvent(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleGetSuccessfully(t, ExpectedEvent1.ID)

	actual, err := events.Get(fake.ServiceClient(), ExpectedEvent1.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedEvent1, *actual)
}
