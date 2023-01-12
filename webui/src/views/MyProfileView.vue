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
			comment: "",
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
		seeMediaComments: async function(m){
            this.$router.push({ path: "/media/"+m.id+"/comments/"})
        },
		seeMyStream: async function(){
            this.$router.push({ path: "/stream/"})
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
		},

		commentMedia(m) {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.post("/media/:mediaid="+ m.id+"/comments/", {
					content: this.comment,});
				this.refresh();
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
			this.refresh();
        },
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
			<button v-if="!loading&&this.profile.userid == this.logged" type="button" class="btn btn-primary" @click="seeMyStream">
                see my stream
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
					ncomments: {{ m.ncomments }}<br />
					liked: {{ this.comment }}
				</p>
				<button v-if="!loading&&(this.profile.userid == logged)" type="button" class="btn btn-primary" @click="deleteMedia(m)">
                delete media
            	</button>

				<div class="mb-3">
        		<label for="description" class="form-label">Comment</label>
            	<input type="text" class="form-control" id="comment" v-model="comment" placeholder="comment here">
        		</div>



				<button v-if="!loading&&(this.profile.userid != logged)" type="button" class="btn btn-primary" @click="likeMedia(m)">
                like media
            	</button>
				<button v-if="!loading&&(this.profile.userid != logged)" type="button" class="btn btn-primary" @click="unlikeMedia(m)">
                unlike media
            	</button>
				<button v-if="!loading" type="button" class="btn btn-primary" @click="commentMedia(m)">
                comment media
            	</button>
				<button v-if="!loading" type="button" class="btn btn-primary" @click="seeMediaComments(m)">
                see media comments
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
