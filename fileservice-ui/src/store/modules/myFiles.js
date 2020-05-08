export default {
  namespaced: true,
  state() {
    return {
      selectedFile: null
    };
  },
  mutations: {
    updateSelectedFile(state, selectedFile) {
      state.selectedFile = selectedFile;
    }
  },
  actions: {
    updateSelectedFile({ commit }, selectedFile) {
      commit('updateSelectedFile', selectedFile);
    },
    deleteSelectedFile({ commit }) {
      commit('updateSelectedFile', null);
    }
  }
};
