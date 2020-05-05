<template>
  <v-toolbar
    dense
  >
    <v-toolbar-title>{{ pageTitle }}</v-toolbar-title>
    <v-spacer />
    <v-btn v-if="selectedFile" icon>
      <v-icon>mdi-trash-can-outline</v-icon>
    </v-btn>
    <v-btn v-if="selectedFile" icon>
      <v-icon>mdi-download-outline</v-icon>
    </v-btn>
    <v-divider
      class="mx-4"
      vertical
    ></v-divider>
    <upload-button
      @change="changeInputFiles"
    ></upload-button>
  </v-toolbar>
</template>

<script>
  import { mapState, mapActions } from 'vuex';
  import UploadButton from '@/components/ToolBar/UploadButton';

  export default {
    name: 'ToolBar',
    components: {
      UploadButton
    },
    data() {
      return {
        pageTitle: 'MyFiles'
      }
    },
    computed: {
      ...mapState('myFiles', ['selectedFile'])
    },
    methods: {
      ...mapActions('inputFile', ['updateInputFiles', 'setUploading']),
      changeInputFiles(refInput) {
        this.updateInputFiles(refInput.files);
        this.setUploading(true);
        refInput.value = '';
      }
    }
  };
</script>

<style scoped>
</style>
