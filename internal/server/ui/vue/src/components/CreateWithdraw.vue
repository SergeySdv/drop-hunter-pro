<template>
  <v-row justify="center">
    <v-dialog
      v-model="show"
      persistent
      width="1024"
    >
      <v-card>
        <v-card-title>
          <span class="text-h5">Create withdraw account</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-form validate-on="submit" ref="form">
              <v-row>
                <v-col cols="12" sm="6" md="4">
                  <v-select
                    :items="exchanges"
                    v-model="exchange"
                    label="Exchange"
                    variant="outlined"
                    density="comfortable"
                  ></v-select>
                </v-col>
                <v-col cols="12" sm="6" md="8">
                  <v-text-field
                    ref="label-input"
                    v-model="item.label"
                    label="label"
                    density="comfortable"
                    variant="outlined"
                    :rules="labelRules()"
                    @input="labelChanged"
                  ></v-text-field>
                </v-col>
                <v-col cols="12" sm="12" md="12">
                  <v-text-field
                    ref="secret-input"
                    :rules="secretKeyRules()"
                    @input="apiSecretChanged"
                    v-model="item.secretKey"
                    density="comfortable"
                    variant="outlined"
                    label="exchange api secret key*"
                    hint="exchange api secret key"/>
                </v-col>
                <v-col cols="12" sm="12" md="12">
                  <v-text-field
                    ref="api-key-input"
                    :rules="apiKeyRules()"
                    @input="apiKeyChanged"
                    v-model="item.apiKey"
                    density="comfortable"
                    variant="outlined"
                    label="exchange api key*"
                    hint="exchange api key"/>
                </v-col>
                <v-col cols="12">
                  <ProxyInput @changed="proxyChanged" :required="true"/>
                </v-col>
              </v-row>
            </v-form>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue-darken-1"
            variant="text"
            @click="Close"
          >
            Close
          </v-btn>
          <v-btn
            type="submit"
            color="blue-darken-1"
            variant="text"
            :disabled="saveLoading"
            :loading="saveLoading"
            @click="CreateWithdrawer">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {helperService, withdrawerService} from "@/generated/services"
import {CreateWithdrawerRequest, ExchangeType} from "@/generated/withdraw";
import ProxyInput from "@/components/ProxyInput.vue";

export default defineComponent({
  name: "CreateWithdraw",
  components: {
    ProxyInput
  },
  props: {
    showProp: Boolean,
  },
  watch: {
    showProp: {
      immediate: true,
      handler(newValue, oldValue) {
        this.show = newValue
      },
    }
  },
  data() {
    return {
      exchange: ExchangeType.Binance,
      exchanges: [
        ExchangeType.Binance,
      ],
      show: this.showProp,
      item: {} as CreateWithdrawerRequest,
      proxyRules: [
        async (v: string) => {
          if (!v) {
            return true
          }
          const res = await helperService.helperServiceValidateProxy({body: {proxy: v}})
          if (res.valid) {
            return true
          } else {
            return res.errorMessage
          }
        }
      ],
      mmskWalletId: "",
      proxyStat: "",
      proxyLoading: false,
      proxyErrorMessage: "",
      proxyTimer: 0,
      saveLoading: false,
    }
  },
  methods: {
    labelRules() {
      return [
        (v: string) => this.item.label ? true : 'label is required'
      ]
    },
    apiKeyRules() {
      return [
        (v: string) => {
          if (this.item.apiKey) {
            return true
          }
          return 'api key is required'
        },
      ]
    },
    secretKeyRules() {
      return [
        (v: string) => {
          if (this.item.secretKey) {
            return true
          }
          return 'api secret key is required'
        },
      ]
    },
    async labelChanged(e: InputEvent) {
      await this.validate()
    },
    async proxyChanged(proxy: string) {
      this.item.proxy = proxy
    },
    async apiSecretChanged(e: InputEvent) {
      await this.validateApiSecret()
    },
    async apiKeyChanged(e: InputEvent) {
      await this.validateApiKey()
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['label-input'].validate()

      return valid
    },
    async validateApiKey(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['api-key-input'].validate()

      return valid
    },
    async validateApiSecret(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['secret-input'].validate()

      return valid
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"].form.validate()

      return valid
    },
    async CreateWithdrawer() {
      if (!(await this.validateForm())) return

      try {
        this.saveLoading = true
        await withdrawerService.withdrawerServiceCreateWithdrawer({body: this.item})
        this.$emit("change", false)
        this.$emit("profileCreated")
      } catch (e) {
        this.saveLoading = false
      }

    },
    Close() {
      this.$emit("change", false)
    }
  }
})


</script>


<style>


</style>

