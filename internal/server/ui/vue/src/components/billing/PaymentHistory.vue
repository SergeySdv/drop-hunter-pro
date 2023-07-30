<template>
  <div class="text-h5">Payment history</div>
  <div class="text-center" v-if="items.length === 0">history is empty</div>
  <v-list v-else>
    <v-list-item
      density="compact"
      variant="outlined"
      class="my-1"
      v-for="item in items"
      :key="item.id"
      :title="item.status"
    >
      <div class="d-flex justify-space-between">
        <span>{{ `amount: ${item.am} USD` }}</span>
        <span>Net: {{ item.net }} to: {{ item.toWallet }} status: {{ item.status }}</span>
      </div>

    </v-list-item>
  </v-list>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";
import {TaskType} from "@/generated/flow";
import {helperService} from "@/generated/services";
import {Order, TaskHistoryRecord} from "@/generated/helper";

export default defineComponent({
  name: "PaymentHistory",
  created() {
    this.getHistory()
  },
  computed: {},
  data() {
    return {
      loading: false,
      items: [] as Order[]
    }
  },
  methods: {
    async getHistory() {
      try {
        this.loading = true
        const res = await helperService.helperServiceGetOrderHistory()
        this.items.push(...res.orders)
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

