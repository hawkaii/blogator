package main

import (
	"database/sql"
	"fmt"

	"log"
	"os"

	"github.com/hawkaii/blogator/internal/config"
	"github.com/hawkaii/blogator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	dbQueries *database.Queries
	cfg       *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.Db_Url)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		dbQueries: dbQueries,
		cfg:       &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handler_agg)
	cmds.register("addfeed", middleareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)
	cmds.register("unfollow", middleareLoggedIn(handlerUnfollow))
	cmds.register("browse", middleareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		return

	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		fmt.Println("Error:", err)
	}

}
