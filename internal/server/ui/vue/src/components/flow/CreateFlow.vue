<template>
  <div class="position-fixed"
       style="
       margin: auto; height: 50px; z-index: 100; right: 2px; left: 56px;
        background-color: white; ;
      box-shadow: rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;">
    <div class="mx-2 my-2 d-flex justify-space-between ma-auto">
      <span class="text-h5"><b>Flow creation</b></span>
      <v-btn
        :disabled="saveLoading"
        :loading="saveLoading"
        @click="CreateFlow">
        Next
      </v-btn>
    </div>
  </div>
  <v-card style="padding-top: 70px">
    <v-card-text>
      <v-form validate-on="submit" ref="flow-form">
        <FlowForm @flow-changed="flowChanged"/>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService} from "@/generated/services"
import {CreateFlowRequest, Task} from "@/generated/flow";
import {TaskArg, taskTypes} from '@/components/tasks/tasks'
import FlowForm from "@/components/flow/FlowForm.vue";

export default defineComponent({
  name: "CreateFlow",
  components: {FlowForm},
  data() {
    return {
      tasks: [] as Task[],
      stepTypes: taskTypes,
      show: this.showProp,
      item: {} as CreateFlowRequest,
      saveLoading: false,
    }
  },
  methods: {
    flowChanged(label: string, tasks: TaskArg[]) {
      this.tasks = []
      tasks.forEach(t => {
        if (t.task) {
          this.tasks.push(t.task)
        }
      })
      this.item.label = label
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
        this.item.tasks = this.tasks

        await flowService.flowServiceCreateFlow({body: this.item})
        this.$router.push({name: "Constructor"})
      } finally {
        this.saveLoading = false
      }

    },
  }
})


</script>


<style>


</style>

