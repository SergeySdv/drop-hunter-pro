<template>
  <v-container>
    <div class="d-flex justify-space-between">
      <div>
        <div class="my-1 text-h4"> User: {{ store.email }}</div>
        <div class="d-flex flex-row my-1 text-h5 ">
          <div class="mr-5">Funds: <b>{{ store.ass }}</b></div>
          <TopUp/>
        </div>


        <div class="my-1 text-h5">Task execution price: <b>{{ store.taskPrice }}</b></div>
      </div>
      <div>
        <b> Payable Tasks</b>
        <br/>
        <ul>
          <li v-for="task in store.payableTasks">
            {{ task }}
          </li>
        </ul>
        <b> Non-payable Tasks</b>
        <br/>
        <ul>
          <li v-for="task in store.nonpayableTasks">
            {{ task }}
          </li>
        </ul>
      </div>
    </div>
    <TaskHistory/>
    <PaymentHistory/>
  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";
import TaskHistory from "@/components/billing/TaskHistory.vue";
import TopUp from "@/components/billing/TopUp.vue";
import PaymentHistory from "@/components/billing/PaymentHistory.vue";

export default defineComponent({
  name: "Billing",
  components: {PaymentHistory, TopUp, TaskHistory},
  created() {
  },
  data() {
    return {}
  },
  computed: {
    ...mapStores(useUserStore),
    store() {
      return this.userStore
    }
  }
})


</script>


<style>
</style>

