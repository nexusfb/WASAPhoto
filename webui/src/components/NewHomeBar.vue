// Home bar to move between three main pages: stream, search page, logged user profile
<script>
export default {
    name: "NavBar",
    data: function () {
        return {
            errormsg: null,
            loading: false,
            logged: localStorage.getItem('Authorization'),
            name: "",
            myProfilePath:'/users/'+ this.name,

        }
    },
    props:{
        profilo: "",
    },
    methods: {
        async GetUsername() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                let response = await this.$axios.get("/id")
				this.name = response.data;
            } catch (e) {
                this.errormsg = e.toString();
				
            }
            this.loading = false;
        },
        
        refresh() {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
			this.GetUsername();
            this.loading = false;
        },
        rect() {
            const rectangle = document.querySelector('.rectangle');
            rectangle.classList.toggle('active')
        },
    },
    mounted() {
        this.refresh()
    }
}
</script>


<template>
    <nav>
        <div class="nav-wrapper center">

            <div id="nav-home-section" class="nav-section">
                <router-link :to="'/users/'+ this.name" class="font-style">home</router-link>
            </div>
            <div id="nav-home-section" class="nav-section">
                <router-link to="/stream" class="font-style">Stream</router-link>
            </div>
            <div id="nav-home-section" class="nav-section" >
                <router-link to="/search" class="font-style">Search</router-link>
            </div>
    

        </div>
       

    </nav>
</template>
<style>
:root {
    --ba1: rgb(6, 12, 24);
    --bo1: rgba(255, 255, 255, 0.2);
}
nav {
    max-width: 604px;
    margin-left: auto;
    margin-right: auto;
}
.font-style {
    font-size: 1.3em;
    font-family: "Copperplate", sans-serif;
    font-weight: 400;
    color: #fcecd4;
    text-decoration: none;
}
.nav-wrapper {
    background-color: var(--ba1);
    border: 2px solid var(--bo1);
    overflow: hidden;
    position: fixed;
    width: 1450px;
    height: 5vh;
    top: 0;
    z-index: 99;
    background-color: #DDBEA8;
    border-radius: 35px;
}
.center {
    display: flex;
    align-items: center;
    justify-content: center;
}
.nav-wrapper>.nav-section {
    padding: 4rem 0;
    display: flex;
    gap: 1rem;
    border-left: 1px solid var(--bo1);
    justify-content: center;
    color:#DDBEA8
}
#nav-home-section,
#nav-search-section,
#nav-profile-section {
    flex-basis: calc(100% / 3);
}
.icons {
    text-decoration: none;
    color: white;
    cursor: pointer;
}
.rectangle {
    position: fixed;
    bottom: 0;
    width: 400px;
    background-color: rgb(10, 12, 24);
    min-height: 12vh;
    height: 12vh;
    border-radius: 10px 10px 0 0;
    transform: translateY(0px);
    transition: transform 1.5s ease-in;
    z-index: 0;
}
.rectangle.active {
    transform: translateY(-116px);
    transition: transform 1.5s ease-out;
    z-index: 98;
}
.input-wrapper {
    display: flex;
    align-items: center;
    justify-content: space-between;
    background: rgb(100, 100, 100);
    box-shadow: 0.25rem 0.25rem 0rem rgb(189, 189, 189);
    width: 75%;
    height: 40px;
    padding: 0.5rem;
    border-radius: 0.5rem;
    position: absolute;
    left: 10%;
    top: 50%;
}
#search {
    margin: 0 0.5rem 0 0.5rem;
    width: 100%;
    height: 20px;
    padding: 0.5rem;
    border: none;
    outline: none;
    background: rgb(34, 34, 34);
    color: white;
    font-size: 1.5rem;
}
.title {
    font-size: 1.8em;
    color: white;
    padding-left: 30px;
    padding-top: 10px;
}
</style>