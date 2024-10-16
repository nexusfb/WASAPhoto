package api

import (
	"errors"
	"net/http"

	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/database"
)

// Delete media with mediaid in the path
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take mediaid from the path
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
		// media does not exist -> don't do anything
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// 3 - check if media belongs to logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if !rt.db.Check("media", "mediaid", "authorid", mediaID, token) {
		ctx.Logger.Error("error: only the author of the media can delete the media")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// here if the logged user is the owner of the media

	// 4 - delete image from tmp directory
	err := rt.deleteImageFromFolder(mediaID, w, ctx)
	if err != nil {
		ctx.Logger.WithError(err).Error("error in deleting images from folder")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5 - call delete photo database function
	err = rt.db.DeletePhoto(mediaID)
	if errors.Is(err, database.ErrMediaDoesNotExists) {
		// database function returned no media exists -> don't do anything
		ctx.Logger.WithError(err).WithField("mediaID", mediaID)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// database function returned error while deleting -> return error
		ctx.Logger.WithError(err).WithField("mediaID", mediaID).Error("can't delete media")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 6 - return success (no content)
	w.WriteHeader(http.StatusNoContent)
}
