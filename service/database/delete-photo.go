package database

import (
	"fmt"
)

// Delete media instance and its attachementsfrom database with a specific mediaID
func (db *appdbimpl) DeletePhoto(mediaid string) error {
	// 1 - create query array
	queryArray := []string{
		// delete likes
		`DELETE FROM like WHERE mediaid=?`,
		// delete comments
		`DELETE FROM comment WHERE mediaid=?`,
		// delete media
		`DELETE FROM media WHERE mediaid=?`}

	// 2 - execute query
	for _, query := range queryArray {
		err := db.deleteMedia(query, mediaid)
		if err != nil {
			// delete media function returned error -> return error
			return err
		}
	}

	// 3 - return success
	return nil
}

// auxiliaty function to execute all query necessary to delete a media and its attachments
func (db *appdbimpl) deleteMedia(q string, mediaid string) error {
	// 1 - execute delete query
	res, err := db.c.Exec(q, mediaid)
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
