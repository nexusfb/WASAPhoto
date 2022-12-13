package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// Add a ban from bannerid user to bannedid user
func (db *appdbimpl) BanUser(bannerID string, bannedID string) error {

	// 1 - check if user already banned
	// should never happen since it is already checked in the API
	err := db.c.QueryRow(`SELECT * FROM ban WHERE bannerid = ? AND bannedid = ?`, bannerID, bannedID).Scan()
	if !errors.Is(err, sql.ErrNoRows) {
		return ErrUserAlreadyBanned
	}

	// here only if user has not banned the other user yet

	// 2 - create new ban
	_, err = db.c.Exec(`INSERT INTO ban (bannerid, bannedid) VALUES (?,?)`, bannerID, bannedID)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error when creating new follow: %w", err)
	}

	// 3 - return nil error
	return nil
}
