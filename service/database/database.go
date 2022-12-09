package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	// errors of userprofile struct
	ErrUserProfileAlreadyExists = errors.New("user already exists")
	ErrUserProfileDoesNotExists = errors.New("user does not exists")

	//errors of media struct
	ErrMediaDoesNotExists = errors.New("media does not exists")
)

// STRUCTS
// 1 - database struct for user profile
type UserProfileDB struct {
	UserID     string
	Username   string
	Bio        string
	ProfilePic string
	NMedia     uint32
	NFollowers uint32
	NFollowing uint32
}

// 2 - database struct for media
type MediaDB struct {
	MediaID   string
	AuthorID  string
	Date      string
	Caption   string
	Photo     string
	NLikes    uint32
	NComments uint32
}

type AppDatabase interface {
	// 1) take username -> create user profile -> return userID
	DoLogin(username string) (string, error)

	// 2) take username -> return user profile
	GetUserProfile(username string) (UserProfileDB, error)

	// 3) take new user profile -> update user profile -> return new user profile
	UpdateUserProfile(profile UserProfileDB) (UserProfileDB, error)

	// 5) take userID -> delete user profile -> return //
	DeleteUserProfile(userid string) error

	// 6) take userID -> return username
	GetUserName(userid string) (string, error)

	// 7) take userID -> create media -> return mediaID
	UploadPhoto(userid string, media MediaDB) (string, error)

	// 8) take mediaID -> delete media -> return //
	DeletePhoto(mediaid string) error

	// 9) take mediaID -> reuturn media
	GetMedia(mediaid string) (MediaDB, error)

	// 10) take userID -> return array of media
	GetUserMedia(userid string) ([]MediaDB, error)

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

	//DROP
	//f, _ := db.Exec(`DROP TABLE users IF EXISTS `)
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
		ncomments INTEGER DEFAULT 0 NOT NULL);`

	// 4 - create table media
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
