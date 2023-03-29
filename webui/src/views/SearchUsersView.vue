// this is the search view where logged user can search for other users
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
        // get the list of users the logged user can search for
        async GetUserList() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.get("/search/:username=").then(response => (this.users = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

        // go to a specific profile
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

        // filter the list
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
        <h3>SEARCH USER</h3>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    <input type="text" v-model="input" placeholder="Search users..." />
   <div class="item-user2" v-for="user in filteredList()" :key="user">
    <ShortProfile  :username="user.username" :pic="user.pic"/>
   </div>
   <div class="item-error" v-if="!filteredList().length">
      <h2>No results found!</h2>
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
.summary_page_b input{
    width: 500px;
    border-radius: 35px;
    position: relative;
    margin-top: 30px;
	font-size: 20px;
    text-align: center;
    font-family: "Copperplate";
    text-transform: uppercase;
    background-color: #e7d7cb;
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