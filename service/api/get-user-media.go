package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get all media of a user with userid in the path
func (rt *_router) getUserMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from path
	fmt.Println("CALLED GET USER MEDIA")
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	fmt.Println("userid=" + userID)
	if len(userID) == 0 {
		// userid is empty -> return error
		ctx.Logger.Error("error: userid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - check if user exists
	if !rt.db.ExistenceCheck(userID, "user") {
		// user does not exist
		ctx.Logger.Error("error: user does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3 - get logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	// 4 - check if logged user has been banned by the author of the media
	if rt.db.Check("ban", "bannerid", "bannedid", userID, token) {
		ctx.Logger.Error("error: could not get user profile because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// here if logged user has not been banned by the user in the path

	// 5 - call get user media database function
	mediaDBArray, err := rt.db.GetUserMedia(userID)
	if err != nil {
		/// get user media database function returned error -> return error
		ctx.Logger.WithError(err).WithField("userID", userID).Error("error: can't get user media from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 6 - map the result mediaDB array into media struct array
	var mediaArray []structs.Media
	for _, mediaDB := range mediaDBArray {
		var media structs.Media
		err = media.FromDatabase(mediaDB, rt.db, token)
		if err != nil {
			ctx.Logger.WithError(err).Error("error: can't map media from database to API")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		mediaArray = append(mediaArray, media)
	}

	// 7 - return array of media structs
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(mediaArray)
}
