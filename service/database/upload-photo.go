package database

import (
	"fmt"
	"time"
)

// Create new media with userid and media DB
func (db *appdbimpl) UploadPhoto(media MediaDB) error {
	// 1 - create timestamp
	now := time.Now().Format(time.RFC3339)

	// 2 - execute query
	_, err := db.c.Exec(`INSERT INTO media (mediaid, authorid, date, caption, photo) VALUES (?,?,?,?,?)`,
		media.MediaID, media.AuthorID, now, media.Caption, media.Photo)
	fmt.Println(media.MediaID)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error when creating new media: %w", err)
	}

	// 4 - return mediaID
	return nil
}
