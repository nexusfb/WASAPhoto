// View of WASA Photo login page

<script>

export default {

    data: function() {
        return {
            errormsg: null,
            loading: false,
            // data of the logged user
            User: {
                userid: "",
                username: "",
            }
        }
    },
    methods: {
        LoginUser: async function () {
            this.loading = true;
            this.errormsg = null;
            try {
                // backend method to login
                let response = await this.$axios.post("/session/", {
                    username: this.User.username,
                });
                // save the response in data
                this.User.userid  = response.data,
                // use view data to set the authorization item in the local storage
                localStorage.setItem('Authorization', this.User.userid),
                // redirect to stream page (default view)
                //this.$router.push({ path: '/stream/' })
                this.$router.push({ path: '/users/'+this.User.username })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        
    }
}
</script>

<template>
    <div class="background">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="login-container">
            <h1> WASA Photo </h1>
            <div class="register">
                <input type="text" v-model="User.username" placeholder="username">
                <button v-if="!loading" class="login-button" type="button" @click="LoginUser">LOGIN</button>
                <LoadingSpinner v-if="loading"></LoadingSpinner>
            </div>
        </div>
    </div>
    
</template>

<style scoped>
.background {
  background-color: #02587b;
  margin: -10px;
  height: 100vh;
  display: flex;
  flex-direction: column;
  text-align: center;
}
.background div{
    margin-right: auto;
    margin-left: auto;
}
.logo {
    width: 100px;
    margin-top: 50px;
}
.login-container {
    margin-top: 350px;
    width: 300px;
    height: 200px;
    background-color:  #e0e8eb;
    position: relative;
    border-radius: 25px;
}
.login-container h1 {
    margin-top: 15px;
}
.register input{
    position: absolute;
    top: 66px;
    left: 45px;
    width: 70%;
    height: 40px;
    padding-left: 15px; 
    border: 1px solid hsl(0, 0%, 100%);
}
.login-button{
    border-radius: 20px;
    margin-top: 80px;
    width: 70%;
    height: 40px;
    border: 1px solid rgb(255, 255, 255);
    background: #f4ba00;
    font-size: 14px;
    font-family: "Rubik", sans-serif;
    font-weight: 400;
    color: white;
    letter-spacing: 4px;
    text-transform: uppercase;
    text-decoration: none;
    border-radius: 25px;
}

</style>