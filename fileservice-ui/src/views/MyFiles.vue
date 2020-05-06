<template>
  <div>
    <toolbar />
    <file-list />
    <file-upload />
  </div>
</template>

<script>
  import FileList from '@/components/FileList';
  import Toolbar from '@/components/Toolbar';
  import FileUpload from '@/components/FileUpload';
  import files from '@/store/modules/files';

  export default {
    name: 'MyFiles',
    components: {
      FileList,
      Toolbar,
      FileUpload
    },
    created () {
      const store = this.$store;
      store.registerModule('files', files);
    },
    beforeRouteLeave (to, from, next) {
      const store = this.$store;
      store.unregisterModule('files');
      next();
    },
  };
</script>
