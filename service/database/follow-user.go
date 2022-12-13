package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// Add a follow from user with userid to user with followid
func (db *appdbimpl) FollowUser(userid string, followid string) error {

	// 1 - check if user already follows other user
	err := db.c.QueryRow(`SELECT * FROM follow WHERE followerid = ? AND followedid = ?`, userid, followid).Scan()
	if !errors.Is(err, sql.ErrNoRows) {
		// should never happen because it is already checked in the API
		return ErrUserAlreadyFollowed
	}

	// here only if user does not follow the other user yet

	// 2 - create new follow
	_, err = db.c.Exec(`INSERT INTO follow (followerid, followingid) VALUES (?,?)`, userid, followid)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error when creating new follow: %w", err)
	}

	// 3 - return nil error
	return nil
}
