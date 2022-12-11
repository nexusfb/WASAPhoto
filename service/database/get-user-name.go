package database

// Get username with userid
func (db *appdbimpl) GetUserName(userid string) (string, error) {
	// 1 - execute query
	var name string
	err := db.c.QueryRow(`SELECT username FROM user WHERE userid =?`, userid).Scan(&name)
	if err != nil {
		// queryrow returned error which means no user exists -> return error
		return "", ErrUserProfileDoesNotExists
	}

	// 2 - return result username
	return name, nil
}
