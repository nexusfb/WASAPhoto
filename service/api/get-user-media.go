package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get all media of a user with userid in the path
func (rt *_router) getUserMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from path
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	if len(userID) == 0 {
		// userid is empty -> return error
		ctx.Logger.Error("error: userid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - call get user media database function
	mediaDBArray, err := rt.db.GetUserMedia(userID)
	if err != nil {
		/// get user media database function returned error -> return error
		ctx.Logger.WithError(err).WithField("userID", userID).Error("error: can't get user media from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3 - map the result mediaDB array into media struct array
	var mediaArray []structs.Media
	for _, mediaDB := range mediaDBArray {
		var media structs.Media
		media.FromDatabase(mediaDB, rt.db)
		mediaArray = append(mediaArray, media)
	}

	// 4 - return array of media structs
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(mediaArray)
}
