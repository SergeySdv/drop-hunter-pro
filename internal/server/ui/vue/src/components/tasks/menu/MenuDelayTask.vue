<template>
  <div v-if="item.random">{{ `Settings: [${humanDuration(item.minRandom)}:${humanDuration(item.maxRandom)}]` }}</div>
  <div v-if="item.duration">Duration: <b>{{
      humanDuration(item.duration)
    }}</b></div>
  <div v-if="item.waitFor">Wait for: <b>{{ formatTime(item.waitFor) }}</b></div>
</template>

<script lang="ts">
import {AmUni, DelayTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {humanDuration, formatTime} from "@/components/helper";

export default defineComponent({
  name: "MenuDelayTask",
  props: {
    task: {
      type: Object as PropType<Task>,
      required: true,
    }
  },
  watch: {
    item: {
      handler() {
        if (this.task?.delayTask) {
          this.item = this.task.delayTask
        }
      },
      deep: true
    }
  },
  data() {
    return {
      item: {} as DelayTask,
    }
  },
  computed: {},
  methods: {humanDuration, formatTime},
  async created() {
    if (this.task?.delayTask) {
      this.item = this.task.delayTask
    }
  }
})
</script>

<style scoped>

</style>
