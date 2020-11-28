<template>
  <v-app>
    <app-bar />
    <v-content
      style="position: fixed; width: 100%"
    >
      <router-view />
    </v-content>
  </v-app>
</template>

<script>
import AppBar from '@/components/layout/AppBar';
import { mapActions } from 'vuex';

export default {
  name: 'App',
  components: { AppBar },
  methods: {
    ...mapActions('config', ['fetchConfig'])
  },
  async created() {
    console.log('Fetch config');
    await this.fetchConfig();
  }
};
</script>

<style></style>
