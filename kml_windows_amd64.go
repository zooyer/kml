package kml

import (
	_ "embed"
	"io/ioutil"
	"syscall"
)

const (
	name = "kml64.dll"
)

var kml = syscall.NewLazyDLL(name)

//go:embed kml64.dll
var lib []byte

func init() {
	if err := ioutil.WriteFile(name, lib, 0644); err != nil {
		panic(err)
	}
}
