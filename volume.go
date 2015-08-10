package goxtremio

import xmsv3 "github.com/emccode/goxtremio/api/v3"

//GetVolume returns a specific volume by name or ID
func GetVolume(id string, name string) (*xmsv3.Volume, error) {
	volume, err := xms.GetVolume(id, name)
	if err != nil {
		return nil, err
	}
	return volume.Content, nil
}

//GetVolumes returns a list of volumes
func GetVolumes() ([]*xmsv3.Ref, error) {
	volumes, err := xms.GetVolumes()
	if err != nil {
		return nil, err
	}
	return volumes.Volumes, nil
}

//NewVolume creates a volume
func NewVolume(opts *xmsv3.PostVolumesReq) (resp *xmsv3.PostVolumesResp, err error) {
	return xms.PostVolumes(opts)
}

//DeleteVolume deletes a volume
func DeleteVolume(id string, name string) error {
	return xms.DeleteVolumes(id, name)
}
