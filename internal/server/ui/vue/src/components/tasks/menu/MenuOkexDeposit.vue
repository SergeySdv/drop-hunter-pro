<template>
  <div><b>Network:</b> {{ item.network }}</div>
  <div><b>Token:</b> {{ item.token }}</div>
  <div v-if="item.address"><b>Okex wallet address:</b> {{ item.address }}</div>
  <div><b>Status: </b> <span :style="'color: ' + getStatusColor(getStatus)">{{ getStatus }}</span></div>
  <GasOptions :item="item.tx" :network="item.network"/>
</template>

<script lang="ts">
import {OkexDepositTask, Task} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {formatTime, humanDuration} from "@/components/helper";
import {ProcessStatus} from "@/generated/process";
import {getStatusColor} from "@/components/tasks/menu/helper";
import GasOptions from "@/components/tasks/menu/GasOptions.vue";

export default defineComponent({
  name: "MenuDelayTask",
  components: {GasOptions},
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
  watch: {
    item: {
      handler() {
        if (this.task?.okexDepositTask) {
          this.item = this.task.okexDepositTask
        }
      },
      deep: true
    }
  },
  data() {
    return {
      item: {} as OkexDepositTask,
    }
  },
  computed: {
    getStatus(): string {
      if (this.status == ProcessStatus.StatusError) {
        return 'error'
      }

      if (this.status == ProcessStatus.StatusDone) {
        return 'completed'
      }

      if (!this.item.tx) {
        return 'sending transaction'
      }

      if (!this.item.subMainTransfer) {
        return 'transfer from sub account to main account'
      }

      return 'not started'
    }
  },
  methods: {getStatusColor, humanDuration, formatTime},
  async created() {
    if (this.task?.okexDepositTask) {
      this.item = this.task.okexDepositTask
    }
  }
})
</script>

<style scoped>

</style>
