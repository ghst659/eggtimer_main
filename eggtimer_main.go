package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"github.com/ghst659/eggtimer"
)

type RegexpDef struct {
	name string
	reStart *regexp.Regexp
	reFinish *regexp.Regexp
}

func (d RegexpDef) TypeName() string {
	return d.name
}

func (d RegexpDef) IsStart(line string) string {
	matches := d.reStart.FindStringSubmatch(line)
	if matches == nil {
		return ""
	}
	return matches[1]
}

func (d RegexpDef) IsFinish(line string) string {
	matches := d.reFinish.FindStringSubmatch(line)	
	if matches == nil {
		return ""
	}
	return matches[1]
}

func main() {
	r := eggtimer.NewRunner(eggtimer.RealClock{})
	c := exec.Command("/home/tsc/dda/dev/go/hourglass/bin/sandmock")
	events := make(chan eggtimer.Event)
	smd := RegexpDef {
		name: "SandMock",
		reStart: regexp.MustCompile(`^Begin\s+(\w+)`),
		reFinish: regexp.MustCompile(`Finished\s+(\w+)`),
	}
	var segmenter eggtimer.Segmenter
	segmenter.AddDefinition(smd)
	go r.Run(c, events)
	activities, err := segmenter.Collect(events)
	if err != nil {
		fmt.Printf("Error: %q", err)
	}
	for key, segment := range activities {
		fmt.Printf("%s:\n\t%q\n", key, *segment)
	}
}
