<script>
import { ref } from 'vue';
import ShortProfile from "@/components/ShortProfile.vue"
import NavBar from "@/components/NewHomeBar.vue";
export default {
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
        async GetUserList() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.get("/users/:userid="+localStorage.getItem('Authorization')+"/banned/").then(response => (this.users = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

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
    <div class="Home">
		<NavBar :profilo="this.$route.params.username"/>
		
	</div>
    <h3>BANNED</h3>
   <div class="item user" v-for="user in filteredList()" :key="user">
    <ShortProfile  :username="user.username" :pic="user.pic"/>
    
   </div>
   <div class="item error" v-if="input&&!filteredList().length">
      <p>No results found!</p>
   </div>
 </template>
 <style>
 @import url("https://fonts.googleapis.com/css2?family=Montserrat&display=swap");
 
 * {
   padding: 0;
   margin: 0;
   box-sizing: border-box;
   font-family: "Montserrat", sans-serif;
 }
 
 body {
   padding: 20px;
   min-height: 100vh;
   background-color: rgb(234, 242, 255);
 }
 
 input {
   display: block;
   width: 350px;
   margin: 20px auto;
   padding: 10px 45px;
   background-size: 15px 15px;
   font-size: 16px;
   border: none;
   border-radius: 5px;
   box-shadow: rgba(50, 50, 93, 0.25) 0px 2px 5px -1px,
     rgba(0, 0, 0, 0.3) 0px 1px 3px -1px;
 }
 .titolo{
    text-align: center;
    background-color: red;
 }
 .item {
   width: 350px;
   margin: 0 auto 10px auto;
   padding: 10px 20px;
   color: white;
   border-radius: 5px;
   box-shadow: rgba(0, 0, 0, 0.1) 0px 1px 3px 0px,
     rgba(0, 0, 0, 0.06) 0px 1px 2px 0px;
 }
 
 .user {
   cursor: pointer;
 }
 
 .error {
   background-color: tomato;
 }
 </style>