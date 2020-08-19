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
    name: 'our-files',
    component: createFileListView('OurFiles')
  }, {
    path: '/shared-with-us/:id(\\w+:\\w+:\\d+:.+)?',
    name: 'shared-with-us',
    component: createFileListView('SharedWithUs')
  }, {
    path: '/shared-with-others/:id(\\w+:\\w+:\\d+:.+)?',
    name: 'shared-with-others',
    component: createFileListView('SharedWithOthers')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
