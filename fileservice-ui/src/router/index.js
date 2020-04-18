import Vue from 'vue'
import VueRouter from 'vue-router'

import MyFiles from '../views/MyFiles.vue'
import SharedWithUs from '../views/SharedWithUs.vue'
import SharedWithCompanies from '../views/SharedWithCompanies.vue'

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
    path: '/shared-with-us',
    name: 'SharedWithUs',
    component: SharedWithUs
  }, {
    path: '/shared-with-companies',
    name: 'SharedWithCompanies',
    component: SharedWithCompanies
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
