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
	var username structs.Username
	err := json.NewDecoder(r.Body).Decode(&username.Name)
	if err != nil {
		// the body was not a parseable JSON -> return error
		fmt.Println("Error: username is not a parseable JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !username.IsValid() {
		// the username is invalid -> return error
		fmt.Println("Error: username is invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - call dologin database function
	newUserID, err := rt.db.DoLogin(username.Name)
	if err != nil {
		// dologin database function returned error -> return error
		ctx.Logger.WithError(err).Error("can't log you in")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3- return new userID
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newUserID)
}
