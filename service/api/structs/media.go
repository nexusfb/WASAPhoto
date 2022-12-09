package structs

import (
	"github.com/nexusfb/WASAPhoto/service/database"
)

var (
	// Photo regex is a url pattern for a png/jpg/jpeg image
	PhotoRx = ProfilePicRx
	// Caption regex is a general string pattern
	CaptionRx = BioRx
)

// Media struct represents a Photo with unique ID, authorID, date of creation, caption, url of the photo
type Media struct {
	MediaID   string `json:"mediaid"`
	AuthorID  string `json:"authorid"`
	Date      string `json:"date,omitempty"`
	Caption   string `json:"caption,omitempty"`
	Photo     string `json:"photo,omitempty"`
	NLikes    uint32 `json:"nlikes,omitempty"`
	NComments uint32 `json:"ncomments,omitempty"`
}

// Function to map a database media to the media struct
func (m *Media) FromDatabase(media database.MediaDB) {
	m.MediaID = media.MediaID
	m.AuthorID = media.AuthorID
	m.Date = media.Date
	m.Caption = media.Caption
	m.Photo = media.Photo
	m.NLikes = media.NLikes
	m.NComments = media.NComments
}

// Function to map a Media struct to a database media
func (media *Media) ToDatabase() database.MediaDB {
	return database.MediaDB{
		MediaID:   media.MediaID,
		AuthorID:  media.AuthorID,
		Date:      media.Date,
		Caption:   media.Caption,
		Photo:     media.Photo,
		NLikes:    media.NLikes,
		NComments: media.NComments,
	}
}

// Function to check the validity of a Media struct
func (m *Media) IsValid() bool {
	return len(m.MediaID) == 27 && len(m.AuthorID) == 27 &&
		PhotoRx.MatchString(m.Photo) && len(m.Photo) >= 0 && len(m.Photo) <= 200 &&
		CaptionRx.MatchString(m.Caption) && len(m.Caption) >= 0 && len(m.Caption) <= 150

}
