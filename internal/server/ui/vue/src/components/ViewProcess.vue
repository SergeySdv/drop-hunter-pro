<template>

  <v-row class="text-h5 mt-4 ml-5 d-flex justify-sm-space-between">
    <v-col cols="12" sm="8" md="8" lg="8">
      {{ `${process.flowLabel} ` }}

      <StatusCard class=" ml-3" :status="process.status"/>
      <br/>
      <span style="font-size: 13px">{{ `(${process.id})` }}</span>
    </v-col>

    <v-col cols="12" sm="4" md="4" lg="4">
      <v-switch
        density="comfortable"
        focused
        v-model="showLess"
        :label="getCheckboxLabel"
        @input="showLessChanged">
      </v-switch>
      <BtnProcessStopResume size="comfortable" :item="process" @updated="processUpdated"/>
      <v-btn density="comfortable" class="ml-3" @click="$router.push({name: 'ViewFlow', params: {id: process.flowId}})">
        Show Flow
      </v-btn>
    </v-col>
  </v-row>


  <v-row class="mt-4 ml-5 mr-5 align-center" v-for="profile in process.profiles">
    <v-col cols="12" sm="2" md="2" lg="1">
      <ProfileCard :profile-id="profile.profileId"/>
    </v-col>
    <v-col cols="12" sm="1" md="1" lg="1">
      <StatusCard :status="profile.status"/>
    </v-col>
    <v-col cols="12" sm="9" md="9" lg="9">
      <v-timeline direction="horizontal" :density="density" line-inset="12">
        <v-timeline-item
          :dot-color="profile.status === ProcessStatus.StatusReady ? 'grey' : 'green'"
        >
          <template v-slot:opposite>
            Start
          </template>
        </v-timeline-item>


        <v-timeline-item v-for="task in profile.tasks">

          <template v-slot:icon>
            <ProcessTaskMenu :task="task" :profile-id="profile.profileId" :process-id="processId"
                             :process-profile-id="profile.id"/>
          </template>

          <template v-slot:opposite>
            {{ task.task.taskType }}
          </template>

        </v-timeline-item>

        <v-timeline-item
          :dot-color="profile.status === ProcessStatus.StatusDone ? 'green' : 'grey'"
        >
          <template v-slot:opposite>
            End
          </template>
        </v-timeline-item>
      </v-timeline>
    </v-col>
  </v-row>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {processService} from "@/generated/services"
import {Process, ProcessProfile, ProcessStatus, ProcessTask} from "@/generated/process";
import ViewFlow from "@/components/ViewFlow.vue";
import StatusCard from "@/components/StatusCard.vue";
import {GetStatusColor} from "@/components/helper";
import ProcessTaskMenu from "@/components/ProcessTaskMenu.vue";
import ProfileCard from "@/components/ProfileCard.vue";
import BtnProcessStopResume from "@/components/BtnProcessStopResume.vue";

export default defineComponent({
  name: "ViewProcess",
  components: {BtnProcessStopResume, ProfileCard, StatusCard, ViewFlow, ProcessTaskMenu},
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
      showLess: true,
      menu: false,
      processId: this.$route.params.id.toString(),
      process: {} as Process,
      profiles: [] as ProcessProfile[],
      updatedAt: new Date(1),
    }
  },
  computed: {
    getCheckboxLabel(): string {
      return !this.showLess ? 'more details' : 'less details'
    },
    ProcessStatus() {
      return ProcessStatus
    },
    getProfileList(): ProcessProfile[] {
      return this.profiles
    },

  },
  methods: {
    showLessChanged() {
      if (this.showLess) {
        this.density = "compact"
      } else {
        this.density = "comfortable"
      }
    },
    GetStatusColor,
    errorHappen(task: ProcessTask): boolean {
      return task.status == ProcessStatus.StatusError
    },
    processUpdated() {

    }
  },
  async created() {
    const init = async () => {
      if (this.$route.params.id) {

        const upadtedRes = await processService.processServiceGetProcessUpdatedAt({body: {processId: this.processId}})

        if (new Date(upadtedRes.updatedAt).getTime() > this.updatedAt.getTime()) {
          this.updatedAt = new Date(upadtedRes.updatedAt)
          const res = await processService.processServiceGetProcess({body: {id: this.processId}})
          this.process = res.process
          this.profiles = this.process.profiles
        }
      }
    }
    await init()
    setInterval(async () => {
      await init()
    }, 1000)

  },
})


</script>


<style>


</style>

