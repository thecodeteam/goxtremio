package main

import (
	"fmt"
	"testing"

	xmsv3 "github.com/emccode/goxtremio/api/v3"
)

func TestGetInitiatorByID(*testing.T) {
	initiator, err := GetInitiator("21", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiator))
}

func TestGetInitiatorByName(*testing.T) {
	initiator, err := GetInitiator("", "VPLEX-ee20")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiator))
}

func TestGetInitiators(*testing.T) {
	initiators, err := GetInitiators()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiators))
}

func TestNewInitiator(*testing.T) {
	req := &xmsv3.PostInitiatorsReq{
		IgID: 4,
		// IgName
		InitiatorName: "iqn.1993-08.org.debian:01:dca5bccb64",
		PortAddress:   "iqn.1993-08.org.debian:01:dca5bccb64",
		// InitiatorAuthenticationPassword
		// InitiatorAuthenticationUserName
		// InitiatorDiscoveryPassword
		// InitiatorDiscoveryUserName
		// InitiatorChapDiscoveryCredentials
	}
	postInitiatorsResp, err := NewInitiator(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", postInitiatorsResp))
}

func TestDeleteInitiator(*testing.T) {
	err := DeleteInitiator("", "ubuntu1")
	if err != nil {
		panic(err)
	}
}
