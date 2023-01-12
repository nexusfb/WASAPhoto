// Page of user profile with username, caption, profile picture, nmedia, nfollowers, nfollowing, user media
<script>
export default {
    data: function() {
        return {
            mediaid: this.$route.params.mediaid,
            comments: "",
            comment: "",

        }
    },
    methods: {
        commentMedia() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.post("/media/:mediaid="+ this.mediaid+"/comments/", {
					content: this.comment,});
				this.refresh();
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
			this.refresh();
        },

		GetMediaComments() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.get("/media/:mediaid="+this.mediaid+"/comments/").then(response => (this.comments = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
		
	
		refresh() {
			this.GetMediaComments()
		},
		
        deleteComment(id) {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.delete("/comments/:commentid="+id);
				this.refresh();
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
			this.refresh();
        },
    },
    mounted() {
		this.refresh();
    },
}
</script>
<template>
	<div>
		

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="card" v-if="comments.length === 0">
			<div class="card-body">
				<p>No comments for this media.</p>


			</div>
		</div>

		<div class="card" v-if="!loading" v-for="c in comments">
			<div class="card-header">
				Comment
			</div>
			<div class="card-body">
				<p class="card-text">
					author: {{ c.author }}<br />
                    date: {{ c.date }}<br />
                    content: {{ c.content }}
				</p>
				<a href="javascript:" class="btn btn-danger" @click="deleteComment(c.commentid)">Delete</a>
			</div>
		</div>

        <div class="mb-3">
        		<label for="description" class="form-label">Comment</label>
            	<input type="text" class="form-control" id="comment" v-model="comment" placeholder="comment here">
                <a href="javascript:" class="btn btn-primary" @click="this.commentMedia()">Create a new comment</a>
        </div>

	</div>
</template>

<style scoped>
.card {
    margin-bottom: 20px;
}
</style>
