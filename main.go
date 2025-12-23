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
	}

	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: cfg,
	}

	c := NewCommand()

	// Users commands
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerGetUsers)

	// Feeds commands
	c.register("agg", agg)
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerGetFeeds)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("following", middlewareLoggedIn(handlerFollowing))
	c.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	// Posts commands
	c.register("browse", middlewareLoggedIn(handlerBrowse))

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("not enough arguments were provided")
	}

	if err := c.Run(programState, command{name: args[1], args: args[2:]}); err != nil {
		log.Fatalf("error: %v\n", err)
	}
}
