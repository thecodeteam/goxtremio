package goxtremio

import xmsv3 "github.com/emccode/goxtremio/api/v3"

//GetInitiatorGroup returns a specific initiator by name or ID
func GetInitiatorGroup(id string, name string) (*xmsv3.InitiatorGroup, error) {
	initiator, err := xms.GetInitiatorGroup(id, name)
	if err != nil {
		return nil, err
	}
	return initiator.Content, nil
}

//GetInitiatorGroups returns a list of initiators
func GetInitiatorGroups() ([]*xmsv3.Ref, error) {
	initiators, err := xms.GetInitiatorGroups()
	if err != nil {
		return nil, err
	}
	return initiators.InitiatorGroups, nil
}
