package database

import (
	"fmt"
)

// Remove ban from bannerid to bannedid
func (db *appdbimpl) UnbanUser(bannerID string, bannedID string) error {
	// 1 - execute delete query
	res, err := db.c.Exec(`DELETE FROM ban WHERE bannerid=? AND bannedid=?`, bannerID, bannedID)
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
		// should never happen since it is already checked in the API
		return ErrUserNotBanned
	}

	// 3 - return success
	return nil
}
