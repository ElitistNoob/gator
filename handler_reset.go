package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, c command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("couldn't delete rows in table: %w", err)
	}

	fmt.Println("rows deleted successfully")
	return nil
}
