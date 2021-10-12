<template>
<div class="app-bar">
    <div class="left-links">
        <router-link class="app-bar-item" v-for="link in activeLinks" :key="link.name" :to="link.to">
            {{ link.name }}
        </router-link>
    </div>
    <div class="right-links ">
        <!-- <a class="app-bar-item" href="#" v-if="!loggedIn" @click.prevent="login">LOGIN</a>
         -->
        <router-link class="app-bar-item" href="#" v-if="!loggedIn" @click.prevent :to="{ name: 'Login' }" LOGIN>LOGIN</router-link>
        <a class="app-bar-item" href="#" v-if="loggedIn" @click.prevent="logoutButtonClicked">LOGOUT</a>
    </div>
</div>
</template>

<script>
import {
    mapGetters
} from "vuex";
import {
    mapActions
} from "vuex";

export default {
    data() {
        return {
            links: [{
                    visibleIfLoggedOut: true,
                    name: "Posts",
                    to: {
                        name: "Posts"
                    }
                },
                {
                    visibleIfLoggedOut: false,
                    name: "User",
                    to: {
                        name: "User",
                        params: {
                            username: this.$store.getters["auth/currentUser"].username
                        }
                    }
                }
            ]
        }
    },
    computed: {
        ...mapGetters({
            loggedIn: "auth/isLoggedIn",
        }),
        activeLinks() {
            return this.links.filter(
                link => link.visibleIfLoggedOut || this.loggedIn
            );
        }
    },
    methods: {
        ...mapActions({
            login: "auth/login",
            logout: "auth/logout"
        }),
        logoutButtonClicked() {
            this.logout().then(() => {
                this.$router.push({
                    name: "Login"
                });
            });
        },

    }
};
</script>

<style scoped>
.app-bar {
    height: 5vh;
    width: 100%;
    position: fixed;
    z-index: 1;
    top: 0;
    left: 0;
    background-color: rgb(29, 22, 13);
    overflow: hidden;
    display: block;
    justify-content: space-between;
}

.left-links {
    padding-left: 10%;
}

.left-links a {
    float: left;
}

a {
    float: left;
    display: block;
    color: #f2f2f2;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
    font-size: 17px;
}

.right-links {
    float: right;
}
</style>
