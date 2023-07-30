<template>
  <div class="mx-8 my-8">
    <div class="text-h3 text-center my-5">Exchanges</div>
    <div>
      <v-table>
        <thead>
        <tr>
          <th class="text-left">Product</th>
          <th class="text-left">Status</th>
          <th class="text-left">Actions</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="item in integrated" :key="item.name">
          <td><a :href="item.link" class="d-inline-flex">
            <v-img height="20px" :src="item.img" class="mr-2"/>
            {{ item.name }}
          </a></td>
          <td>{{ item.status }}</td>
          <td>{{ item.action }}</td>
        </tr>
        </tbody>
      </v-table>
    </div>

  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {link} from "@/components/tasks/links";

enum Status {
  'integrated' = 'integrated',
  'soon' = 'soon',
  'not planned' = 'not planned',
  'no decision' = 'no decision'
}

export default defineComponent({
  name: "Exchange",
  created() {
  },
  data() {
    return {
      integrated: [
        {
          name: 'Okex',
          link: link.okx,
          status: Status.integrated,
          action: 'Full support of all withdrawals available on pff site + multi-sub-account deposit support',
          quest: '',
          img: '/icons/binance.ico'
        },
        {
          name: 'Binance',
          link: link.binance,
          status: Status.integrated,
          action: 'Full support of all withdrawals available on pff site',
          quest: '',
          img: '/icons/okx.png',
        },
      ],
    }
  }
})


</script>


<style>
</style>

