package valence

// GetLTILinks returns all LTI links for an org unit.
// GET /d2l/api/le/{leVersion}/lti/link/{orgUnitId}/
func (c *Client) GetLTILinks(orgUnitId int64) ([]LTILink, error) {
	var out []LTILink
	err := c.get(c.lePath("lti/link/%d/", orgUnitId), nil, &out)
	return out, err
}

// GetLTILinkSharing returns sharing info for an LTI link.
// GET /d2l/api/le/{leVersion}/lti/link/{orgUnitId}/{linkId}/sharing/
func (c *Client) GetLTILinkSharing(orgUnitId, linkId int64) ([]LTISharingData, error) {
	var out []LTISharingData
	err := c.get(c.lePath("lti/link/%d/%d/sharing/", orgUnitId, linkId), nil, &out)
	return out, err
}

// GetLTIToolProviderSharing returns sharing info for an LTI tool provider.
// GET /d2l/api/le/{leVersion}/lti/tp/{sourceOrgUnitId}/{tpId}/sharing/
func (c *Client) GetLTIToolProviderSharing(sourceOrgUnitId, tpId int64) ([]LTISharingData, error) {
	var out []LTISharingData
	err := c.get(c.lePath("lti/tp/%d/%d/sharing/", sourceOrgUnitId, tpId), nil, &out)
	return out, err
}

// GetLTIAdvantageLinks returns all LTI Advantage links for an org unit.
// GET /d2l/api/le/{leVersion}/ltiadvantage/links/orgunit/{orgUnitId}/
func (c *Client) GetLTIAdvantageLinks(orgUnitId int64) ([]LTIAdvantageLink, error) {
	var out []LTIAdvantageLink
	err := c.get(c.lePath("ltiadvantage/links/orgunit/%d/", orgUnitId), nil, &out)
	return out, err
}

// GetLTIAdvantageQuickLink returns a specific LTI Advantage quick link.
// GET /d2l/api/le/{leVersion}/ltiadvantage/quicklinks/orgunit/{orgUnitId}/link/{linkId}
func (c *Client) GetLTIAdvantageQuickLink(orgUnitId, linkId int64) (*LTIAdvantageLink, error) {
	var out LTIAdvantageLink
	err := c.get(c.lePath("ltiadvantage/quicklinks/orgunit/%d/link/%d", orgUnitId, linkId), nil, &out)
	return &out, err
}

// GetLTIAdvantageDeploymentSharing returns sharing info for an LTI Advantage deployment.
// GET /d2l/api/le/{leVersion}/ltiadvantage/deployment/{deploymentId}/sharing/
func (c *Client) GetLTIAdvantageDeploymentSharing(deploymentId int64) ([]LTIDeploymentSharingData, error) {
	var out []LTIDeploymentSharingData
	err := c.get(c.lePath("ltiadvantage/deployment/%d/sharing/", deploymentId), nil, &out)
	return out, err
}

// GetLTIAdvantageDeploymentOrgUnitSharing returns sharing for a specific org unit on a deployment.
// GET /d2l/api/le/{leVersion}/ltiadvantage/deployment/{deploymentId}/sharing/{sharingOrgUnitId}
func (c *Client) GetLTIAdvantageDeploymentOrgUnitSharing(deploymentId, sharingOrgUnitId int64) (*LTIDeploymentSharingData, error) {
	var out LTIDeploymentSharingData
	err := c.get(c.lePath("ltiadvantage/deployment/%d/sharing/%d", deploymentId, sharingOrgUnitId), nil, &out)
	return &out, err
}
