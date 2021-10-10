<template>
<div class="container">
    <h1>{{ user.username }}</h1>
    <p>{{ user.description }}</p>

    <post-list :posts="userPosts" title="Last Posts"></post-list>
</div>
</template>

<script>
import PostList from "../components/PostList.vue";

export default {
    components: {
        PostList
    },
    props: {
      userid: { 
        type: String, 
        default: "" 
      }
    },
    data: function(){
      return {};
    },
    computed: {
        userPosts() {
            return this.$store.getters["posts/userPosts"](this.user.username);
        },
        user() {
            return this.$store.getters["users/getUser"](this.userid)
        }
    }
};
</script>
