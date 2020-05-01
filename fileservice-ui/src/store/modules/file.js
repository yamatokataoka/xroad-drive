export default {
  namespaced: true,

  state: {
    selectedFile: null
  },
  mutations: {
    updateSelectedFile(state, payload) {
      state.selectedFile = payload;
    }
  },
  actions: {
  }
};