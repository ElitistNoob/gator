package main

import (
	"log"
	"os"

	"github.com/ElitistNoob/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}

	programState := &state{
		cfg: cfg,
	}

	c := NewCommand()
	c.register("login", handlerLogin)

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
