package valence

// GetConfigVariableValue returns the value of a config variable for an org unit.
// GET /d2l/api/lp/{lpVersion}/configVariables/{variableUUID}/values/orgUnits/{orgUnitId}
func (c *Client) GetConfigVariableValue(variableUUID string, orgUnitId int64) (*ConfigVariableValue, error) {
	var out ConfigVariableValue
	err := c.get(c.lpPath("configVariables/%s/values/orgUnits/%d", variableUUID, orgUnitId), nil, &out)
	return &out, err
}
