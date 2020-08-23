export default {
  namespaced: true,
  state() {
    return {
      selectedXRoadMember: null
    };
  },
  mutations: {
    updateSelectedXRoadMember(state, selectedXRoadMember) {
      state.selectedXRoadMember = selectedXRoadMember;
    }
  },
  actions: {
    updateSelectedXRoadMember({ commit }, selectedXRoadMember) {
      commit('updateSelectedXRoadMember', selectedXRoadMember);
    }
  }
};
