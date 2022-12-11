package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nexusfb/WASAPhoto/service/database/databaseStructs"
)

var (
	// errors of userprofile struct
	ErrUserProfileAlreadyExists = errors.New("user already exists")
	ErrUserProfileDoesNotExists = errors.New("user does not exists")
)

/*
// 3 - database struct for follow
type FollowDB struct {
	FollowerID string
	FollowedID string
}

// 4 - database struct for likes
type LikeDB struct {
	MediaID  string
	AuthorID string
}
*/

/*
// 6 - database struct for bans
type BanDB struct {
	BannerID string
	BannedID string
}
*/

type AppDatabase interface {
	// takes username -> creates user profile -> returns userID
	DoLogin(username string) (string, error)

	// takes username -> returns user profile
	GetUserProfile(username string) (databaseStructs.UserProfileDB, error)

	// takes new user profile -> updates user profile -> returns new user profile
	UpdateUserProfile(profile UserProfileDB) (UserProfileDB, error)

	// takes userID -> deletes user profile -> returns //
	DeleteUserProfile(userid string) error

	// takes userID -> returns username
	GetUserName(userid string) (string, error)

	// takes username -> returns userID
	GetUserID(username string) (string, error)

	// takes userID -> creates media -> returns mediaID
	UploadPhoto(userid string, media MediaDB) (string, error)

	// takes mediaID -> deletes media -> returns //
	DeletePhoto(mediaid string) error

	// takes mediaID -> reuturns media
	GetMedia(mediaid string) (MediaDB, error)

	// takes userID -> returns array of media
	GetUserMedia(userid string) ([]MediaDB, error)

	// takes userID -> adds user to following list -> returns error
	FollowUser(userid string, followid string) error

	// takes userID -> deletes user from following list -> returns error
	UnfollowUser(userid string, followedid string) error

	// takes userID -> returns array of followers
	GetUserFollowers(userid string) ([]string, error)

	// takes userID -> returns array of followings
	GetUserFollowings(userid string) ([]string, error)

	// takes table, column, event -> returns the number of rows in the table for which column==event
	CountRows(table string, column string, event string) uint32

	LikePhoto(mediaid string, userid string) error
	UnlikePhoto(mediaid string, userid string) error
	GetMediaLikes(mediaid string) []string

	CommentPhoto(mediaid string, userid string) (string, error)
	UncommentPhoto(mediaid string, userid string) error
	GetMediaComments(mediaid string) ([]CommentDB, error)

	BanUser(bannerID string, bannedID string) error
	UnbanUser(bannerID string, bannedID string) error
	GetBannedUsers(bannerID string) ([]string, error)

	GetMyStream(userID string) ([]MediaDB, error)

	// default
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// DROP
	//f, _ := db.Exec(`DROP TABLE users IF EXISTS `)
	//f, _ = db.Exec(`DROP TABLE media IF EXISTS `)
	//fmt.Println(f)

	// 1 - table users
	var tableName string = "users"
	sqlStmt := `CREATE TABLE users (
		userid TEXT NOT NULL PRIMARY KEY,
		username TEXT NOT NULL,
		bio TEXT DEFAULT "" NOT NULL,
		profilepic TEXT DEFAULT "" NOT NULL,
		nmedia INTEGER DEFAULT 0 NOT NULL,
		nfollowers INTEGER DEFAULT 0 NOT NULL,
		nfollowing INTEGER DEFAULT 0 NOT NULL);`

	// 2 - create table users
	err := createTables(tableName, sqlStmt, db)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	// 3 - table media
	tableName = "media"
	sqlStmt = `CREATE TABLE media (
		mediaid INTEGER NOT NULL PRIMARY KEY,
		authorid TEXT NOT NULL,
		date TEXT DEFAULT "" NOT NULL,
		caption TEXT DEFAULT "" NOT NULL,
		photo TEXT DEFAULT "" NOT NULL,
		nlikes INTEGER DEFAULT 0 NOT NULL,
		ncomments INTEGER DEFAULT 0 NOT NULL,
		FOREIGN KEY (authorid) REFERENCES users(userid));`

	// 4 - create table media
	err = createTables(tableName, sqlStmt, db)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	// 5 - table media
	tableName = "follow"
	sqlStmt = `CREATE TABLE follow (
		followerid INTEGER NOT NULL,
		followedid TEXT NOT NULL,
		FOREIGN KEY (followerid) REFERENCES users(userid),
		FOREIGN KEY (followerid) REFERENCES users(userid),
		PRIMARY KEY (followerid,followingid));`

	// 6 - create table media
	err = createTables(tableName, sqlStmt, db)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

// function that given a table name -> creates table
func createTables(tableName string, sqlStmt string, db *sql.DB) error {
	var table string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='` + tableName + `';`).Scan(&table)
	fmt.Println(err)
	if errors.Is(err, sql.ErrNoRows) {
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return fmt.Errorf("error creating database: %w", err)
		}
	}
	return err
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
