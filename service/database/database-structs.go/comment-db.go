package databaseStructs

// database struct of a comment
type CommentDB struct {
	MediaID  string // identifier of the media where the comment was created
	AuthorID string // identifier of the author of the comment
	Date     string // date of creation of the comment
	Content  string // content of the comment (i.e short text content)
}
