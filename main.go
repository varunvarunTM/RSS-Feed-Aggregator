package main

import (
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"RSS-feed-aggregator/config"
	"RSS-feed-aggregator/internal/database"
	"database/sql"
)

func main() {
	cfg,err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	dbUrl := cfg.DbUrl

	db,err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	dbQueries := database.New(db)

	s := &state{
		db: dbQueries,
		cfg: &cfg,
	}

	cmds := &commands{
		commandMap: make(map[string]func(*state,command)error),
	}

	cmds.register("login",loginHandler)
	cmds.register("register",registerHandler)

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

	err = cmds.run(s,cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}