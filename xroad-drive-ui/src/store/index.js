import Vue from 'vue';
import Vuex from 'vuex';
import selectedFile from './modules/selectedFile';
import uploadFiles from './modules/uploadFiles';
import fileList from './modules/fileList';
import xroadMetadata from './modules/xroadMetadata';
import selectedXRoadMember from './modules/selectedXRoadMember';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    selectedFile,
    uploadFiles,
    fileList,
    xroadMetadata,
    selectedXRoadMember
  }
});
