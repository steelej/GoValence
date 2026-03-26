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
