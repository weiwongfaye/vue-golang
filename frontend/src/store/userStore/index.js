export default {
    namespaced: true,
    state: {
      loadedUsers: [
        {
          id: 1,
          username: "idalmasso",
          description: "Here is the description",
          //here there will be the logic for auth and so on...
          loggedIn: false
        }
      ]
    },
    mutations: {
      ADD_USER(state, user) {
        if (state.loadedUsers.some(u => u.username == user.username)) {
          state.loadedUsers.splice(
            state.loadedUsers.indexOf(u => u.username == user.username),
            1
          );
        }
        state.loadedUsers.push(user);
      }
    },
    actions: {},
    getters: {
      getUser: state => username => {
        if (state.loadedUsers.some(user => user.username == username)) {
          return state.loadedUsers.find(user => user.username == username);
        } else {
          //Here I'll have to request from the server!!
          console.log("Here I'll have to request from the server!!")
          return {};
        }
      }
    }
  };