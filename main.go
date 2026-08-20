package main

import (
	"github.com/joaooliveira247/todo_cli/cmd"
)

func main() {
	app := cmd.NewApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
