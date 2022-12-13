package database

import (
	"fmt"
)

// Delete media instance and its attachementsfrom database with a specific mediaID
func (db *appdbimpl) DeletePhoto(mediaid string) error {
	// 1 - create query array
	queryArray := []string{
		`DELETE FROM like WHERE mediaid=?`,    // delete media likes
		`DELETE FROM comment WHERE mediaid=?`, // delete media comments
		`DELETE FROM media WHERE mediaid=?`}   // delete media

	// 2 - execute query
	for _, query := range queryArray {
		_, err := db.c.Exec(query, mediaid)
		if err != nil {
			// exec returned error -> return error
			return fmt.Errorf("error encountered while executing a delete query: %w", err)
		}
	}

	// 3 - return success
	return nil
}
