package database

import (
	"fmt"
)

// TODO: cacella relativi likes e comments
// Delete media instance from database with mediaID
func (db *appdbimpl) DeletePhoto(mediaid string) error {
	// 1 - execute delete query
	res, err := db.c.Exec(`DELETE FROM media WHERE mediaid=?`, mediaid)
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
		// no row has been deleted which means no media has that input mediaid -> return error
		return ErrMediaDoesNotExists
	}
	return nil
}
