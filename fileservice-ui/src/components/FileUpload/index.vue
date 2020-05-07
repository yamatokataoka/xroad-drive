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
    <file-list
      :files="uploadFiles"
      :listOpen="listOpen"
    ></file-list>
  </v-snackbar>
</template>

<script>
  import { mapState, mapActions } from "vuex";
  import FileListbar from '@/components/FileUpload/FileListbar';
  import FileList from '@/components/FileUpload/FileList';

  export default {
    name: 'FileUpload',
    components: {
      FileListbar,
      FileList
    },
    data() {
      return {
        timeout: 0,
        listOpen: true
      }
    },
    computed: {
      ...mapState('uploadFiles', ['uploading', 'uploadFiles'])
    },
    methods: {
      ...mapActions('uploadFiles', ['deleteUploadFiles'])
    },
    beforeDestroy() {
      this.deleteUploadFiles();
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