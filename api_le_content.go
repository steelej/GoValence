package valence

import (
	"io"
	"net/url"
)

// GetTableOfContents returns the full content table of contents for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/content/toc
func (c *Client) GetTableOfContents(orgUnitId int64, params url.Values) (*TableOfContents, error) {
	var out TableOfContents
	err := c.get(c.lePath("%d/content/toc", orgUnitId), params, &out)
	return &out, err
}

// GetContentTopicFile returns the raw bytes of a content file topic.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/content/topics/{topicId}/file
func (c *Client) GetContentTopicFile(orgUnitId, topicId int64) (io.ReadCloser, error) {
	return c.getRaw(c.lePath("%d/content/topics/%d/file", orgUnitId, topicId), nil)
}
