<template>

  <v-card class="text-center">
    <v-card-title>
      <span class="text-h5">How to connect okex account?</span>
    </v-card-title>
    <v-card-text>
      <ul>
        <li><h2 style="color: blue">1) go to <a href="https://www.okx.com/account/my-api"
                                                target="_blank">https://www.okx.com/account/my-api</a></h2></li>
        <br/>
        <br/>
        <li><h2 style="color: blue">2) Switch to main account</h2></li>
        <br/>
        <br/>
        <img src="/onboarding/okx/step1.png">

        <br/>
        <br/>
        <li><h2 style="color: blue">3) Click user icon -> API Keys
        </h2></li>
        <br/>
        <br/>
        <img src="/onboarding/okx/step2.png">
        <br/>
        <br/>
        <li><h2 style="color: blue">4) Click Create API key
        </h2></li>
        <br/>
        <br/>
        <img src="/onboarding/okx/step33.png">

        <br/>
        <br/>
        <li><h2 style="color: blue">5) fill up the form
        </h2></li>
        <br/>
        <br/>
        <img src="/onboarding/okx/step3.png">
        <br/>
        <br/>
        <li>
          <h2 style="color: blue">6) Copy Secret and API keys
          </h2>
        </li>
        <br/>
        <br/>
        <img src="/onboarding/okx/step4.png">
        <br/>
        <br/>

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
  name: "ExchangeAccountOkexDocs",
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

