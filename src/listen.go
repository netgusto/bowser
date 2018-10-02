package main

// From https://gist.github.com/nathankerr/38d8b0d45590741b57f5f79be336f07c

/*
#cgo CFLAGS: -x objective-c -Wno-incompatible-pointer-types-discards-qualifiers
#cgo LDFLAGS: -framework Foundation
#include "handler.h"
*/
import "C"

import (
	"github.com/andlabs/ui"
)

//export ReceiveURL
func ReceiveURL(u *C.char) {
	urlStream <- C.GoString(u)
}

func listen(config Config) chan bool {

	urlStream = make(chan string, 1) // the event handler blocks!, so buffer the channel at least once to get the first message
	C.StartURLHandler()

	done := make(chan bool)

	ui.Main(func() {
		ui.QueueMain(func() { ui.Quit() })

		go func() {
			for url := range urlStream {
				<-handleURL(config, url)
				done <- true
			}
		}()
	})

	return done
}
