<template>
  <v-container>
    <v-card>
      {{ network }}
      <v-card-text>
        <v-row>
          <v-col>
            <v-text-field
              ref="network-input"
              v-model="params.rpcEndpoint"
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
              v-model="params.id"
              label="network id"
              density="compact"
              variant="outlined"
              :loading="loading"
              :disabled="true"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-checkbox
          density="compact"
          focused
          v-model="params.useLimitGas"
          label="use custom limits for gas"
          @input="useLimitGasChanged">
        </v-checkbox>

        <div v-if="params.useLimitGas">
          <v-row>
            <v-col>
              <v-text-field
                v-model="params.gasTotal"
                label="max gas used for tx (wei)"
                density="compact"
                variant="outlined"
              ></v-text-field>
            </v-col>
            <v-col>
              <v-text-field
                v-model="suggestedGasPrice"
                label="network gas price (wei)"
                density="compact"
                variant="outlined"
                :loading="loading"
                :disabled="true"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            Gas in USD: {{ limitGasUSD }}
          </v-row>
        </div>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {settingsService} from "@/generated/services"
import {Delay} from "@/components/helper";
import {SettingsNetwork} from "@/generated/settings";
import {Network} from "@/generated/flow";

export default defineComponent({
  name: "SettingsNetwork",
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
  data() {
    return {
      loading: false,
      suggestedGasPrice: 0,
      ETHPrice: 2000,
    }
  },
  computed: {
    params: {
      get(): SettingsNetwork {
        return this.modelValue
      },
      async set(value: SettingsNetwork) {
        this.$emit('update:modelValue', value)
      }
    },
    limitGasUSD() {
      return this.ETHPrice / 10e18 * Number(this.params.gasTotal)

    }
  },
  methods: {
    async useLimitGasChanged() {

    },
    async endpointChanged(init: boolean) {
      console.log('endpointChanged', init)
      if (!init) {
        const before = this.params.rpcEndpoint
        await Delay(1000)
        if (this.params.rpcEndpoint !== before) {
          return
        }
      }

      this.loading = true
      if (!await this.validate()) {
        this.loading = false
        return
      }

      this.loading = false
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
          if (!this.params?.rpcEndpoint) {
            return 'rpc required'
          }

          try {
            const res = await settingsService.settingsServiceGetNetworkByRpc({
              body: {
                network: this.network,
                endpoint: this.params.rpcEndpoint
              }
            })

            if (res.valid) {
              this.params.id = res.id
              this.suggestedGasPrice = Number(res.suggestedGasPrice)
              return true
            } else {
              this.params.id = ''
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
    const res = await settingsService.settingsServiceGetNetworkByRpc({
      body: {
        network: this.network,
        endpoint: this.params.rpcEndpoint
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

