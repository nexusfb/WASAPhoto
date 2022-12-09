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

// Post a new media with userid in the path
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from path
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	if len(userID) == 0 {
		// userid is empty -> return error
		fmt.Println("Error: userID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get media from request body
	var media structs.Media
	err := json.NewDecoder(r.Body).Decode(&media)
	if err != nil {
		// media is not a parseable JSON -> return error
		fmt.Println("Error: media is not a parseable JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !media.IsValid() {
		// media is not valid -> return error
		fmt.Println("Error: new media is invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3 - call upload photo database function with userID and converted media struct to database media
	newMediaID, err := rt.db.UploadPhoto(userID, media.ToDatabase())
	if err != nil {
		// upload photo database function returned error -> return error
		ctx.Logger.WithError(err).Error("can't uplad photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 4 - return new media ID
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newMediaID)
}
