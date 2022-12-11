package database

import "fmt"

// Get comments of media with specific mediaid
func (db *appdbimpl) GetMediaComments(mediaID string) ([]CommentDB, error) {
	// 1 - execute query
	rawComments, err := db.c.Query(`SELECT * FROM comment WHERE mediaid = ?`, mediaID)
	if err != nil {
		// query returned error -> return error
		return nil, fmt.Errorf("error encountered while executing a select query: %w", err)
	}
	defer func() { _ = rawComments.Close() }()

	// 2 - retrieve rows and build followers array result
	var ret []CommentDB
	for rawComments.Next() {
		var c CommentDB
		err = rawComments.Scan(&c.CommentID, &c.MediaID, &c.AuthorID, &c.Date, &c.Content)
		if err != nil {
			// scan returned error -> return error
			return nil, err
		}
		ret = append(ret, c)
	}
	if err = rawComments.Err(); err != nil {
		return nil, err
	}
	// 3 - return result media array
	return ret, nil
}
