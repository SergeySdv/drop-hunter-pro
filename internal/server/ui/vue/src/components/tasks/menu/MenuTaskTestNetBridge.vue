<template>
  <a :href="link.testnetbridge" target="_blank">TestNetBridge</a>
  <div>Network: <b>{{ `${item.network}` }}</b></div>
  <div>Settings: <b>{{ `[ ${item.minAmount} : ${item.maxAmount} ] ETH` }}</b></div>
  <div v-if="item.amount">Amount ETH: <b>{{ `${item.amount}` }}</b></div>
  <div>Status:
    <span :style="`color: ${getStatusColor(getStatus)}`">{{ getStatus }}</span>
  </div>
</template>

<script lang="ts">
import {
  AmUni,
  Network,
  StargateBridgeTask,
  SyncSwapTask,
  Task,
  TaskType,
  TestNetBridgeSwapTask,
  Token
} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ca} from "vuetify/locale";
import {processService} from "@/generated/services";
import {GetTaskSettingsResponse, ProcessStatus} from "@/generated/process";
import TaskSettings from "@/components/tasks/menu/TaskSettings.vue";
import {getStatusColor} from "@/components/tasks/menu/helper";
import {link} from "@/components/tasks/links";

export default defineComponent({
  name: "MenuTaskTestNetBridge",
  methods: {getStatusColor},
  components: {TaskSettings},
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
      item: {} as TestNetBridgeSwapTask,
    }
  },
  watch: {
    task: {
      handler() {
        if (this.task?.testNetBridgeSwapTask) {
          this.item = this.task.testNetBridgeSwapTask
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

      if (this.status == ProcessStatus.StatusReady) {
        return 'not started'
      }


      return 'waiting'
    },
  },
  async created() {
    if (this.task?.testNetBridgeSwapTask) {
      this.item = this.task.testNetBridgeSwapTask
    }
  }
})
</script>

<style scoped>

</style>
