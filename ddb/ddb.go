package ddb

import (
	"database/sql"
	"fmt"

	_ "github.com/marcboeker/go-duckdb"
)

var (
	DuckDB *sql.DB
)

func ConnectDDB() (err error) {
	DuckDB, err = sql.Open("duckdb", "")
	if err != nil {
		return fmt.Errorf("error opening duckdb: %w", err)
	}
	return
}
