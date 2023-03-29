package database

import "fmt"

// Get user list given username
func (db *appdbimpl) GetUserList(username string) ([]string, error) {
	// 1 - execute query
	// next level: users, err := db.c.Query(`SELECT username FROM user WHERE username LIKE "` + username + `%"`)
	users, err := db.c.Query(`SELECT userid FROM user`)
	if err != nil {
		// query returned error -> return error
		return nil, fmt.Errorf("error encountered while executing a select query: %w", err)
	}
	defer func() { _ = users.Close() }()

	// 2 - retrieve rows and build users array result
	var ret []string
	for users.Next() {
		var f string
		err = users.Scan(&f)
		if err != nil {
			// scan returned error -> return error
			return nil, err
		}
		ret = append(ret, f)
	}
	if err = users.Err(); err != nil {
		return nil, fmt.Errorf("error encountered while iterating on rows: %w", err)
	}

	// 3 - return result users array
	return ret, nil
}
