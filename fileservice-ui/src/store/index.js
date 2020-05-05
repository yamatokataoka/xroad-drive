import Vue from 'vue';
import Vuex from 'vuex';
import inputFile from './modules/inputFile.js';
import myFiles from './modules/myFiles.js';
import axios from 'axios';

axios.defaults.baseURL = 'https://8082-dot-11052537-dot-devshell.appspot.com';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    inputFile,
    myFiles
  }
});
