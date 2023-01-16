// Short profile component
// It is a short version of the user profile with just profile pic and username

<script>
export default {
    name: 'ShortProfile',
    props: {
        pic: { 
            type: String,
            default:
                // select default profile pic here??
                ""
        },
        size:{
            type: Number,
            default: 50
        },
        username:{
            type: String,
            default: 'unknown'
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
    }
}
</script>

<template>
    <header class="header">
        <div class="header-content">
            <figure class="profilePic">
                <img :src="pic" :width="size" :height="size" />
                <div class="short-profile-username">
                    <button v-if="!loading" class="miao" @click="ToProfile(this.username)">
        {{ this.username }}
            	</button>
                </div>
                
            </figure>
        </div>
    </header>
</template>

<style>
.header {
    display: flex;
    align-items: center;
    height: 60px;
    padding-left: 10px;
    padding-right: 16px;
    background-color:rgb(237, 246, 249);
    border-radius: 20px;

}
.miao {
   color:rgb(237, 246, 249);
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
  margin-top: 14px;
}
.profilePic img {
  border: 2px solid #f4ba00;
  border-radius: inherit;
}
.short-profile-username {
    margin-top: 10px;
    margin-left: 9px;
    font-size: 16px;
    align-items: center;
    color: black;
    background-color: rgb(237, 246, 249);
}

</style>