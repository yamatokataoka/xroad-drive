<template>
  <div>
    <input
      style="display: none"
      ref="input"
      type="file"
      multiple
      @change="emitInputFiles()"
    >
    <v-btn icon @click="clickInput">
      <v-icon>mdi-file-upload-outline</v-icon>
    </v-btn>
  </div>
</template>

<script>
  export default {
    name: 'UploadButton',
    methods: {
      clickInput() {
        return this.$refs.input.click();
      },
      emitInputFiles() {
        this.$emit('change', this.$refs.input);
      }
    }
  };
</script>
