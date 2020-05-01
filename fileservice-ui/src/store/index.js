import Vue from 'vue';
import Vuex from 'vuex';
import file from './modules/file.js';
import axios from 'axios';

axios.defaults.baseURL = '/api';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    file
  }
});
