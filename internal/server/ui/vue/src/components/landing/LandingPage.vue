<template>

  <div class="text-center text-h5 text-blue-darken-2 my-5">
    <h2
      class="">
      Automate your airdrop hunting strategy
    </h2>
  </div>

  <v-card
    theme=""
    flat
    rounded="0"
  >

    <v-img src="/icons/process.png"/>

    <div class="d-flex justify-center">
      <v-btn height="60px" width="200px" style="font-size: 30px" @click="$router.push({name:'Modules'})">Modules
      </v-btn>
    </div>

  </v-card>

  <div class="mx-5 my-8">
    <div class="d-flex flex-wrap justify-space-between mr-5">

      <div v-for="item in items">

        <div style="width: 300px" class="mx-5">
          <div class="mr-4">
            <div class="">

            </div>
          </div>
          <div><h6 class="text-h5 text-blue">
            {{ item.title }}
          </h6>
            <p class="my-3">
              {{ item.details }}
            </p></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';

export default defineComponent({
  name: "LandingPage",
  methods: {},
  data() {
    return {
      onboarding: 0,
      items: [
        {
          title: 'Security',
          details: 'All sensitive data such as: private keys and secret keys are stored in encrypted format in database.\n' +
            '            They do not appear anywhere after creation. Users can not see any sensitive information about other users.'
        },
        {
          title: 'Friendly user interface',
          details: ' No special skills needed. You can manage multiple accounts in one place'
        },
        {
          title: 'Full support',
          details: 'Our support team will help you with any questions'
        },
        {
          title: 'Create and run you custom flow',
          details: ' You can build your own flow (example: withdraw -> swap -> bridge -> deposit -> repeat)'
        },
        {
          title: 'Scalability',
          details: 'Can support large amount of flows and accounts'
        },
        {
          title: 'Withdrawal',
          details: 'Withdraw any coins on any network is available from Binance and Okex.'
        },
        {
          title: 'Deposit okex',
          details: ' Use advantage of okex sub-accounts to make deposits to lower risks'
        },
        {
          title: ' Supports many blockchains',
          details: 'Etherium, Binance BNB, Optimism, Arbitrum, Polygon, ZkSync, Avalanche'
        },
        {
          title: 'Snapshot.org voting',
          details: 'Easily subscribe and vote any topic you want'
        }
      ]
    }
  }
})


</script>


<style>
</style>

