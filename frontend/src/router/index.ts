import { createRouter, createWebHistory } from 'vue-router'

import HelloView from '@/views/HelloView.vue'
import IndexView from '@/views/IndexView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'route',
      component: IndexView,
    },
    {
      path: '/hello',
      name: 'hello',
      component: HelloView,
    },
  ],
})

export default router
