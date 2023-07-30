<template>
  <v-text-field
    :type="show1 ? 'text' : 'password'"
    ref="proxy-input"
    label="SOCKS5 proxy in format <ip:port:login:password>"
    :messages="proxyStat"
    :rules="proxyRules()"
    @input="proxyChanged"
    v-model="proxy"
    density="comfortable"
    variant="outlined"
    :loading="proxyLoading"
    :disabled="proxyLoading"
    :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
    @click:append="show1 = !show1"
  ></v-text-field>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {helperService} from "@/generated/services"
import {Timer} from "@/components/helper";

export default defineComponent({
  name: "ProxyInput",
  props: {
    required: {
      type: Boolean,
      required: true,
    },
    modelValue: {
      type: String,
      required: false
    }
  },
  emits: ['update:modelValue'],
  computed: {
    proxy: {
      get(): string {
        return this.modelValue ? this.modelValue : ""
      },
      async set(value: string) {
        await this.proxyChanged()
        this.$emit('update:modelValue', value)
      }
    }
  },
  data() {
    return {
      show1: false,
      timer: new Timer(),
      proxyStat: "",
      proxyLoading: false,
    }
  },
  methods: {
    proxyRules() {
      return [
        async (v: string) => {
          if (this.required) {
            if (!this.proxy) {
              return 'Proxy required'
            }
          }
          if (!this.proxy) {
            return true
          }
          try {
            const res = await helperService.helperServiceValidateProxy({body: {proxy: this.proxy}})
            if (res.valid) {
              this.proxyStat = `${res.ip} ${res.countryName} [${res.countryCode}]`
              return true
            } else {
              return res.errorMessage
            }
          } catch (e) {
            return 'invalid proxy'
          }
        }
      ]
    },
    async validateProxy(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['proxy-input'].validate()

      return valid
    },
    async proxyChanged() {

      this.timer.add(1000)
      this.timer.cb(async () => {
        this.proxyLoading = true
        if (!await this.validateProxy()) {
          this.proxyLoading = false
          return
        }
        this.proxyLoading = false
      })
    },
  }
})


</script>


<style>


</style>

