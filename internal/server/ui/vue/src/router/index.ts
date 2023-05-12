// Composables
import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'

const routes = [
  {
    name: 'Profiles',
    path: '/profiles',
    component: () => import('@/components/Profiles.vue'),
  },
  {
    name: 'Profiles',
    path: '/',
    component: () => import('@/components/Profiles.vue'),
  },
  {
    name: 'Withdraw',
    path: '/withdraw',
    component: () => import('@/components/Withdraw.vue'),
  },
  {
    name: 'Constructor',
    path: '/constructor',
    component: () => import('@/components/Constructor.vue'),
  },
  {
    name: 'Processes',
    path: '/processes',
    component: () => import('@/components/Processes.vue'),
  },
  {
    name: 'ViewFlow',
    path: '/flow/:id',
    component: () => import('@/components/ViewFlow.vue'),
    props: {
      id: String,
    }
  },
  {
    name: 'ViewProcess',
    path: '/process/:id',
    component: () => import('@/components/ViewProcess.vue'),
    props: {
      id: String,
    }
  },
  {
    name: 'Settings',
    path: '/settings',
    component: () => import('@/components/Settings.vue'),
    props: {
      id: String,
    }
  },
] as RouteRecordRaw[]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
