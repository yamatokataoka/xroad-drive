<template>
  <!-- todo: hard-coded height -->
  <v-data-table
    :headers="headers"
    :items="fileListWithDateGroupIndex"
    group-by="DateGroupIndex"
    hide-default-footer
    fixed-header
    single-select
    @click:row="clickRow"
    height="calc(100vh - 96px)"
    item-key="id"
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
        {{ group | getDateGroupByIndex }}
      </th>
    </template>
  </v-data-table>
</template>

<script>
  import { mapState, mapActions } from 'vuex';

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
      fileListWithDateGroupIndex: function() {
        return this.fileList.map(item => ({
          ...item, DateGroupIndex: this.$options.filters.getDateGroupIndex(item.createdDateTime)
        }));
      }
    },
    methods: {
      ...mapActions('selectedFile', ['updateSelectedFile']),
      ...mapActions('fileList', ['fetchFileList']),
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
      isToday(dateString) {
        const inputDate = this.$options.filters.formatDate(dateString);
        const today = new Date().toDateString();
        return (today == inputDate);
      }
    },
    created() {
      this.pollFileList()
    },
    beforeRouteLeave() {
      clearInterval(this.polling)
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
