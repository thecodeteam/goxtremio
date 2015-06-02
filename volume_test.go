package main

import (
	"fmt"
	"testing"

	xmsv3 "github.com/emccode/goxtremio/api/v3"
)

func TestGetVolumeByID(*testing.T) {
	volume, err := GetVolume("24", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))
}

func TestGetVolumeByName(*testing.T) {
	volume, err := GetVolume("", "ubuntu-vol4")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))
}

func TestGetVolumes(*testing.T) {
	volumes, err := GetVolumes()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumes))
}

func TestNewVolume(*testing.T) {
	req := &xmsv3.PostVolumesReq{
		VolName: "testing1",
		VolSize: 1073741824,
	}
	postVolumesResp, err := NewVolume(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", postVolumesResp))
}

func TestDeleteVolume(*testing.T) {
	err := DeleteVolume("", "testing1")
	if err != nil {
		panic(err)
	}
}
