package database

import (
	"fmt"
)

// Delete followedid from userid followings
func (db *appdbimpl) UnlikePhoto(userID string, mediaID string) error {
	// 1 - execute delete query
	_, err := db.c.Exec(`DELETE FROM like WHERE userid=? AND mediaid=?`, userID, mediaID)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error encountered while executing a delete query: %w", err)
	}

	// 3 - return success
	return nil
}
