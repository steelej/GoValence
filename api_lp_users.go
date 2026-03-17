package valence

import "net/url"

// WhoAmI returns the current user's basic info.
// GET /d2l/api/lp/{lpVersion}/users/whoami
func (c *Client) WhoAmI() (*WhoAmIUser, error) {
	var out WhoAmIUser
	err := c.get(c.lpPath("users/whoami"), nil, &out)
	return &out, err
}

// GetUser returns a specific user's data.
// GET /d2l/api/lp/{lpVersion}/users/{userId}
func (c *Client) GetUser(userId int64) (*UserData, error) {
	var out UserData
	err := c.get(c.lpPath("users/%d", userId), nil, &out)
	return &out, err
}

// GetUsers returns a paged list of users, optionally filtered.
// Supported params: orgDefinedId, userName, bookmark.
// GET /d2l/api/lp/{lpVersion}/users/
func (c *Client) GetUsers(params url.Values) (*PagedResultSet[UserData], error) {
	var out PagedResultSet[UserData]
	err := c.get(c.lpPath("users/"), params, &out)
	return &out, err
}
