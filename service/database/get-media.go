package database

import (
	"fmt"
)

func (db *appdbimpl) GetMedia(mediaid string) (MediaDB, error) {
	const query = `
	SELECT *
	FROM media
	WHERE mediaid = ?`

	var ret MediaDB
	rawMedia, err := db.c.Query(query, mediaid)
	if err != nil {
		return MediaDB{}, err
	}
	defer func() { _ = rawMedia.Close() }()
	result := rawMedia.Next()
	fmt.Println(result)
	if !result {
		return MediaDB{}, ErrMediaDoesNotExists
	}

	err = rawMedia.Scan(&ret.MediaID, &ret.AuthorID, &ret.Date, &ret.Caption, &ret.Photo, &ret.NLikes, &ret.NComments)
	if err != nil {
		return MediaDB{}, fmt.Errorf("error encountered while scanning media : %w", err)
	}
	return ret, nil

}
