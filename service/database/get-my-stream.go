package database

import "fmt"

// Get array of media of users followed by user with a specific userid
func (db *appdbimpl) GetMyStream(userid string) ([]MediaDB, error) {
	// 1 - execute query
	query := `
	SELECT media.mediaid, media.authorid, media.date, media.caption, follow.followerid
	FROM media 
	INNER JOIN follow on media.authorid=follow.followedid 
	AND follow.followerid = ?
	ORDER BY media.date DESC;`

	rawMedia, err := db.c.Query(query, userid)
	if err != nil {
		// query returned errror -> return error
		return nil, fmt.Errorf("error encountered while executing a select query: %w", err)
	}
	defer func() { _ = rawMedia.Close() }()

	// 2 - retrieve rows and build media array result
	var ret []MediaDB
	for rawMedia.Next() {
		var m MediaDB
		err = rawMedia.Scan(&m.MediaID, &m.AuthorID, &m.Date, &m.Caption, &m.Photo)
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
