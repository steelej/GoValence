package valence

import "net/url"

// GetMyEnrollments returns the current user's enrollments.
// GET /d2l/api/lp/{lpVersion}/enrollments/myenrollments/
func (c *Client) GetMyEnrollments(params url.Values) (*PagedResultSet[MyOrgUnitInfo], error) {
	var out PagedResultSet[MyOrgUnitInfo]
	err := c.get(c.lpPath("enrollments/myenrollments/"), params, &out)
	return &out, err
}

// GetOrgUnitEnrollments returns all users enrolled in an org unit.
// GET /d2l/api/lp/{lpVersion}/enrollments/orgUnits/{orgUnitId}/users/
func (c *Client) GetOrgUnitEnrollments(orgUnitId int64, params url.Values) (*PagedResultSet[OrgUnitUser], error) {
	var out PagedResultSet[OrgUnitUser]
	err := c.get(c.lpPath("enrollments/orgUnits/%d/users/", orgUnitId), params, &out)
	return &out, err
}

// GetUserOrgUnitEnrollment returns a specific user's enrollment in an org unit.
// GET /d2l/api/lp/{lpVersion}/enrollments/orgUnits/{orgUnitId}/users/{userId}
func (c *Client) GetUserOrgUnitEnrollment(orgUnitId, userId int64) (*EnrollmentData, error) {
	var out EnrollmentData
	err := c.get(c.lpPath("enrollments/orgUnits/%d/users/%d", orgUnitId, userId), nil, &out)
	return &out, err
}

// GetUserEnrollments returns all org unit enrollments for a specific user.
// GET /d2l/api/lp/{lpVersion}/enrollments/users/{userId}/orgUnits/
func (c *Client) GetUserEnrollments(userId int64, params url.Values) (*PagedResultSet[UserEnrollmentData], error) {
	var out PagedResultSet[UserEnrollmentData]
	err := c.get(c.lpPath("enrollments/users/%d/orgUnits/", userId), params, &out)
	return &out, err
}
