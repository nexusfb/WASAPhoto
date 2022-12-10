package database

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Create new media with userid in path and media DB in request body
func (db *appdbimpl) UploadPhoto(userid string, media MediaDB) (string, error) {
	// 1 - create new mediaID
	rawMid, err := uuid.NewV1()
	if err != nil {
		// newV4 returned error -> return error
		return "00000000", fmt.Errorf("error encountered while creating new mediaID: %w", err)
	}
	Mid := rawMid.String()

	// 2 - create timestamp
	now := time.Now().Format(time.RFC3339)
	media.Date = now

	// 3 - execute query
	_, err = db.c.Exec(`INSERT INTO media (mediaid, authorid, date, caption, photo, nlikes, ncomments ) VALUES (?,?,?,?,?,?,?)`,
		Mid, userid, media.Date, media.Caption, media.Photo, 0, 0)
	if err != nil {
		// exec returned error -> return error
		return "00000000", fmt.Errorf("error when creating new media: %w", err)
	}

	// 4 - return mediaID
	return Mid, nil
}
