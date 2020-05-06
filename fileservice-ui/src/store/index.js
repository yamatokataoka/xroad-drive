import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

axios.defaults.baseURL = 'https://8082-dot-11052537-dot-devshell.appspot.com';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    // register files module dynamically on MyFile view
  }
});
