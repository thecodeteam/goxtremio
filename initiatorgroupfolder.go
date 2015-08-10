package goxtremio

import xmsv3 "github.com/emccode/goxtremio/api/v3"

//GetIGFolder returns a specific initiator by name or ID
func GetIGFolder(id string, name string) (*xmsv3.IGFolder, error) {
	igFolder, err := xms.GetIGFolder(id, name)
	if err != nil {
		return nil, err
	}
	return igFolder.Content, nil
}

//GetIGFolders returns a list of initiators
func GetIGFolders() ([]*xmsv3.Ref, error) {
	igFolders, err := xms.GetIGFolders()
	if err != nil {
		return nil, err
	}
	return igFolders.Folders, nil
}
