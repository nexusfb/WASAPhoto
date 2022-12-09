package database

import (
	"fmt"
)

func (db *appdbimpl) GetUserProfile(username string) (ProfileDB, error) {
	const query = `
	SELECT *
	FROM users
	WHERE username = ?`

	var ret ProfileDB
	user, err := db.c.Query(query, username)
	if err != nil {
		return ProfileDB{}, err
	}
	defer func() { _ = user.Close() }()
	result := user.Next()
	fmt.Println(result)
	if !result {
		return ProfileDB{}, ErrUserProfileDoesNotExists
	}

	err = user.Scan(&ret.UserID, &ret.Username, &ret.Bio, &ret.ProfilePic, &ret.NMedia, &ret.NFollowers, &ret.NFollowing)
	if err != nil {
		return ProfileDB{}, fmt.Errorf("error encountered while scanning user profile: %w", err)
	}
	return ret, nil

}
