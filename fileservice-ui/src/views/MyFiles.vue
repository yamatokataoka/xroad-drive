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
  import { mapState } from 'vuex';
  import FileList from '@/components/FileList';
  import Toolbar from '@/components/Toolbar';
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
    }
  };
</script>
