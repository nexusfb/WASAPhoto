// Page of user profile with username, caption, profile picture, nmedia, nfollowers, nfollowing, user media
<script>
export default {
    data: function() {
        return {
            loading : false,
            errmsg : null,
            profile: this.GetProfile(),
			media:[],
			logged: localStorage.getItem('Authorization'),
        }
    },
    methods: {
        GetProfile() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.get("/users/?username="+this.$route.params.username).then(response => (this.profile = response.data));
            } catch (e) {
                this.errormsg = e.toString();
				
            }
            this.loading = false;
        },

		GetUserMedia() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.get("/users/:userid="+this.profile.userid+"/media/").then(response => (this.media = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
		searchUsers: async function(){
            this.$router.push({ path: '/search'})
        },
        createMedia: async function(){
            this.$router.push({ path: '/users/'+this.profile.username+"/newMedia"})
        },
		updateProfile: async function(){
            this.$router.push({ path: '/users/'+this.profile.username+"/updateProfile"})
        },
		changeUsername: async function(){
            this.$router.push({ path: '/users/'+this.profile.username+"/changeUsername"})
        },
	
		refresh() {
			this.GetProfile();
			this.GetUserMedia();
		},
		deleteMedia(m) {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.delete("/media/:mediaid="+ m.id);
				this.$router.push({ path: '/users/'+this.profile.username })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
			this.refresh();
        },
		likeMedia(m) {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.put("/media/:mediaid="+ m.id+"/likes/");
				this.refresh();
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
			this.refresh();
        },
		unlikeMedia(m) {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.delete("/media/:mediaid="+ m.id+"/likes/");
				this.refresh();
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
			this.refresh();
        },
		followUser() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.put("/users/:userid="+this.logged+"/followings/:followingid="+this.profile.userid);
				this.$router.push({ path: '/users/'+this.profile.username })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
			this.refresh();
        },

		unfollowUser() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.delete("/users/:userid="+this.logged+"/followings/:followingid="+this.profile.userid);
				this.$router.push({ path: '/users/'+this.profile.username })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
			this.refresh();
        },

		async toggle(){
			if (this.profile.followed==true){
				this.unfollowUser();
				this.profile.followed = false;
			}
			else{
				this.followUser();
				this.profile.followed = true;
			}
		}
    },
    mounted() {
		this.refresh();
    },
}
</script>

<template>
    <div> <h1> Profile</h1></div>
            <div  class="card-body">
                <p class="card-text">
					<img :src= this.profile.profilepic ><br />
                    name: {{ this.profile.username }}<br />
                    bio: {{ this.profile.bio }}<br />
					media: {{ this.profile.nmedia}}<br />
                    followers: {{ this.profile.nfollowers}}<br />
                    followings: {{ this.profile.nfollowing }}
                </p>
            </div>
            <div>
            <button v-if="!loading" type="button" class="btn btn-primary" @click="createMedia">
                Create new media
            </button><br /><br />
			<button v-if="!loading" type="button" class="btn btn-primary" @click="updateProfile">
                Update profile
            </button><br /><br />
			<button v-if="!loading" type="button" class="btn btn-primary" @click="changeUsername">
                change username
            </button><br /><br />
			<button v-if="!loading" type="button" class="btn btn-primary" @click="searchUsers">
                search users
            </button><br /><br />
			<div v-if= "this.profile.userid != logged">
				<button v-if="!loading" type="button" class="ui button big" @click="toggle">
					{{profile.followed ? 'unfollow' : 'follow'}}
            	</button>

			</div>


            <LoadingSpinner v-if="loading"></LoadingSpinner>


			<div class="card" v-if="media === null">
				<div class="card-body">
					<p>No media in the database.</p>
				</div>
			</div>

		<div class="card" v-if="!loading" v-for="m in media">
			<div class="card-header">
				Media
			</div>
			<div class="card-body">
				<p class="card-text">
					<img :src=m.photo><br />
					Caption: {{ m.caption }}<br />
					nlikes: {{ m.nlikes }}<br />
					liked: {{ m.liked }}
				</p>
				<button v-if="!loading&&(this.profile.userid == logged)" type="button" class="btn btn-primary" @click="deleteMedia(m)">
                delete media
            	</button>
				<button v-if="!loading&&(this.profile.userid != logged)" type="button" class="btn btn-primary" @click="likeMedia(m)">
                like media
            	</button>
				<button v-if="!loading&&(this.profile.userid != logged)" type="button" class="btn btn-primary" @click="unlikeMedia(m)">
                unlike media
            	</button>
			</div>
		</div>
    </div>

</template>

<style scoped>
.card {
    margin-bottom: 20px;
}
</style>
