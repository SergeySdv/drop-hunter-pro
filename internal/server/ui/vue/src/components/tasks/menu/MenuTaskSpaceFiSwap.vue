<template>
  <a :href="taskProps.SpaceFISwap.service.link" target="_blank">{{ taskProps.SpaceFISwap.service.name }}</a>
  <div>Network: <b>{{ item.network }}</b></div>
  <div>Swap: <b>{{ `${item.fromToken} to ${item.toToken}` }}</b></div>
  <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
  <div>Slippage: 0.1%</div>
  <div>Status:
    <span :style="`color: ${getTxStatusColor}`">{{ getTxStatus }}</span>
  </div>
  <GasOptions :item="item.tx" :network="item.network"/>
</template>

<script lang="ts">
import {AmUni, MaverickSwapTask, Network, SpaceFiSwapTask, SyncSwapTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ca} from "vuetify/locale";
import {processService} from "@/generated/services";
import {GetTaskSettingsResponse, ProcessStatus} from "@/generated/process";
import TaskSettings from "@/components/tasks/menu/TaskSettings.vue";
import {getAmountSend} from "./helper";
import {link} from "@/components/tasks/links";
import {taskProps} from "@/components/tasks/tasks";
import GasOptions from "@/components/tasks/menu/GasOptions.vue";

export default defineComponent({
  name: "MenuTaskSpaceFiSwap",
  components: {GasOptions, TaskSettings},
  props: {
    task: {
      type: Object as PropType<Task>,
      required: true,
    },
    status: {
      type: String as PropType<ProcessStatus>,
      required: true,
    }
  },
  data() {
    return {
      item: {} as SpaceFiSwapTask,
    }
  },
  watch: {
    task: {
      handler() {
        if (this.task?.spaceFiSwapTask) {
          this.item = this.task.spaceFiSwapTask
        }
      },
      deep: true
    }
  },
  computed: {
    taskProps() {
      return taskProps
    },
    link() {
      return link
    },
    getTxStatusColor(): string {
      const s = this.getTxStatus
      switch (s) {
        case 'error':
          return 'red'
        case 'not started':
          return 'grey'
        case 'completed':
          return 'green'
        default:
          return 'blue'
      }
    },
    getTxStatus(): string {
      if (this.status == ProcessStatus.StatusError) {
        return 'error'
      }
      if (!this.item.tx) {
        return 'not started'
      }

      if (this.item.tx.txCompleted) {
        return 'completed'
      }

      return 'waiting'
    },
  },
  methods: {
    getAmountSend,
    getBalance(am?: AmUni): string {
      if (!am) {
        return ''
      }

      return `${Number(am.usd).toPrecision(3)} USD`
      // return `${Number(am.eth).toPrecision(3)} ETH ${Number(am.usd).toPrecision(3)} USD`
    },

  },
  async created() {
    if (this.task?.spaceFiSwapTask) {
      this.item = this.task.spaceFiSwapTask
    }
  }
})
</script>

<style scoped>

</style>
