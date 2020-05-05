<template>
  <div>
    <tool-bar />
    <file-list />
    <file-upload-snack-bar />
  </div>
</template>

<script>
  import FileList from '@/components/FileList';
  import ToolBar from '@/components/ToolBar';
  import FileUploadSnackBar from '@/components/FileUploadSnackBar';

  export default {
    name: 'MyFiles',
    components: {
      FileList,
      ToolBar,
      FileUploadSnackBar
    },
  };
</script>
