package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
	"github.com/nexusfb/WASAPhoto/service/database"
)

func (rt *_router) getUserMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - prendere userid dal path
	userid := ps.ByName("userid")
	userid = strings.TrimPrefix(userid, ":userid=")
	if userid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - chiamare funzione che con userid ti ritorna foto
	var media []database.MediaDB
	var m []structs.Media

	media, err := rt.db.GetUserMedia(userid)

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't provide user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, element := range media {
		var s structs.Media
		s.FromDatabase(element)
		m = append(m, s)
	}

	// Send the user profile to the user
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(m)
}
