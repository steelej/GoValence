package valence

// GetSelfAssessments returns all self assessments for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/selfassessments/
func (c *Client) GetSelfAssessments(orgUnitId int64) ([]SelfAssessment, error) {
	var out []SelfAssessment
	err := c.get(c.lePath("%d/selfassessments/", orgUnitId), nil, &out)
	return out, err
}

// GetSelfAssessment returns a specific self assessment.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/selfassessments/{selfAssessmentId}
func (c *Client) GetSelfAssessment(orgUnitId, selfAssessmentId int64) (*SelfAssessment, error) {
	var out SelfAssessment
	err := c.get(c.lePath("%d/selfassessments/%d", orgUnitId, selfAssessmentId), nil, &out)
	return &out, err
}

// GetSelfAssessmentAttempts returns all attempts for a self assessment.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/selfassessments/{selfAssessmentId}/attempts/
func (c *Client) GetSelfAssessmentAttempts(orgUnitId, selfAssessmentId int64) ([]SelfAssessmentAttempt, error) {
	var out []SelfAssessmentAttempt
	err := c.get(c.lePath("%d/selfassessments/%d/attempts/", orgUnitId, selfAssessmentId), nil, &out)
	return out, err
}
