<template>
  <v-menu
    v-model="showTooltip"
    width="350px"
    persistent
    :close-on-back="false"
    :close-on-content-click="false"
  >
    <template v-slot:activator="{ props }">
      <v-btn
        class="ml-1"
        :rounded="false"
        density="compact"
        v-bind="props"
      >
        Estimate
      </v-btn>
    </template>

    <v-skeleton-loader
      :loading="loading"
      type="table"
      min-height="200px"
      elevation="10"
    >
      <v-responsive>

        <v-card>
          <v-card-title class="d-flex justify-space-between">
            <div>Market Estimation</div>
            <v-icon icon="mdi-close" @click="showTooltip = false"/>
          </v-card-title>
          <v-card-text style="white-space: pre-wrap">
            <div v-if="errorMessage" style="color: red">{{ errorMessage }}</div>
            <div v-else class="d-flex justify-space-between">
              <div>
                <!--                <div><b> balance</b>: {{ getBalance(estimation.balance, false) }}</div>-->
                <!--                <div><b>value</b>: {{ getBalance(estimation.value, false) }}</div>-->
                <div><b>gas</b>: {{ getBalance(estimation.gas, true) }}</div>
              </div>
              <div>
                <div><b>gas limit</b>: {{ getBalance(estimation.gasLimit, false) }}</div>
                <div><b>gas price</b>: {{ getBalance(estimation.gasPrice, false) }}</div>
                <!--                <div><b>gas/value: </b>: {{ `${Number(estimation.gasValuePercent).toPrecision(2)}%` }}</div>-->
              </div>
            </div>
          </v-card-text>
        </v-card>
      </v-responsive>
    </v-skeleton-loader>
  </v-menu>
</template>

<script lang="ts">
import {AmUni, Network, Token} from "@/generated/flow";
import {defineComponent} from "vue";
import {processService} from "@/generated/services";
import {EstimationTx} from "@/generated/process";

export default defineComponent({
  name: "EstimateTask",
  props: {
    taskId: {
      type: String,
      required: true,
    },
    profileId: {
      type: String,
      required: true,
    },
    processId: {
      type: String,
      required: true,
    },
  },
  watch: {
    showTooltip: {
      handler() {
        if (this.showTooltip) {
          this.estimate()
        }
      }
    }
  },
  data() {
    return {
      errorMessage: '',
      loading: false,
      showTooltip: false,
      estimation: {} as EstimationTx
    }
  },
  methods: {
    getBalance(am?: AmUni, gas?: boolean): string {
      if (!am) {
        return ''
      }

      if (gas) {
        let token = Token.ETH
        switch (am.network) {
          case Network.BinanaceBNB:
            token = Token.BNB
            break
          case Network.POLIGON:
            token = Token.MATIC
            break
          case  Network.AVALANCHE:
            token = Token.AVAX
            break
        }
        return `\n ${token}: ${Number(am.eth).toPrecision(3)}\n  USD: ${Number(am.usd)} \n GWEI: ${Number(am.gwei)}`
      }
      return `\n GWEI: ${Number(am.gwei)}`
    },
    async estimate() {
      try {
        this.loading = true
        this.errorMessage = ''
        const res = await processService.processServiceEstimateCost({
          body: {
            profileId: this.profileId,
            taskId: this.taskId,
            processId: this.processId,
          }
        })
        if (res.error) {
          this.errorMessage = res.error
        } else if (res.data) {
          this.estimation = res.data
        }
      } catch (e) {
        this.errorMessage = 'error estimating'
      } finally {
        this.loading = false
      }
    }
  },
})
</script>

<style scoped>

</style>
