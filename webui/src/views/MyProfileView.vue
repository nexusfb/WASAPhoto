// Page of user profile with username, caption, profile picture, nmedia, nfollowers, nfollowing, user media
<script>
import NewMedia from "@/components/NewMedia.vue";
import NavBar from "@/components/NewHomeBar.vue";
export default {
    components: { NewMedia, NavBar},
    data: function() {
        return {
            loading : false,
            errmsg : true,
            profile: {},
			media:[],
			logged: localStorage.getItem('Authorization'),
			comment: "",
			creatingMedia: false,
			changingProfile: false,
			changingusername: false,
			photo: "",
         	caption:"",
         	preview: "",
			newppic:null,

        }
    },
    methods: {
        async GetProfile() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                let response = await this.$axios.get("/users/?username="+this.$route.params.username)
				this.profile = response.data;
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
			if (this.creatingMedia==false){
            this.creatingMedia = true;
			this.changingProfile = false;
			this.changingusername = false;}
			else{
            this.creatingMedia = false;}
        },
		
		updateProfile: async function(){
            if (this.changingProfile==false){
            this.changingProfile = true;
			this.changingusername = false;
			this.creatingMedia = false;}
			else{
            this.changingProfile = false;}
        },
		changeUsername: async function(){
            if (this.changingusername==false){
            this.changingusername = true;
			this.creatingMedia = false;
			this.changingProfile = false;}
			else{
            this.changingusername = false;}
        },
		seeMediaComments: async function(m){
            this.$router.push({ path: "/media/"+m.id+"/comments/"})
        },
		seeMyStream: async function(){
            this.$router.push({ path: "/stream"})
        },
	
		refresh() {
			this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
			this.GetProfile().then(() => this.GetUserMedia());
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
		async toggleBan(){
			if (this.profile.banned==true){
				this.unbanUser();
				this.profile.banned = false;
			}
			else{
				this.banUser();
				this.profile.banned = true;
			}
		},

		banUser() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.put("/users/:userid="+this.logged+"/banned/:bannedid="+this.profile.userid);
				this.$router.push({ path: '/stream/' })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
			this.refresh();
        },

		unbanUser() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.delete("/users/:userid="+this.logged+"/banned/:bannedid="+this.profile.userid);
				this.$router.push({ path: '/stream/' })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
			this.refresh();
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
		handleImageUpload(event) {
            this.previewImage(event)
            this.onFileUpload(event)
        },
        previewImage(event) {
            let input = event.target;
            if (input.files && input.files[0]) {
                let reader = new FileReader();
                reader.onload = (e) => {
                    this.preview = e.target.result;
                }
                reader.readAsDataURL(input.files[0]);
            }
        },
        onFileUpload (event) {
          this.photo = event.target.files[0]
        },
        async onSubmit() {
          // upload file
          this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
          const formData = new FormData()
          formData.append('pic', this.photo)
          formData.append('cap', this.caption)
          
		  try {
			await this.$axios.post("/users/"+ this.profile.userid+ "/media/", formData, {
          }).then((res) => {
            console.log(res)
          });
		  this.$router.push({ path: '/users/'+this.profile.username })
		  
            } catch (e) {
                this.errormsg = e.toString();
            }
			this.refresh();
			this.creatingMedia = false;
            this.loading = false;
        },
		async submitProfile() {
            this.loading = true;
            this.error = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let formData = new FormData();
            	formData.append('bio', this.profile.bio)
            	formData.append('username', this.profile.username)
                if (!this.$refs.newppic) return
                formData.append('pic', this.$refs.newppic.files[0]);
                await this.$axios.put("/users/:userid="+this.profile.userid, formData, {
                    headers: { 'Content-Type': 'multipart/form-data' }
                });
                this.$router.push({ path: '/users/'+this.profile.username })
            } catch (error) {
                this.error = error;
            }
			this.refresh();
            this.loading = false;
			this.changingProfile = false;
        },
		changename: async function () {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                await this.$axios.patch("/users/:userid="+this.profile.userid, {
					username: this.profile.username,})
                this.$router.push({ path: '/users/'+this.profile.username })
            } catch (e) {
                this.errormsg = e.toString();
            }
			this.refresh();
            this.loading = false;
			this.changingusername = false;
        },
		async getFollowers() {
            this.$router.push({ path: '/users/'+this.profile.username+'/followers/'})
        },
		async getFollowings() {
            this.$router.push({ path: '/users/'+this.profile.username+'/followings/'})
        },
		async getBanned() {
            this.$router.push({ path: '/users/'+this.profile.username+'/bans/'})
        }
    },
    mounted() {
		this.refresh();
    },
}
</script>

<template>
	<div class="Home">
		<NavBar :profilo="this.$route.params.username"/>
		
	</div>
    <header class="micio">
        <div class="header-content">
			<div class="user-title">
			{{this.profile.username}}
        </div>
        </div>
		<div class="profilepic">
		<img :src= this.profile.profilepic :width="300" :height="300" ><br />
        </div>
		<div class="user-numbers">
			<div class="row">
				<div class="column"><h3>{{ this.profile.nmedia }}</h3></div>
				<div class="column"><h3>{{this.profile.nfollowers}}</h3></div>
				<div class="column"><h3>{{this.profile.nfollowing}}</h3></div>
			</div>
			<div class="row">
				<div class="column">media</div>
				<div class="column"><button v-if="!loading"  @click="getFollowers">
        			followers
            		</button></div>
				<div class="column">
					<button v-if="!loading"  @click="getFollowings">
        			followings
            		</button>
				</div>
			</div>

        </div>
		<div class="buttons">
			<div class="buttons2" v-if= "this.profile.userid != logged">
				<button v-if="!loading" type="button" class="login-button" @click="toggle">
					{{profile.followed ? 'unfollow' : 'follow'}}
            	</button>
				<button v-if="!loading" type="button" class="login-button" @click="toggleBan">
					{{this.profile.banned ? 'unban' : 'ban'}}
            	</button>

			</div>
			<button v-if="!loading&&this.profile.userid == logged" type="button" class="login-button" @click="createMedia">
                Create new media
            </button>
			<button v-if="!loading&&this.profile.userid == logged" type="button" class="login-button" @click="updateProfile">
                Update profile
            </button>
			<button v-if="!loading&&this.profile.userid == logged" type="button" class="login-button" @click="changeUsername">
                change username
            </button>
			<button v-if="!loading&&this.profile.userid == this.logged" type="button" class="login-button" @click="getBanned">
                see banned users
            </button>
			
			
		</div>
    </header>

	<div v-if= "this.creatingMedia==true">
        <div class="newmedia">
          <form @submit.prevent="onSubmit">
			<div class="upload-space">
				<h2>create new media</h2>
              <div class="input">
				<input type="file" id="post-image" @change="handleImageUpload" accept="image/*">
				</div>
			  <div class="preview-space">
              <img id="preview-image" v-if="preview" :src="preview" :width="400" :height="400">
				</div> 
			</div>
			  <div class="form-group2">
                  <input type="text" v-model="caption" placeholder="Write a caption here" class="form-control">
              </div>
              <div class="form-group2">
                  <button class="login-button">Upload File</button>
              </div>
          </form>
		</div>
	</div>
	<div v-if= "this.changingusername==true">
        <div class="newusername">
          <form @submit.prevent="changename">

				<h2>change username</h2>


			  <div class="form-group2">
				<label for="description" class="form-label">Username</label>
            	<input type="text" class="form-control" id="Username" v-model="profile.username" placeholder= this.profile.username>
              
              </div>
              <div class="form-group2">
                  <button class="login-button">Upload File</button>
              </div>
          </form>
		</div>
	</div>
	<div v-if= "this.changingProfile==true">
        <div class="change-profile">
          <form @submit.prevent="submitProfile">
			<div class="upload-space">
				<h2>change profile</h2>
              <div class="input">
				<input type="file" id="newppic" name="newppic" ref="newppic" @change="handleImageUpload" accept="image/*">
				</div>
			  <div class="preview-space2">
              <img id="preview-image2" v-if="preview" :src="preview" :width="400" :height="400">
				</div> 
			</div>
			  <div class="form-group2">
				<label for="description" class="form-label">Username</label>
            	<input type="text" class="form-control" id="Username" v-model="profile.username" placeholder= this.profile.username>
              </div>
			  <div class="form-group2">
				<label for="description" class="form-label">Bio</label>
				  <input type="text" class="form-control" id="bio" v-model="profile.bio" placeholder= this.profile.bio >
              </div>
              <div class="form-group2">
                  <button class="login-button">change profile</button>
              </div>
          </form>
		</div>
		
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
					<img :src=m.photo :width="500" :height="500"><br />
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

</template>

<style>

.micio {
    height: 600px;
    padding-left: 10px;
    padding-right: 16px;
	padding-top: 10px;
    background-color:rgb(34, 145, 182);
    border-radius: 20px;

}
.newmedia {
    height: 650px;
    padding-left: 10px;
    padding-right: 16px;
	padding-top: 10px;
    background-color:rgb(34, 182, 41);
    border-radius: 20px;
	align-items: center;
	margin: auto;

}
.newusername {
    height: 200px;
    padding-left: 10px;
    padding-right: 16px;
	padding-top: 10px;
    background-color:rgb(34, 172, 182);
    border-radius: 20px;
	align-items: center;
	margin: auto;

}
.upload-space {
    height: 500px;
    padding-left: 10px;
    padding-right: 16px;
	padding-top: 20px;
    background-color:rgb(34, 135, 182);
    border-radius: 20px;
 	text-align-last:  center;
	margin:auto;

}
.change-profile {
    height: 800px;
    padding-left: 10px;
    padding-right: 16px;
	padding-top: 20px;
    background-color:rgb(34, 182, 39);
    border-radius: 20px;
 	text-align-last:  center;
	margin:auto;

}
.preview-space {
    height: 400px;
    border-radius: 20px;
	border: 1px solid hsl(0, 73%, 41%);
	margin-left: 600px;
	margin-top: -225px;

}
.preview-space img{
    display: flex;
 align-items: center;
	margin:auto;

}
.preview-space2 {
    height: 400px;
    border-radius: 20px;
	border: 1px solid hsl(0, 73%, 41%);
	margin-left: 600px;
	margin-top: -225px;

}
.preview-space2 img{
    display: flex;
 align-items: center;
	margin:auto;
	border-radius: 50pc;

}
.header-content {
	margin-top: 5px;
}
.user-title {
	font-size: 45px;
    text-align: center;
	margin:auto;
	color:beige;
}
.profilepic {
  padding: 2px;
  margin: auto;
  align-items: center;
}
.profilepic img {
  border: 2px solid #f4ba00;
  border-radius: 50pc;
  display: flex;
 align-items: center;
	margin:auto;


}
.input{
	width: 300px;
	height: 50px;
	margin-top: 200px;
	margin-left: 150px;
}
.user-numbers {
  height: 60px;
  width: 300px;
  align-items: center;
	margin:auto;
}
.buttons {
	height: 60px;
  margin-bottom: -200px;
  margin-top: 50px;
  width: 800px;
	margin:auto;
	background-color:rgb(34, 145, 182);
}
.buttons2 {
	height: 60px;
  margin-bottom: -200px;
  margin-top: 50px;
  width: 400px;
	margin:auto;
}
.column {
  padding: 2px;
  width: 100px;
  color:beige;
margin: auto;
  text-align: center;
}
.row:after {
  content: "";
  display: table;
  clear: both;
}
.short-profile-username {
    margin-top: 10px;
    margin-left: 9px;
    font-size: 16px;
    align-items: center;
    color: black
}
.profile-page{
    background-color: #02587b;
}
.card {
    margin-bottom: 20px;
}
.login-button{
    border-radius: 20px;
	align-items: center;
	margin-top: 20px;
    width: 200px;
    height: 40px;
    border: 1px solid rgb(255, 255, 255);
    background: #f4ba00;
    font-size: 10px;
    font-family: "Rubik", sans-serif;
    font-weight: 400;
    color: white;
    letter-spacing: 4px;
    text-transform: uppercase;
    text-decoration: none;
    border-radius: 25px;

}
.Home {
	max-width: 601px;
	margin-left: auto;
	margin-right: auto;
	padding-bottom: 10vh;
}
.sidebar {
	margin-top: 10px;
	background: #f4ba00;
	border: 2px solid #f4ba00;
}
@media (--t) {
	.sidebar {
		display: block;
		margin-top: 16px;
		border: 2px solid #f4ba00;
	}
	.sidebar p {
		position: sticky;
		top: calc(53px + 30px + 18px);
	}
}
</style>
