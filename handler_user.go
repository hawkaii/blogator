package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hawkaii/blogator/internal/database"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]

	_, err := s.dbQueries.GetUser(context.Background(), name)
	if err == nil {
		fmt.Println("user already exists")

		os.Exit(1)
	}

	user, err := s.dbQueries.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.dbQueries.GetUser(context.Background(), name)
	if err != nil {
		fmt.Print("couldn't find user: %w", err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerReset(s *state, cmd command) error {
	// Call the DeleteUsers query
	err := s.dbQueries.DeleteUsers(context.Background())

	// Handle any errors
	if err != nil {
		fmt.Println("error deleting users")
		os.Exit(1)
	}
	// Print a success or error message
	fmt.Println("users deleted")
	// Exit with the appropriate code
	return nil
}

func handlerGetUsers(s *state, cmd command) error {

	// call the listusers query
	users, err := s.dbQueries.GetUsers(context.Background())
	// handle any errors
	if err != nil {
		fmt.Println("error listing users")
		os.Exit(1)
	}
	// Print the users
	fmt.Println("Users:")
	for _, user := range users {
		if user.Name == s.cfg.Current_User_Name {
			fmt.Printf(" * Name:    %v (current)\n", user.Name)
		} else {
			printUser(user)
		}
	}
	// Exit with the appropriate code
	return nil
}

func printUser(user database.User) {

	fmt.Printf(" * Name:    %v\n", user.Name)
}
