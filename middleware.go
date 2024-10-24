package main

import (
	"context"

	"github.com/hawkaii/blogator/internal/database"
)

func middleareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		user, err := s.dbQueries.GetUser(context.Background(), s.cfg.Current_User_Name)

		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
