package valence

// GetRoles returns all roles defined in the organization.
// GET /d2l/api/lp/{lpVersion}/roles/
func (c *Client) GetRoles() ([]RoleInfo, error) {
	var out []RoleInfo
	err := c.get(c.lpPath("roles/"), nil, &out)
	return out, err
}

// GetRolesUnstable returns all roles defined in the organization, including
// fields that are only available in the LP unstable contract.
// GET /d2l/api/lp/unstable/roles/
func (c *Client) GetRolesUnstable() ([]Role, error) {
	var out []Role
	err := c.get(c.lpUnstablePath("roles/"), nil, &out)
	return out, err
}

// GetRole returns a specific role by ID.
// GET /d2l/api/lp/{lpVersion}/roles/{roleId}
func (c *Client) GetRole(roleId int64) (*RoleInfo, error) {
	var out RoleInfo
	err := c.get(c.lpPath("roles/%d", roleId), nil, &out)
	return &out, err
}
