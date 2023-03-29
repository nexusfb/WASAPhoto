// this is the followers view where all the user followed by a profile are listed
<script>
import { ref } from 'vue';
import ShortProfile from "@/components/ShortProfile.vue"
import NavBar from "@/components/NewHomeBar.vue";
export default {
    props:['username'],
    components: {
        ShortProfile,
        NavBar
    },
    data: function() {
        return {
            loading : false,
            errmsg : null,
            input: ref(""),
            profile:"",
			users:[],
        }
    },
    methods: {
        // get followers list
        async GetUserList() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.get("/users/:userid="+this.$route.params.username+"/followers/").then(response => (this.users = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        // go to profile
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

        filteredList() {
        return this.users.filter((user) => user.username.toLowerCase().includes(this.input.toLowerCase()) );

        
    },
},
	mounted() {
		this.GetUserList();
        this.filteredList();
	}
}
</script>
<template>
   <div class="page_b">
	<div class="Bar_b">
		<NavBar :profilo="this.$route.params.username"/>
	</div>
    <header class="summary_page_b">
        <h3>FOLLOWERS</h3>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="item-user2" v-for="user in filteredList()" :key="user">
    <ShortProfile  :username="user.username" :pic="user.pic"/>
   </div>
   <div class="item-error" v-if="!filteredList().length">
      <h2>This profile has no followers.</h2>
   </div>
    </header>
</div>
 </template>

 <style>
 .page_b{
 background-color: #160F29;
   margin: -50px;
   display: flex;
   flex-direction: column;
   text-align: center;
 }
 .Bar_b {
     position: relative;
     margin-top:50px;
     max-width: 601px;
     margin-left: 120px;
 }
 .summary_page_b{
     position: relative;
     margin-top: 50px;
     height: 6000px;
     padding-left: 10px;
     padding-right: 16px;
     background-color:#246A73;
     border-radius: 20px;
 }
 .page_b h3 {
     position: relative;
     margin-top: 30px;
     font-size: 45px;
     text-align: center;
     color:beige;
     font-family: "Copperplate";
     text-transform: uppercase;
 }
 .item-error h2{
     position: relative;
     margin-top: 30px;
     font-size: 25px;
     text-align: center;
     color:beige;
     font-family: "Copperplate";
     text-transform: uppercase;
 }

 .item-user2{
    width: 500px;
    margin-left: 640px;
}

  </style>