<template>
  <v-dialog
    v-model="dialog"
    fullscreen="true"
  >
    <template v-slot:activator="{ props }">
      <v-btn
        color="primary"
        v-bind="props"
      >
        Export
      </v-btn>
    </template>
    <v-card>
      <v-card-title class="d-flex justify-space-between">
        <div>
          <span class="text-h5 mr-3">Export exchange accounts</span>
        </div>
        <div>
          <v-icon icon="mdi-close" @click="dialog = false"/>
        </div>

      </v-card-title>
      <v-card-text style="white-space: pre-wrap">
        {{ text }}
      </v-card-text>

    </v-card>
  </v-dialog>
</template>

<script lang="ts">

import {profileService, withdrawerService} from '@/generated/services';
import {defineComponent} from 'vue';

export default defineComponent({
  name: "ExportExchangeAccounts",
  watch: {
    dialog: {
      async handler() {
        if (this.dialog) {
          const res = await withdrawerService.withdrawerServiceExportExchangeAccounts()
          this.text = res.text
        }
      }
    }
  },
  data() {
    return {
      text: '',
      dialog: false,
    }
  },
  methods: {
    async copyClipBoard(text: string) {
      console.log(window.navigator.clipboard)
      await navigator.clipboard.writeText(text)
    }
  },
  async created() {

  }
})
</script>

<style>
</style>

