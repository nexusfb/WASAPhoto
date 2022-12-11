package database

import (
	"fmt"
)

// Delete user profile instance from database with a specific userid
func (db *appdbimpl) DeleteUserProfile(userid string) error {
	// notice: 2 steps -> 1 delete user media 2 delete user profilen and its attachments

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
		`DELETE FROM follow WHERE followedid=?`, // delete likes of the user
		`DELETE FROM follow WHERE followedid=?`, // delete comments of the user
		`DELETE FROM user WHERE userid=?`,       // delete user profile
	}

	// 2 - execute query
	for _, query := range queryArray {
		err := db.deleteUser(query, userid)
		if err != nil {
			return err
		}
	}

	// 3 - return success
	return nil
}

// auxiliaty function to execute all query necessary to delete a user profile and its attachments
func (db *appdbimpl) deleteUser(q string, userid string) error {
	// 1 - execute delete query
	res, err := db.c.Exec(q, userid)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error encountered while executing a delete query: %w", err)
	}
	// 2 - check if row has been deleted successfully
	affected, err := res.RowsAffected()
	if err != nil {
		// rowsAffected returned error -> return error
		return fmt.Errorf("error encountered while checking if row has been deleted: %w", err)
	} else if affected == 0 {
		// no row has been deleted -> return error
		return fmt.Errorf("error: no row has been deleted which means no corresponding row was found: %w", err)
	}

	// 3 - return success
	return nil
}
