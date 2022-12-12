package database

import (
	"fmt"
)

// Delete followedid from userid followings
func (db *appdbimpl) UnlikePhoto(userID string, mediaID string) error {
	// 1 - execute delete query
	res, err := db.c.Exec(`DELETE FROM like WHERE followerid=? AND followedid=?`, userID, mediaID)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error encountered while executing a delete query: %w", err)
	}
	// 2 - check if row has been deleted successfully
	affected, err := res.RowsAffected()
	if err != nil {
		// rowsAffected returned error -> return error
		return fmt.Errorf("error encountered while checking if row has been deleted: %w", err)
	} else if affected == 0 {
		// no row has been deleted -> return error
		// should never happen because it is already checked in the API
		return ErrUserLikeDoesNotExist
	}

	// 3 - return success
	return nil
}
