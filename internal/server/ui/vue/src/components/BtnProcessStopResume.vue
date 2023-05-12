<template>
  <v-btn :loading="loading" class="mr-2" v-if="processCanBeStopped()" :density="size"
         @click="Stop(item.id)"
         color="orange">Stop
  </v-btn>
  <v-btn v-else-if="item.status !== ProcessStatus.StatusDone" :loading="loading" class="mr-2" :density="size"
         @click="Resume(item.id)" color="green">
    {{ processNotStarted() ? "start" : "resume" }}
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
        await Delay(1000)
        this.loading = false
      }
    },
    async Resume(id: string) {
      try {
        this.loading = true
        await processService.processServiceResumeProcess({body: {processId: id}})
        this.$emit("updated")
      } finally {
        await Delay(1000)
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

