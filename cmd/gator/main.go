package main

import (
	"log"
	"os"

	"github.com/ElitistNoob/gator/internal/cli"
	"github.com/ElitistNoob/gator/internal/tui"
)

func main() {
	if len(os.Args) > 1 {
		_, err := cli.Run(os.Args)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if err := tui.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
