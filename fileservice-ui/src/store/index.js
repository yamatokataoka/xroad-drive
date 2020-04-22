import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    selectedFile: null
  },
  mutations: {
    setSelectedFile(state, payload) {
      state.selectedFile = payload;
    }
  },
  actions: {
  },
  modules: {
  }
})
