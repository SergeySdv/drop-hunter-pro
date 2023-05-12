<template>
  <div>
    <div class="d-flex justify-end">
      <v-btn v-if="selectedSome" color="black" @click=DeleteFlows class="ma-3">Delete flow</v-btn>
      <v-btn @click=CreateFlow class="ma-3">Create flow</v-btn>
    </div>
    <v-table>
      <thead>
      <tr>
        <th></th>
        <th class="text-left">Label</th>
        <th class="text-left">Details</th>
        <th class="text-left">View</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="item in getList" :key="item.id">
        <td>
          <CheckBox style="height: 40px" density="compact" :id="item.id" focused
                    @CheckboxChanged="CheckboxChanged"></CheckBox>
        </td>
        <td>{{ item.label }}</td>
        <td>{{ showStep(item) }}</td>
        <td>
          <v-btn color="blue" density="compact" @click="viewRoute(item.id)">View</v-btn>
        </td>
      </tr>
      </tbody>
    </v-table>
    <CreateFlow :showProp="showCreateFlowDialog" @change="CreateFlowChange" @Created="UpdateList"/>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService} from "@/generated/services"
import {flow_Flow as Flow, TaskType} from "@/generated/flow";
import CreateFlow from "@/components/CreateFlow.vue";
import CheckBox from "@/components/CheckBox.vue";

export default defineComponent({
  name: "Constructor",
  components: {
    CheckBox,
    CreateFlow
  },
  props: {},
  data() {
    return {
      list: [] as Flow[],
      showCreateFlowDialog: false,
      selected: new Set<string>(),
    }
  },
  computed: {
    getList(): Flow[] {
      return this.list
    },
    selectedSome(): boolean {
      return this.selected.size > 0
    }
  },
  methods: {
    viewRoute(id: string) {
      this.$router.push({name: "ViewFlow", params: {id: id}})
    },
    showStep(flow: Flow): string {
      if (!flow || !flow.tasks) {
        return "-"
      }
      let res = ""
      flow.tasks.forEach((task) => {
        switch (task.taskType) {
          case TaskType.StargateBridge: {
            res += `${task.weight}) ${TaskType.StargateBridge}`
          }
        }
      })
      return res
    },
    CheckboxChanged(id: string, value: boolean) {
      if (value) {
        this.selected.add(id)
      } else {
        this.selected.delete(id)
      }
    },
    async UpdateList() {
      const res = await flowService.flowServiceListFlow()
      this.list = res.flows
      this.selected = new Set<string>()
    },
    CreateFlow() {
      this.showCreateFlowDialog = true
    },
    CreateFlowChange(b: boolean) {
      this.showCreateFlowDialog = b
    },
    async DeleteFlows() {
      for (let id of this.selected.keys()) {
        await flowService.flowServiceDeleteFlow({body: {id: id}})
      }
      await this.UpdateList()
    }
  },
  created() {
    this.UpdateList()
  }
})


</script>


<style>


</style>

