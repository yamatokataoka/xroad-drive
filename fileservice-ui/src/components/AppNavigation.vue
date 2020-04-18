<template>
  <nav>
    <v-navigation-drawer
      app
      clipped
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
      <v-toolbar-title>{{ appTitle }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn outlined>SIGN IN</v-btn>
    </v-app-bar>
  </nav>
</template>

<script>
export default {
  name: 'AppNavigation',
  data() {
    return {
      appTitle: 'File Service',
      nav_lists: [{
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
      }]
    };
  }
};
</script>

<style scoped></style>
