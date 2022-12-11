package structs

import (
	"regexp"

	"github.com/nexusfb/WASAPhoto/service/database"
)

var (
	// Username regex is a case-sensitive alfanumeric string + (._)
	UsernameRx = regexp.MustCompile(`^[a-zA-Z0-9._]{5,20}$`)
	// ProfilePic regex is a url pattern for a png/jpg/jpeg image
	ProfilePicRx = regexp.MustCompile(`^(https?:\/\/.*\.(?:png|jpg|jpeg))$`)
	// Bio regex is a general string pattern
	BioRx = regexp.MustCompile("^.*$")
)

// Profile struct represents a Profile with unique user ID, username, bio, url of the profile image
type UserProfile struct {
	UserID     string `json:"userid"`
	Username   string `json:"username"`
	Bio        string `json:"bio,omitempty"`
	ProfilePic string `json:"profilepic,omitempty"`
	NMedia     uint32 `json:"nmedia,omitempty"`
	NFollowers uint32 `json:"nfollowers,omitempty"`
	NFollowing uint32 `json:"nfollowing,omitempty"`
}

// Function to map a database profile to the struct profile
func (p *UserProfile) FromDatabase(profile database.UserProfileDB, db database.AppDatabase) {
	p.UserID = profile.UserID
	p.Username = profile.Username
	p.Bio = profile.Bio
	p.ProfilePic = profile.ProfilePic
	p.NMedia = db.CountRows("media", "authorid", p.UserID)
	p.NFollowers = db.CountRows("follow", "followedid", p.UserID)
	p.NFollowing = db.CountRows("follow", "followerid", p.UserID)
}

// Function to map the struct profile to a database profile
func (profile *UserProfile) ToDatabase() database.UserProfileDB {
	return database.UserProfileDB{
		UserID:     profile.UserID,
		Username:   profile.Username,
		Bio:        profile.Bio,
		ProfilePic: profile.ProfilePic,
	}
}

// Function to check if a media struct is valid
func (p *UserProfile) IsValid() bool {
	return len(p.UserID) == 27 &&
		len(p.Username) >= 5 && len(p.Username) <= 20 && UsernameRx.MatchString(p.Username) && len(p.ProfilePic) <= 200 && ProfilePicRx.MatchString(p.ProfilePic) &&
		len(p.Bio) <= 150 && BioRx.MatchString(p.Bio)
}

// Username structs represents a username
type Username struct {
	Name string `json:"username"`
}

// Function to check if a username is valid
func (u Username) IsValid() bool {
	return len(u.Name) >= 5 && len(u.Name) <= 20 && UsernameRx.MatchString(u.Name)
}
