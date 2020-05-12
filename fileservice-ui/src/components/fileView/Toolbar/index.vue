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
      <v-btn
        icon
        @click="clickDownload"
      >
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
  import axios from 'axios';

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
      ...mapActions('uploadFiles', ['updateUploadFiles', 'updateUploading']),
      changeUploadFiles(refInput) {
        this.updateUploadFiles(refInput.files);
        this.updateUploading(true);
        refInput.value = '';
      },
      async clickDownload() {
        if (selectedFile === null) {
          console.log('File is not selected');
          return;
        }
        const { selectedFile } = this;
        const id = selectedFile.id;

        try {
          const response = await axios({
            url: `/api/download/${encodeURIComponent(id)}`,
            method: 'get',
            responseType: 'blob'
          });
          const regexp = /filename="(.*)"/;
          const contentDisposition = response.headers['content-disposition'];
          const fileName = regexp.exec(contentDisposition)[1];
          const url = URL.createObjectURL(new Blob([response.data]));
          const link = window.document.createElement("a");
          link.href = url;
          if (window.navigator.msSaveBlob) {
            return window.navigator.msSaveBlob(response, fileName);
          } else {
            link.download = fileName;
            link.click();
            link.remove();
          }
          setTimeout(function() {
            window.URL.revokeObjectURL(url);
          }, 1000);
        } catch (error) {
          console.log('Failed to download file');
          console.log(error);
        }
      }
    }
  };
</script>

<style scoped>
</style>
