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

export default {
  name: 'App',
  components: { AppBar },
};
</script>

<style></style>
