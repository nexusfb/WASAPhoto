package database

import "fmt"

// Get users who are banned by logged user
func (db *appdbimpl) GetBannedUsers(userid string) ([]string, error) {
	// 1 - execute query
	rawBanned, err := db.c.Query(`SELECT bannedid FROM ban WHERE bannerid = ?`, userid)
	if err != nil {
		// query returned error -> return error
		return nil, fmt.Errorf("error encountered while executing a select query: %w", err)
	}
	defer func() { _ = rawBanned.Close() }()

	// 2 - retrieve rows and build banned array result
	var ret []string
	for rawBanned.Next() {
		var f string
		err = rawBanned.Scan(&f)
		if err != nil {
			// scan returned error -> return error
			return nil, err
		}
		ret = append(ret, f)
	}
	if err = rawBanned.Err(); err != nil {
		return nil, err
	}

	// 3 - return result media array
	return ret, nil
}
