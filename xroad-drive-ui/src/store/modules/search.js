export default {
  namespaced: true,
  state() {
    return {
      search: ""
    };
  },
  mutations: {
    updateSearch(state, search) {
      state.search = search;
    }
  },
  actions: {
    updateSearch({ commit }, search) {
      commit('updateSearch', search);
    }
  }
};
