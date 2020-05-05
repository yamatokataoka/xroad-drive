import axios from 'axios';

export default {
  namespaced: true,
  state: {
    inputFiles: []
  },
  mutations: {
    updateInputFiles(state, inputFiles) {
      state.inputFiles = [...state.inputFiles, ...inputFiles];
    }
  },
  actions: {
    updateInputFiles({ commit, state }, files) {
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
        inputFiles.push(inputFile);
      }
      commit('updateInputFiles', inputFiles);
    }
  }
};