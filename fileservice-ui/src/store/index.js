import Vue from 'vue';
import Vuex from 'vuex';
import myFiles from './modules/myFiles.js';
import axios from 'axios';

axios.defaults.baseURL = '/api';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    myFiles
  }
});
