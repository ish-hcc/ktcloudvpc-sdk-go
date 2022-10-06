package imageimport

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// ImportMethod represents valid Import API method.
type ImportMethod string

const (
	// GlanceDirectMethod represents glance-direct Import API method.
	GlanceDirectMethod ImportMethod = "glance-direct"

	// WebDownloadMethod represents web-download Import API method.
	WebDownloadMethod ImportMethod = "web-download"
)

// Get retrieves Import API information data.
func Get(c *ktvpcsdk.ServiceClient) (r GetResult) {
	resp, err := c.Get(infoURL(c), &r.Body, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToImportCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies parameters of a new image import.
type CreateOpts struct {
	Name ImportMethod `json:"name"`
	URI  string       `json:"uri"`
}

// ToImportCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToImportCreateMap() (map[string]interface{}, error) {
	b, err := ktvpcsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"method": b}, nil
}

// Create requests the creation of a new image import on the server.
func Create(client *ktvpcsdk.ServiceClient, imageID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToImportCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(importURL(client, imageID), b, nil, &ktvpcsdk.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
