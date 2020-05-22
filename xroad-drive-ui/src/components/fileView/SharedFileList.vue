<template>
  <!-- todo: hard-coded height -->
  <v-data-table
    :headers="headers"
    :items="fileList"
    group-by="sharedAt"
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
    <template v-slot:item.sharedAt="{ item }">
      <span v-if="isToday(item.sharedAt)">today {{ item.sharedAt | formatHoursMins }}</span>
      <span v-else>{{ item.sharedAt | formatDate }}</span>
    </template>
    <template v-slot:item.filesize="{ item }">
      <span>{{ item.filesize | formatBytes }}</span>
    </template>
    <template v-slot:group.header="{ group, headers }">
      <th :colspan="headers.length">
        {{ group | formatDate }}
      </th>
    </template>
  </v-data-table>
</template>

<script>
  import { mapActions } from 'vuex';

  export default {
    name: 'SharedFileList',
    props: {
      selectedFile: {
        required: true
      }
    },
    data() {
      return {
        headers: [
          { text: 'Name', value: 'filename' },
          { text: 'Shared By', value: 'sharedBy' },
          { text: 'Shared At', value: 'sharedAt' },
          { text: 'Size', value: 'filesize' }
        ],
        fileList: [
          {
            id: 1,
            filename: 'file.txt',
            filesize: 333333,
            sharedBy: 'Company A',
            sharedAt: '2020-05-18T00:50:47.222'
          },
          {
            id: 2,
            filename: 'sample.txt',
            filesize: 3333330,
            sharedBy: 'Company B',
            sharedAt: '2020-05-18T00:50:47.222'
          },
          {
            id: 3,
            filename: 'details.docx',
            filesize: 400000,
            sharedBy: 'Company B',
            sharedAt: '2020-05-10T00:50:47.222'
          },
          {
            id: 4,
            filename: 'company-a-file.txt',
            filesize: 333333,
            sharedBy: 'Company A',
            sharedAt: '2020-05-20T00:50:47.222'
          }
        ]
      }
    },
    computed: {
      // ...mapState('fileList', ['fileList'])
    },
    methods: {
      ...mapActions('myFiles', ['updateSelectedFile']),
      // ...mapActions('fileList', ['fetchFileList']),
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
      // this.pollFileList()
    },
    beforeRouteLeave() {
      // clearInterval(this.polling)
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
