package valence

import (
	"io"
	"net/url"
	"path/filepath"
)

// GetCourseCopyJobStatus returns the status of a course copy job.
// The jobToken is returned when a copy job is initiated (POST).
// This GET variant polls the job status.
// GET /d2l/api/le/{leVersion}/import/{orgUnitId}/copy/
func (c *Client) GetCourseCopyJobs(orgUnitId int64) ([]CourseImportJobData, error) {
	var out []CourseImportJobData
	err := c.get(c.lePath("import/%d/copy/", orgUnitId), nil, &out)
	return out, err
}

// CreateCourseImportJob uploads a course package and creates a new import job.
// POST /d2l/api/le/{leVersion}/import/{orgUnitId}/imports/
func (c *Client) CreateCourseImportJob(orgUnitId int64, fileName string, body io.Reader, callbackURL string) (*CourseImportJobData, error) {
	var out CourseImportJobData
	params := url.Values{}
	if callbackURL != "" {
		params.Set("callbackUrl", callbackURL)
	}

	err := c.postMultipartFile(
		c.lePath("import/%d/imports/", orgUnitId),
		params,
		"file",
		filepath.Base(fileName),
		"application/zip",
		body,
		&out,
	)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

// GetCourseImportJob returns the current status of an import job.
// GET /d2l/api/le/{leVersion}/import/{orgUnitId}/imports/{jobToken}
func (c *Client) GetCourseImportJob(orgUnitId int64, jobToken string) (*CourseImportJobStatus, error) {
	var out CourseImportJobStatus
	err := c.get(c.lePath("import/%d/imports/%s", orgUnitId, jobToken), nil, &out)
	return &out, err
}
