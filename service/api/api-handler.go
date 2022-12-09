package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	// mine
	rt.router.POST("/session/", rt.wrap(rt.doLogin))
	rt.router.PATCH("/users/:userid", rt.wrap(rt.setMyUserName))
	rt.router.PUT("/users/:userid", rt.wrap(rt.updateUserProfile))
	rt.router.DELETE("/users/:userid", rt.wrap(rt.deleteUserProfile))
	rt.router.GET("/users/", rt.wrap(rt.getUserProfile))

	rt.router.POST("/media/", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/media/:mediaid", rt.wrap(rt.deletePhoto))
	rt.router.GET("/media/mediaid", rt.wrap(rt.getUserProfile))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
