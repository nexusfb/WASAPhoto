package api

import (
	"errors"
	"fmt"
	"net/http"

	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/database"
)

// Delete media with mediaid in the path
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take mediaid from the path
	mediaID := ps.ByName("media")
	mediaID = strings.TrimPrefix(mediaID, ":mediaid=")
	if len(mediaID) == 0 {
		// mediaid is empty -> return error
		fmt.Println("Error: mediaid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - call delete photo database function
	err := rt.db.DeletePhoto(mediaID)
	if errors.Is(err, database.ErrMediaDoesNotExists) {
		// database function returned no media exists -> return
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// database function returned error while deleting -> return error
		ctx.Logger.WithError(err).WithField("mediaID", mediaID).Error("can't delete media")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3 - return success (no content)
	w.WriteHeader(http.StatusNoContent)
}
