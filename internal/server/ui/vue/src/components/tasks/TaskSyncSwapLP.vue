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
            :disabled="true"
            hide-details
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-autocomplete
            density="compact"
            variant="outlined"
            label="pool token a"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="tokens"
            v-model="item.a"
            :disabled="true"
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
            label="pool token b"
            :items="getToTokens"
            :rules="[required]"
            v-model="item.b"
            :disabled="true"
            item-title="code"
            hide-details
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          Operation:
          {{ item.add ? 'deposit' : 'withdraw' }}
          <v-btn
            :disabled="disabled"
            @click="item.add = !item.add"
          >Change
          </v-btn>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          {{ `50% ${item.a} 50% ${item.b}` }}
        </v-col>
      </v-row>
      <AmountInput :coin="item.a" :disabled="disabled" v-model="item.amount"/>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, SyncSwapLPTask, SyncSwapTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {link} from "@/components/tasks/links";
import {required} from "@/components/tasks/menu/helper";
import {taskProps} from "@/components/tasks/tasks";

export default defineComponent({
  name: "TaskSyncSwapLP",
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

      if (this.item.fromToken === Token.ETH) {
        return this.tokens.filter((n) => n !== this.item.toToken)
      }

      return [Token.ETH]
    }
  },
  watch: {
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
      networks: [Network.ZKSYNCERA] as Network[],
      tokens: [Token.ETH, Token.USDC] as Token[],
      item: {
        network: Network.ZKSYNCERA,
        amount: {
          sendAll: true,
        },
        a: Token.USDC,
        b: Token.ETH,
      } as SyncSwapLPTask,
    }
  },
  methods: {
    required,
    getTask(): Task {
      const task = {
        weight: this.weight.toString(),
        syncSwapLPTask: this.item,
        taskType: TaskType.SyncSwapLP,
        description: ''
      }
      task.description = taskProps[TaskType.SyncSwapLP].descFn(task)

      return task
    },
    inputChanged() {
    },
    syncTask() {
      if (this.task) {
        if (this.task.syncSwapLPTask) {
          this.item = this.task.syncSwapLPTask
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
