<template>
  <div class="mx-8 my-8">
    <div class="text-h3 text-center my-5">LayerZero</div>
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
      <div>
        <v-expansion-panels>
          <v-expansion-panel
            title="Мнения"
          >
            <v-expansion-panel-text>

              <div>

                Мнение от <a href="https://t.me/+oRCEEsa_d-Y4ODQy" target="_blank">ProMint</a>
                Возможные критерии дропа
                <br/>

                <ul class="mt-4">
                  Наибольшая вероятность:
                  <li> Количество транзакций или взаимодействий с контрактами L0: 10/20/50+</li>
                  <li> Суммарный объём транзакций: $1к/$10к/$50к/$100к+</li>
                  <li>Количество затронутых сетей: 3/5/8+</li>
                  <li> Количество активных месяцев: 1/3/5+</li>
                </ul>
                <br/>


                Практически на все свои аккаунты я делаю активность с объёмами, стараюсь попасть в топ Leaderboard, он
                немного кривой, но создатель хорошо проработал рейтинговую систему аккаунтов. Обязательно прогоняю пачки
                повторно: через 48 часов, каждый месяц. Сейчас подгоняю каждую неделю,так как верю что Snapshot близко
                или уже был.

                <ul class="mt-4">
                  Высокая вероятность:
                  <li>Использование официальных мостов (Stargate, BTC Bridge, Testnet Bridge, USDC Bridge, Aptos
                    Bridge);
                  </li>
                  <li>Бридж NFT (через любой сервис)</li>
                </ul>
                <br/>

                <ul class="mt-4">
                  Средняя вероятность:
                  <li>Количество взаимодействий со сторонними dApp поддерживающих L0</li>
                  <li>Ликвидность для Testnet Bridge</li>
                  <li></li>
                </ul>
                <br/>

                <ul class="mt-4">
                  Маловероятно:
                  <li>Стейкинг/холд STG</li>
                  <li>Различные роли в дискорд</li>
                  <li>Ликвидность для Stargate</li>
                  <li>Ликвидность для Stargate</li>
                </ul>
                <br/>


                Разберем пирамидальную концепцию, которая состоит из 3-ех типов аккаунтов:

                <ul class="mt-4">
                  Жирные:
                  <li>30-50 транзакций;</li>
                  <li>$50r+ суммарный объём;</li>
                  <li>10+ сетей взаимодействия;</li>
                  <li>30-50 транзакций;</li>
                  <li>Взаимодействие со всеми официальными мостами (Stargate, BTC Bridge, Testnet Bridge, USDC Bridge,
                    Aptos
                    Bridge);
                  </li>
                  <li>Бридж NFT;</li>
                  <li>После основного актива проводить по 1 транзакции раз в 1-2 недели;</li>
                  <li>Взаимодействие со сторонними dApp на L0: 10+;</li>
                  <li>Холд STG, голосования, стейкинг STG, дать ликвидность Stargate 1к USDT/USDC;</li>
                </ul>


                <ul class="mt-4">
                  Средние:
                  <li>20+ транзакций;</li>
                  <li>$10к+ суммарный объём;</li>
                  <li>6+ сетей взаимодействия, минимум 3 L1-сети;</li>
                  <li>Взаимодействие со всеми официальными мостами (Stargate, BTC Bridge, Testnet Bridge, USDC Bridge,
                    Aptos
                    Bridge);
                  </li>
                  <li>Бридж NFT;</li>
                  <li>После основного актива проводить по 1 транзакции раз в 1-2 недели;</li>
                </ul>

                <ul>
                  Сибил:
                  <li>8+ транзакций;</li>
                  <li>$500+ объём;</li>
                  <li>3+ сети взаимодействия;</li>
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
  name: "LayerZero",
  created() {
  },
  data() {
    return {
      integrated: [
        {
          name: 'Stargate.finance',
          link: taskProps.StargateBridge.service.link,
          img: taskProps.StargateBridge.service.img,
          status: Status.integrated,
          action: 'Swaps (ETH, STG, USDT, USDC)',
          quest: ''
        },
        {
          name: 'TestNetBridge',
          link: taskProps.TestNetBridgeSwap.service.link,
          img: taskProps.TestNetBridgeSwap.service.img,
          status: Status.integrated,
          action: 'Arbitrum, Optimism, Ethereum (ETH) -> Goerli (ETH)',
          quest: ''
        },
      ],
    }
  }
})


</script>


<style>
</style>

