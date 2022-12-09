package database

import (
	"fmt"
)

func (db *appdbimpl) DeleteUserProfile(id string) error {
	res, err := db.c.Exec(`DELETE FROM users WHERE uid=?`, id)
	if err != nil {
		return fmt.Errorf("error encountered while executing a delete query: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrUserProfileDoesNotExists
	}
	return nil
}
