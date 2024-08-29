package database

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func InitDbConnection() {
	filename := "deadline_cal.sqlite3"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
	} else {
		_, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			panic(err)
		}
	}

	ctx := context.Background()

	db, err := sql.Open("sqlite3", "deadline_cal.sqlite3")
	if err != nil {
		panic(err)
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		fmt.Println(err)
	}

	queries := New(db)

	createTask, err := queries.CreateTask(ctx, CreateTaskParams{
		TaskName:  "test Task",
		StartDate: "09/08/2024",
		EndDate:   "09/08/2024",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(createTask)

}
