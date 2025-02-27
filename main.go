package main

import (
	"fmt"
	"log"
	"os"

	"github.com/benKapl/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	state := state{cfg: &cfg}

	cmds := commands{
		handlerMap: make(map[string]handler),
	}

	cmds.handlerMap["login"] = handlerLogin // Register login command

	cliArgs := os.Args
	if len(cliArgs) < 2 {
		log.Fatalf("not enough arguments supplied")
	}

	cmd := command{
		name: cliArgs[1],
		args: cliArgs[2:],
	}
	fmt.Printf("Command: %+v\n", cmd)

	err = cmds.run(&state, cmd)
	if err != nil {
		log.Fatalf("%v", err)
	}

}
