package goxtremio

import xmsv3 "github.com/emccode/goxtremio/api/v3"

//GetInitiator returns a specific initiator by name or ID
func GetInitiator(id string, name string) (*xmsv3.Initiator, error) {
	initiator, err := xms.GetInitiator(id, name)
	if err != nil {
		return nil, err
	}
	return initiator.Content, nil
}

//GetInitiators returns a list of initiators
func GetInitiators() ([]*xmsv3.Ref, error) {
	initiators, err := xms.GetInitiators()
	if err != nil {
		return nil, err
	}
	return initiators.Initiators, nil
}

//NewInitiator creates a volume
func NewInitiator(opts *xmsv3.PostInitiatorsReq) (resp *xmsv3.PostInitiatorsResp, err error) {
	return xms.PostInitiators(opts)
}

//DeleteInitiator deletes a volume
func DeleteInitiator(id string, name string) error {
	return xms.DeleteInitiators(id, name)
}
