package goxtremio

import (
	"fmt"
	"testing"
)

var xms *XMS

func init() {
	xms = New("https://10.5.132.140/api/json", true, "admin", "Xtrem10")
}

func TestGetVolumes(*testing.T) {

	volumes, err := xms.getVolumes()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumes))

}

func TestGetVolumeById(*testing.T) {

	volume, err := xms.getVolume("24", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))

}

func TestGetVolumeByName(*testing.T) {

	volume, err := xms.getVolume("", "ubuntu-vol4")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))

}

func TestGetLunMaps(*testing.T) {

	lunMaps, err := xms.getLunMaps()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", lunMaps))

}

func TestGetLunMapById(*testing.T) {

	lunMap, err := xms.getLunMap("1", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", lunMap))

}

func TestGetLunMapByName(*testing.T) {

	lunMap, err := xms.getLunMap("", "24_4_1")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", lunMap))

}

func TestGetInitiators(*testing.T) {

	initiators, err := xms.getInitiators()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiators))

}

func TestGetInitiatorById(*testing.T) {

	initiator, err := xms.getInitiator("11", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiator))

}

func TestGetInitiatorByName(*testing.T) {

	initiator, err := xms.getInitiator("", "VPLEX-ee30")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiator))

}
