package flavors

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("flavors")
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("flavors", id)
}
