package valence

// GetVersions returns the supported API versions for all products.
func (c *Client) GetVersions() ([]ProductVersions, error) {
	var out []ProductVersions
	err := c.get("/d2l/api/versions/", nil, &out)
	return out, err
}
