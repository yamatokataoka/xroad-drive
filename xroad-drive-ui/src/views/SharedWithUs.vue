<template>
  <div>
    <toolbar
      :pageTitle="pageTitle"
      :selectedFile="selectedFile"
    ></toolbar>
    <shared-file-list
      :selectedFile="selectedFile"
    ></shared-file-list>
    <file-upload-dialog />
  </div>
</template>

<script>
  import { mapState, mapActions } from 'vuex';
  import SharedFileList from '@/components/fileView/SharedFileList';
  import Toolbar from '@/components/fileView/Toolbar';
  import FileUploadDialog from '@/components/ui/FileUploadDialog';

  export default {
    name: 'SharedWithUs',
    components: {
      SharedFileList,
      Toolbar,
      FileUploadDialog
    },
    data() {
      return {
        pageTitle: 'Shared With Us'
      }
    },
    computed: {
      ...mapState('myFiles', ['selectedFile'])
    },
    methods: {
      ...mapActions('myFiles', ['updateSelectedFile'])
    },
    beforeRouteLeave(to, from, next) {
      this.updateSelectedFile(null);
      next();
    }
  };
</script>
