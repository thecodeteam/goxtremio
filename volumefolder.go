package goxtremio

import xms "github.com/emccode/goxtremio/api/v3"

type VolumeFolder *xms.VolumeFolder

//GetVolumeFolder returns a specific initiator by name or ID
func (c *Client) GetVolumeFolder(id string, name string) (VolumeFolder, error) {
	volumeFolder, err := c.api.GetVolumeFolder(id, name)
	if err != nil {
		return nil, err
	}
	return volumeFolder.Content, nil
}

//GetVolumeFolders returns a list of initiators
func (c *Client) GetVolumeFolders() (Refs, error) {
	volumeFolders, err := c.api.GetVolumeFolders()
	if err != nil {
		return nil, err
	}
	return volumeFolders.Folders, nil
}
