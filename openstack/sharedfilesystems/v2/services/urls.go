package services

import "github.com/innodreamer/ktvpc-sdk_poc"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("services")
}