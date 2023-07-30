<template>
  <v-card-actions>
    <v-container>
      <div class="mb-3">use <a href="https://stargate.finance/transfer" target="_blank">stargate swap</a> to see
        available swap
        options
      </div>
      <v-row>
        <v-col>
          <v-select
            ref="stargate-bridge-form"
            density="compact"
            variant="outlined"
            label="network"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="networks"
            v-model="network"
            :disabled="disabled"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-autocomplete
            density="compact"
            variant="outlined"
            label="from token"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="tokens"
            v-model="fromToken"
            :disabled="disabled"
            :loading="tokenLoading"
            item-title="code"
            return-object
          />
        </v-col>
        <v-col>
          <v-autocomplete
            return-object
            density="compact"
            variant="outlined"
            v-on:change="inputChanged"
            label="to token"
            :items="tokens"
            :rules="[required]"
            v-model="toToken"
            :disabled="disabled"
            :loading="tokenLoading"
            item-title="code"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-text-field
            density="compact"
            variant="outlined"
            type="number"
            label="estimated amount in ETH"
            :rules="[required]"
            v-model="estimatedAmount"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-autocomplete
            return-object
            density="compact"
            variant="outlined"
            v-on:change="inputChanged"
            label="select protocol"
            :items="protocols"
            :rules="[required]"
            v-model="protocol"
            :disabled="disabled"
            :loading="protocolLoading"
            item-title="name"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, StargateBridgeTask, Swap1inchTask, Task, TaskType, Token} from "@/generated/flow";
import {swap1inchService} from "@/generated/services";
import {defineComponent, PropType} from "vue";
import {GetSwapOptionsProtocol, Swap1inchToken} from "@/generated/swap1inch";
import {fi} from "vuetify/locale";

export default defineComponent({
  name: "TaskSwap1inch",
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
    }
  },

  watch: {
    network: {
      handler() {
        if (this.network) {
          this.item.network = this.network
        }
      },
    },
    fromToken: {
      handler() {
        if (this.fromToken) {
          this.item.fromTokenAddr = this.fromToken.addr
          this.item.fromTokenName = this.fromToken.name
          this.item.fromTokenCode = this.fromToken.code
          this.listSwapOptions()
        }
      },
    },
    toToken: {
      handler() {
        if (this.toToken) {
          this.item.toTokenAddr = this.toToken.addr
          this.item.toTokenName = this.toToken.name
          this.item.toTokenCode = this.toToken.code
          this.listSwapOptions()
        }
      },
    },
    estimatedAmount: {
      handler() {
        this.listSwapOptions()
      },
    },
    item: {
      handler() {
        this.$emit('taskChanged', this.getTask())
      },
      deep: true
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
      protocol: null as GetSwapOptionsProtocol | null,
      protocols: [] as GetSwapOptionsProtocol[],
      protocolLoading: false,
      estimatedAmount: null as Number | null,
      tokenLoading: false,
      networks: [] as Network[],
      tokens: [] as Swap1inchToken[],
      fromToken: null as Swap1inchToken | null,
      toToken: null as Swap1inchToken | null,
      network: null as Network | null,
      item: {
        network: Network.ARBITRUM,
        fromTokenAddr: '',
        fromTokenCode: '',
        fromTokenName: '',
        toTokenAddr: '',
        toTokenCode: '',
        toTokenName: '',
      } as Swap1inchTask,
      required: (v: any) => !!v || 'required'
    }
  },
  methods: {
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        swap1inchTask: this.item,
        taskType: TaskType.Swap1inch,
        description: JSON.stringify(this.item),
      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['stargate-bridge-form'].validate()

      return valid
    },

    inputChanged() {
      return this.validateForm()
    },

    async listSwapOptions() {
      if (!(this.fromToken && this.toToken && this.estimatedAmount)) {
        return
      }
      try {
        this.protocolLoading = true
        const res = await swap1inchService.swap1InchServiceGetSwapOptions({
          body: {
            network: this.item.network,
            amount: String(this.estimatedAmount),
            fromTokenAddr: this.item.fromTokenAddr,
            toTokenAddr: this.item.toTokenAddr,
          }
        })
        this.protocols = res.protocols
      } finally {
        this.protocolLoading = false
      }

    },
    async listTokens() {
      try {
        this.tokenLoading = true
        const res = await swap1inchService.swap1InchServiceGetTokens({body: {network: this.item.network}})
        this.tokens = res.tokens
      } finally {
        this.tokenLoading = false
      }
    },
    async listNetworks() {
      const res = await swap1inchService.swap1InchServiceGetNetworks()
      this.networks = res.networks
    },
    syncTask() {
      if (this.task) {
        if (this.task.swap1inchTask) {
          this.item = this.task.swap1inchTask
          this.$emit('taskChanged', this.getTask())
        }
      }
    }
  },
  created() {
    this.listNetworks()
    this.listTokens()
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
})
</script>

<style scoped>

</style>
