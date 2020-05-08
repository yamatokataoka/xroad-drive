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
      ...mapActions('uploadFiles', [
        'upload',
        'updateProgressById',
        'updateIsDoneById',
        'updateIndeterminateById'
      ])
    },
    async mounted() {
      const { file } = this;
      const id = file.id;

      if (!file) return;
      console.log(`Start to upload file: ${file.name}`)
      try {
        await this.upload(file);
        this.updateIsDoneById({ id, isDone: true });
        this.updateProgressById({ id, progress: 100 });
        console.log(`Succeeded to upload file: ${file.name}`);
      } catch (error) {
        this.updateIsDoneById({ id, isDone: false });
        this.updateProgressById({ id, progress: 100 });
        this.updateIndeterminateById({ id, indeterminate: false });
        console.log(`Failed to upload file: ${file.name}`);
      }
    },
    render: () => null,
  };
</script>