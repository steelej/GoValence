package valence

import "net/url"

// GetNewsItems returns news items for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/news/
func (c *Client) GetNewsItems(orgUnitId int64, params url.Values) ([]NewsItem, error) {
	var out []NewsItem
	err := c.get(c.lePath("%d/news/", orgUnitId), params, &out)
	return out, err
}

// GetNewsItem returns a specific news item.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/news/{newsItemId}
func (c *Client) GetNewsItem(orgUnitId, newsItemId int64) (*NewsItem, error) {
	var out NewsItem
	err := c.get(c.lePath("%d/news/%d", orgUnitId, newsItemId), nil, &out)
	return &out, err
}
