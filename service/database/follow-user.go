package database

import (
	"fmt"
)

// Add a follow from user with userid to user with followid
func (db *appdbimpl) FollowUser(userid string, followid string) error {

	// 1 - check if user already follows other user
	r, err := db.c.Query(`SELECT * FROM follow WHERE followerid = ? AND followedid = ?`, userid, followid)
	if err != nil {
		return err
	}
	if !r.Next() == false {
		// should never happpen since it is checked in the API
		return ErrUserAlreadyFollowed
	}
	if err = r.Err(); err != nil {
		return err
	}
	defer func() { _ = r.Close() }()
	// here only if user does not follow the other user yet

	// 2 - create new follow
	_, err = db.c.Exec(`INSERT INTO follow (followerid, followedid) VALUES (?,?)`, userid, followid)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error when creating new follow: %w", err)
	}

	// 3 - return nil error
	return nil
}
