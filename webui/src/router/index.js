

import {createRouter, createWebHashHistory} from 'vue-router'
import EnrollView from "../views/EnrollView.vue";
import ResultsView from "../views/ResultsView.vue";
import LoginView from "../views/LoginView.vue";
import MyProfileView from "../views/MyProfileView.vue";
import NewMediaView from "../views/NewMediaView.vue";
import UpdateProfileView from "../views/UpdateProfileView.vue";
import ChangeUsernameView from "../views/ChangeUsernameView.vue";
import SearchUsersView from "../views/SearchUsersView.vue";

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        {path: '/', component: ResultsView},
        {path: '/enroll', component: EnrollView},
        {path: '/login', component: LoginView},
        {path: '/users/:username', component: MyProfileView, props: true},
        {path: '/users/:username/newMedia', component: NewMediaView, props: true},
		{path: '/users/:username/updateProfile', component: UpdateProfileView, props: true},
		{path: '/users/:username/changeUsername', component: ChangeUsernameView, props: true},
		{path: '/search', component: SearchUsersView, props: true},
    ]
})

export default router



