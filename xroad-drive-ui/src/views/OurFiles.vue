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
    name: 'OurFiles',
    components: {
      FileList,
      Toolbar,
      FileUploadDialog
    },
    data() {
      return {
        pageTitle: 'Our Files'
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
