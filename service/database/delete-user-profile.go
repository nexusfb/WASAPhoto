package database

import (
	"fmt"
)

// Delete user profile instance from database with a specific userid
func (db *appdbimpl) DeleteUserProfile(userid string) error {
	// notice: 2 steps -> 1 delete user media 2 delete user profile and its attachments

	// STEP 1: delete user media
	// 1 - get all media published by user
	media, err := db.GetUserMedia(userid)
	if err != nil {
		// get user media returned error -> return error
		return err
	}
	// 2 - iterate over each media and delete it
	for _, m := range media {
		err = db.DeletePhoto(m.MediaID)
		if err != nil {
			// delete photo returned error -> return error
			return err
		}
	}

	// STEP 2: delete user profile
	// 1 - create query array
	queryArray := []string{
		`DELETE FROM follow WHERE followerid=?`, // delete followers
		`DELETE FROM follow WHERE followedid=?`, // delete followings
		`DELETE FROM ban WHERE bannerid=?`,      // delete bans (users banned by user)
		`DELETE FROM ban WHERE bannedid=?`,      // delete bans (bans of the user by other users)
		`DELETE FROM like WHERE userid=?`,       // delete likes of the user
		`DELETE FROM comment WHERE userid=?`,    // delete comments of the user
		`DELETE FROM user WHERE userid=?`,       // delete user profile
	}

	// 2 - execute query
	for _, query := range queryArray {
		_, err := db.c.Exec(query, userid)
		if err != nil {
			// exec returned error -> return error
			return fmt.Errorf("error encountered while executing a delete query: %w", err)
		}
	}

	// 3 - return success
	return nil
}
