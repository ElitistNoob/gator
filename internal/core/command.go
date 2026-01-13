package core

import (
	"context"
)

type Command struct {
	Name string
	Args []string
	Run  func(ctx context.Context, s *State, args []string) (string, error)
}
