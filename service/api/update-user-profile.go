package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/api/structs"
	"github.com/nexusfb/WASAPhoto/service/database"
)

// Update user profile with userid in the path
func (rt *_router) updateUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take userid from path
	userID := ps.ByName("userid")
	userID = strings.TrimPrefix(userID, ":userid=")
	if len(userID) == 0 {
		// userid is empty -> return error
		fmt.Println("Error: userID is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - take new profile from requesy body
	var newProfile structs.UserProfile
	newProfile.UserID = userID
	if err := json.NewDecoder(r.Body).Decode(&newProfile); err != nil {
		// new profile is not a parseable JSON -> return error
		fmt.Println("Error: new profile is not a parseable JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !newProfile.IsValid() {
		// new profile is not valid -> return error
		fmt.Println("Error: new profile is invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3 - call update user profile database function
	_, err := rt.db.UpdateUserProfile(newProfile.ToDatabase())
	if errors.Is(err, database.ErrUserProfileDoesNotExists) {
		// database function returns user profile does not exist -> return error
		fmt.Println("Error: cannot updaqte user profile because user profile does not exist")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// database function returns an error -> return error
		ctx.Logger.WithError(err).WithField("userid", userID).Error("Can't update user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 4 - return updated user profile
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newProfile)
}
