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
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	token = strings.TrimPrefix(token, "Bearer ")

	// 3 - logged user can unfollow only from his own profile, check if it is his profile
	if token != followerID {
		ctx.Logger.Error("error: could not unfollow user because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// notice that in this case I don't check if the requested user exists because the only case in which he unfollows
	// the user is if it correspond to the logged user which I take as an assumption to exist

	// here only if logged user is trying to unfollow a user

	// 4 - get followedID from path
	followedID := ps.ByName("followingid")
	followedID = strings.TrimPrefix(followedID, ":followingid=")
	if len(followedID) == 0 {
		// followedID is empty -> return error
		ctx.Logger.Error("error: mediaid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 5 - check if user is valid
	if !rt.db.ExistenceCheck(followedID, "user") {
		// user does not exist
		ctx.Logger.Error("error: user does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 6 - check validity of the follow
	if followerID == followedID {
		ctx.Logger.Error("error: you cannot unfollow yourself")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 7 - check if logged user does not follow the specified user
	if !rt.db.Check("follow", "followerid", "followedid", token, followedID) {
		// logged user does not follow the specified user
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// 8 - check if logged user has been banned by requested user profile
	if rt.db.Check("ban", "bannerid", "bannedid", followedID, token) {
		ctx.Logger.Error("error: could unfollow user because you are not authorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// 9 - call delete photo database function
	err := rt.db.UnfollowUser(followerID, followedID)
	if errors.Is(err, database.ErrUserNotFollowed) {
		// database function returned user not followed -> don't do anything
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// database function returned error while deleting -> return error
		ctx.Logger.WithError(err).WithField("userID", followedID).Error("can't unfollow user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 10 - return success (no content)
	w.WriteHeader(http.StatusNoContent)
}
