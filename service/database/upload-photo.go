package database

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Create new media with userid and media DB
func (db *appdbimpl) UploadPhoto(userid string, media MediaDB) (string, error) {
	// 1 - create new mediaID
	rawMid, err := uuid.NewV4()
	if err != nil {
		// newV4 returned error -> return error
		return "", fmt.Errorf("error encountered while creating new mediaID: %w", err)
	}
	Mid := rawMid.String()

	// 2 - create timestamp
	media.Date = time.Now().Format(time.RFC3339)

	// 3 - execute query
	_, err = db.c.Exec(`INSERT INTO media (mediaid, authorid, date, caption, photo) VALUES (?,?,?,?,?,?,?)`,
		Mid, userid, media.Date, media.Caption, media.Photo)
	if err != nil {
		// exec returned error -> return error
		return "", fmt.Errorf("error when creating new media: %w", err)
	}

	// 4 - return mediaID
	return Mid, nil
}
