<template>
  <v-data-table
    :headers="headers"
    :items="fileList"
    hide-default-footer
    fixed-header
    single-select
    @click:row="clickRow"
    class="elevation-1"
  >
    <template v-slot:no-data>No Files</template>
  </v-data-table>
</template>

<script>
  import { mapState, mapActions } from 'vuex';

  export default {
    name: 'FileList',
    data() {
      return {
        fileList: [
          {
            id: 1,
            name: 'file_a.txt',
            owner: 'test_user',
            lastModifiedDate: '09.05.2020',
            size: '100 MB'
          }, {
            id: 2,
            name: 'file_b.txt',
            owner: 'test_user',
            lastModifiedDate: '09.06.2020',
            size: '150 KB'
          }, {
            id: 3,
            name: 'file_c.txt',
            owner: 'test_user',
            lastModifiedDate: '09.07.2020',
            size: '250 KB'
          }
        ],
        headers: [
          { text: 'Name', value: 'name' },
          { text: 'Owner', value: 'owner' },
          { text: 'Last modified', value: 'lastModifiedDate' },
          { text: 'Size', value: 'size' }
        ]
      }
    },
    computed: {
      ...mapState({
        selectedFile: state => state.myFiles.fileList.selectedFile
      })
    },
    methods: {
      ...mapActions('myFiles', ['updateSelectedFile']),
      clickRow(item, row) {
        if (this.selectedFile && row.isSelected) {
          row.select(false);
          this.updateSelectedFile(null);
        } else {
          row.select(true);
          this.updateSelectedFile(item);
        }
      }
    },
  };
</script>

<style>
</style>
