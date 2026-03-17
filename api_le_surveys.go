package valence

import "net/url"

// GetSurveys returns all surveys for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/surveys/
func (c *Client) GetSurveys(orgUnitId int64, params url.Values) (*PagedResultSet[Survey], error) {
	var out PagedResultSet[Survey]
	err := c.get(c.lePath("%d/surveys/", orgUnitId), params, &out)
	return &out, err
}

// GetSurvey returns a specific survey.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/surveys/{surveyId}
func (c *Client) GetSurvey(orgUnitId, surveyId int64) (*Survey, error) {
	var out Survey
	err := c.get(c.lePath("%d/surveys/%d", orgUnitId, surveyId), nil, &out)
	return &out, err
}

// GetSurveyAttempts returns all attempts for a survey.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/surveys/{surveyId}/attempts/
func (c *Client) GetSurveyAttempts(orgUnitId, surveyId int64, params url.Values) (*PagedResultSet[SurveyAttempt], error) {
	var out PagedResultSet[SurveyAttempt]
	err := c.get(c.lePath("%d/surveys/%d/attempts/", orgUnitId, surveyId), params, &out)
	return &out, err
}

// GetSurveyAttempt returns a specific survey attempt.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/surveys/{surveyId}/attempts/{attemptId}
func (c *Client) GetSurveyAttempt(orgUnitId, surveyId, attemptId int64) (*SurveyAttempt, error) {
	var out SurveyAttempt
	err := c.get(c.lePath("%d/surveys/%d/attempts/%d", orgUnitId, surveyId, attemptId), nil, &out)
	return &out, err
}
