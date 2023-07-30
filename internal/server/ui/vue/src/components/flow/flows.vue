<template>
  <div>
    <Loader v-if="loading"/>
    <div v-else>
      <div class="d-flex justify-end">
        <v-btn @click=CreateFlow class="ma-3">Create</v-btn>
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
          active-color="black"
          @click="viewFlow(item.id)"
        >
          <div>
            <div class="mr-2"><b>{{ `name: ${item.label}` }}</b></div>
            <div><b>Flow:</b>
              <div class="mr-2" v-for="(d, i) in getFlow(item)">
                <b>{{ i + 1 }})</b> {{ d }}
              </div>
            </div>
          </div>
        </v-list-item>
      </v-list>
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService} from "@/generated/services"
import {flow_Flow as Flow, TaskType} from "@/generated/flow";
import CreateFlow from "@/components/flow/CreateFlow.vue";
import CheckBox from "@/components/CheckBox.vue";
import {getFlow} from "@/components/tasks/tasks";
import Loader from "@/components/Loader.vue";

export default defineComponent({
  name: "Constructor",
  components: {
    Loader,
    CheckBox,
    CreateFlow
  },
  props: {},
  data() {
    return {
      list: [] as Flow[],
      loading: true,
      loadingError: false,
    }
  },
  computed: {
    getList(): Flow[] {
      return this.list
    },
  },
  methods: {
    getFlow,
    viewFlow(id: string) {
      window.open("/flow/" + id, "_blank")
    },
    showStep(flow: Flow): string {
      if (!flow || !flow.tasks) {
        return "-"
      }
      let res = ""
      flow.tasks.forEach((task) => {
        res += ` ${task.weight}) ${task.taskType}`
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
      try {
        this.loadingError = false
        this.loading = true
        const res = await flowService.flowServiceListFlow()
        this.list = res.flows
        this.selected = new Set<string>()
      } catch (e) {
        this.loadingError = true
      } finally {
        this.loading = false
      }

    },
    CreateFlow() {
      this.$router.push({name: 'CreateFlow'})
    },
  },
  created() {
    this.UpdateList()
  }
})


</script>


<style>


</style>

