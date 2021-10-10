import { createStore } from 'vuex'
import postnameNotMatter from "./postStore/index.js";
import usernameNotMatter from "./userStore/index.js";
import authModule from "./authStore/index.js"

export default createStore({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    posts: postnameNotMatter, users: usernameNotMatter, auth: authModule
  }
})
