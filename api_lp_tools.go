package valence

import "net/url"

// GetOrgUnitToolsPage returns one page of tool information for an org unit.
// Supported params include "bookmark" and "namesOnly". If namesOnly is true,
// use GetOrgUnitToolNamesPage instead.
// GET /d2l/api/lp/{lpVersion}/tools/orgUnits/{orgUnitId}
func (c *Client) GetOrgUnitToolsPage(orgUnitId int64, params url.Values) (*PagedResultSet[ToolInfo], error) {
	var out PagedResultSet[ToolInfo]
	err := c.get(c.lpPath("tools/orgUnits/%d", orgUnitId), params, &out)
	return &out, err
}

// GetOrgUnitToolNamesPage returns one page of tool IDs and names for an org unit.
// This uses the namesOnly query parameter on the org-unit tools route.
// GET /d2l/api/lp/{lpVersion}/tools/orgUnits/{orgUnitId}?namesOnly=true
func (c *Client) GetOrgUnitToolNamesPage(orgUnitId int64, params url.Values) (*PagedResultSet[ToolWithName], error) {
	if params == nil {
		params = url.Values{}
	}
	params.Set("namesOnly", "true")

	var out PagedResultSet[ToolWithName]
	err := c.get(c.lpPath("tools/orgUnits/%d", orgUnitId), params, &out)
	return &out, err
}

// GetOrgUnitTools returns all tools configured for an org unit, paging through
// all result pages automatically.
// GET /d2l/api/lp/{lpVersion}/tools/orgUnits/{orgUnitId}
func (c *Client) GetOrgUnitTools(orgUnitId int64) ([]ToolInfo, error) {
	var all []ToolInfo
	bookmark := ""
	for {
		params := url.Values{}
		if bookmark != "" {
			params.Set("bookmark", bookmark)
		}
		page, err := c.GetOrgUnitToolsPage(orgUnitId, params)
		if err != nil {
			return nil, err
		}
		all = append(all, page.Items...)
		if !page.PagingInfo.HasMoreItems {
			break
		}
		bookmark = page.PagingInfo.Bookmark
	}
	return all, nil
}

// GetOrgUnitToolNames returns all tool IDs and display names for an org unit,
// paging through all result pages automatically.
// GET /d2l/api/lp/{lpVersion}/tools/orgUnits/{orgUnitId}?namesOnly=true
func (c *Client) GetOrgUnitToolNames(orgUnitId int64) ([]ToolWithName, error) {
	var all []ToolWithName
	bookmark := ""
	for {
		params := url.Values{}
		if bookmark != "" {
			params.Set("bookmark", bookmark)
		}
		page, err := c.GetOrgUnitToolNamesPage(orgUnitId, params)
		if err != nil {
			return nil, err
		}
		all = append(all, page.Items...)
		if !page.PagingInfo.HasMoreItems {
			break
		}
		bookmark = page.PagingInfo.Bookmark
	}
	return all, nil
}
