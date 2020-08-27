<template>
  <v-dialog
    persistent
    no-click-animation
    hide-overlay
    max-width="324"
    :value="uploading"
    :retain-focus="focus"
  >
    <file-listbar
      @clickChevron="listOpen = !listOpen"
      :listOpen="listOpen"
    ></file-listbar>
    <file-list
      :files="files"
      :listOpen="listOpen"
    ></file-list>
  </v-dialog>
</template>

<script>
  import { mapState, mapActions } from "vuex";
  import FileListbar from '@/components/ui/FileUploadDialog/FileListbar';
  import FileList from '@/components/ui/FileUploadDialog/FileList';
  import Vue from 'vue';

  export default {
    name: 'FileUploadDialog',
    components: {
      FileListbar,
      FileList
    },
    data() {
      return {
        timeout: 0,
        listOpen: true,
        focus: false,
        files: []
      }
    },
    watch: {
      // Reflect changes of uploadFiles into child props
      uploadFiles: {
        handler() {
          Vue.set(this, 'files', this.uploadFiles)
        }, deep: true
      }
    },
    computed: {
      ...mapState('uploadFiles', ['uploading', 'uploadFiles'])
    },
    methods: {
      ...mapActions('uploadFiles', ['deleteUploadFiles', 'updateUploading'])
    },
    beforeDestroy() {
      this.updateUploading(false);
      this.deleteUploadFiles();
    }
  };
</script>

<style scoped>
  >>>.v-dialog {
    position: absolute;
    bottom: 0;
    right: 0;
  }
  >>>.v-list-group__header {
    display: none;
  }
</style>