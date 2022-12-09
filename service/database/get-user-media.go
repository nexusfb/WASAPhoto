package database

import "fmt"

// Get user array media with userid in path
func (db *appdbimpl) GetUserMedia(userid string) ([]MediaDB, error) {
	// 1 - execute query
	rawMedia, err := db.c.Query(`SELECT * FROM media WHERE authorid = ?`, userid)
	if err != nil {
		// query returned errror -> return error
		return nil, fmt.Errorf("error encountered while executing a select query: %w", err)
	}
	defer func() { _ = rawMedia.Close() }()

	// 2 - retrieve rows and build media array result
	var ret []MediaDB
	for rawMedia.Next() {
		var m MediaDB
		err = rawMedia.Scan(&m.MediaID, &m.AuthorID, &m.Date, &m.Caption, &m.Photo, &m.NLikes, &m.NComments)
		if err != nil {
			// scan returned error -> return error
			return nil, err
		}
		ret = append(ret, m)
	}
	if err = rawMedia.Err(); err != nil {
		return nil, err
	}
	// 3 - return result media array
	return ret, nil
}
