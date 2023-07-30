<template>
  <v-dialog
    v-model="dialog"
    width="800"
    :close-on-content-click="false"
    :close-on-back="false"
    persistent="true"
  >
    <template v-slot:activator="{ props }">
      <v-btn density="comfortable" @click="dialog=true">TopUp balance</v-btn>
    </template>

    <v-card>
      <v-card-title>Balance TopUp</v-card-title>
      <v-card-text>
        <div class="d-flex justify-space-between my-4">
          <v-text-field
            class="mr-3"
            type="number"
            v-model="amInput"
            density="compact"
            variant="outlined"
            placeholder="enter amount in USDT"
            :disabled="amInputDisabled"
            :rules="[onlyInteger,required]"
            suffix="USDT"
            hide-details
          />
          <v-btn v-if="!amInputDisabled" rounded="false" @click="amountFilled">Next</v-btn>
        </div>
        <div v-if="amInputDisabled">
          <span>Send exact amount of <b>{{ amToSend }} <a :href="coinAddrUrl"
                                                          target="_blank">USDT</a></b>
            <br/>
            to address: <b>{{
                walletToSend
              }}</b> using Arbitrum Network</span>
          <div class="text-center my-10">
            <div v-if="finished" style="color: green" class="text-h6 font-weight-bold">Payment confirmed</div>
            <div v-else>
              Waiting for transaction confirmation
              <v-progress-linear
                color="deep-purple-accent-4"
                indeterminate
                rounded
                height="6"
              ></v-progress-linear>
            </div>
          </div>
        </div>

      </v-card-text>
      <v-card-actions class="d-flex justify-end">
        <v-btn v-if="finished" @click="cancel">Close</v-btn>
        <v-btn v-else @click="cancel">Cancel</v-btn>
      </v-card-actions>
    </v-card>

  </v-dialog>
</template>


<script lang="ts">

import {defineComponent} from 'vue';
import {helperService} from "@/generated/services";
import {onlyInteger, required} from "@/components/tasks/menu/helper";

export default defineComponent({
  name: "TopUp",
  created() {
  },
  data() {
    return {
      amInput: "",
      dialog: false,
      amInputDisabled: false,
      orderCreating: false,

      coinAddrUrl: "",
      orderId: "",
      walletToSend: "",
      amToSend: 0,
      errorMsg: "",
      finished: false,
      pollerTimer: null as NodeJS.Timer | null
    }
  },
  methods: {
    required,
    onlyInteger,
    async amountFilled() {
      try {
        this.orderCreating = true
        const res = await helperService.helperServiceCreateOrder({body: {am: this.amInput}})
        this.amInputDisabled = true
        this.orderId = res.id
        this.walletToSend = res.toWallet
        this.coinAddrUrl = res.coinAddrUrl
        this.amToSend = res.am
        this.startPolling()
      } finally {
        this.orderCreating = false
      }
    },
    cancel() {
      this.stopPolling()
      this.amInput = ""
      this.amInputDisabled = false
      this.orderCreating = false
      this.errorMsg = ""
      this.finished = false
      this.coinAddrUrl = ''
      this.dialog = false
      this.pollerTimer = null
    },
    startPolling() {
      this.pollerTimer = setInterval(() => {
        this.checkOrder()
      }, 1000)

    },
    stopPolling() {
      if (this.pollerTimer) {
        clearInterval(this.pollerTimer)
      }
    },
    async checkOrder() {
      try {
        const res = await helperService.helperServiceGetOrderStatus({body: {orderId: this.orderId}})
        if (res.status === 'Error') {
          this.errorMsg = 'error happened'
          this.stopPolling()

        } else if (res.status === 'Processed') {
          this.finished = true
          this.stopPolling()
        }
      } finally {

      }

    }
  }

})


</script>


<style>
</style>

