package databaseStructs

// database struct for user profile
type UserProfileDB struct {
	UserID     string // identifier of the owner of the profile
	Username   string // username chosen by the owner of the profile
	Bio        string // biography chosen by the owner of the profile (i.e a short description)
	ProfilePic string // url of the profile picture
}
