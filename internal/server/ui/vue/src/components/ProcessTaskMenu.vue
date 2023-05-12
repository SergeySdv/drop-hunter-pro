<template>
  <v-menu
    v-model="menu"
    :close-on-content-click="false"
    location="end"
  >

    <template v-slot:activator="{ props }">
      <v-btn
        v-bind="props"
        :color="GetStatusColor(task.status)"
        @click="loadHistory"
      >Info
      </v-btn>
    </template>

    <template v-slot:default="{ props }">
      <v-card>
        <v-card-text>
          <b>{{ errorHappen(task) ? 'error:' + task.error : '' }}</b>
          <br v-if="errorHappen(task)"/>
          <v-divider/>
          <div style="white-space: pre-wrap;">{{ task.task.description.replaceAll(" ", "\n") }}</div>
          <div v-if="task.task.taskType === TaskType.StargateBridge">
            {{ GetStargateTaskFee }}
            <v-divider/>
          </div>

          <div v-for="record in records" :key="record.id">
            <v-divider/>
            {{ `${dayjs(record.startedAt).format('YYYY-MM-DD HH:mm:ss')}` }}
            <StatusCard :status="record.startStatus"/>
            ->
            <StatusCard :status="record.finishStatus ? record.finishStatus : ProcessStatus.StatusReady"/>
            {{ record.msg }}
          </div>

          <div v-if=" task.transactions.length > 0">
            Transactions:
            <div v-for="(tx, i) in task.transactions">
              <a :href="tx" target="_blank">tx: {{ i }}</a>
            </div>
          </div>

        </v-card-text>
        <v-card-actions>
          <v-btn v-if="errorHappen(task)" @click="retry(task, processProfileId)" :loading="retryLoading">Retry
          </v-btn>
        </v-card-actions>
      </v-card>
    </template>
  </v-menu>

</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {helperService, processService} from "@/generated/services"
import {Network, ProcessStatus, ProcessTask, ProcessTaskHistoryRecord, TaskType} from "@/generated/process";
import ViewFlow from "@/components/ViewFlow.vue";
import StatusCard from "@/components/StatusCard.vue";
import {GetStatusColor} from "@/components/helper";
import dayjs from "dayjs";

export default defineComponent({
  name: "ProcessTaskMenu",
  components: {StatusCard, ViewFlow},
  props: {
    task: {
      type: Object as PropType<ProcessTask>,
      required: true,
    },
    processProfileId: {
      type: String,
      required: true
    },
    processId: {
      type: String,
      required: true
    },
    profileId: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      menu: false,
      disabled: true,
      retryLoading: false,
      records: [] as ProcessTaskHistoryRecord[],
      stargateTaskFee: ""
    }
  },
  computed: {
    TaskType() {
      return TaskType
    },
    ProcessStatus() {
      return ProcessStatus
    },
    GetStargateTaskFee() {
      return this.stargateTaskFee
    }
  },
  methods: {
    dayjs,
    GetStatusColor,
    errorHappen(task: ProcessTask): boolean {
      return task.status == ProcessStatus.StatusError
    },
    async loadHistory() {
      if (!this.menu) {
        return
      }
      try {
        const res = await processService.processServiceGetProcessTaskHistory({body: {taskId: this.task.id}})
        this.records = res.records
      } finally {

      }
    },
    async getStargateBridgeFee(from: Network, to: Network, profileId: string) {
      try {
        const res = await helperService.helperServiceEstimateStargateBridgeFee({
          body: {
            from: from,
            to: to,
            profileId: profileId
          }
        })
        return JSON.stringify(res)
      } finally {

      }
    },
    async retry(task: ProcessTask, profileId: string) {
      try {
        this.retryLoading = true
        await processService.processServiceRetryProcess({
          body: {
            processId: this.processId,
            profileId: profileId,
            taskId: task.id,
          }
        })
      } finally {
        await new Promise(resolve => {
          setTimeout(() => {
            resolve(1)
          }, 1000)
        })
        this.retryLoading = false
      }

    },
  },
  async mounted() {
    await this.loadHistory()
    if (this.task.task.taskType === TaskType.StargateBridge) {

      const from = this.task.task.stargateBridgeTask?.fromNetwork
      const to = this.task.task.stargateBridgeTask?.toNetwork
      if (from && to) {
        this.stargateTaskFee = await this.getStargateBridgeFee(from, to, this.profileId)
      }
    }

  },
})


</script>


<style>


</style>

