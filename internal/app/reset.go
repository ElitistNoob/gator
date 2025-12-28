package app

import (
	"context"
	"fmt"

	"github.com/ElitistNoob/gator/internal/core"
)

func ResetDB(s *core.State, c core.Command) (string, error) {
	if err := s.DB.DeleteUsers(context.Background()); err != nil {
		return "", fmt.Errorf("couldn't delete rows in table: %w", err)
	}

	return "rows deleted successfully", nil
}
