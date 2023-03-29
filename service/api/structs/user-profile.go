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
	// ID of the owner of the profile
	UserID string `json:"userid"`
	// username of the owner of the profile
	Username string `json:"username"`
	// short text content
	Bio string `json:"bio"`
	// url of the profile picture
	ProfilePic string `json:"profilepic"`
	// number of media published by the owner of the profile
	NMedia uint32 `json:"nmedia"`
	// number of followers of the owner of the profile
	NFollowers uint32 `json:"nfollowers"`
	// number of followings of the owner of the profile
	NFollowing uint32 `json:"nfollowing"`
	// true if the authenticated user follows this userid
	Followed bool `json:"followed"`
	// true if the authenticated user banned this userid
	Banned bool `json:"banned"`
}

// Function to map a database user profile to the api struct user profile
func (p *UserProfile) FromDatabase(profile database.UserProfileDB, db database.AppDatabase, token string) {
	// userid is the same as database userid
	p.UserID = profile.UserID
	// username is the same as database username
	p.Username = profile.Username
	// bio is the same as database bio
	p.Bio = profile.Bio
	// profile pic is the same as database profile pic
	p.ProfilePic = profile.ProfilePic
	// the number of media of a user is computed counting the rows of the media table with this authorid
	p.NMedia = db.CountRows("media", "authorid", p.UserID)
	// the number of followers of a user is computed counting the rows of the follow table with this authorid
	p.NFollowers = db.CountRows("follow", "followedid", p.UserID)
	// the number of followings of a user is computed counting the rows of the follow table with this authorid
	p.NFollowing = db.CountRows("follow", "followerid", p.UserID)
	// retrive the follow of the authenticated user
	p.Followed = db.Check("follow", "followerid", "followedid", token, profile.UserID)
	// retrive the ban of the authenticated user
	p.Banned = db.Check("ban", "bannerid", "bannedid", token, profile.UserID)
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
	return len(p.Username) >= 5 && len(p.Username) <= 20 && UsernameRx.MatchString(p.Username) && len(p.ProfilePic) <= 200 && ProfilePicRx.MatchString(p.ProfilePic) &&
		len(p.Bio) <= 150 && BioRx.MatchString(p.Bio)
}

// Username structs represents a username
type Username struct {
	Name string `json:"Username"`
}

// Function to check if a username is valid
func (u Username) IsValid() bool {
	return len(u.Name) >= 5 && len(u.Name) <= 20 && UsernameRx.MatchString(u.Name)
}
