package kml

import (
	_ "embed"
	"io/ioutil"
	"syscall"
)

var kml = syscall.NewLazyDLL("kml32.dll")

//go:embed kml32.dll
var lib []byte

func init() {
	if err := ioutil.WriteFile("./kml32.dll", lib, 0644); err != nil {
		panic(err)
	}
}