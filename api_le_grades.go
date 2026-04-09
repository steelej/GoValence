package valence

import "net/url"

// GetGradeObjects returns all grade objects for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/
func (c *Client) GetGradeObjects(orgUnitId int64) ([]GradeObject, error) {
	var out []GradeObject
	err := c.get(c.lePath("%d/grades/", orgUnitId), nil, &out)
	return out, err
}

// GetGradeCategories returns all grade categories for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/categories/
func (c *Client) GetGradeCategories(orgUnitId int64) ([]GradeCategory, error) {
	var out []GradeCategory
	err := c.get(c.lePath("%d/grades/categories/", orgUnitId), nil, &out)
	return out, err
}

// GetGradeSchemes returns all grade schemes for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/schemes/
func (c *Client) GetGradeSchemes(orgUnitId int64) ([]GradeScheme, error) {
	var out []GradeScheme
	err := c.get(c.lePath("%d/grades/schemes/", orgUnitId), nil, &out)
	return out, err
}

// GetGradeSetup returns gradebook setup information for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/setup/
func (c *Client) GetGradeSetup(orgUnitId int64) (*GradeSetupInfo, error) {
	var out GradeSetupInfo
	err := c.get(c.lePath("%d/grades/setup/", orgUnitId), nil, &out)
	return &out, err
}

// GetGradeValue returns a specific user's value for a grade object.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/{gradeObjectId}/values/{userId}
func (c *Client) GetGradeValue(orgUnitId, gradeObjectId, userId int64) (*GradeValue, error) {
	var out GradeValue
	err := c.get(c.lePath("%d/grades/%d/values/%d", orgUnitId, gradeObjectId, userId), nil, &out)
	return &out, err
}

// GetGradeValues returns grade values for all users for a grade object.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/{gradeObjectId}/values/
func (c *Client) GetGradeValues(orgUnitId, gradeObjectId int64, params url.Values) (*ObjectListPage[GradeValueEntry], error) {
	var out ObjectListPage[GradeValueEntry]
	err := c.get(c.lePath("%d/grades/%d/values/", orgUnitId, gradeObjectId), params, &out)
	return &out, err
}

// GetUserGradeValues returns all grade values for a specific user in an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/values/{userId}/
func (c *Client) GetUserGradeValues(orgUnitId, userId int64) ([]GradeValue, error) {
	var out []GradeValue
	err := c.get(c.lePath("%d/grades/values/%d/", orgUnitId, userId), nil, &out)
	return out, err
}

// GetFinalGradeValue returns the final grade value for a specific user.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/final/values/{userId}
func (c *Client) GetFinalGradeValue(orgUnitId, userId int64) (*FinalGradeValue, error) {
	var out FinalGradeValue
	err := c.get(c.lePath("%d/grades/final/values/%d", orgUnitId, userId), nil, &out)
	return &out, err
}

// GetFinalCalculatedGradeValue returns the final calculated grade for a specific user.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/final/values/{userId}?gradeType=calculated
func (c *Client) GetFinalCalculatedGradeValue(orgUnitId, userId int64) (*FinalGradeValue, error) {
	var out FinalGradeValue
	params := url.Values{"gradeType": []string{"calculated"}}
	err := c.get(c.lePath("%d/grades/final/values/%d", orgUnitId, userId), params, &out)
	return &out, err
}

// GetFinalGradeValues returns final grade values for all users in an org unit (paginated).
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/final/values/
func (c *Client) GetFinalGradeValues(orgUnitId int64, params url.Values) (*ObjectListPage[FinalGradeValueEntry], error) {
	var out ObjectListPage[FinalGradeValueEntry]
	err := c.get(c.lePath("%d/grades/final/values/", orgUnitId), params, &out)
	return &out, err
}
