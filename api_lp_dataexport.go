package valence

import "io"

const bdsAdvancedDataSetUUID = "49ac9b6f-8cbc-4a98-a95c-6ce0d89bca57"

// GetBDSDownload downloads the Brightspace Data Set (BDS) for the advanced data set.
// GET /d2l/api/lp/{lpVersion}/dataExport/bds/download/49ac9b6f-8cbc-4a98-a95c-6ce0d89bca57
func (c *Client) GetBDSDownload() (io.ReadCloser, error) {
	return c.getRaw(c.lpPath("dataExport/bds/download/%s", bdsAdvancedDataSetUUID), nil)
}
