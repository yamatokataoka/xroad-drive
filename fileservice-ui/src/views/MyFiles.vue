<template>
  <div>
    <toolbar />
    <file-list />
    <file-upload-snackbar />
  </div>
</template>

<script>
  import FileList from '@/components/FileList';
  import Toolbar from '@/components/Toolbar';
  import FileUploadSnackbar from '@/components/FileUploadSnackbar';

  export default {
    name: 'MyFiles',
    components: {
      FileList,
      Toolbar,
      FileUploadSnackbar
    },
  };
</script>
