package structs

// IMPORTS
import (
	"regexp"

	"github.com/nexusfb/WASAPhoto/service/database"
)

// REGEX VARIABLES
var (
	UsernameRx   = regexp.MustCompile(`^[a-zA-Z0-9._]{5,20}$`)               // Username regex is a case-sensitive alfanumeric string + (._)
	ProfilePicRx = regexp.MustCompile(`^(https?:\/\/.*\.(?:png|jpg|jpeg))$`) // ProfilePic regex is a url pattern for a png/jpg/jpeg image
	BioRx        = regexp.MustCompile("^.*$")                                // Bio regex is a general string pattern
)

// User profile struct
type UserProfile struct {
	UserID     string `json:"userid"`               // ID of the owner of the profile
	Username   string `json:"username"`             // username of the owner of the profile
	Bio        string `json:"bio,omitempty"`        // short text content
	ProfilePic string `json:"profilepic,omitempty"` // url of the profile picture
	NMedia     uint32 `json:"nmedia,omitempty"`     // number of media published by the owner of the profile
	NFollowers uint32 `json:"nfollowers,omitempty"` // number of followers of the owner of the profile
	NFollowing uint32 `json:"nfollowing,omitempty"` // number of followings of the owner of the profile
}

// Function to map a database user profile to the api struct user profile
func (p *UserProfile) FromDatabase(profile database.UserProfileDB, db database.AppDatabase) {
	p.UserID = profile.UserID
	p.Username = profile.Username
	p.Bio = profile.Bio
	p.ProfilePic = profile.ProfilePic
	p.NMedia = db.CountRows("media", "authorid", p.UserID)
	p.NFollowers = db.CountRows("follow", "followedid", p.UserID)
	p.NFollowing = db.CountRows("follow", "followerid", p.UserID)
}

// Function to map the api struct user profile to a database user profile
func (profile *UserProfile) ToDatabase() database.UserProfileDB {
	return database.UserProfileDB{
		UserID:     profile.UserID,
		Username:   profile.Username,
		Bio:        profile.Bio,
		ProfilePic: profile.ProfilePic,
	}
}

// Function to check if a user profile struct is valid
func (p *UserProfile) IsValid() bool {
	return len(p.UserID) == 27 && len(p.Username) >= 5 && len(p.Username) <= 20 && UsernameRx.MatchString(p.Username) && len(p.ProfilePic) <= 200 && ProfilePicRx.MatchString(p.ProfilePic) &&
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
