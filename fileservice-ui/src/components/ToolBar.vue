<template>
  <v-toolbar
    app
    dense
  >
    <v-toolbar-title>{{ pageTitle }}</v-toolbar-title>
    <v-spacer />
    <v-btn icon>
      <v-icon>mdi-trash-can-outline</v-icon>
    </v-btn>
    <v-btn icon>
      <v-icon>mdi-download-outline</v-icon>
    </v-btn>
    <v-divider
      class="mx-4"
      vertical
    ></v-divider>
    <v-btn icon>
      <v-icon>mdi-file-upload-outline</v-icon>
    </v-btn>
  </v-toolbar>
</template>

<script>
export default {
  name: 'ToolBar',
  data: () => ({
    pageTitle: 'MyFiles',
  }),
};
</script>

<style scoped>
</style>