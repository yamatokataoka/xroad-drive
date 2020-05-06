<template>
  <v-snackbar
    :timeout="timeout"
    right
    v-if="inputFiles.length"
    :value="uploading"
  >
    <upload-toolbar
      @clickChevron="listOpen = !listOpen"
      :listOpen="listOpen"
    ></upload-toolbar>
    <upload-item-list
      :files="inputFiles"
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
        inputFiles: state => state.myFiles.inputFiles.inputFiles,
        uploading: state => state.myFiles.inputFiles.uploading
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