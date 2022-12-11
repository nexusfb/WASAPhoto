package database

import "fmt"

// Get user who are followed by user with userid
func (db *appdbimpl) GetUserFollowings(userid string) ([]string, error) {
	// 1 - execute query
	rawFollowings, err := db.c.Query(`SELECT followedid FROM follow WHERE followerid = ?`, userid)
	if err != nil {
		// query returned error -> return error
		return nil, fmt.Errorf("error encountered while executing a select query: %w", err)
	}
	defer func() { _ = rawFollowings.Close() }()

	// 2 - retrieve rows and build followings array result
	var ret []string
	for rawFollowings.Next() {
		var f string
		err = rawFollowings.Scan(&f)
		if err != nil {
			// scan returned error -> return error
			return nil, err
		}
		ret = append(ret, f)
	}
	if err = rawFollowings.Err(); err != nil {
		return nil, fmt.Errorf("error encountered while iterating on rows: %w", err)
	}
	// 3 - return result followings array
	return ret, nil
}
