package databaseStructs

// database struct of a media
type MediaDB struct {
	MediaID  string // media identifier
	AuthorID string // identifier of the user that published the media
	Date     string // date of publication
	Caption  string // caption of the media (i.e a short description of the media)
	Photo    string // url of the photo of the media
}
