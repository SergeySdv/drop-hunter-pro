<template>
  <div>
    <Loader v-if="loading"/>
    <div v-else>
      <div class="d-inline-flex">
        <v-checkbox class="mx-2" color="blue" v-model="StatusReadyW" direction="horizontal" hide-details
                    density="compact">
          <template v-slot:label>
            <StatusCard :status="ProcessStatus.StatusReady"/>
          </template>
        </v-checkbox>
        <v-checkbox class="mx-2" color="blue" v-model="StatusRetryW" direction="horizontal" hide-details
                    density="compact">
          <template v-slot:label>
            <StatusCard :status="ProcessStatus.StatusRetry"/>
          </template>
        </v-checkbox>
        <v-checkbox class="mx-2" color="blue" v-model="StatusStopW" direction="horizontal" hide-details
                    density="compact">
          <template v-slot:label>
            <StatusCard :status="ProcessStatus.StatusStop"/>
          </template>
        </v-checkbox>
        <v-checkbox class="mx-2" color="blue" v-model="StatusErrorW" direction="horizontal" hide-details
                    density="compact">
          <template v-slot:label>
            <StatusCard :status="ProcessStatus.StatusError"/>
          </template>
        </v-checkbox>
        <v-checkbox class="mx-2" color="blue" v-model="StatusRunningW" direction="horizontal" hide-details
                    density="compact">
          <template v-slot:label>
            <StatusCard :status="ProcessStatus.StatusRunning"/>
          </template>
        </v-checkbox>
        <v-checkbox class="mx-2" color="blue" v-model="StatusDoneW" direction="horizontal" hide-details
                    density="compact">
          <template v-slot:label>
            <StatusCard :status="ProcessStatus.StatusDone"/>
          </template>
        </v-checkbox>
      </div>

      <v-list max-width="96vw" class="px-5">
        <v-list-item
          density="compact"
          variant="outlined"
          class="my-1 my-4"
          v-for="item in getList"
          :key="item.id"
          rounded
          height="auto"
          elevation="1"
          @click="viewProcess(item.id)"
          style="border: 1px solid "
        >
          <div class="d-flex justify-start">
            <StatusCard :status="item.status" class="mr-2"/>
            <div class="mr-2">
              <v-progress-linear
                :model-value="item.progress"
                color="blue"
                height="25"
                style="width: 100px"
              >
                <template v-slot:default="{ value }">
                  <strong>{{ Math.ceil(value) }}%</strong>
                </template>
              </v-progress-linear>
            </div>

            <div class="mr-2"><b>{{ `flow name: ${item.flowLabel}` }}</b></div>
            <div class="mr-2">Created: {{ dayjs(item.createdAt).format('YYYY-MM-DD HH:mm:ss') }}</div>
            <div class="mr-2">Updated: {{ dayjs(item.updatedAt).format('YYYY-MM-DD HH:mm:ss') }}</div>

          </div>
          <div class="mr-2">Profiles: <b>[{{ getProfiles(item) }}]</b></div>
          <div><b>Flow:</b>
            <div class="mr-2" v-for="(d, i) in getFlow(item.flow)">
              <b>{{ i + 1 }})</b> {{ d }}
            </div>
          </div>
        </v-list-item>
        <v-btn v-if="showNext" @click="next" width="100%">More</v-btn>
      </v-list>

    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {processService} from "@/generated/services"
import CreateFlow from "@/components/flow/CreateFlow.vue";
import CheckBox from "@/components/CheckBox.vue";
import {Process, ProcessStatus} from "@/generated/process";
import StatusCard from "@/components/StatusCard.vue";
import dayjs from "dayjs";
import BtnProcessStopResume from "@/components/BtnProcessStopResume.vue";
import {getFlow} from "@/components/tasks/tasks";
import Loader from "@/components/Loader.vue";

export default defineComponent({
  name: "Processes",
  watch: {
    StatusRetryW: {
      handler() {
        this.reset()
        this.UpdateList()
      }
    },
    StatusStopW: {
      handler() {
        this.reset()
        this.UpdateList()
      }
    },
    StatusErrorW: {
      handler() {
        this.reset()
        this.UpdateList()
      }
    },
    StatusRunningW: {
      handler() {
        this.reset()
        this.UpdateList()
      }
    },
    StatusDoneW: {
      handler() {
        this.reset()
        this.UpdateList()
      }
    },
    StatusReadyW: {
      handler() {
        this.reset()
        this.UpdateList()
      }
    },
  },
  components: {
    Loader,
    BtnProcessStopResume,
    StatusCard,
    CheckBox,
    CreateFlow
  },
  props: {},
  data() {
    return {
      StatusRetryW: false,
      StatusStopW: true,
      StatusErrorW: true,
      StatusRunningW: true,
      StatusDoneW: false,
      StatusReadyW: true,

      offset: 0,
      loading: false,
      loadingError: false,
      list: [] as Process[],
      showCreateFlowDialog: false,
      selected: new Set<string>(),
    }
  },
  computed: {
    showNext(): boolean {
      if (this.list.length === 0) {
        return false
      }
      return this.list.length % 15 === 0
    },
    ProcessStatus() {
      return ProcessStatus
    },
    dayjs() {
      return dayjs
    },
    getList(): Process[] {
      return this.list
    },
    selectedSome(): boolean {
      return this.selected.size > 0
    }
  },
  methods: {
    getFlow,
    viewProcess(id: string) {
      const routeData = this.$router.resolve({name: 'ViewProcess', params: {id: id}});
      window.open(routeData.href, '_blank');
    },
    getProfiles(p: Process): string {
      let labels: number[] = []
      p.profiles.forEach((pp) => {
        labels.push(Number(pp.profileLabel))
      })

      labels = labels.sort((a, b) => a - b)

      let out = ""
      labels.forEach((label) => {
        out += label + ", "
      })
      if (out.length) {
        out = out.substring(0, out.length - 2)
      }
      return out
    },
    next() {
      this.offset += 15
      this.UpdateList()
    },
    reset() {
      this.offset = 0
      this.list = []
    },
    processProgress(p: Process): number {
      let totalProfiles = p.profiles.length
      if (p.profiles[0]) {
        let totalTask = p.profiles[0].tasks.length
        let total = totalProfiles * totalTask
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

        return Math.ceil(finished / total)
      }
      return 0
    },
    viewRoute(id: string) {
      window.open("/process/" + id, "_blank")
    },
    getStatuses() {
      const statuses: ProcessStatus[] = []
      if (this.StatusDoneW) {
        statuses.push(ProcessStatus.StatusDone)
      }
      if (this.StatusReadyW) {
        statuses.push(ProcessStatus.StatusReady)
      }
      if (this.StatusErrorW) {
        statuses.push(ProcessStatus.StatusError)
      }
      if (this.StatusRetryW) {
        statuses.push(ProcessStatus.StatusRetry)
      }
      if (this.StatusStopW) {
        statuses.push(ProcessStatus.StatusStop)
      }
      if (this.StatusRunningW) {
        statuses.push(ProcessStatus.StatusRunning)
      }
      return statuses
    },
    async UpdateList() {
      try {

        this.loadingError = false
        this.loading = true
        const res = await processService.processServiceListProcess({
          body: {
            offset: this.offset.toString(),
            statuses: this.getStatuses(),
          }
        })
        this.list.push(...res.processes)
        this.selected = new Set<string>()
      } catch (e) {
        this.loadingError = true
      } finally {
        this.loading = false
      }
    },
    async cancel(id: string) {
      try {
        await processService.processServiceCancelProcess({body: {processId: id}})
      } finally {
        await this.UpdateList()
      }
    },
    processUpdated() {
      this.UpdateList()
    }
  },
  created() {
    this.UpdateList()
  }
})


</script>


<style scoped>

.header {
  margin: auto;
  height: auto;
  z-index: 100;
  right: 2px;
  left: 56px;
  background-color: white;;
  box-shadow: rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;
}
</style>

