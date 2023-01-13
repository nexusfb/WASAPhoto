// Page of user profile with username, caption, profile picture, nmedia, nfollowers, nfollowing, user media
<script>
export default {
    data: function() {
        return {
            posts:[],

        }
    },
    methods: {


		GetMyStream() {
            this.loading = true;
            this.errormsg = null;
			this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            try {
                this.$axios.get("/users/:userid="+localStorage.getItem('Authorization')+"/stream/").then(response => (this.posts = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
		
	
		refresh() {
			this.GetMyStream()
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

		<div class="card" v-if="posts.length === 0">
			<div class="card-body">
				<p>Your stream is empty.</p>
			</div>
		</div>

		<div class="card" v-if="!loading" v-for="p in posts">
			<div class="card-header">
				{{ p.author }}
			</div>
			<div class="card-body">
				<p class="card-text">
                    <img :src= p.photo ><br />
					date: {{ p.date }}<br />
                    caption: {{ p.caption }}<br />
                    likes: {{ p.nlikes }}<br />
                    comments: {{ p.ncomments }}<br />
                    liked: {{ p.liked }}
				</p>
				
			</div>
		</div>

       

	</div>
</template>

<style scoped>
.card {
    margin-bottom: 20px;
}
</style>
