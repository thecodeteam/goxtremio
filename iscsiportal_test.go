package goxtremio

import (
	"fmt"
	"testing"
)

func TestGetISCSIPortalByID(*testing.T) {
	iSCSIPortal, err := GetISCSIPortal("4", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", iSCSIPortal))
}

func TestGetISCSIPortalByName(*testing.T) {
	iSCSIPortal, err := GetISCSIPortal("", "192.168.1.64/24")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", iSCSIPortal))
}

func TestGetISCSIPortals(*testing.T) {
	iSCSIPortals, err := GetISCSIPortals()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", iSCSIPortals))
}
