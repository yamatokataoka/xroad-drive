<template>
  <div>
    <toolbar
      :pageTitle="pageTitle"
      :selectedFile="selectedFile"
    ></toolbar>
    <file-list />
    <file-upload />
  </div>
</template>

<script>
  import { mapState, mapActions } from 'vuex';
  import FileList from '@/components/FileList';
  import Toolbar from '@/components/fileView/Toolbar';
  import FileUpload from '@/components/FileUpload';

  export default {
    name: 'MyFiles',
    components: {
      FileList,
      Toolbar,
      FileUpload
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
