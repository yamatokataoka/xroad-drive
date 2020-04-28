export default {
  namespaced: true,

  state: {
    selectedFile: null
  },
  mutations: {
    setSelectedFile(state, payload) {
      state.selectedFile = payload;
    }
  },
  actions: {
  }
};