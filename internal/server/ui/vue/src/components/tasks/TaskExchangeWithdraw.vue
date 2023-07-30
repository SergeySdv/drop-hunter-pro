<template>
  <v-card-actions class="mt-3">
    <v-container>
      <v-row>
        <WithdrawerSelect v-model="item.withdrawerId" :disabled="disabled"/>
      </v-row>
      <v-row v-if="item.withdrawerId">
        <v-autocomplete
          ref="exchange-withdraw-form-token"
          density="compact"
          variant="outlined"
          label="select token"
          :rules="tokenRules()"
          :items="tokens"
          item-title="title"
          item-value="token"
          v-model="item.token"
          :disabled="disabled"
          :loading="optionsLoading"
        />
      </v-row>
      <v-row>
        <v-select
          ref="exchange-withdraw-form-network"
          density="compact"
          variant="outlined"
          label="network"
          item-value="network"
          item-title="title"
          :rules="networkRules()"
          :items="GetNetworks"
          v-model="item.network"
          :disabled="disabled"
          :loading="optionsLoading"
        />
      </v-row>
      <v-row>
        <v-checkbox
          :disabled="disabled"
          density="compact"
          focused
          name="send all coins on balance"
          v-model="item.sendAllCoins"
          :label="'send all coins from balance'"
          @input="CheckboxChanged">
        </v-checkbox>
      </v-row>
      <v-row v-if="!item.sendAllCoins">
        <v-col>
          <v-text-field
            ref="exchange-withdraw-form-min"
            type="number"
            label="min amount"
            density="compact"
            variant="outlined"
            :rules="getMinRules()"
            :disabled="disabled"
            v-model="item.amountMin"/>
        </v-col>
        <v-col>
          <v-text-field
            ref="exchange-withdraw-form-max"
            type="number"
            label="max amount"
            density="compact"
            variant="outlined"
            :rules="getMaxRules()"
            :disabled="disabled"
            v-model="item.amountMax"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-checkbox
          :disabled="disabled"
          density="compact"
          focused
          name="use external addr"
          v-model="item.useExternalAddr"
          :label="'use external address'"
          @input="CheckboxChanged">
        </v-checkbox>
      </v-row>
      <v-row v-if="item.useExternalAddr">
        <v-col>
          <div class="mb-5">You can use your address to receive coins <a
            href="https://www.binance.com/en/my/wallet/account/main/deposit/crypto/">binance</a>
          </div>
          <v-text-field
            ref="withdraw-external-addr-input"
            density="compact"
            variant="outlined"
            label="address to receive coins"
            :rules="binanceAddrRules()"
            v-model="item.withdrawAddr"
            :disabled="disabled"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">

import WithdrawerSelect from "@/components/exchange.acc/WithdrawerSelect.vue";

interface Token {
  token: string
  title: string
}

interface Network {
  network: string
  title: string
}

import {Task, TaskType, WithdrawExchangeTask} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {withdrawerService} from "@/generated/services";
import {ExchangeWithdrawNetwork, ExchangeWithdrawOptions, Withdrawer} from "@/generated/withdraw";

export default defineComponent({
  name: "TaskExchangeWithdraw",
  components: {WithdrawerSelect},
  emits: ['taskChanged'],
  props: {
    weight: {
      type: Number,
      required: true,
    },
    disabled: {
      type: Boolean,
      required: true,
    },
    task: {
      type: Object as PropType<Task>,
      required: false,
      deep: true
    }
  },

  watch: {
    'item.withdrawerId': {
      handler(b, a) {
        this.withdrawerChanged()
        this.validateWithdrawId()
        this.$emit('taskChanged', this.getTask())
      },
    },
    'item.token': {
      handler(b, a) {
        this.tokenChanged()
        this.validateToken()
        this.$emit('taskChanged', this.getTask())
      },
    },
    'item.network': {
      handler(b, a) {
        this.validateNetwork()
        this.$emit('taskChanged', this.getTask())
      },
    },
    'item.amountMax': {
      handler(b, a) {
        this.validateMax()
        this.$emit('taskChanged', this.getTask())
      },
    },
    'item.amountMin': {
      handler(b, a) {
        this.validateMin()
        this.$emit('taskChanged', this.getTask())
      },
    },
    'item.withdrawAddr': {
      handler(b, a) {
        this.validateAddr()
        this.$emit('taskChanged', this.getTask())
      },
    },
    'item.sendAllCoins': {
      handler(b, a) {
        this.$emit('taskChanged', this.getTask())
      },
    },
    task: {
      handler(b, a) {
        if (JSON.stringify(a) !== JSON.stringify(b)) {
          this.syncTask()
        }
      },
      deep: true
    }
  },

  data() {
    return {
      optionsLoading: false,
      networks: [] as Network[],
      tokens: [] as Token[],
      withdrawers: [] as Withdrawer[],
      item: {
        token: "",
        amountMax: "",
        amountMin: "",
        withdrawerId: "",
        network: "",
        withdrawAddr: "",
        useExternalAddr: false,
        sendAllCoins: false,
      } as WithdrawExchangeTask,
      options: new Map<string, ExchangeWithdrawNetwork[]>(),
    }
  },
  methods: {
    CheckboxChanged() {

    },
    binanceAddrRules() {
      return [(v: any) => !!this.item.withdrawAddr || 'required',]
    },
    withdrawIdRules() {
      return [(v: any) => !!this.item.withdrawerId || 'required',]
    },
    tokenRules() {
      return [(v: any) => !!this.item.token || 'required',]
    },
    networkRules() {
      return [(v: any) => !!this.item.network || 'required',]
    },
    getMinRules() {
      return [
        (v: any) => !!this.item.amountMin || 'required',
        (v: any) => {
          if (!this.item.amountMax) {
            return true
          }
          if (Number(this.item.amountMin) > Number(this.item.amountMax)) {
            return 'min > max'
          }
          return true
        }
      ]
    },
    getMaxRules() {
      return [
        (v: any) => !!this.item.amountMin || 'required',
        (v: any) => {
          if (Number(this.item.amountMin) > Number(this.item.amountMax)) {
            return 'min > max'
          }
          return true
        }
      ]
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
    tokenChanged() {
      if (this.item.token) {
        const networks = this.options.get(this.item.token)
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
      if (this.item.withdrawerId) {
        try {
          this.tokens = []
          this.networks = []
          this.optionsLoading = true
          const res = await withdrawerService.withdrawerServiceGetExchangeWithdrawOptions({body: {withdrawerId: this.item.withdrawerId}})

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
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        withdrawExchangeTask: this.item,
        taskType: TaskType.WithdrawExchange,
        description: "",
      }
    },
    async validateWithdrawId(): Promise<boolean> {
      try {
        // @ts-ignore попизди мне еще что руки из жопы у меня ага
        // спасибо китайцам скажи лучше
        const {valid} = await this["$refs"]['exchange-withdraw-form-withdraw-id'].validate()

        return valid
      } catch (e) {
        return false
      }
    },
    async validateToken(): Promise<boolean> {
      try {
        // @ts-ignore попизди мне еще что руки из жопы у меня ага
        // спасибо китайцам скажи лучше
        const {valid} = await this["$refs"]['exchange-withdraw-form-token'].validate()

        return valid
      } catch (e) {
        return false
      }
    },
    async validateNetwork(): Promise<boolean> {
      try {
        // @ts-ignore попизди мне еще что руки из жопы у меня ага
        // спасибо китайцам скажи лучше
        const {valid} = await this["$refs"]['exchange-withdraw-form-network'].validate()

        return valid
      } catch (e) {
        return false
      }
    },
    async validateAddr(): Promise<boolean> {
      try {
        // @ts-ignore попизди мне еще что руки из жопы у меня ага
        // спасибо китайцам скажи лучше
        const {valid} = await this["$refs"]['withdraw-external-addr-input'].validate()

        return valid
      } catch (e) {
        return false
      }
    },
    async validateMin(): Promise<boolean> {
      try {
        // @ts-ignore попизди мне еще что руки из жопы у меня ага
        // спасибо китайцам скажи лучше
        const {valid} = await this["$refs"]['exchange-withdraw-form-min'].validate()

        return valid
      } catch (e) {
        return false
      }
    },
    async validateMax(): Promise<boolean> {
      try {
        // @ts-ignore попизди мне еще что руки из жопы у меня ага
        // спасибо китайцам скажи лучше
        const {valid} = await this["$refs"]['exchange-withdraw-form-max'].validate()

        return valid
      } catch (e) {
        return false
      }

    },
    syncTask() {
      if (this.task && this.task.withdrawExchangeTask) {
        this.item = this.task.withdrawExchangeTask
        this.$emit('taskChanged', this.getTask())
      }
    }
  },
  computed: {
    GetNetworks(): Network[] {
      return this.networks
    },
  },
  async created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
    try {
      const res = await withdrawerService.withdrawerServiceListWithdrawer()
      this.withdrawers = res.withdrawers.map((w) => {
        w.label = `${w.label} [${w.exchangeType}]`
        return w
      })
    } finally {

    }
  }
})
</script>

<style scoped>

</style>
