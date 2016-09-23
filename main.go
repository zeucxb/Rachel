package main

import (
	"Rachel/mind"
	"flag"
	"log"
	"os"
)

func main() {

	_, err := os.Stat("/commands/")
	if err != nil {
		if err := os.Mkdir("/commands/", 0777); err != nil {
			log.Fatalln("Permision problem, please execute rachel as root. (sudo rachel)")
		}
	}

	// Teach - teach a command to Rachel
	teachFlag := flag.String("teach", "", "Teach a command to rachel")

	// Run - run a Rachel command
	runFlag := flag.String("run", "", "Run a rachel command")

	// Teach - teach a command to Rachel
	changeFlag := flag.String("change", "", "Change a rachel command")

	flag.Parse()

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
