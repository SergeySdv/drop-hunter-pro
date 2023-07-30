<template>
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
        <span class="text-h5">What is profile?</span>
      </v-card-title>
      <v-card-text>
        <b>Profile</b> is a custom entity that allows user to connect metamask wallet.
        <br/>
        <br/>
        <b>Purpose</b> - it is needed to make process
        creation easier
        <br/>
        <br/>

        <ul>
          <li><b>Label</b> - (custom profile name, it is better to name it with digits for easier usage)</li>
          <li><b>Metamask private key (PK)</b> - PK is stored in encrypted format and it is used only during runtime.
            It does not shown up in logs or internal net traffic. So it is protected. No-one can change/see it after
            profile creation
          </li>
          <li><b>Proxy</b> - optional socks5 proxy, if set - all request embraced with profile will be done through
            that proxy
          </li>
          <li><b>Meta</b> - you can write your notes here</li>

          If everything is set correct - you can easily check metamask balance.
        </ul>

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
</template>

<script lang="ts">

import {defineComponent} from 'vue';

export default defineComponent({
  name: "ProfilesDocs",
  data() {
    return {
      dialog: false,
    }
  },
})
</script>

<style>
</style>

