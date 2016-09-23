package main

import (
	"Rachel/mind"
	"flag"
)

func main() {

	// Teach - teach a command to Rachel
	teachFlag := flag.String("teach", "", "Teach a command to Rachel")

	// Run - run a Rachel command
	runFlag := flag.String("run", "", "Run a Rachel command")

	flag.Parse()

	if *teachFlag != "" {
		mind.Teach(*teachFlag)
	}

	if *runFlag != "" {
		mind.Run(*runFlag)
	}
}
