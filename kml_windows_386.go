package kml

import "syscall"

var kml = syscall.NewLazyDLL("kml32.dll")
