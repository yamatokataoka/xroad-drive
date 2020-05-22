import Vue from 'vue';
import Vuex from 'vuex';
import selectedFile from './modules/selectedFile';
import uploadFiles from './modules/uploadFiles';
import fileList from './modules/fileList';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    selectedFile,
    uploadFiles,
    fileList
  }
});
