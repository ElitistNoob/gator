package main

import (
	"log"
	"os"

	"github.com/ElitistNoob/gator/internal/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
