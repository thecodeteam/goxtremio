package goxtremio

import xms "github.com/emccode/goxtremio/api/v3"

type InitiatorGroupFolder *xms.IGFolder

//GetIGFolder returns a specific initiator by name or ID
func (c *Client) GetIGFolder(id string, name string) (InitiatorGroupFolder, error) {
	igf, err := c.api.GetIGFolder(id, name)
	if err != nil {
		return nil, err
	}
	return InitiatorGroupFolder(igf.Content), nil
}

//GetIGFolders returns a list of initiators
func (c *Client) GetIGFolders() (Refs, error) {
	igFolders, err := c.api.GetIGFolders()
	if err != nil {
		return nil, err
	}
	return igFolders.Folders, nil
}
