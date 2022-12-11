package api

import (
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
	followerID := ps.ByName("userid")
	followerID = strings.TrimPrefix(followerID, ":userid=")
	if len(followerID) == 0 {
		// followerID is empty -> return error
		ctx.Logger.Error("error: followerID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 2 - get followedID from path
	followedID := ps.ByName("followingid")
	followedID = strings.TrimPrefix(followedID, ":followingid=")
	if len(followedID) == 0 {
		// followedID is empty -> return error
		ctx.Logger.Error("error: mediaid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3 - check validity of the follow
	if followerID == followedID {
		ctx.Logger.Error("error: you cannot follow yourself")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 4 - call follow user database function
	err := rt.db.FollowUser(followerID, followedID)
	if err != nil {
		// folloe user database function returned error -> return error
		ctx.Logger.WithError(err).Error("error: can't follow user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5 - return success (no content)
	w.WriteHeader(http.StatusCreated)
}
