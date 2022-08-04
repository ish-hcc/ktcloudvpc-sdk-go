package testing

import (
	"testing"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/placement/v1/resourceproviders"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/pagination"
	th "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
	fake "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper/client"
)

func TestListResourceProviders(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleResourceProviderList(t)

	count := 0
	err := resourceproviders.List(fake.ServiceClient(), resourceproviders.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := resourceproviders.ExtractResourceProviders(page)
		if err != nil {
			t.Errorf("Failed to extract resource providers: %v", err)
			return false, err
		}
		th.AssertDeepEquals(t, ExpectedResourceProviders, actual)

		return true, nil
	})

	th.AssertNoErr(t, err)

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestCreateResourceProvider(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleResourceProviderCreate(t)

	expected := ExpectedResourceProvider1

	opts := resourceproviders.CreateOpts{
		Name:               ExpectedResourceProvider1.Name,
		UUID:               ExpectedResourceProvider1.UUID,
		ParentProviderUUID: ExpectedResourceProvider1.ParentProviderUUID,
	}

	actual, err := resourceproviders.Create(fake.ServiceClient(), opts).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &expected, actual)
}

func TestGetResourceProvider(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleResourceProviderGet(t)

	expected := ExpectedResourceProvider1

	actual, err := resourceproviders.Get(fake.ServiceClient(), ExpectedResourceProvider1.UUID).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &expected, actual)
}

func TestDeleteResourceProvider(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleResourceProviderDelete(t)

	res := resourceproviders.Delete(fake.ServiceClient(), "b99b3ab4-3aa6-4fba-b827-69b88b9c544a")
	th.AssertNoErr(t, res.Err)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleResourceProviderUpdate(t)

	name := "new_name"
	parentProviderUUID := "b99b3ab4-3aa6-4fba-b827-69b88b9c544a"

	options := resourceproviders.UpdateOpts{
		Name:               &name,
		ParentProviderUUID: &parentProviderUUID,
	}
	rp, err := resourceproviders.Update(fake.ServiceClient(), "4e8e5957-649f-477b-9e5b-f1f75b21c03c", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, rp.Name, name)
	th.AssertEquals(t, rp.ParentProviderUUID, parentProviderUUID)
}

func TestGetResourceProvidersUsages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleResourceProviderGetUsages(t)

	actual, err := resourceproviders.GetUsages(fake.ServiceClient(), ResourceProviderTestID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedUsages, *actual)
}

func TestGetResourceProvidersInventories(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleResourceProviderGetInventories(t)

	actual, err := resourceproviders.GetInventories(fake.ServiceClient(), ResourceProviderTestID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedInventories, *actual)
}

func TestGetResourceProvidersAllocations(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleResourceProviderGetAllocations(t)

	actual, err := resourceproviders.GetAllocations(fake.ServiceClient(), ResourceProviderTestID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedAllocations, *actual)
}

func TestGetResourceProvidersTraits(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleResourceProviderGetTraits(t)

	actual, err := resourceproviders.GetTraits(fake.ServiceClient(), ResourceProviderTestID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedTraits, *actual)
}
