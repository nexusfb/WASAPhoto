// Short profile component
// It is a short version of the user profile with just profile pic and username

<script>
export default {
    name: 'ShortProfile',
    props: {
        pic: { 
            type: String,
            default:
                ""
        },
        size:{
            type: Number,
            default: 50
        },
        username:{
            type: String,
            default: 'unknown'
        },
    },
    data: function () {
        return {
            sp: "",
        }
    },
    methods: {
        async ToProfile(name) {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$router.push({ path: '/users/'+name })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async getImage() {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.get("/images/?image_name=" + this.pic, { responseType: 'blob' })
                // Get the image data as a Blob object
                var imgBlob = response.data;
                // Create an object URL from the Blob object
                this.sp = URL.createObjectURL(imgBlob);
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
        },
    },
    mounted() {
        if (this.pic) {
            this.getImage()
        }
    },
}
</script>

<template>
    <header class="header">
        <div class="header-content">
            <figure class="profilePic">
                <img :src=this.sp :width="size" :height="size" />
                <div class="short-profile-username">
                    <button v-if="!loading" class="sp-border" @click="ToProfile(this.username)">
        {{ this.username }}
            	</button>
                </div>
                
            </figure>
        </div>
    </header>
</template>

<style>
.header {
    margin-top: 10px;
    display: flex;
    align-items: center;
    height: 60px;
    padding-left: 10px;
    padding-right: 16px;
    background-color:#DDBEA8;
    border-radius: 30px ;


}
.sp-border {
   border: 2px solid #ffffff;
   border-radius: 20px;
   padding-left: 10px;
    padding-right: 10px;
}
.header-content {
    display: flex;
    align-items: center;
    margin-left: 2px;
}
.profilePic {
  padding: 2px;
  border-radius: 50%;
  display: inline-flex;
  margin-top: 10px;
}
.profilePic img {
  border: 2px solid #2b1e4f;
  border-radius: inherit;
}
.short-profile-username {
    margin-top: 10px;
    margin-left: 9px;
    font-size: 16px;
    align-items: center;
    color: black;
}
.short-profile-username button {

    margin-left: 9px;
    font-size: 16px;
    align-items: center;
    color: beige;
    background-color: #2b1e4f;
    font-family: "Copperplate";
    text-transform: uppercase;
    border-color: #2b1e4f;
}

</style>