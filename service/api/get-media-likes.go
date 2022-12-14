package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get all media likes with mediaid in the path
func (rt *_router) getMediaLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take mediaid from path
	mediaID := ps.ByName("mediaid")
	mediaID = strings.TrimPrefix(mediaID, ":mediaid=")
	if len(mediaID) == 0 {
		// userid is empty -> return error
		ctx.Logger.Error("error: userid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - check if media exists
	if !rt.db.ExistenceCheck(mediaID, "media") {
		// media does not exist
		ctx.Logger.Error("error: media does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3 - take media
	mediaDB, err := rt.db.GetMedia(mediaID)
	if err != nil {
		// get media database function returned error -> return error
		ctx.Logger.WithError(err).WithField("mediaID", mediaID).Error("error: can't get media from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 4 - check if logged user has been banned from the author of the media
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	res := rt.db.Check("ban", "bannerid", "bannedid", mediaDB.AuthorID, token)
	if res {
		ctx.Logger.Error("error: could not get user profile because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// here if logged user has not been banned by the media's owner

	// 5 - call get user media database function
	likes, err := rt.db.GetMediaLikes(mediaID)
	if err != nil {
		/// get user media database function returned error -> return error
		ctx.Logger.WithError(err).WithField("mediaID", mediaID).Error("error: can't get media likes from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 6 - map the result likes array into short profile struct array
	var shortProfileArray []structs.ShortProfile
	for _, l := range likes {
		var sp structs.ShortProfile
		err = sp.FromDatabase(rt.db, l)
		if err != nil {
			ctx.Logger.WithError(err).Error("error: can't map likes to short profile")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		shortProfileArray = append(shortProfileArray, sp)
	}

	// 7 - return array of likes
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(shortProfileArray)
}
