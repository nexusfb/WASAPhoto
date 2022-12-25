<script>
export default {
    data: function() {
        return {
            errormsg: null,
            detailedmsg: null,
            loading: false,
            id : 10,
            User: {
                UserID: null,
                Username: null,

            }
        }
    },
    methods: {
        LoginUser: async function () {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.post("/session/", {
                    username: this.Username,
                });
				this.UserID  = response.data,
                localStorage.setItem('Authorization', this.UserID),
                this.$router.push({ path: '/users/'+this.Username })

            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        }


    }
}
</script>

<template>
    <div>
        <div
            class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
            <h1 class="h2">Login</h1>
        </div>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

        <div class="mb-3">
            <label for="description" class="form-label">Username</label>
            <input type="string" class="form-control" id="username" v-model="Username" placeholder="please insert username">
        </div>

        <div>
            <button v-if="!loading" type="button" class="btn btn-primary" @click="LoginUser">
                Login
            </button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
        </div>
    </div>

</template>

<style>
</style>