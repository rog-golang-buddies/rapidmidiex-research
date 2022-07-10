package main

import (
	"log"

	"github.com/rog-golang-buddies/realtime-midi/examples/wspingpong"
)

func main() {
	// Make log print a datetime and a filename:linenumber
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	wspingpong.StartServer(":9876", wspingpong.LogLevelBasicWithHeaders)
}
