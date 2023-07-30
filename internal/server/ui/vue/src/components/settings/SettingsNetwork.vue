<template>
  <v-container>
    <v-card>
      {{ network }}
      <v-card-text>
        <v-row>
          <v-col>
            <v-text-field
              ref="network-input"
              v-model="settings.rpcEndpoint"
              label="rpc endpoint"
              density="compact"
              variant="outlined"
              :rules="networkRules()"
              @input="endpointChanged(false)"
              :loading="loading"
              :disabled="loading"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-text-field
              v-model="settings.id"
              label="network id"
              density="compact"
              variant="outlined"
              :loading="loading"
              :disabled="true"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <NetworkGasMultiplier v-model="settings.gasMultiplier"/>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <NetworkMaxGas v-model="settings.maxGas" :network="network"/>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {settingsService} from "@/generated/services"
import {Timer} from "@/components/helper";
import {SettingsNetwork} from "@/generated/settings";
import {Network} from "@/generated/flow";
import NetworkGasMultiplier from "@/components/settings/NetworkGasMultiplier.vue";
import deepEqual from "deep-equal";
import NetworkMaxGas from "@/components/settings/NetworkMaxGas.vue";

export default defineComponent({
  name: "SettingsNetwork",
  components: {NetworkMaxGas, NetworkGasMultiplier},
  props: {
    modelValue: {
      type: Object as PropType<SettingsNetwork>,
      required: true,
    },
    network: {
      type: Object as PropType<Network>,
      required: true,
    }
  },
  emits: ['update:modelValue'],
  watch: {
    settings: {
      handler() {
        if (!deepEqual(this.settings, this.modelValue)) {

          this.$emit('update:modelValue', this.settings)
        }

      },
      deep: true,
    }
  },
  data() {
    return {
      timer: new Timer(),
      loading: false,
      suggestedGasPrice: 0,
      ETHPrice: 2000,
      settings: {} as SettingsNetwork,
    }
  },
  computed: {
    Network() {
      return Network
    },
    limitGasUSD() {
      return this.ETHPrice / 10e17 * Number(this.settings.gasTotal)
    }
  },
  methods: {
    async useLimitGasChanged() {

    },
    async endpointChanged(init: boolean) {
      if (this.network == Network.ZKSYNCLITE) {
        return
      }
      if (!init) {

        this.timer.add(1000)
        this.timer.cb(async () => {
          this.loading = true
          if (!await this.validate()) {
            this.loading = false
            return
          }

          this.loading = false
        })
      }
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['network-input'].validate()

      return valid
    },
    networkRules() {
      return [
        async (v: string) => {
          if (this.network == Network.ZKSYNCLITE) {
            return true
          }
          if (!this.settings?.rpcEndpoint) {
            return 'rpc required'
          }

          try {
            const res = await settingsService.settingsServiceGetNetworkByRpc({
              body: {
                network: this.network,
                endpoint: this.settings.rpcEndpoint
              }
            })

            if (res.valid) {
              this.settings.id = res.id
              this.suggestedGasPrice = Number(res.suggestedGasPrice)
              return true
            } else {
              this.settings.id = ''
              this.suggestedGasPrice = 0
              return res.error ? res.error : 'call an ambulance. api error'
            }
          } catch (e) {
            return 'call an ambulance. api error'
          }
        }
      ]
    },
  },
  async created() {
    this.settings = Object.assign({}, this.modelValue)

    const res = await settingsService.settingsServiceGetNetworkByRpc({
      body: {
        network: this.network,
        endpoint: this.settings.rpcEndpoint
      }
    })

    if (res.valid) {
      this.suggestedGasPrice = Number(res.suggestedGasPrice)
    }
  }
})


</script>


<style>


</style>

