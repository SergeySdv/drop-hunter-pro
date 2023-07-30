<template>
  <v-card-actions>
    <v-container>
      <div class="mb-3">use <a target="_blank" :href="link.orbiter">Orbiter.finance</a> to see
        available options
      </div>
      <v-row>
        <v-col cols="6">
          <v-select
            density="compact"
            variant="outlined"
            label="from network"
            :rules="[required]"
            :items="getFromNetworks"
            v-model="item.fromNetwork"
            :disabled="disabled"
            hide-details
          />
        </v-col>
        <v-col cols="6">
          <v-select
            density="compact"
            variant="outlined"
            label="to network"
            :rules="[required]"
            :items="getToNetworks"
            v-model="item.toNetwork"
            :disabled="disabled"
            hide-details
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-autocomplete
            density="compact"
            variant="outlined"
            label="from token"
            :rules="[required]"
            @click="fromTokenChanged"
            :items="fromTokens"
            v-model="item.fromToken"
            :disabled="disabled || !networkPairSet"
            item-title="code"
            return-object
            hide-details
          />
        </v-col>
        <v-col>
          <v-autocomplete
            return-object
            density="compact"
            variant="outlined"
            label="to token"
            :items="toTokens"
            :rules="[required]"
            v-model="item.toToken"
            :disabled="disabled || !networkPairSet || !item.fromToken"
            item-title="code"
            hide-details
          />
        </v-col>
      </v-row>
      <div class="my-2" v-if="networkPairSet && item.fromToken && item.toToken">{{
          `Swap possible - ${swapOpt.error ? "NO" : `YES [fee ${swapOpt.fee} USD min: ${swapOpt.min} max: ${swapOpt.max}]`}`
        }}
      </div>
      <AmountInput :coin="item.fromToken" :disabled="disabled" v-model="item.amount"/>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, OrbiterBridgeTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {orbiterService} from "@/generated/services";
import {required} from "@/components/tasks/menu/helper";
import {GetSwapOptionsRes} from "@/generated/orbiter";
import {link} from "@/components/tasks/links";

export default defineComponent({
  name: "TaskSwap1inch",
  components: {AmountInput, WEIInputField},
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
  computed: {
    link() {
      return link
    },
    networkPairSet(): boolean {
      return !!(this.item.toNetwork && this.item.fromNetwork)
    },
    Token() {
      return Token
    },
    getFromNetworks(): Network[] {
      return this.networks.filter((n) => n !== Network.ZKSYNCLITE)
    },
    getToNetworks(): Network[] {
      return this.networks.filter((n) => n !== this.item.fromNetwork)
    }
  },
  watch: {
    item: {
      handler() {
        this.getFromTokens()
        this.getToTokens()
        this.getSwapOpt()

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
      networksLoading: false,
      networks: [] as Network[],

      fromTokens: [] as Token[],
      fromTokensLoading: false,

      toTokens: [] as Token[],
      toTokensLoading: false,

      swapOpt: {min: "", max: "", fee: '', error: 'fill up the form above'} as GetSwapOptionsRes,
      swapOptLoading: false,

      item: {} as OrbiterBridgeTask,
    }
  },
  methods: {
    fromTokenChanged() {
      //@ts-ignore
      this.item.toToken = undefined
    },
    required,
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        orbiterBridgeTask: this.item,
        taskType: TaskType.OrbiterBridge,
        description: "",
      }
    },
    async getSwapOpt() {
      if (!this.networkPairSet) {
        return
      }
      if (!this.item.fromToken) {
        return
      }
      if (!this.item.toToken) {
        return
      }

      try {
        this.swapOptLoading = true
        const res = await orbiterService.orbiterServiceGetSwapOptions({
          body: {
            fromNetwork: this.item.fromNetwork,
            toNetwork: this.item.toNetwork,
            fromToken: this.item.fromToken,
            toToken: this.item.toToken,
          }
        })

        this.swapOpt.max = res.max
        this.swapOpt.min = res.min
        this.swapOpt.fee = res.fee
        this.swapOpt.error = res.error

      } finally {
        this.swapOptLoading = false
      }
    },
    async getToTokens() {
      if (!this.networkPairSet) {
        return
      }
      if (!this.item.fromToken) {
        return
      }
      try {
        this.toTokensLoading = true
        const res = await orbiterService.orbiterServiceGetToTokens({
          body: {
            fromNetwork: this.item.fromNetwork,
            toNetwork: this.item.toNetwork,
            fromToken: this.item.fromToken,
          }
        })
        this.toTokens = res.tokens
      } finally {
        this.toTokensLoading = false
      }
    },
    async getFromTokens() {
      if (!this.networkPairSet) {
        return
      }

      try {
        this.fromTokensLoading = true
        const res = await orbiterService.orbiterServiceGetFromTokens({
          body: {
            fromNetwork: this.item.fromNetwork,
            toNetwork: this.item.toNetwork,
          }
        })
        this.fromTokens = res.tokens
      } finally {
        this.fromTokensLoading = false
      }
    },
    async getNetworks() {
      try {
        this.networksLoading = true
        const res = await orbiterService.orbiterServiceGetNetworks()
        this.networks = res.networks
      } finally {
        this.networksLoading = false
      }
    },
    syncTask() {
      if (this.task) {
        if (this.task.orbiterBridgeTask) {
          this.item = this.task.orbiterBridgeTask
          this.$emit('taskChanged', this.getTask())
        }
      }
    }
  },
  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
    this.getNetworks()
  }
})
</script>

<style scoped>

</style>
