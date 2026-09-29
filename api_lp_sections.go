package valence

// CreateSection creates a new section in an org unit.
// PUT /d2l/api/lp/{lpVersion}/{orgUnitId}/sections/
func (c *Client) InitalizeSections(orgUnitId int64, data SectionSettingsData) (SectionSettingsData, error) {
	var out SectionSettingsData
	err := c.postJSON(c.lpPath("%d/sections/settings", orgUnitId), nil, data, &out)
	return out, err
}

// GetSections returns all sections for an org unit.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/sections/
func (c *Client) GetSections(orgUnitId int64) ([]Section, error) {
	var out []Section
	err := c.get(c.lpPath("%d/sections/", orgUnitId), nil, &out)
	return out, err
}

// CreateSection creates a new section in an org unit.
// POST /d2l/api/lp/{lpVersion}/{orgUnitId}/sections/
func (c *Client) CreateSection(orgUnitId int64, data SectionData) (*Section, error) {
	var out Section
	err := c.postJSON(c.lpPath("%d/sections/", orgUnitId), nil, data, &out)
	return &out, err
}

// GetSection returns a specific section.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/sections/{sectionId}
func (c *Client) GetSection(orgUnitId, sectionId int64) (*Section, error) {
	var out Section
	err := c.get(c.lpPath("%d/sections/%d", orgUnitId, sectionId), nil, &out)
	return &out, err
}

// GetSectionEnrollments returns the enrollments for a specific section.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/sections/{sectionId}/enrollments/
func (c *Client) GetSectionEnrollments(orgUnitId, sectionId int64) ([]GroupEnrollment, error) {
	var out []GroupEnrollment
	err := c.get(c.lpPath("%d/sections/%d/enrollments/", orgUnitId, sectionId), nil, &out)
	return out, err
}

// DeleteSection deletes the specified section.
// DELETE /d2l/api/lp/{lpVersion}/{orgUnitId}/sections/{sectionId}
func (c *Client) DeleteSection(orgUnitId, sectionId int64) error {
	return c.delete(c.lpPath("%d/sections/%d", orgUnitId, sectionId), nil, nil)
}
