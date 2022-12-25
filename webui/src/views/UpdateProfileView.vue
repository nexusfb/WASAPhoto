<script>
export default {
    props: [`userid`],
    data: function() {
        return {
            errormsg: null,
            loading: false,

            bio: "a",
            profilepic: "b",
            username: "c",

        }
    },
    methods: {
        UpdateProfile: async function () {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.put("/users/:userid="+this.$route.params.userid, {
					bio: this.bio,
					profilepic: this.profilepic,
					username: this.username,})
                this.$router.push({ path: '/users/'+this.username })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

        async GetProfile() {
            this.loading = true;
            this.errormsg = null;
            try {
                this.$axios.get("/users/?username="+this.$route.params.username).then(response => (this.userid = response.userid));
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
    <div>
        <div
            class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
            <h1 class="h2">Update profile</h1>
        </div>

        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

        <div class="mb-3">
            <label for="description" class="form-label">Username</label>
            <input type="text" class="form-control" id="username" v-model="username" placeholder="new username">
        </div>
        <div class="mb-3">
            <label for="description" class="form-label">Bio</label>
            <input type="text" class="form-control" id="bio" v-model="bio" placeholder="what are you? describe yourself!">
        </div>
        <div class="mb-3">
            <label for="description" class="form-label">Profile pic</label>
            <input type="text" class="form-control" id="profilepic" v-model="profilepic" placeholder="change profile pic">
        </div>

        <div>
            <button v-if="!loading" type="button" class="btn btn-primary" @click="UpdateProfile">
                Update Profile
            </button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
        </div>

        <div class="card-body">
                <p class="card-text">
                    bio: {{ this.username }}<br />
                    name: {{ this.bio }}<br />
                    foto: {{ this.profilepic }}

                </p>
            </div>
        
    </div>
</template>

<style>
</style>