<template>
  <v-toolbar
    dense
  >
    <v-breadcrumbs :items="breadcrumbs">
      <template v-slot:divider>
        <v-icon size="20">mdi-chevron-right</v-icon>
      </template>
      <template v-slot:item="{ item }">
        <v-btn
          text
          class="subtitle-1 breadcrumbs-btn--active"
          :to="item.to"
          exact
        >
          {{ item.text }}
        </v-btn>
      </template>
    </v-breadcrumbs>
    <v-spacer />
    <template v-if="selectedFile">
      <v-btn
        icon
        @click="clickDelete"
      >
        <v-icon>mdi-trash-can-outline</v-icon>
      </v-btn>
      <v-btn
        icon
        @click="clickDownload"
      >
        <v-icon>mdi-download-outline</v-icon>
      </v-btn>
    </template>
    <v-divider
      class="mx-4"
      vertical
    ></v-divider>
    <upload-button
      @change="changeUploadFiles"
    ></upload-button>
  </v-toolbar>
</template>

<script>
  import { mapState, mapActions } from 'vuex';
  import UploadButton from '@/components/fileView/Toolbar/UploadButton';
  import axios from 'axios';

  export default {
    name: 'Toolbar',
    components: {
      UploadButton
    },
    props: {
      selectedFile: {
        required: true
      }
    },
    computed: {
      breadcrumbs() {
        let pathArray = this.$route.path.split("/")

        // not need the first item as it's empty
        pathArray.shift()

        const breadcrumbs = pathArray.reduce((breadcrumbArray, path, i) => {
          breadcrumbArray.push({
            text: this.strToTile(path),
            to: breadcrumbArray[i - 1]
            ? "/" + pathArray[i - 1] + "/" + path
            : "/" + path,
          });
          return breadcrumbArray;
        }, [])
        return breadcrumbs;
      },
      ...mapState('selectedXRoadMember', ['selectedXRoadMember']),
      ...mapState('xroadMetadata', ['providers', 'clients']),
    },
    methods: {
      ...mapActions('uploadFiles', ['updateUploadFiles', 'updateUploading']),
      ...mapActions('selectedFile', ['updateSelectedFile']),
      changeUploadFiles(refInput) {
        this.updateUploadFiles(refInput.files);
        this.updateUploading(true);
        refInput.value = '';
      },
      strToTile(str) {
        let splitedArray = str.split('-');
        let outArray = [];
        const searchArray = this.providers.concat(this.clients);

        const matchedMember = searchArray.find(xroadMember => xroadMember && xroadMember.id === str);

        if (matchedMember) {
          return matchedMember.name;
        }
        for(var x = 0; x < splitedArray.length; x++){
          outArray.push(splitedArray[x].charAt(0).toUpperCase() + splitedArray[x].slice(1));
        }

        let outStr = outArray.join(' ');

        if (str == 'our-files') {
          return outStr
        }
        return outStr.charAt(0) + outStr.slice(1).toLowerCase();
      },
      async clickDownload() {
        if (selectedFile === null) {
          console.log('File is not selected');
          return;
        }
        const { selectedFile } = this;
        const id = selectedFile.id;

        let serviceId = null;
        if (this.selectedXRoadMember) {
          serviceId = this.selectedXRoadMember + ':' + 'XRoadDrive';
        }

        const xroadMemberId = process.env.VUE_APP_XROAD_MEMBER_ID;
        let path = `/api/download/${encodeURIComponent(id)}`;
        let headers = {};

        if (serviceId) {
          const servicePath = serviceId.replace(/:/g, '/');
          if (!xroadMemberId) {
            console.log('Failed to set config for axios: xroadMemberId: ' + xroadMemberId);
            return;
          }
          // TODO: URL join helper is needed.
          path = '/security-server/r1/' + servicePath + path;
          headers['X-Road-Client'] = xroadMemberId;
        }

        try {
          const response = await axios({
            url: path,
            method: 'get',
            headers: headers,
            responseType: 'blob'
          });
          const regexp = /filename="(.*)"/;
          const contentDisposition = response.headers['content-disposition'];
          const fileName = regexp.exec(contentDisposition)[1];
          const url = URL.createObjectURL(new Blob([response.data]));
          const link = window.document.createElement("a");
          link.href = url;
          if (window.navigator.msSaveBlob) {
            return window.navigator.msSaveBlob(response, fileName);
          } else {
            link.download = fileName;
            link.click();
            link.remove();
          }
          setTimeout(function() {
            window.URL.revokeObjectURL(url);
          }, 1000);
          console.log('Succeeded to download file');
        } catch (error) {
          console.log('Failed to download file');
          console.log(error);
        }
      },
      async clickDelete() {
        if (selectedFile === null) {
          console.log('File is not selected');
          return;
        }

        const { selectedFile } = this;
        const id = selectedFile.id;

        let serviceId = null;
        if (this.selectedXRoadMember) {
          serviceId = this.selectedXRoadMember + ':' + 'XRoadDrive';
        }

        const xroadMemberId = process.env.VUE_APP_XROAD_MEMBER_ID;
        let path = `/api/delete/${encodeURIComponent(id)}`;
        let headers = {};

        if (serviceId) {
          const servicePath = serviceId.replace(/:/g, '/');
          if (!xroadMemberId) {
            console.log('Failed to set config for axios: xroadMemberId: ' + xroadMemberId);
            return;
          }
          // TODO: URL join helper is needed.
          path = '/security-server/r1/' + servicePath + path;
          headers['X-Road-Client'] = xroadMemberId;
        }

        try {
          await axios({
            url: path,
            method: 'delete',
            headers: headers
          });
          this.updateSelectedFile(null);
          console.log('Succeeded to delete file');
        } catch (error) {
          console.log('Failed to delete file');
          console.log(error);
        }
      }
    }
  };
</script>

<style scoped>
.v-btn {
  text-transform: none !important;
}
ul.v-breadcrumbs {
  padding-left: 0;
}
.v-breadcrumbs >>> li:nth-child(even) {
  padding: 0;
}
.breadcrumbs-btn--active::before {
  background-color: transparent;
  opacity: 0.08;
}
.breadcrumbs-btn--active:hover::before {
  background-color: currentColor;
  opacity: 0.08;
}
</style>
