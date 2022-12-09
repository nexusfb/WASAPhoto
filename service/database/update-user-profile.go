package database

import (
	"fmt"
)

func (db *appdbimpl) UpdateUserProfile(p ProfileDB) (ProfileDB, error) {

	// prendere profilo vecchio
	// 1 - id vecchio (uguale a nuovo)
	id := p.UserID
	// 2- get con userID
	const query = `
	SELECT username, bio, profilepic, nmedia, nfollows, nfollowing
	FROM users
	WHERE uid = ?`
	var old ProfileDB
	user, _ := db.c.Query(query, id)
	defer func() { _ = user.Close() }()
	err := user.Scan(&old.Username, &old.Bio, &old.ProfilePic, &old.NMedia, &old.NFollowers, &old.NFollowing)
	if err != nil {
		return ProfileDB{}, fmt.Errorf("error encountered while scanning user profile: %w", err)
	}
	// a questo punto il profilo vecchio sta dentro old

	// update profilo
	_, err = db.c.Exec(`UPDATE users SET username=?, bio=?, profilepic=?, WHERE uid=?`,
		p.Username, p.Bio, p.ProfilePic, p.UserID)
	if err != nil {
		return old, err
	}

	// se no errori e cambiato -> ritorna nuovo
	return p, nil
}
