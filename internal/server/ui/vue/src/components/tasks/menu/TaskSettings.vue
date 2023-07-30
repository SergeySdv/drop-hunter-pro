<template>
  <v-skeleton-loader
    :loading="settingsLoading"
    type="list-item"
  >
    <v-responsive>
      <span class="mr-1"> Max tx cost:</span>
      <span v-if="errorLoading" style="color: red">{{ errorLoading }}</span>
      <span v-else>
          <span v-if="useGasLimit()">{{ getBalance(settings.gasLimit) }}</span>
          <v-dialog
            v-model="showSetTxMaxCostDialog"
            width="350px"
            persistent
          >
            <template v-slot:activator="{ props }">
              <v-btn
                class="ml-1"
                :rounded="false"
                density="compact"
                variant="outlined"
                height="20"
                width="80"
                v-bind="props"
              >
                update
              </v-btn>
            </template>


            <v-card>
              <v-card-title class="d-flex justify-space-between">
                <div>Change max tx cost</div>
                <v-icon icon="mdi-close" @click="showSetTxMaxCostDialog = false"/>
              </v-card-title>
              <v-container>
                <v-form ref="gas-limit-input">
                  <v-text-field
                    type="number"
                    variant="outlined"
                    :messages="gasLimitInputMessage"
                    density="compact"
                    label="set max transaction cost in WEI"
                    v-model="gasLimit"
                    :rules="[onlyInteger]"
                  />
                </v-form>
              </v-container>
              <v-card-actions class="d-flex justify-end">
                <v-btn @click="updateGasLimit">Update</v-btn>
              </v-card-actions>
            </v-card>

          </v-dialog>
        </span>
    </v-responsive>
  </v-skeleton-loader>
</template>

<script lang="ts">
import {AmUni, Network, SyncSwapTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ca} from "vuetify/locale";
import {helperService, processService} from "@/generated/services";
import {GetTaskSettingsResponse} from "@/generated/process";
import {mdiPencilOutline} from "@mdi/js";
import {Delay} from "@/components/helper";

export default defineComponent({
  name: "TaskSettings",
  props: {
    taskType: {
      type: Object as PropType<TaskType>,
      required: true,
    },
    network: {
      type: Object as PropType<Network>,
      required: true,
    }
  },
  watch: {
    gasLimit: {
      handler() {
        if (this.gasLimit) {
          this.gasLimitChanged()
        }
      }
    },
    showSetTxMaxCostDialog: {
      handler() {
        if (this.showSetTxMaxCostDialog) {
          this.gasLimitChanged()
        }
      }
    }
  },
  data() {
    return {
      gasLimitInputMessage: '',
      gasLimit: null as null | number,
      errorLoading: null as null | string,
      settingsLoading: false,
      settingsUpdating: false,
      showSetTxMaxCostDialog: false,
      settings: {
        gasLimit: undefined,
      } as GetTaskSettingsResponse,
      onlyInteger: (v: number) => {
        if (!this.gasLimit) {
          return 'required'
        }

        if (Number.isInteger(Number(this.gasLimit))) {
          return true
        }
        return 'number must be integer'
      },
    }
  },
  methods: {
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['gas-limit-input'].validate()

      return valid
    },
    async gasLimitChanged() {
      if (this.gasLimit && this.showSetTxMaxCostDialog) {
        await Delay(100)
        if (!await this.validateForm()) {
          return
        }

        try {
          this.gasLimitInputMessage = ''
          const res = await helperService.helperServiceCastWei({
            body: {
              wei: this.gasLimit.toString(),
              network: this.network,
            }
          })
          this.gasLimitInputMessage = 'equals to: ' + this.getBalance(res.am)
        } catch (e) {
          this.gasLimitInputMessage = 'error estimating price'
        }

      }
    },
    async updateGasLimit() {
      if (this.gasLimit) {
        this.settingsUpdating = true
        if (!await this.validateForm()) {
          this.settingsUpdating = false
          return
        }
        try {
          await processService.processServiceSetTaskSettings({
            body: {
              wei: this.gasLimit.toString(),
              taskType: this.taskType,
            }
          })
          this.showSetTxMaxCostDialog = false
          await this.getTaskSettings()
        } catch (e) {

        } finally {
          this.settingsUpdating = false
        }
      }
    },
    getBalance(am?: AmUni): string {
      if (!am) {
        return ''
      }

      return `${Number(am.usd).toPrecision(3)} USD`
      // return `${Number(am.eth).toPrecision(3)} ETH ${Number(am.usd).toPrecision(3)} USD`
    },
    useGasLimit(): boolean {
      if (!this.settings) {
        return false
      }
      return this.settings.gasLimit !== undefined
    },
    async getTaskSettings() {
      try {
        this.errorLoading = null
        this.settingsLoading = true
        this.settings = await processService.processServiceGetTaskSettings({
          body: {
            network: this.network,
            taskType: this.taskType,
          }
        })
        if (this.settings && this.settings.gasLimit) {
          this.gasLimit = Number(this.settings.gasLimit.wei)
        }

      } catch (e) {
        this.errorLoading = 'error loading settings'
      } finally {
        this.settingsLoading = false
      }
    }
  },
  async created() {
    await this.getTaskSettings()
  }
})
</script>

<style scoped>

</style>
