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
      const xroadMemberId = process.env.VUE_APP_XROAD_MEMBER_ID;
      let config = {};
      let path = '/api/metadata';

      if (serviceId) {
        const servicePath = serviceId.replace(/:/g, '/');
        if (!xroadMemberId) {
          console.log('Failed to set config for axios: xroadMemberId: ' + xroadMemberId);
          return;
        }
        // TODO: URL join helper is needed.
        path = '/security-server/r1/' + servicePath + path;
        config.headers = { 'X-Road-Client': xroadMemberId};
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