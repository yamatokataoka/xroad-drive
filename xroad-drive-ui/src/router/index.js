import Vue from 'vue';
import VueRouter from 'vue-router';

import createFileListView from '@/views/CreateFileListView.js';

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/our-files'
  }, {
    path: '/our-files',
    component: createFileListView('OurFiles'),
    meta: {
      id: 'our-files'
    }
  }, {
    path: '/shared-with-us/:id(\\w+:\\w+:\\d+:.+)?',
    component: createFileListView('SharedWithUs'),
    meta: {
      id: route => route.params.id ? 'provider' + ':' + route.params.id : 'provider',
      dynamic: true
    }
  }, {
    path: '/shared-with-others/:id(\\w+:\\w+:\\d+:.+)?',
    component: createFileListView('SharedWithOthers'),
    meta: {
      id: route => route.params.id ? 'client' + ':' + route.params.id : 'client',
      dynamic: true
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
