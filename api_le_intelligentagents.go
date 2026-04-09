package valence

import "net/url"

// GetIntelligentAgents returns all intelligent agents for an org unit.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/agents
func (c *Client) GetIntelligentAgents(orgUnitId int64, params url.Values) (*ObjectListPage[IntelligentAgent], error) {
	var out ObjectListPage[IntelligentAgent]
	err := c.get(c.lePath("%d/agents", orgUnitId), params, &out)
	return &out, err
}

// GetIntelligentAgent returns a specific intelligent agent.
// GET /d2l/api/le/{leVersion}/{orgUnitId}/agents/{agentId}
func (c *Client) GetIntelligentAgent(orgUnitId, agentId int64) (*IntelligentAgent, error) {
	var out IntelligentAgent
	err := c.get(c.lePath("%d/agents/%d", orgUnitId, agentId), nil, &out)
	return &out, err
}
