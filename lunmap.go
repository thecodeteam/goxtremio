package goxtremio

import xms "github.com/emccode/goxtremio/api/v3"

type LunMap *xms.LunMap
type NewLunMapOptions xms.PostLunMapsReq
type NewLunMapResult *xms.PostLunMapsResp

//GetLunMap returns a specific lunMap by name or ID
func (c *Client) GetLunMap(id string, name string) (LunMap, error) {
	lunMap, err := c.api.GetLunMap(id, name)
	if err != nil {
		return nil, err
	}
	return lunMap.Content, nil
}

//GetLunMaps returns a list of lunMaps
func (c *Client) GetLunMaps() (Refs, error) {
	lunMaps, err := c.api.GetLunMaps()
	if err != nil {
		return nil, err
	}
	return lunMaps.LunMaps, nil
}

//NewLunMap creates a volume
func (c *Client) NewLunMap(opts *NewLunMapOptions) (NewLunMapResult, error) {
	req := xms.PostLunMapsReq(*opts)
	return c.api.PostLunMaps(&req)
}

//DeleteLunMap deletes a volume
func (c *Client) DeleteLunMap(id string, name string) error {
	return c.api.DeleteLunMaps(id, name)
}
