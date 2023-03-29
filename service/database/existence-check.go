package database

import (
	"errors"
)

// function to check if user/media exists or not
func (db *appdbimpl) ExistenceCheck(id string, what string) bool {
	// user case
	if what == "user" {
		// 1 - get username of the user
		_, err := db.GetUserName(id)
		if errors.Is(err, ErrUserProfileDoesNotExists) {
			// err user does not exist -> return false
			return false
		}
		if err == nil {
			// user exists -> return true
			return true
		}
		// 2 - error occurred when looking for username -> return false ( no possibly harmful operation is performed)
		return false
	}

	// username case
	if what == "username" {
		// 1 - get userid of the user
		_, err := db.GetUserID(id)
		if errors.Is(err, ErrUserProfileDoesNotExists) {
			// err user does not exist -> return false
			return false
		}
		if err == nil {
			// user exists -> return true
			return true
		}
		// 2 - error occurred when looking for username -> return false (no possibly harmful operation is performed)
		return false
	}

	// media case
	if what == "media" {
		// 1 - get media
		_, err := db.GetMedia(id)
		if errors.Is(err, ErrMediaDoesNotExists) {
			// err media does not exist -> return false
			return false
		}
		if err == nil {
			// media exists -> return true
			return true
		}

		// 2 - error occurred when looking for username -> return false ( no possibly harmful operation is performed)
		return false
	}

	// comment case
	if what == "comment" {
		// 1 - execute query
		comment, err := db.c.Query(`SELECT * FROM comment WHERE commentid = ?`, id)
		if err != nil {
			// query returned error -> return false
			return false
		}
		defer func() { _ = comment.Close() }()
		// 2 - check if at least one row
		if !comment.Next() {
			return false
		}
		if err = comment.Err(); err != nil {
			return false
		}
		return true
	}

	// never used
	return false

}
