package cli

import (
	"fmt"
	"log"

	"github.com/ElitistNoob/gator/internal/app"
	"github.com/ElitistNoob/gator/internal/core"
)

func Run(args []string) (string, error) {
	if len(args) < 2 {
		log.Fatalln("not enough arguments were provided")
	}

	state, err := app.Initialize(core.ModeCLI)
	if err != nil {
		log.Fatalf("failed to initialize app: %s", err)
	}
	defer state.SQLDB.Close()

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

	out, err := c.Run(state, core.Command{Name: args[1], Args: args[2:]})
	if err != nil {
		log.Fatal(err)
	}

	if state.Mode == core.ModeCLI {
		fmt.Print(out)
	}

	return out, err
}
