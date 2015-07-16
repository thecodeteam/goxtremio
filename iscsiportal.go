package goxtremio

import xmsv3 "github.com/emccode/goxtremio/api/v3"

//GetISCSIPortal returns a specific iSCSI portal by name or ID
func GetISCSIPortal(id string, name string) (*xmsv3.ISCSIPortal, error) {
	iSCSIPortal, err := xms.GetISCSIPortal(id, name)
	if err != nil {
		return nil, err
	}
	return iSCSIPortal.Content, nil
}

//GetISCSIPortals returns a list of iSCSI portals
func GetISCSIPortals() ([]*xmsv3.Ref, error) {
	iSCSIPortals, err := xms.GetISCSIPortals()
	if err != nil {
		return nil, err
	}
	return iSCSIPortals.ISCSIPortals, nil
}
