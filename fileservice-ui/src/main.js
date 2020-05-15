import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import vuetify from './plugins/vuetify';
import * as filters from './filters';
import axios from 'axios';

Vue.config.productionTip = false;

Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key]);
});

axios.defaults.baseURL = process.env.VUE_APP_API_URL || 'http://localhost:8082';

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app');
