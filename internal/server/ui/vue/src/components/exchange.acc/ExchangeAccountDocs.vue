<template>
  <div>
    <v-dialog
      v-model="dialog"
      width="800"
    >
      <template v-slot:activator="{ props }">
        <v-btn
          color="primary"
          v-bind="props"
        >
          Docs
          <v-icon icon="mdi-information-outline"/>
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="text-h5">What is exchange account?</span>
        </v-card-title>
        <v-card-text>
          <b>Exchange account</b> is a account of binance/okex/bybit decentralized crypto exchange (DEX).
          <br/>
          <br/>
          <b>Purpose</b> - it is needed to make withdraw and deposit easier. You can make these processes completely
          automatic.
          <br/>
          <br/>
          <b>General usage</b> withdraw from binance -> do work -> deposit okex sub-account wallet -> transfer from okex
          sub-account to okex main account ->
          transfer from okex -> binance. Repeat as many times as you want.
          <br/>
          <br/>
          Supported DEX:
          <ul>
            <li><b>Binance</b>
              - withdraw only
              - <a href="/docs/binance_account_instruction" target="_blank">instruction</a>
            </li>
            <li><b>Okex</b>
              - withdraw/deposit (sub-accounts support) <a href="/docs/okex_account_instruction" target="_blank">instruction</a>
            </li>
          </ul>

          <br/>
          <b>Since almost all DEX allows to use api for withdraw for a certain IP address it is necessary to set socks5
            proxy with static IP for each account.</b>

          <br/>
          <br/>


        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="green-darken-1"
            variant="text"
            @click="dialog = false"
          >
            OK
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';

export default defineComponent({
  name: "ExchangeAccountDocs",
  data() {
    return {
      dialog: false,
    }
  },
})
</script>

<style>
</style>

