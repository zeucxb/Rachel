package mind

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func Run(name string) {
	commandFile := fmt.Sprintf("./commands/%s.rachel", name)

	command, err := ioutil.ReadFile(commandFile)
	if err != nil {
		fmt.Println(err)
	}

	output, err := exec.Command("sh", "-c", string(command)).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}
