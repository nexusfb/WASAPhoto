package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get media with mediaid in the path
func (rt *_router) getMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - get mediaid from path
	mediaID := ps.ByName("mediaid")
	mediaID = strings.TrimPrefix(mediaID, ":mediaid=")
	if len(mediaID) == 0 {
		// mediaid is empty -> return error
		ctx.Logger.Error("error: mediaid is empty")
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

	// 3 - get logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	// 4 - take media
	mediaDB, err := rt.db.GetMedia(mediaID)
	if err != nil {
		// get media database function returned error -> return error
		ctx.Logger.WithError(err).WithField("mediaID", mediaID).Error("error: can't get media from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5 - check if logged user has been banned from the author of the media
	res := rt.db.Check("ban", "bannerid", "bannedid", mediaDB.AuthorID, token)
	if res {
		ctx.Logger.Error("error: could not get user profile because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// 6 - map returned mediaDB into media struct
	var media structs.Media
	media.FromDatabase(mediaDB, rt.db, token)

	// 7 - return the mapped media struct
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(media)
}
