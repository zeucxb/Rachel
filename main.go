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
