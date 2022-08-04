package providers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath     = "lbaas"
	resourcePath = "providers"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}
