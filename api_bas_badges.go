package valence

import "net/url"

// GetIssuedBadges returns issued badges for a specific user.
// GET /d2l/api/bas/{basVersion}/issued/users/{userId}/
func (c *Client) GetIssuedBadges(userId int64, params url.Values) (*ObjectListPage[IssuedBadge], error) {
	var out ObjectListPage[IssuedBadge]
	err := c.get(c.basPath("issued/users/%d/", userId), params, &out)
	return &out, err
}
