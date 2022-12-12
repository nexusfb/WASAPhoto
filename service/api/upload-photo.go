package api

import (
	"encoding/json"
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
		ctx.Logger.Error("error: userID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get logged user
	token := r.Header.Get("Authorization")

	// 3 - logged user can change username only of his own profile, check if it is his profile
	if token != userID {
		ctx.Logger.Error("error: could not change username because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// notice that in this case I don't check if the requested user exists because the only case in which he uploads
	// a photo in this profile is if it correspond to the logged user which I take as an assumption to exist

	// here only if logged user is trying to upload a photo in his profile

	// 4 - get media from request body
	var media structs.Media
	err := json.NewDecoder(r.Body).Decode(&media)
	if err != nil {
		// media is not a parseable JSON -> return error
		ctx.Logger.WithError(err).WithField("media", media).Error("error: media is not a parseable JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !media.IsValid() {
		// media is not valid -> return error
		ctx.Logger.Error("error: new media is not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 5 - call upload photo database function with userID and converted media struct to database media
	newMediaID, err := rt.db.UploadPhoto(userID, media.ToDatabase())
	if err != nil {
		// upload photo database function returned error -> return error
		ctx.Logger.WithError(err).Error("can't uplad photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 6 - return new media ID
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newMediaID)
}
