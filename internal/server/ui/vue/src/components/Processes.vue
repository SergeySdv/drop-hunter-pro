<template>
  <div>
    <v-table>
      <thead>
      <tr>
        <th class="text-left">Flow label</th>
        <th class="text-left">Status</th>
        <th class="text-left">Updated</th>
        <th class="text-left">Actions</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="item in getList" :key="item.id">
        <td>{{ `${item.flowLabel}` }}</td>
        <td>
          <StatusCard :status="item.status"/>
        </td>
        <td>{{ dayjs(item.updatedAt).format('YYYY-MM-DD HH:mm:ss') }}</td>
        <td>

          <v-btn class="mr-2" density="compact" @click="viewRoute(item.id)" color="blue">View</v-btn>
          <BtnProcessStopResume size="compact" :item="item" @updated="processUpdated"/>
          <!--          <v-btn class="mr-2" density="compact" @click="Cancel(item.id)" color="red">Cancel</v-btn>-->
        </td>
      </tr>
      </tbody>
    </v-table>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService, processService} from "@/generated/services"
import {flow_Flow as Flow, TaskType} from "@/generated/flow";
import CreateFlow from "@/components/CreateFlow.vue";
import CheckBox from "@/components/CheckBox.vue";
import {Process, ProcessStatus} from "@/generated/process";
import StatusCard from "@/components/StatusCard.vue";
import dayjs from "dayjs";
import BtnProcessStopResume from "@/components/BtnProcessStopResume.vue";

export default defineComponent({
  name: "Processes",
  components: {
    BtnProcessStopResume,
    StatusCard,
    CheckBox,
    CreateFlow
  },
  props: {},
  data() {
    return {
      list: [] as Process[],
      showCreateFlowDialog: false,
      selected: new Set<string>(),
    }
  },
  computed: {
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
    viewRoute(id: string) {
      this.$router.push({name: "ViewProcess", params: {id: id}})
    },
    async UpdateList() {
      const res = await processService.processServiceListProcess()
      this.list = res.processes
      this.selected = new Set<string>()
    },
    async Cancel(id: string) {
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


<style>


</style>

