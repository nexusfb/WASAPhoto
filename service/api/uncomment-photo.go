package api

import (
	"errors"
	"net/http"

	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
	"github.com/nexusfb/WASAPhoto/service/database"
)

// Delete comment to media with mediaid in the path
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// 1 - take commentid from the path
	commentID := ps.ByName("commentid")
	commentID = strings.TrimPrefix(commentID, ":commentid=")
	if len(commentID) == 0 {
		// commentID is empty -> return error
		ctx.Logger.Error("error: commentid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3 - check if comment exists
	if !rt.db.ExistenceCheck(commentID, "comment") {
		// comment does not exist -> don't do anything
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// 2 - check if comment belongs to logged user
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if !rt.db.Check("comment", "commentid", "userid", commentID, token) {
		ctx.Logger.Error("error: only the author of the comment can delete the comment")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// here if the logged user is the owner of the comment

	// 3 - call delete comment database function
	err := rt.db.UncommentPhoto(commentID)
	if errors.Is(err, database.ErrCommentDoesNotExist) {
		// database function returned no media exists -> don't do anything
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// database function returned error while deleting -> return error
		ctx.Logger.WithError(err).WithField("commentID", commentID).Error("can't delete comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 4 - return success (no content)
	w.WriteHeader(http.StatusNoContent)
}
