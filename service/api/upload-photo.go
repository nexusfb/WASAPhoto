package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Post a new media with userid in the path
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from path
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	if len(userID) == 0 {
		// userid is empty -> return error
		ctx.Logger.Error("error: userID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get authenticated user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	// 3 - authenticated user can upload photos only in his own profile, check if it is his profile
	if token != userID {
		ctx.Logger.Error("error: could not upload photo because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// notice that in this case I don't check if the requested user exists because the only case in which he uploads
	// a photo in this profile is if it correspond to the logged user which I take as an assumption to exist

	// here only if logged user is trying to upload a photo in his profile

	// 4 - decode information inserted by the user in the request body
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error: could not parse form")
		return
	}

	// 5 - create new media
	var newUserMedia structs.Media

	// 6 - get photo from request body
	photo, fileHeader, err := r.FormFile("pic")
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not parse photo")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer photo.Close()
	buff := make([]byte, 512)
	_, err = photo.Read(buff)
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not read photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 7 - check if photo is valid
	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" {
		ctx.Logger.WithError(err).Error("error: The provided file format is not allowed. Please upload a JPEG,JPG or PNG image")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = photo.Seek(0, io.SeekStart)
	if err != nil {
		ctx.Logger.WithError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 8 - create new mediaID
	rawMid, err := uuid.NewV4()
	if err != nil {
		// newV4 returned error -> return error
		ctx.Logger.WithError(err).Error("error encountered while creating new mediaID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Mid := rawMid.String()
	newUserMedia.MediaID = Mid

	// 9 - Save the photo in the images folder
	imageDir := "/tmp"
	folderName := "images"
	fileName := fmt.Sprintf("%s%s", Mid, filepath.Ext(fileHeader.Filename))
	path := filepath.Join(imageDir, folderName, fileName)
	f, err := os.Create(path)
	if err != nil {
		ctx.Logger.WithError(err)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, photo)
	if err != nil {
		ctx.Logger.WithError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 10 - Fill media struct
	newUserMedia.Photo = fileName
	newUserMedia.Caption = r.FormValue("cap")
	newUserMedia.AuthorID = token

	// 11 - convert into database struct
	mediadb := newUserMedia.ToDatabase(rt.db)

	// 12 - call upload photo database function
	err = rt.db.UploadPhoto(mediadb)
	if err != nil {
		// upload photo database function returned error -> return error
		ctx.Logger.WithError(err).Error("can't uplad photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 13 - return new media ID
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Mid)
}
