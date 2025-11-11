// Package db is not a typical db wrapper, well it is a wrapper but rather a warpper over the os package
package db

import (
	"encoding/csv"
	"fmt"
	"os"
)

const (
	Location = "/tmp/listit/db.csv"
	Todo     = "Todo"
	Doing    = "Doing"
	Done     = "Done"
)

// Init func creates a file in the tmp directory by default
// and all your data gets stored there
func Init() error {
	if _, err := os.Stat(Location); err != nil {
		fmt.Println("Database does not exist initializing it... (One time thing)")
		// this would mean it does not exist and we must create the db
		if err := os.MkdirAll("/tmp/listit/", 0o777); err != nil {
			panic(err)
		}
		file, err := os.Create(Location)
		if err != nil {
			panic(err)
		}
		// deferring makes it run at the end of the function call no matter what
		defer file.Close()

		writer := csv.NewWriter(file)
		// at the end, it will write all the data written into the buffer
		defer writer.Flush()

		// these are the columns and should be added first
		header := []string{"Task", "Status"}
		if err := writer.Write(header); err != nil {
			panic(err)
		}
	}
	return nil
}

func Add(record [2]string) error {
	file, err := os.OpenFile(Location, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err = writer.Write(record[:]); err != nil {
		return err
	}
	if err = writer.Error(); err != nil {
		return err
	}
	return nil
}

func ReadAll() ([][]string, error) {
	file, err := os.Open(Location)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
