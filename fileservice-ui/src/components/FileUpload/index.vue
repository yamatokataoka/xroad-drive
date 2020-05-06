<template>
  <v-snackbar
    :timeout="timeout"
    right
    v-if="uploadFiles.length"
    :value="uploading"
  >
    <file-listbar
      @clickChevron="listOpen = !listOpen"
      :listOpen="listOpen"
    ></file-listbar>
    <upload-file-list
      :files="uploadFiles"
      :listOpen="listOpen"
    ></upload-file-list>
  </v-snackbar>
</template>

<script>
  import { mapState } from "vuex";
  import FileListbar from '@/components/FileUpload/FileListbar';
  import UploadFileList from '@/components/FileUpload/UploadFileList';

  export default {
    name: 'FileUpload',
    components: {
      FileListbar,
      UploadFileList
    },
    data() {
      return {
        timeout: 0,
        listOpen: true
      }
    },
    computed: {
      ...mapState({
        uploadFiles: state => state.myFiles.uploadFiles.uploadFiles,
        uploading: state => state.myFiles.uploadFiles.uploading
      })
    }
  };
</script>

<style scoped>
  >>>.v-snack__content {
    padding: 0;
    display: block;
  }
  >>>.v-list-group__header {
    display: none;
  }
</style>