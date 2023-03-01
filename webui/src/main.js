
// default
import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'

// my assets
import './assets/dashboard.css'
import './assets/main.css'

// default
const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.use(router)
app.mount('#app')
/*
import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js'
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'

import './assets/main.css'



const eventBus = new reactive({});
export { eventBus }

const app = createApp(App)

app.config.globalProperties.$axios = axios;
app.component('font-awesome-icon', FontAwesomeIcon)
app.component("ErrorMsg", ErrorMsg);
app.component("WelcomeMsg", WelcomeMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.use(router)
app.mount('#app')
*/