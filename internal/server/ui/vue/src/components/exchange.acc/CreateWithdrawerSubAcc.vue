<template>
  <v-container>
    <v-form validate-on="submit" ref="form">
      <v-row>
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
          <v-text-field
            ref="pass-phrase-input"
            :rules="okexPassPhraseRules()"
            @input="passPhraseChanged"
            v-model="okexPassPhrase"
            density="comfortable"
            variant="outlined"
            label="exchange pass phrase*"
            hint="exchange pass phrase*"/>
        </v-col>
        <v-col cols="12">
          <ProxyInput v-model="item.proxy" :required="true"/>
        </v-col>
      </v-row>
      <span v-if="error" class="text-red">{{ error }}</span>
      <v-btn
        type="submit"
        color="blue-darken-1"
        variant="text"
        :disabled="saveLoading"
        :loading="saveLoading"
        @click="CreateWithdrawer">
        Save
      </v-btn>
    </v-form>
  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {helperService, withdrawerService} from "@/generated/services"
import {CreateOkexWithdrawerRequest, CreateWithdrawerRequest, ExchangeType, Withdrawer} from "@/generated/withdraw";
import ProxyInput from "@/components/ProxyInput.vue";
import OkexWithdrawOptionSubAcc from "@/components/exchange.acc/OkexWithdrawOptionSubAcc.vue";

export default defineComponent({
  name: "CreateWithdrawerSubAcc",
  emits: ["withdrawerCreated"],
  computed: {
    ExchangeType() {
      return ExchangeType
    },
  },
  components: {
    OkexWithdrawOptionSubAcc,
    ProxyInput
  },
  props: {
    id: {
      required: true,
      type: String
    },
  },
  data() {
    return {
      error: null as any,
      item: {} as CreateOkexWithdrawerRequest,
      okexPassPhrase: "",
      saveLoading: false,
    }
  },
  methods: {
    okexPassPhraseRules() {
      return [
        (v: string) => this.okexPassPhrase ? true : 'pass phrase is required'
      ]
    },
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
    async apiSecretChanged(e: InputEvent) {
      await this.validateApiSecret()
    },
    async apiKeyChanged(e: InputEvent) {
      await this.validateApiKey()
    },
    async passPhraseChanged(e: InputEvent) {
      await this.validateOkexPassPhrase()
    },
    async validateOkexPassPhrase(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['pass-phrase-input'].validate()

      return valid
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
      this.error = ""
      this.saveLoading = true
      if (!(await this.validateForm())) {
        this.saveLoading = false
        return
      }

      try {

        let arg = Object.assign({}, this.item) as CreateOkexWithdrawerRequest;

        arg.apiKey = `${this.item.apiKey}:${this.okexPassPhrase}`
        arg.parentId = this.id
        await withdrawerService.withdrawerServiceCreateSubWithdrawer({body: arg})
        this.ResetForm()
        this.$emit('withdrawerCreated')
      } catch (e) {
        this.error = e
      } finally {
        this.saveLoading = false
      }
    },
    ResetForm() {
      this.item.parentId = ""
      this.item.apiKey = ""
      this.item.label = ""
      this.item.secretKey = ""
      this.item.proxy = ""
    },
  },
})


</script>


<style>


</style>

