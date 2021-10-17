import { createRouter, createWebHistory } from 'vue-router'
import store from "@/store/index.js";
import Posts from "../views/Posts.vue";
import Login from "../views/Login.vue"
import User from "../views/User.vue"
// import NotFoundComponent from "../components/404.vue"


const routes = [

  {
    path: '/',
    name: 'Posts',
    component: Posts
  },
  // {
  //   path: '/user',
  //   name: 'User',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/User.vue'),
  //   props: true
  // },
  {
    path: "/users/:username",
    name: "Users",
    component: User,
    props: true,
    beforeEnter: (to, from, next) => {
      if (!store.getters["auth/isLoggedIn"]) {
        next({ name: "Login" });
      } else {
        next();
      }
    }
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
    //This is not needed right by now, because the store is  refreshed on page refresh... Will be needed!
    beforeEnter: (to, from, next) => {
      if (store.getters["auth/isLoggedIn"]) {
        next({ name: "Posts" });
      } else {
        next();
      }
    }
  },

  // {
  //   path: '/:pathMatch(.*)*',
  //   name: '404',
  //   component: NotFoundComponent
  // }
]


const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
