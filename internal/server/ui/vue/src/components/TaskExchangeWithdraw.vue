<template>
  <div class="pa-3">
    <h3>{{ `${weight}) Exchange withdraw` }}</h3>
    <div>
      <v-container>
        <v-row>
          <v-col>
            <v-select
              ref="exchange-withdraw-form-withdraw-id"
              density="compact"
              variant="outlined"
              label="select withdrawer"
              :rules="withdrawIdRules()"
              :items="withdrawers"
              item-value="id"
              item-title="label"
              v-model="item.withdrawerId"
              :disabled="disabled"
            />
          </v-col>
        </v-row>

        <v-row v-if="item.withdrawerId">
          <v-col>
            <v-select
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
          </v-col>
          <v-col>
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
          </v-col>
        </v-row>
        <v-row>
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
      </v-container>
    </div>
  </div>
</template>

<script lang="ts">

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
      handler() {
        this.withdrawerChanged()
        this.validateWithdrawId()
        this.$emit('stepChanged', this.getTask())
      },
    },
    'item.token': {
      handler() {
        this.tokenChanged()
        this.validateToken()
        this.$emit('stepChanged', this.getTask())
      },
    },
    'item.network': {
      handler() {
        this.validateNetwork()
        this.$emit('stepChanged', this.getTask())
      },
    },
    'item.amountMax': {
      handler() {
        this.validateMax()
        this.$emit('stepChanged', this.getTask())
      },
    },
    'item.amountMin': {
      handler() {
        this.validateMin()
        this.$emit('stepChanged', this.getTask())
      },
    },
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
      } as WithdrawExchangeTask,
      options: new Map<string, ExchangeWithdrawNetwork[]>(),
    }
  },
  methods: {
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
          this.optionsLoading = true
          const res = await withdrawerService.withdrawerServiceGetExchangeWithdrawOptions({body: {withdrawerId: this.item.withdrawerId}})
          res.tokens.forEach(token => {
            const t: Token = {
              token: token.token,
              title: this.GetTokenTitle(token)
            }
            this.tokens.push(t)
            this.options.set(token.token, token.networks)
          })
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
        description: JSON.stringify(this.item),
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
  },
  computed: {
    GetNetworks(): Network[] {
      return this.networks
    },
  },
  async created() {
    this.$emit('stepChanged', this.getTask())
    if (this.task && this.task.withdrawExchangeTask) {
      this.item = this.task.withdrawExchangeTask
    }
    try {
      const withdrawers = await withdrawerService.withdrawerServiceListWithdrawer()
      this.withdrawers = withdrawers.withdrawers
    } finally {

    }

    if (this.task) {
      if (this.task.withdrawExchangeTask) {
        this.item = this.task.withdrawExchangeTask
      }
    }
  }
})
</script>

<style scoped>

</style>
