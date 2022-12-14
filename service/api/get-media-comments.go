package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get all media comments with mediaid in the path
func (rt *_router) getMediaComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	if rt.db.Check("ban", "bannerid", "bannedid", mediaDB.AuthorID, token) {
		ctx.Logger.Error("error: could not get user profile because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// here if logged user has not been banned by the media's owner

	// 5 - call get user media database function
	commentsDBArray, err := rt.db.GetMediaComments(mediaID)
	if err != nil {
		/// get user media database function returned error -> return error
		ctx.Logger.WithError(err).WithField("mediaID", mediaID).Error("error: can't get media likes from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 6 - map the result likes array into short profile struct array
	var commentsArray []structs.Comment
	for _, c := range commentsDBArray {
		var comment structs.Comment
		err = comment.FromDatabase(c, rt.db)
		if err != nil {
			ctx.Logger.WithError(err).Error("error: can't map comments from database to API")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		commentsArray = append(commentsArray, comment)
	}

	// 7 - return array of followers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(commentsArray)
}
