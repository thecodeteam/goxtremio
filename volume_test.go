package goxtremio

import (
	"fmt"
	"testing"
)

func TestGetVolumeByID(*testing.T) {
	volume, err := c.GetVolume("24", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))
}

func TestGetVolumeByName(*testing.T) {
	volume, err := c.GetVolume("", "ubuntu-vol4")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))
}

func TestGetVolumes(*testing.T) {
	volumes, err := c.GetVolumes()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumes))
}

func TestNewVolume(*testing.T) {
	req := &NewVolumeOptions{
		VolName: "testing1",
		VolSize: 1073741824,
	}
	res, err := c.NewVolume(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", res))
}

func TestDeleteVolume(*testing.T) {
	err := c.DeleteVolume("", "testing1")
	if err != nil {
		panic(err)
	}
}
