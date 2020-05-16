import axios from 'axios';

export default {
  namespaced: true,
  state() {
    return {
      fileList: []
    };
  },
  mutations: {
    updateFileList(state, fileList) {
      state.fileList = fileList;
    }
  },
  actions: {
    async fetchFileList({ commit }) {
      try {
        const response = await axios.get('/api/metadata');

        commit('updateFileList', response.data);
      } catch (error) {
        console.log('Failed to fetch file list');
      }
    }
  }
};
