package database

import (
	"fmt"
)

// TODO: cacella relativi follows & bans
// Delete user profile instance from database with userid
func (db *appdbimpl) DeleteUserProfile(userid string) error {
	// 1 - create query array
	queryArray := []string{
		// delete followers
		`DELETE FROM follow WHERE followerid=?`,
		// delete followings
		`DELETE FROM follow WHERE followedid=?`,
		// delete bans
		// delete likes
		// delete comments
		// delete posts
		`DELETE FROM media WHERE authorid=?`,
		// delete user profile
		`DELETE FROM users WHERE userid=?`}

	// 2 - execute query
	for _, query := range queryArray {
		err := db.deleteUser(query, userid)
		if err != nil {
			return err
		}
	}

	return nil
}

// 3 - create function to delete user given a query
func (db *appdbimpl) deleteUser(q string, userid string) error {
	res, err := db.c.Exec(q, userid)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error encountered while executing a delete query: %w", err)
	}
	// 4 - check if row has been deleted successfully
	affected, err := res.RowsAffected()
	if err != nil {
		// rowsAffected returned error -> return error
		return fmt.Errorf("error encountered while checking if row has been deleted: %w", err)
	} else if affected == 0 {
		// no row has been deleted -> return error
		return fmt.Errorf("error: no row has been deleted which means no corresponding row was found: %w", err)
	}
	return nil
}
