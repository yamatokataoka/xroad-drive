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
    name: 'OurFiles',
    component: createFileListView('OurFiles')
  }, {
    path: '/providers',
    name: 'SharedWithUs',
    component: createFileListView('SharedWithUs')
  }, {
    path: '/clients',
    name: 'SharedWithOthers',
    component: createFileListView('SharedWithOthers')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
