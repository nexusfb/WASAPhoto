<script>
import ShortProfile from "@/components/ShortProfile.vue"
export default {
    name: "CommentForm",
    props: {
        pp:String,
        commentid:  String,
        owner: String,
        timestamp: String,
        body: String,
        logged: String,
        authorid: String,
    },
    components: {
        ShortProfile
    },
    data: function () {
        return {
            header: localStorage.getItem('Authorization'),
            loading: false,
            errormsg: null,
            username: "",
            textComment: "",
            myProfilePic: "",
            myPP: "",
            ppUrl: "",
            imgUrl: "",
        }
    },
    methods: {
        async get_user_profile(name) {
            this.$router.push({ path: "/users/", query: { username: name } })
        },
        async Get_my_profile() {
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            if (this.header) {
                try {
                    let response = await this.$axios.get("/users/?username=" + this.owner)
                    this.myPP = response.data.profilepic
                } catch (e) {
                    this.errormsg = e.response.data.error.toString();
                }
            }
            this.loading = false;
        },
        async getImages() {
            if (this.pp) {
                let uri = await this.GetImage(this.pp)
                this.ppUrl = uri
            }
        },
        async GetImage(url) {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.get("/images/?image_name=" + url, { responseType: 'blob' })
                // Get the image data as a Blob object
                var imgBlob = response.data;
                // Create an object URL from the Blob object
                var uri = URL.createObjectURL(imgBlob);
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
            return uri
        },
        async deleteComment() {
            this.loading = true;
            this.errormsg = null;
            try {
                this.$axios.delete('/comments/' + this.commentid).then(() => (this.$emit('refresh-parent'))).catch(e => this.errormsg = e.response.data.error.toString())
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
        },
    },
    computed: {
        timeAgo() {
            var dateString = this.timestamp;
            var date = new Date(dateString);
            var year = date.getFullYear();
            var month = date.getMonth(); // getMonth() returns a number between 0 and 11
            var day = date.getDate();
            var hours = date.getHours();
            var minutes = date.getMinutes();
            var seconds = date.getSeconds();
            /* console.log("Timestamp Year: " + year + " Month: " + month + " Day: " + day);
            console.log("Timestamp Hours: " + hours + " Minutes: " + minutes + " Seconds: " + seconds); */
            var currentDate = new Date();
            var c_year = currentDate.getFullYear();// current year
            var c_month = currentDate.getMonth();
            var c_day = currentDate.getDate();
            var c_hours = currentDate.getHours();
            var c_minutes = currentDate.getMinutes();
            var c_seconds = currentDate.getSeconds();
            /* console.log("Current Year: " + c_year + " Month: " + c_month + " Day: " + c_day)
            console.log("Current Hours:" + c_hours + " Minutes: " + c_minutes + " Seconds: " + c_seconds); */
            var timeAgo = "";
            var diffYear = c_year - year;
            var diffMonth = c_month - month;
            var diffDay = c_day - day;
            var diffHour = c_hours - hours;
            var diffMinutes = c_minutes - minutes;
            var diffSeconds = c_seconds - seconds;
            if (diffYear !== 0) {
                timeAgo = diffYear + " years ago";
            } else if (diffMonth !== 0) {
                timeAgo = diffMonth + " months ago";
            } else if (diffDay !== 0) {
                timeAgo = diffDay + " days ago";
            } else if (diffHour !== 0) {
                timeAgo = diffHour + " hours ago";
            } else if (diffMinutes !== 0) {
                timeAgo = diffMinutes + " minutes ago";
            } else if (diffSeconds !== 0) {
                timeAgo = diffSeconds + " seconds ago";
            } else {
                timeAgo = "Just now";
            }
            /* console.log(timeAgo); */
            return timeAgo;
        },
        isMine() {
            return (this.authorid === this.logged)
        }
    },
    mounted() {
        this.Get_my_profile().then(() => this.getImages())
    }
}
</script>

<template>
    <div class="post">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <!-- header -->
        <header class="header section">
            <div class="header-author">
                <ShortProfile  :username="this.owner" :pic="this.pp"/>
            </div>
            <div class="header-more">
                <button v-if=isMine type="delete" @click="deleteComment">Delete Comment</button>
            </div>
        </header>
            <div class="caption">
                <span class="caption-span">{{ this.body }}</span>
                <span class="caption-span2">{{ timeAgo }}</span>
                
            </div>
        </div>
</template>


<style scoped>
.post {
    border-radius: 3px;
    border: 1px solid rgba(219, 219, 219, 1);
    max-width: 604px !important;
    margin-bottom: 30px;
    background-color: rgb(245, 239, 220);
    border-radius: 25px;
}
.post .section {
    padding-left: 16px;
    padding-right: 16px;
}
.post .header {
    display: flex;
    align-items: center;
    height: 55px;
    margin-top: 2px;
}
.post .header-author {
    display: flex;
    align-items: center;
}
.post .header-author-info {
    margin-left: 8px;
    font-size: 16px;
}
.post .header-more {
    margin-left: auto;
}
.post .header-more button[type="delete"] {
    color: white;
    padding: 6px 10px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    background-color: #911b1b;
}
.post .post-media {
    width: 600px;
    height: 400px;
}
.post .post-media .post-image {
    margin-top: 5px;
    width: 100%;
    height: 100%;
    object-fit: cover;
}
.post .two-col {
    display: flex;
    /* outline: solid; */
}
.post .action-buttons {
    margin-top: 0px;
    height: 60px;
}
.post .action-buttons button {
    padding-top: 6px;
    display: flex;
    align-items: center;
}
.post .action-buttons .num {
    display: inline;
    align-items: center;
    padding-left: 8px;
    font-size: 15px;
    font-weight: 600;
    font-family: Georgia, 'Times New Roman', Times, serif;
    color: #333;
}
.post .action-buttons .icon {
    height: 25px;
    width: 25px;
}
#like:hover {
    color: #555;
    /* background-color: rgb(232, 62, 79) */
}
#comment:hover {
    color: #555
}
.post .caption {
    flex-wrap: wrap;
    margin-top: 90px;
    padding-left: 25%;
    /* outline: solid; */
}
.post .caption li b:hover {
    text-decoration: underline;
    cursor: pointer;
}
.post .caption .caption-span {
    margin-left: -350px;
    overflow: auto;
    position: absolute;
    margin-top: -75px;
    color: #2b1e4f;
    font-family: "Copperplate";
    text-transform: uppercase;
    
}
.post .caption .caption-span2 {
    margin-left: 60px;
    overflow: auto;
    position: absolute;
    margin-top: -35px;
    color: #6b6972;
    font-family: "Copperplate";
    text-transform: uppercase;
    
}
.post .comments-list .time {
    margin-top: 8px;
}
.post .comments-list .time-ago {
    color: rgba(142, 142, 142, 1);
    text-transform: uppercase;
}
.post .comments-list .comment {
    max-width: inherit;
    border-top: 1px solid #efefef;
    margin-top: 30px;
    display: flex;
    height: 55px;
    align-items: center;
}
.post .comments-list .comment .text-body {
    margin-left: -10px;
    margin-top: -10px;
    max-width: 500px;
    background-color: #f4e8d7;
    font-size: 17px;
}
.post .comments-list .comment input {

    flex: 1;
}
.post .comments-list .comment input:focus {
    outline: none;
}
.post .comments-list .comment input::placeholder {
    background-color: -internal-light-dark(rgb(255, 255, 255), rgb(59, 59, 59));
}
.post .comments-list .comment button[type="submit"] {
    background-color: #fafafa;
    margin-left: 16px;
    font-size: 16px;
    color: rgb(23, 92, 122);
}
.post .comments-list .comment button[type="submit"]:hover {
    text-decoration: underline;
    cursor: pointer;
}
.like-content {
    display: inline-block;
    position: absolute;
    width: 90px;
    margin-top: -30px;
    margin-left: -10px;
    font-size: 18px;
    text-align: center;
}
.like-content .btn-secondary {
	  display: block;
	  margin: 40px auto 0px;
    text-align: center;
    background: #2b1e4f;
    border-radius: 35px;
    padding: 10px 17px;
    font-size: 15px;
    cursor: pointer;
    width: 90px;
    border: none;
    outline: none;
    color: #ffffff;
    text-decoration: none;
    color: beige;
    background-color: #2b1e4f;
    font-family: "Copperplate";
    text-transform: uppercase;
}
.like-content .btn-secondary2 {
    position: absolute;
    margin-left: 100px;
    margin-top: -35px;
	display: block;
    text-align: center;
    background: #4d4859;
    border-radius: 20px;
    font-size: 15px;
    cursor: pointer;
    width: 100px;
    height: 30px;
    border: none;
    outline: none;
    color: #ffffff;
    text-decoration: none;
    color: beige;
    background-color: #2b1e4f;
    font-family: "Copperplate";
    text-transform: uppercase;
}

.like-content .btn-secondary3 {
    position: absolute;
    margin-left: 210px;
    margin-top: -35px;
	display: block;
    text-align: center;
    background: #2b1e4f;
    border-radius: 20px;
    font-size: 15px;
    cursor: pointer;
    width: 140px;
    height: 30px;
    border: none;
    outline: none;
    color: #ffffff;
    text-decoration: none;
    color: beige;
    background-color: #2b1e4f;
    font-family: "Copperplate";
    text-transform: uppercase;
}
.caption .btn-secondary4 {
    position: absolute;
    margin-left: -150px;
    margin-top: 5px;
	display: block;
    text-align: center;
    background: #2b1e4f;
    border-radius: 20px;
    font-size: 15px;
    cursor: pointer;
    width: 140px;
    height: 20px;
    border: none;
    outline: none;
    color: #ffffff;
    text-decoration: none;
    color: beige;
    background-color: #2b1e4f;
    font-family: "Copperplate";
    text-transform: uppercase;
}
.btn-secondary5 {
    position: absolute;
    margin-left: 498px;
    margin-top: -10px;
	display: block;
    text-align: center;
    background: #2b1e4f;
    border-radius: 20px;
    font-size: 20px;
    cursor: pointer;
    width: 80px;
    height: 30px;
    border: none;
    outline: none;
    color: #ffffff;
    text-decoration: none;
    color: beige;
    background-color: #2b1e4f;
    font-family: "Copperplate";
    text-transform: uppercase;
}

</style>