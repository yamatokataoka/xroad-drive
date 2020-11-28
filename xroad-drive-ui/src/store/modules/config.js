import axios from 'axios';

export default {
  namespaced: true,
  state() {
    return {
      xroadMemberId: '',
      commonServiceCode: ''
    };
  },
  mutations: {
    updateXRoadMemberId(state, xroadMemberId) {
      state.xroadMemberId = xroadMemberId;
    },
    updateCommonServiceCode(state, commonServiceCode) {
      state.commonServiceCode = commonServiceCode;
    }
  },
  actions: {
    async fetchConfig({ commit }) {
      try {
        const response = await axios.get('/api/config');
        await commit('updateXRoadMemberId', response.data.memberId);
        await commit('updateCommonServiceCode', response.data.commonServiceCode);
        console.log('Update is finished');
      } catch (error) {
        console.log('Failed to fetch X-Road config: ' + error);
      }
    }
  }
};
