package main

import (
	"fmt"
)

type (
	command struct {
		name string
		args []string
	}
	commands struct {
		handlerMap map[string]handler
	}
	handler func(*state, command) error
)

func (c *commands) register(name string, f handler) error {
	if _, exists := c.handlerMap[name]; exists {
		return fmt.Errorf("Command %s already registered\n", name)
	}

	c.handlerMap[name] = f
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	if _, exists := c.handlerMap[cmd.name]; !exists {
		return fmt.Errorf("Command %s does not exist\n", cmd.name)
	}
	err := c.handlerMap[cmd.name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}
