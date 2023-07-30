<template>
  <div class="mx-8 my-8">
    <div class="text-h3 text-center my-5">Snapshot.org</div>
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
          <td><a :href="item.link">{{ item.name }}</a></td>
          <td>{{ item.status }}</td>
          <td>{{ item.action }}</td>
        </tr>
        </tbody>
      </v-table>
      <div>
        <v-expansion-panels>
          <v-expansion-panel
            title="Мнения"
          >
            <v-expansion-panel-text>
              <div>
                Некоторые дропы зависят от голосовалок, например
                <a href="https://www.bitget.com/academy/how-to-claim-mav-airdrop" target="_blank">maverick airdrop</a>
              </div>
            </v-expansion-panel-text>
          </v-expansion-panel>
        </v-expansion-panels>


      </div>

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
  name: "Snapshot",
  created() {
  },
  data() {
    return {
      integrated: [
        {
          name: 'Vote on snapshot',
          link: link.snapshot,
          status: Status.integrated,
          action: 'follow and vote automatic on any guild',
          quest: ''
        },

      ],
    }
  }
})


</script>


<style>
</style>

