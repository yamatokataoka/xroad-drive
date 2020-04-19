<template>
  <nav>
    <v-navigation-drawer
      app
      clipped
      fixed
      v-model="drawer"
      mobile-break-point="960"
    >
      <v-treeview
        :items="nav_lists"
        activatable
        shaped
        dense
        item-key="name"
        hoverable
      >
        <template v-slot:prepend="{ item }">
          <v-icon
            v-if="item.icon"
            dense
          >
            {{ item.icon }}
          </v-icon>
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
  </nav>
</template>

<script>
export default {
  name: 'AppNavigation',
  data: () => ({
    drawer: null,
    appTitle: 'File Service',
    nav_lists: [
      {
        name: 'My Files',
        icon: 'mdi-file-multiple-outline'
      }, {
        name: 'Shared with us',
        icon: 'mdi-account-multiple-outline',
        children: [{
          name: 'Company A'
        }, {
          name: 'Company B' 
        }, {
          name: 'Something'
        }]
      }, {
        name: 'Shared with companies',
        icon: 'mdi-office-building',
        children: [{
          name: 'Company C'
        }, {
          name: 'Company D' 
        }, {
          name: 'Something'
        }]
      }
    ],
  }),
};
</script>

<style scoped></style>
