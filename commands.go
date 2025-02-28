package main

import (
	"errors"
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) error {
	if _, exists := c.registeredCommands[name]; exists {
		return fmt.Errorf("Command %s already registered\n", name)
	}

	c.registeredCommands[name] = f
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	if _, exists := c.registeredCommands[cmd.Name]; !exists {
		return errors.New("command not found")
	}
	err := c.registeredCommands[cmd.Name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}
