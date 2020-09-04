<template>
  <div>
    <toolbar
      :selectedFile="selectedFile"
    ></toolbar>
    <file-list
      :selectedFile="selectedFile"
      :headers="headers"
      :group-by="group"
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
    name: 'SharedWithUs',
    components: {
      FileList,
      Toolbar,
      FileUploadDialog
    },
    data() {
      return {
        headers: [
          { text: 'Name', value: 'filename' },
          { text: 'Shared At', value: 'createdDateTime' },
          { text: 'Size', value: 'filesize' }
        ],
        group: 'createdDateTime'
      }
    },
    computed: {
      ...mapState('selectedFile', ['selectedFile'])
    },
    methods: {
      ...mapActions('selectedFile', ['updateSelectedFile'])
    }
  };
</script>
