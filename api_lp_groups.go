package valence

// GetGroupCategories returns all group categories for an org unit.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/
func (c *Client) GetGroupCategories(orgUnitId int64) ([]GroupCategory, error) {
	var out []GroupCategory
	err := c.get(c.lpPath("%d/groupcategories/", orgUnitId), nil, &out)
	return out, err
}

// CreateGroupCategory starts creation of a new group category for an org unit.
// POST /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/
func (c *Client) CreateGroupCategory(orgUnitId int64, data GroupCategoryData) (*GroupsJobData, error) {
	var out GroupsJobData
	err := c.postJSON(c.lpPath("%d/groupcategories/", orgUnitId), nil, data, &out)
	return &out, err
}

// GetGroupCategory returns a specific group category.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}
func (c *Client) GetGroupCategory(orgUnitId, groupCategoryId int64) (*GroupCategory, error) {
	var out GroupCategory
	err := c.get(c.lpPath("%d/groupcategories/%d", orgUnitId, groupCategoryId), nil, &out)
	return &out, err
}

// GetGroupCategoryStatus returns the async creation status for a group category.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}/status
func (c *Client) GetGroupCategoryStatus(orgUnitId, groupCategoryId int64) (*GroupCategoryJobStatus, error) {
	var out GroupCategoryJobStatus
	err := c.get(c.lpPath("%d/groupcategories/%d/status", orgUnitId, groupCategoryId), nil, &out)
	return &out, err
}

// UpdateGroupCategory replaces a group category for an org unit.
// PUT /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}
func (c *Client) UpdateGroupCategory(orgUnitId, groupCategoryId int64, data GroupCategoryData) (*GroupCategory, error) {
	var out GroupCategory
	err := c.putJSON(c.lpPath("%d/groupcategories/%d", orgUnitId, groupCategoryId), nil, data, &out)
	return &out, err
}

// GetGroups returns all groups in a group category.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}/groups/
func (c *Client) GetGroups(orgUnitId, groupCategoryId int64) ([]Group, error) {
	var out []Group
	err := c.get(c.lpPath("%d/groupcategories/%d/groups/", orgUnitId, groupCategoryId), nil, &out)
	return out, err
}

// CreateGroup creates a new group within a group category.
// POST /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}/groups/
func (c *Client) CreateGroup(orgUnitId, groupCategoryId int64, data GroupData) (*Group, error) {
	var out Group
	err := c.postJSON(c.lpPath("%d/groupcategories/%d/groups/", orgUnitId, groupCategoryId), nil, data, &out)
	return &out, err
}

// DeleteGroup deletes a group within a group category.
// DELETE /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}/groups/{groupId}
func (c *Client) DeleteGroup(orgUnitId, groupCategoryId, groupId int64) error {
	return c.delete(c.lpPath("%d/groupcategories/%d/groups/%d", orgUnitId, groupCategoryId, groupId), nil)
}

// GetGroup returns a specific group.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}/groups/{groupId}
func (c *Client) GetGroup(orgUnitId, groupCategoryId, groupId int64) (*Group, error) {
	var out Group
	err := c.get(c.lpPath("%d/groupcategories/%d/groups/%d", orgUnitId, groupCategoryId, groupId), nil, &out)
	return &out, err
}

// UpdateGroup replaces a group within a group category.
// PUT /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}/groups/{groupId}
func (c *Client) UpdateGroup(orgUnitId, groupCategoryId, groupId int64, data GroupData) (*Group, error) {
	var out Group
	err := c.putJSON(c.lpPath("%d/groupcategories/%d/groups/%d", orgUnitId, groupCategoryId, groupId), nil, data, &out)
	return &out, err
}

// GetGroupEnrollments returns the enrollments for a specific group.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}/groups/{groupId}/enrollments/
func (c *Client) GetGroupEnrollments(orgUnitId, groupCategoryId, groupId int64) ([]GroupEnrollment, error) {
	var out []GroupEnrollment
	err := c.get(c.lpPath("%d/groupcategories/%d/groups/%d/enrollments/", orgUnitId, groupCategoryId, groupId), nil, &out)
	return out, err
}

// EnrollUserInGroup enrolls a user in a group.
// POST /d2l/api/lp/{lpVersion}/{orgUnitId}/groupcategories/{groupCategoryId}/groups/{groupId}/enrollments/
func (c *Client) EnrollUserInGroup(orgUnitId, groupCategoryId, groupId int64, data GroupEnrollment) error {
	return c.postJSON(c.lpPath("%d/groupcategories/%d/groups/%d/enrollments/", orgUnitId, groupCategoryId, groupId), nil, data, nil)
}
