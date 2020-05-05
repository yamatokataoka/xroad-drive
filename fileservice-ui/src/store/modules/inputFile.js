import axios from 'axios';

export default {
  namespaced: true,
  state() {
    return {
      uploading: false,
      inputFiles: []
    };
  },
  mutations: {
    updateInputFiles(state, inputFiles) {
      state.inputFiles = [...state.inputFiles, ...inputFiles];
    },
    updateInputFileById(state, {
      inputFile,
      id = inputFile.id,
      file = inputFile.file,
      name = inputFile.name,
      progress = inputFile.progress,
      isDone = inputFile.isDone,
      indeterminate = inputFile.indeterminate
    }) {
      const index = state.inputFiles.findIndex(item => item.id === id);
      console.log('progress', progress)
      state.inputFiles.splice(index, 1, {
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
    deleteInputFiles(state) {
      state.inputFiles = [];
    }
  },
  actions: {
    updateInputFiles({ commit, state }, files) {
      console.log('called updateInputFiles');
      let inputFiles = [];
      const id = state.inputFiles.length === 0 ? 0 : (() => {
        const ids = state.inputFiles.map((element) => element.id);
        return Math.max(...ids) + 1;
      })(id);

      for (let i = 0; i < files.length; i++) {
        const file = files[i];
        const inputFile = {
          id: id + i,
          file: file,
          name: file.name,
          progress: 0,
          isDone: false,
          indeterminate: true
        };
        console.log('id', inputFile.id);
        inputFiles.push(inputFile);
      }
      commit('updateInputFiles', inputFiles);
    },
    async upload({ dispatch }, inputFile) {
      console.log('start to upload');
      const formData = new FormData();
      formData.append('file', inputFile.file);

      await axios({
        url: '/api/upload',
        method: 'POST',
        data: formData,
        headers: { 'Content-Type': 'multipart/form-data' },
        onUploadProgress: (progressEvent) => {
          const totalLength = progressEvent.lengthComputable ? progressEvent.total : progressEvent.target.getResponseHeader('content-length') || progressEvent.target.getResponseHeader('x-decompressed-content-length');
          if (totalLength !== null) {
            inputFile.indeterminate = false;
            let progress = Math.round( (progressEvent.loaded * 100) / totalLength );
            // To determine v-progress-linear's color correctly
            progress = progress === 100 ? 95 : progress;
            dispatch('updateProgressById', { id: inputFile.id, progress });
          }
        }
      });
    },
    updateProgressById({ state, commit }, { id, progress }) {
      const currentFile = state.inputFiles.find(item => item.id === id);
      commit('updateInputFileById', { inputFile: currentFile, progress });
    },
    updateIsDoneById({ state, commit }, { id, isDone }) {
      const currentFile = state.inputFiles.find(item => item.id === id);
      commit('updateInputFileById', { inputFile: currentFile, isDone });
    },
    updateIndeterminateById({ state, commit }, { id, indeterminate }) {
      const currentFile = state.inputFiles.find(item => item.id === id);
      commit('updateInputFileById', { inputFile: currentFile, indeterminate });
    },
    setUploading({ commit }, uploading) {
      commit('setUploading', uploading);
    },
    deleteInputFiles({ commit }) {
      commit('deleteInputFiles');
    },
  }
};
