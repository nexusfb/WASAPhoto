package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get all followers of a user with userid in the path
func (rt *_router) getUserFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from path
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	if len(userID) == 0 {
		// userid is empty -> return error
		ctx.Logger.Error("error: userid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - check if user is valid
	if !rt.db.ExistenceCheck(userID, "user") {
		// user does not exist
		ctx.Logger.Error("error: user does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3 - get logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	// 4 - check if logged user has been banned by requested user profile
	res := rt.db.Check("ban", "bannerid", "bannedid", userID, token)
	if res {
		ctx.Logger.Error("error: could not get user followers because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// here only if logged user has not been banned by user profile's owner

	// 5 - call get user followers database function
	followers, err := rt.db.GetUserFollowers(userID)
	if err != nil {
		/// get user followers database function returned error -> return error
		ctx.Logger.WithError(err).WithField("userID", userID).Error("error: can't get user followers from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 6 - map the result followers array into short profile struct array
	var shortProfileArray []structs.ShortProfile
	for _, f := range followers {
		var sp structs.ShortProfile
		sp.FromDatabase(rt.db, f)
		shortProfileArray = append(shortProfileArray, sp)
	}

	// 7 - return array of followers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(shortProfileArray)
}
