package database

// IMPORTS
import (
	"database/sql"
	"errors"
	"fmt"
)

// ERRORS
var (
	// errors of userprofile struct
	ErrUserProfileAlreadyExists = errors.New("error: user already exists")
	ErrUserProfileDoesNotExists = errors.New("error: user does not exists")

	// errors of media struct
	ErrMediaDoesNotExists = errors.New("error: media does not exists")

	// error of follow
	ErrUserAlreadyFollowed = errors.New("error: user already followed")
	ErrUserNotFollowed     = errors.New("error: user not followed")

	// error of ban
	ErrUserAlreadyBanned = errors.New("error: user already banned")
	ErrUserNotBanned     = errors.New("error: user not banned")

	// error of like
	ErrUserLikeDoesNotExist = errors.New("error: user like does not exist")
)

// STRUCTS
// database struct for user profile
type UserProfileDB struct {
	UserID     string // identifier of the owner of the profile
	Username   string // username chosen by the owner of the profile
	Bio        string // biography chosen by the owner of the profile (i.e a short description)
	ProfilePic string // url of the profile picture
}

// database struct of a media
type MediaDB struct {
	MediaID  string // media identifier
	AuthorID string // identifier of the user that published the media
	Date     string // date of publication
	Caption  string // caption of the media (i.e a short description of the media)
	Photo    string // url of the photo of the media
}

// database struct of a comment
type CommentDB struct {
	CommentID string // identifier of the comment
	MediaID   string // identifier of the media where the comment was created
	AuthorID  string // identifier of the author of the comment
	Date      string // date of creation of the comment
	Content   string // content of the comment (i.e short text content)
}

// HIGH LEVEL INTERFACE OF THE DATABASE
type AppDatabase interface {
	// USER PROFILE
	DoLogin(username string) (string, error)                        // takes username -> creates user profile -> returns userID
	GetUserProfile(username string) (UserProfileDB, error)          // takes username -> returns user profile
	UpdateUserProfile(profile UserProfileDB) (UserProfileDB, error) // takes new user profile -> updates user profile -> returns new user profile
	DeleteUserProfile(userid string) error                          // takes userID -> deletes user profile -> returns
	FollowUser(userid string, followid string) error                // takes userID -> adds user to following list -> returns error
	UnfollowUser(userid string, followedid string) error            // takes userID -> deletes user from following list -> returns error
	GetUserFollowers(userid string) ([]string, error)               // takes userID -> returns array of followers
	GetUserFollowings(userid string) ([]string, error)              // takes userID -> returns array of followings
	BanUser(bannerID string, bannedID string) error                 // takes bannerID and bannedID -> creates ban -> returns error
	UnbanUser(bannerID string, bannedID string) error               // takes bannerID and bannedID -> deletes ban -> returns error
	GetBannedUsers(bannerID string) ([]string, error)               // takes userID -> returns list of banned users by this userID
	GetMyStream(userID string) ([]MediaDB, error)                   // takes userID -> returns collection of media of followed users

	// UTILITY
	GetUserName(userid string) (string, error)                  // takes userID -> returns username
	GetUserID(username string) (string, error)                  // takes username -> returns userID
	CountRows(table string, column string, event string) uint32 // takes table, column, event -> returns the number of rows in the table for which column==event

	// MEDIA
	UploadPhoto(userid string, media MediaDB) (string, error)                      // takes userID -> creates media -> returns mediaID
	DeletePhoto(mediaid string) error                                              // takes mediaID -> deletes media -> returns
	GetMedia(mediaid string) (MediaDB, error)                                      // takes mediaID -> reuturns media
	GetUserMedia(userid string) ([]MediaDB, error)                                 // takes userID -> returns array of media
	LikePhoto(mediaid string, userid string) error                                 // takes userID and mediaID -> creates like -> reuturns error
	UnlikePhoto(mediaid string, userid string) error                               // takes userID and mediaID -> deletes like -> reuturns error
	GetMediaLikes(mediaid string) []string                                         // takes mediaID -> returns array of users who liked that media
	CommentPhoto(userid string, mediaid string, comment CommentDB) (string, error) // takes comment, mediaID and userID -> creates comment -> reuturns commentID and error
	UncommentPhoto(commentid string) error                                         // takes mediaID and userID -> deletes comment -> reuturns error
	GetMediaComments(mediaid string) ([]CommentDB, error)                          // takes mediaID -> returns array of comments to that media

	// DEFAULT
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
	//fmt.Println(f)

	// 1 - define user table
	tableName := "users"
	sqlStmt := `CREATE TABLE user (
		userid TEXT NOT NULL PRIMARY KEY,
		username TEXT NOT NULL,
		bio TEXT DEFAULT "" ,
		profilepic TEXT DEFAULT "");`

	// 2 - create user table
	err := createTables(tableName, sqlStmt, db)
	if err != nil {
		return nil, fmt.Errorf("error creating user table: %w", err)
	}

	// 3 - define follow table
	tableName = "follow"
	sqlStmt = `CREATE TABLE follow (
		followerid TEXT NOT NULL,
		followedid TEXT NOT NULL,
		FOREIGN KEY (followerid) REFERENCES user(userid),
		FOREIGN KEY (followerid) REFERENCES user(userid),
		PRIMARY KEY (followerid, followingid));`

	// 4 - create follow table
	err = createTables(tableName, sqlStmt, db)
	if err != nil {
		return nil, fmt.Errorf("error creating follow table: %w", err)
	}

	// 5 - define ban table
	tableName = "ban"
	sqlStmt = `CREATE TABLE ban (
		bannerid TEXT NOT NULL,
		bannedid TEXT NOT NULL,
		FOREIGN KEY (bannerid) REFERENCES user(userid),
		FOREIGN KEY (bannedid) REFERENCES user(userid),
		PRIMARY KEY (bannerid, bannedid));`

	// 6 - create ban table
	err = createTables(tableName, sqlStmt, db)
	if err != nil {
		return nil, fmt.Errorf("error creating ban table: %w", err)
	}

	// 7 - define media table
	tableName = "media"
	sqlStmt = `CREATE TABLE media (
		mediaid INTEGER NOT NULL PRIMARY KEY,
		authorid TEXT NOT NULL,
		date TEXT DEFAULT "",
		caption TEXT DEFAULT "",
		photo TEXT DEFAULT "",
		FOREIGN KEY (authorid) REFERENCES user(userid));`

	// 8 - create media table
	err = createTables(tableName, sqlStmt, db)
	if err != nil {
		return nil, fmt.Errorf("error creating media table: %w", err)
	}

	// 9 - define like table
	tableName = "like"
	sqlStmt = `CREATE TABLE like (
			userid TEXT NOT NULL,
			mediaid TEXT NOT NULL,
			FOREIGN KEY (userid) REFERENCES user(userid),
			FOREIGN KEY (mediaid) REFERENCES media(mediaid),
			PRIMARY KEY (userid, mediaid));`

	// 10 - create like table
	err = createTables(tableName, sqlStmt, db)
	if err != nil {
		return nil, fmt.Errorf("error creating like table: %w", err)
	}

	// 11 - define comment table
	tableName = "comment"
	sqlStmt = `CREATE TABLE comment (
			commentid TEXT NOT NULL PRIMARY KEY,
			userid TEXT NOT NULL,
			mediaid TEXT NOT NULL,
			date TEXT DEFAULT "",
			content TEXT DEFAULT "",
			FOREIGN KEY (userid) REFERENCES user(userid),
			FOREIGN KEY (mediaid) REFERENCES media(mediaid));`

	// 12 - create comment table
	err = createTables(tableName, sqlStmt, db)
	if err != nil {
		return nil, fmt.Errorf("error creating comment table: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

// function that given a table name and sql statement -> creates table
func createTables(tableName string, sqlStmt string, db *sql.DB) error {
	var table string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='` + tableName + `';`).Scan(&table)
	if errors.Is(err, sql.ErrNoRows) {
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return fmt.Errorf("error creating table: %w", err)
		}
	}
	return err
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
