import axios from 'axios';

export default {
  state() {
    return {
      uploading: false,
      uploadFiles: []
    };
  },
  mutations: {
    updateUploadFiles(state, uploadFiles) {
      state.uploadFiles = [...state.uploadFiles, ...uploadFiles];
    },
    updateUploadFileById(state, {
      uploadFile,
      id = uploadFile.id,
      file = uploadFile.file,
      name = uploadFile.name,
      progress = uploadFile.progress,
      isDone = uploadFile.isDone,
      indeterminate = uploadFile.indeterminate
    }) {
      const index = state.uploadFiles.findIndex(item => item.id === id);
      console.log('progress', progress)
      state.uploadFiles.splice(index, 1, {
        id,
        file,
        name,
        progress,
        isDone,
        indeterminate
      });
    },
    setUploading(state, uploading) {
      state.uploading = uploading;
    },
    deleteUploadFiles(state) {
      state.uploadFiles = [];
    }
  },
  actions: {
    updateUploadFiles({ commit, state }, files) {
      console.log('called updateUploadFiles');
      let uploadFiles = [];
      const id = state.uploadFiles.length === 0 ? 0 : (() => {
        const ids = state.uploadFiles.map((element) => element.id);
        return Math.max(...ids) + 1;
      })(id);

      for (let i = 0; i < files.length; i++) {
        const file = files[i];
        const uploadFile = {
          id: id + i,
          file: file,
          name: file.name,
          progress: 0,
          isDone: false,
          indeterminate: true
        };
        console.log('id', uploadFile.id);
        uploadFiles.push(uploadFile);
      }
      commit('updateUploadFiles', uploadFiles);
    },
    async upload({ dispatch }, uploadFile) {
      console.log('start to upload');
      const formData = new FormData();
      formData.append('file', uploadFile.file);

      await axios({
        url: '/api/upload',
        method: 'POST',
        data: formData,
        headers: { 'Content-Type': 'multipart/form-data' },
        onUploadProgress: (progressEvent) => {
          const totalLength = progressEvent.lengthComputable ? progressEvent.total : progressEvent.target.getResponseHeader('content-length') || progressEvent.target.getResponseHeader('x-decompressed-content-length');
          if (totalLength !== null) {
            uploadFile.indeterminate = false;
            let progress = Math.round( (progressEvent.loaded * 100) / totalLength );
            // To determine v-progress-linear's color correctly
            progress = progress === 100 ? 95 : progress;
            dispatch('updateProgressById', { id: uploadFile.id, progress });
          }
        }
      });
    },
    updateProgressById({ state, commit }, { id, progress }) {
      const currentFile = state.uploadFiles.find(item => item.id === id);
      commit('updateUploadFileById', { uploadFile: currentFile, progress });
    },
    updateIsDoneById({ state, commit }, { id, isDone }) {
      const currentFile = state.uploadFiles.find(item => item.id === id);
      commit('updateUploadFileById', { uploadFile: currentFile, isDone });
    },
    updateIndeterminateById({ state, commit }, { id, indeterminate }) {
      const currentFile = state.uploadFiles.find(item => item.id === id);
      commit('updateUploadFileById', { uploadFile: currentFile, indeterminate });
    },
    setUploading({ commit }, uploading) {
      commit('setUploading', uploading);
    },
    deleteUploadFiles({ commit }) {
      commit('deleteUploadFiles');
    },
  }
};
