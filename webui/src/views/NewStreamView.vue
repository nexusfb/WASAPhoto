// View of logged user stream
// This is the default view

<script>
// import post & nav bar
import Media from '@/components/Media.vue'
import HomeBar from '@/components/NewHomeBar.vue'
export default {
	name: 'stream',
	components: {
		Media,
		HomeBar
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
	<div class="stream">
        <div class="timeline">
		    <Media v-for="m in this.media" :media="m"/>
        </div>
       
		<div class="sidebar">
			<HomeBar />
		</div>
	</div>
</template>

<style>
.stream {
	margin-left: auto;
	margin-right: auto;
	padding-bottom: 10vh;
    background-color: #02587b;
}
.timeline {
    align-items: center;
    max-width: 700px;
	margin-left: 400px;
	margin-right: auto;
	padding-bottom: 10vh;
    background-color: #6bd5ff;
}
@media (--t) {
	.stream {
		max-width: none;
		display: grid;
		grid-template-columns: 1fr 295px;
		grid-gap: 28px;
	}
}
.sidebar {
	display: none;
}
@media (--t) {
	.sidebar {
		display: block;
		margin-top: 16px;
	}
	.sidebar p {
		position: sticky;
		top: calc(53px + 30px + 18px);
	}
}
</style>