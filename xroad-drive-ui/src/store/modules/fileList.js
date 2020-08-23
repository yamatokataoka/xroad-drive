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
    async fetchFileList({ commit }, serviceId=null) {
      let config = {};
      let path = '/api/metadata';

      if (serviceId) {
        const servicePath = serviceId.replace(/:/g, '/');
        if (!process.env.VUE_APP_XROAD_MEMBER_ID) {
          console.log('The process.env.VUE_APP_XROAD_MEMBER_ID is ' + process.env.VUE_APP_XROAD_MEMBER_ID);
          return;
        }
        // TODO: URL join helper is needed.
        path = '/security-server/r1/' + servicePath + path;
        config.headers = { 'X-Road-Client': process.env.VUE_APP_XROAD_MEMBER_ID};
      }

      try {
        const response = await axios.get(path, config);

        commit('updateFileList', response.data);
      } catch (error) {
        console.log('Failed to fetch file list');
      }
    }
  }
};