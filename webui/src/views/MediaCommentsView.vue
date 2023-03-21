<script>
import { ref } from 'vue';
import CommentForm from "@/components/CommentForm.vue"
import NavBar from "@/components/NewHomeBar.vue";
export default {
    props:['mediaid'],
    components: {
        CommentForm,
        NavBar
    },
    data: function() {
        return {
            loading : false,
            errmsg : null,
            input: ref(""),
            name:"",
			comments:[],
            logged: localStorage.getItem('Authorization'),
        }
    },
    methods: {
        async GetCommentsList() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.get("/media/:mediaid="+this.$route.params.mediaid+"/comments/").then(response => (this.comments = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

        async GetUsername() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                let response = await this.$axios.get("/id")
				this.name = response.data;
            } catch (e) {
                this.errormsg = e.
                toString();
				
            }
            this.loading = false;
        },
        async refresh() {
			await this.GetCommentsList()
		},

        
    },

	mounted() {
		this.GetCommentsList().then(this.GetUsername());
        this.refresh();
	}
}
</script>
<template>
   <div class="page_b">
    
	<div class="Bar_b">
		<NavBar :profilo="this.name"/>
	</div>
    <header class="summary_page_b">
        <h3>COMMENTS</h3>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="comment_page1">
            <CommentForm	v-on:refresh-parent="refresh"  v-for="comment in comments" :key="comment.commentid" :commentid="comment.commentid" :pp="comment.authorpic" :owner="comment.author" :timestamp="comment.date" :body="comment.content" :logged="this.logged" :authorid="comment.authorid"/>
        </div>
    <div class="item-error" v-if="!comments.length">
      <h2>No comments yet!</h2>
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

 .page_b h3 {
     position: relative;
     margin-top: 30px;
     font-size: 45px;
     text-align: center;
     color:beige;
     font-family: "Copperplate";
     text-transform: uppercase;
 }
 .comment_page1 {
    margin-left: 550px;
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