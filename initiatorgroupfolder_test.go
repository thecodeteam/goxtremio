package main

import (
	"fmt"
	"testing"
)

func TestGetIGFolderByID(*testing.T) {
	igFolder, err := GetIGFolder("5", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", igFolder))
}

func TestGetIGFolderByName(*testing.T) {
	igFolder, err := GetIGFolder("", "/ubuntu")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", igFolder))
}

func TestGetIGFolders(*testing.T) {
	igFolders, err := GetIGFolders()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", igFolders))
}
