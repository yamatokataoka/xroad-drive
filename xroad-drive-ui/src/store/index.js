import Vue from 'vue';
import Vuex from 'vuex';
import selectedFile from './modules/selectedFile';
import uploadFiles from './modules/uploadFiles';
import fileList from './modules/fileList';
import xroadMetadata from './modules/xroadMetadata';
import selectedXRoadMember from './modules/selectedXRoadMember';
import search from './modules/search';
import config from './modules/config';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    selectedFile,
    uploadFiles,
    fileList,
    xroadMetadata,
    selectedXRoadMember,
    search,
    config
  }
});
