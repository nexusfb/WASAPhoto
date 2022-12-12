package structs

import (
	"github.com/nexusfb/WASAPhoto/service/database"
)

var (
	NameRx = UsernameRx   // Username regex is a case-sensitive alfanumeric string + (._)
	PicRx  = ProfilePicRx // Pic regex is a url pattern for a png/jpg/jpeg image
)

// Short profile struct represents a username and its
type ShortProfile struct {
	Username string // username of the owner of the profile
	Pic      string // URL of the profile picture
}

// Function to map a database profile to the struct profile
func (s *ShortProfile) FromDatabase(db database.AppDatabase, userID string) error {
	// 1 - get username
	name, err := db.GetUserName(userID)
	if err != nil {
		// get username returned error -> return error
		return err
	}
	s.Username = name

	// 2 - get profile picture
	profile, err := db.GetUserProfile(name)
	if err != nil {
		// get user profile returned error -> return error
		return err
	}
	s.Pic = profile.ProfilePic
	return nil
}
