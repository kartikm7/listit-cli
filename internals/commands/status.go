package commands

import (
	"context"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kartikm7/listit-cli/internals/db"
	"github.com/urfave/cli/v3"
)

func Status(_ context.Context, cmd *cli.Command) error {
	records, err := db.ReadAll()
	if err != nil {
		panic(err)
	}
	t := table.NewWriter()
	headers := table.Row{}
	for _, v := range records[0] {
		headers = append(headers, v)
	}
	t.AppendHeader(headers)
	for _, record := range records[1:] {
		row := table.Row{}
		for _, v := range record {
			row = append(row, v)
		}
		t.AppendRow(row)
	}
	fmt.Println(t.Render())
	return nil
}
