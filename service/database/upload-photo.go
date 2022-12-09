package database

import (
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

func (db *appdbimpl) UploadPhoto(userid string, media MediaDB) (string, error) {
	//crea mediaID
	rawMid := ksuid.New()
	Mid := rawMid.String()
	//metti dentro media il timestamp
	now := time.Now().Format(time.RFC3339)
	media.Date = now
	//metti media mel DB
	_, err := db.c.Exec(`INSERT INTO media (mediaid, authorid, date, caption, photo, nlikes, ncomments ) VALUES (?,?,?,?,?,?,?)`,
		Mid, userid, media.Date, media.Caption, media.Photo, 0, 0)
	if err != nil {
		return "00000000", fmt.Errorf("error when creating new media: %w", err)
	}
	return Mid, nil
}
