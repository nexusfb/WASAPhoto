package database

import (
	"fmt"
)

// Add a like from user with userid to media with mediaid
func (db *appdbimpl) LikePhoto(userID string, mediaID string) error {
	// we do not check if user has already liked this photo because it is already checked in the API before calling this function
	// here only if user has not liked this photo yet

	// 1 - create new like
	_, err := db.c.Exec(`INSERT INTO like (userid, mediaid) VALUES (?,?)`, userID, mediaID)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error when creating new like: %w", err)
	}

	// 2 - return nil error
	return nil
}
