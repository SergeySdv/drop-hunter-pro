<template>

  <v-card class="text-center">
    <v-card-title>
      <span class="text-h5">How to connect binance account?</span>
    </v-card-title>
    <v-card-text>
      <ul>
        <li><h2 style="color: blue">1) go to <a href="https://www.binance.com/"
                                                target="_blank">https://www.binance.com/</a>
        </h2></li>
        <br/>
        <br/>
        <img src="/onboarding/binance/binance_main_page_api.png">

        <br/>
        <br/>
        <li><h2 style="color: blue">2) Click user icon -> API management -> Create API
        </h2></li>
        <br/>
        <br/>
        <img src="/onboarding/binance/binance_create_api_key.png">
        <br/>
        <br/>
        <li><h2 style="color: blue">3) System generated -> Next
        </h2></li>
        <br/>
        <br/>
        <img src="/onboarding/binance/binance_create_api_gen.png">

        <br/>
        <br/>
        <li><h2 style="color: blue">4) Enter any label
        </h2></li>
        <br/>
        <br/>
        <img src="/onboarding/binance/binance_create_api_label.png">

        <br/>
        <br/>
        <li><h2 style="color: blue">5) Copy Secret and API keys
        </h2></li>
        <br/>
        <br/>
        <img src="/onboarding/binance/binance_create_api_secret.png">
        <br/>
        <br/>

        <li><h2 style="color: blue">6) Set restrictions
        </h2></li>
        <br/><br/>
        <img src="/onboarding/binance/binance_create_api_restrictions.png">
        <br/><br/>

        <li><h2 style="color: blue">7) Go <a href="/exchange_accounts" target="_blank">Exchange accounts sections</a> ->
          Add Exchange Account
        </h2></li>
        <br/><br/>
        <img src="/onboarding/binance/binance_add_account.png">
        <br/><br/>
      </ul>

    </v-card-text>
  </v-card>
</template>

<script lang="ts">

import {defineComponent} from 'vue';

export default defineComponent({
  name: "ExchangeAccountBinanceDocs",
  components: {},
  props: {},
  data() {
    return {
      dialog: false,
    }
  },
  computed: {},
  methods: {},
})
</script>

<style>
</style>

