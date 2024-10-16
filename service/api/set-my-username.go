package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"strings"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
	"github.com/nexusfb/WASAPhoto/service/database"
)

// Change username with username in request body
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	token = strings.TrimPrefix(token, "Bearer ")

	// 3 - logged user can change username only of his own profile, check if it is his profile
	if token != userID {
		ctx.Logger.Error("error: could not change username because you are not authorized ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// notice that in this case I don't check if the requested user exists because the only case in which he updates
	// the username is if it correspond to the logged user which I take as an assumption to exist

	// here only ifs logged user is trying to change his username

	// 4 - take json from request body
	var newUsernameJson structs.Username
	if err := json.NewDecoder(r.Body).Decode(&newUsernameJson); err != nil {
		// username is not a parseable JSON -> return error
		ctx.Logger.WithError(err).WithField("username", newUsernameJson).Error("error: username is not a parseable JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !newUsernameJson.IsValid() {
		// userid is not valid -> return error
		ctx.Logger.WithError(err).WithField("username", newUsernameJson).Error("error: username is not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 5 - take old username using userID
	oldName, err := rt.db.GetUserName(userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not get old name")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 6 - take old profile using old username
	oldProfile, err := rt.db.GetUserProfile(oldName)
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not get old profile")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 7 - convert old profile to array of byte
	oldProfileByte, err := json.Marshal(&oldProfile)
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not marshal old profile to JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 8 - create patch
	patchJson := `[{"op": "replace", "path": "/username", "value": "` + newUsernameJson.Name + `"}]`

	// 9 - convert patch to array of byte
	patchByte := []byte(patchJson)

	// 10 - decode patch
	patch, err := jsonpatch.DecodePatch(patchByte)
	if err != nil {
		// decode patch returned error -> return error
		ctx.Logger.WithError(err).WithField("patch", patch).Error("error: could not decode patch")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 11- apply patch
	newProfileByte, err := patch.Apply(oldProfileByte)
	if err != nil {
		// apply returned error -> return error
		ctx.Logger.WithError(err).WithField("newProfile", newProfileByte).Error("error: could not apply patch")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 12 - conver new profile from array of byte to json and put it into a new profile struct
	var newProfileJson structs.UserProfile
	err = json.Unmarshal(newProfileByte, &newProfileJson)
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not unmarshal new profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 13 - call update user profile with the new profile translated to database profile
	_, err = rt.db.UpdateUserProfile(newProfileJson.ToDatabase())
	if errors.Is(err, database.ErrUserProfileDoesNotExists) {
		// user profile does not exist -> return error
		// should never happen since it is already checked in the API
		ctx.Logger.WithError(err).WithField("username", newProfileJson.Username).Error("error: cannot change username because user profile does not exist")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("userID", userID).Error("Can't update user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 14 - return updated profile
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newProfileJson)
}
