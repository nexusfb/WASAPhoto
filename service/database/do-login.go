package database

import (
	"fmt"

	"github.com/gofrs/uuid"
)

// If username does not exist create new user profile, if it exists login
func (db *appdbimpl) DoLogin(username string) (string, error) {
	// 1 - check if user exists by trying to get userprofile
	profile, err := db.GetUserProfile(username)
	if err == nil {
		// user profile already exists, return userID
		return profile.UserID, nil
	} else if err != ErrUserProfileDoesNotExists {
		// get user profile database function returned error during execution -> return error
		return "00000000", fmt.Errorf("error encountered while checking if profile exists: %w", err)
	}

	// here only if get user profile database function returner error user profile does not exist

	// 2 - create new userID
	rawUid, err := uuid.NewV4()
	if err != nil {
		// newV4 returned error -> return error
		return "00000000", fmt.Errorf("error encountered while creating new userID: %w", err)
	}
	uid := rawUid.String()

	// 3 - execute query
	_, err = db.c.Exec(`INSERT INTO users (userid, username) VALUES (?,?)`, uid, username)
	if err != nil {
		// exec returned error -> return error
		return "00000000", fmt.Errorf("error when creating new user profile: %w", err)
	}

	// 4 - return new userID
	return uid, nil
}
