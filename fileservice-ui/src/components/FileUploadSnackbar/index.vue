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
    <upload-item-list
      :files="uploadFiles"
      :listOpen="listOpen"
    ></upload-item-list>
  </v-snackbar>
</template>

<script>
  import { mapState } from "vuex";
  import UploadToolbar from '@/components/FileUploadSnackbar/UploadToolbar';
  import UploadItemList from '@/components/FileUploadSnackbar/UploadItemList';

  export default {
    name: 'FileUploadSnackbar',
    components: {
      UploadToolbar,
      UploadItemList
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