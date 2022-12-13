package database

import (
	"fmt"
)

// Delete comment with a specific commentID
func (db *appdbimpl) UncommentPhoto(commentID string) error {
	// 1 - execute delete query
	_, err := db.c.Exec(`DELETE FROM comment WHERE commentid=?`, commentID)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error encountered while executing a delete query: %w", err)
	}

	// 3 - return success
	return nil
}
