package goxtremio

import (
	"fmt"
	"testing"
)

func TestGetLunMapByID(*testing.T) {
	volume, err := c.GetLunMap("21", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))
}

func TestGetLunMapByName(*testing.T) {
	volume, err := c.GetLunMap("", "18_1_1")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))
}

func TestGetLunMaps(*testing.T) {
	volumes, err := c.GetLunMaps()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumes))
}

func TestNewLunMaps(*testing.T) {
	req := &NewLunMapOptions{
		VolID: 24,
		IgID:  4,
	}
	res, err := c.NewLunMap(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", res))
}

func TestDeleteLunMapsByID(*testing.T) {
	err := c.DeleteLunMap("21", "")
	if err != nil {
		panic(err)
	}

}

func TestDeleteLunMapsByName(*testing.T) {
	err := c.DeleteLunMap("", "testing1")
	if err != nil {
		panic(err)
	}
}
