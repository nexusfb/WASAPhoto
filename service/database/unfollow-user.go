package database

import (
	"fmt"
)

// Delete followedid from userid followings
func (db *appdbimpl) UnfollowUser(userid string, followedid string) error {
	// 1 - execute delete query
	res, err := db.c.Exec(`DELETE FROM follow WHERE followerid=? AND followedid=?`, userid, followedid)
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
		return ErrUserNotFollowed
	}

	// 3 - return success
	return nil
}