import {createPinia, mapStores} from 'pinia'
import router from "@/router";

export const pinia = createPinia()


import {defineStore} from 'pinia'
import {helperService} from "@/generated/services";
import {TaskType} from "@/generated/flow";
import {Delay} from "@/components/helper";

export const useUserStore = defineStore('user', {
  state: () => ({
    login: false,
    email: '',
    ass: '0 USD',
    taskPrice: '0 USD',
    nonpayableTasks: [] as TaskType[],
    payableTasks: [] as TaskType[],
    impact: {
      my: "0",
      top: "0",
      total: "0"
    }
  }),
  actions: {
    async redirectToGeneralPage() {
      const path = window.location.pathname
      if (path === "/" || path === '/modules') {
        return
      }

      if (window.location.hostname === 'localhost') {
        return
      }

      await router.push({name: 'LandingPage'})
    },
    async syncUser(): Promise<void> {
      try {
        const res = await helperService.helperServiceGetUser()
        this.$patch((state) => {
          state.login = true
          state.email = res.email
          state.ass = Number(res.funds).toPrecision(6) + " USD"
          state.taskPrice = Number(res.taskPrice).toPrecision(2) + " USD"
          state.payableTasks = res.payableTasks
          state.nonpayableTasks = res.nonpayableTasks
        })
      } catch (e) {
        this.$patch((state) => {
          state.login = false
          state.email = ''
          state.ass = '0 USD'
          state.taskPrice = '0 USD'
        })


        await this.redirectToGeneralPage()
      }
    },
    async syncDailyImpact(): Promise<void> {
      try {
        const res = await helperService.helperServiceTransactionsDailyImpact()
        this.$patch((state) => {
          state.impact.my = res.myImpact
          state.impact.top = res.topImpact
          state.impact.total = res.totalImpact
        })
      } catch (e) {
        this.$patch((state) => {
          state.impact.my = '0'
          state.impact.top = '0'
          state.impact.total = '0'
        })
      }
    }
  }
})

export const useSysStore = defineStore('sys', {
  state: () => ({
    snackbar: false,
    snackbarText: '',
    error: false
  }),
  actions: {
    showSnackBar(text: string, error: boolean) {
      this.$patch((state) => {
        state.snackbar = true
        console.log('state.snackbartext', text)
        state.snackbarText = text
        state.error = error
      })
      Delay(5000, () => {
        this.$patch((state) => {
          state.snackbar = false
          state.snackbarText = ''
          state.error = error
        })
      })

    }
  }
})
