//@ts-nocheck
// @formatter:off
import axios from 'axios'
import dayjs from "dayjs";
import {mapActions,} from "pinia";
import {useUserStore} from "@/plugins/pinia";

// @ts-ignore
import CryptoJS, {Hex} from "crypto-js";

import {serviceOptions as rso, ProfileService} from './profile'
import {serviceOptions as so1, HelperService} from './helper'
import {serviceOptions as so2, WithdrawerService} from './withdraw'
import {serviceOptions as so3, FlowService} from './flow'
import {serviceOptions as so4, ProcessService} from './process'
import {serviceOptions as so5, SettingsService} from './settings'
import {serviceOptions as so6, Swap1InchService} from './swap1inch'
import {serviceOptions as so7, OrbiterService} from './orbiter'


let ivs = "1234567812345678"

let seceret = ''
if (import.meta.env.DEV) {
  seceret = 'jd92mdld93mdlemr'
  ivs = 'jd42mdlf93mdlemr'
} else {
  seceret = ['ZEZqACAp3TGYNQHe'][0]
  ivs = ['Fk2J79xYcZLhhZpM'][0]
}

export const instance = axios.create({
  baseURL: import.meta.env.DEV ? 'http://localhost:8083/' : '/',
  timeout: 30000,
  withCredentials: true,
  transformRequest: [
    (data, headers) => {

      if (!data) {
        data = {}
      }
      headers['Grpc-Metadata-tz'] = dayjs().format('Z')
      headers.setContentType('application/base64', true)
      headers.setAccept('application/base64', true)

      console.log("request body: ", data)
      return encrypt(data, seceret);
    },
  ],
  transformResponse: [
    (data, headers) => {
      if (data) {
        data = decrypt(data, seceret);
      }
      console.log("response body: ", data)
      return data
    }
  ]
});


const store = mapActions(useUserStore, ['redirectToGeneralPage'])


const encrypt = (str: any, key: string) => {
  let b = JSON.stringify(str)
  //@ts-ignore
  key = CryptoJS.enc.Utf8.parse(key);
  let iv = CryptoJS.enc.Utf8.parse(ivs);
  let encrypted = CryptoJS.AES.encrypt(b, key, {
    iv: iv,
    padding: CryptoJS.pad.Pkcs7
  });
  return encrypted.toString();
}
const decrypt = (str: any, key: string) => {
  //@ts-ignore
  key = CryptoJS.enc.Utf8.parse(key);
  let iv = CryptoJS.enc.Utf8.parse(ivs);
  let encrypted = CryptoJS.AES.decrypt(str.toString(), key, {
    iv: iv,
    padding: CryptoJS.pad.Pkcs7
  });
  const b = encrypted.toString(CryptoJS.enc.Utf8);

  return JSON.parse(b)
}

instance.interceptors.response.use(res => {
  return res
}, err => {
  console.error(err)

  if (err.response?.status === 401) {
    store.redirectToGeneralPage()
  } else if (err.response?.status === 403) {

  }

  throw new Error(err.response?.data?.message)
})

export const loginGoogleHref = () => {
  return import.meta.env.DEV ? "http://localhost:8083/api/google/login" : '/api/google/login'
}

rso.axios = instance
so1.axios = instance
so2.axios = instance
so3.axios = instance
so4.axios = instance
so5.axios = instance
so6.axios = instance
so7.axios = instance


export const profileService = new ProfileService()
export const helperService = new HelperService()
export const withdrawerService = new WithdrawerService()
export const flowService = new FlowService()
export const processService = new ProcessService()
export const settingsService = new SettingsService()
export const swap1inchService = new Swap1InchService()
export const orbiterService = new OrbiterService()





