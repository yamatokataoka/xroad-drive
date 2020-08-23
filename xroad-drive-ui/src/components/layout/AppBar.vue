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
  import { mapGetters, mapActions } from 'vuex';

  export default {
    name: 'AppBar',
    data() {
      return {
        drawer: null,
        appTitle: 'X-Road Drive',
        active: [],
        open: [],
        rootNavItems: [
          {
            id: 'our-files',
            name: 'Our Files',
            icon: 'mdi-file-multiple-outline',
            to: '/our-files'
          }, {
            id: 'shared-with-us',
            name: 'Shared with us',
            icon: 'mdi-account-multiple-outline',
            to: '/shared-with-us',
            children: []
          }, {
            id: 'shared-with-others',
            name: 'Shared with others',
            icon: 'mdi-office-building',
            to: '/shared-with-others',
            children: []
          }
        ]
      }
    },
    computed: {
      ...mapGetters('xroadMetadata', ['providersForNav', 'clientsForNav']),
      navItems() {
        let navItems = [];
        for (let i = 0; i < this.rootNavItems.length; i++) {
          navItems[i] = {...this.rootNavItems[i]};
          if (this.rootNavItems[i].children) {
            if (this.rootNavItems[i].id === 'shared-with-us') {
              navItems[i].children = this.providersForNav(this.rootNavItems[i].id);
            } else if (this.rootNavItems[i].id === 'shared-with-others') {
              navItems[i].children = this.clientsForNav(this.rootNavItems[i].id);
            }
          }
        }
        return navItems;
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
      }
    },
    methods: {
      ...mapActions('xroadMetadata', ['fetchProviders', 'fetchClients']),
      selected(event) {
        if (!event[0] || !this.active[0] || event[0] === this.active[0]) return;
        const item = this.findObjectById(this.navItems, event[0]);
        if (item) {
          this.active = [item.id];
          this.openParent(item.id);
          this.$router.push(item.to);
        }
      },
      findObjectById(object, id) {
        if (typeof object !== 'object' || object === null) return null;
        if (object.id === id) return object;
        for(let key in object) {
          if (Object.prototype.hasOwnProperty.call(object, key)){
            const foundObject = this.findObjectById(object[key], id);
            if (foundObject) return foundObject;
          }
        }
        return null;
      },
      openParent(id) {
        const item = this.findObjectById(this.navItems, id);
        if (!item) return;

        if (!this.open.length && item.parent) {
          this.open = [item.parent];
        }
      },
      async pollNavitems() {
        await this.fetchProviders();
        await this.fetchClients();

        this.polling = setInterval(() => {
          this.fetchProviders();
          this.fetchClients();
        }, 3000)
      }
    },
    async created() {
      await this.pollNavitems();

      const currentRoute = this.$router.currentRoute;
      const id = currentRoute.params.id
        ? currentRoute.name + ':' + currentRoute.params.id
        : currentRoute.name;
      if (!id) return;

      this.active = [id];
      this.openParent(id);
    },
    beforeDestroy() {
      clearInterval(this.polling)
    }
  };
</script>

<style scoped></style>
