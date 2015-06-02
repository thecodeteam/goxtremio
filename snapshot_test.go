package main

import (
	"fmt"
	"testing"

	xmsv3 "github.com/emccode/goxtremio/api/v3"
)

func TestGetSnapshotByID(*testing.T) {
	snapshot, err := GetSnapshot("25", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", snapshot))
}

func TestGetSnapshotByName(*testing.T) {
	snapshot, err := GetSnapshot("", "ubuntu-vol4.snap.06022015-09:58")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", snapshot))
}

func TestGetSnapshots(*testing.T) {
	snapshots, err := GetSnapshots()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", snapshots))
}

func TestNewSnapshot(*testing.T) {
	var snapList []*xmsv3.SnapListItem
	snapList = append(snapList, &xmsv3.SnapListItem{
		AncestorVolID: "ubuntu-vol4",
		SnapVolName:   "ubuntu-vol4-snap1",
	})

	req := &xmsv3.PostSnapshotsReq{
		SnapList: snapList,
		FolderID: "/Ubuntu",
	}

	postSnapshotsResp, err := NewSnapshot(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", postSnapshotsResp))
}

func TestDeleteSnapshot(*testing.T) {
	err := DeleteSnapshot("", "ubuntu-vol4-snap1")
	if err != nil {
		panic(err)
	}
}
