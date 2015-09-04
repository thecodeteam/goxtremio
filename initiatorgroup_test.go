package goxtremio

import (
	"fmt"
	"testing"
)

func TestGetInitiatorGroupByID(*testing.T) {
	initiator, err := c.GetInitiatorGroup("4", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiator))
}

func TestGetInitiatorGroupByName(*testing.T) {
	initiator, err := c.GetInitiatorGroup("", "VPLEX-ee20")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiator))
}

func TestGetInitiatorGroups(*testing.T) {
	initiators, err := c.GetInitiatorGroups()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiators))
}
