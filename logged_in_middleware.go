package main

import (
	"context"
	"fmt"

	"github.com/mauricebonnesdev/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("User with name %s not found: %v", s.cfg.CurrentUserName, err)
		}

		return handler(s, cmd, user)
	}
}
