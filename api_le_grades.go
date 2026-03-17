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

// GetGradeValue returns a specific user's value for a grade object.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/{gradeObjectId}/values/{userId}
func (c *Client) GetGradeValue(orgUnitId, gradeObjectId, userId int64) (*GradeValue, error) {
	var out GradeValue
	err := c.get(c.lePath("%d/grades/%d/values/%d", orgUnitId, gradeObjectId, userId), nil, &out)
	return &out, err
}

// GetGradeValues returns grade values for all users for a grade object.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/grades/{gradeObjectId}/values/
func (c *Client) GetGradeValues(orgUnitId, gradeObjectId int64, params url.Values) (*PagedResultSet[GradeValue], error) {
	var out PagedResultSet[GradeValue]
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
