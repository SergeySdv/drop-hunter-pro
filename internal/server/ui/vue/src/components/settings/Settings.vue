<template>
  <v-container>
    <Loader v-if="loading"/>
    <div v-else>

      <div class="header">
        <div class="d-flex justify-end">
          <v-btn v-if="settingsChanged" @click=Update :loading="updating" class="ma-3">Update</v-btn>
          <v-btn @click=Reset :loading="reseting" class="ma-3">Reset</v-btn>
        </div>
      </div>

      <div style="margin-top: 90px" class="text-h4 text-center">Network settings</div>

      <v-form ref="settings-form">
        <SettingsNetwork v-if="!loading" v-model="settings.arbitrum" :network="Network.ARBITRUM"/>
        <SettingsNetwork v-if="!loading" v-model="settings.optimism" :network="Network.OPTIMISM"/>
        <SettingsNetwork v-if="!loading" v-model="settings.etherium" :network="Network.Etherium"/>
        <SettingsNetwork v-if="!loading" v-model="settings.bnb" :network="Network.BinanaceBNB"/>
        <SettingsNetwork v-if="!loading" v-model="settings.polygon" :network="Network.POLIGON"/>
        <SettingsNetwork v-if="!loading" v-model="settings.avalanche" :network="Network.AVALANCHE"/>
        <SettingsNetwork v-if="!loading" v-model="settings.zksyncMainNet" :network="Network.ZKSYNCERA"/>
        <SettingsNetwork v-if="!loading" v-model="settings.zksyncLite" :network="Network.ZKSYNCLITE"/>
      </v-form>
    </div>
  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {helperService, settingsService} from "@/generated/services"
import {Settings} from "@/generated/settings";
import {Network} from "@/generated/flow";
import SettingsNetwork from "@/components/settings/SettingsNetwork.vue";
import {GetUserResponse} from "@/generated/helper";
import {TaskType} from "@/generated/process";
import {castGasLimitMap} from "@/components/helper";
import WEIInputField from "@/components/WEIInputField.vue";
import {estimatedTaskMap} from "@/components/tasks/tasks";
import deepEqual from 'deep-equal';
import Loader from "@/components/Loader.vue";
import {mapActions, mapStores} from "pinia";
import {useSysStore, useUserStore} from "@/plugins/pinia";

interface TaskGasLimit {
  task: TaskType
  limitWei: number
}

export default defineComponent({
  name: "Settings",
  components: {Loader, WEIInputField, SettingsNetwork},
  data() {
    return {
      orgiSettings: {} as Settings,
      taskGasLimits: [] as TaskGasLimit[],
      loading: false,
      loadingError: false,
      settings: {} as Settings,
      updating: false,
      reseting: false,
      changed: false,
      user: {
        email: "unknown",
        id: "unknown"
      } as GetUserResponse
    }
  },
  watch: {
    settings: {
      handler() {
        this.changed = !deepEqual(this.settings, this.orgiSettings)
      },
      deep: true,
    }
  },
  computed: {
    ...mapStores(useSysStore),
    settingsChanged(): boolean {
      return this.changed
    },
    Network() {
      return Network
    },
    GetTelegramSubscriptionLink(): string {
      //t.me/drop_hunter_alert_bot
      return `https://t.me/drop_hunter_alert_bot`
    }
  },
  methods: {
    GetTaskGasLimits(): TaskGasLimit[] {
      const result: TaskGasLimit[] = []
      const m = castGasLimitMap(this.settings.taskGasLimitMap)

      const taskTypes = [] as TaskType[]
      estimatedTaskMap.forEach((hasLimit, taskType) => {
        if (hasLimit) {
          taskTypes.push(taskType)
        }
      })

      for (const taskType of taskTypes) {
        const item = m.get(taskType)
        if (!item) {
          result.push({limitWei: 0, task: taskType})
        } else {
          result.push({limitWei: item, task: taskType})
        }
      }
      return result
    },
    async loadSettings() {
      try {
        this.loadingError = false
        this.loading = true
        const res = await settingsService.settingsServiceGetSettings()
        this.settings = res.settings

        this.orgiSettings = Object.assign({}, res.settings)
      } catch (e) {
        this.loadingError = true
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
        let taskLimitMap = {} as object
        for (const item of this.taskGasLimits) {
          //@ts-ignore
          taskLimitMap[item.task] = item.limitWei.toString()
        }

        this.settings.taskGasLimitMap = taskLimitMap

        const res = await settingsService.settingsServiceUpdateSettings({body: {settings: this.settings}})
        this.settings = res.settings
        this.orgiSettings = Object.assign({}, this.settings)
        this.sysStore.showSnackBar("settings updated", false)
      } catch (e) {
        this.sysStore.showSnackBar("settings update error", true)
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
    this.taskGasLimits = this.GetTaskGasLimits()
    this.user = await helperService.helperServiceGetUser().catch()
  }
})


</script>


<style scoped>
.header {
  top: 0px;
  position: fixed;
  height: auto;
  min-height: 20px;
  z-index: 100;
  right: 2px;
  left: 56px;
  background-color: white;;
  box-shadow: rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;
}

</style>

