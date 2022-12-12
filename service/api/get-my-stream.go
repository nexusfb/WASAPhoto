package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get stream of user with userid in the path
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from path
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	if len(userID) == 0 {
		// userid is empty -> return error
		ctx.Logger.Error("error: userid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get logged user
	token := r.Header.Get("Authorization")

	// 3 - logged user can get stream only from his own profile, check if it is his profile
	if token != userID {
		ctx.Logger.Error("error: could not get stream because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// notice that in this case I don't check if the requested user exists because the only case in which he gets
	// the stream is if it correspond to the logged user which I take as an assumption to exist

	// here only if logged user is trying to get his stream

	// 4 - call get my stream database function
	mediaDBArray, err := rt.db.GetMyStream(userID)
	if err != nil {
		/// get my stream database function returned error -> return error
		ctx.Logger.WithError(err).WithField("userID", userID).Error("error: can't get your stream from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5 - map the result media array into the api version
	var mediaArray []structs.Media
	for _, mediaDB := range mediaDBArray {
		var media structs.Media
		media.FromDatabase(mediaDB, rt.db, token)
		mediaArray = append(mediaArray, media)
	}

	// 6 - return array of followings
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(mediaArray)
}
