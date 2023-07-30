<template>
  <v-card-actions>
    <v-container>
      <div class="mb-3">use <a :href="link.syncSwap" target="_blank">SyncSwap</a> to see
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
            v-model="item.network"
            :disabled="disabled"
          />
        </v-col>
        <v-col>
          <v-autocomplete
            return-object
            density="compact"
            variant="outlined"
            v-on:change="inputChanged"
            label="direction"
            :items="pairs"
            :rules="[required]"
            v-model="pair"
            :disabled="disabled"
            item-title="name"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <AmountInput :coin="item.fromToken" :disabled="disabled" v-model="item.amount"/>
        </v-col>
        <v-col>
          <v-col>Slippage: 0.1%</v-col>
        </v-col>
      </v-row>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, SyncSwapTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {link} from "@/components/tasks/links";
import {SwapPair, tokenSwapPair} from "@/components/helper";
import {required} from "@/components/tasks/menu/helper";

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
    Token() {
      return Token
    },
  },
  watch: {
    pair: {
      handler() {
        this.item.fromToken = this.pair.from
        this.item.toToken = this.pair.to
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
      pairs: [
        tokenSwapPair(Token.ETH, Token.USDC),
        tokenSwapPair(Token.USDC, Token.ETH),

        tokenSwapPair(Token.ETH, Token.LUSD),
        tokenSwapPair(Token.LUSD, Token.ETH),

        tokenSwapPair(Token.ETH, Token.LSD),
        tokenSwapPair(Token.LSD, Token.ETH),

        tokenSwapPair(Token.ETH, Token.MUTE),
        tokenSwapPair(Token.MUTE, Token.ETH),

      ] as SwapPair[],
      pair: null as SwapPair | null,
      networks: [Network.ZKSYNCERATESTNET, Network.ZKSYNCERA] as Network[],
      item: {
        network: Network.ZKSYNCERA,
        amount: {
          sendAll: true,
        },
        toToken: Token.USDC,
        fromToken: Token.ETH,
      } as SyncSwapTask,
    }
  },
  methods: {
    required,
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        syncSwapTask: this.item,
        taskType: TaskType.SyncSwap,
        description: "",
      }
    },
    inputChanged() {
    },
    syncTask() {
      if (this.task) {
        if (this.task.syncSwapTask) {
          this.item = this.task.syncSwapTask
          this.pair = tokenSwapPair(this.item.fromToken, this.item.toToken)
          this.$emit('taskChanged', this.getTask())
        }
      }
    }
  },
  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
})
</script>

<style scoped>

</style>
