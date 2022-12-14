package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Post a new comment with mediaid in the path
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - get mediaid from path
	mediaID := ps.ByName("mediaid")
	mediaID = strings.TrimPrefix(mediaID, ":mediaid=")
	if len(mediaID) == 0 {
		// followerID is empty -> return error
		ctx.Logger.Error("error: followerID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - check if media exists
	if !rt.db.ExistenceCheck(mediaID, "media") {
		// media does not exist
		ctx.Logger.Error("error: user media not exist")
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

	// 5 - check if logged user has been banned from the author of the media ONLY IF he is not the author of the media
	if !(mediaDB.AuthorID == token) {
		if rt.db.Check("ban", "bannerid", "bannedid", mediaDB.AuthorID, token) {
			ctx.Logger.Error("error: could not get user profile because you are not authorized ")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	// here only if logged user has not been banned by media owner

	// 6 - get comment from request body
	var comment structs.Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// comment is not a parseable JSON -> return error
		ctx.Logger.WithError(err).WithField("comment", comment).Error("error: comment is not a parseable JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !comment.IsValid() {
		// comment is not valid -> return error
		ctx.Logger.Error("error: new comment is not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 6 - call comment photo database function with userID and converted media struct to database media
	newCommentID, err := rt.db.CommentPhoto(token, mediaID, comment.ToDatabase(rt.db))
	if err != nil {
		// comment photo database function returned error -> return error
		ctx.Logger.WithError(err).Error("can't comment photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 6 - return new media ID
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newCommentID)
}
