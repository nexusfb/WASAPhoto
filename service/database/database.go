package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// errori comuni
var (
	// errori comuni di user profile
	ErrUserProfileAlreadyExists = errors.New("user already exists")
	ErrUserProfileDoesNotExists = errors.New("user does not exists")

	//errori comuni media
	ErrMediaDoesNotExists = errors.New("media does not exists")
)

// struct utili
// 1) struct profilo
type ProfileDB struct {
	UserID     string
	Username   string
	Bio        string
	ProfilePic string
	NMedia     uint32
	NFollowers uint32
	NFollowing uint32
}

// 1) struct profilo
type MediaDB struct {
	MediaID   string
	AuthorID  string
	Date      string
	Caption   string
	Photo     string
	NLikes    uint32
	NComments uint32
}

// interfaccia del database
// RK: nomi scelti dal prof ma con la prima lettera maiuscola
type AppDatabase interface {
	// 1) username -> login (userid)
	DoLogin(username string) (string, error)
	// 2) username -> profilo
	GetUserProfile(username string) (ProfileDB, error)
	// 3) profilo -> profilo
	UpdateUserProfile(profile ProfileDB) (ProfileDB, error)
	// 5) delete user profile
	DeleteUserProfile(userid string) error
	// 6) get user name
	GetUserName(userid string) (string, error)

	// 7) upload media
	UploadPhoto(userid string, media MediaDB) (string, error)
	// 8) delete media
	DeletePhoto(mediaid string) error
	// 9) get media
	GetMedia(mediaid string) (MediaDB, error)

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

	tablesQueryArray := [2]string{`CREATE TABLE users (
		userid TEXT NOT NULL PRIMARY KEY,
		username TEXT NOT NULL,
		bio TEXT DEFAULT "" NOT NULL,
		profilepic TEXT DEFAULT "" NOT NULL,
		nmedia INTEGER DEFAULT 0 NOT NULL,
		nfollowers INTEGER DEFAULT 0 NOT NULL,
		nfollowing INTEGER DEFAULT 0 NOT NULL);`,

		`CREATE TABLE media (
		mediaid INTEGER NOT NULL PRIMARY KEY,
		authorid TEXT NOT NULL,
		date TEXT DEFAULT "" NOT NULL,
		caption TEXT DEFAULT "" NOT NULL,
		photo TEXT DEFAULT "" NOT NULL,
		nlikes INTEGER DEFAULT 0 NOT NULL,
		ncomments INTEGER DEFAULT 0 NOT NULL);`}

	//DROP
	f, _ := db.Exec(`DROP TABLE users IF EXISTS `)
	fmt.Println(f)
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users' AND name='media';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		for i := 0; i < len(tablesQueryArray); i++ {
			_, err = db.Exec(tablesQueryArray[i])
			if err != nil {
				return nil, fmt.Errorf("error creating database structure: %w", err)
			}
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
