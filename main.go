package main

import (
	"log"

	"github.com/ElitistNoob/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}

	state := &State{
		cfg: cfg,
	}

	cmd := &commands{
		cmds: make(map[string]func(*State, command) error),
	}
}
