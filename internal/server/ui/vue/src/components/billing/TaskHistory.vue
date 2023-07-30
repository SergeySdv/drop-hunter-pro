<template>
  <div class="text-h5">Task execution history</div>
  <div class="text-center" v-if="items.length === 0">history is empty</div>
  <v-list v-else>
    <v-list-item
      density="compact"
      variant="outlined"
      class="my-1"
      v-for="item in items"
      :key="item.taskId"
      :title="item.taskType"
    >
      <div class="d-flex justify-space-between">
        <span>{{ `Price: ${item.taskPrice} USD` }}</span>
        <a :href="`/process/${item.processId}`" target="_blank">show process</a>
      </div>

    </v-list-item>
  </v-list>
  <v-btn v-if="showNext" @click="next" width="100%">More</v-btn>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";
import {TaskType} from "@/generated/flow";
import {helperService} from "@/generated/services";
import {TaskHistoryRecord} from "@/generated/helper";

export default defineComponent({
  name: "TaskHistory",
  created() {
    this.getHistory()
  },
  computed: {
    showNext(): boolean {
      if (this.items.length === 0) {
        return false
      }
      return this.items.length % 10 === 0
    }
  },
  data() {
    return {
      loading: false,
      offset: 0,
      items: [] as TaskHistoryRecord[]
    }
  },
  methods: {
    next() {
      this.offset += 10
      this.getHistory()
    },
    async getHistory() {
      try {
        this.loading = true
        const res = await helperService.helperServiceGetBillingHistory({body: {offset: this.offset.toString()}})
        this.items.push(...res.records)
      } catch (e) {
      } finally {
        this.loading = false
      }
    }
  }

})


</script>


<style>
</style>

