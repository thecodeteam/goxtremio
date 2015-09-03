package goxtremio

import (
	"errors"
	xms "github.com/emccode/goxtremio/api/v3"
)

type Snapshot *xms.Snapshot
type SnapListItems []*xms.SnapListItem
type NewSnapshotOptions xms.PostSnapshotsReq
type NewSnapshotResult *xms.PostSnapshotsResp

//GetSnapshot returns a specific snapshot by name or ID
func (c *Client) GetSnapshot(id string, name string) (Snapshot, error) {
	initiator, err := c.api.GetSnapshot(id, name)
	if err != nil {
		return nil, err
	}
	return initiator.Content, nil
}

//GetSnapshots returns a list of snapshots
func (c *Client) GetSnapshots() (Refs, error) {
	initiators, err := c.api.GetSnapshots()
	if err != nil {
		return nil, err
	}
	return initiators.Snapshots, nil
}

func NewSnapListItems(ancestorIdsAndNames ...string) SnapListItems {
	if ancestorIdsAndNames == nil {
		panic(errors.New("ancestorIdsAndNames is required"))
	}

	y := 0
	l := len(ancestorIdsAndNames)
	arr := make([]*xms.SnapListItem, l/2)

	for x := 0; x < l; x += 2 {
		arr[y] = &xms.SnapListItem{
			AncestorVolID: ancestorIdsAndNames[x],
			SnapVolName:   ancestorIdsAndNames[x+1],
		}
		y += 1
	}

	return arr
}

//Constructs a new Snapshot instance
func SnapshotCtor() Snapshot {
	return &xms.Snapshot{}
}

//Constructs a new Snapshot instance
func SnapshotCtorNameIndex(name string, index int) Snapshot {
	return &xms.Snapshot{Name: name, Index: index}
}

//NewSnapshot creates a volume
func (c *Client) NewSnapshot(opts *NewSnapshotOptions) (NewSnapshotResult, error) {
	req := xms.PostSnapshotsReq(*opts)
	return c.api.PostSnapshots(&req)
}

//DeleteSnapshot deletes a volume
func (c *Client) DeleteSnapshot(id string, name string) error {
	return c.api.DeleteSnapshots(id, name)
}
