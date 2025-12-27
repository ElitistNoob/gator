package app

import (
	"database/sql"
	"log"

	"github.com/ElitistNoob/gator/internal/config"
	"github.com/ElitistNoob/gator/internal/core"
	"github.com/ElitistNoob/gator/internal/database"
)

func Initialize() (*core.State, error) {
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

	return state, nil
}
