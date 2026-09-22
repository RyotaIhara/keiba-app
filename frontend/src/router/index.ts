import { createRouter, createWebHistory } from 'vue-router'

import IndexView from '@/views/IndexView.vue'
import RaceCourseView from '@/views/RaceCourseView.vue'
import UserView from '@/views/UserView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'route',
      component: IndexView,
    },
    {
      path: '/user',
      name: 'user',
      component: UserView,
    },
    {
      path: '/race-course',
      name: 'race-course',
      component: RaceCourseView,
    },
  ],
})

export default router
