

import {createRouter, createWebHashHistory} from 'vue-router'
import EnrollView from "../views/EnrollView.vue";
import ResultsView from "../views/ResultsView.vue";
import LoginView from "../views/LoginView.vue";
import MyProfileView from "../views/MyProfileView.vue";
import NewMediaView from "../views/NewMediaView.vue";

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        {path: '/', component: ResultsView},
        {path: '/enroll', component: EnrollView},
        {path: '/login', component: LoginView},
        {path: '/users/:username', component: MyProfileView, props: true},
        {path: '/users/:username/newMedia', component: NewMediaView, props: true},
    ]
})

export default router



