package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kartikm7/listit-cli/internals/commands"
	"github.com/kartikm7/listit-cli/internals/db"
	"github.com/urfave/cli/v3"
)

func main() {
	// initializing the db first
	db.Init()

	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:        "it",
				Description: "list it as in your task down",
				Usage:       "pass a string as an argument",
				Action:      commands.It,
			},
		},
	}

	// this is super interesting, I'm from the javascript world so this is a litte freaky to me but here's the breakdown
	// & is for pointers, so we initialize a Struct that returns it's memory address which is then used to invoke it's
	// Run method, which expects to be called on a Pointer that holds a value of type Command
	// Now, run takes in two arguments a Context and Arguments from the OS
	// 1. Context: It's sole purpose is to have elegant stopping of goroutines running in the background,
	// the Background context is ever running until the cancel() method is invoked and that can get invoked through multiple scenarios
	// like CTRL + C, force quiting the pane, shutting down the OS, and much more.
	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println("fuck hogaya")
	}
}
