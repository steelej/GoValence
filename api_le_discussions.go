package valence

import "net/url"

// GetForums returns all discussion forums for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/discussions/forums/
func (c *Client) GetForums(orgUnitId int64) ([]Forum, error) {
	var out []Forum
	err := c.get(c.lePath("%d/discussions/forums/", orgUnitId), nil, &out)
	return out, err
}

// GetForum returns a specific discussion forum.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/discussions/forums/{forumId}
func (c *Client) GetForum(orgUnitId, forumId int64) (*Forum, error) {
	var out Forum
	err := c.get(c.lePath("%d/discussions/forums/%d", orgUnitId, forumId), nil, &out)
	return &out, err
}

// GetTopics returns all topics in a discussion forum.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/discussions/forums/{forumId}/topics/
func (c *Client) GetTopics(orgUnitId, forumId int64) ([]Topic, error) {
	var out []Topic
	err := c.get(c.lePath("%d/discussions/forums/%d/topics/", orgUnitId, forumId), nil, &out)
	return out, err
}

// GetTopic returns a specific discussion topic.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/discussions/forums/{forumId}/topics/{topicId}
func (c *Client) GetTopic(orgUnitId, forumId, topicId int64) (*Topic, error) {
	var out Topic
	err := c.get(c.lePath("%d/discussions/forums/%d/topics/%d", orgUnitId, forumId, topicId), nil, &out)
	return &out, err
}

// GetPosts returns all posts in a discussion topic.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/discussions/forums/{forumId}/topics/{topicId}/posts/
func (c *Client) GetPosts(orgUnitId, forumId, topicId int64, params url.Values) (*PagedResultSet[Post], error) {
	var out PagedResultSet[Post]
	err := c.get(c.lePath("%d/discussions/forums/%d/topics/%d/posts/", orgUnitId, forumId, topicId), params, &out)
	return &out, err
}

// GetPost returns a specific discussion post.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/discussions/forums/{forumId}/topics/{topicId}/posts/{postId}
func (c *Client) GetPost(orgUnitId, forumId, topicId, postId int64) (*Post, error) {
	var out Post
	err := c.get(c.lePath("%d/discussions/forums/%d/topics/%d/posts/%d", orgUnitId, forumId, topicId, postId), nil, &out)
	return &out, err
}

// GetPostReplies returns all replies to a specific post.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/discussions/forums/{forumId}/topics/{topicId}/posts/{postId}/Replies
func (c *Client) GetPostReplies(orgUnitId, forumId, topicId, postId int64, params url.Values) (*PagedResultSet[Post], error) {
	var out PagedResultSet[Post]
	err := c.get(c.lePath("%d/discussions/forums/%d/topics/%d/posts/%d/Replies", orgUnitId, forumId, topicId, postId), params, &out)
	return &out, err
}
