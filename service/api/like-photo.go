package api

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
)

// add like of the logged user to the
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// 5 - check if logged user has been banned from the author of the media
	res := rt.db.Check("ban", "bannerid", "bannedid", mediaDB.AuthorID, token)
	if res {
		ctx.Logger.Error("error: could not get user profile because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// here only if logged user has not been banned by media owner

	// 6 - check validity of the like
	if mediaDB.AuthorID == token {
		ctx.Logger.Error("error: you cannot like your own media")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 7 - check if already liked by the logged user
	if rt.db.Check("like", "mediaid", "userid", mediaID, token) {
		// user already liked this media
		w.WriteHeader(http.StatusOK)
		return
	}

	// here only if logged user could like this media but did not like this media yet

	// 8 - call like Photo database function
	err = rt.db.LikePhoto(mediaID, token)
	if err != nil {
		// like photo user database function returned error -> return error
		ctx.Logger.WithError(err).Error("error: can't like photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 9 - return success (no content)
	w.WriteHeader(http.StatusCreated)
}
