<template>
  <v-menu
    v-model="menu"
    :close-on-content-click="false"
    location="end"
  >

    <template v-slot:activator="{ props }">
      <v-btn variant="outlined" @click="load" density="comfortable" v-bind="props">{{ label }}
      </v-btn>
    </template>

    <template v-slot:default="{ props }">
      <v-card class="px-4 py-4">

        <div class="d-flex justify-space-between">
          <div>
            <a class="pr-5" :href="`https://debank.com/profile/${profile.mmskId}`" target="_blank">Debank</a>
            <TopUpProfile :profile-id="profileId"/>
          </div>
          <v-icon icon="mdi-close" @click="menu = false"/>
        </div>


        <div class="d-flex justify-space-between">
          <div>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.ARBITRUM"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.Etherium"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.OPTIMISM"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.BinanaceBNB"/>
            </v-card-item>
          </div>
          <div>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.POLIGON"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.AVALANCHE"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.ZKSYNCLITE"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.ZKSYNCERA"/>
            </v-card-item>
          </div>
        </div>
        <div>
          <div v-if="transactions.length > 0">
            Recent Transactions:
            <div v-for="(tx, i) in transactions">
              <div>
                <a class="mr-1" :href="tx.url" target="_blank">TX</a>
                <b> {{ tx.code }}</b>
                {{ formatTime(tx.createdAt) }}
                <a :href="`/process/${tx.processId}`" target="_blank">Process</a>
              </div>
            </div>
          </div>
        </div>
      </v-card>
    </template>
  </v-menu>

</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {processService, profileService} from "@/generated/services"
import Balance from "@/components/profile/Balance.vue";
import {Network, Profile} from "@/generated/profile";
import TopUp from "@/components/billing/TopUp.vue";
import Withdraw from "@/components/exchange.acc/ExchangeAccounts.vue";
import TopUpProfile from "@/components/exchange.acc/TopUpProfile.vue";
import {Transaction} from "@/generated/process";
import {formatTime} from "../helper";

export default defineComponent({
  name: "ProfileCard",
  components: {TopUpProfile, Withdraw, TopUp, Balance},
  props: {
    profileId: {
      type: String,
      required: true
    },
    label: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      profile: {} as Profile,
      menu: false,
      transactionsLoading: false,
      transactions: [] as Transaction[]
    }
  },
  computed: {
    Network() {
      return Network
    }
  },
  methods: {
    async load() {
      this.loadTransactions()
      this.loadProfile()
    },
    formatTime,
    async loadTransactions() {
      try {
        this.transactionsLoading = true
        const res = await processService.processServiceGetProfileTransactions({body: {profileId: this.profileId}})
        this.transactions = res.transactions
      } finally {
        this.transactionsLoading = false
      }
    },
    async loadProfile() {
      const res = await profileService.profileServiceGetProfile({body: {profileId: this.profileId}})
      this.profile = res.profile
    }
  },
})


</script>


<style>


</style>

