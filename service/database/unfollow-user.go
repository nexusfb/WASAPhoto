package database

import (
	"fmt"
)

// Delete followedid from userid followings
func (db *appdbimpl) UnfollowUser(userid string, followedid string) error {
	// 1 - execute delete query
	_, err := db.c.Exec(`DELETE FROM follow WHERE followerid=? AND followedid=?`, userid, followedid)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error encountered while executing a delete query: %w", err)
	}

	// 2 - return success
	return nil
}
