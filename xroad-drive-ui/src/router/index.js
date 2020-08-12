import Vue from 'vue'
import VueRouter from 'vue-router'

import OurFiles from '../views/OurFiles.vue'
import SharedWithUs from '../views/SharedWithUs.vue'
import SharedWithOthers from '../views/SharedWithOthers.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    redirect: '/our-files'
  }, {
    path: '/our-files',
    name: 'OurFiles',
    component: OurFiles
  }, {
    path: '/providers',
    name: 'SharedWithUs',
    component: SharedWithUs
  }, {
    path: '/clients',
    name: 'SharedWithOthers',
    component: SharedWithOthers
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
