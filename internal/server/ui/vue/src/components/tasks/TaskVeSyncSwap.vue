<template>
  <v-card-actions>
    <v-container>
      <div class="mb-3">use<a :href="taskProps.VeSyncSwap.service.link"
                              target="_blank">{{ taskProps.VeSyncSwap.service.name }}</a> to see
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
          />
        </v-col>
        <v-col>
          <v-autocomplete
            density="compact"
            variant="outlined"
            label="direction"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="pairs"
            v-model="pair"
            :disabled="disabled"
            item-title="name"
            return-object
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
import {MaverickSwapTask, Network, SpaceFiSwapTask, Task, TaskType, Token, VeSyncSwapTask} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {link} from "@/components/tasks/links";
import {taskProps} from "@/components/tasks/tasks";
import {required} from "@/components/tasks/menu/helper";
import {SwapPair, tokenSwapPair} from "@/components/helper";


export default defineComponent({
  name: "TaskVeSyncSwap",
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
    taskProps() {
      return taskProps
    },
    link() {
      return link
    },
    Token() {
      return Token
    },
  },
  watch: {
    "pair": {
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
      ] as SwapPair[],
      pair: null as SwapPair | null,
      networks: [Network.ZKSYNCERA] as Network[],
      item: {
        network: Network.ZKSYNCERA,
        amount: {
          sendAll: true,
        },
        toToken: Token.USDC,
        fromToken: Token.ETH,
      } as VeSyncSwapTask,
    }
  },
  methods: {
    required,
    getTask(): Task {
      const taskType = TaskType.VeSyncSwap
      const task = {
        weight: this.weight.toString(),
        veSyncSwapTask: this.item,
        taskType: taskType,
        description: "",
      }
      task.description = taskProps[taskType].descFn(task)
      return task
    },
    inputChanged() {
    },
    syncTask() {
      if (this.task) {
        if (this.task.veSyncSwapTask) {
          this.item = this.task.veSyncSwapTask
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
