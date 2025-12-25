package cli

import (
	"database/sql"
	"log"

	"github.com/ElitistNoob/gator/internal/app"
	"github.com/ElitistNoob/gator/internal/config"
	"github.com/ElitistNoob/gator/internal/core"
	"github.com/ElitistNoob/gator/internal/database"
	_ "github.com/lib/pq"
)

func Run(args []string) error {
	if len(args) < 2 {
		log.Fatalln("not enough arguments were provided")
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	state := &core.State{
		DB:  dbQueries,
		Cfg: cfg,
	}

	c := NewCommand()

	// Users commands
	c.register("login", app.Login)
	c.register("register", app.RegisterUser)
	c.register("reset", app.ResetDB)
	c.register("users", app.GetUsers)

	// Feeds commands
	c.register("agg", app.Agg)
	c.register("addfeed", app.MiddlewareLoggedIn(app.AddFeed))
	c.register("feeds", app.GetFeeds)
	c.register("follow", app.MiddlewareLoggedIn(app.FollowFeed))
	c.register("following", app.MiddlewareLoggedIn(app.Following))
	c.register("unfollow", app.MiddlewareLoggedIn(app.Unfollow))

	// Posts commands
	c.register("browse", app.MiddlewareLoggedIn(app.BrowsePosts))

	return c.Run(state, core.Command{Name: args[1], Args: args[2:]})
}
