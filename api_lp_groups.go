package valence

// GetGroupCategories returns all group categories for an org unit.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/
func (c *Client) GetGroupCategories(orgUnitId int64) ([]GroupCategory, error) {
	var out []GroupCategory
	err := c.get(c.lpPath("%d/groupcategories/", orgUnitId), nil, &out)
	return out, err
}

// GetGroupCategory returns a specific group category.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}
func (c *Client) GetGroupCategory(orgUnitId, groupCategoryId int64) (*GroupCategory, error) {
	var out GroupCategory
	err := c.get(c.lpPath("%d/groupcategories/%d", orgUnitId, groupCategoryId), nil, &out)
	return &out, err
}

// GetGroups returns all groups in a group category.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}/groups/
func (c *Client) GetGroups(orgUnitId, groupCategoryId int64) ([]Group, error) {
	var out []Group
	err := c.get(c.lpPath("%d/groupcategories/%d/groups/", orgUnitId, groupCategoryId), nil, &out)
	return out, err
}

// GetGroup returns a specific group.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}/groups/{groupId}
func (c *Client) GetGroup(orgUnitId, groupCategoryId, groupId int64) (*Group, error) {
	var out Group
	err := c.get(c.lpPath("%d/groupcategories/%d/groups/%d", orgUnitId, groupCategoryId, groupId), nil, &out)
	return &out, err
}

// GetGroupEnrollments returns the enrollments for a specific group.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}/groups/{groupId}/enrollments/
func (c *Client) GetGroupEnrollments(orgUnitId, groupCategoryId, groupId int64) ([]GroupEnrollment, error) {
	var out []GroupEnrollment
	err := c.get(c.lpPath("%d/groupcategories/%d/groups/%d/enrollments/", orgUnitId, groupCategoryId, groupId), nil, &out)
	return out, err
}
