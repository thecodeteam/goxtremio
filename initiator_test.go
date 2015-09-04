package goxtremio

import (
	"fmt"
	"testing"
)

func TestGetInitiatorByID(*testing.T) {
	initiator, err := c.GetInitiator("21", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiator))
}

func TestGetInitiatorByName(*testing.T) {
	initiator, err := c.GetInitiator("", "VPLEX-ee20")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiator))
}

func TestGetInitiators(*testing.T) {
	initiators, err := c.GetInitiators()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiators))
}

func TestNewInitiator(*testing.T) {
	req := &NewInitiatorOptions{
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
	res, err := c.NewInitiator(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", res))
}

func TestDeleteInitiator(*testing.T) {
	err := c.DeleteInitiator("", "ubuntu1")
	if err != nil {
		panic(err)
	}
}
