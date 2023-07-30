<template>
  <v-dialog
    v-model="dialog"
    :close-on-content-click="false"
    :close-on-back="false"
    persistent="true"
    width="400px"
    location="center"
  >
    <template v-slot:activator="{ props }">
      <v-btn variant="outlined" @click="dialog = true" density="comfortable" v-bind="props">TopUp</v-btn>
    </template>

    <v-skeleton-loader
      width="100%"
      :loading="loading"
      type="table"
    >
      <v-responsive>

        <div class="my-5 mx-5">
          <div class="my-5 d-flex justify-space-between">
            <span class="text-h5">TopUp</span>
            <v-icon icon="mdi-close" @click="reset()"/>
          </div>
          <div>
            <v-form ref="form">
              <v-row>
                <WithdrawerSelect v-model="withdrawerId"/>
              </v-row>
              <div v-if="withdrawerId">
                <v-row align="center">
                  <v-autocomplete
                    ref="exchange-withdraw-form-token"
                    density="comfortable"
                    variant="outlined"
                    label="select token"
                    :rules="[required]"
                    :items="tokens"
                    item-title="title"
                    item-value="token"
                    v-model="token"
                    :loading="optionsLoading"
                    :disabled="disableWithdraw"
                  />
                </v-row>
                <v-row>
                  <v-select
                    ref="exchange-withdraw-form-network"
                    density="comfortable"
                    variant="outlined"
                    label="network"
                    item-value="network"
                    item-title="title"
                    :rules="[required]"
                    :items="GetNetworks"
                    v-model="network"
                    :loading="optionsLoading"
                    :disabled="disableWithdraw"
                  />
                </v-row>
                <v-row>
                  <v-text-field
                    type="number"
                    label="amount"
                    density="comfortable"
                    variant="outlined"
                    :rules="[required]"
                    :loading="optionsLoading"
                    :disabled="disableWithdraw"
                    v-model="amount"/>
                </v-row>
                <v-row>
                  <v-btn width="100%" @click="withdraw" :loading="withdrawLoading" :disabled="disableWithdraw">TopUp
                  </v-btn>
                </v-row>
                <v-row class="d-flex justify-center">
                  <div style="color: red" v-if="withdrawError">{{ withdrawError }}</div>
                  <div class="text-center text-h6" v-if="withdrawId">Coins sent!</div>
                </v-row>
              </div>
            </v-form>
          </div>
        </div>
      </v-responsive>
    </v-skeleton-loader>
  </v-dialog>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {withdrawerService} from "@/generated/services"
import {
  CreateOkexWithdrawerRequest,
  ExchangeType, ExchangeWithdrawNetwork,
  ExchangeWithdrawOptions,
  Withdrawer
} from "@/generated/withdraw";
import ProxyInput from "@/components/ProxyInput.vue";
import OkexWithdrawOptionSubAcc from "@/components/exchange.acc/OkexWithdrawOptionSubAcc.vue";
import CreateWithdrawerSubAcc from "@/components/exchange.acc/CreateWithdrawerSubAcc.vue";
import {required} from "@/components/tasks/menu/helper";
import ProfileSearch from "@/components/profile/ProfileSearch.vue";
import {Profile} from "@/generated/profile";
import {Delay, Timer} from "@/components/helper";
import WithdrawerSelect from "@/components/exchange.acc/WithdrawerSelect.vue";

interface Network {
  network: string
  title: string
}

interface Token {
  token: string
  title: string
}

export default defineComponent({
  props: {
    profileId: {
      required: true,
      type: String,
    },
  },
  name: "TopUpProfile",
  watch: {
    'withdrawerId': {
      handler(b, a) {
        this.withdrawerChanged()
      },
    },
    'token': {
      handler(b, a) {
        this.tokenChanged()
      },
    },
    'network': {
      handler(b, a) {
      },
    },
  },
  computed: {
    ExchangeType() {
      return ExchangeType
    },
    disableWithdraw(): boolean {
      return !!this.withdrawId;
    },
    GetNetworks(): Network[] {
      return this.networks
    },
  },
  components: {
    WithdrawerSelect,
    ProfileSearch,
    CreateWithdrawerSubAcc,
    OkexWithdrawOptionSubAcc,
    ProxyInput
  },
  data() {
    return {
      token: '',
      withdrawerId: '',
      network: '',
      dialog: false,
      withdrawId: '',
      timer: null as NodeJS.Timer | null,
      withdrawStatus: '',
      withdrawLoading: false,
      optionsLoading: false,
      networks: [] as Network[],
      tokens: [] as Token[],
      loading: false,
      options: new Map<string, ExchangeWithdrawNetwork[]>(),
      withdrawError: '',
      amount: '',
    }
  },
  methods: {
    reset() {
      this.token = ''
      this.withdrawerId = ''
      this.network = ''
      this.withdrawId = ''
      this.withdrawStatus = ''
      this.optionsLoading = false
      this.withdrawError = ''
      this.amount = ''
      this.dialog = false
    },
    startPolling() {
      this.timer = setInterval(() => {
        this.getWithdrawStatus()
      }, 10000)
    },
    stopPolling() {
      if (this.timer) {
        clearInterval(this.timer)
      }
    },
    async getWithdrawStatus() {
      if (!this.withdrawId) {
        return ''
      }

      try {
        const res = await withdrawerService.withdrawerServiceWithdrawStatus({body: {withdrawId: this.withdrawId}})
        this.withdrawStatus = res.status
        if (this.withdrawStatus === 'done' || this.withdrawStatus === 'error') {
          this.stopPolling()
        }
      } finally {

      }
    },
    async withdraw() {
      try {
        this.withdrawLoading = true
        const valid = await this.validateWithdrawForm()

        if (!valid) {
          this.withdrawLoading = false
          return
        }
        this.withdrawError = ''
        this.withdrawId = ''
        const res = await withdrawerService.withdrawerServiceWithdraw({
          body: {
            amount: this.amount,
            network: this.network,
            profileId: this.profileId,
            token: this.token,
            withdrawerId: this.withdrawerId,
          }
        })
        if (res.errorMessage) {
          this.withdrawError = res.errorMessage
        } else {
          this.withdrawId = res.withdrawId
        }
      } finally {
        this.withdrawLoading = false
      }
    },
    required,
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['formm'].validate()

      return valid
    },
    async validateWithdrawForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async getWithdrawerOptions() {
      try {
        this.tokens = []
        this.networks = []
        this.optionsLoading = true
        const res = await withdrawerService.withdrawerServiceGetExchangeWithdrawOptions({body: {withdrawerId: this.withdrawerId}})

        res.tokens.sort((a, b) => {
          return Number(b.amount) - Number(a.amount)
        })

        res.tokens.forEach(token => {
          const t: Token = {
            token: token.token,
            title: this.GetTokenTitle(token)
          }
          this.tokens.push(t)
          this.options.set(token.token, token.networks)
        })


        this.tokenChanged()
      } finally {
        this.optionsLoading = false
      }
    },
    tokenChanged() {
      if (this.token) {
        const networks = this.options.get(this.token)
        if (networks) {
          this.networks = []
          networks.forEach((n) => {
            this.networks.push({
              network: n.network,
              title: this.GetNetworkTitle(n)
            })
          })
        }
      }
    },
    async withdrawerChanged() {
      if (this.withdrawerId) {
        try {
          this.tokens = []
          this.networks = []
          this.optionsLoading = true
          const res = await withdrawerService.withdrawerServiceGetExchangeWithdrawOptions({body: {withdrawerId: this.withdrawerId}})

          res.tokens.sort((a, b) => {
            return Number(b.amount) - Number(a.amount)
          })

          res.tokens.forEach(token => {
            const t: Token = {
              token: token.token,
              title: this.GetTokenTitle(token)
            }
            this.tokens.push(t)
            this.options.set(token.token, token.networks)
          })


          this.tokenChanged()
        } finally {
          this.optionsLoading = false
        }
      }
    },
    GetTokenTitle(item: ExchangeWithdrawOptions): string {
      if (item) {
        return item.token + " - " + item.amount
      }
      return ""
    },
    GetNetworkTitle(item: ExchangeWithdrawNetwork): string {
      if (item) {
        return `${item.network} min:${item.minAmount} max:${item.maxAmount} fee:${item.fee}`
      }
      return ""
    },
  },
  async created() {
  }
})


</script>


<style>


</style>

