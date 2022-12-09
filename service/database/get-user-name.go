package database

// funzione che da userID -> username
func (db *appdbimpl) GetUserName(userid string) (string, error) {
	var name string
	err := db.c.QueryRow(`SELECT username FROM users WHERE userid =?`, userid).Scan(&name)
	if err != nil {
		return "", ErrUserProfileDoesNotExists
	}
	return name, nil
}
