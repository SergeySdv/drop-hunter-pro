<template>
  <div>
    <div><b>Gas:</b>
      <v-dialog
        v-model="menu"
        width="600px"
        height="300px"
        :close-on-content-click="false"
        :close-on-back="false"
        persistent="true"
        location="center"
        location-strategy="static"
      >
        <template v-slot:activator="{ props }">
          <v-btn class="ml-3" variant="outlined" height="18px" width="100px" @click="openSettings">Settings</v-btn>
        </template>

        <v-card class="px-4 py-4">
          <v-card-title class="d-flex justify-space-between">
            <div>Gas settings</div>
            <v-icon icon="mdi-close" @click="menu = false"/>
          </v-card-title>
          <v-card-text>
            <v-row>
              <v-col>
                <NetworkGasMultiplier v-model="settings.gasMultiplier"/>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <NetworkMaxGas v-model="settings.maxGas" :network="network"/>
              </v-col>
            </v-row>
          </v-card-text>
          <v-card-actions class="d-flex justify-end">
            <v-btn v-if="changed" @click="updateSettings">Update</v-btn>
          </v-card-actions>

        </v-card>
      </v-dialog>
    </div>

    <div v-if="item" class="pl-2">
      <div v-if="item.gasLimit">
        <i>Max:</i>
        {{ Number(item.gasLimit?.usd).toFixed(2) }} USD
      </div>
      <div v-if="item.multiplier">
        <i>Multiplier:</i>
        {{ Number(item.multiplier).toFixed(2) }}
      </div>
      <div v-if="item.gasEstimated">
        <i>Estimated:</i>
        {{ Number(item.gasEstimated?.usd).toFixed(2) }} USD
      </div>
      <div v-if="item.gasResult">
        <i>Paid:</i>
        {{ Number(item?.gasResult?.usd).toFixed(2) }} USD
      </div>
    </div>
  </div>

</template>

<script lang="ts">
import {Amount, AmUni, Network, TaskTx, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import NetworkGasMultiplier from "@/components/settings/NetworkGasMultiplier.vue";
import NetworkMaxGas from "@/components/settings/NetworkMaxGas.vue";
import {helperService, settingsService} from "@/generated/services";
import {Settings, SettingsNetwork as Sn} from "@/generated/settings";
import SettingsNetwork from "@/components/settings/SettingsNetwork.vue";
import deepEqual from "deep-equal";
import {mapStores} from "pinia";
import {useSysStore} from "@/plugins/pinia";

export default defineComponent({
  name: "GasOptions",
  components: {NetworkMaxGas, NetworkGasMultiplier, SettingsNetwork},
  props: {
    network: {
      type: String as PropType<Network>,
      required: true,
    },
    item: {
      default: {} as TaskTx,
      type: Object as PropType<TaskTx>,
      required: true,
    },

  },
  watch: {
    settings: {
      handler() {
        if (!deepEqual(this.settings, this.origSettings)) {
          this.changed = true
        }
      },
      deep: true
    }
  },
  data() {
    return {
      menu: false,
      settings: {} as Sn,
      origSettings: {} as Sn,
      global: {} as Settings,
      changed: false,
    }
  },
  computed: {
    ...mapStores(useSysStore),
  },
  methods: {
    async updateSettings() {
      let settings = this.getNetworkSettings(this.global)
      settings = this.settings
      try {
        await settingsService.settingsServiceUpdateSettings({body: {settings: this.global}})
        this.sysStore.showSnackBar("settings updated", false)
        this.menu = false
        this.changed = false
      } catch (e) {
        this.sysStore.showSnackBar("settings update error", true)
      }

    },
    async openSettings() {
      const res = await settingsService.settingsServiceGetSettings()

      this.settings = this.getNetworkSettings(res.settings)
      this.global = res.settings
      this.origSettings = Object.assign({}, this.settings)
      this.menu = true
    },
    getNetworkSettings(s: Settings) {
      switch (this.network) {
        case Network.AVALANCHE:
          return s.avalanche
        case Network.BinanaceBNB:
          return s.bnb
        case Network.POLIGON:
          return s.polygon
        case Network.ZKSYNCLITE:
          return s.zksyncLite
        case Network.ZKSYNCERATESTNET:
          return s.zksyncTestNet
        case Network.ZKSYNCERA:
          return s.zksyncMainNet
        case Network.Etherium:
          return s.etherium
        case Network.ARBITRUM:
          return s.arbitrum
        case Network.OPTIMISM:
          return s.optimism
      }
    }
  },
})
</script>

<style scoped>

</style>
