<template>
  <div>
    <v-chip v-if="deposit.profileId" closable close-icon="mdi-close"
            @click:close="remove(deposit.profileId, deposit.addr)">
      {{ `${deposit.profileLabel}` }}
    </v-chip>
    <span v-else @click="searchProfiles">
    <v-autocomplete
      return-object
      v-model="selectedProfile"
      :loading="profilesLoading"
      :items="suggestedProfiles"
      item-title="num"
      item-value="id"
      label="select profiles"
      density="compact"
      no-data-text="no detached profiles"
    />
  </span>
  </div>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {helperService, profileService, withdrawerService} from "@/generated/services"
import {
  DepositAddresses,
} from "@/generated/withdraw";
import {Profile} from "@/generated/profile";
import {Timer} from "@/components/helper";

export default defineComponent({
  name: "OkexWithdrawProfiles",
  emits: ['update:modelValue', 'updated'],
  props: {
    modelValue: {
      required: true,
      type: Object as PropType<DepositAddresses>
    },
    withdrawerId: {
      required: true,
      type: String
    }
  },
  watch: {
    deposit: {
      handler() {
        this.$emit('update:modelValue', this.deposit)
      },
      deep: true
    },
    selectedProfile: {
      handler(p: Profile) {
        this.attach(p)
      }
    }
  },
  data() {
    return {
      deposit: {} as DepositAddresses,
      selectedProfile: null as Profile | null,
      suggestedProfile: "",
      suggestedProfiles: [] as Profile[],
      profilesLoading: false,
      timer: new Timer()
    }
  },
  methods: {
    async remove(profileId: string | undefined, addr: string) {
      try {
        await withdrawerService.withdrawerServiceOkexDepositAddrDetach({
          body: {
            profileId: profileId ? profileId : '',
            okexDepositAddr: addr,
            withdrawerId: this.withdrawerId
          }
        })
        this.deposit.profileId = undefined
        this.deposit.profileLabel = undefined
      } catch (e) {

      }
    },
    async attach(p: Profile) {
      if (p && p.id) {
        await withdrawerService.withdrawerServiceOkexDepositAddrAttach({
          body: {
            profileId: p.id,
            okexDepositAddr: this.deposit.addr,
            withdrawerId: this.withdrawerId
          }
        })

        this.deposit.profileId = p.id
        this.deposit.profileLabel = p.label

        this.$emit('updated')
      }
    },
    async searchProfiles() {
      this.timer.add(0)
      this.timer.cb(async () => {
        try {
          this.profilesLoading = true
          const res = await profileService.profileServiceSearchProfilesNotConnectedToOkexDeposit()

          this.suggestedProfiles = res.profiles
        } finally {
          this.profilesLoading = false
        }
      })
    },
  },
  async created() {
    Object.assign(this.deposit, this.modelValue)
    if (this.modelValue.profileId) {
      console.log('zuydewfe')
    }
  }
})


</script>


<style>


</style>

