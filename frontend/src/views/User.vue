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
      username: { 
        type: String, 
        default: "" 
      }
    },
    data: function(){
      return {};
    },
    mounted() {
      this.$store.dispatch("users/addUser", {username: this.username});  
    },  
    computed: {
        userPosts() {
            return this.$store.getters["posts/userPosts"](this.username);
        },
        user() {
            console.log(this.username)
            return this.$store.getters["users/getUser"](this.username)
        }
    }
};
</script>
