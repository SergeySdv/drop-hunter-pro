<template>
  <v-select
    density="compact"
    variant="outlined"
    label="select exchange account"
    :rules="[required]"
    :items="withdrawers"
    item-value="id"
    item-title="label"
    v-model="withdrawerId"
    :loading="loading"
    :disabled="disabled"
  >
    <template v-slot:no-data>
      <div class="mx-2">No exchange account found. <a href="/exchange_accounts">Create one</a></div>
    </template>
  </v-select>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {withdrawerService} from "@/generated/services"
import {
  CreateOkexWithdrawerRequest,
  ExchangeType, ExchangeWithdrawNetwork,
  ExchangeWithdrawOptions,
  Withdrawer
} from "@/generated/withdraw";
import ProxyInput from "@/components/ProxyInput.vue";
import OkexWithdrawOptionSubAcc from "@/components/exchange.acc/OkexWithdrawOptionSubAcc.vue";
import CreateWithdrawerSubAcc from "@/components/exchange.acc/CreateWithdrawerSubAcc.vue";
import {required} from "@/components/tasks/menu/helper";
import ProfileSearch from "@/components/profile/ProfileSearch.vue";
import {Profile} from "@/generated/profile";
import {Delay, Timer} from "@/components/helper";

export default defineComponent({
  name: "WithdrawerSelect",
  emits: ['update:modelValue'],
  props: {
    disabled: {
      type: Boolean,
      required: false,
      default: false,
    },
    modelValue: {
      type: String,
      required: true
    }
  },
  computed: {
    withdrawerId: {
      get(): string {
        return this.modelValue ? this.modelValue : ""
      },
      async set(value: string) {
        this.$emit('update:modelValue', value)
      }
    }
  },
  components: {},
  data() {
    return {
      withdrawers: [] as Withdrawer[],
      loading: false,
    }
  },
  methods: {
    required,
    async GetWithdrawerList() {
      try {
        this.loading = true
        const res = await withdrawerService.withdrawerServiceListWithdrawer()
        this.withdrawers = res.withdrawers.map((w) => {
          w.label = `${w.label} [${w.exchangeType}]`
          return w
        })
      } finally {
        this.loading = false
      }
    }
  },
  async created() {
    try {
      this.loading = true
      await this.GetWithdrawerList()

    } finally {
      this.loading = false
    }

  }
})


</script>


<style>


</style>

