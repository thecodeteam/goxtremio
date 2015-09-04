package goxtremio

import xms "github.com/emccode/goxtremio/api/v3"

type ISCSIPortal *xms.ISCSIPortal

//GetISCSIPortal returns a specific iSCSI portal by name or ID
func (c *Client) GetISCSIPortal(id string, name string) (ISCSIPortal, error) {
	iSCSIPortal, err := c.api.GetISCSIPortal(id, name)
	if err != nil {
		return nil, err
	}
	return iSCSIPortal.Content, nil
}

//GetISCSIPortals returns a list of iSCSI portals
func (c *Client) GetISCSIPortals() (Refs, error) {
	iSCSIPortals, err := c.api.GetISCSIPortals()
	if err != nil {
		return nil, err
	}
	return iSCSIPortals.ISCSIPortals, nil
}
