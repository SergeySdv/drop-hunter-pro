<template>
  <v-slider
    v-model="multiplier"
    density="compact"
    label="max gas"
    :min="getMin"
    step="1000"
    :max="getMax"
    persistent-hint
    thumb-label="always"
    hide-details
  >
    <template v-slot:append>

      <v-tooltip text="Tooltip">
        <template v-slot:activator="{ props }">
          <v-icon v-bind="props" icon="mdi-information-outline"></v-icon>
        </template>

      </v-tooltip>


    </template>
    <template v-slot:thumb-label="{ modelValue }">
      <div style="width: 80px" class="text-center">{{ Number(getPriceUSD).toFixed(2) }} USD</div>

    </template>
  </v-slider>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {Network} from "@/generated/flow";
import {Timer} from "@/components/helper";
import {helperService} from "@/generated/services";

export default defineComponent({
  name: "NetworkMaxGas",
  props: {
    modelValue: {
      type: String,
      required: false,
    },
    network: {
      type: String as PropType<Network>,
      required: true
    }
  },
  emits: ['update:modelValue'],
  data() {
    return {
      priceUSD: "-1",
      timer: new Timer()
    }
  },
  computed: {
    getMin(): string {
      if (this.network == Network.BinanaceBNB) {
        return "100000000000000"
      }

      if (this.network == Network.POLIGON) {
        return "10000000000000000"
      }

      if (this.network == Network.AVALANCHE) {
        return "1000000000000000"
      }

      return "10000000000000"
    },
    getMax(): string {
      if (this.network == Network.BinanaceBNB) {
        return "100000000000000000"
      }
      if (this.network == Network.POLIGON) {
        return "40000000000000000000"
      }

      if (this.network == Network.AVALANCHE) {
        return "1000000000000000000"
      }

      return "10000000000000000"
    },
    getPriceUSD() {
      return this.priceUSD
    },
    Network() {
      return Network
    },
    multiplier: {
      get(): string {
        return this.modelValue ? this.modelValue : "10000000000000"
      },
      async set(value: any) {
        this.loadPriceUSD()
        this.$emit('update:modelValue', String(value))
      }
    },
  },
  methods: {
    async loadPriceUSD() {

      this.timer.add(100)
      this.timer.cb(async () => {
        const res = await helperService.helperServiceCastWei({
          body: {
            wei: this.multiplier.toString(),
            network: this.network
          }
        })
        if (res.am.usd) {
          this.priceUSD = res.am.usd
        }
      })
    },
  },
  created() {
    this.loadPriceUSD()
  }
})


</script>


<style>


</style>

