package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// Add a follow from user with userid in path to user with followid in path
func (db *appdbimpl) FollowUser(userid string, followid string) error {
	// 1 - check if user already follows followid
	var follow FollowDB
	err := db.c.QueryRow(`SELECT * FROM follow WHERE followerid = ? AND followedid = ?`, userid, followid).Scan(&follow)
	if !errors.Is(err, sql.ErrNoRows) {
		return ErrUserAlreadyFollowed
	}

	// 2 - create new follow
	_, err = db.c.Exec(`INSERT INTO follow (followerid, followingid) VALUES (?,?)`, userid, followid)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error when creating new follow: %w", err)
	}

	// 3 - return nil error
	return nil
}
