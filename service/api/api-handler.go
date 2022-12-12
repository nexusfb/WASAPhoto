package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	// USER PROFILE routes
	rt.router.POST("/session/", rt.wrap(rt.doLogin))                  // user login
	rt.router.GET("/users/", rt.wrap(rt.getUserProfile))              // get user profile
	rt.router.PATCH("/users/:userid", rt.wrap(rt.setMyUserName))      // change username
	rt.router.PUT("/users/:userid", rt.wrap(rt.updateUserProfile))    // change user profile
	rt.router.DELETE("/users/:userid", rt.wrap(rt.deleteUserProfile)) // delete user profile
	rt.router.GET("/users/:userid/stream/", rt.wrap(rt.getMyStream))  // get logged user stream

	// FOLLOW routes
	rt.router.PUT("/users/:userid/followings/:followingid", rt.wrap(rt.followUser))      // follow user
	rt.router.DELETE("/users/:userid/followings/:followingid", rt.wrap(rt.UnfollowUser)) // unfollow user
	rt.router.GET("/users/:userid/followers/", rt.wrap(rt.getUserFollowers))             // get followers
	rt.router.GET("/users/:userid/followings/", rt.wrap(rt.getUserFollowings))           // get following

	// BAN routes
	rt.router.PUT("/users/:userid/banned/:bannedid", rt.wrap(rt.banUser))      // ban user
	rt.router.DELETE("/users/:userid/banned/:bannedid", rt.wrap(rt.unbanUser)) // unban user
	rt.router.GET("/users/:userid/banned/", rt.wrap(rt.getBannedUsers))        // get banned users

	// MEDIA routes
	rt.router.POST("/users/:userid/media/", rt.wrap(rt.uploadPhoto)) // upload media
	rt.router.DELETE("/media/:mediaid", rt.wrap(rt.deletePhoto))     // delete media
	rt.router.GET("/media/:mediaid", rt.wrap(rt.getMedia))           // get media
	rt.router.GET("/users/:userid/media/", rt.wrap(rt.getUserMedia)) // get user media

	// LIKES routes
	rt.router.PUT("/media/:mediaid/likes/", rt.wrap(rt.likePhoto))      // like a media
	rt.router.DELETE("/media/:mediaid/likes/", rt.wrap(rt.unlikePhoto)) // unlike a media
	rt.router.GET("/media/:mediaid/likes/", rt.wrap(rt.getMediaLikes))  // get media likes

	// COMMENTS routes
	rt.router.POST("/media/:mediaid/comments/", rt.wrap(rt.commentPhoto))    // comment a media
	rt.router.DELETE("/comments/:commentid", rt.wrap(rt.uncommentPhoto))     // uncomment a media
	rt.router.GET("/media/:mediaid/comments/", rt.wrap(rt.getMediaComments)) // get media comments

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
