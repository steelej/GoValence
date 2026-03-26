package valence

import (
	"io"
	"net/url"
	"strings"
)

// GetDropboxFolders returns all dropbox folders for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/dropbox/folders/
func (c *Client) GetDropboxFolders(orgUnitId int64, params url.Values) ([]DropboxFolder, error) {
	var out []DropboxFolder
	err := c.get(c.lePath("%d/dropbox/folders/", orgUnitId), params, &out)
	return out, err
}

// GetDropboxFolder returns a specific dropbox folder.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/dropbox/folders/{folderId}
func (c *Client) GetDropboxFolder(orgUnitId, folderId int64) (*DropboxFolder, error) {
	var out DropboxFolder
	err := c.get(c.lePath("%d/dropbox/folders/%d", orgUnitId, folderId), nil, &out)
	return &out, err
}

// GetDropboxSubmissions returns all submissions for a dropbox folder, grouped by user.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/dropbox/folders/{folderId}/submissions/
func (c *Client) GetDropboxSubmissions(orgUnitId, folderId int64) ([]UserSubmissions, error) {
	var out []UserSubmissions
	err := c.get(c.lePath("%d/dropbox/folders/%d/submissions/", orgUnitId, folderId), nil, &out)
	return out, err
}

// GetDropboxUserSubmissions returns submissions for a specific user in a dropbox folder.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/dropbox/folders/{folderId}/submissions/user/{userId}
func (c *Client) GetDropboxUserSubmissions(orgUnitId, folderId, userId int64) ([]DropboxSubmissionEntry, error) {
	var out []DropboxSubmissionEntry
	err := c.get(c.lePath("%d/dropbox/folders/%d/submissions/user/%d", orgUnitId, folderId, userId), nil, &out)
	return out, err
}

// GetDropboxSubmissionFile returns the raw bytes of a submission file.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/dropbox/folders/{folderId}/submissions/{submissionId}/files/{fileId}
func (c *Client) GetDropboxSubmissionFile(orgUnitId, folderId, submissionId, fileId int64) (io.ReadCloser, error) {
	return c.getRaw(c.lePath("%d/dropbox/folders/%d/submissions/%d/files/%d", orgUnitId, folderId, submissionId, fileId), nil)
}

// GetDropboxFeedbackFile returns the raw bytes of a feedback attachment file.
// entityType should be "user" or "group" (case-insensitive; will be lowercased).
// GET /d2l/api/le/{leVersion}/{orgUnitId}/dropbox/folders/{folderId}/feedback/{entityType}/{entityId}/attachments/{fileId}
func (c *Client) GetDropboxFeedbackFile(orgUnitId, folderId int64, entityType string, entityId, fileId int64) (io.ReadCloser, error) {
	return c.getRaw(c.lePath("%d/dropbox/folders/%d/feedback/%s/%d/attachments/%d", orgUnitId, folderId, strings.ToLower(entityType), entityId, fileId), nil)
}

// GetDropboxCategories returns all dropbox categories for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/dropbox/categories/
func (c *Client) GetDropboxCategories(orgUnitId int64) ([]DropboxCategory, error) {
	var out []DropboxCategory
	err := c.get(c.lePath("%d/dropbox/categories/", orgUnitId), nil, &out)
	return out, err
}
