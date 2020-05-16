import Vue from 'vue';
import Vuex from 'vuex';
import myFiles from './modules/myFiles';
import uploadFiles from './modules/uploadFiles';
import fileList from './modules/fileList';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    myFiles,
    uploadFiles,
    fileList
  }
});
