import { createRouter, createWebHistory } from 'vue-router'

import DashboardView from '@/views/DashboardView.vue'
import RaceCourseView from '@/views/RaceCourseView.vue'
import RaceView from '@/views/RaceView.vue'
import UserView from '@/views/UserView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: DashboardView,
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
    {
      path: '/race',
      name: 'race',
      component: RaceView,
    },
  ],
})

export default router
