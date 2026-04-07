package valence

import "net/url"

// GetTableOfContents returns the full content table of contents for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/content/toc
func (c *Client) GetTableOfContents(orgUnitId int64, params url.Values) (*TableOfContents, error) {
	var out TableOfContents
	err := c.get(c.lePath("%d/content/toc", orgUnitId), params, &out)
	return &out, err
}
