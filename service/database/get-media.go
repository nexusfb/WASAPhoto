package database

import (
	"fmt"
)

// Get media with mediaid
func (db *appdbimpl) GetMedia(mediaid string) (MediaDB, error) {
	// 1 - execute query
	rawMedia, err := db.c.Query(`SELECT * FROM media WHERE mediaid = ?`, mediaid)
	if err != nil {
		// query returned error -> return error
		return MediaDB{}, fmt.Errorf("error encountered while executing a select query: %w", err)
	}
	defer func() { _ = rawMedia.Close() }()

	// 2 - retrieve rows
	result := rawMedia.Next()
	if !result {
		// next() returned error which means there is no row (media does not exist) -> return error
		// should never happen -> existence of the media is already checked in the API
		return MediaDB{}, ErrMediaDoesNotExists
	}
	if err = rawMedia.Err(); err != nil {
		return MediaDB{}, err
	}

	// 3 - create media result
	var ret MediaDB
	err = rawMedia.Scan(&ret.MediaID, &ret.AuthorID, &ret.Date, &ret.Caption, &ret.Photo)
	if err != nil {
		// scan returned error -> return error
		return MediaDB{}, fmt.Errorf("error encountered while scanning media : %w", err)
	}

	// 4 - return media result
	return ret, nil
}
