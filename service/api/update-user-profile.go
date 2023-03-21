package api

import (
	"encoding/json"
	"errors"
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
	"github.com/nexusfb/WASAPhoto/service/database"
)

// Update user profile
func (rt *_router) updateUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from path
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	if userID == "" {
		// userid is empty -> return error
		ctx.Logger.Error("error: userID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get authenticated user
	token := r.Header.Get("Authorization")

	// 3 - verify that authenticated user and userid correspond, only the owner of the profile is allowed to update it
	if token != userID {
		ctx.Logger.Error("error: could not update profile because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// notice that in this case I don't check if the requested user exists because the only case in which he updates
	// the profile is if it correspond to the logged user which I take as an assumption to exist
	// here only if logged user is trying to change his profile

	// 4 - decode information inserted by the user in the request body
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error: could not parse form")
		return
	}

	// 5 - create new profile
	var newUserProfile structs.UserProfile

	// 6 - get new username
	var newUsername structs.Username
	// get username from the request body
	newUsername.Name = r.FormValue("username")
	// get old username from the database
	oldUsername, _ := rt.db.GetUserName(userID)
	// check if new username does not correspond to another existing profile
	if newUsername.Name != oldUsername && rt.db.ExistenceCheck(newUsername.Name, "username") {
		// username already exists
		ctx.Logger.Error("error: username already exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !newUsername.IsValid() {
		// username is invalid
		ctx.Logger.WithError(err).Error("error: username format is invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// assign new username to new profile
	newUserProfile.Username = newUsername.Name

	// 7 - get new bio
	newUserProfile.Bio = r.FormValue("bio")
	if !(len(newUserProfile.Bio) <= 200 && structs.BioRx.MatchString(newUserProfile.Bio)) {
		// bio is invalid
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("error: bio format is invalid")
		return
	}

	// 8 - get new profile pic
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
	// check if photo is valid
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
	// create new mediaID
	rawMid, err := uuid.NewV4()
	if err != nil {
		// newV4 returned error -> return error
		ctx.Logger.WithError(err).Error("error encountered while creating new mediaID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Mid := rawMid.String()
	// 6 - save photo in photos folder and save with image id
	/*f, err := os.Create(fmt.Sprintf("./webui/src/assets/photos/%s%s", Mid, filepath.Ext(fileHeader.Filename)))
	if err != nil {
		ctx.Logger.WithError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}*/
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
	newUserProfile.ProfilePic = fileName
	// 7 - create picture url
	// newPicURL := fmt.Sprintf("./src/assets/photos/%s%s", Mid, filepath.Ext(fileHeader.Filename))

	// 9 - set id
	newUserProfile.UserID = userID

	// 10 - call update user profile database function
	_, err = rt.db.UpdateUserProfile(newUserProfile.ToDatabase())
	if errors.Is(err, database.ErrUserProfileDoesNotExists) {
		// database function returns user profile does not exist -> return error
		ctx.Logger.WithError(err).WithField("username", newUserProfile.Username).Error("error: cannot update user profile because user profile does not exist")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// database function returns an error -> return error
		ctx.Logger.WithError(err).WithField("userid", userID).Error("Can't update user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 11 - return updated user profile
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newUserProfile)
}
