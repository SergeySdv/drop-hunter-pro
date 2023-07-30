// Composables
import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'

const routes = [
  {
    name: 'Modules',
    path: '/modules',
    component: () => import('@/components/modules/General.vue'),
  },
  {
    name: 'About',
    path: '/about',
    component: () => import('@/components/About.vue'),
  },
  {
    name: 'LandingPage',
    path: '/',
    component: () => import('@/components/landing/LandingPage.vue'),
  },
  {
    name: 'GetStarted',
    path: '/get_started',
    component: () => import('@/components/guides/GetStarted.vue'),
  },
  {
    name: 'Profiles',
    path: '/profiles',
    component: () => import('@/components/profile/Profiles.vue'),
  },
  {
    name: 'Withdraw',
    path: '/exchange_accounts',
    component: () => import('@/components/exchange.acc/ExchangeAccounts.vue'),
  },
  {
    name: 'Constructor',
    path: '/constructor',
    component: () => import('@/components/flow/flows.vue'),
  },
  {
    name: 'Processes',
    path: '/processes',
    component: () => import('@/components/process/Processes.vue'),
  },
  {
    name: 'ViewFlow',
    path: '/flow/:id',
    component: () => import('@/components/flow/ViewFlow.vue'),
    props: {
      id: String,
    }
  },
  {
    name: 'ViewProcess',
    path: '/process/:id',
    component: () => import('@/components/process/ViewProcess.vue'),
    props: {
      id: String,
    }
  },
  {
    name: 'Settings',
    path: '/settings',
    component: () => import('@/components/settings/Settings.vue'),
    props: {
      id: String,
    }
  },
  {
    name: 'CreateFlow',
    path: '/flow/create',
    component: () => import('@/components/flow/CreateFlow.vue'),
  },
  {
    name: 'ExchangeAccountView',
    path: '/exchange_account',
    component: () => import('@/components/exchange.acc/ExchangeAccountView.vue'),
  },
  {
    name: 'ExchangeAccountBinanceDocs',
    path: '/docs/binance_account_instruction',
    component: () => import('@/components/exchange.acc/ExchangeAccountBinanceDocs.vue'),
  },
  {
    name: 'ExchangeAccountOkexDocs',
    path: '/docs/okex_account_instruction',
    component: () => import('@/components/exchange.acc/ExchangeAccountOkexDocs.vue'),
  },
  {
    name: 'CreateProfileBatch',
    path: '/profiles/batch/create',
    component: () => import('@/components/profile/CreateProfile.vue'),
  },
  {
    name: 'Billing',
    path: '/billing',
    component: () => import('@/components/billing/Billing.vue'),
  },


] as RouteRecordRaw[]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
