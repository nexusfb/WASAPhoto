package structs

import (
	"github.com/nexusfb/WASAPhoto/service/database"
)

var (
	PhotoRx   = ProfilePicRx // Photo regex is a url pattern for a png/jpg/jpeg image
	CaptionRx = BioRx        // Caption regex is a general string pattern
)

type Media struct {
	// unique identifier of the media
	MediaID string `json:"id"`
	// unique identifier of the author of the media
	AuthorID string `json:"authorid"`
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
	// declare variables
	var err error
	// media id is the same as the database media id
	m.MediaID = media.MediaID
	//retrive author name using the author id
	m.AuthorName, err = db.GetUserName(media.AuthorID)
	if err != nil {
		return err
	}
	// author id is the same as the database author id
	m.AuthorID = media.AuthorID
	// retrive the profile using the author name
	profile, err := db.GetUserProfile(m.AuthorName)
	if err != nil {
		return err
	}
	// take the profile picture
	m.AuthorPic = profile.ProfilePic
	// date is the same as the database date
	m.Date = media.Date
	// caption is the same as the database caption
	m.Caption = media.Caption
	// photo is the same as the database photo
	m.Photo = media.Photo
	// retrive the number of likes counting the rows of the like table with that mediaid
	m.NLikes = db.CountRows("like", "mediaid", m.MediaID)
	// retrive the number of comments counting the rows of the comment table with that mediaid
	m.NComments = db.CountRows("comment", "mediaid", m.MediaID)
	// retrive the like of the logged user for that mediaid
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
