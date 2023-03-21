// Page of user profile with username, caption, profile picture, nmedia, nfollowers, nfollowing, user media
<script>
import FinalMedia from "@/components/FinalMedia.vue";
import NavBar from "@/components/NewHomeBar.vue";
export default {
    components: { FinalMedia, NavBar},
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
         	photoCaption:"",
         	preview: "",
			newppic:null,
            new:"",

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
                this.errormsg = e.
                toString();
				
            }
            this.loading = false;
        },
        async deleteProfile() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.delete("/users/:userid="+this.logged).then(this.logout)
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
        logout: async function(){
            this.$router.push({ path: '/'})
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
	
		async refresh() {
			this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
			await this.GetProfile()
            if (this.profile.profilepic){await this.getImage(this.profile.profilepic)}
            this.GetUserMedia();
            this.new = this.profile.username
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
                this.$axios.put("/users/:userid="+this.logged+"/followings/:followingid="+this.profile.userid).then(() => (this.refresh()));
				// this.$router.push({ path: '/users/'+this.profile.username });
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

		unfollowUser() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.delete("/users/:userid="+this.logged+"/followings/:followingid="+this.profile.userid).then(() => (this.refresh()));
				//this.$router.push({ path: '/users/'+this.profile.username })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
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
            if (this.comment.length==0){
                if (e.response && e.response.status == 400){
                this.errormsg = "Error: the inserted username is invalid. Try again."}
            }
            else{
            try {
                this.$axios.post("/media/:mediaid="+ m.id+"/comments/", {
					content: this.comment,});
				this.refresh();
            } catch (e) {
                this.errormsg = e.toString();
            }
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
          formData.append('cap', this.photoCaption)
          if (!this.photo){
                this.errormsg = "Error: you must provide a picture. Please try again."
                return
            }
          else{
		  try {
			this.$axios.post("/users/"+ this.profile.userid+ "/media/", formData, {
          }).then(() => (this.refresh()));
		  
            } catch (e) {
                this.errormsg = e.toString();
            }}
			this.refresh();
			this.creatingMedia = false;
            this.loading = false;
        },
		async submitProfile() {
            this.loading = true;
            this.error = null;
            
            if (this.new.length<5){
                this.errormsg = "Error: usernames have to be at least 5 characters long. Please try again."
                return
            }
            if (this.new.length>20){
                this.errormsg = "Error: usernames can have a maximum lenght of 20 characters. Please try again."
                return
            }
            if (!this.$refs.newppic.files[0]){
                this.errormsg = "Error: you must provide a profile picture. Please try again."
                return
            }
            try {
                let formData = new FormData();
            	formData.append('bio', this.profile.bio)
            	formData.append('username', this.new)
                formData.append('pic', this.$refs.newppic.files[0]);
                await this.$axios.put("/users/:userid="+this.profile.userid, formData, {
                    headers: { 'Content-Type': 'multipart/form-data' }
                });
                this.$router.push({ path: '/users/'+this.new }).then(() => (this.refresh()));
            } catch (e) {
                if (e.response && e.response.status == 400){
                this.errormsg = "Error: the inserted profile is invalid. Please try again."
                }

            }
            this.refresh = true;
            this.loading = false;
			this.changingProfile = false;
        },
		changename: async function () {
            this.loading = true;
            this.errormsg = null;
            if (this.new.length<5){
                this.errormsg = "Error: usernames have to be at least 5 characters long. Please try again."
                return
            }
            if (this.new.length>20){
                this.errormsg = "Error: usernames can have a maximum lenght of 20 characters. Please try again."
                return
            }
            try {
                this.$axios.patch("/users/:userid="+this.profile.userid, {
					username: this.new,})
                this.$router.push({ path: '/users/'+this.new }).then(() => (this.refresh()));
            } catch (e) {
                if (e.response && e.response.status == 400){
                this.errormsg = "Error: the inserted username is invalid. Please try again."}
                if (e.response && e.response.status == 500){
                this.errormsg = "Error: internal error. Please try again."}
            }
			this.refresh();
            this.loading = false;
			this.changingusername = false;
        },
		async getFollowers() {
            this.$router.push({ path: '/users/'+this.profile.userid+'/followers/', props: true})
        },
		async getFollowings() {
            this.$router.push({ path: '/users/'+this.profile.userid+'/followings/'})
        },
		async getBanned() {
            this.$router.push({ path: '/users/'+this.profile.username+'/bans/'})
        },
        async getImage(p) {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.get("/images/?image_name=" + p, { responseType: 'blob' })
                // Get the image data as a Blob object
                var imgBlob = response.data;
                // Create an object URL from the Blob object
                this.profile.profilepic = URL.createObjectURL(imgBlob);
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
        },
    },
    mounted() {
        this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
            error => { return Promise.reject(error); });
		this.refresh();
    },
}
</script>

<template>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    <div class="page">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	<div class="Bar">
		<NavBar :profilo="this.$route.params.username"/>
	</div>
    <header class="summary_page">
        <div class="header-content">
			<div class="user-title">
			{{this.profile.username}}
            </div>
        </div>
        <button v-if="this.profile.userid==logged" class="button1" @click="deleteProfile">Delete Profile</button>
            <button v-if="this.profile.userid==logged"  class="button2" @click="logout">logout</button>
		<div class="profilepic">
		<img :src= this.profile.profilepic :width="300" :height="300" ><br />
        </div>
        <div class="user-bio">
			{{this.profile.bio}}
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
			<div class="generalbuttons" v-if= "this.profile.userid != logged">
				<button v-if="!loading" type="button" class="login-button" @click="toggle">
					{{profile.followed ? 'unfollow' : 'follow'}}
            	</button>
				<button v-if="!loading" type="button" class="login-button" @click="toggleBan">
					{{this.profile.banned ? 'unban' : 'ban'}}
            	</button>
			</div>
            <div class="myprofilebuttons" v-if= "this.profile.userid == logged">
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
		</div>
    </header>

	<div v-if= "this.creatingMedia==true">
        <div class="newmedia">
          <form @submit.prevent="onSubmit">
			<div class="upload-space">
				<h2>create new media</h2>
              <div class="input">
                <h3>pick an image</h3>
				<input type="file" id="post-image" @change="handleImageUpload" accept="image/*">
			  </div>
			  <div class="preview-space">
                <img id="preview-image" v-if="preview" :src="preview" :width="400" :height="400">
			  </div> 
			  <div class="form-caption">
                  <input type="text" v-model="photoCaption" placeholder="Write a caption here" class="form-control">
                  <button class="login-button">Share</button>
              </div>
            </div>
          </form>
		</div>
	</div>
	<div v-if= "this.changingusername==true">
        <div class="newusername">
            <div class="upload-space-username">
				<h2>change username
                </h2>
			  <div class="form-username">
				<h3>insert username</h3>
            	<input type="text" v-model="this.new" placeholder= this.profile.username>
                <button class="login-button" @click="changename">change</button>
              </div>
            </div>
		</div>
	</div>
	<div v-if= "this.changingProfile==true">
        <div class="newmedia">
          <form @submit.prevent="submitProfile">
			<div class="upload-space">
			    <h2>change profile</h2>
                <div class="input">
                    <h3>pick an image</h3>
				    <input type="file" id="newppic" name="newppic" ref="newppic" @change="handleImageUpload" accept="image/*">
				</div>
			    <div class="preview-space2">
                    <img id="preview-image2" v-if="preview" :src="preview" :width="400" :height="400">
				</div>
			  <div class="form-profile">
				<label for="description" class="form-label">Username</label>
            	<input type="text" class="form-control" id="Username" v-model="this.new" placeholder= this.profile.username>
				<label for="description" class="form-label">Bio</label>
				<input type="text" class="form-control" id="bio" v-model="profile.bio" placeholder= this.profile.bio >
                <button class="login-button">change profile</button>
              </div>
            </div>
          </form>
		</div>
		
	</div>

            <LoadingSpinner v-if="loading"></LoadingSpinner>


			<div class="item-error" v-if="!media">
            <h2>No media has been posted yet!</h2>
            </div>

			<header class="summary_page_b">
		 <div class="stream_timeline">
		 <FinalMedia	v-on:refresh-parent="refresh" v-for="post in media" :key="post.id" :pp="post.authorpic" :photoId="post.id"
				:owner="post.author" :image="post.photo"
				:timestamp="post.date" :caption="post.caption" :likesCount="post.nlikes"
				:commentsCount="post.ncomments" :liked="post.liked" :logged="this.logged" :authorid="post.authorid"/>
	 </div>
	 </header>
    </div>

</template>

<style>
.page{
background-color: #160F29;
  margin: -50px;
  display: flex;
  flex-direction: column;
  text-align: center;
}
.Bar {
    position: relative;
    margin-top:50px;
	max-width: 601px;
	margin-left: 120px;
}
.summary_page {
    position: relative;
    margin-top: 50px;
    height: 600px;
    padding-left: 10px;
    padding-right: 16px;
    background-color:#246A73;
    border-radius: 20px;
}

.header-content {
	margin-top: 5px;
}
.user-title {
	font-size: 45px;
    text-align: center;
	margin:auto;
	color:beige;
    font-family: "Copperplate";
    text-transform: uppercase;
}
.profilepic {
  padding: 2px;
  margin: auto;
  align-items: center;
}
.profilepic img {
  border: 2px solid #1b5158;
  border-radius: 50pc;
  display: flex;
 align-items: center;
	margin:auto;
}
.user-bio{
    font-size: 20px;
    text-align: center;
	margin:auto;
	color:beige;
    font-family: "Copperplate";
}
.user-numbers {
  height: 60px;
  width: 300px;
  align-items: center;
  margin-top:20px;
  margin-left: 710px;
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
.column button{
    background-color:#246A73;
    color: beige;
}
.buttons {
	height: 60px;
  margin-bottom: -200px;
  margin-top: 50px;
  width: 800px;
	margin:auto;
    background-color:#246A73;
}
.generalbuttons {
  height: 60px;
  margin-bottom: -200px;
  margin-top: 50px;
  width: 800px;
	margin:auto;
}
.generalbuttons button {
  background-color: #1b5158;
  border-color: beige;
  color: beige;

}
.button1{
    background-color: rgb(139, 14, 2);
    margin-left: 550px;
    margin-top: -45px;
    position: absolute;
    color:beige;
    font-family: "Copperplate";
    font-size: 17px;
    width: 140px;
    border-radius: 20px;

}
.button2{
    background-color: rgb(66, 2, 139);
    margin-left: 700px;
    margin-top: -45px;
    position: absolute;
    color:beige;
    font-family: "Copperplate";
    font-size: 17px;
    width: 100px;
    border-radius: 20px;

}
.myprofilebuttons {
  height: 60px;
  margin-bottom: -200px;
  margin-top: 50px;
  width: 800px;
	margin:auto;
}
.myprofilebuttons button {
  background-color: #1b5158;
  border-color: beige;
  color: beige;

}
.newmedia {
    position: relative;
    height: 650px;
    margin-left: 20px;
    padding-left: 70px;
	padding-top: 10px;
    border-radius: 20px;
	align-items: center;
	margin: auto;
    background-color: #160F29;
}
.upload-space {
    height: 600px;
    padding-left: 10px;
    padding-right: 10px;
	padding-top: 20px;
    background-color:#DDBEA8;
    border-radius: 20px;
 	text-align-last:  center;
	margin:auto;

}
.upload-space h2{
    font-size: 35px;
    text-align: center;
	margin:auto;
	color:beige;
    font-family: "Copperplate";
    text-transform: uppercase;
}
.input{
	width: 300px;
	height: 50px;
	margin-top: 200px;
	margin-left:450px;
}
.input h3{
    position: absolute;
    font-size: 20px;
    text-align: center;
	color:beige;
    font-family: "Copperplate";
    text-transform: uppercase;
    margin-left: 30px;
    margin-top: 245px;

}
.input input {
    position: absolute;
    width: 200px;
    background-color: #e8d1c1;
    color: #e8d1c1;
    margin-left: -140px;
    margin-top: 270px;
}
.preview-space {
    position:relative;
    height: 400px;
    width: 400px;
	border: #e8d1c1;
	margin-left: 350px;
	margin-top: -220px;

}
.form-caption{
    position: relative;
    margin-top: -300px;
    margin-left: 900px;
    width: 600px;
}
.form-caption input{
    height: 100px;
    font-size: 15px;
}
.form-caption button{
    font-size: 20px;
    text-align: center;
	color:beige;
    font-family: "Copperplate";
    text-transform: uppercase;
    background-color: #246A73;
}
.newusername {
    position: relative;
    height: 270px;
    margin-left: 20px;
    padding-left: 70px;
	padding-top: 10px;
    border-radius: 20px;
	align-items: center;
	margin: auto;
    background-color: #160F29;

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
.preview-space img{
    display: flex;
 align-items: center;
	margin:auto;

}
.preview-space2 {
    position:relative;
    height: 400px;
    width: 400px;
	border: #e8d1c1;
	margin-left: 350px;
	margin-top: -220px;

}
.preview-space2 img{
    display: flex;
 align-items: center;
	margin:auto;
	border-radius: 50pc;

}
.form-profile{
    position: relative;
    margin-top: -300px;
    margin-left: 900px;
    width: 600px;
}
.form-profile label {
    font-size: 20px;
    text-align: center;
	color:beige;
    font-family: "Copperplate";
    text-transform: uppercase;
}
.form-profile input{
    margin-top: -5px;
    height: 50px;
    font-size: 15px;
}
.form-profile button{
    width: 300px;
    font-size: 20px;
    text-align: center;
	color:beige;
    font-family: "Copperplate";
    text-transform: uppercase;
    background-color: #246A73;
}
.upload-space-username {
    height: 250px;
    padding-left: 10px;
    padding-right: 10px;
	padding-top: 20px;
    background-color:#DDBEA8;
    border-radius: 20px;
 	text-align-last:  center;
	margin:auto;
}
.upload-space-username h2{
    font-size: 35px;
    text-align: center;
	margin:auto;
	color:beige;
    font-family: "Copperplate";
    text-transform: uppercase;
}
.form-username{
    position: absolute;
    margin-top: 10px;
    margin-left: 550px;
    width: 600px;
}
.form-username input{
    position: relative;
    height: 50px;
    width: 300px;
    font-size: 15px;
    margin-top: 30px;
    margin-left: -10px;
}
.form-username button{
    font-size: 20px;
    text-align: center;
	color:beige;
    font-family: "Copperplate";
    text-transform: uppercase;
    background-color: #246A73;
    margin-top: -5px;
    margin-left: 20px;
}
.form-username h3{
    position: absolute;
    font-size: 20px;
    text-align: center;
	color:beige;
    font-family: "Copperplate";
    text-transform: uppercase;
    margin-left: 180px;

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
