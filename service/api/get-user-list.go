package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get all user with username in the path
func (rt *_router) getUserList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take username from path
	username := ps.ByName("username")
	username = strings.TrimPrefix(username, ":username=")

	// 2 - get logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	// 3 - call get user list database function
	users, err := rt.db.GetUserList(username)
	if err != nil {
		/// get user list database function returned error -> return error
		ctx.Logger.WithError(err).WithField("username", username).Error("error: can't get user list from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 4 - filter user list with ban conditions
	var finalUserList []structs.ShortProfile
	for _, u := range users {

		// 1 - check if user banned logged user
		if rt.db.Check("ban", "bannerid", "bannedid", u, token) {
			continue
		}

		// 2 - check if logged user banned user
		if rt.db.Check("ban", "bannerid", "bannedid", token, u) {
			continue
		}

		// here only if user and logged user did not ban each other

		// translate in short profile
		var sp structs.ShortProfile
		err = sp.FromDatabase(rt.db, u)
		if err != nil {
			ctx.Logger.WithError(err).Error("error: can't map followings to short profile")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 3 - add user to final list
		finalUserList = append(finalUserList, sp)
	}

	// 5 - return array of names
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(finalUserList)
}
