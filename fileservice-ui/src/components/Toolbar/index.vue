<template>
  <v-toolbar
    dense
  >
    <v-toolbar-title>{{ pageTitle }}</v-toolbar-title>
    <v-spacer />
    <template v-if="selectedFile">
      <v-btn icon>
        <v-icon>mdi-trash-can-outline</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>mdi-download-outline</v-icon>
      </v-btn>
    </template>
    <v-divider
      class="mx-4"
      vertical
    ></v-divider>
    <upload-button
      @change="changeUploadFiles"
    ></upload-button>
  </v-toolbar>
</template>

<script>
  import { mapState, mapActions } from 'vuex';
  import UploadButton from '@/components/Toolbar/UploadButton';

  export default {
    name: 'Toolbar',
    components: {
      UploadButton
    },
    props: {
      pageTitle: {
        required: true
      }
    },
    computed: {
      ...mapState({
        selectedFile: state => state.files.fileList.selectedFile
      })
    },
    methods: {
      ...mapActions('files', ['updateUploadFiles', 'setUploading']),
      changeUploadFiles(refInput) {
        this.updateUploadFiles(refInput.files);
        this.setUploading(true);
        refInput.value = '';
      }
    }
  };
</script>

<style scoped>
</style>
