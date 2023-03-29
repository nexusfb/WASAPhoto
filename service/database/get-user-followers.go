package database

import "fmt"

// Get users who follow user with userid
func (db *appdbimpl) GetUserFollowers(userid string) ([]string, error) {
	// 1 - execute query
	rawFollowers, err := db.c.Query(`SELECT followerid FROM follow WHERE followedid = ?`, userid)
	if err != nil {
		// query returned error -> return error
		return nil, fmt.Errorf("error encountered while executing a select query: %w", err)
	}
	defer func() { _ = rawFollowers.Close() }()

	// 2 - retrieve rows and build followers array result
	var ret []string
	for rawFollowers.Next() {
		var f string
		err = rawFollowers.Scan(&f)
		if err != nil {
			// scan returned error -> return error
			return nil, err
		}
		ret = append(ret, f)
	}
	if err = rawFollowers.Err(); err != nil {
		return nil, err
	}

	// 3 - return result media array
	return ret, nil
}
