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
            v-on:change="inputChanged"
            :rules="[required]"
            :items="tokens"
            v-model="fromToken"
            :disabled="disabled"
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
            v-on:change="inputChanged"
            label="to token"
            :items="getToTokens"
            :rules="[required]"
            v-model="toToken"
            :disabled="true"
            item-title="code"
            hide-details
          />
        </v-col>
      </v-row>
      <AmountInput :coin="fromToken" :disabled="disabled" v-model="item.amount"/>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, Task, TaskType, Token, WETHTask} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {link} from "@/components/tasks/links";
import {required} from "@/components/tasks/menu/helper";

export default defineComponent({
  name: "TaskWethSwap",
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
    getToTokens(): Token[] {
      return this.tokens.filter((n) => n !== this.toToken)
    }
  },
  watch: {
    fromToken: {
      handler() {
        if (this.fromToken === Token.WETH) {
          this.toToken = Token.ETH
          this.item.wrap = false
        } else {
          this.toToken = Token.WETH
          this.item.wrap = true
        }
        this.$emit('taskChanged', this.getTask())
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
      fromToken: Token.ETH,
      toToken: Token.WETH,
      networks: [Network.ZKSYNCERA] as Network[],
      tokens: [Token.WETH, Token.ETH] as Token[],
      item: {} as WETHTask,
    }
  },
  methods: {
    required,
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        wETHTask: this.item,
        taskType: TaskType.WETH,
        description: "",
      }
    },
    inputChanged() {
    },
    syncTask() {
      if (this.task) {
        if (this.task.wETHTask) {
          this.item = this.task.wETHTask
          this.fromToken = this.item.wrap ? Token.ETH : Token.WETH
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
