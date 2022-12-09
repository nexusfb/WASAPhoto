package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
)

// Get media with mediaid in the path
func (rt *_router) getMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - get mediaid from path
	mediaID := ps.ByName("media")
	mediaID = strings.TrimPrefix(mediaID, ":mediaid=")
	if len(mediaID) == 0 {
		// mediaid is empty -> return error
		fmt.Println("Error: mediaid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - call get media database function
	mediaDB, err := rt.db.GetMedia(mediaID)
	if err != nil {
		// get media database function returned error -> return error
		ctx.Logger.WithError(err).Error("Can't provide media")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3 - map returned mediaDB into media struct
	var media structs.Media
	media.FromDatabase(mediaDB, rt.db)

	// 4 - return the mapped media struct
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(media)
}
