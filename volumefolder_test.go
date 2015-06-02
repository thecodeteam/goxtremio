package main

import (
	"fmt"
	"testing"
)

func TestGetVolumeFolderByID(*testing.T) {
	volumeFolder, err := GetVolumeFolder("7", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumeFolder))
}

func TestGetVolumeFolderByName(*testing.T) {
	volumeFolder, err := GetVolumeFolder("", "/Ubuntu")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumeFolder))
}

func TestGetVolumeFolders(*testing.T) {
	volumeFolders, err := GetVolumeFolders()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumeFolders))
}
