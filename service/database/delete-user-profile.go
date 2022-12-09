package database

import (
	"fmt"
)

// TODO: cacella relativi follows & bans
// Delete user profile instance from database with userid
func (db *appdbimpl) DeleteUserProfile(id string) error {
	// 1 - execute delete query
	res, err := db.c.Exec(`DELETE FROM users WHERE userid=?`, id)
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
		// no row has been deleted which means no user has that input userid -> return error
		return ErrUserProfileDoesNotExists
	}
	return nil
}
