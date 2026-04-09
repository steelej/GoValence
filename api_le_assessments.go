package valence

import "net/url"

// GetRubrics returns rubrics for an org unit, optionally filtered by object type and object id.
// GET /d2l/api/le/unstable/{orgUnitId}/rubrics/
func (c *Client) GetRubrics(orgUnitId int64, params url.Values) ([]Rubric, error) {
	var out []Rubric
	err := c.get(c.leUnstablePath("%d/rubrics/", orgUnitId), params, &out)
	return out, err
}
