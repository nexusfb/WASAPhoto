package database

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Create new comment with userid and mediaid
func (db *appdbimpl) CommentPhoto(userID string, mediaID string, comment CommentDB) (string, error) {
	// 1 - create new commentID
	rawCid, err := uuid.NewV4()
	if err != nil {
		// newV4 returned error -> return error
		return "", fmt.Errorf("error encountered while creating new commentID: %w", err)
	}
	cID := rawCid.String()

	// 2 - create timestamp
	now := time.Now().Format(time.RFC3339)

	// 3 - execute query
	_, err = db.c.Exec(`INSERT INTO comment (commentid, mediaid, authorid, date, content) VALUES (?,?,?,?,?,?,?)`,
		cID, mediaID, userID, now, comment.Content)
	if err != nil {
		// exec returned error -> return error
		return "", fmt.Errorf("error when creating new comment: %w", err)
	}

	// 4 - return mediaID
	return cID, nil
}
