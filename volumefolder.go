package main

import xmsv3 "github.com/emccode/goxtremio/api/v3"

//GetVolumeFolder returns a specific initiator by name or ID
func GetVolumeFolder(id string, name string) (*xmsv3.VolumeFolder, error) {
	volumeFolder, err := xms.GetVolumeFolder(id, name)
	if err != nil {
		return nil, err
	}
	return volumeFolder.Content, nil
}

//GetVolumeFolders returns a list of initiators
func GetVolumeFolders() ([]*xmsv3.Ref, error) {
	volumeFolders, err := xms.GetVolumeFolders()
	if err != nil {
		return nil, err
	}
	return volumeFolders.Folders, nil
}
