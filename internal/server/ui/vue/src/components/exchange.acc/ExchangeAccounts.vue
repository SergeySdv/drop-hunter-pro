<template>
  <div>
    <Loader v-if="loading"/>
    <div v-else>
      <div class="d-flex justify-end ma-3">
        <!--            <ExportExchangeAccounts class="mr-3"/>-->
        <ExchangeAccountDocs class="mr-3"/>
        <v-btn v-if="selectedSome" color="black" @click=DeleteWithdrawers class="mr-3">Delete exchange accounts
        </v-btn>
        <v-btn @click=CreateWithdrawer class="mr-3">Add exchange account</v-btn>
      </div>

      <v-table fixed-header height="90vh">
        <thead>
        <tr>
          <th></th>
          <th class="text-left">Label</th>
          <th class="text-left">Exchange</th>
          <th class="text-left">Proxy</th>
          <th class="text-left">Options</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="item in getList" :key="item.id">
          <td>
            <CheckBox style="height: 40px" density="compact" :id="item.id" focused
                      @CheckboxChanged="CheckboxChanged"></CheckBox>
          </td>
          <td>{{ item.label }}</td>
          <td>{{ item.exchangeType }}</td>
          <td>
            <BtnCheckProxy v-if="item.proxy" :proxy="item.proxy"/>
            <span v-else>{{ 'disabled' }}</span></td>
          <td>
            <v-btn @click="$router.push({name:'ExchangeAccountView', query: {wId: item.id}})">view</v-btn>
          </td>
        </tr>
        </tbody>
      </v-table>
      <CreateWithdraw :showProp="showCreateWithdrawerDialog" @change="CreateWithdrawerChange"
                      @profileCreated="UpdateList"/>
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {withdrawerService} from "@/generated/services"
import CreateWithdraw from "@/components/exchange.acc/CreateWithdraw.vue";
import CheckBox from "@/components/CheckBox.vue";
import {ExchangeType, Withdrawer} from "@/generated/withdraw";
import OkexWithdrawOption from "@/components/exchange.acc/ExchangeAccountView.vue";
import ExchangeAccountDocs from "@/components/exchange.acc/ExchangeAccountDocs.vue";
import BtnCheckProxy from "@/components/BtnCheckProxy.vue";
import ExportExchangeAccounts from "@/components/exchange.acc/ExportExchangeAccounts.vue";
import Loader from "@/components/Loader.vue";

export default defineComponent({
  name: "Withdraw",
  components: {
    Loader,
    ExportExchangeAccounts,
    BtnCheckProxy,
    ExchangeAccountDocs,
    OkexWithdrawOption,
    CheckBox,
    CreateWithdraw,
  },
  data() {
    return {
      list: [] as Withdrawer[],
      showCreateWithdrawerDialog: false,
      selected: new Set<string>(),
      loading: true,
      loadingError: false,
    }
  },
  computed: {
    ExchangeType() {
      return ExchangeType
    },
    getList(): Withdrawer[] {
      return this.list
    },
    selectedSome(): boolean {
      return this.selected.size > 0
    }
  },
  methods: {
    CheckboxChanged(id: string, value: boolean) {
      if (value) {
        this.selected.add(id)
      } else {
        this.selected.delete(id)
      }
    },
    async UpdateList() {
      try {
        this.loadingError = false
        this.loading = true
        const res = await withdrawerService.withdrawerServiceListWithdrawer()
        this.list = res.withdrawers
        this.selected = new Set<string>()
      } catch (e) {
        this.loadingError = true
      } finally {
        this.loading = false
      }

    },
    CreateWithdrawer() {
      this.showCreateWithdrawerDialog = true
    },
    CreateWithdrawerChange(b: boolean) {
      this.showCreateWithdrawerDialog = b
    },
    async DeleteWithdrawers() {
      for (let id of this.selected.keys()) {
        await withdrawerService.withdrawerServiceDeleteWithdrawer({body: {id: id}})
      }
      await this.UpdateList()
    }
  },

  created() {
    this.UpdateList()
  }
})


</script>


<style>


</style>

