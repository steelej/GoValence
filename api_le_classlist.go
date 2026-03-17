package valence

// GetClasslist returns the classlist for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/classlist/
func (c *Client) GetClasslist(orgUnitId int64) ([]ClasslistUser, error) {
	var out []ClasslistUser
	err := c.get(c.lePath("%d/classlist/", orgUnitId), nil, &out)
	return out, err
}
