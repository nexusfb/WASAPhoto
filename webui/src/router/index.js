

import {createRouter, createWebHashHistory} from 'vue-router'
import EnrollView from "../views/EnrollView.vue";
import ResultsView from "../views/ResultsView.vue";
import NewLoginView from "../views/NewLoginView.vue";
import MyProfileView from "../views/MyProfileView.vue";
import NewMediaView from "../views/NewMediaView.vue";
import UpdateProfileView from "../views/UpdateProfileView.vue";
import ChangeUsernameView from "../views/ChangeUsernameView.vue";
import SearchUsersView from "../views/SearchUsersView.vue";
import MediaComments from "../views/MediaCommentsView.vue";
import MyStream from "../views/NewStreamView.vue";

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        {path: '/login', component: NewLoginView},
        {path: '/users/:username', component: MyProfileView, props: true},
        {path: '/users/:username/newMedia', component: NewMediaView, props: true},
		{path: '/users/:username/updateProfile', component: UpdateProfileView, props: true},
		{path: '/users/:username/changeUsername', component: ChangeUsernameView, props: true},
		{path: '/search', component: SearchUsersView, props: true},
		{path: '/media/:mediaid/comments/', component: MediaComments, props: true},
		{path: '/stream/', component: MyStream, props: true},
    ]
})

export default router



