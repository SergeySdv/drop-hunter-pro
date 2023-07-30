<template>
  <v-btn :loading="loading" v-if="processCanBeStopped() && !errorHappen()" :density="size"
         @click="Stop(item.id)"
         color="orange">Stop
  </v-btn>
  <v-btn v-else-if="item.status !== ProcessStatus.StatusDone && !errorHappen()" :loading="loading"
         :density="size"
         @click="Resume(item.id)" color="green">
    {{ processNotStarted() ? "Start" : "Resume" }}
  </v-btn>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {processService} from "@/generated/services"
import {Process, ProcessStatus} from "@/generated/process";
import {Delay} from "@/components/helper";

export default defineComponent({
  name: "BtnProcessStopResume",
  props: {
    item: {
      type: Object as PropType<Process>,
      required: true,
    },
    size: {
      type: String as PropType<'compact' | 'comfortable'>,
      required: true,
    },
  },

  data() {
    return {
      loading: false
    }
  },
  computed: {
    ProcessStatus() {
      return ProcessStatus
    },
  },
  methods: {
    errorHappen(): boolean {
      if (this.item.status === ProcessStatus.StatusError) {
        return true
      }
      return false
    },
    processCanBeStopped(): boolean {

      if (this.item.status === ProcessStatus.StatusStop) {
        return false
      }

      if (this.item.status === ProcessStatus.StatusDone) {
        return false
      }
      return true
    },

    processNotStarted(): boolean {
      if (this.item.status !== ProcessStatus.StatusStop) {
        return false
      }

      return !this.item.profiles.some(p => p.status !== ProcessStatus.StatusReady)
    },

    async Stop(id: string) {
      try {
        this.loading = true
        await processService.processServiceStopProcess({body: {processId: id}})
        this.$emit("updated")
      } finally {
        this.loading = false
      }
    },
    async Resume(id: string) {
      try {
        this.loading = true
        await processService.processServiceResumeProcess({body: {processId: id}})
        this.$emit("updated")
      } finally {
        this.loading = false
      }
    },
    async mounted() {
    },
  }
})


</script>


<style>


</style>

