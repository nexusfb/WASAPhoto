// Page of user profile with username, caption, profile picture, nmedia, nfollowers, nfollowing, user media
<script>
export default {
    data: function() {
        return {
            loading : false,
            errmsg : null,
            profile:"",
			media:[],
        }
    },
    methods: {
        async GetProfile() {
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
		async GetUserMedia() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.get("/users/:userid="+localStorage.getItem('Authorization')+"/media/").then(response => (this.media = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
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
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
			try {
				this.$axios.get("/users/:userid="+localStorage.getItem('Authorization')+"/media/").then(response => (this.media = response.data));
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async deleteMedia(m) {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.delete("/media/:mediaid="+ m.id).then(response => (this.media = response.data));
				this.$router.push({ path: '/users/'+this.profile.username })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        }
    },
    mounted() {
		this.GetUserMedia();
        this.GetProfile();
		this.refresh();
    }
}
</script>

<template>
    <div>
        <h1> Profile</h1>
            </div>
            <div class="card-body">
                <p class="card-text">
					<img :src=this.profile.profilepic><br />
                    name: {{ this.profile.username }}<br />
                    bio: {{ this.profile.bio }}<br />
					media: {{ this.profile.nmedia}}<br />
                    followers: {{ this.profile.nfollowers}}<br />
                    followings: {{ this.profile.nfollowing}}
                </p>
            </div>
            <div>
            <button v-if="!loading" type="button" class="btn btn-primary" @click="createMedia">
                Create new media
            </button>
			<button v-if="!loading" type="button" class="btn btn-primary" @click="updateProfile">
                Update profile
            </button>
			<button v-if="!loading" type="button" class="btn btn-primary" @click="changeUsername">
                change username
            </button>
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
					Caption: {{ m.caption }}
				</p>
				<button v-if="!loading" type="button" class="btn btn-primary" @click="deleteMedia(m)">
                delete media
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
