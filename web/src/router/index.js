import { createRouter, createWebHistory } from 'vue-router'
import AboutView from '@/views/AboutView.vue'
import IndexView from '@/views/IndexView.vue'
import LoginView from '@/views/LoginView.vue'
import UserView from '@/views/UserView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', component: IndexView },
    { path: '/index', component: IndexView },
    { path: '/user', component: UserView },
    { path: '/login', component: LoginView },
  ]
})

export default router
