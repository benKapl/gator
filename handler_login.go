package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no username specified")
	}

	userName := cmd.args[0]

	err := s.cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("%s is logged in\n", userName)
	return nil
}
