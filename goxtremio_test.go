package goxtremio

import (
	"os"
	"testing"
)

var c *Client

func TestMain(m *testing.M) {
	var err error
	c, err = NewClient()
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
