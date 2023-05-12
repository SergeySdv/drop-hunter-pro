<template>
  <v-card-title>
    <span class="text-h5">Flow: {{ flow.label }}</span>
  </v-card-title>
  <v-container fluid>
    <v-row class="mb-8">
      <v-col>
        <v-autocomplete
          clearable
          return-object
          :loading="profilesLoading"
          v-model="selectedProfiles"
          :items="suggestedProfiles"
          item-children
          item-title="label"
          item-value="id"
          @update:search="ProfileSuggestChanged"
          multiple
          ref="label-input"
          label="select profiles"
          density="comfortable"
          variant="outlined"
        />
        <div v-for="profile in getProfileList">
          <ProfileCard :profile-id="profile.id"/>
        </div>
      </v-col>

      <v-col>
        <v-btn :disabled="selectedProfiles.length === 0" @click="startProcess">Start flow</v-btn>
      </v-col>
    </v-row>
    <v-spacer/>
    <v-form validate-on="submit" ref="flow-form">
      <v-row>
        <v-col cols="12" sm="6" md="4">
          <v-text-field
            ref="label-input"
            v-model="flow.label"
            label="label"
            density="comfortable"
            variant="outlined"
            :disabled="disabled"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-spacer/>
      <v-row>
        <v-col>
        </v-col>
        <v-col>

        </v-col>
      </v-row>
      <div v-for="item in getTasks">
        <component
          class="mb-5"
          @stepChanged="stepChanged"
          :task="item.task"
          :disabled="disabled"
          style="box-shadow: rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;"
          :is="item.c" :weight="item.i"></component>
      </div>
    </v-form>
  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService, processService, profileService} from "@/generated/services"
import {flow_Flow as Flow, Task, TaskType} from "@/generated/flow";
import TaskStargateBridge from "@/components/TaskStargateBridge.vue";
import {Profile} from "@/generated/profile";
import TaskDelay from "@/components/TaskDelay.vue";
import ProfileCard from "@/components/ProfileCard.vue";
import {stepComponentMap} from "@/components/tasks";

export default defineComponent({
  name: "ViewFlow",
  components: {ProfileCard},
  props: {
    propId: {
      type: String,
      required: false,
    }
  },

  data() {
    return {
      profilesLoading: false,
      flowId: "" as string,
      selectedProfiles: [] as Profile[],
      suggestedProfile: "",
      suggestedProfiles: [] as Profile[],
      flow: {} as Flow,
      disabled: true,
      stepComponents: [] as any[],
      stepMap: new Map<string, Task>(),
      stepTypes: [TaskType.StargateBridge, TaskType.Delay],
      saveLoading: false,
    }
  },
  computed: {
    getTasks() {
      return this.stepComponents
    },
    getProfileList() {
      return this.selectedProfiles
    }
  },
  methods: {
    async ProfileSuggestChanged(a: string) {

      // delay 1s
      this.suggestedProfile = a
      try {
        const snapshot = a
        await new Promise((resolve, reject) => {
          setTimeout(() => {
            if (this.suggestedProfile === snapshot) {
              resolve({})
            } else {
              reject({})
            }
          }, 1000)
        })
      } catch (e) {
        return
      }

      try {
        this.profilesLoading = true
        const res = await profileService.profileServiceSearchProfile({body: {pattern: a}})

        this.suggestedProfiles = res.profiles
      } finally {
        this.profilesLoading = false
      }

    },
    stepChanged(step: Task) {
      console.log(step)
      this.stepMap.set(step.weight, step)
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
  },
  async mounted() {

    console.log(this.propId)
    if (this.propId) {
      this.flowId = this.propId
    } else if (this.$route.params.id) {
      this.flowId = this.$route.params.id.toString()
    }


    if (this.$route.params.id) {
      const res = await flowService.flowServiceGetFlow({body: {id: this.flowId}})
      this.flow = res.flow
      this.flow.tasks.forEach((s) => {
        this.stepMap.set(s.weight, s)
        this.stepComponents.push({c: stepComponentMap.get(s.taskType), i: s.weight, task: s})
      })
    }

    const res = await profileService.profileServiceSearchProfile({body: {pattern: ""}})
    this.suggestedProfiles = res.profiles
  }
})


</script>


<style>


</style>

