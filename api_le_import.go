package valence

import (
	"io"
	"net/url"
	"path/filepath"
)

// CreateCourseCopyJob queues a Brightspace course copy job.
// POST /d2l/api/le/{leVersion}/import/{orgUnitId}/copy/
func (c *Client) CreateCourseCopyJob(orgUnitId int64, data CreateCopyJobRequest) (*CreateCopyJobResponse, error) {
	var out CreateCopyJobResponse
	err := c.postJSON(c.lePath("import/%d/copy/", orgUnitId), nil, data, &out)
	return &out, err
}

// GetCourseCopyJobs returns the copy jobs for a target course offering.
// GET /d2l/api/le/{leVersion}/import/{orgUnitId}/copy/
func (c *Client) GetCourseCopyJobs(orgUnitId int64) ([]CourseImportJobData, error) {
	var out []CourseImportJobData
	err := c.get(c.lePath("import/%d/copy/", orgUnitId), nil, &out)
	return out, err
}

// GetCourseCopyJob returns the status of a queued course copy job request.
// GET /d2l/api/le/{leVersion}/import/{orgUnitId}/copy/{jobToken}
func (c *Client) GetCourseCopyJob(orgUnitId int64, jobToken string) (*GetCopyJobResponse, error) {
	var out GetCopyJobResponse
	err := c.get(c.lePath("import/%d/copy/%s", orgUnitId, jobToken), nil, &out)
	return &out, err
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
