<template>
  <v-row>
    <v-col>
      <v-radio-group
        density="comfortable"
        direction="horizontal"
        inline="true"
        v-model="kind"
        @change="amountKindChanged"
        :label="`Choose amount of ${coin} to send:`"
        :disabled="disabled"
        hide-details
      >
        <v-radio label="All" value="all"></v-radio>
        <!--        <v-radio label="Value" value="value"></v-radio>-->
        <v-radio label="Percent" value="percent"></v-radio>
      </v-radio-group>
    </v-col>
  </v-row>
  <v-row>
    <v-col v-if="kind === 'percent'">
      <v-text-field
        type="number"
        v-model="amount.sendPercent"
        variant="outlined"
        density="compact"
        :label="`percent of ${coin} to use`"
        :rules="[percent]"
        :disabled="disabled"
        hide-details
      />
    </v-col>
    <v-col v-if="kind === 'value'">
      <v-text-field
        type="number"
        variant="outlined"
        density="compact"
        :label="`value of ${coin} to use`"
        v-model="amount.sendValue"
        :rules="[required]"
        :disabled="disabled"
        hide-details
      />
    </v-col>
  </v-row>
</template>

<script lang="ts">
import {Amount, AmUni} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {processService} from "@/generated/services";
import {EstimationTx} from "@/generated/process";
import {percent} from "@/components/helper";
import {ca} from "vuetify/locale";

export default defineComponent({
  name: "AmountInput",
  emits: ['update:modelValue'],
  props: {
    coin: {
      type: String,
      required: true,
    },
    disabled: {
      type: Boolean,
      required: false,
    },
    modelValue: {
      type: Object as PropType<Amount>,
      required: false
    }
  },
  watch: {
    modelValue: {
      immediate: true,
      handler() {
        if (!this.modelValue) {
          return
        }

        switch (true) {
          case this.modelValue.sendValue !== undefined:
            this.kind = 'value';
            break;
          case this.modelValue.sendPercent !== undefined:
            this.kind = 'percent';
            break;
          case this.modelValue.sendAll !== undefined:
            this.kind = 'all';
            break;
        }

        this.amount = this.modelValue
      },
      deep: true
    },
    amount: {
      async handler(a, b) {
        this.$emit('update:modelValue', this.amount)
      },
      deep: true
    }
  },
  data() {
    return {
      kind: 'all' as 'all' | 'value' | 'percent',
      amount: {} as Amount
    }
  },
  methods: {
    percent,
    required: (v: any) => !!v || 'required',
    amountKindChanged() {
      switch (this.kind) {
        case 'all':
          this.amount.sendAll = true
          this.amount.sendValue = undefined
          this.amount.sendPercent = undefined
          this.amount.sendAmount = undefined
          break
        case "percent":
          this.amount.sendAll = undefined
          this.amount.sendValue = undefined
          this.amount.sendAmount = undefined
          break
        case "value":
          this.amount.sendAll = undefined
          this.amount.sendPercent = undefined
          this.amount.sendAmount = undefined
      }
    },
  }
})
</script>

<style scoped>

</style>
