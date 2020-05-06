<template>
  <v-snackbar
    :timeout="timeout"
    right
    v-if="uploadFiles.length"
    :value="uploading"
  >
    <upload-toolbar
      @clickChevron="listOpen = !listOpen"
      :listOpen="listOpen"
    ></upload-toolbar>
    <upload-file-list
      :files="uploadFiles"
      :listOpen="listOpen"
    ></upload-file-list>
  </v-snackbar>
</template>

<script>
  import { mapState } from "vuex";
  import UploadToolbar from '@/components/FileUploadSnackbar/UploadToolbar';
  import UploadFileList from '@/components/FileUploadSnackbar/UploadFileList';

  export default {
    name: 'FileUploadSnackbar',
    components: {
      UploadToolbar,
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