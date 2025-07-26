package main

import (
	"errors"
	"fmt"
)

type command struct {
	name string
	arguments []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func loginHandler(s *state , cmd command ) error {
	if len(cmd.arguments) == 0 {
		return errors.New("A username is required")
	}

	username := cmd.arguments[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}

func (c *commands) run(s *state , cmd command ) error {
	function,ok := c.commandMap[cmd.name]
	if ok {
		return function(s,cmd)
	} else {
		return errors.New("command doesn't exist")
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}