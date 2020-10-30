package main

import (
	"fmt"

	"github.com/LukasKnuth/sonos_ctrl/sonos"
	"github.com/LukasKnuth/sonos_ctrl/sonos/rx"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	fmt.Println("Starting Sonos discovery")

	observable := rxgo.FromChannel(sonos.Discover()).Distinct(rx.StatefulDistinctController().Distinct)

	for controller := range observable.Observe() {
		fmt.Println("Found: ", controller.V)
	}
}
