<template>
  <div>
    <div class="d-flex justify-end">
      <v-btn @click=Update :loading="updating" class="ma-3">Update settings</v-btn>
      <v-btn @click=Reset :loading="reseting" class="ma-3">Reset settings</v-btn>
    </div>
    <v-form ref="settings-form">
      <SettingsNetwork v-if="!loading" v-model="settings.arbitrum" :network="Network.ARBITRUM"/>
      <SettingsNetwork v-if="!loading" v-model="settings.optimism" :network="Network.OPTIMISM"/>
      <SettingsNetwork v-if="!loading" v-model="settings.etherium" :network="Network.Etherium"/>
      <SettingsNetwork v-if="!loading" v-model="settings.bnb" :network="Network.BinanaceBNB"/>
      <!--      <SettingsNetwork v-if="!loading" v-model="settings.polygon" :network="Network.POLIGON"/>-->
      <SettingsNetwork v-if="!loading" v-model="settings.avalanche" :network="Network.AVALANCHE"/>
    </v-form>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {settingsService} from "@/generated/services"
import {Settings} from "@/generated/settings";
import {Network} from "@/generated/flow";
import SettingsNetwork from "@/components/SettingsNetwork.vue";

export default defineComponent({
  name: "Settings",
  components: {SettingsNetwork},
  data() {
    return {
      loading: false,
      settings: {} as Settings,
      error: false,
      updating: false,
      reseting: false
    }
  },
  computed: {
    Network() {
      return Network
    }
  },
  methods: {
    async loadSettings() {
      try {
        this.error = false
        this.loading = true
        const res = await settingsService.settingsServiceGetSettings({
          body: {}
        })
        this.settings = res.settings
      } catch (e) {
        this.error = true
      } finally {
        this.loading = false
      }
    },
    async Reset() {
      this.reseting = true
      try {
        const res = await settingsService.settingsServiceReset()
        this.settings = res.settings
      } finally {
        this.reseting = false
      }
    },
    async Update() {
      this.updating = true
      if (!await this.validate()) {
        this.updating = false
      }

      try {
        const res = await settingsService.settingsServiceUpdateSettings({body: {settings: this.settings}})
        this.settings = res.settings
      } finally {
        this.updating = false
      }
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['settings-form'].validate()

      return valid
    },
  },
  async created() {
    await this.loadSettings()
  }
})


</script>


<style>


</style>

