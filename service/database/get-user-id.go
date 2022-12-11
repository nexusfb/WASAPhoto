package database

// Get username with userid
func (db *appdbimpl) GetUserID(username string) (string, error) {
	// 1 - execute query
	var name string
	err := db.c.QueryRow(`SELECT userid FROM user WHERE username =?`, username).Scan(&name)
	if err != nil {
		// queryrow returned error which means no user exists -> return error
		return "", ErrUserProfileDoesNotExists
	}

	// 2 - return result userID
	return name, nil
}
