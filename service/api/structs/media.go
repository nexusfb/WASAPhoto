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
	MediaID    string
	AuthorID   string
	AuthorName string // notice that author name was not stored in the database struct of media but it is needed here in order to display it
	Date       string
	Caption    string `json:"caption,omitempty"`
	Photo      string `json:"photo,omitempty"`
	NLikes     uint32
	NComments  uint32
	liked      bool
}

// Function to map a database media to the media struct
func (m *Media) FromDatabase(media database.MediaDB, db database.AppDatabase, token string) error {
	var err error
	m.MediaID = media.MediaID
	m.AuthorID = media.AuthorID
	m.AuthorName, err = db.GetUserName(m.AuthorID)
	if err != nil {
		return err
	}
	m.Date = media.Date
	m.Caption = media.Caption
	m.Photo = media.Photo
	m.NLikes = db.CountRows("like", "mediaid", m.MediaID)
	m.NComments = db.CountRows("comment", "mediaid", m.MediaID)
	m.liked = db.Check("like", "mediaid", "userid", m.MediaID, token)
	return nil
}

// Function to map a Media struct to a database media
func (media *Media) ToDatabase() database.MediaDB {
	return database.MediaDB{
		Caption: media.Caption,
		Photo:   media.Photo,
	}
}

// Function to check the validity of a Media struct
func (m *Media) IsValid() bool {
	return PhotoRx.MatchString(m.Photo) && len(m.Photo) <= 200 &&
		CaptionRx.MatchString(m.Caption) && len(m.Caption) <= 150

}
