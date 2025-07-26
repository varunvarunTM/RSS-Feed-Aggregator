package main

import (
	"fmt"
	"os"
)

func main() {

	config,err := Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	s := &state{
		cfg: &config,
	}

	c := &commands{
		commandMap: make(map[string]func(*state,command)error),
	}

	c.register("login",loginHandler)

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments were provided")
		os.Exit(1)
	}
	
	commandName := os.Args[1]
	arguments := os.Args[2:]
	
	cmd := command {
		name: commandName,
		arguments: arguments,
	}

	err = c.run(s,cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}