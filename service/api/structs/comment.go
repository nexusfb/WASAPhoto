package structs

import (
	"github.com/nexusfb/WASAPhoto/service/database"
)

var (
	ContentRx = BioRx // Content regex is a general string pattern
)

// Comment struct
type Comment struct {
	// unique identifier for the comment
	CommentID string `json:"commentid"`
	// username of the author of the comment
	// notice that author name is not stored in the database but it is needed here in order to display it
	AuthorName string `json:"author"`
	// date and time of the comment
	Date string `json:"date"`
	// body of the comment
	Content string `json:"content"`
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
	return database.CommentDB{
		Content: c.Content,
	}
}

// Function to check the validity of a comment api struct
func (c *Comment) IsValid() bool {
	return len(c.Content) <= 150 && ContentRx.MatchString(c.Content)

}
