

import {createRouter, createWebHashHistory} from 'vue-router'
import NewLoginView from "../views/NewLoginView.vue";
import MyProfileView from "../views/MyProfileView.vue";
import SearchUsersView from "../views/SearchUsersView.vue";
import MediaComments from "../views/MediaCommentsView.vue";
import MyStream from "../views/NewStreamView.vue";
import UserBans from "../views/UserBans.vue";
import UserFollowers from "../views/UserFollowers.vue";
import UserFollowings from "../views/UserFollowings.vue";
import Likes from "../views/Likes.vue";

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        {path: '/', component: NewLoginView},
        {path: '/users/:username', component: MyProfileView, props: true},
		{path: '/users/:username/followings/', component: UserFollowings, props: true},
		{path: '/users/:username/followers/', component: UserFollowers, props: true},
		{path: '/users/:username/bans/', component: UserBans, props: true},
		{path: '/search', component: SearchUsersView, props: true},
		{path: '/media/:mediaid/comments/', component: MediaComments, props: true},
		{path: '/media/:mediaid/likes/', component: Likes, props: true},
		{path: '/stream', component: MyStream, props: true},
    ]
})

export default router



