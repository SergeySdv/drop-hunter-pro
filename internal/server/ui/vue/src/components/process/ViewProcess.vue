<template>
  <div>
    <v-skeleton-loader width="100%" :loading="loading" type="table">
      <v-responsive>
        <div class="header">
          <v-progress-linear
            active
            :model-value="processProgress(process)"
            :color="getProgressColor"
            height="15"
          >
            <template v-slot:default="{ value }">
              <strong>{{ process.progress }}%</strong>
            </template>
          </v-progress-linear>

          <v-row class="pt-5 ml-5 d-flex justify-sm-space-between">
            <div class="pl-3 text-h5">
              {{ `${process.flowLabel} ` }}
              <StatusCard :status="process.status"/>
              <div>
                <span class="text-h6 text-blue-grey">Обновлено: {{ formatTime(process.updatedAt) }}</span>
              </div>
            </div>
            <div class="mr-5">
              <div>
                <v-checkbox
                  style="height: 30px"
                  color="blue"
                  v-model="showLess"
                  :label="getCheckboxLabel"
                  @input="showLessChanged">
                </v-checkbox>
                <v-checkbox
                  style="height: 30px"
                  color="blue"
                  :disabled="autoRetryLoading"
                  v-model="process.autoRetry"
                  :label="getAutoRetryLabel"
                  @input="autoRetryChanged">
                </v-checkbox>
              </div>
              <div class="mt-4">
                <BtnProcessStopResume class="mx-1" size="comfortable" :item="process" @updated="processUpdated"/>
                <v-btn class="mx-1" density="comfortable"
                       @click="$router.push({name: 'ViewFlow', params: {id: process.flowId}})">Flow
                </v-btn>
                <DeleteProcess class="mx-1" v-if="canCancel" :process-id="processId"/>
              </div>
            </div>
          </v-row>
        </div>

        <div style="padding-top: 170px;">
          <v-row class="mt-4 ml-5 mr-5 align-center" v-for="profile in process.profiles">
            <ProfileCard class="ml-5" :label="profile.profileLabel" :profile-id="profile.profileId"/>
            <StatusCard class="ml-5" :status="profile.status"/>

            <v-timeline
              class="mt-4"
              direction="horizontal"
              style="overflow-x: scroll; overflow-y: scroll; white-space:nowrap; width: 90vw"
              :density="density"
            >
              <v-timeline-item v-for="task in profile.tasks">
                <template v-slot:icon>
                  <ProcessTaskMenu :task="task" :profile-id="profile.profileId" :process-id="processId"
                                   :pp-id="profile.id"/>
                </template>
                <template v-slot:opposite>
                  {{ task.task.taskType }}
                </template>
              </v-timeline-item>
            </v-timeline>
          </v-row>
        </div>
      </v-responsive>
    </v-skeleton-loader>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {processService} from "@/generated/services"
import {Process, ProcessProfile, ProcessStatus, ProcessTask} from "@/generated/process";
import ViewFlow from "@/components/flow/ViewFlow.vue";
import StatusCard from "@/components/StatusCard.vue";
import {formatTime, GetStatusColor, Timer} from "@/components/helper";
import ProcessTaskMenu from "@/components/process/ProcessTaskMenu.vue";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import BtnProcessStopResume from "@/components/BtnProcessStopResume.vue";
import dayjs from "dayjs";
import DeleteProcess from "@/components/process/DeleteProcess.vue";

export default defineComponent({
  name: "ViewProcess",
  components: {DeleteProcess, BtnProcessStopResume, ProfileCard, StatusCard, ViewFlow, ProcessTaskMenu},
  props: {},
  watch: {
    loadingMap: {
      handler: () => {
      },
      deep: true
    }
  },

  data() {
    return {
      checkboxLabel: 'detailed',
      density: "compact" as 'compact' | 'comfortable',
      showLess: false,
      menu: false,
      processId: this.$route.params.id.toString(),
      process: {autoRetry: false} as Process,
      profiles: [] as ProcessProfile[],
      updatedAt: new Date(1),
      poller: {} as any,
      loading: true,
      autoRetryLoading: false,
    }
  },
  computed: {
    canCancel(): boolean {
      if (this.process.status == ProcessStatus.StatusStop) {
        return true
      }
      if (this.process.status == ProcessStatus.StatusDone) {
        return true
      }
      if (this.process.status == ProcessStatus.StatusError) {
        return true
      }
      return false
    },
    getProgressColor(): string {
      if (this.process.status === ProcessStatus.StatusRunning) {
        return 'blue'
      }
      if (this.process.status === ProcessStatus.StatusDone) {
        return 'green'
      }
      if (this.process.status === ProcessStatus.StatusError) {
        return 'red'
      }
      return 'grey'
    },
    getCheckboxLabel(): string {
      return !this.showLess ? 'details on' : 'details off'
    },
    getAutoRetryLabel(): string {
      return this.process.autoRetry ? 'Auto retry on' : 'Auto retry off'
    },
    ProcessStatus() {
      return ProcessStatus
    },
    getProfileList(): ProcessProfile[] {
      return this.profiles
    },

  },
  methods: {
    formatTime,
    dayjs,
    processProgress(p: Process): number {
      let totalProfiles = p.profiles.length
      if (p.profiles[0]) {
        let totalTask = p.profiles[0].tasks.length
        let total = totalProfiles * totalTask

        console.log(total)
        let finished = 0

        p.profiles.forEach((p) => {
          p.tasks.forEach((t) => {
            if (t.status == ProcessStatus.StatusDone) {
              finished++
            }
          })
        })

        if (finished == 0) {
          return 0
        }

        return Math.ceil(finished / total * 100)
      }
      return 0
    },
    showLessChanged() {
      if (this.showLess) {
        this.density = "compact"
      } else {
        this.density = "comfortable"
      }
    },
    async autoRetryChanged() {
      try {
        this.autoRetryLoading = true
        if (!this.process.autoRetry) {
          await processService.processServiceDisableAutoRetry({body: {processId: this.processId}})
        } else {
          await processService.processServiceEnableAutoRetry({body: {processId: this.processId}})
        }
      } finally {
        this.autoRetryLoading = false
      }
    },
    GetStatusColor,
    errorHappen(task: ProcessTask): boolean {
      return task.status == ProcessStatus.StatusError
    },
    processUpdated() {

    },
  },
  beforeUnmount() {
    clearTimeout(this.poller)
  },
  created: async function () {
    this.showLessChanged()
    const init = async () => {
      if (this.$route.params.id) {

        const upadtedRes = await processService.processServiceGetProcessUpdatedAt({body: {processId: this.processId}})

        if (new Date(upadtedRes.updatedAt).getTime() > this.updatedAt.getTime()) {
          this.updatedAt = new Date(upadtedRes.updatedAt)
          try {
            const res = await processService.processServiceGetProcess({body: {id: this.processId}})
            this.process = res.process
            this.profiles = this.process.profiles
            this.loading = true
          } finally {
            this.loading = false
          }
        }
      }
    }

    const period = import.meta.env.DEV ? 1000 : 1000

    await init()
    this.poller = setInterval(async () => {
      await init()
    }, period)

  },
})


</script>


<style>
.header {
  position: fixed;
  height: auto;
  min-height: 150px;
  z-index: 100;
  right: 2px;
  left: 56px;
  background-color: white;;
  box-shadow: rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;
}

</style>

