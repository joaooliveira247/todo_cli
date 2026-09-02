package main

import (
	"github.com/joaooliveira247/todo_cli/cmd"
	"github.com/joaooliveira247/todo_cli/internal/db"
)

func main() {
	database, err := db.InitDB()
	defer database.Close()
	if err != nil {
		panic(err)
	}

	app := cmd.NewApp(database)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
