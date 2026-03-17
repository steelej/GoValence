package valence

// GetOrgUnitTools returns the tools available for an org unit.
// GET /d2l/api/lp/{lpVersion}/tools/orgUnits/{orgUnitId}
func (c *Client) GetOrgUnitTools(orgUnitId int64) ([]ToolInfo, error) {
	var out []ToolInfo
	err := c.get(c.lpPath("tools/orgUnits/%d", orgUnitId), nil, &out)
	return out, err
}

// GetOrgUnitTool returns a specific tool for an org unit.
// GET /d2l/api/lp/{lpVersion}/tools/orgUnits/{orgUnitId}/{toolId}
func (c *Client) GetOrgUnitTool(orgUnitId, toolId int64) (*ToolInfo, error) {
	var out ToolInfo
	err := c.get(c.lpPath("tools/orgUnits/%d/%d", orgUnitId, toolId), nil, &out)
	return &out, err
}
