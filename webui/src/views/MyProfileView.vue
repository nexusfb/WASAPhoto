
<script>
export default {
    props: [`username`],
    data: function() {
        return {
            loading : false,
            errmsg : null,
            profile:"",

        }
    },
    methods: {
        async GetProfile() {
            this.loading = true;
            this.errormsg = null;
            try {
                this.$axios.get("/users/?username="+this.$route.params.username).then(response => (this.profile = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        createMedia: async function(){
            this.$router.push({ path: '/users/'+this.profile.username+"/newMedia"})
        }
    },
    mounted() {
        this.GetProfile();
    }
}
</script>

<template>
    <div>
        <h1> Profile</h1>
            </div>
            <div class="card-body">
                <p class="card-text">
                    id: {{ this.profile.userid }}<br />
                    name: {{ this.profile.username }}<br />
                    bio: {{ this.profile.bio }}<br />
                    media: {{ this.profile.nmedia }}<br />
                    followers: {{ this.profile.nfollowers}}<br />
                    followings: {{ this.profile.nfollowing}}

                </p>
            </div>
            <div>
            <button v-if="!loading" type="button" class="btn btn-primary" @click="createMedia">
                Create new media
            </button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
        </div>
</template>

<style scoped>
.card {
    margin-bottom: 20px;
}
</style>
