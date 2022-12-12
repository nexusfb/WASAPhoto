package database

// function to check if user/media exists or not
func (db *appdbimpl) ExistenceCheck(id string, what string) bool {
	// user case
	if what == "user" {
		// 1 - get username of the user
		_, err := db.GetUserName(id)
		if err == ErrUserProfileDoesNotExists {
			// err user does not exist -> return false
			return false
		}
		if err == nil {
			// user exists -> return true
			return true
		}
		// 2 - error occurred when looking for username -> return false ( no possibly harmful operation is performed)
		return false
	}
	// media case
	if what == "media" {
		// 1 - get media
		_, err := db.GetMedia(id)
		if err == ErrMediaDoesNotExists {
			// err media does not exist -> return false
			return false
		}
		if err == nil {
			// media exists -> return true
			return true
		}

		// 2 - error occurred when looking for username -> return false ( no possibly harmful operation is performed)
		return false
	}

	// never used
	return false

}
