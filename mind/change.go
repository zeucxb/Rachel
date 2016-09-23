package mind

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Change - change a rachel command
func Change(name string) {
	// Command - the command to be change
	command := strings.Join(flag.Args(), " ")

	commandFile := fmt.Sprintf("/commands/%s.rachel", name)

	_, err := os.Stat(commandFile)
	if err != nil {
		log.Fatalf("%s is not a valid command. Try to teach it. (sudo rachel -teach %s %s)", name, name, command)
	}

	f, err := os.Create(commandFile)
	if err != nil {
		log.Fatalf("Permision problem, please execute rachel as root. (sudo rachel -change %s %s)", name, command)
	}
	defer f.Close()

	err = ioutil.WriteFile(commandFile, []byte(command), 0644)
	if err != nil {
		log.Fatalf("Permision problem, please execute rachel as root. (sudo rachel -change %s %s)", name, command)
	}
}
