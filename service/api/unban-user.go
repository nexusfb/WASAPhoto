package api

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
)

// Delete user from logged user's banned users list
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - get bannerID from path
	bannerID := ps.ByName("userid")
	bannerID = strings.TrimPrefix(bannerID, ":userid=")
	if len(bannerID) == 0 {
		// bannerID is empty -> return error
		ctx.Logger.Error("error: bannerID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	// 3 - logged user can ban only from his own profile, check if it is his profile
	if token != bannerID {
		ctx.Logger.Error("error: could not change username because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// here only if logged user is trying to ban another user

	// 5 - get bannedID from path
	bannedID := ps.ByName("bannedid")
	bannedID = strings.TrimPrefix(bannedID, ":bannedid=")
	if len(bannedID) == 0 {
		// bannedID is empty -> return error
		ctx.Logger.Error("error: bannedID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 6 - check if user is valid
	if !rt.db.ExistenceCheck(bannedID, "user") {
		// banned user does not exist
		ctx.Logger.Error("error: banned user does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 7 - check validity of the ban
	if bannerID == bannedID {
		ctx.Logger.Error("error: you cannot unban yourself")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 8 - check if logged user has banned the specified user
	if !rt.db.Check("ban", "bannerid", "bannedid", bannerID, bannedID) {
		// logged user has not banned the specified user -> don't do anything
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// 9 - call unban user database function
	err := rt.db.UnbanUser(bannerID, bannedID)
	if err != nil {
		// folloe user database function returned error -> return error
		ctx.Logger.WithError(err).Error("error: can't follow user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 10 - return success (no content)
	w.WriteHeader(http.StatusNoContent)
}
