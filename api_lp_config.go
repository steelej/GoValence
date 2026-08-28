package valence

// GetConfigVariableValue returns the value of a config variable for an org unit.
// GET /d2l/api/lp/{lpVersion}/configVariables/{variableUUID}/values/orgUnits/{orgUnitId}
func (c *Client) GetConfigVariableValue(variableUUID string, orgUnitId int64) (*ConfigVariableValue, error) {
	var out ConfigVariableValue
	err := c.get(c.lpPath("configVariables/%s/values/orgUnits/%d", variableUUID, orgUnitId), nil, &out)
	return &out, err
}

// GetEffectiveConfigVariableValue returns the resolved value of a config variable for an org unit.
// GET /d2l/api/lp/{lpVersion}/configVariables/{variableUUID}/effectiveValues/orgUnits/{orgUnitId}
func (c *Client) GetEffectiveConfigVariableValue(variableUUID string, orgUnitId int64) (*ConfigVariableValue, error) {
	var out ConfigVariableValue
	err := c.get(c.lpPath("configVariables/%s/effectiveValues/orgUnits/%d", variableUUID, orgUnitId), nil, &out)
	return &out, err
}

// UpdateConfigVariableValue updates a config variable value for an org unit.
// PUT /d2l/api/lp/{lpVersion}/configVariables/{variableUUID}/values/orgUnits/{orgUnitId}
func (c *Client) UpdateConfigVariableValue(variableUUID string, orgUnitId int64, data SpecifiedOrgUnitValue) (*ConfigVariableValue, error) {
	var out ConfigVariableValue
	err := c.putJSON(c.lpPath("configVariables/%s/values/orgUnits/%d", variableUUID, orgUnitId), nil, data, &out)
	return &out, err
}
