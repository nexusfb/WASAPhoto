

import {createRouter, createWebHashHistory} from 'vue-router'
import NewLoginView from "../views/NewLoginView.vue";
import MyProfileView from "../views/MyProfileView.vue";
import NewMediaView from "../views/NewMediaView.vue";
import UpdateProfileView from "../views/UpdateProfileView.vue";
import ChangeUsernameView from "../views/ChangeUsernameView.vue";
import SearchUsersView from "../views/SearchUsersView.vue";
import MediaComments from "../views/MediaCommentsView.vue";
import MyStream from "../views/NewStreamView.vue";
import UserBans from "../views/UserBans.vue";
import UserFollowers from "../views/UserFollowers.vue";
import UserFollowings from "../views/UserFollowings.vue";

/*
import { createRouter, createWebHashHistory } from 'vue-router'
// import PostForm from '../components/PostForm.vue'
import Login from '../views/aleLoginView.vue'
import ProfileView from '../views/aleProfileView.vue'
// import list from '../components/UserList.vue'
import comments from '../components/Comments.vue'
import Home from '../views/aleHomeView.vue'
// import Edit from '../components/EditPage.vue'
import changed from '../components/ChangeUsername.vue'
// import SinglePhoto from '@/components/SinglePhoto.vue'
// import ErrorMsg from '@/components/ErrorMsg.vue'
*/

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        {path: '/', component: NewLoginView},
        {path: '/users/:username', component: MyProfileView, props: true},
        {path: '/users/:username/newMedia', component: NewMediaView, props: true},
		{path: '/users/:username/followings/', component: UserFollowings, props: true},
		{path: '/users/:username/followers/', component: UserFollowers, props: true},
		{path: '/users/:username/bans/', component: UserBans, props: true},
		{path: '/users/:username/updateProfile', component: UpdateProfileView, props: true},
		{path: '/users/:username/changeUsername', component: ChangeUsernameView, props: true},
		{path: '/search', component: SearchUsersView, props: true},
		{path: '/media/:mediaid/comments/', component: MediaComments, props: true},
		{path: '/stream', component: MyStream, props: true},
		/*
		{ path: '/', component: Login, name: 'Login', alias: "/login" },
		{ path: '/users/:user_id/stream/', component: Home, name: 'Home'},
		// get user profile by name
		{ path: '/users/', component: ProfileView, name: 'Profile', query: { username: { type: String, default: '' } }},
		// change username
		{ path: '/users/:user_id/changeUsername/', component: changed, name: 'username'},
		// change profile
		{ path: '/users/:user_id/editProfile/', component: Edit, name: 'EditPage'},
		// delete profile

		// upload a photo
		{ path: '/users/:user_id/form/', component: PostForm, name: 'PostForm'},
		// Get Single Photo
		{ path: '/post/:photo_id', component: SinglePhoto},

		// See comments
		{ path: '/photos/:photo_id/comments/', component: comments},

		// Get list of the users that added a like/follow/ban
		{ path: '/:listType(likes|followers|following|bans)/', component: list},

		{ path: '/error/:msg', component: ErrorMsg, props: true}*/
    ]
})

export default router



