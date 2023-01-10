package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
)

// Add user to logged user's followings
/* notice: the path is /users/A/followings/B
my idea is that A is the logged user and he adds B to his followings list
I could have done the same by considering A to be the user I want to follow and B myself and add me to his followers list
At the end followers and following are in the same table so it does not matter
I chose this way because users/userid/banned will need userid to be the logged user so I wanted to keep continuity between them */
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - get followerID from path
	fmt.Println("FOLLOW")
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

	// 3 - logged user can follow only from his own profile, check if it is his profile
	if token != followerID {
		ctx.Logger.Error("error: could not follow user because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// notice that in this case I don't check if the requested user exists because the only case in which it follows another profile
	// is if it correspond to the logged user which I take as an assumption to exist

	// here only if logged user is trying to follow another user

	// 4 - get followedID from path
	followedID := ps.ByName("followingid")
	followedID = strings.TrimPrefix(followedID, ":followingid=")
	if len(followedID) == 0 {
		// followedID is empty -> return error
		ctx.Logger.Error("error: followedID is empty")
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
		ctx.Logger.Error("error: you cannot follow yourself")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 7 - check if logged user already follows the specified user

	if rt.db.Check("follow", "followerid", "followedid", token, followedID) {
		// logged user already follows the specified user
		w.WriteHeader(http.StatusOK)
		return
	}

	// 8 - check if logged user has been banned by requested user profile
	if rt.db.Check("ban", "bannerid", "bannedid", followedID, token) {
		ctx.Logger.Error("error: could not follow user because you are not authorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// here only if user exists, logged user does not follow him yet and has not been banned by him

	// 9 - call follow user database function
	err := rt.db.FollowUser(followerID, followedID)
	if err != nil {
		// folloe user database function returned error -> return error
		ctx.Logger.WithError(err).Error("error: can't follow user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 9 - return success (no content)
	w.WriteHeader(http.StatusCreated)
}
