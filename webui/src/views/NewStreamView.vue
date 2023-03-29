// this is the stream view where the logged user can see the media of its followings
<script>
import FinalMedia from '@/components/FinalMedia.vue'
import NavBar from '@/components/NewHomeBar.vue'
export default {
	name: 'stream',
	components: {
		FinalMedia,
		NavBar
	},
	data: function () {
		return {
			errormsg: null,
			loading: false,
			media:[],
		}
	},

	methods: {
        // get logged user stream
		async refresh() {
			this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.get("/users/:userid="+localStorage.getItem('Authorization')+"/stream/").then(response => (this.media = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>



<template>
	<div class="stream_page">
	 <div class="Bar_b">
		 <NavBar :profilo="this.$route.params.username"/>
	 </div>
	 <header class="summary_page_b">
		 <h3>MY STREAM</h3>
		 <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		 <div class="stream_timeline">
		 <FinalMedia	v-on:refresh-parent="refresh" v-for="post in this.media" :key="post.id" :pp="post.authorpic" :photoId="post.id"
				:owner="post.author" :image="post.photo"
				:timestamp="post.date" :caption="post.caption" :likesCount="post.nlikes"
				:commentsCount="post.ncomments" :liked="post.liked" :logged="this.logged" :authorid="post.authorid"/>
	 	 </div>
		  <div class="item-error" v-if="!this.media.length">
      		<h2>Your stream is empty.</h2>
   </div>
	 </header>
 </div>
  </template>
 
  <style>
 .stream_page{
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
	 height: 10000px;
	 padding-left: 10px;
	 padding-right: 16px;
	 background-color:#246A73;
	 border-radius: 20px;
 }
 .stream_page h3 {
	 position: relative;
	 margin-top: 30px;
	 font-size: 45px;
	 text-align: center;
	 color:beige;
	 font-family: "Copperplate";
	 text-transform: uppercase;
 }
 .stream_timeline {
	margin-left: 580px;
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