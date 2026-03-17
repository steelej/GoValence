package valence

// GetCourseCopyJobStatus returns the status of a course copy job.
// The jobToken is returned when a copy job is initiated (POST).
// This GET variant polls the job status.
// GET /d2l/api/le/{leVersion}/import/{orgUnitId}/copy/
func (c *Client) GetCourseCopyJobs(orgUnitId int64) ([]CourseImportJobData, error) {
	var out []CourseImportJobData
	err := c.get(c.lePath("import/%d/copy/", orgUnitId), nil, &out)
	return out, err
}
