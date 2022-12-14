package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
	"github.com/nexusfb/WASAPhoto/service/database"
)

// Get user profile with username in query
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - check if username in query
	if !r.URL.Query().Has("username") {
		// no username in query -> return error
		ctx.Logger.Error("error: no username given")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get username from query
	name := r.URL.Query().Get("username")

	// 3 - get logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	// 4 - take user id
	id, err := rt.db.GetUserID(name)
	if errors.Is(err, database.ErrUserProfileDoesNotExists) {
		ctx.Logger.Error("error: user profile does not exist")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		ctx.Logger.Error("error: could not get userID %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5 - check if user is valid
	if !rt.db.ExistenceCheck(id, "user") {
		// user does not exist
		ctx.Logger.Error("error: user does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 6 - check if logged user has been banned by requested user profile
	if rt.db.Check("ban", "bannerid", "bannedid", id, token) {
		ctx.Logger.Error("error: could not get user profile because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// here only if logged user has not been banned by user profile's owner

	// 7 - call get user profile database function
	userProfileDB, err := rt.db.GetUserProfile(name)
	if err != nil {
		// get user profile database function returned error -> return error
		ctx.Logger.WithError(err).Error("Can't provide user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 8 - map returned user profile database to profile struct
	var userProfile structs.UserProfile
	userProfile.FromDatabase(userProfileDB, rt.db)

	// 9 - return user profile struct
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(userProfile)
}
