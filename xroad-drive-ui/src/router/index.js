import Vue from 'vue';
import VueRouter from 'vue-router';

import createFileListView from '@/views/CreateFileListView.js';

Vue.use(VueRouter)

const idRegex = '\\w+:\\w+:\\d+:.+'

const routes = [
  {
    path: '/',
    redirect: '/our-files'
  }, {
    path: '/our-files',
    component: createFileListView('OurFiles')
  }, {
    path: `/providers/:id(${idRegex})?`,
    component: createFileListView('SharedWithUs')
  }, {
    path: `/clients/:id(${idRegex})?`,
    component: createFileListView('SharedWithOthers')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
