package main

import (
	"fmt"
	"testing"

	xmsv3 "github.com/emccode/goxtremio/api/v3"
)

func TestGetLunMapByID(*testing.T) {
	volume, err := GetLunMap("21", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))
}

func TestGetLunMapByName(*testing.T) {
	volume, err := GetLunMap("", "18_1_1")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))
}

func TestGetLunMaps(*testing.T) {
	volumes, err := GetLunMaps()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumes))
}

func TestNewLunMaps(*testing.T) {
	req := &xmsv3.PostLunMapsReq{
		VolID: 24,
		IgID:  4,
	}
	postLunMapsResp, err := NewLunMap(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", postLunMapsResp))
}

func TestDeleteLunMapsByID(*testing.T) {
	err := DeleteLunMap("21", "")
	if err != nil {
		panic(err)
	}

}

func TestDeleteLunMapsByName(*testing.T) {
	err := DeleteLunMap("", "testing1")
	if err != nil {
		panic(err)
	}
}
