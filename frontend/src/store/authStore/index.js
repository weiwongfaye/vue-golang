export default {
    namespaced: true,
    state: {
      user: {
        username: "",
        //here there will be the logic for auth and so on...
        loggedIn: false,
        token: ""

      }
    },
    mutations: {
    //   LOGIN(state) {
    //     state.user.loggedIn = true;
    //   },
    //   LOGOUT(state) {
    //     state.user.loggedIn = false;
    //   }
    LOGIN(state, { username, token }) {
        state.user.loggedIn = true;
        state.user.username = username;
        state.user.token = token;
    },
    LOGOUT(state) {
        state.user.loggedIn = false;
        state.user.username = "";
        state.user.token = "";
    }
    },
    actions: {
    //   async login(context) {
    //     console.log("login in...");
    //     context.commit("LOGIN");
    //   },
    //   async logout(context) {
    //     context.commit("LOGOUT");
    //   }
    async login(context, { username, password }) {
        return fetch("http://localhost:3000/api/auth/login", {
          method: "POST",
          body: JSON.stringify({
                                 username: username, 
                                 password: password })
          })
          .then(response => {
            if (!response.ok) {
              throw new Error("Cannot login!");
            }
            return response.json();
          }).then(data => {
             context.commit("LOGIN",
                  { username: username, token: data.token });
          }).catch(error => {
              context.commit("LOGOUT");
              throw error;
            });
        },
    async logout(context) {
        context.commit("LOGOUT");
    },
    async signup(context, { username, password }) {
        return fetch("http://localhost:3000/api/auth/create-user", {
          method: "POST",
          body: JSON.stringify(
              { username: username, password: password })
          }).then(response => {
            if (!response.ok) {
              throw new Error("Cannot signup!");
            }
            return response.json();
          }).then(data => {
            context.commit("LOGIN", 
                  { username: username, token: data.token });
          }).catch(error => {
            context.commit("LOGOUT");
            error.read().then((data, done) => {
                console.log(done)
                throw Error(data);
            });
          });
        }
    },
    getters: {
      currentUser(state) {
        return state.user;
      },
      isLoggedIn(state) {
        if (!state.user) return false;
        return state.user.loggedIn;
      },
      getTokenHeader(state) {
        return "Bearer " + state.user.token;
      }
    }
};