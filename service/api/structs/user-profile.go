package structs

import (
	"regexp"

	"github.com/nexusfb/WASAPhoto/service/database"
)

var (
	UsernameRx   = regexp.MustCompile(`^[a-zA-Z0-9._]{5,20}$`)
	ProfilePicRx = regexp.MustCompile(`^(https?:\/\/.*\.(?:png|jpg|jpeg))$`)
	BioRx        = regexp.MustCompile("^.*$")
)

type Profile struct {
	UserID     string `json:"userid"`
	Username   string `json:"username"`
	Bio        string `json:"bio,omitempty"`
	ProfilePic string `json:"profilepic,omitempty"`
	NMedia     uint32 `json:"nmedia,omitempty"`
	NFollowers uint32 `json:"nfollowers,omitempty"`
	NFollowing uint32 `json:"nfollowing,omitempty"`
}

type Username struct {
	Name string `json:"username"`
}

func (u Username) IsValid() bool {
	return len(u.Name) >= 5 && len(u.Name) <= 20 && UsernameRx.MatchString(u.Name)
}

func (p *Profile) FromDatabase(profile database.ProfileDB) {
	p.UserID = profile.UserID
	p.Username = profile.Username
	p.Bio = profile.Bio
	p.ProfilePic = profile.ProfilePic
	p.NMedia = profile.NMedia
	p.NFollowers = profile.NFollowers
	p.NFollowing = profile.NFollowing
}

func (profile *Profile) ToDatabase() database.ProfileDB {
	return database.ProfileDB{
		UserID:     profile.UserID,
		Username:   profile.Username,
		Bio:        profile.Bio,
		ProfilePic: profile.ProfilePic,
		NMedia:     profile.NMedia,
		NFollowers: profile.NFollowers,
		NFollowing: profile.NFollowing,
	}
}

func (p *Profile) IsValid() bool {
	return len(p.UserID) >= 5 && len(p.UserID) <= 20 && UsernameRx.MatchString(p.UserID) &&
		len(p.Username) >= 5 && len(p.Username) <= 20 && UsernameRx.MatchString(p.Username) &&
		p.NMedia >= 0 &&
		p.NFollowers >= 0 &&
		p.NFollowing >= 0 &&
		len(p.ProfilePic) >= 0 && len(p.ProfilePic) <= 200 && ProfilePicRx.MatchString(p.ProfilePic) &&
		len(p.Bio) >= 0 && len(p.Bio) <= 150 && BioRx.MatchString(p.Bio)

}
