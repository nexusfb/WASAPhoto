package database

import "fmt"

// Get likes of a media with a specific mediaid
func (db *appdbimpl) GetMediaLikes(mediaID string) ([]string, error) {
	// 1 - execute query
	rawLikes, err := db.c.Query(`SELECT userid FROM like WHERE mediaid = ?`, mediaID)
	if err != nil {
		// query returned error -> return error
		return nil, fmt.Errorf("error encountered while executing a select query: %w", err)
	}
	defer func() { _ = rawLikes.Close() }()

	// 2 - retrieve rows and build followers array result
	var ret []string
	for rawLikes.Next() {
		var l string
		err = rawLikes.Scan(&l)
		if err != nil {
			// scan returned error -> return error
			return nil, err
		}
		ret = append(ret, l)
	}
	if err = rawLikes.Err(); err != nil {
		return nil, err
	}
	// 3 - return result media array
	return ret, nil
}
