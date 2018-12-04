package main

import (
	"fmt"
	"os/exec"
	"github.com/ghst659/eggtimer"
)

func main() {
	r := eggtimer.NewRunner(eggtimer.RealClock{})
	c := exec.Command("/home/tsc/dda/dev/go/hourglass/bin/sandmock")
	events := make(chan eggtimer.Event)
	var segmenter eggtimer.Segmenter
	segmenter.AddDefinition("SandMock", `^Begin\s+(\w+)`, `Finished\s+(\w+)`)
	go r.Run(c, events)
	activities, err := segmenter.Collect(events)
	if err != nil {
		fmt.Printf("Error: %q", err)
	}
	for key, segment := range activities {
		fmt.Printf("%s:\n\t%q\n", key, *segment)
	}
}
