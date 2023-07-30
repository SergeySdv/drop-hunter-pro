<template>
  <a :href="link.stargatefinance" target="_blank">stargate.finance</a>
  <div>Network: <b>{{ `${item.fromNetwork} to ${item.toNetwork}` }}</b></div>
  <div>Token: <b>{{ `${item.fromToken} to ${item.toToken}` }}</b></div>
  <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
  <div>Status:
    <span :style="`color: ${getStatusColor(getStatus)}`">{{ getStatus }}</span>
  </div>
  <div v-if="item.lzscanUrl"><a :href="item.lzscanUrl">LayerZero scan</a></div>
  <GasOptions :item="item.tx" :network="item.fromNetwork"/>
</template>

<script lang="ts">
import {AmUni, Network, StargateBridgeTask, SyncSwapTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ca} from "vuetify/locale";
import {processService} from "@/generated/services";
import {GetTaskSettingsResponse, ProcessStatus} from "@/generated/process";
import TaskSettings from "@/components/tasks/menu/TaskSettings.vue";
import {getAmountSend, getStatusColor} from "@/components/tasks/menu/helper";
import {link} from "@/components/tasks/links";
import GasOptions from "@/components/tasks/menu/GasOptions.vue";

export default defineComponent({
  name: "MenuTaskStargateBridge",
  methods: {getAmountSend, getStatusColor},
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
      item: {} as StargateBridgeTask,
    }
  },
  watch: {
    task: {
      handler() {
        if (this.task?.stargateBridgeTask) {
          this.item = this.task.stargateBridgeTask
        }
      },
      deep: true
    }
  },
  computed: {
    link() {
      return link
    },
    getStatus(): string {
      if (this.status == ProcessStatus.StatusDone) {
        return 'completed'
      }
      if (this.status == ProcessStatus.StatusError) {
        return 'error'
      }

      if (this.item.layerZeroStatus) {
        return 'waiting stargate swap complete'
      }

      if (!this.item.tx) {
        return 'not started'
      }

      if (this.item.txCompleted) {
        return 'transaction send'
      }

      return 'waiting'
    },
  },
  async created() {
    if (this.task?.stargateBridgeTask) {
      this.item = this.task.stargateBridgeTask
    }
  }
})
</script>

<style scoped>

</style>
