package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
)

// Get username from id
func (rt *_router) getUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - get logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	// 2 - take username from database
	Username, err := rt.db.GetUserName(token)
	if err != nil {
		// get username database function returned error -> return error
		ctx.Logger.WithError(err).WithField("userid", token).Error("error: can't get username from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 7 - return the mapped media struct
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Username)
}
