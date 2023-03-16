package api

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/database"
)

func (rt *_router) deleteImageFromFolder(photo_id string, w http.ResponseWriter, ctx reqcontext.RequestContext) error {
	photo, err := rt.db.GetMedia(photo_id)
	if errors.Is(err, database.ErrMediaDoesNotExists) {
		// The photo (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "The photo does not exist"}`))
		return err
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't provide photo")
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	path := filepath.Join("/tmp", "images", photo.Photo)
	err = os.Remove(path)
	if err != nil {
		// handle the error
		ctx.Logger.WithError(err).WithField("id", photo_id).Error("image could not be removed from the folder")
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}
