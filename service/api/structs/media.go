package structs

import (
	"regexp"

	"github.com/nexusfb/WASAPhoto/service/database"
)

var (
	PhotoRx   = regexp.MustCompile(`^(https?:\/\/.*\.(?:png|jpg|jpeg))$`)
	CaptionRx = regexp.MustCompile("^.*$")
)

type Media struct {
	MediaID   string `json:"mediaid"`
	AuthorID  string `json:"authorid"`
	Date      string `json:"date,omitempty"`
	Caption   string `json:"caption,omitempty"`
	Photo     string `json:"photo,omitempty"`
	NLikes    uint32 `json:"nlikes,omitempty"`
	NComments uint32 `json:"ncomments,omitempty"`
}

func (m *Media) FromDatabase(media database.MediaDB) {
	m.MediaID = media.MediaID
	m.AuthorID = media.AuthorID
	m.Date = media.Date
	m.Caption = media.Caption
	m.Photo = media.Photo
	m.NLikes = media.NLikes
	m.NComments = media.NComments
}

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

func (m *Media) IsValid() bool {
	return len(m.MediaID) >= 5 && len(m.MediaID) <= 20 &&
		len(m.AuthorID) >= 5 && len(m.AuthorID) <= 20 && PhotoRx.MatchString(m.Photo) &&
		m.NComments >= 0 &&
		m.NLikes >= 0 &&
		len(m.Photo) >= 0 && len(m.Photo) <= 200 && CaptionRx.MatchString(m.Caption) &&
		len(m.Caption) >= 0 && len(m.Caption) <= 150

}
