package valence

// GetReleaseConditions retrieves release conditions for a target in an org unit.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/conditionalRelease/conditions/{targetType}/{targetId}
func (c *Client) GetReleaseConditions(orgUnitId int64, targetType string, targetId int64) (*ReleaseConditionsData, error) {
	var out ReleaseConditionsData
	err := c.get(c.lpPath("%d/conditionalRelease/conditions/%s/%d", orgUnitId, targetType, targetId), nil, &out)
	return &out, err
}
