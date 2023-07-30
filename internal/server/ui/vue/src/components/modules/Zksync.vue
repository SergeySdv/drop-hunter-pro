<template>
  <div class="mx-8 my-8">
    <div class="text-h3 text-center my-5">ZkSync</div>
    <div>
      <v-table>
        <thead>
        <tr>
          <th class="text-left">Product</th>
          <th class="text-left">Status</th>
          <th class="text-left">Actions</th>
          <th class="text-left">Quest</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="item in integratedZksync" :key="item.name">
          <td><a :href="item.link" class="d-inline-flex">
            <v-img height="20px" :src="item.img" class="mr-2"/>
            {{ item.name }}
          </a></td>
          <td>{{ item.status }}</td>
          <td>{{ item.action }}</td>
          <td>{{ item.quest }}</td>
        </tr>
        </tbody>
      </v-table>
      <div>
        <!--        <div v-for="item in links">-->
        <!--          <a :href="item.link">{{ item.name }}</a>-->
        <!--        </div>-->

        <v-expansion-panels>
          <v-expansion-panel
            title="Мнения"
          >
            <v-expansion-panel-text>
              links:
              <a href="https://dune.com/sixdegree/zksync-airdrop-simulation-ranking" target="_blank">Dune dashboard</a>
              <br/>
              <div>
                Мнение от <a href="https://t.me/+oRCEEsa_d-Y4ODQy" target="_blank">ProMint</a>
                <ul>
                  <li>100+ транзакций в Era</li>
                  <li>10+ транзакций в Lite</li>
                  <li>Взаимодействие с zkBridge с обеих сторон</li>
                  <li>20-30 смарт-контрактов</li>
                  <li>Баланс кошелька 0.005 ETH</li>
                  <li>Активность в каждом месяце</li>
                  <li>2-10к объём</li>
                </ul>
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
import {taskProps} from "@/components/tasks/tasks";

enum Status {
  'integrated' = 'integrated',
  'soon' = 'soon',
  'not planned' = 'not planned',
  'no decision' = 'no decision'
}

export default defineComponent({
  name: "Zksync",
  created() {
  },
  data() {
    return {
      integratedZksync: [
        {
          name: 'zk bridge',
          link: taskProps.ZkSyncOfficialBridgeFromEthereum.service.link,
          img: taskProps.ZkSyncOfficialBridgeFromEthereum.service.img,
          status: Status.integrated,
          action: 'Bridge ZkSync (ETH) <-> Ethereum (ETH)',
          quest: 'https://layer3.xyz/quests/welcome-to-a-new-era'
        },
        {
          name: 'SyncSwap',
          link: taskProps.SyncSwap.service.link,
          img: taskProps.SyncSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC, LUSD, LSD, MUTE). LP deposit/withdraw (ETH/USDC)',
          quest: 'https://layer3.xyz/quests/syncswap-on-zksync'
        },
        {
          name: 'MuteIO',
          link: taskProps.MuteioSwap.service.link,
          img: taskProps.MuteioSwap.service.img,
          status: Status.integrated,
          action: 'Swap ZkSync (ETH <-> USDC)',
          quest: ''
        },
        {
          name: 'Maverick',
          link: taskProps.MaverickSwap.service.link,
          img: taskProps.MaverickSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> MAV, LUSD, USDC)',
          quest: ''
        },
        {
          name: 'SpaceFi',
          link: taskProps.SpaceFISwap.service.link,
          img: taskProps.SpaceFISwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> SPACE, USDC)',
          quest: ''
        },
        {
          name: taskProps.VelocoreSwap.service.name,
          link: taskProps.VelocoreSwap.service.link,
          img: taskProps.VelocoreSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> VC, USDC)',
          quest: ''
        },
        {
          name: taskProps.IzumiSwap.service.name,
          link: taskProps.IzumiSwap.service.link,
          img: taskProps.IzumiSwap.service.img,
          status: Status.integrated,
          action: 'Swap (WETH/ETH <-> USDC/IZI)',
          quest: ''
        },
        {
          name: taskProps.VeSyncSwap.service.name,
          link: taskProps.VeSyncSwap.service.link,
          img: taskProps.VeSyncSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC)',
          quest: ''
        },
        {
          name: taskProps.EzkaliburSwap.service.name,
          link: taskProps.EzkaliburSwap.service.link,
          img: taskProps.EzkaliburSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC)',
          quest: ''
        },
        {
          name: taskProps.ZkSwap.service.name,
          link: taskProps.ZkSwap.service.link,
          img: taskProps.ZkSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC)',
          quest: ''
        }
      ],
      links: [
        {
          name: 'explorer',
          link: 'https://explorer.zksync.io/',
          desc: 'default explorer'
        },
        {
          name: 'zksync activity stats',
          link: 'https://byfishh.github.io/zk-flow/',
          desc: 'shows platform activity stats'
        },

      ]
    }
  }
})


</script>


<style>
</style>

