package goxtremio

import (
	"fmt"
	"testing"
)

func TestGetIGFolderByID(*testing.T) {
	igFolder, err := c.GetIGFolder("5", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", igFolder))
}

func TestGetIGFolderByName(*testing.T) {
	igFolder, err := c.GetIGFolder("", "/ubuntu")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", igFolder))
}

func TestGetIGFolders(*testing.T) {
	igFolders, err := c.GetIGFolders()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", igFolders))
}
