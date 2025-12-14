package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ElitistNoob/gator/internal/config"
	"github.com/ElitistNoob/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}
	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: cfg,
	}

	c := NewCommand()
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("not enough arguments were provided")
		os.Exit(1)
	}

	if err := c.Run(programState, command{name: args[1], args: args[2:]}); err != nil {
		log.Fatalf("error: %v\n", err)
		os.Exit(1)
	}
}
