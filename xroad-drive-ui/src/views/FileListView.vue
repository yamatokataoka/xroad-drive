<template>
  <div>
    <toolbar
      :selectedFile="selectedFile"
    ></toolbar>
    <file-list
      v-if="type === 'OurFiles'"
      :selectedFile="selectedFile"
    ></file-list>
    <file-list-shared-with-us
      v-else-if="type === 'SharedWithUs'"
      :selectedFile="selectedFile"
    ></file-list-shared-with-us>
    <file-list-shared-with-others
      v-else
      :selectedFile="selectedFile"
    ></file-list-shared-with-others>
    <file-upload-dialog />
  </div>
</template>

<script>
  import { mapState, mapActions } from 'vuex';
  import FileList from '@/components/fileView/FileList';
  import Toolbar from '@/components/fileView/Toolbar';
  import FileUploadDialog from '@/components/ui/FileUploadDialog';
  import FileListSharedWithUs from '@/components/fileView/FileListSharedWithUs';
  import FileListSharedWithOthers from '@/components/fileView/FileListSharedWithOthers';

  export default {
    name: 'FileListView',
    components: {
      FileList,
      Toolbar,
      FileUploadDialog,
      FileListSharedWithUs,
      FileListSharedWithOthers
    },
    props: {
      type: String
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
