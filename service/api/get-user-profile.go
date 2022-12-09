package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
	"github.com/nexusfb/WASAPhoto/service/database"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	var profile database.ProfileDB
	var p structs.Profile

	if r.URL.Query().Has("username") {
		name := r.URL.Query().Get("username")
		profile, err = rt.db.GetUserProfile(name)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't provide user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	p.FromDatabase(profile)
	// Send the user profile to the user
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(p)
}
