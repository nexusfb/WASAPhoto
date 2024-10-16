package database

import (
	"fmt"
)

// Get user profile with username
func (db *appdbimpl) GetUserProfile(username string) (UserProfileDB, error) {
	// 1 - execute query
	user, err := db.c.Query(`SELECT * FROM user WHERE username = ?`, username)
	if err != nil {
		// query returned error -> return error
		return UserProfileDB{}, fmt.Errorf("error encountered while executing select query: %w", err)
	}
	defer func() { _ = user.Close() }()

	// 2 - create result user profile
	var ret UserProfileDB
	if !user.Next() {
		return UserProfileDB{}, ErrUserProfileDoesNotExists
	}
	if err = user.Err(); err != nil {
		return UserProfileDB{}, err
	}
	err = user.Scan(&ret.UserID, &ret.Username, &ret.Bio, &ret.ProfilePic)
	if err != nil {
		// scan returned error -> return error
		return UserProfileDB{}, fmt.Errorf("error encountered while scanning user profile: %w", err)
	}

	// 3 - return result user profile
	return ret, nil

}
