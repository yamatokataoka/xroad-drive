<template>
  <v-list
    style="max-height: 260px; padding: 0"
    class="overflow-y-auto"
  >
    <v-list-group
      :value="listOpen"
    >
      <v-list-item
        v-for="file in files"
        :key="file.id"
      >
        <v-list-item-content>
          <v-list-item-title v-text="file.name" class="body-2"></v-list-item-title>
        </v-list-item-content>
        <v-list-item-action>
          <v-icon color="success" v-if="file.isDone">mdi-check-circle</v-icon>
        </v-list-item-action>
        <v-progress-linear
          absolute
          bottom
          :value="file.progress"
          :indeterminate="file.indeterminate"
          :color="file.progress === 100 ? file.isDone ? 'success' : 'error' : 'primary'"
        ></v-progress-linear>
        <file-upload
          :file="file"
        ></file-upload>
      </v-list-item>
    </v-list-group>
  </v-list>
</template>

<script>
  import FileUpload from '@/components/ui/FileUploadDialog/FileList/FileUpload';

  export default {
    name: 'FileList',
    components: {
      FileUpload
    },
    props: {
      files: {
        type: Array,
        required: true
      },
      listOpen: {
        type: Boolean,
        required: true
      }
    }
  };
</script>