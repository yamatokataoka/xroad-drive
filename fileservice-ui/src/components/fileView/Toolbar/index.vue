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
  import { mapActions } from 'vuex';
  import UploadButton from '@/components/fileView/Toolbar/UploadButton';

  export default {
    name: 'Toolbar',
    components: {
      UploadButton
    },
    props: {
      pageTitle: {
        type: String,
        required: true
      },
      selectedFile: {
        required: true
      }
    },
    methods: {
      ...mapActions('uploadFiles', ['updateUploadFiles', 'setUploading']),
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
