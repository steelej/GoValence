package valence

import "net/url"

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
		var page PagedResultSet[ToolInfo]
		if err := c.get(c.lpPath("tools/orgUnits/%d", orgUnitId), params, &page); err != nil {
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
