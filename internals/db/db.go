// Package db is not a typical db wrapper, well it is a wrapper but rather a warpper over the os package
package db

import (
	"fmt"
	"os"
)

// Init func creates a file in the tmp directory by default
// and all your data gets stored there
func Init() error {
	if _, err := os.Stat("/tmp/listit/db.csv"); err != nil {
		fmt.Println("Database does not exist initializing it... (One time thing)")
		// this would mean it does not exist and we must create the db
		if err := os.MkdirAll("/tmp/listit/", 0o755); err != nil {
			panic(err)
		}
		if err := os.WriteFile("/tmp/listit/db.csv", []byte(""), 0o644); err != nil {
			panic(err)
		}
	}
	return nil
}
