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
        :items="navItems"
        activatable
        shaped
        dense
        item-key="id"
        hoverable
        @update:active="selected"
        class="body-2"
        :active.sync="active"
        :open.sync="open"
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
  import { mapState } from 'vuex';

  export default {
    name: 'AppBar',
    data() {
      return {
        drawer: null,
        appTitle: 'X-Road Drive',
        active: [],
        open: []
      }
    },
    computed: {
      ...mapState('xroadMetadata', ['providers', 'clients']),
      navItems() {
        return [
          {
            id: 'our-files',
            name: 'Our Files',
            icon: 'mdi-file-multiple-outline',
            to: '/our-files'
          }, {
            id: 'provider',
            name: 'Shared with us',
            icon: 'mdi-account-multiple-outline',
            to: '/shared-with-us',
            children: this.providers
          }, {
            id: 'client',
            name: 'Shared with others',
            icon: 'mdi-office-building',
            to: '/shared-with-others',
            children: this.clients
          }
        ];
      }
    },
    watch: {
      // Keep active state, Even if the item was clicked again.
      active: {
        handler(newActive, oldActive) {
          if (!newActive.length && oldActive.length) {
            newActive[0] = oldActive[0];
          }
        }, deep: true
      },
      $route: {
        handler(to, from) {
          const id = to.meta.dynamic ? to.meta.id(to) : to.meta.id;
          if (!id) return;

          // When from === 'undefined', it's just after reload.
          if (typeof from === 'undefined' || to.path !== from.path) {
            this.active = [id];
            this.openParent(id);
          }
        }, immediate: true
      }
    },
    methods: {
      selected(event) {
        if (!event[0] || !this.active[0] || event[0] === this.active[0]) return;
        const item = this.findObjectById(this.navItems, event[0]);
        if (item) {
          this.$router.push(item.to);
        }
      },
      findObjectById(object, id) {
        if (typeof object !== 'object' || object === null) return null;
        if (object.id === id) return object;
        for(let i in object) {
          if (Object.prototype.hasOwnProperty.call(object, i)){
            const foundObject = this.findObjectById(object[i], id);
            if (foundObject) return foundObject;
          }
        }
        return null;
      },
      openParent(id) {
        const item = this.findObjectById(this.navItems, id);
        if (item === null) return;

        if (!this.open.length && item.type) {
          this.open = [item.type];
        }
      }
    }
  };
</script>

<style scoped></style>
