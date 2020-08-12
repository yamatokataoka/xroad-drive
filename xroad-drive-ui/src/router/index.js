import Vue from 'vue'
import VueRouter from 'vue-router'

import MyFiles from '../views/MyFiles.vue'
import SharedWithUs from '../views/SharedWithUs.vue'
import SharedWithOthers from '../views/SharedWithOthers.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    redirect: '/my-files'
  }, {
    path: '/my-files',
    name: 'MyFiles',
    component: MyFiles
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
