package main

import (
	"syscall/js"

	"github.com/n4ze3m/zippy/zippy"
)

func main() {
	// create zip file
	js.Global().Set("Zippy", zippy.CreateArchiveFile())
	<-make(chan bool)
}
