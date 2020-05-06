<template>
  <v-toolbar>
    <v-btn
      icon
      @click="clickClose"
    >
      <v-icon>mdi-close</v-icon>
    </v-btn>
    <v-toolbar-title class="body-2">Uploading...</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn
      icon
      v-if="listOpen"
      @click="emitClickChevron"
    ><v-icon>mdi-chevron-up</v-icon>
    </v-btn>
    <v-btn
      icon
      v-else
      @click="emitClickChevron"
    ><v-icon>mdi-chevron-down</v-icon>
    </v-btn>
  </v-toolbar>
</template>

<script>
  import { mapActions } from "vuex";

  export default {
    name: 'FileListbar',
    props: {
      listOpen: {
        required: true
      }
    },
    data() {
      return {
        value: 50,
        uploadFinished: false
      }
    },
    methods: {
      ...mapActions('myFiles', ['setUploading', 'deleteUploadFiles']),
      clickClose() {
        this.setUploading(false);
        this.deleteUploadFiles();
      },
      emitClickChevron() {
        this.$emit('clickChevron');
      }
    }
  };
</script>