package database

func (db *appdbimpl) GetUserMedia(userid string) ([]MediaDB, error) {
	// 1 - query
	const query = `
	SELECT *
	FROM media
	WHERE userid = ?`

	var ret []MediaDB
	rawMedia, err := db.c.Query(query, userid)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rawMedia.Close() }()

	// 2 - costruisci array
	for rawMedia.Next() {
		var m MediaDB
		err = rawMedia.Scan(&m.MediaID, &m.AuthorID, &m.Date, &m.Caption, &m.Photo, &m.NLikes, &m.NComments)
		if err != nil {
			return nil, err
		}
		ret = append(ret, m)
	}
	if err = rawMedia.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
