import Vue from 'vue';
import VueRouter from 'vue-router';

import OurFiles from '@/views/OurFiles';
import SharedWithUs from '@/views/SharedWithUs';
import SharedWithOthers from '@/views/SharedWithOthers';

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/our-files'
  }, {
    path: '/our-files',
    name: 'our-files',
    component: OurFiles
  }, {
    path: '/shared-with-us/:id(\\w+:\\w+:\\d+:.+)?',
    name: 'shared-with-us',
    component: SharedWithUs
  }, {
    path: '/shared-with-others/:id(\\w+:\\w+:\\d+:.+)?',
    name: 'shared-with-others',
    component: SharedWithOthers
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
