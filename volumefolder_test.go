package goxtremio

import (
	"fmt"
	"testing"
)

func TestGetVolumeFolderByID(t *testing.T) {
	volumeFolder, err := c.GetVolumeFolder("7", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumeFolder))
}

func TestGetVolumeFolderByName(*testing.T) {
	volumeFolder, err := c.GetVolumeFolder("", "/Ubuntu")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumeFolder))
}

func TestGetVolumeFolders(*testing.T) {
	volumeFolders, err := c.GetVolumeFolders()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumeFolders))
}
