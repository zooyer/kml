package kml

import "syscall"

var kml = syscall.NewLazyDLL("kml64.dll")
