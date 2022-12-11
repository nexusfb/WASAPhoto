package database

import "fmt"

// Get user who follow user with userid in path
func (db *appdbimpl) GetUserFollowers(userid string) ([]string, error) {
	// 1 - execute query
	rawMedia, err := db.c.Query(`SELECT followerid FROM follow WHERE followedid = ?`, userid)
	if err != nil {
		// query returned errror -> return error
		return nil, fmt.Errorf("error encountered while executing a select query: %w", err)
	}
	defer func() { _ = rawMedia.Close() }()

	// 2 - retrieve rows and build media array result
	var ret []string
	for rawMedia.Next() {
		var f string
		err = rawMedia.Scan(&f)
		if err != nil {
			// scan returned error -> return error
			return nil, err
		}
		ret = append(ret, f)
	}
	if err = rawMedia.Err(); err != nil {
		return nil, err
	}
	// 3 - return result media array
	return ret, nil
}
