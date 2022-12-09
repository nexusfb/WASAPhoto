package database

import (
	"fmt"
)

// Update user profile with profile in the request body
func (db *appdbimpl) UpdateUserProfile(newProfile UserProfileDB) (UserProfileDB, error) {
	// 1 - take old profile
	oldName, err := db.GetUserName(newProfile.UserID)
	oldProfile, err := db.GetUserProfile(oldName)

	// 2 - execute query
	_, err = db.c.Exec(`UPDATE users SET username=?, bio=?, profilepic=?, WHERE uid=?`,
		newProfile.Username, newProfile.Bio, newProfile.ProfilePic, newProfile.UserID)
	if err != nil {
		// exec returned error -> return error
		return oldProfile, fmt.Errorf("error encountered while executing update query: %w", err)
	}

	// 3 - return new profile
	return newProfile, nil
}
