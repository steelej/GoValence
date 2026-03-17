package valence

import "net/url"

// GetQuizzes returns all quizzes for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/quizzes/
func (c *Client) GetQuizzes(orgUnitId int64, params url.Values) (*PagedResultSet[QuizReadData], error) {
	var out PagedResultSet[QuizReadData]
	err := c.get(c.lePath("%d/quizzes/", orgUnitId), params, &out)
	return &out, err
}

// GetQuiz returns a specific quiz.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/quizzes/{quizId}
func (c *Client) GetQuiz(orgUnitId, quizId int64) (*QuizReadData, error) {
	var out QuizReadData
	err := c.get(c.lePath("%d/quizzes/%d", orgUnitId, quizId), nil, &out)
	return &out, err
}

// GetQuizAttempts returns all attempts for a quiz.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/quizzes/{quizId}/attempts/
func (c *Client) GetQuizAttempts(orgUnitId, quizId int64, params url.Values) (*PagedResultSet[QuizAttemptData], error) {
	var out PagedResultSet[QuizAttemptData]
	err := c.get(c.lePath("%d/quizzes/%d/attempts/", orgUnitId, quizId), params, &out)
	return &out, err
}

// GetQuizAttempt returns a specific quiz attempt.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/quizzes/{quizId}/attempts/{attemptId}
func (c *Client) GetQuizAttempt(orgUnitId, quizId, attemptId int64) (*QuizAttemptData, error) {
	var out QuizAttemptData
	err := c.get(c.lePath("%d/quizzes/%d/attempts/%d", orgUnitId, quizId, attemptId), nil, &out)
	return &out, err
}

// GetQuizQuestions returns all questions for a quiz.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/quizzes/{quizId}/questions/
func (c *Client) GetQuizQuestions(orgUnitId, quizId int64, params url.Values) (*PagedResultSet[QuizQuestion], error) {
	var out PagedResultSet[QuizQuestion]
	err := c.get(c.lePath("%d/quizzes/%d/questions/", orgUnitId, quizId), params, &out)
	return &out, err
}

// GetQuizQuestion returns a specific quiz question.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/quizzes/{quizId}/questions/{questionId}
func (c *Client) GetQuizQuestion(orgUnitId, quizId, questionId int64) (*QuizQuestion, error) {
	var out QuizQuestion
	err := c.get(c.lePath("%d/quizzes/%d/questions/%d", orgUnitId, quizId, questionId), nil, &out)
	return &out, err
}

// GetQuizSpecialAccess returns special access settings for a user on a quiz.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/quizzes/{quizId}/specialaccess/{userId}
func (c *Client) GetQuizSpecialAccess(orgUnitId, quizId, userId int64) (*QuizSpecialAccessData, error) {
	var out QuizSpecialAccessData
	err := c.get(c.lePath("%d/quizzes/%d/specialaccess/%d", orgUnitId, quizId, userId), nil, &out)
	return &out, err
}
