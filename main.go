package main

import (
	"Rachel/mind"
	"flag"
	"log"
	"os"
	"strings"
)

// Teach - teach a command to Rachel
var teachFlag = flag.String("teach", "", "Teach a command to rachel")

// Run - run a Rachel command
var runFlag = flag.String("run", "", "Run a rachel command")

// Change - change a rachel command
var changeFlag = flag.String("change", "", "Change a rachel command")

func init() {
	flag.Parse()
}

func main() {
	_, err := os.Stat("/rachel/")
	if err != nil {
		if err := os.Mkdir("/rachel/", 0777); err != nil {
			command := strings.Join(os.Args, " ")
			log.Fatalf("Permision problem, please execute rachel as root. (sudo %s)", command)
		}
	}

	if *teachFlag != "" {
		mind.Teach(*teachFlag)
	}

	if *runFlag != "" {
		mind.Run(*runFlag)
	}

	if *changeFlag != "" {
		mind.Change(*changeFlag)
	}
}
