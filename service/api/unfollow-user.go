package api

import (
	"errors"
	"net/http"

	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/database"
)

// Delete user from logged user's followings list
func (rt *_router) UnfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - get followerID from path
	followerID := ps.ByName("userid")
	followerID = strings.TrimPrefix(followerID, ":userid=")
	if len(followerID) == 0 {
		// followerID is empty -> return error
		ctx.Logger.Error("error: followerID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 2 - get logged user
	token := r.Header.Get("Authorization")

	// 3 - logged user can unfollow only from his own profile, check if it is his profile
	if token != followerID {
		ctx.Logger.Error("error: could not change username because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// notice that in this case I don't check if the requested user exists because the only case in which he unfollows
	// the user is if it correspond to the logged user which I take as an assumption to exist

	// here only if logged user is trying to unfollow a user

	// 2 - get followedID from path
	followedID := ps.ByName("followingid")
	followedID = strings.TrimPrefix(followedID, ":followingid=")
	if len(followedID) == 0 {
		// followedID is empty -> return error
		ctx.Logger.Error("error: mediaid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3 - check if user is valid
	if !rt.db.ExistenceCheck(followedID, "user") {
		// user does not exist
		ctx.Logger.Error("error: user does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 4 - call delete photo database function
	err := rt.db.UnfollowUser(followerID, followedID)
	if errors.Is(err, database.ErrUserNotFollowed) {
		// database function returned user not followed -> return
		ctx.Logger.WithError(err).WithField("userID", followedID)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// database function returned error while deleting -> return error
		ctx.Logger.WithError(err).WithField("userID", followedID).Error("can't unfollow user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5 - return success (no content)
	w.WriteHeader(http.StatusNoContent)
}