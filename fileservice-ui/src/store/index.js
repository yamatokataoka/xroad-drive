import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';
import myFiles from './modules/myFiles';
import uploadFiles from './modules/uploadFiles';

axios.defaults.baseURL = 'https://8082-dot-11052537-dot-devshell.appspot.com';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    myFiles,
    uploadFiles
  }
});
