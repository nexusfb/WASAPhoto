package api

// function to check if logged user corresponds to given userid (i.e if user is interacting with his own profile)
func authentication(token string, userid string) bool {
	return (token == userid)

}
