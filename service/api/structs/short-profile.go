package structs

import (
	"github.com/nexusfb/WASAPhoto/service/database"
)

var (
	// Username regex is a case-sensitive alfanumeric string + (._)
	NameRx = UsernameRx
	// Pic regex is a url pattern for a png/jpg/jpeg image
	PicRx = ProfilePicRx
)

// Short profile struct represents a username and its
type ShortProfile struct {
	Username string `json:"username"`
	Pic      string `json:"profilepic,omitempty"`
}

// Function to map a database profile to the struct profile
func (s *ShortProfile) FromDatabase(db database.AppDatabase, name string) error {
	profile, err := db.GetUserProfile(name)
	if err != nil {
		return err
	}
	s.Username = profile.Username
	s.Pic = profile.ProfilePic
	return nil
}
