<template>
  <!-- todo: hard-coded height -->
  <v-data-table
    :headers="headers"
    :items="fileListWithDateGroup"
    group-by="dateGroup"
    fixed-header
    single-select
    @click:row="clickRow"
    height="calc(100vh - 96px - 58px)"
    item-key="id"
    :search="search"
  >
    <template v-slot:no-data>No Files</template>
    <template v-slot:item.filename="{ item }">
      <v-icon class="me-4">mdi-file-outline</v-icon>
      <span>{{ item.filename }}</span>
    </template>
    <template v-slot:item.createdDateTime="{ item }">
      <span v-if="isToday(item.createdDateTime)">today {{ item.createdDateTime | formatHoursMins }}</span>
      <span v-else>{{ item.createdDateTime | formatDate }}</span>
    </template>
    <template v-slot:item.filesize="{ item }">
      <span>{{ item.filesize | formatBytes }}</span>
    </template>
    <template v-slot:group.header="{ group, headers }">
      <th :colspan="headers.length">
        {{ group }}
      </th>
    </template>
  </v-data-table>
</template>

<script>
  import { mapState, mapActions } from 'vuex';
  import { isToday } from '@/utils';

  export default {
    name: 'FileListSharedWithOthers',
    props: {
      selectedFile: {
        required: true
      }
    },
    data() {
      return {
        headers: [
          { text: 'Name', value: 'filename' },
          { text: 'Shared At', value: 'createdDateTime' },
          { text: 'Size', value: 'filesize' }
        ]
      }
    },
    computed: {
      ...mapState('fileList', ['fileList']),
      ...mapState('search', ['search']),
      fileListWithDateGroup: function() {
        return this.fileList.map(item => ({
          ...item, dateGroup: this.$options.filters.dateGroup(item.createdDateTime)
        }));
      }
    },
    methods: {
      ...mapActions('selectedFile', ['updateSelectedFile']),
      ...mapActions('fileList', ['fetchFileList']),
      ...mapActions('search', ['updateSearch']),
      clickRow(item, row) {
        if (this.selectedFile && row.isSelected) {
          row.select(false);
          this.updateSelectedFile(null);
        } else {
          row.select(true);
          this.updateSelectedFile(item);
        }
      },
      pollFileList() {
        this.fetchFileList();

        this.polling = setInterval(() => {
          this.fetchFileList();
        }, 3000)
      },
      isToday
    },
    created() {
      this.pollFileList()
    },
    beforeDestroy() {
      clearInterval(this.polling);
      this.updateSelectedFile(null);
      this.updateSearch('');
    }
  };
</script>

<style scoped>
  /* todo: hard-coded fixes */
  >>>.v-row-group__header {
    background: unset!important;
  }
  >>>td:not(.v-data-table__mobile-row) {
    border-bottom: unset!important;
  }
</style>
