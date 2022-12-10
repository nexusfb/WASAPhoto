package api

import (
	"encoding/json"
	"errors"
	"fmt"
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
	fmt.Println(strings.TrimPrefix(userID, ":userid="))
	if len(userID) == 0 {
		// userid is empty -> return error
		ctx.Logger.Error("error: userID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - take json from request body
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

	// 3 - take old username using userID
	oldName, _ := rt.db.GetUserName(userID)

	// 4 - take old profile using old username
	oldProfile, _ := rt.db.GetUserProfile(oldName)

	// 5 - convert old profile to array of byte
	oldProfileByte, err := json.Marshal(&oldProfile)

	// 6 - create patch
	patchJson := `[{"op": "replace", "path": "/username", "value": "` + newUsernameJson.Name + `"}]`

	// 7 - convert patch to array of byte
	patchByte := []byte(patchJson)

	// 7 - decode patch
	patch, err := jsonpatch.DecodePatch(patchByte)
	if err != nil {
		// decode patch returned error -> return error
		ctx.Logger.WithError(err).WithField("patch", patch).Error("error: could not decode patch")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 8 - apply patch
	newProfileByte, err := patch.Apply(oldProfileByte)
	if err != nil {
		// apply returned error -> return error
		ctx.Logger.WithError(err).WithField("newProfile", newProfileByte).Error("error: could not apply patch")
		return
	}

	// 9 - conver new profile from array of byte to json and put it into a new profile struct
	var newProfileJson structs.UserProfile
	newProfileJson.UserID = userID
	err = json.Unmarshal(newProfileByte, &newProfileJson)
	// 10 - call update user profile with the new profile translated to database profile
	_, err = rt.db.UpdateUserProfile(newProfileJson.ToDatabase())
	if errors.Is(err, database.ErrUserProfileDoesNotExists) {
		// user profile does not exist -> return error
		ctx.Logger.WithError(err).WithField("username", newProfileJson.Username).Error("error: cannot change username because user profile does not exist")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("userID", userID).Error("Can't update user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 11 - return updated profile
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newProfileJson)
}
