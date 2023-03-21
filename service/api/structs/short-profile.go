package structs

import (
	"github.com/nexusfb/WASAPhoto/service/database"
)

// Short profile struct
type ShortProfile struct {
	// username of the owner of the profile
	Username string `json:"username"`
	// URL of the profile picture
	Pic string `json:"pic,omitempty"`
}

// Function to map a database profile to the struct profile
func (s *ShortProfile) FromDatabase(db database.AppDatabase, userID string) error {
	name, err := db.GetUserName(userID)
	if err != nil {
		return err
	}
	s.Username = name
	profile, err := db.GetUserProfile(name)
	if err != nil {
		return err
	}
	s.Pic = profile.ProfilePic
	return nil
}
