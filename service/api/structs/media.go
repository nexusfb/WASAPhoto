package structs

import (
	"github.com/nexusfb/WASAPhoto/service/database"
)

var (
	PhotoRx   = ProfilePicRx // Photo regex is a url pattern for a png/jpg/jpeg image
	CaptionRx = BioRx        // Caption regex is a general string pattern
)

// Media struct
type Media struct {
	// unique identifier of the media
	MediaID string `json:"id"`
	// unique identifier of the author of the media
	AuthorID string
	// username of the author of the media
	// notice that author name was not stored in the database struct of media but it is needed here in order to display it
	AuthorName string `json:"author"`
	// profile picture of the author
	AuthorPic string `json:"authorpic"`
	// date and time of the media
	Date string `json:"date"`
	// description of the media
	Caption string `json:"caption"`
	// URL of the photo in the media
	Photo string `json:"photo"`
	// likes count
	NLikes uint32 `json:"nlikes"`
	// comments count
	NComments uint32 `json:"ncomments"`
	// true if logged user has liked the media or not
	Liked bool `json:"liked"`
}

// Function to map a database media to the media struct
func (m *Media) FromDatabase(media database.MediaDB, db database.AppDatabase, token string) error {
	var err error
	m.MediaID = media.MediaID
	m.AuthorName, err = db.GetUserName(media.AuthorID)
	if err != nil {
		return err
	}
	profile, err := db.GetUserProfile(m.AuthorName)
	if err != nil {
		return err
	}
	m.AuthorPic = profile.ProfilePic
	m.Date = media.Date
	m.Caption = media.Caption
	m.Photo = media.Photo
	m.NLikes = db.CountRows("like", "mediaid", m.MediaID)
	m.NComments = db.CountRows("comment", "mediaid", m.MediaID)
	m.Liked = db.Check("like", "mediaid", "userid", m.MediaID, token)
	return nil
}

// Function to map a Media struct to a database media
func (media *Media) ToDatabase(db database.AppDatabase) database.MediaDB {
	return database.MediaDB{
		AuthorID: media.AuthorID,
		MediaID:  media.MediaID,
		Caption:  media.Caption,
		Photo:    media.Photo,
	}
}

// Function to check the validity of a Media struct
func (m *Media) IsValid() bool {
	return PhotoRx.MatchString(m.Photo) && len(m.Photo) <= 200 &&
		CaptionRx.MatchString(m.Caption) && len(m.Caption) <= 150

}
