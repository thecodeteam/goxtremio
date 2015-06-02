package main

import xmsv3 "github.com/emccode/goxtremio/api/v3"

//GetSnapshot returns a specific snapshot by name or ID
func GetSnapshot(id string, name string) (*xmsv3.Snapshot, error) {
	initiator, err := xms.GetSnapshot(id, name)
	if err != nil {
		return nil, err
	}
	return initiator.Content, nil
}

//GetSnapshots returns a list of snapshots
func GetSnapshots() ([]*xmsv3.Ref, error) {
	initiators, err := xms.GetSnapshots()
	if err != nil {
		return nil, err
	}
	return initiators.Snapshots, nil
}

//NewSnapshot creates a volume
func NewSnapshot(opts *xmsv3.PostSnapshotsReq) (resp *xmsv3.PostSnapshotsResp, err error) {
	return xms.PostSnapshots(opts)
}

//DeleteSnapshot deletes a volume
func DeleteSnapshot(id string, name string) error {
	return xms.DeleteSnapshots(id, name)
}
