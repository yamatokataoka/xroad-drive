<template>
</template>

<script>
  import { mapActions } from "vuex";

  export default {
    name: 'FileUpload',
    props: {
      file: {
        required: true
      }
    },
    methods: {
      ...mapActions('files', [
        'upload',
        'updateProgressById',
        'updateIsDoneById',
        'updateIndeterminateById'
      ])
    },
    async mounted() {
      const { file } = this;
      if (!file) return;
      try {
        await this.upload(file);
        this.updateIsDoneById({ id: file.id, isDone: true });
        this.updateProgressById({ id: file.id, progress: 100 });
        console.log("Succeeded to upload");
      } catch (error) {
        this.updateIsDoneById({ id: file.id, isDone: false });
        this.updateProgressById({ id: file.id, progress: 100 });
        this.updateIndeterminateById({ id: file.id, indeterminate: false });
        console.log("Failed to upload");
      }
    }
  };
</script>