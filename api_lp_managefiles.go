package valence

import (
	"io"
	"net/url"
)

// GetManageFiles returns the direct child folders and files at an org unit path.
// Set the optional "path" query parameter to list a folder relative to the course path.
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/managefiles/
func (c *Client) GetManageFiles(orgUnitId int64, params url.Values) (*ObjectListPage[FileSystemObject], error) {
	var out ObjectListPage[FileSystemObject]
	err := c.get(c.lpPath("%d/managefiles/", orgUnitId), params, &out)
	return &out, err
}

// GetManageFile returns the raw bytes of a course file at path.
// The path must be relative to the course path, for example "/topic1/topic1.html".
// GET /d2l/api/lp/{lpVersion}/{orgUnitId}/managefiles/file
func (c *Client) GetManageFile(orgUnitId int64, path string) (io.ReadCloser, error) {
	params := url.Values{}
	params.Set("path", path)
	return c.getRaw(c.lpPath("%d/managefiles/file", orgUnitId), params)
}
