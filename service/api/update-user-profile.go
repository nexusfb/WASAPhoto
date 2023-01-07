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

// Update user profile with userid in the path
func (rt *_router) updateUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from path
	fmt.Println("update in atto")
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	if len(userID) == 0 {
		// userid is empty -> return error
		ctx.Logger.Error("error: userID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	fmt.Println(token)

	// 3 - logged user can change only his own profile, check if it is his profile
	if token != userID {
		fmt.Println(token)
		fmt.Println(userID)
		ctx.Logger.Error("error: could not change username because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// notice that in this case I don't check if the requested user exists because the only case in which he updates
	// the profile is if it correspond to the logged user which I take as an assumption to exist

	// here only if logged user is trying to change his profile

	// 4 - take new profile pic
	/* RK: even if profile picture is not changed I have to save the file again because it is more efficient than other solutions I have found.
	In fact there are two main factors to consider:
		1) pictures are identified by and ID which is random, thus the only way to see if the picture already exists will need linear time
			wrt the number of picure in the folder, since I will have to scan them all to check if it already exists.
			time to upload again picture (and waisting space) <<< scan all photos list
		2) I expect that the update user profile won't be a common operation, usually profiles are updated not too often
		3) It could be dangerous if I choose a internet picture which also another user added to the photos folder, it would be a mistake to have just one copy of the photo

	In conclusion I think that uploading the picture again without considering if the content of the picture has been changed is the best option
	*/

	photo, fileHeader, err := r.FormFile("pic")
	fmt.Println("foto")
	fmt.Println(photo)
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

	//5 - check if photo is valid
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
	fmt.Println("check1")
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
	f, err := os.Create(fmt.Sprintf("./photos/%s%s", Mid, filepath.Ext(fileHeader.Filename)))
	if err != nil {
		ctx.Logger.WithError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, photo)
	if err != nil {
		ctx.Logger.WithError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 7 - create picture url
	newPicURL := fmt.Sprintf("http://localhost:3000/photos/%s%s", Mid, filepath.Ext(fileHeader.Filename))

	// 8 - take bio
	newBio := r.FormValue("bio")

	// 9 - take username
	newUsername := r.FormValue("username")

	// 10 - build new profile
	var newProfile structs.UserProfile
	newProfile.UserID = userID
	newProfile.Bio = newBio
	newProfile.Username = newUsername
	newProfile.ProfilePic = newPicURL

	// 11 - call update user profile database function
	_, err = rt.db.UpdateUserProfile(newProfile.ToDatabase())
	if errors.Is(err, database.ErrUserProfileDoesNotExists) {
		// database function returns user profile does not exist -> return error
		ctx.Logger.WithError(err).WithField("username", newProfile.Username).Error("error: cannot update user profile because user profile does not exist")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// database function returns an error -> return error
		ctx.Logger.WithError(err).WithField("userid", userID).Error("Can't update user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 12 - return updated user profile
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newProfile)
}
