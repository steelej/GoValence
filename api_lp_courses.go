package valence

import (
	"io"
	"net/url"
)

// GetCourse returns a specific course offering.
// GET /d2l/api/lp/{lpVersion}/courses/{orgUnitId}
func (c *Client) GetCourse(orgUnitId int64) (*CourseOffering, error) {
	var out CourseOffering
	err := c.get(c.lpPath("courses/%d", orgUnitId), nil, &out)
	return &out, err
}

// CreateCourse creates a new course offering.
// POST /d2l/api/lp/{lpVersion}/courses/
func (c *Client) CreateCourse(data CreateCourseOffering) (*CourseOffering, error) {
	var out CourseOffering
	err := c.postJSON(c.lpPath("courses/"), nil, data, &out)
	return &out, err
}

// UpdateCourse replaces the information for a course offering.
// PUT /d2l/api/lp/{lpVersion}/courses/{orgUnitId}
func (c *Client) UpdateCourse(orgUnitId int64, data CourseOfferingInfo) (*CourseOffering, error) {
	var out CourseOffering
	err := c.putJSON(c.lpPath("courses/%d", orgUnitId), nil, data, &out)
	return &out, err
}

// DeleteCourse deletes a course offering.
// DELETE /d2l/api/lp/{lpVersion}/courses/{orgUnitId}
func (c *Client) DeleteCourse(orgUnitId int64) error {
	return c.delete(c.lpPath("courses/%d", orgUnitId), nil)
}

// GetCourseImage returns the raw bytes of a course's image.
// GET /d2l/api/lp/{lpVersion}/courses/{orgUnitId}/image
func (c *Client) GetCourseImage(orgUnitId int64) (io.ReadCloser, error) {
	return c.getRaw(c.lpPath("courses/%d/image", orgUnitId), nil)
}

// GetCourseTemplates returns a paged list of course templates.
// GET /d2l/api/lp/{lpVersion}/coursetemplates/
func (c *Client) GetCourseTemplates(params url.Values) (*PagedResultSet[CourseTemplate], error) {
	var out PagedResultSet[CourseTemplate]
	err := c.get(c.lpPath("coursetemplates/"), params, &out)
	return &out, err
}

// GetCourseTemplate returns a specific course template.
// GET /d2l/api/lp/{lpVersion}/coursetemplates/{courseTemplateId}
func (c *Client) GetCourseTemplate(courseTemplateId int64) (*CourseTemplate, error) {
	var out CourseTemplate
	err := c.get(c.lpPath("coursetemplates/%d", courseTemplateId), nil, &out)
	return &out, err
}

// CreateCourseTemplate creates a new course template.
// POST /d2l/api/lp/{lpVersion}/coursetemplates/
func (c *Client) CreateCourseTemplate(data CreateCourseTemplate) (*CourseTemplate, error) {
	var out CourseTemplate
	err := c.postJSON(c.lpPath("coursetemplates/"), nil, data, &out)
	return &out, err
}

// UpdateCourseTemplate replaces the basic information for a course template.
// PUT /d2l/api/lp/{lpVersion}/coursetemplates/{courseTemplateId}
func (c *Client) UpdateCourseTemplate(courseTemplateId int64, data CourseTemplateInfo) (*CourseTemplate, error) {
	var out CourseTemplate
	err := c.putJSON(c.lpPath("coursetemplates/%d", courseTemplateId), nil, data, &out)
	return &out, err
}

// DeleteCourseTemplate deletes a course template.
// DELETE /d2l/api/lp/{lpVersion}/coursetemplates/{courseTemplateId}
func (c *Client) DeleteCourseTemplate(courseTemplateId int64) error {
	return c.delete(c.lpPath("coursetemplates/%d", courseTemplateId), nil)
}
