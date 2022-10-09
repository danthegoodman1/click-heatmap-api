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
	DuckDB, err = sql.Open("duckdb", "./.ddb")
	if err != nil {
		return fmt.Errorf("error opening duckdb: %w", err)
	}
	_, err = DuckDB.Exec("INSTALL 'json'")
	if err != nil {
		return fmt.Errorf("error installing JSON extension: %w", err)
	}
	_, err = DuckDB.Exec("LOAD 'json'")
	if err != nil {
		return fmt.Errorf("error loading JSON extension: %w", err)
	}
	return
}
