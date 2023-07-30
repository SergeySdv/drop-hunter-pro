<template>
  <div><b>Network:</b> {{ item.network }}</div>
  <div><b>Token:</b> {{ item.token }}</div>
  <div><b>Random amount range:</b> {{ `[ ${item.amountMin} : ${item.amountMax} ] ${item.token}` }}</div>
  <div v-if="item.amount"><b>Amount:</b> {{ item.sendAllCoins ? 'all coins on balance' : `${item.amount}` }}</div>
  <div><b>Withdraw to:</b> {{ item.useExternalAddr ? item.withdrawAddr : 'profile wallet' }}</div>
  <div><b>Status: </b> <span :style="'color: ' + getStatusColor(getStatus)">{{ getStatus }}</span></div>
</template>

<script lang="ts">
import {Task, WithdrawExchangeTask} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ProcessStatus} from "@/generated/process";
import {getStatusColor} from "@/components/tasks/menu/helper";

export default defineComponent({
  name: "MenuExchangeWithdraw",
  methods: {getStatusColor},
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
        if (this.task?.withdrawExchangeTask) {
          this.item = this.task.withdrawExchangeTask
        }
      },
      deep: true
    }
  },
  data() {
    return {
      item: {} as WithdrawExchangeTask,
    }
  },
  computed: {
    getStatus(): string {
      if (this.status == ProcessStatus.StatusError) {
        return 'error'
      }
      if (this.item.txId) {
        return 'completed'
      }

      if (this.item.withdrawOrderId) {
        return 'awaiting approval confirmation'
      }

      if (this.status === ProcessStatus.StatusReady) {
        return 'not started'
      }

      return 'sending transaction'
    }
  },
  async created() {
    if (this.task?.withdrawExchangeTask) {
      this.item = this.task.withdrawExchangeTask
    }
  }
})
</script>

<style scoped>

</style>
