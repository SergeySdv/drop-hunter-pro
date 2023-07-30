<template>
  <v-skeleton-loader
    width="100%"
    :loading="listLoading"
    type="table"
  >
    <v-responsive>

      <v-form ref="formm">
        <v-row>
          <v-col><b>Status</b>: {{ withdrawerStatus }}</v-col>
        </v-row>
        <v-row>
          <v-col cols="12">
            <v-text-field
              ref="label-input"
              v-model="withdrawer.label"
              label="label"
              density="comfortable"
              variant="outlined"
              :rules="labelRules()"
              @input="labelChanged"
            ></v-text-field>
          </v-col>
          <v-col cols="12">
            <ProxyInput v-model="withdrawer.proxy" :required="true"/>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-btn v-if="WithdrawerChanged" class="my-3" :loading="updating" @click="update">Update</v-btn>
            <div v-if="updatingError">{{ updatingError }}</div>
          </v-col>
        </v-row>
      </v-form>

      <v-table fixed-header height="90vh">
        <thead>
        <tr>
          <th class="text-left">address</th>
          <th class="text-left">profile</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="item in items">
          <td>{{ `${item.addr}` }}</td>
          <td class="align-center" style="width: 200px">
            <OkexWithdrawProfiles :model-value="item" :withdrawer-id="id"
                                  @updated="OkexWithdrawProfileUpdated"/>
          </td>
        </tr>
        </tbody>
      </v-table>
    </v-responsive>
  </v-skeleton-loader>


</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {helperService, profileService, withdrawerService} from "@/generated/services"
import {
  CreateOkexWithdrawerRequest,
  CreateWithdrawerRequest,
  DepositAddresses,
  ExchangeType,
  Withdrawer
} from "@/generated/withdraw";
import ProxyInput from "@/components/ProxyInput.vue";
import {Profile} from "@/generated/profile";
import {Timer} from "@/components/helper";
import OkexWithdrawProfiles from "@/components/exchange.acc/OkexWithdrawProfiles.vue";

export default defineComponent({
  name: "OkexWithdrawOptionSubAcc",
  computed: {
    WithdrawerChanged() {
      if (this.withdrawer.proxy !== this.originalWithdrawer.proxy) {
        return true
      }
      if (this.withdrawer.label !== this.originalWithdrawer.label) {
        return true
      }

      return false
    },
  },
  components: {
    OkexWithdrawProfiles,
    ProxyInput
  },
  props: {
    id: {
      required: true,
      type: String
    },
  },
  data() {
    return {
      updatingError: "" as string | undefined,
      listLoading: false as boolean,
      items: [] as DepositAddresses[],
      withdrawerStatus: "",
      withdrawer: {
        label: '',
        proxy: '',
        id: ''
      } as Withdrawer,
      originalWithdrawer: {
        label: '',
        proxy: '',
        id: ''
      } as Withdrawer,
      updating: false,
    }
  },
  methods: {
    labelRules() {
      return [
        (v: string) => this.withdrawer.label ? true : 'label is required'
      ]
    },
    async labelChanged(e: InputEvent) {
      await this.validate()
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['label-input'].validate()

      return valid
    },
    async getWithdrawer() {
      const res = await withdrawerService.withdrawerServiceGetWithdrawer({body: {withdrawerId: this.id}})
      this.withdrawer = res.withdrawer
      this.withdrawerStatus = res.error ? res.error : 'OK'
      this.originalWithdrawer = Object.assign({}, res.withdrawer)
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['formm'].validate()

      return valid
    },
    async update() {

      try {
        this.updating = true
        this.updatingError = ""

        if (!await this.validateForm()) {
          this.updating = false
          return
        }

        const res = await withdrawerService.withdrawerServiceUpdateWithdrawer({
          body: {
            proxy: this.withdrawer.proxy,
            label: this.withdrawer.label,
            withdrawerId: this.withdrawer.id,
          }
        })

        this.updatingError = res.error

        await this.getWithdrawer()
      } finally {
        this.updating = false
      }
    },

    async OkexWithdrawProfileUpdated() {
      // await this.List()
    },
    async List() {
      try {
        this.listLoading = true
        const res = await withdrawerService.withdrawerServiceListDepositAddresses({body: {withdrawerId: this.id}})
        this.items = res.items.sort((a, b) => {
          const nameA = a.addr.toUpperCase(); // ignore upper and lowercase
          const nameB = b.addr.toUpperCase(); // ignore upper and lowercase
          if (nameA < nameB) {
            return -1;
          }
          if (nameA > nameB) {
            return 1;
          }

          // names must be equal
          return 0;
        })
      } finally {
        this.listLoading = false
      }
    },
  },
  async created() {
    this.List()
    this.getWithdrawer()
  }
})


</script>


<style>


</style>

