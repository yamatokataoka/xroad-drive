<template>
  <div>
    <toolbar
      :pageTitle="pageTitle"
      :selectedFile="selectedFile"
    ></toolbar>
    <file-list
      :selectedFile="selectedFile"
    ></file-list>
    <file-upload-dialog />
  </div>
</template>

<script>
  import { mapState, mapActions } from 'vuex';
  import FileList from '@/components/fileView/FileList';
  import Toolbar from '@/components/fileView/Toolbar';
  import FileUploadDialog from '@/components/ui/FileUploadDialog';

  export default {
    name: 'MyFiles',
    components: {
      FileList,
      Toolbar,
      FileUploadDialog
    },
    data() {
      return {
        pageTitle: 'My Files'
      }
    },
    computed: {
      ...mapState('myFiles', ['selectedFile'])
    },
    methods: {
      ...mapActions('myFiles', ['deleteSelectedFile'])
    },
    beforeRouteLeave(to, from, next) {
      this.deleteSelectedFile();
      next();
    }
  };
</script>
