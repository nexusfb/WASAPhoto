package structs

import (
	"github.com/nexusfb/WASAPhoto/service/database"
)

var (
	ContentRx = BioRx      // Content regex is a general string pattern
	AuthorRx  = UsernameRx // Author regex is a case-sensitive alfanumeric string + (._)
)

// Comment struct
type Comment struct {
	CommentID string `json:"commentid"`
	// notice that author name was not stored in the database struct of media but it is needed here in order to display it
	MediaID    string `json:"mediaid"`
	AuthorName string `json:"authorname"`
	Date       string `json:"date,omitempty"`
	Content    string `json:"content,omitempty"`
}

// Function to map a database comment to the comment api struct
func (c *Comment) FromDatabase(comment database.CommentDB, db database.AppDatabase) error {
	var err error
	c.CommentID = comment.CommentID
	c.AuthorName, err = db.GetUserName(comment.AuthorID)
	if err != nil {
		return err
	}
	c.Date = comment.Date
	c.Content = comment.Content
	return nil
}

// Function to map a comment api struct to a database comment struct
func (c *Comment) ToDatabase(db database.AppDatabase) database.CommentDB {
	id, err := db.GetUserID(c.AuthorName)
	if err != nil {
		return database.CommentDB{}
	}
	return database.CommentDB{
		MediaID:  c.MediaID,
		AuthorID: id,
		Content:  c.Content,
	}
}

// Function to check the validity of a comment api struct
func (c *Comment) IsValid() bool {
	return len(c.CommentID) == 27 && len(c.AuthorName) >= 5 && len(c.AuthorName) <= 20 &&
		AuthorRx.MatchString(c.AuthorName) && len(c.Content) <= 150 &&
		ContentRx.MatchString(c.Content)

}
