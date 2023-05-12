import axios from 'axios'
import {notify} from '@kyvg/vue3-notification'
import dayjs from "dayjs";

import {ProfileService,} from "@/generated/profile"
import {HelperService} from "@/generated/helper";
import {WithdrawerService} from "@/generated/withdraw";
import {FlowService} from "@/generated/flow";
import {ProcessService} from "@/generated/process";
import {SettingsService} from "@/generated/settings";

import {serviceOptions as rso} from './profile'
import {serviceOptions as so1} from './helper'
import {serviceOptions as so2} from './withdraw'
import {serviceOptions as so3} from './flow'
import {serviceOptions as so4} from './process'
import {serviceOptions as so5} from './settings'


const instance = axios.create({
  baseURL: import.meta.env.DEV ? 'http://localhost:8083/' : '/',
  timeout: 30000,
  withCredentials: true,
});

instance.interceptors.request.use(res => {
  if (res) {
    if (res.headers) {
      // @ts-ignore
      res.headers['Grpc-Metadata-tz'] = dayjs().format('Z')
    }
  }


  return res
}, err => {
  console.log(err)
  console.log(err.response.status)
  if (err.response.status === 401) {
    notify({text: "not authenticated", type: 'error'})
  }
})

export const login = () => {
  window.location.href = import.meta.env.DEV ? "http://localhost:8083/api/google/login" : '/api/google/login'
}

rso.axios = instance
so1.axios = instance
so2.axios = instance
so3.axios = instance
so4.axios = instance
so5.axios = instance


export const profileService = new ProfileService()
export const helperService = new HelperService()
export const withdrawerService = new WithdrawerService()
export const flowService = new FlowService()
export const processService = new ProcessService()
export const settingsService = new SettingsService()





