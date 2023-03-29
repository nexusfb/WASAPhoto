package api

import (
	_ "image/jpeg"
	"io"
	"path/filepath"

	_ "image/png"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
)

// get photo from photos folder
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var media string
	// 1 - get media name from query
	if r.URL.Query().Has("media") {
		media = r.URL.Query().Get("media")

	} else {
		// No image field found
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "query has no field media"}`))
		return
	}

	// 2 - get the path for the image in the server
	image_directory := "/tmp"
	folder_name := "images"
	path := filepath.Join(image_directory, folder_name, media)

	// 3 - open the image
	img, err := os.Open(path)
	if err != nil {
		// error opening file
		ctx.Logger.WithError(err).Error("error: could not open photo file")
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer img.Close()

	// 4 - Send the image back to the front
	w.Header().Set("Content-Type", "blob")
	_, err = io.Copy(w, img)
	if err != nil {
		// error copying image
		ctx.Logger.WithError(err).Error("error: could not copy photo")
		w.WriteHeader(http.StatusInternalServerError)
	}

}
