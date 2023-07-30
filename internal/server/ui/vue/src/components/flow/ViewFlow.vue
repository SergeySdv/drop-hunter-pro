<template>
  <div class="position-fixed header">
    <div class="mx-2 my-2 d-flex justify-space-between ma-auto">
      <span class="text-h5">Flow: {{ flow.label }}</span>
      <div>
        <v-btn v-if="!flow.deletedAt && !flow.nextId" color="grey" @click="DeleteFlow">Delete</v-btn>
        <v-btn class="mx-2" :disabled="selectedProfiles.length === 0" @click="startProcess">Start process</v-btn>
        <span v-if="!flow.deletedAt && !flow.nextId">
          <v-btn v-if="!editMode" focused @click="editModeChanged">Update</v-btn>
          <v-btn v-else @click="updateFlow">Save changed</v-btn>
          </span>
      </div>
    </div>
    <div class="mx-8 my-3">
      <ProfileSearch
        v-model="selectedProfiles"
        :disabled="editMode"
      />
      <span class="ml-3 mb-2" v-for="profile in getProfileList">
          <ProfileCard :label="profile.num" :profile-id="profile.id"/>
        </span>
    </div>
  </div>


  <v-spacer class="my-6" style="padding-top: 200px"/>
  <v-form validate-on="submit" ref="flow-form">
    <FlowForm v-if="!flowLoading" :disable="disabled" :label-value="flow.label" :tasks-value="tasks"
              @flow-changed="flowChanged"/>
  </v-form>


</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService, processService, profileService} from "@/generated/services"
import {flow_Flow as Flow, Task, TaskType} from "@/generated/flow";
import TaskStargateBridge from "@/components/tasks/TaskStargateBridge.vue";
import {Profile} from "@/generated/profile";
import TaskDelay from "@/components/tasks/TaskDelay.vue";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import {taskComponentMap, TaskArg, taskTypes} from "@/components/tasks/tasks";
import {Delay, Timer} from "@/components/helper";
import FlowForm from "@/components/flow/FlowForm.vue";
import ProfileSearch from "@/components/profile/ProfileSearch.vue";

export default defineComponent({
  name: "ViewFlow",
  components: {ProfileSearch, FlowForm, ProfileCard},
  props: {
    propId: {
      type: String,
      required: false,
    }
  },
  data() {
    return {
      selectedProfiles: [] as Profile[],
      editMode: false,
      flowId: "" as string,
      flow: {} as Flow,
      disabled: true,
      tasks: [] as Task[],
      showUpdateBtn: false,
      updatingFlow: false,
      timer: new Timer(),
      flowLoading: true
    }
  },
  computed: {
    getProfileList() {
      return this.selectedProfiles.map(e => e)
    }
  },
  methods: {
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['flow-form'].validate()

      return valid
    },
    flowChanged(label: string, tasks: TaskArg[]) {
      this.flow.label = label
      this.tasks = []
      tasks.forEach(t => {
        if (t.task) {
          this.tasks.push(t.task)
        }
      })
    },
    async updateFlow() {
      if (!(await this.validateForm())) {
        return
      }
      try {
        this.updatingFlow = true
        this.flow.tasks = this.tasks

        const res = await flowService.flowServiceUpdateFlow({body: {flow: this.flow}})
        this.flow = res.flow
        this.flowId = this.flow.id
        this.editModeChanged()
        this.$router.push({name: 'ViewFlow', params: {id: this.flow.id}})

      } finally {
        this.updatingFlow = false
      }
    },
    editModeChanged() {
      this.editMode = !this.editMode
      this.disabled = !this.editMode
      this.showUpdateBtn = this.editMode
    },
    getCheckboxLabel(): string {
      return this.editMode ? "sw" : 'switch tp editing mode'
    },
    async startProcess() {
      const res = await processService.processServiceCreateProcess({
        body: {
          flowId: this.flowId,
          profileIds: this.selectedProfiles.map((p) => p.id),
        }
      })
      this.$router.push({name: "ViewProcess", params: {id: res.process.id}})
    },
    async DeleteFlow() {
      try {
        await flowService.flowServiceDeleteFlow({body: {id: this.flowId}})
        this.$router.push({name: 'Constructor'})
      } catch (e) {

      }

    }
  },
  async mounted() {

    if (this.propId) {
      this.flowId = this.propId
    } else if (this.$route.params.id) {
      this.flowId = this.$route.params.id.toString()
    }

    try {
      this.flowLoading = true
      const res = await flowService.flowServiceGetFlow({body: {id: this.flowId}})
      this.flow = res.flow
      this.tasks = this.flow.tasks
    } finally {
      this.flowLoading = false
    }
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

