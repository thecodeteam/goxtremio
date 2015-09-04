package goxtremio

import (
	"fmt"
	"testing"
)

func TestGetEventsBySeverity(*testing.T) {
	event, err := c.GetEvents("information", "", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", event))
}

func TestGetEventsByEventCode(*testing.T) {
	event, err := c.GetEvents("", "1400201", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", event))
}

func TestGetEventsByDescription(*testing.T) {
	event, err := c.GetEvents("", "", "^InfiniBand Switch port state for port")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", event))
}

func TestGetEvents(*testing.T) {
	events, err := c.GetEvents("", "", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", events))
}
