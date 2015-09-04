package goxtremio

import (
	"fmt"
	"testing"
)

func TestGetSnapshotByID(*testing.T) {
	snapshot, err := c.GetSnapshot("25", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", snapshot))
}

func TestGetSnapshotByName(*testing.T) {
	snapshot, err := c.GetSnapshot("", "ubuntu-vol4.snap.06022015-09:58")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", snapshot))
}

func TestGetSnapshots(*testing.T) {
	snapshots, err := c.GetSnapshots()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", snapshots))
}

func TestNewSnapshot(*testing.T) {

	req := &NewSnapshotOptions{
		SnapList: NewSnapListItems("ubuntu-vol4", "ubuntu-vol4-snap1"),
		FolderID: "/Ubuntu",
	}

	res, err := c.NewSnapshot(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", res))
}

func TestDeleteSnapshot(*testing.T) {
	err := c.DeleteSnapshot("", "ubuntu-vol4-snap1")
	if err != nil {
		panic(err)
	}
}
