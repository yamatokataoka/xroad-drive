<template>
  <v-snackbar
    :timeout="timeout"
    right
    v-if="inputFiles.length"
    :value="uploading"
  >
    <upload-tool-bar
      @clickChevronUp="listOpen = false"
      @clickChevronDown="listOpen = true"
      :listOpen="listOpen"
    ></upload-tool-bar>
    <upload-item-list
      :files="inputFiles"
      :listOpen="listOpen"
    ></upload-item-list>
  </v-snackbar>
</template>

<script>
  import { mapState } from "vuex";
  import UploadToolBar from '@/components/FileUploadSnackBar/UploadToolBar';
  import UploadItemList from '@/components/FileUploadSnackBar/UploadItemList';

  export default {
    name: 'FileUploadSnackBar',
    components: {
      UploadToolBar,
      UploadItemList
    },
    data() {
      return {
        timeout: 0,
        listOpen: true
      }
    },
    computed: {
      ...mapState('inputFile', ['inputFiles', 'uploading'])
    },
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