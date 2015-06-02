package main

import xmsv3 "github.com/emccode/goxtremio/api/v3"

//GetLunMap returns a specific lunMap by name or ID
func GetLunMap(id string, name string) (*xmsv3.LunMap, error) {
	lunMap, err := xms.GetLunMap(id, name)
	if err != nil {
		return nil, err
	}
	return lunMap.Content, nil
}

//GetLunMaps returns a list of lunMaps
func GetLunMaps() ([]*xmsv3.Ref, error) {
	lunMaps, err := xms.GetLunMaps()
	if err != nil {
		return nil, err
	}
	return lunMaps.LunMaps, nil
}

//NewLunMap creates a volume
func NewLunMap(opts *xmsv3.PostLunMapsReq) (resp *xmsv3.PostLunMapsResp, err error) {
	return xms.PostLunMaps(opts)
}

//DeleteLunMap deletes a volume
func DeleteLunMap(id string, name string) error {
	return xms.DeleteLunMaps(id, name)
}
