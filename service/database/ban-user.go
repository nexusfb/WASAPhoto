package database

import (
	"fmt"
)

// Add a ban from bannerid user to bannedid user
func (db *appdbimpl) BanUser(bannerID string, bannedID string) error {
	// 1 - check if user already banned other user
	r, err := db.c.Query(`SELECT * FROM ban WHERE bannerid = ? AND bannedid = ?`, bannerID, bannedID)
	if err != nil {
		return err
	}
	if !r.Next() == false {
		// should never happpen since it is checked in the API
		return ErrUserAlreadyFollowed
	}
	if err = r.Err(); err != nil {
		return err
	}
	defer func() { _ = r.Close() }()
	// here only if user has not banned the other user yet

	// 2 - check if user follows the user he wants to ban
	if db.Check("follow", "followerid", "followedid", bannerID, bannedID) {
		// banner should unfollow the other user first
		err = db.UnfollowUser(bannerID, bannedID)
		if err != nil {
			return err
		}
	}

	// 3 - check if banned follows the logged user
	if db.Check("follow", "followerid", "followedid", bannedID, bannerID) {
		// banner should unfollow the other user first
		err = db.UnfollowUser(bannedID, bannerID)
		if err != nil {
			return err
		}
	}

	// 4 - create new ban
	_, err = db.c.Exec(`INSERT INTO ban (bannerid, bannedid) VALUES (?,?);`, bannerID, bannedID)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error when creating new ban: %w", err)
	}

	// 5 - delete all likes of banned user from logged user
	query :=
		`DELETE FROM like WHERE mediaid in (SELECT like.mediaid FROM like JOIN media WHERE media.authorid = ? and like.userid = ?)
		AND userid in (SELECT like.userid FROM like JOIN media WHERE media.authorid = ? and like.userid = ?)`

	_, err = db.c.Exec(query, bannerID, bannedID, bannerID, bannedID)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error encountered while executing a delete query: %w", err)
	}

	// 6 - delete all comments of banned user from logged user
	query =
		`DELETE FROM comment WHERE mediaid in (SELECT comment.mediaid FROM comment  JOIN media WHERE media.authorid = ? and comment.userid = ?)
		AND userid in (SELECT comment.userid FROM comment JOIN media WHERE media.authorid = ? and comment.userid = ?)`

	_, err = db.c.Exec(query, bannerID, bannedID, bannerID, bannedID)
	if err != nil {
		// exec returned error -> return error
		return fmt.Errorf("error encountered while executing a delete query: %w", err)
	}

	// 7 - return nil error
	return nil
}
