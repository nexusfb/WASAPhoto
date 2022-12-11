package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get all followings of a user with userid in the path
func (rt *_router) getUserFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from path
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	if len(userID) == 0 {
		// userid is empty -> return error
		ctx.Logger.Error("error: userid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - call get user followings database function
	followings, err := rt.db.GetUserFollowers(userID)
	if err != nil {
		/// get user followings database function returned error -> return error
		ctx.Logger.WithError(err).WithField("userID", userID).Error("error: can't get user followers from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3 - map the result followers array into short profile struct array
	var shortProfileArray []structs.ShortProfile
	for _, f := range followings {
		var sp structs.ShortProfile
		sp.FromDatabase(rt.db, f)
		shortProfileArray = append(shortProfileArray, sp)
	}

	// 3 - return array of followings
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(shortProfileArray)
}
