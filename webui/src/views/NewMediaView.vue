<script>
export default {
    props: [`username`],
    data: function() {
      return {
         photo: "",
         caption:"",
         profile: "",
      }
    },
    methods: {
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
          this.photo = event.target.files[0]
        },
        onSubmit() {
          // upload file
          this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
          const formData = new FormData()
          formData.append('pic', this.photo)
          formData.append('cap', this.caption)
          this.$axios.post("/users/"+ this.profile.userid+ "/media/", formData, {
          }).then((res) => {
            console.log(res)
          })
          this.$router.push({ path: '/users/'+this.profile.username });
        },
    },
    mounted() {
        this.GetProfile();
    }
}
</script>

<template>
    <div>
      <div class="container">
          <form @submit.prevent="onSubmit">
              <div class="form-group">
                  <input type="file" @change="onFileUpload">
              </div>
              <div class="form-group">
                  <input type="text" v-model="caption" placeholder="Write a caption here" class="form-control">
              </div>
              <div class="form-group">
                  <button class="btn btn-primary btn-block btn-lg">Upload File</button>
              </div>
          </form>
      </div>    
    </div>
  </template>

<style>
</style>