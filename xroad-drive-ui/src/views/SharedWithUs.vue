<template>
  <div>
    <toolbar
      :pageTitle="pageTitle"
      :selectedFile="selectedFile"
    ></toolbar>
    <file-list-shared-with-us
      :selectedFile="selectedFile"
    ></file-list-shared-with-us>
    <file-upload-dialog />
  </div>
</template>

<script>
  import { mapState, mapActions } from 'vuex';
  import FileListSharedWithUs from '@/components/fileView/FileListSharedWithUs';
  import Toolbar from '@/components/fileView/Toolbar';
  import FileUploadDialog from '@/components/ui/FileUploadDialog';

  export default {
    name: 'SharedWithUs',
    components: {
      FileListSharedWithUs,
      Toolbar,
      FileUploadDialog
    },
    data() {
      return {
        pageTitle: 'Shared With Us'
      }
    },
    computed: {
      ...mapState('selectedFile', ['selectedFile'])
    },
    methods: {
      ...mapActions('selectedFile', ['updateSelectedFile'])
    },
    beforeRouteLeave(to, from, next) {
      this.updateSelectedFile(null);
      next();
    }
  };
</script>
