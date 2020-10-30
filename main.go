package main

import (
	"fmt"

	"github.com/LukasKnuth/sonos_ctrl/sonos"
)

func main() {
	fmt.Println("Starting Sonos discovery")
	discovery := sonos.Discover()
	for true {
		found := <-discovery
		fmt.Println("Found: ", found)
	}
}
