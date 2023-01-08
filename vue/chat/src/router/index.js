import Vue from 'vue'
import VueRouter from 'vue-router'
import AuthView from '@/views/Auth'
import MainView from '@/views/Main'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    component: AuthView
  },
  {
    path: '/home',
    component: MainView
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
