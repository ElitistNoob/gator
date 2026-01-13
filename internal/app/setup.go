package app

import (
	"database/sql"
	"log"

	"github.com/ElitistNoob/gator/internal/config"
	"github.com/ElitistNoob/gator/internal/core"
	"github.com/ElitistNoob/gator/internal/database"
	_ "github.com/lib/pq"
)

func Initialize(m core.RunMode) (*core.State, error) {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	state := &core.State{
		DB:    database.New(db),
		SQLDB: db,
		Cfg:   cfg,
		Mode:  m,
	}

	return state, nil
}
