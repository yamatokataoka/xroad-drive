import axios from 'axios';

export default {
  namespaced: true,
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
    updateUploadFile(state, {
      uploadFile,
      id = uploadFile.id,
      file = uploadFile.file,
      name = uploadFile.name,
      progress = uploadFile.progress,
      isDone = uploadFile.isDone,
      indeterminate = uploadFile.indeterminate
    }) {
      const index = state.uploadFiles.findIndex(item => item.id === id);
      state.uploadFiles.splice(index, 1, {
        id,
        file,
        name,
        progress,
        isDone,
        indeterminate
      });
    },
    updateUploading(state, uploading) {
      state.uploading = uploading;
    },
    deleteUploadFiles(state) {
      state.uploadFiles = [];
    }
  },
  actions: {
    updateUploadFiles({ commit, getters }, files) {
      let uploadFiles = [];
      const id = getters.nextId;

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
        uploadFiles.push(uploadFile);
      }
      commit('updateUploadFiles', uploadFiles);
    },
    async upload({ dispatch }, uploadFile) {
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
            dispatch('updateIndeterminateById', { id: uploadFile.id, indeterminate: false });
            let progress = Math.round( (progressEvent.loaded * 100) / totalLength );
            // To determine v-progress-linear's color correctly
            progress = progress === 100 ? 90 : progress;
            dispatch('updateProgressById', { id: uploadFile.id, progress });
          }
        }
      });
    },
    updateProgressById({ commit, getters }, { id, progress }) {
      const uploadFile = getters.getFileById(id);
      commit('updateUploadFile', {
        uploadFile,
        progress
      });
    },
    updateIsDoneById({ commit, getters }, { id, isDone }) {
      const uploadFile = getters.getFileById(id);
      commit('updateUploadFile', {
        uploadFile,
        isDone
      });
    },
    updateIndeterminateById({ commit, getters }, { id, indeterminate }) {
      const uploadFile = getters.getFileById(id);
      commit('updateUploadFile', {
        uploadFile,
        indeterminate
      });
    },
    updateUploading({ commit }, uploading) {
      commit('updateUploading', uploading);
    },
    deleteUploadFiles({ commit }) {
      commit('deleteUploadFiles');
    },
  },
  getters: {
    nextId(state) {
      const { uploadFiles } = state;

      if (uploadFiles.length === 0) {
        return 0;
      } else {
        const ids = uploadFiles.map((element) => element.id);
        return Math.max(...ids) + 1;
      }
    },
    getFileById: (state) => (id) => {
      return state.uploadFiles.find(uploadFile => uploadFile.id === id);
    }
  }
};
