package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// User login with username in request body
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take username from request body
	fmt.Println("LOGIN")
	var username structs.Username
	err := json.NewDecoder(r.Body).Decode(&username)
	fmt.Println("username=" + username.Name)
	if err != nil {
		// the body was not a parseable JSON -> return error
		ctx.Logger.WithError(err).WithField("username", username).Error("error: username is not a parseable JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !username.IsValid() {
		// the username is invalid -> return error
		ctx.Logger.WithError(err).WithField("username", username).Error("error: username is not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - call dologin database function
	newUserID, err := rt.db.DoLogin(username.Name)
	if err != nil {
		// dologin database function returned error -> return error
		ctx.Logger.WithError(err).WithField("username", username.Name).Error("error: can't log you in")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3 - return new userID
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newUserID)
}
