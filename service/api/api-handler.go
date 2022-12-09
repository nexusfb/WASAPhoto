package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	// WASAPhoto routes

	// 1 - user login
	rt.router.POST("/session/", rt.wrap(rt.doLogin))
	// curl -X POST "http://localhost:3000/session/" -H 'Content-Type: application/json' -d '{"username":"Aldo"}'

	// 2 - change username
	rt.router.PATCH("/users/:userid", rt.wrap(rt.setMyUserName))
	// curl -X PATCH http://localhost:3000/users/:user_id=2IeFY94UsNvbeRIC9KFXoeABais/ -H 'Content-Type: application/json' -H 'Accept: application/json' -d '{"username":"paola"}'

	// 3 - change user profile
	rt.router.PUT("/users/:userid", rt.wrap(rt.updateUserProfile))
	// curl -X PUT -H http://localhost:3000/users/:user_id=2IeFY94UsNvbeRIC9KFXoeABais/ "Content-Type: application/json" -H 'Accept: application/json' -d '{"username":"paola"}'

	// 4 - delete user profile
	rt.router.DELETE("/users/:userid", rt.wrap(rt.deleteUserProfile))
	// curl -X DELETE -H http://localhost:3000/users/:user_id=2IeFY94UsNvbeRIC9KFXoeABais/""

	// 5 - get user profile
	rt.router.GET("/users/", rt.wrap(rt.getUserProfile))
	// curl "http://localhost:3000/users/?username=alessandro"

	// 6 - upload media
	rt.router.POST("/media/", rt.wrap(rt.uploadPhoto))
	// curl -X POST "http://localhost:3001/users/:user_id=2IdzRfVifWvS1US33tMpnPZUWun/photos/" -H 'Content-Type: application/json' -d '{"caption":"Vodka may not be the answer but it's worth a shot", "image":"http://lab.it/logo.png"}'

	// 7 - delete media
	rt.router.DELETE("/media/:mediaid", rt.wrap(rt.deletePhoto))
	// curl -X DELETE "http://localhost:3001/users/:mediaid=""

	// 8 - get media
	rt.router.GET("/media/mediaid", rt.wrap(rt.getUserProfile))
	// curl "http://localhost:3000/media/:mediaid="

	// 9 - get user media
	rt.router.GET("/users/:userid/media/", rt.wrap(rt.getUserProfile))
	// curl "http://localhost:3000/users/:user_id=2IdzR

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
