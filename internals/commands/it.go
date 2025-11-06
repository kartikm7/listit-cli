// Package commands
// Is to bifurcate each operation into it's own fucntion, and keep the code loosely coupled
package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func It(_ context.Context, cmd *cli.Command) error {
	args := cmd.Args().Slice()
	for _, v := range args {
		fmt.Println(v)
	}
	return nil
}
