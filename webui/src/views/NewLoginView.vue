// this is the login view where user can insert username to log back in or create a nw account
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
        // given input string login user
        LoginUser: async function () {
            this.loading = true;
            this.errormsg = null;
            // emptiness check
            if (this.User.username.length == 0) {
                    this.errormsg = "Error: empty usernames are not valid. Please try again."}
            else{
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
                if (e.response && e.response.status == 400){
                this.errormsg = "Error: the inserted username is invalid. Please try again."}
                if (e.response && e.response.status == 500){
                this.errormsg = "Error: internal error. Please try again."}
            }}
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
                <input type="text" v-model="User.username" placeholder="enter username">
                <button v-if="!loading" class="login-button" type="button" @click="LoginUser">LOGIN</button>
                <LoadingSpinner v-if="loading"></LoadingSpinner>
            </div>
        </div>
    </div>
    
</template>

<style scoped>
.background {
  background-color: #160F29;
  margin: -50px;
  height: 110vh;
  display: flex;
  flex-direction: column;
  text-align: center;
}
.background div{
    margin-right: auto;
    margin-left: auto;
}
.login-container {
    margin-top: 370px;
    width: 400px;
    height: 300px;
    background-color:  #DDBEA8;
    position: relative;
    border-radius: 25px;
}
.login-container h1 {
    margin-top: 80px;
    color:#160F29;
    font-family: "Copperplate";
    text-transform: uppercase;
    
}
.register input{
    position: relative;
    top: 10px;
    left: 5px;
    width: 70%;
    height: 40px;
    padding-left: 15px; 
    border: 1px solid #DDBEA8;
    background-color:  #f3dfc1;
    color:#160F29;
}
.login-button{
    position: relative;
    border-radius: 20px;
    margin-top: 20px;
    width: 50%;
    height: 40px;
    border: 1px solid rgb(255, 255, 255);
    background: #368F8B;
    font-size: 15px;
    font-family: "Copperplate";
    font-weight: 400;
    color: #f3dfc1;
    letter-spacing: 4px;
    text-transform: uppercase;
    text-decoration: none;
    border-radius: 25px;
}
</style>