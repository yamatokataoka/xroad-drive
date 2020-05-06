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

  export default {
    name: 'MyFiles',
    components: {
      FileList,
      Toolbar,
      FileUpload
    },
  };
</script>
