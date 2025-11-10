// Package commands
// Is to bifurcate each operation into it's own fucntion, and keep the code loosely coupled
package commands

import (
	"context"

	"github.com/kartikm7/listit-cli/internals/db"
	"github.com/urfave/cli/v3"
)

func It(_ context.Context, cmd *cli.Command) error {
	args := cmd.Args().Slice()
	record := [2]string{}
	copy(record[:], args[:])
	// incase status is not passed we default to Todo
	if record[1] == "" {
		record[1] = db.Todo
	}
	err := db.Add(record)
	if err != nil {
		panic(err)
	}
	return nil
}
