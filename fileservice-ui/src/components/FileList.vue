<template>
  <v-data-table
    :headers="headers"
    :items="fileList"
    hide-default-footer
    fixed-header
    single-select
    @click:row="rowClick"
    class="elevation-1"
  >
    <template v-slot:no-data>No Files</template>
  </v-data-table>
</template>

<script>
export default {
  name: 'FileList',
  data: () => ({
    selectedId: -1,
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
      { text: 'Size', value: 'size' },
    ],
  }),
  methods: {
    rowClick: function (item, row) {
      if (item.name == this.selectedId) {
        row.select(false);
        this.selectedId='';
      } else {
        row.select(true);
        this.selectedId=item.name;
      }
    }
  },
};
</script>
<style scoped>
</style>