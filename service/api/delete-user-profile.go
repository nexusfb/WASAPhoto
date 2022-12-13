package api

import (
	"errors"
	"net/http"

	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/database"
)

// Delete user profile with userid in the path
func (rt *_router) deleteUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from the path
	user := ps.ByName("userid")
	user = strings.TrimPrefix(user, ":userid=")
	if len(user) == 0 {
		// userid is empty -> return error
		ctx.Logger.Error("error: userid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	// 3 - logged user can delete only its own profile, check if it is his profile
	if token != user {
		ctx.Logger.Error("error: could not delete user because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// notice that in this case I don't check if the requested user exists because the only case in which he can delete
	// the profile is if it correspond to the logged user which I take as an assumption to exist

	// here only if logged user is trying to delete his profile

	// 7 - check if logged user already follows the specified user
	if !rt.db.ExistenceCheck(token, "user") {
		// logged user does not follow the specified user
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// 4 - call delete user profile database function
	err := rt.db.DeleteUserProfile(user)
	if errors.Is(err, database.ErrUserProfileDoesNotExists) {
		// should never happen
		// database function returned no user profile exists -> return
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// database function returned error while deleting -> return error
		ctx.Logger.WithError(err).WithField("userid", user).Error("can't delete the user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 5 - return success (no content)
	w.WriteHeader(http.StatusNoContent)
}
