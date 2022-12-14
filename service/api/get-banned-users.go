package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get all banned user by user with userid in the path
func (rt *_router) getBannedUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from path
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	if len(userID) == 0 {
		// userid is empty -> return error
		ctx.Logger.Error("error: userid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	// 3 - logged user can get bans only from his own profile, check if it is his profile
	if token != userID {
		ctx.Logger.Error("error: could not get bans because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// notice that in this case I don't check if the requested user exists because the only case in which he gets
	// the banned users list is if it correspond to the logged user which I take as an assumption to exist

	// here only if logged user is trying to get bans from his profile

	// 4 - call get banned users database function
	banned, err := rt.db.GetBannedUsers(userID)
	if err != nil {
		//  get banned users database function returned error -> return error
		ctx.Logger.WithError(err).WithField("userID", userID).Error("error: can't get banned users from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5 - map the result banned users array into short profile struct array
	var shortProfileArray []structs.ShortProfile
	for _, f := range banned {
		var sp structs.ShortProfile
		err = sp.FromDatabase(rt.db, f)
		if err != nil {
			ctx.Logger.WithError(err).Error("error: can't map banned to short profile")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		shortProfileArray = append(shortProfileArray, sp)
	}

	// 6 - return array of followings
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(shortProfileArray)
}
