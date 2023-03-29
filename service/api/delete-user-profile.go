package api

import (
	"net/http"

	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
)

// Delete user profile with userid in the path
func (rt *_router) deleteUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from the path
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	if userID == "" {
		// userid is empty -> return error
		ctx.Logger.Error("error: userid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	// 3 - logged user can delete only its own profile, check if it is his profile
	if token != userID {
		ctx.Logger.Error("error: could not delete user because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// notice that in this case I don't check if the requested user exists because the only case in which he can delete
	// the profile is if it correspond to the logged user which I take as an assumption to exist

	// here only if logged user is trying to delete his profile

	// 4 - get user media
	userMedia, err := rt.db.GetUserMedia(userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get user media")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5 - delete all media
	for _, media := range userMedia {
		err := rt.deleteImageFromFolder(media.MediaID, w, ctx)
		if err != nil {
			ctx.Logger.WithError(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// 6 - call delete user profile database function
	err = rt.db.DeleteUserProfile(userID)
	if err != nil {
		// database function returned error while deleting -> return error
		ctx.Logger.WithError(err).WithField("userid", userID).Error("can't delete the user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 7 - return success (no content)
	w.WriteHeader(http.StatusNoContent)
}
