package goxtremio

import xms "github.com/emccode/goxtremio/api/v3"

type InitiatorGroup xms.InitiatorGroup

//GetInitiatorGroup returns a specific initiator by name or ID
func (c *Client) GetInitiatorGroup(
	id string, name string) (*InitiatorGroup, error) {
	ig, err := c.api.GetInitiatorGroup(id, name)
	if err != nil {
		return nil, err
	}
	igg := InitiatorGroup(*ig.Content)
	return &igg, nil
}

//GetInitiatorGroups returns a list of initiators
func (c *Client) GetInitiatorGroups() (Refs, error) {
	ig, err := c.api.GetInitiatorGroups()
	if err != nil {
		return nil, err
	}
	return ig.InitiatorGroups, nil
}
