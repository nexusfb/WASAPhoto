package database

import (
	"fmt"
)

func (db *appdbimpl) GetUserProfile(username string) (UserProfileDB, error) {
	const query = `
	SELECT *
	FROM users
	WHERE username = ?`

	var ret UserProfileDB
	user, err := db.c.Query(query, username)
	if err != nil {
		return UserProfileDB{}, err
	}
	defer func() { _ = user.Close() }()
	result := user.Next()
	fmt.Println(result)
	if !result {
		return UserProfileDB{}, ErrUserProfileDoesNotExists
	}

	err = user.Scan(&ret.UserID, &ret.Username, &ret.Bio, &ret.ProfilePic, &ret.NMedia, &ret.NFollowers, &ret.NFollowing)
	if err != nil {
		return UserProfileDB{}, fmt.Errorf("error encountered while scanning user profile: %w", err)
	}
	return ret, nil

}
