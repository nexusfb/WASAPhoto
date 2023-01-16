// Page of user profile with username, caption, profile picture, nmedia, nfollowers, nfollowing, user media
<script>
import NewMedia from "@/components/NewMedia.vue";
export default {
    components: { NewMedia },
    data: function() {
        return {
            loading : false,
            errmsg : null,
            profile: this.GetProfile(),
			media:[],
			logged: localStorage.getItem('Authorization'),
			comment: "",
			creating: false,
			photo: "",
         caption:"",
         preview: "",
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
			if (this.creating==false){
            this.creating = true;}
			else{
            this.creating == false;}
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
			this.$axios.post("/users/"+ this.profile.userid+ "/media/", formData, {
          }).then((res) => {
            console.log(res)
          })
				await this.refresh();
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
    },
    mounted() {
		this.refresh();
    },
}
</script>

<template>

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
				<div class="column">followers</div>
				<div class="column">{{ this.creating }}</div>
			</div>

        </div>
		<div class="buttons">
			<button v-if="!loading&&this.profile.userid == logged" type="button" class="login-button" @click="createMedia">
                Create new media
            </button>
			<button v-if="!loading&&this.profile.userid == logged" type="button" class="login-button" @click="updateProfile">
                Update profile
            </button>
			<button v-if="!loading&&this.profile.userid == logged" type="button" class="login-button" @click="changeUsername">
                change username
            </button>
			<button v-if="!loading" type="button" class="login-button" @click="searchUsers">
                search users
            </button>
			<button v-if="!loading&&this.profile.userid == this.logged" type="button" class="login-button" @click="seeMyStream">
                see my stream
            </button>
			<div v-if= "this.profile.userid != logged">
				<button v-if="!loading" type="button" class="login-button" @click="toggle">
					{{profile.followed ? 'unfollow' : 'follow'}}
            	</button>

			</div>

			<div v-if= "this.profile.userid != logged">
				<button v-if="!loading" type="button" class="login-button" @click="toggleBan">
					{{this.profile.banned ? 'unban' : 'ban'}}
            	</button>

			</div>
		</div>
    </header>
	<div v-if= "this.creating==true">
        <div class="newmedia">
          <form @submit.prevent="onSubmit">
			<div class="upload-space">
				<h2>create new media</h2>
              <div class="form-group">
				<label for="fileupload">Choose photo to post</label>
                  <input id="fileupload"  type="file" accept="image/png, image/jpeg" @change="handleImageUpload">
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


			<div v-if= "this.profile.userid != logged">
				<button v-if="!loading" type="button" class="btn btn-primary" @click="toggle">
					{{profile.followed ? 'unfollow' : 'follow'}}
            	</button>

			</div>

			<div v-if= "this.profile.userid != logged">
				<button v-if="!loading" type="button" class="btn btn-primary" @click="toggleBan">
					{{this.profile.banned ? 'unban' : 'ban'}}
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
    background-color:rgb(182, 34, 135);
    border-radius: 20px;

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
.preview-space {
    height: 400px;
    border-radius: 20px;
	margin-left: 600px;
	margin-top: -225px;

}
.preview-space img{
    display: flex;
 align-items: center;
	margin:auto;

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
.form-group{
	width: 300px;
	height: 50px;
    background-color:#f4ba00;
    border-radius: 20px;
	margin-top: 200px;
	margin-left: 150px;
	border: 1px solid rgb(255, 255, 255);
}
.user-numbers {
  height: 60px;
  width: 300px;
  align-items: center;
	margin:auto;
}
.buttons {
  margin-left: 200px;
  margin-right: 200px;
  margin-bottom: 200px;
  background-color:rgb(182, 34, 34);
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
    margin-top: 8px;
	align-items: center;
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
</style>
