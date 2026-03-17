package valence

import "net/url"

// GetOrganizationInfo returns top-level organization information.
// GET /d2l/api/lp/{lpVersion}/organization/info
func (c *Client) GetOrganizationInfo() (*OrganizationInfo, error) {
	var out OrganizationInfo
	err := c.get(c.lpPath("organization/info"), nil, &out)
	return &out, err
}

// GetOrgUnit returns a specific org unit by ID.
// GET /d2l/api/lp/{lpVersion}/orgstructure/{orgUnitId}
func (c *Client) GetOrgUnit(orgUnitId int64) (*OrgUnitProperties, error) {
	var out OrgUnitProperties
	err := c.get(c.lpPath("orgstructure/%d", orgUnitId), nil, &out)
	return &out, err
}

// GetOrgUnits returns a paged list of org units.
// GET /d2l/api/lp/{lpVersion}/orgstructure/
func (c *Client) GetOrgUnits(params url.Values) (*PagedResultSet[OrgUnitProperties], error) {
	var out PagedResultSet[OrgUnitProperties]
	err := c.get(c.lpPath("orgstructure/"), params, &out)
	return &out, err
}

// GetOrgUnitParents returns the parents of a given org unit.
// GET /d2l/api/lp/{lpVersion}/orgstructure/{orgUnitId}/parents/
func (c *Client) GetOrgUnitParents(orgUnitId int64) ([]OrgUnitProperties, error) {
	var out []OrgUnitProperties
	err := c.get(c.lpPath("orgstructure/%d/parents/", orgUnitId), nil, &out)
	return out, err
}

// GetOrgRecycleBin returns org units in the recycle bin.
// GET /d2l/api/lp/{lpVersion}/orgstructure/recyclebin/
func (c *Client) GetOrgRecycleBin(params url.Values) (*PagedResultSet[OrgUnitProperties], error) {
	var out PagedResultSet[OrgUnitProperties]
	err := c.get(c.lpPath("orgstructure/recyclebin/"), params, &out)
	return &out, err
}

// GetSemesterOrgUnitType returns the org unit type info for semesters.
// GET /d2l/api/lp/{lpVersion}/outypes/semester
func (c *Client) GetSemesterOrgUnitType() (*OrgUnitTypeInfo, error) {
	var out OrgUnitTypeInfo
	err := c.get(c.lpPath("outypes/semester"), nil, &out)
	return &out, err
}
