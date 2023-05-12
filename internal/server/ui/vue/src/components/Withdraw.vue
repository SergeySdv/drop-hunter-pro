<template>
  <div>
    <div class="d-flex justify-end">
      <v-btn v-if="selectedSome" color="black" @click=DeleteWithdrawers class="ma-3">Delete withdraw accounts</v-btn>
      <v-btn @click=CreateWithdrawer class="ma-3">Create withdraw account</v-btn>
    </div>

    <v-table>
      <thead>
      <tr>
        <th></th>
        <th class="text-left">Label</th>
        <th class="text-left">Exchange</th>
        <th class="text-left">Proxy</th>
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
        <td>{{ item.proxy || 'disabled' }}</td>
      </tr>
      </tbody>
    </v-table>
    <CreateWithdraw :showProp="showCreateWithdrawerDialog" @change="CreateWithdrawerChange"
                    @profileCreated="UpdateList"/>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {withdrawerService} from "@/generated/services"
import CreateWithdraw from "@/components/CreateWithdraw.vue";
import CheckBox from "@/components/CheckBox.vue";
import {Withdrawer} from "@/generated/withdraw";

export default defineComponent({
  name: "Withdraw",
  components: {
    CheckBox,
    CreateWithdraw,
  },
  props: {},
  data() {
    return {
      list: [] as Withdrawer[],
      showCreateWithdrawerDialog: false,
      selected: new Set<string>(),
    }
  },
  computed: {
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
      const res = await withdrawerService.withdrawerServiceListWithdrawer()
      this.list = res.withdrawers
      this.selected = new Set<string>()
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

