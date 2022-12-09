package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

func (rt *_router) getMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - get userid from path
	mediaID := ps.ByName("media")
	mediaID = strings.TrimPrefix(mediaID, ":mediaid=")
	if mediaID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mediaDB, err := rt.db.GetMedia(mediaID)

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't provide media")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var media structs.Media
	media.FromDatabase(mediaDB)
	// Send the user profile to the user
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(media)
}
