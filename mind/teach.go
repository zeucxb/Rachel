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

	f, err := os.Create(commandFile)
	if err != nil {
		log.Fatalln("Permision problem, please execute rachel as root. (sudo rachel)")
	}
	defer f.Close()

	err = ioutil.WriteFile(commandFile, []byte(command), 0644)
	if err != nil {
		log.Fatalln("Permision problem, please execute rachel as root. (sudo rachel)")
	}
}
