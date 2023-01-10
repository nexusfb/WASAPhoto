<script>
 
// to do: mettere in automatico foto vecchia -> se metto this.file = this.profilepic non va :( ma se faccio this.profile.profilepic si WHYY)
export default {
    data: function() {
        return {
            errormsg: null,
            loading: false,
            profile:"",
            file: "ciao"
        }
    },
    methods: {
        onFileChange(e) {
            const selectedFile = e.target.files[0]; // accessing file
            this.profile.profilepic = selectedFile;
        },
        selectImage () {
            this.$refs.fileInput.click()

        },

        pickFile () {

            let input = this.$refs.fileInput

            let file = input.files

            if (file && file[0]) {

                let reader = new FileReader

                reader.onload = e => {

                this.previewImage = e.target.result }

                reader.readAsDataURL(file[0])

                this.$emit('input', file[0])

            }

        },

        UpdateProfile: async function () {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            let formData = new FormData();
            formData.append('pic', this.profile.profilepic)
            formData.append('bio', this.profile.bio)
            formData.append('username', this.profile.username)
            try {
                this.$axios.put("/users/:userid="+this.profile.userid, formData)
                this.$router.push({ path: '/users/'+this.profile.username })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

        async GetProfile() {
            this.loading = true;
            this.errormsg = null;
            try {
                this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
                this.$axios.get("/users/?username="+this.$route.params.username).then(response => (this.profile = response.data))
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
            
        },
        onFileUpload (event) {
          this.file = event.target.files[0]
        },
        onSubmit() {
          // upload file
          this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
          const formData = new FormData()
          formData.append('pic', this.file)
            formData.append('bio', this.profile.bio)
            formData.append('username', this.profile.username)
          this.$axios.put("/users/:userid="+this.profile.userid, formData, {
          }).then((res) => {
            console.log(res)
          })
          this.$router.push({ path: '/users/'+this.profile.username })
        },
    },
    mounted() {
        this.GetProfile();
    }
}
</script>

<template>


        <div
            class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
            <h1 class="h2">Update profile</h1>
        </div>

        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

        <div class="mb-3">
            <label for="description" class="form-label">Username</label>
            <input type="text" class="form-control" id="Username" v-model="profile.username" placeholder= this.profile.username>
        </div>
        <div class="mb-3">
            <label for="description" class="form-label">Bio</label>
            <input type="text" class="form-control" id="bio" v-model="profile.bio" placeholder=this.profile.bio>
        </div>

        <label for="description" class="form-label">profile</label>
        <div class="container">
        <form @submit.prevent="onSubmit">
            <div class="form-group">
                <input type="file" @change="onFileUpload">
            </div>
            <div class="form-group">
                <button class="btn btn-primary btn-block btn-lg">Upload File</button>
            </div>
        </form>
    </div> 

        
        
</template>

<style>
</style>