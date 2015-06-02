package main

import (
	"log"
	"os"
	"strconv"

	xmsv3 "github.com/emccode/goxtremio/api/v3"
)

var xms *xmsv3.XMS

func init() {
	endpoint := os.Getenv("GOXTREMIO_ENDPOINT")
	insecure, _ := strconv.ParseBool(os.Getenv("GOXTREMIO_INSECURE"))
	username := os.Getenv("GOXTREMIO_USERNAME")
	password := os.Getenv("GOXTREMIO_PASSWORD")

	var err error
	xms, err = xmsv3.New(endpoint, insecure, username, password)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {

}
