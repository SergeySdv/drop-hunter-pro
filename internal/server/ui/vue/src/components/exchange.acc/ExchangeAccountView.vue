<template>
  <v-container>
    <v-skeleton-loader
      width="100%"
      :loading="loading"
      type="table"
    >
      <v-responsive>
        <v-card>
          <v-card-title>
            <span class="text-h5">Exchange account options</span>
          </v-card-title>
          <v-card-text>

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
                  <div class="d-flex justify-center my-2 mb-5">
                    <v-btn width="100%" v-if="WithdrawerChanged" class="my-3" :loading="updating" @click="update">Update
                    </v-btn>
                    <div v-if="updatingError">{{ updatingError }}</div>
                  </div>
                </v-col>

              </v-row>
            </v-form>

            <v-expansion-panels v-if="withdrawer.exchangeType === ExchangeType.Okex" variant="accordion" multiple>
              <v-expansion-panel
                title="Add sub account"
              >
                <v-expansion-panel-text>
                  <CreateWithdrawerSubAcc :id="withdrawerId" @withdrawer-created="List"/>
                </v-expansion-panel-text>
              </v-expansion-panel>
              <v-expansion-panel v-for="sub in subaccs" :title="`Sub account: ${sub.label}`">
                <v-expansion-panel-text>
                  <b class="d-flex justify-center">Okex deposit sub-acc addresses</b>
                  <OkexWithdrawOptionSubAcc :id="sub.id"/>
                </v-expansion-panel-text>
              </v-expansion-panel>
            </v-expansion-panels>

          </v-card-text>
        </v-card>
      </v-responsive>
    </v-skeleton-loader>
  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {withdrawerService} from "@/generated/services"
import {
  CreateOkexWithdrawerRequest,
  ExchangeType, ExchangeWithdrawNetwork,
  ExchangeWithdrawOptions,
  Withdrawer
} from "@/generated/withdraw";
import ProxyInput from "@/components/ProxyInput.vue";
import OkexWithdrawOptionSubAcc from "@/components/exchange.acc/OkexWithdrawOptionSubAcc.vue";
import CreateWithdrawerSubAcc from "@/components/exchange.acc/CreateWithdrawerSubAcc.vue";
import {required} from "@/components/tasks/menu/helper";
import ProfileSearch from "@/components/profile/ProfileSearch.vue";
import {Profile} from "@/generated/profile";
import {Timer} from "@/components/helper";

interface Network {
  network: string
  title: string
}

interface Token {
  token: string
  title: string
}

export default defineComponent({
  name: "ExchangeAccountView",
  computed: {
    ExchangeType() {
      return ExchangeType
    },
    withdrawerId: {
      set(id: string) {
        this.$router.push({query: {wId: id, ...this.$route.query}})
      },
      get(): string {
        const s = this.$route.query.wId
        if (s !== null && !Array.isArray(s)) {
          return s
        }
        return ''
      }
    },
    WithdrawerChanged() {
      if (this.withdrawer.proxy !== this.originalWithdrawer.proxy) {
        return true
      }
      if (this.withdrawer.label !== this.originalWithdrawer.label) {
        return true
      }
      return false
    },
    subaccs(): Withdrawer[] {
      return this.withdrawers
    }
  },
  components: {
    ProfileSearch,
    CreateWithdrawerSubAcc,
    OkexWithdrawOptionSubAcc,
    ProxyInput
  },
  data() {
    return {
      loading: false,
      updatingError: "" as string | undefined,
      error: null as any,
      show: false,
      item: {} as CreateOkexWithdrawerRequest,
      saveLoading: false,
      listLoading: false,
      withdrawers: [] as Withdrawer[],
      withdrawerStatus: 'OK',
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
    required,
    async add() {
      this.show = true
      await this.List()
    },
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

    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['formm'].validate()

      return valid
    },

    async validateWithdrawForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async List() {
      try {
        this.listLoading = true
        const res = await withdrawerService.withdrawerServiceListSubWithdrawer({body: {withdrawerId: this.withdrawerId}})
        this.withdrawers = res.withdrawers
        this.$emit("change", false)
      } finally {
        this.listLoading = false
      }
    },
    Close() {
      this.show = false
    },
    async getWithdrawer() {
      const res = await withdrawerService.withdrawerServiceGetWithdrawer({body: {withdrawerId: this.withdrawerId}})
      this.withdrawer = res.withdrawer
      this.withdrawerStatus = res.error ? res.error : 'OK'
      this.originalWithdrawer = Object.assign({}, res.withdrawer)
    },
    async update() {

      try {
        this.updating = true
        this.updatingError = ''

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
    GetTokenTitle(item: ExchangeWithdrawOptions): string {
      if (item) {
        return item.token + " - " + item.amount
      }
      return ""
    },
    GetNetworkTitle(item: ExchangeWithdrawNetwork): string {
      if (item) {
        return `${item.network} min:${item.minAmount} max:${item.maxAmount} fee:${item.fee}`
      }
      return ""
    },
  },
  async created() {
    try {
      this.loading = true
      await this.getWithdrawer()
      await this.List()
    } finally {
      this.loading = false
    }

  }
})


</script>


<style>


</style>

