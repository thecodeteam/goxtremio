package goxtremio

import xms "github.com/emccode/goxtremio/api/v3"

type Volume *xms.Volume
type NewVolumeOptions xms.PostVolumesReq
type NewVolumeResult *xms.PostVolumesResp

//GetVolume returns a specific volume by name or ID
func (c *Client) GetVolume(id string, name string) (Volume, error) {
	volume, err := c.api.GetVolume(id, name)
	if err != nil {
		return nil, err
	}
	return volume.Content, nil
}

//GetVolumes returns a list of volumes
func (c *Client) GetVolumes() (Refs, error) {
	volumes, err := c.api.GetVolumes()
	if err != nil {
		return nil, err
	}
	return volumes.Volumes, nil
}

//Constructs a new Volume instance
func VolumeCtor() Volume {
	return &xms.Volume{}
}

//Constructs a new Volume instance
func VolumeCtorNameIndex(name string, index int) Volume {
	return &xms.Volume{Name: name, Index: index}
}

//NewVolume creates a volume
func (c *Client) NewVolume(opts *NewVolumeOptions) (NewVolumeResult, error) {
	req := xms.PostVolumesReq(*opts)
	return c.api.PostVolumes(&req)
}

//DeleteVolume deletes a volume
func (c *Client) DeleteVolume(id string, name string) error {
	return c.api.DeleteVolumes(id, name)
}
