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

// questa funzione fa il patch dello username
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - prendere userid dal path
	userid := ps.ByName("userid")
	userid = strings.TrimPrefix(userid, ":userid=")
	if userid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 2 - prendere json dal body
	var newUsernameJson string
	if err := json.NewDecoder(r.Body).Decode(&newUsernameJson); err != nil {
		// body non era un json -> errore
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if 4 == 3 {
		// TODO -> check se nome va bene
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 4 - prendere il vecchio profilo
	oldName, _ := rt.db.GetUserName(userid)
	oldProfile, _ := rt.db.GetUserProfile(oldName)

	// 5 - farlo diventare un json
	oldProfileByte, err := json.Marshal(&oldProfile)

	// 6 - create patch
	patchJson := `[{"op": "replace", "path": "/username", "value": "` + newUsernameJson + `"}]`
	patchByte := []byte(patchJson)

	// 7 - applica patch
	patch, err := jsonpatch.DecodePatch(patchByte)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 8 - modifica versione json profilo
	newProfileByte, err := patch.Apply(oldProfileByte)
	if err != nil {
		panic(err)
	}

	// 9 - new profile
	var newProfileJson structs.Profile
	err = json.Unmarshal(newProfileByte, &newProfileJson)

	// 9 - mandiamo nuovo profilo alla funzione update user profile
	_, err = rt.db.UpdateUserProfile(newProfileJson.ToDatabase())

	// 10 - errori eventuali
	if errors.Is(err, database.ErrUserProfileDoesNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("userid", userid).Error("Can't update user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newProfileJson)
}
