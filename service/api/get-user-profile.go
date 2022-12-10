package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get user profile with username in query
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take username from query
	if !r.URL.Query().Has("username") {
		// no username in query -> return error
		ctx.Logger.Error("error: username not given")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// username in query -> get it
	name := r.URL.Query().Get("username")
	// 2 - call get user profile database function
	userProfileDB, err := rt.db.GetUserProfile(name)
	if err != nil {
		// get user profile database function returned error -> return error
		ctx.Logger.WithError(err).Error("Can't provide user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3 - map returned user profile database to profile struct
	var userProfile structs.UserProfile
	userProfile.FromDatabase(userProfileDB)

	// 4 - return user profile struct
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(userProfile)
}
