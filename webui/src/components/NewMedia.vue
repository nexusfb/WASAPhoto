<script>
export default {
	name: 'NewMedia',
    props: {
        pic: { 
            type: String,
            default:""
        },
        caption:{
            type: String,
            default: ""
        },

    },
    data: {
        photo: "",
        caption:"",
        preview: "",
    },
    methods: {
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
          });	this.creating = false;
				await this.refresh();
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
    }
}
</script>

<template>
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
  width: 1000px;
	margin:auto;
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
</style>
