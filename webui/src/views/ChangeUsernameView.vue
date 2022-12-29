<script>
export default {
    data: function() {
        return {
            errormsg: null,
            loading: false,

            profile:"",

        }
    },
    methods: {
        ChangeUsername: async function () {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.patch("/users/:userid="+this.profile.userid, {
					username: this.profile.username,})
                //this.$router.push({ path: '/users/'+this.profile.username })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

        async GetProfile() {
            this.loading = true;
            this.errormsg = null;
            try {
                this.$axios.get("/users/?username="+this.$route.params.username).then(response => (this.profile = response.data))
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        }
    },
    mounted() {
        this.GetProfile()
    }
}
</script>

<template>
        <div
            class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
            <h1 class="h2">Change Username</h1>
        </div>

        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

        <div class="mb-3">
            <label for="description" class="form-label">Username</label>
            <input type="text" class="form-control" id="Username" v-model="profile.username" placeholder= this.profile.username>
        </div>
    

        <div>
            <button v-if="!loading" type="button" class="btn btn-primary" @click="ChangeUsername">
                Change
            </button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
        </div>

        
</template>

<style>
</style>