<template>
  <div>
    <v-navigation-drawer
      app
      clipped
      fixed
      v-model="drawer"
      mobile-break-point="960"
    >
      <v-treeview
        :items="nav_items"
        activatable
        shaped
        dense
        item-key="name"
        hoverable
        @update:active="OnUpdateActive"
        class="body-2"
      >
        <template v-slot:prepend="{ item }">
          <v-icon v-if="item.icon" dense>{{ item.icon }}</v-icon>
        </template>
      </v-treeview>
    </v-navigation-drawer>
    <v-app-bar
      app
      clipped-left
      dense
      flat
    >
      <v-app-bar-nav-icon class="hidden-md-and-up" @click="drawer = !drawer" />
      <v-toolbar-title>{{ appTitle }}</v-toolbar-title>
      <v-spacer />
      <v-text-field
        flat
        solo-inverted
        hide-details
        prepend-inner-icon="mdi-magnify"
        label="Search"
      />
      <v-spacer />
      <v-btn outlined class="hidden-sm-and-down">SIGN IN</v-btn>
    </v-app-bar>
  </div>
</template>

<script>
  export default {
    name: 'AppBar',
    data() {
      return {
        drawer: null,
        appTitle: 'X-Road Drive',
        nav_items: [
          {
            name: 'My Files',
            icon: 'mdi-file-multiple-outline',
            to: '/my-files'
          }, {
            name: 'Shared with us',
            icon: 'mdi-account-multiple-outline',
            to: '/shared-with-us',
            children: [
              {
                name: 'Company A'
              }, {
                name: 'Company B'
              }, {
                name: 'Something'
              }
            ]
          }, {
            name: 'Shared with companies',
            icon: 'mdi-office-building',
            to: '/shared-with-companies',
            children: [
              {
                name: 'Company C'
              }, {
                name: 'Company D'
              }, {
                name: 'Something'
              }
            ]
          }
        ],
      }
    },
    methods: {
      OnUpdateActive(event) {
        let item = this.nav_items.find((item) => {
          return (item.name === event[0]);
        });
        this.$router.push(item.to);
      }
    }
  };
</script>

<style scoped></style>
