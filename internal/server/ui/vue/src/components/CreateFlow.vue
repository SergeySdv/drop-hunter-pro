<template>
  <v-row justify="center">
    <v-dialog
      v-model="show"
      persistent
      width="80%"
      height="80%"
      min-height="80"
    >
      <v-card>
        <v-card-title>
          <span class="text-h5">Create flow</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-form validate-on="submit" ref="flow-form">
              <v-row>
                <v-col cols="12" sm="6" md="4">
                  <v-text-field
                    ref="label-input"
                    v-model="item.label"
                    label="label"
                    density="comfortable"
                    variant="outlined"
                    :rules="labelRules()"
                    @input="labelChanged"
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-spacer/>
              <v-row>
                <v-col>
                  <v-select
                    v-model="selectedStep"
                    :items="stepTypes"
                  />
                </v-col>
                <v-col>
                  <v-btn class="ma-2" @click="addStep">Add step</v-btn>
                </v-col>
              </v-row>
              <div v-for="item in getSteps">
                <component
                  class="mb-5"
                  @stepChanged="stepChanged"
                  style="box-shadow: rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;"
                  :disabled="false"
                  :is="item.c" :weight="item.i"></component>
              </div>
            </v-form>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue-darken-1"
            variant="text"
            @click="Close"
          >
            Close
          </v-btn>
          <v-btn
            type="submit"
            color="blue-darken-1"
            variant="text"
            :disabled="saveLoading"
            :loading="saveLoading"
            @click="CreateFlow">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService} from "@/generated/services"
import {CreateFlowRequest, Task, TaskType} from "@/generated/flow";
import {stepComponentMap, taskTypes} from '@/components/tasks'

export default defineComponent({
  name: "CreateFlow",
  props: {
    showProp: Boolean,
  },
  watch: {
    showProp: {
      handler(newValue, oldValue) {
        this.show = newValue
      },
    }
  },
  data() {
    return {
      selectedStep: TaskType.Delay,
      stepComponents: [] as any[],
      stepMap: new Map<string, Task>(),
      stepTypes: taskTypes,
      show: this.showProp,
      item: {} as CreateFlowRequest,
      saveLoading: false,
    }
  },
  computed: {
    getSteps() {
      return this.stepComponents
    }
  },
  methods: {
    async addStep() {
      this.stepComponents.push({c: stepComponentMap.get(this.selectedStep), i: this.stepComponents.length + 1});
    },
    stepChanged(step: Task) {
      console.log(step)
      this.stepMap.set(step.weight, step)
    },
    labelRules() {
      return [
        (v: string) => this.item.label ? true : 'label is required'
      ]
    },
    async labelChanged(e: InputEvent) {
      await this.validate()
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['label-input'].validate()

      return valid
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['flow-form'].validate()

      return valid
    },
    async CreateFlow() {
      if (!(await this.validateForm())) {
        return
      }

      try {
        this.saveLoading = true

        const tasks: Task[] = []
        this.stepMap.forEach((step) => {
          tasks.push(step)
        })
        this.item.tasks = tasks

        await flowService.flowServiceCreateFlow({body: this.item})
        this.$emit("change", false)
        this.$emit("Created")
      } catch (e) {
        console.error(e)
        this.saveLoading = false
      }

    },
    Close() {
      this.$emit("change", false)
    }
  }
})


</script>


<style>


</style>

