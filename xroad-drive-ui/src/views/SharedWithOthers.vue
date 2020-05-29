<template>
  <div>
    <toolbar
      :pageTitle="pageTitle"
      :selectedFile="selectedFile"
    ></toolbar>
    <file-list-shared-with-others
      :selectedFile="selectedFile"
    ></file-list-shared-with-others>
    <file-upload-dialog />
  </div>
</template>

<script>
  import { mapState, mapActions } from 'vuex';
  import FileListSharedWithOthers from '@/components/fileView/FileListSharedWithOthers';
  import Toolbar from '@/components/fileView/Toolbar';
  import FileUploadDialog from '@/components/ui/FileUploadDialog';

  export default {
    name: 'SharedWithOthers',
    components: {
      FileListSharedWithOthers,
      Toolbar,
      FileUploadDialog
    },
    data() {
      return {
        pageTitle: 'Shared With Others'
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
