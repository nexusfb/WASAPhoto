package database

import (
	"database/sql"
	"errors"
)

// function to check if occurrence exists in database
func (db *appdbimpl) Check(table string, column1 string, v1 string, column2 string, v2 string) bool {
	// 1 - build query to check if an occurrence exists or not
	query := `SELECT * FROM ` + table + ` WHERE ` + column1 + `=? AND ` + column2 + `=?`

	// 2 - execute query
	err := db.c.QueryRow(query, v1, v2).Scan()
	if errors.Is(err, sql.ErrNoRows) {
		// no rows found means no occurrence
		return false
	}
	if err == nil {
		// row found means occurrence
		return true
	}

	// 3 - error occurred when executing query -> return false (for sure no potentially harmful operation is performed)
	return false

}
