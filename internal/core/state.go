package core

import (
	"database/sql"

	"github.com/ElitistNoob/gator/internal/config"
	"github.com/ElitistNoob/gator/internal/database"
)

type RunMode int

const (
	ModeCLI RunMode = iota
	ModeTUI
)

type State struct {
	DB    *database.Queries
	SQLDB *sql.DB
	Cfg   *config.Config
	Mode  RunMode
}
