package main

import (
	"errors"
	"fmt"
	"time"
	"context"
	"github.com/google/uuid"
	"os"
	"database/sql"
	"RSS-feed-aggregator/internal/database"
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

	user,err := s.db.GetUser(context.Background(),username)
	if err == sql.ErrNoRows {
    	os.Exit(1)
	}
	if err != nil {
		fmt.Println("Error in GetUser:", err)
    	return err
	}
	
	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User has been set")
	fmt.Printf("User data: %+v\n",user)
	return nil
}

func registerHandler(s *state , cmd command ) error {
	if len(cmd.arguments) == 0 {
		return errors.New("A username is required")
	}
	username := cmd.arguments[0]

	user,err := s.db.GetUser(context.Background(),username)
	if err == nil {
		fmt.Println("Duplicate user found; will exit 1")
    	os.Exit(1)
	}
	if err != sql.ErrNoRows {
		fmt.Println("Error in GetUser:", err)
    	return err
	}

	user,err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: username,
	})
	if err != nil {
		fmt.Println("Error creating user:", err)
		return err
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User was created")
	fmt.Printf("User data: %+v\n",user)
	return nil
}

func deleteHandler(s *state , cmd command ) error {
	err := s.db.DeleteUsers(context.Background())
	if err !=  nil {
		fmt.Println("Failure to delete users")
		return err
	}
	fmt.Println("Success, users deleted")
	return nil
}

func users(s *state , cmd command ) error {
	users,err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Failure to get all users")
		return err
	}
	
	for _,user := range users {
		fmt.Printf("* %s",user)
		if user == s.cfg.CurrentUserName {
			fmt.Printf(" (current)")
		}
		fmt.Printf("\n")
	}
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