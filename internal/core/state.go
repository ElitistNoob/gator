package core

import (
	"github.com/ElitistNoob/gator/internal/config"
	"github.com/ElitistNoob/gator/internal/database"
)

type State struct {
	DB  *database.Queries
	Cfg *config.Config
}
