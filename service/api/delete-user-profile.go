package api

import (
	"errors"
	"fmt"
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
		fmt.Println("Error: userid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - call delete user profile database function
	err := rt.db.DeleteUserProfile(user)
	if errors.Is(err, database.ErrUserProfileDoesNotExists) {
		// database function returned no user profile exists -> return
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// database function returned error while deleting -> return error
		ctx.Logger.WithError(err).WithField("userid", user).Error("can't delete the user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 3 - return success (no content)
	w.WriteHeader(http.StatusNoContent)
}
