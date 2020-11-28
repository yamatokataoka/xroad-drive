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
    async fetchFileList({ commit, rootState }, serviceId=null) {
      const xroadMemberId = rootState.config.xroadMemberId;
      let config = {};
      let path = '/api/metadata';

      if (serviceId) {
        const servicePath = serviceId.replace(/:/g, '/');
        if (!xroadMemberId) {
          console.log('Failed to set config for axios: xroadMemberId: ' + xroadMemberId);
          return;
        }

        const xroadMemberPath = xroadMemberId.replace(/:/g, '/');
        // TODO: URL join helper is needed.
        path = '/security-server/r1/' + servicePath + path;
        config.headers = { 'X-Road-Client': xroadMemberPath };
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
