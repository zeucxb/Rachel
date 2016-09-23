package mind

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Teach - teach a command to Rachel
func Teach(name string) {
	// Command - the command to be teach
	command := strings.Join(flag.Args(), " ")

	commandFile := fmt.Sprintf("/commands/%s.rachel", name)

	_, err := os.Stat(commandFile)
	if err == nil {
		log.Fatalf("The %s command already exists. Try to change it. (sudo rachel -change %s %s)", name, name, command)
	}

	f, err := os.Create(commandFile)
	if err != nil {
		log.Fatalf("Permision problem, please execute rachel as root. (sudo rachel -teach %s %s)", name, command)
	}
	defer f.Close()

	err = ioutil.WriteFile(commandFile, []byte(command), 0644)
	if err != nil {
		log.Fatalf("Permision problem, please execute rachel as root. (sudo rachel -teach %s %s)", name, command)
	}
}
