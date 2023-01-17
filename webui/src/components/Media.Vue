<script>
import ShortProfile from "@/components/ShortProfile.vue"
export default {
    name:Media;
    props: {
        media:"",

    },
    name: 'Media',
    components: {
        ShortProfile
    }
}
</script>


<template>
    <div class="post">
        <header class="header section">
            <div class="header-author">
                <ShortProfile  :username="media.author" :pic="this.media.authorpic"/>
            </div>
        </header>

        <div class="post-media">
            <img :src="media.photo" alt="" />
        </div>

        <div class="two-col section">
            <div class="action-buttons">
                <ul>
                    <li>
                        <button v-if="!loading" class="login-button" type="button">{{ media.nlikes }}</button>
                        
                    </li>
                    <li>
                        <button type="button">
                            <font-awesome-icon class="icon-comment" icon="fa-regular fa-comment" size="2x" />
                            <span class="num"> {{media.ncomments}} </span> <!-- {{ post.comments_count }} -->
                        </button>
                    </li>
                </ul>
            </div>

            <div class="caption">
                <li>
                    <span class="caption-span">{{media.caption}}</span> <!-- {{ post.caption }} -->
                </li>
            </div>
        </div>

        <div class="comments-list">
            <!-- datetime-->
            <div class="time section">
                <CustomText size="xxsmall" class="time-ago">10 MINUTES AGO</CustomText> <!-- {{ post.timestamp }} -->
            </div>

            <!-- comments form -->
            <div class="comment section">
                <input class="text-body" type="text" placeholder="Add a comment...">
                <a href="#" type="button">Post</a>
            </div>
        </div>
    </div>
</template>


<style scoped>
.post {
    border-radius: 3px;
    width: 700px;
    height: 600px;
    margin-bottom: 60px;
    background-color:rgb(237, 246, 249);
    border-radius: 20px;
}
.post .section {
    padding-left: 16px;
    padding-right: 16px;
}
.post .header {
    display: flex;
    align-items: center;
    height: 60px;
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
.post .post-media {
    width: 700px;
    height: 300px;
}
.post .two-col  {
    margin-top: 150px;
    align-items: center;
    background-color:rgb(233, 44, 27);
}
.post .action-buttons {
    margin-top: 0px;
    height: 60px;
    display: flex;
    align-items: center;
}
.post .action-buttons button {
    padding-top: 5px;
    display: flex;
    align-items: center;
   /*outline: solid;*/
}
.post .action-buttons .num {
    display: inline;
    align-items: center;
    padding-left: 8px;
    font-size: 15px;
    font-weight: 600;
    font-family:Georgia, 'Times New Roman', Times, serif;
    color: rgb(40,40,40);
}
.post .action-buttons .icon-like {
    height: 25px;
    width: 25px;
}
.post .action-buttons .icon-like:hover {
    color: #555;
    /* background-color: rgb(232, 62, 79) */
}
.post .action-buttons .icon-comment {
    height: 25px;
    width: 25px;
}
.post .action-buttons .icon-comment:hover {
    color: #555
}
.post .caption{
    display: flex;
    flex-wrap: wrap;
    align-items: flex-start;
   /*outline: solid;*/
    margin-bottom: 0%;
    margin-left: 10%;
}
.post .caption li b:hover {
    cursor: pointer;
}
.post .caption li .caption-span {
    margin-left: 4px;
    overflow: auto;
}
.post .comments-list .time {
    margin-top: 8px;
}
.post .comments-list .time-ago {
    color: rgba(142, 142, 142, 1);
}
.post .comments-list .comment {
    max-width: inherit;
    border-top: 1px solid #efefef;
    margin-top: 4px;
    display: flex;
    height: 55px;
    align-items: center;
}
.post .comments-list .comment .text-body {
    margin-left: 16px;
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
.post .comments-list .comment a {
    margin-left: 16px;
    font-size: 16px;
    color: rgba(0, 160, 230, 1);
}
.post .comments-list .comment a:hover {
    text-decoration: underline;
    cursor: pointer;
}
.login-button{
    border-radius: 20px;
    margin-top: 80px;
    width: 70%;
    height: 40px;
    border: 1px solid rgb(255, 255, 255);
    background: #f4ba00;
    font-size: 14px;
    font-family: "Rubik", sans-serif;
    font-weight: 400;
    color: white;
    letter-spacing: 4px;
    text-transform: uppercase;
    text-decoration: none;
    border-radius: 25px;
}
</style>